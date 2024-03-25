package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/lightsparkdev/go-sdk/utils"
	"github.com/uma-universal-money-address/uma-go-sdk/uma"
	umaprotocol "github.com/uma-universal-money-address/uma-go-sdk/uma/protocol"
	umautils "github.com/uma-universal-money-address/uma-go-sdk/uma/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// Vasp1 is an implementation of the sending VASP in the UMA protocol.
type Vasp1 struct {
	config       *UmaConfig
	pubKeyCache  uma.PublicKeyCache
	requestCache *Vasp1RequestCache
	nonceCache   uma.NonceCache
	client       *services.LightsparkClient
}

func NewVasp1(config *UmaConfig, pubKeyCache uma.PublicKeyCache) *Vasp1 {
	oneDayAgo := time.Now().AddDate(0, 0, -1)
	return &Vasp1{
		config:       config,
		pubKeyCache:  pubKeyCache,
		requestCache: NewVasp1RequestCache(),
		nonceCache:   uma.NewInMemoryNonceCache(oneDayAgo),
		client:       services.NewLightsparkClient(config.ApiClientID, config.ApiClientSecret, config.ClientBaseURL),
	}
}

func (v *Vasp1) handleClientUmaLookup(context *gin.Context) {
	receiverAddress := context.Param("receiver")
	err := ValidateUmaAddress(receiverAddress)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": "Invalid receiver address",
		})
		return
	}
	addressParts := strings.Split(receiverAddress, "@")
	receiverId := addressParts[0]
	receiverVasp := addressParts[1]
	signingKey, err := v.config.UmaSigningPrivKeyBytes()

	lnurlpRequest, err := uma.GetSignedLnurlpRequestUrl(
		signingKey, receiverAddress, v.getVaspDomain(context), true, nil)

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
			"reason": "Failed to parse lnurlp response",
		})
		return
	}
	if !lnurlpResponse.IsUmaResponse() {
		// Try to fall back to a non-UMA lnurlp response.
		v.handleNonUmaLnurlpResponse(*lnurlpResponse, receiverId, receiverVasp, context)
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

	receiverSigningPubKey, err := receivingVaspPubKey.SigningPubKey()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to get signing pub key for receiving VASP",
		})
		return
	}
	err = uma.VerifyUmaLnurlpResponseSignature(*lnurlpResponse.AsUmaResponse(), receiverSigningPubKey, v.nonceCache)
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
		"receiverCurrencies": lnurlpResponse.Currencies,
		"minSendSats":        lnurlpResponse.MinSendable,
		"maxSendSats":        lnurlpResponse.MaxSendable,
		"callbackUuid":       callbackUuid,
		"receiverKYCStatus":  lnurlpResponse.Compliance.KycStatus, // You might not actually send this to a client in practice.
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
		v.getVaspDomain(context),
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
	if umautils.IsDomainLocalhost(context.Request.Host) {
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
	if currencyCode == "" {
		currencyCode = "SAT"
	}
	currencySupported := false
	receiverCurrencies := initialRequestData.lnurlpResponse.Currencies
	if receiverCurrencies == nil {
		receiverCurrencies = &[]umaprotocol.Currency{SatsCurrency}
	}
	for i := range *receiverCurrencies {
		if (*receiverCurrencies)[i].Code == currencyCode {
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
	var payerInfo *PayerInfo
	if initialRequestData.lnurlpResponse.RequiredPayerData != nil {
		payerInfoVal := v.getPayerInfo(*initialRequestData.lnurlpResponse.RequiredPayerData, context)
		payerInfo = &payerInfoVal
	}
	isAmountInMsats := strings.ToLower(context.Query("isAmountInMsats")) == "true"
	if !initialRequestData.lnurlpResponse.IsUmaResponse() {
		isAmountInMsats = strings.ToLower(context.Query("isAmountInMsats")) != "false"
	}
	var comment *string
	if commentVal, ok := context.GetQuery("comment"); ok {
		comment = &commentVal
	}

	if !initialRequestData.lnurlpResponse.IsUmaResponse() {
		v.handleNonUmaPayReq(
			context, initialRequestData, amountInt64, callbackUuid, payerInfo, currencyCode, isAmountInMsats, comment)
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

	trInfo := "Here is some fake travel rule info. It's up to you to actually implement this."
	senderUtxos, err := v.client.GetNodeChannelUtxos(v.config.NodeUUID)
	if err != nil {
		log.Printf("Failed to get prescreening UTXOs: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to get prescreening UTXOs",
		})
		return
	}
	vasp2EncryptionPubKey, err := vasp2PubKeys.EncryptionPubKey()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to get encryption pub key for receiving VASP",
		})
		return
	}

	senderNode, err := GetNode(v.client, v.config.NodeUUID)
	if err != nil || senderNode == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to get sender node pub key",
		})
		return
	}
	umaMajorVersion := uma.MAJOR_VERSION
	if initialRequestData.lnurlpResponse.UmaVersion != nil {
		umaVersion, err := uma.ParseVersion(*initialRequestData.lnurlpResponse.UmaVersion)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"status": "ERROR",
				"reason": "Failed to parse uma version",
			})
			return
		}
		umaMajorVersion = umaVersion.Major
	}
	txID := "1234" // In practice, you'd probably use some real transaction ID here.
	// If you are using a standardized travel rule format, you can set this to something like:
	// "IVMS@101.2023".
	var trFormat *umaprotocol.TravelRuleFormat
	payReq, err := uma.GetUmaPayRequest(
		amountInt64,
		vasp2EncryptionPubKey,
		umaSigningPrivateKey,
		currencyCode,
		!isAmountInMsats,
		payerInfo.Identifier,
		umaMajorVersion,
		payerInfo.Name,
		payerInfo.Email,
		&trInfo,
		trFormat,
		umaprotocol.KycStatusVerified,
		&senderUtxos,
		(*senderNode).GetPublicKey(),
		v.getUtxoCallback(context, txID),
		&umaprotocol.CounterPartyDataOptions{
			umaprotocol.CounterPartyDataFieldName.String():  {Mandatory: false},
			umaprotocol.CounterPartyDataFieldEmail.String(): {Mandatory: false},
			// Compliance and Identifier are mandatory fields added automatically.
		},
		nil,
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
	compliance, err := payreqResponse.PayeeData.Compliance()
	var utxoCallback *string
	if compliance != nil && compliance.UtxoCallback != nil && *compliance.UtxoCallback != "" {
		utxoCallback = compliance.UtxoCallback
	}
	v.requestCache.SavePayReqData(
		callbackUuid,
		payreqResponse.EncodedInvoice,
		utxoCallback,
		&invoiceData,
	)

	context.JSON(http.StatusOK, gin.H{
		"encodedInvoice":          payreqResponse.EncodedInvoice,
		"callbackUuid":            callbackUuid,
		"amountMsats":             invoiceData.Amount,
		"amountReceivingCurrency": payreqResponse.PaymentInfo.Amount,
		"conversionRate":          payreqResponse.PaymentInfo.Multiplier,
		"currencyCode":            payreqResponse.PaymentInfo.CurrencyCode,
		"exchangeFeesMsats":       payreqResponse.PaymentInfo.ExchangeFeesMillisatoshi,
		"expiresAt":               invoiceData.ExpiresAt.Unix(),
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
	var utxosWithAmounts []umaprotocol.UtxoWithAmount
	for _, postTransactionData := range *payment.UmaPostTransactionData {
		amountMilliSatoshi, err := utils.ValueMilliSatoshi(postTransactionData.Amount)
		if err != nil {
			continue
		}
		utxosWithAmounts = append(utxosWithAmounts, umaprotocol.UtxoWithAmount{
			Utxo:   postTransactionData.Utxo,
			Amount: amountMilliSatoshi,
		})
	}
	if err != nil {
		log.Fatalf("Failed to marshal UTXOs: %v", err)
	} else if payReqData.utxoCallback != nil {
		log.Printf("Sending UTXOs to %s: %s", *payReqData.utxoCallback, utxosWithAmounts)
		signingPrivateKey, err := v.config.UmaSigningPrivKeyBytes()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"status": "ERROR",
				"reason": err.Error(),
			})
			return
		}
		postTxCallback, err := uma.GetPostTransactionCallback(utxosWithAmounts, v.getVaspDomain(context), signingPrivateKey)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"status": "ERROR",
				"reason": err.Error(),
			})
			return
		}
		requestBody, err := json.Marshal(postTxCallback)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"status": "ERROR",
				"reason": err.Error(),
			})
			return
		}
		utxoCallbackResponse, err := http.Post(
			*payReqData.utxoCallback,
			"application/json",
			bytes.NewBuffer(requestBody),
		)
		if err != nil {
			log.Printf("Failed to send UTXOs to receiver: %v", err)
		} else if utxoCallbackResponse.StatusCode != http.StatusOK {
			log.Printf("Failed to send UTXOs to receiver: %d", utxoCallbackResponse.StatusCode)
		} else {
			log.Printf("Sent UTXOs to receiver")
		}
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
	response := umaprotocol.PubKeyResponse{
		SigningPubKeyHex:    hex.EncodeToString(signingPubKeyBytes),
		EncryptionPubKeyHex: hex.EncodeToString(encryptionPubKeyBytes),
		ExpirationTimestamp: &twoWeeksFromNowSec,
	}

	context.JSON(http.StatusOK, response)
}

