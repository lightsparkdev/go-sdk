package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/lightsparkdev/go-sdk/uma"
	"github.com/lightsparkdev/go-sdk/utils"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Vasp1 is an implementation of the sending VASP in the UMA protocol.
type Vasp1 struct {
	config       *UmaConfig
	pubKeyCache  uma.PublicKeyCache
	requestCache *Vasp1RequestCache
	client       *services.LightsparkClient
}

func NewVasp1(config *UmaConfig, pubKeyCache uma.PublicKeyCache) *Vasp1 {
	return &Vasp1{
		config:       config,
		pubKeyCache:  pubKeyCache,
		requestCache: NewVasp1RequestCache(),
		client:       services.NewLightsparkClient(config.ApiClientID, config.ApiClientSecret, &config.ClientBaseURL),
	}
}

func (v *Vasp1) handleClientUmaLookup(context *gin.Context) {
	receiverAddress := context.Param("receiver")
	addressParts := strings.Split(receiverAddress, "@")
	if len(addressParts) != 2 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": "Invalid receiver address",
		})
		return
	}
	receiverId := addressParts[0]
	receiverVasp := addressParts[1]
	signingKey, err := v.config.UmaSigningPrivKeyBytes()
	lnurlpRequest, err := uma.GetSignedLnurlpRequestUrl(signingKey, receiverAddress, "localhost:8080", true, nil)

	resp, err := http.Get(lnurlpRequest.String())
	if err != nil {
		// TODO: Maybe this should be a 400 depending on the error?
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to fetch receiver's lnurlp",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusPreconditionFailed {
		retryResp, didSendFailure := v.handleUnsupportedVersionResponse(resp, signingKey, receiverAddress, context)
		defer retryResp.Body.Close()
		if didSendFailure {
			return
		}
		resp = retryResp
	}

	if resp.StatusCode != http.StatusOK {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Failed response from receiver: %d", resp.StatusCode),
		})
		return
	}

	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to read lnurlp response from receiver",
		})
		return
	}
	lnurlpResponse, err := uma.ParseLnurlpResponse(responseBodyBytes)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to parse lnurlp response from receiver",
		})
		return
	}

	receivingVaspPubKey, err := uma.FetchPublicKeyForVasp(receiverVasp, v.pubKeyCache)
	if err != nil || receivingVaspPubKey == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to fetch public key for receiving VASP",
		})
		return
	}

	err = uma.VerifyUmaLnurlpResponseSignature(lnurlpResponse, receivingVaspPubKey.SigningPubKey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to verify lnurlp response signature from receiver",
		})
		return
	}

	callbackUuid := v.requestCache.SaveLnurlpResponseData(*lnurlpResponse, receiverId, receiverVasp)

	// In practice a VASP might want to check lnurlpResponse.Compliance for info on travel rule status, kyc state, etc. We don't
	// have anything to do with that data in this demo, though.

	context.JSON(http.StatusOK, gin.H{
		"currencies":        lnurlpResponse.Currencies,
		"minSendSats":       lnurlpResponse.MinSendable,
		"maxSendSats":       lnurlpResponse.MaxSendable,
		"callbackUuid":      callbackUuid,
		"receiverKYCStatus": lnurlpResponse.Compliance.KycStatus, // You might not actually send this to a client in practice.
	})
}

func (v *Vasp1) handleUnsupportedVersionResponse(response *http.Response, signingKey []byte, receiverAddress string, context *gin.Context) (*http.Response, bool) {
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to read lnurlp response from receiver",
		})
		return nil, true
	}
	supportedMajorVersions, err := uma.GetSupportedMajorVersionsFromErrorResponseBody(responseBodyBytes)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to parse supported major versions from error response",
		})
		return nil, true
	}
	highestSupportedVersion := uma.SelectHighestSupportedVersion(supportedMajorVersions)
	if highestSupportedVersion == nil {
		context.JSON(http.StatusPreconditionFailed, gin.H{
			"status": "ERROR",
			"reason": "No compatible UMA version with VASP2",
		})
		return nil, true
	}
	lnurlpRequest, err := uma.GetSignedLnurlpRequestUrl(
		signingKey,
		receiverAddress,
		"localhost:8080",
		true,
		highestSupportedVersion,
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to generate lnurlp request",
		})
		return nil, true
	}

	retryResponse, err := http.Get(lnurlpRequest.String())
	if err != nil {
		// TODO: Maybe this should be a 400 depending on the error?
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to fetch receiver's lnurlp",
		})
		return nil, true
	}
	return retryResponse, false
}