func (v *Vasp1) handleNonUmaLnurlpResponse(
	lnurlpResponse umaprotocol.LnurlpResponse, receiverId string, receiverDomain string, context *gin.Context) {
	callbackUuid := v.requestCache.SaveLnurlpResponseData(lnurlpResponse, receiverId, receiverDomain)
	var serializedCurrencies = []byte("[]")
	if lnurlpResponse.Currencies != nil && len(*lnurlpResponse.Currencies) == 0 {
		var err error
		serializedCurrencies, err = json.Marshal(lnurlpResponse.Currencies)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"status": "ERROR",
				"reason": "Failed to serialize currencies",
			})
			return
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"receiverCurrencies": serializedCurrencies,
		"callbackUuid":       callbackUuid,
		"maxSendSats":        lnurlpResponse.MaxSendable,
		"minSendSats":        lnurlpResponse.MinSendable,
		"receiverKYCStatus":  umaprotocol.KycStatusNotVerified,
	})
}

func (v *Vasp1) handleNonUmaPayReq(
	context *gin.Context,
	initialRequestData Vasp1InitialRequestData,
	amountInt64 int64,
	callbackUuid string,
	payerInfo *PayerInfo,
	currencyCode string,
	isAmountInMsats bool,
	comment *string,
) {
	callbackUrl, err := url.Parse(initialRequestData.lnurlpResponse.Callback)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to parse callback URL",
		})
		return
	}
	var payerData *umaprotocol.PayerData
	if payerInfo != nil {
		payerData = &umaprotocol.PayerData{
			umaprotocol.CounterPartyDataFieldName.String():       payerInfo.Name,
			umaprotocol.CounterPartyDataFieldEmail.String():      payerInfo.Email,
			umaprotocol.CounterPartyDataFieldIdentifier.String(): payerInfo.Identifier,
		}
	}
	var sendingAmountCurrencyCode *string
	if !isAmountInMsats {
		*sendingAmountCurrencyCode = currencyCode
	}
	payreq := umaprotocol.PayRequest{
		SendingAmountCurrencyCode: sendingAmountCurrencyCode,
		ReceivingCurrencyCode:     &currencyCode,
		Amount:                    amountInt64,
		PayerData:                 payerData,
		RequestedPayeeData:        nil,
		Comment:                   comment,
	}

	payreqParams, err := payreq.EncodeAsUrlParams()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to encode payreq as URL params",
		})
		return
	}
	queryParams := callbackUrl.Query()
	for key, values := range *payreqParams {
		for _, value := range values {
			queryParams.Add(key, value)
		}
	}
	callbackUrl.RawQuery = queryParams.Encode()

	payreqResult, err := http.Get(callbackUrl.String())
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

	var payreqResponse umaprotocol.PayReqResponse
	err = json.Unmarshal(payreqResultBytes, &payreqResponse)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Failed to unmarshal payreq response",
		})
		return
	}

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
		nil,
		&invoiceData,
	)

	resp := gin.H{
		"encodedInvoice": payreqResponse.EncodedInvoice,
		"callbackUuid":   callbackUuid,
		"amountMsats":    invoiceData.Amount,
		"currencyCode":   currencyCode,
		"expiresAt":      invoiceData.ExpiresAt.Unix(),
	}
	if payreqResponse.PaymentInfo != nil {
		resp["amountReceivingCurrency"] = payreqResponse.PaymentInfo.Amount
		resp["conversionRate"] = payreqResponse.PaymentInfo.Multiplier
		resp["exchangeFeesMsats"] = payreqResponse.PaymentInfo.ExchangeFeesMillisatoshi
		resp["receivingCurrencyDecimals"] = payreqResponse.PaymentInfo.Decimals
	}
	context.JSON(http.StatusOK, resp)
}

func (v *Vasp1) getVaspDomain(context *gin.Context) string {
	envVaspDomain := v.config.OwnVaspDomain
	if envVaspDomain != "" {
		return envVaspDomain
	}
	requestHost := context.Request.Host
	requestHostWithoutPort := strings.Split(requestHost, ":")[0]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	if port != "80" && port != "443" {
		return fmt.Sprintf("%s:%s", requestHostWithoutPort, port)
	}
	return requestHostWithoutPort
}

type PayerInfo struct {
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Identifier string  `json:"identifier"`
}

// NOTE: In a real application, you'd want to use the authentication context to pull out this information. It's not actually
// always Alice sending the money ;-).
func (v *Vasp1) getPayerInfo(options umaprotocol.CounterPartyDataOptions, context *gin.Context) PayerInfo {
	var name string
	if options[umaprotocol.CounterPartyDataFieldName.String()].Mandatory {
		name = v.config.Username
	}
	var email string
	if options[umaprotocol.CounterPartyDataFieldEmail.String()].Mandatory {
		email = v.config.Username + "@" + v.getVaspDomain(context)
	}
	return PayerInfo{
		Name:       &name,
		Email:      &email,
		Identifier: "$" + v.config.Username + "@" + v.getVaspDomain(context),
	}
}