func (v *Vasp1) getUtxoCallback(context *gin.Context, txId string) string {
	scheme := "https://"
	if strings.HasPrefix(context.Request.Host, "localhost:") {
		scheme = "http://"
	}
	return fmt.Sprintf("%s%s/api/uma/utxocallback?txid=%s", scheme, context.Request.Host, txId)
}

func (v *Vasp1) handleClientPayReq(context *gin.Context) {
	callbackUuid := context.Param("callbackUuid")
	initialRequestData, ok := v.requestCache.GetLnurlpResponseData(callbackUuid)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": "Invalid or missing callback UUID",
		})
		return
	}

	amount := context.Query("amount")
	amountInt64, err := strconv.ParseInt(amount, 10, 64)
	if err != nil || amountInt64 <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": "Invalid amount",
		})
		return
	}

	currencyCode := context.Query("currencyCode")
	currencySupported := false
	for i := range initialRequestData.lnurlpResponse.Currencies {
		if initialRequestData.lnurlpResponse.Currencies[i].Code == currencyCode {
			currencySupported = true
			break
		}
	}
	if !currencySupported {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": "Unsupported currency",
		})
		return
	}

	umaSigningPrivateKey, err := v.config.UmaSigningPrivKeyBytes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	vasp2PubKeys, err := uma.FetchPublicKeyForVasp(initialRequestData.vasp2Domain, v.pubKeyCache)
	if err != nil || vasp2PubKeys == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to fetch public key for receiving VASP",
		})
		return
	}

	payerInfo := getPayerInfo(initialRequestData.lnurlpResponse.RequiredPayerData)
	trInfo := "Here is some fake travel rule info. It's up to you to actually implement this."
	senderUtxos, err := v.client.GetNodeChannelUtxos(v.config.NodeUUID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to get prescreening UTXOs",
		})
		return
	}
	// This is the node pub key of the sender's node. In practice, you'd want to get this from the sender's node.
	senderNodePubKey := "abcdef12345"
	txID := "1234" // In practice, you'd probably use some real transaction ID here.
	payReq, err := uma.GetPayRequest(
		vasp2PubKeys.EncryptionPubKey,
		umaSigningPrivateKey,
		currencyCode,
		amountInt64,
		payerInfo.Identifier,
		payerInfo.Name,
		payerInfo.Email,
		&trInfo,
		uma.KycStatusVerified,
		&senderUtxos,
		&senderNodePubKey,
		v.getUtxoCallback(context, txID),
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to generate payreq",
		})
		return
	}

	payReqBytes, err := json.Marshal(payReq)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to marshal payreq",
		})
		return
	}
	payreqResult, err := http.Post(initialRequestData.lnurlpResponse.Callback, "application/json", bytes.NewBuffer(payReqBytes))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	if payreqResult.StatusCode != http.StatusOK {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Failed payreq response: %d", payreqResult.StatusCode),
		})
		return
	}

	defer payreqResult.Body.Close()

	payreqResultBytes, err := io.ReadAll(payreqResult.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to read payreq response",
		})
		return
	}

	payreqResponse, err := uma.ParsePayReqResponse(payreqResultBytes)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to unmarshal payreq response",
		})
		return
	}

	// TODO: Pre-screen the UTXOs from payreqResponse.Compliance.Utxos

	log.Printf("Received invoice: %s", payreqResponse.EncodedInvoice)
	invoice, err := v.client.DecodePaymentRequest(payreqResponse.EncodedInvoice)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to decode invoice",
		})
		return
	}
	invoiceData := (*invoice).(objects.InvoiceData)
	v.requestCache.SavePayReqData(
		callbackUuid,
		payreqResponse.EncodedInvoice,
		payreqResponse.Compliance.UtxoCallback,
		&invoiceData,
	)

	context.JSON(http.StatusOK, gin.H{
		"encodedInvoice": payreqResponse.EncodedInvoice,
		"callbackUuid":   callbackUuid,
		"amount":         invoiceData.Amount,
		"conversionRate": payreqResponse.PaymentInfo.Multiplier,
		"currencyCode":   payreqResponse.PaymentInfo.CurrencyCode,
		"expiresAt":      invoiceData.ExpiresAt.Unix(),
	})
}

func (v *Vasp1) handleClientPaymentConfirm(context *gin.Context) {
	// NOTE: In a real application, you'd want this request to be authenticated so that only the right user can confirm
	// the payment.
	callbackUuid := context.Param("callbackUuid")
	payReqData, ok := v.requestCache.GetPayReqData(callbackUuid)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": "Invalid or missing callback UUID",
		})
		return
	}

	if payReqData.invoiceData.Amount.OriginalValue == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": "cannot pay zero-amount invoices via UMA",
		})
		return
	}
	seedBytes, err := v.config.NodeMasterSeedBytes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}
	v.client.LoadNodeSigningKey(
		v.config.NodeUUID,
		// Switch this to BitcoinNetworkMainnet if you're testing with a mainnet node:
		*services.NewSigningKeyLoaderFromSignerMasterSeed(seedBytes, objects.BitcoinNetworkRegtest))

	payment, err := v.client.PayUmaInvoice(
		v.config.NodeUUID,
		payReqData.encodedInvoice,
		60,
		1000000,
		nil,
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to pay invoice",
		})
		return
	}
	payment, err = v.waitForPaymentCompletion(payment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed while waiting for payment completion",
		})
		return
	}

	log.Printf("Payment %s completed: %s", payment.Id, payment.Status.StringValue())

	// In practice, you'd want to send the UTXOs to the receiver's UTXO callback URL here.
	// For this demo, we'll just log them.
	var utxosWithAmounts []uma.UtxoWithAmount
	for _, postTransactionData := range *payment.UmaPostTransactionData {
		amountMilliSatoshi, err := utils.ValueMilliSatoshi(postTransactionData.Amount)
		if err != nil {
			continue
		}
		utxosWithAmounts = append(utxosWithAmounts, uma.UtxoWithAmount{
			Utxo:   postTransactionData.Utxo,
			Amount: amountMilliSatoshi,
		})
	}
	utxosWithAmountsBytes, err := json.Marshal(utxosWithAmounts)
	if err == nil {
		log.Printf("Sending UTXOs to %s: %s", payReqData.utxoCallback, utxosWithAmountsBytes)
	}

	context.JSON(http.StatusOK, gin.H{
		"didSucceed": payment.Status == objects.TransactionStatusSuccess,
		"paymentId":  payment.Id,
	})
}

func (v *Vasp1) waitForPaymentCompletion(payment *objects.OutgoingPayment) (*objects.OutgoingPayment, error) {
	attemptLimit := 200
	for payment.Status != objects.TransactionStatusSuccess && payment.Status != objects.TransactionStatusFailed {
		if attemptLimit == 0 {
			return nil, errors.New("payment timed out")
		}
		attemptLimit--
		time.Sleep(100 * time.Millisecond)

		entity, err := v.client.GetEntity(payment.Id)
		if err != nil {
			return nil, err
		}
		castPayment, didCast := (*entity).(objects.OutgoingPayment)
		if !didCast {
			return nil, errors.New("failed to cast payment to OutgoingPayment")
		}
		payment = &castPayment
	}
	return payment, nil
}

func (v *Vasp1) handlePubKeyRequest(context *gin.Context) {
	signingPubKeyBytes, err := v.config.UmaSigningPubKeyBytes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}
	encryptionPubKeyBytes, err := v.config.UmaEncryptionPubKeyBytes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	twoWeeksFromNow := time.Now().AddDate(0, 0, 14)
	twoWeeksFromNowSec := twoWeeksFromNow.Unix()
	response := uma.PubKeyResponse{
		SigningPubKey:       signingPubKeyBytes,
		EncryptionPubKey:    encryptionPubKeyBytes,
		ExpirationTimestamp: &twoWeeksFromNowSec,
	}

	context.JSON(http.StatusOK, response)
}

type PayerInfo struct {
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Identifier string  `json:"identifier"`
}

// NOTE: In a real application, you'd want to use the authentication context to pull out this information. It's not actually
// always Alice sending the money ;-).
func getPayerInfo(options uma.PayerDataOptions) PayerInfo {
	var name string
	if options.NameRequired {
		name = "Alice FakeName"
	}
	var email string
	if options.EmailRequired {
		email = "$alice@vasp1.com"
	}
	return PayerInfo{
		Name:       &name,
		Email:      &email,
		Identifier: "$alice@vasp1.com",
	}
}
