package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/uma-universal-money-address/uma-go-sdk/uma"
	"github.com/uma-universal-money-address/uma-go-sdk/uma/protocol"
	umaprotocol "github.com/uma-universal-money-address/uma-go-sdk/uma/protocol"
	umautils "github.com/uma-universal-money-address/uma-go-sdk/uma/utils"
)

// Vasp2 is an implementation of the receiving VASP in the UMA protocol.
type Vasp2 struct {
	config      *UmaConfig
	pubKeyCache uma.PublicKeyCache
	nonceCache  uma.NonceCache
}

// Note: In a real application, this exchange rate would come from some real oracle.
const MillisatoshiPerUsd = 22883.56

func (v *Vasp2) getLnurlpCallback(context *gin.Context) string {
	scheme := "https://"
	if umautils.IsDomainLocalhost(context.Request.Host) {
		scheme = "http://"
	}
	return fmt.Sprintf("%s%s/api/uma/payreq/%s", scheme, context.Request.Host, v.config.UserID)
}

func (v *Vasp2) getUtxoCallback(context *gin.Context, txId string) string {
	scheme := "https://"
	if umautils.IsDomainLocalhost(context.Request.Host) {
		scheme = "http://"
	}
	return fmt.Sprintf("%s%s/api/uma/utxocallback?txid=%s", scheme, context.Request.Host, txId)
}

func (v *Vasp2) getMetadata() (string, error) {
	metadata := [][]string{
		{"text/plain", fmt.Sprintf("Pay to domain.org user %s", v.config.Username)},
		{"text/identifier", fmt.Sprintf("%s@domain.org", v.config.Username)},
	}

	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		return "", err
	}

	return string(jsonMetadata), nil
}

func (v *Vasp2) handleWellKnownLnurlp(context *gin.Context) {
	username := context.Param("username")

	// Allow with or without the $ for LNURL fallback.
	if username != v.config.Username && username != "$"+v.config.Username {
		context.JSON(http.StatusNotFound, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("User not found: %s", username),
		})
		return
	}

	requestUrl := context.Request.URL
	requestUrl.Host = context.Request.Host

	lnurlpRequest, err := uma.ParseLnurlpRequest(*requestUrl)
	if err != nil {
		var unsupportedVersionErr *uma.UnsupportedVersionError
		if errors.As(err, &unsupportedVersionErr) {
			context.JSON(http.StatusPreconditionFailed, gin.H{
				"status":                 "ERROR",
				"reason":                 fmt.Sprintf("Unsupported version: %s", unsupportedVersionErr.UnsupportedVersion),
				"supportedMajorVersions": unsupportedVersionErr.SupportedMajorVersions,
				"unsupportedVersion":     unsupportedVersionErr.UnsupportedVersion,
			})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	umaLnurlpRequest := lnurlpRequest.AsUmaRequest()
	// Fallback to regular LNURL if the request is not a UMA request.
	if umaLnurlpRequest == nil {
		v.handleNonUmaLnurlRequest(context, *lnurlpRequest)
		return
	}

	lnurlpResponse, hadError := v.handleUmaQueryData(context, *umaLnurlpRequest)
	if hadError {
		return
	}
	context.JSON(http.StatusOK, lnurlpResponse)
}

func (v *Vasp2) handleNonUmaLnurlRequest(context *gin.Context, lnurlpRequest umaprotocol.LnurlpRequest) {
	callback := v.getLnurlpCallback(context)
	metadata, err := v.getMetadata()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}
	response, err := uma.GetLnurlpResponse(
		lnurlpRequest,
		callback,
		metadata,
		1,
		100_000_000,
		nil,
		nil,
		nil,
		&[]umaprotocol.Currency{
			UsdCurrency,
			SatsCurrency,
		},
		nil,
		nil,
		nil,
	)

	context.JSON(http.StatusOK, response)
}

func (v *Vasp2) handleUmaQueryData(context *gin.Context, lnurlpRequest umaprotocol.UmaLnurlpRequest) (*umaprotocol.LnurlpResponse, bool) {
	vaspDomainValidationErr := ValidateDomain(lnurlpRequest.VaspDomain)
	if vaspDomainValidationErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid sending VASP domain: %v", vaspDomainValidationErr),
		})
		return nil, true
	}
	pubKeys, err := uma.FetchPublicKeyForVasp(lnurlpRequest.VaspDomain, v.pubKeyCache)
	if err != nil || pubKeys == nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}

	if err := uma.VerifyUmaLnurlpQuerySignature(lnurlpRequest, *pubKeys, v.nonceCache); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}

	metadata, err := v.getMetadata()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}

	umaPrivateKey, err := v.config.UmaSigningPrivKeyBytes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}

	isSubjectToTravelRule := true
	kycStatus := umaprotocol.KycStatusVerified
	signedResponse, err := uma.GetLnurlpResponse(
		lnurlpRequest.LnurlpRequest,
		v.getLnurlpCallback(context),
		metadata,
		1,
		100_000_000,
		&umaPrivateKey,
		&isSubjectToTravelRule,
		&umaprotocol.CounterPartyDataOptions{
			umaprotocol.CounterPartyDataFieldIdentifier.String(): {Mandatory: true},
			umaprotocol.CounterPartyDataFieldCompliance.String(): {Mandatory: true},
			umaprotocol.CounterPartyDataFieldName.String():       {Mandatory: false},
			umaprotocol.CounterPartyDataFieldEmail.String():      {Mandatory: false},
		},
		&[]umaprotocol.Currency{
			UsdCurrency,
			SatsCurrency,
		},
		&kycStatus,
		nil,
		nil,
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}
	return signedResponse, false
}

// This is the handler for regular (non-UMA) LNURL payreq requests when the request is a GET.
func (v *Vasp2) handleLnurlPayreq(context *gin.Context) {
	uuid := context.Param("uuid")

	if uuid != v.config.UserID {
		context.JSON(http.StatusNotFound, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("User not found: %s", uuid),
		})
		return
	}

	payreq, err := umaprotocol.ParsePayRequestFromQueryParams(context.Request.URL.Query())
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid request: %v", err),
		})
		return
	}

	metadata, err := v.getMetadata()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	lsClient := services.NewLightsparkClient(v.config.ApiClientID, v.config.ApiClientSecret, v.config.ClientBaseURL)
	expirySecs := int32(600) // Expire in 10 minutes
	invoiceCreator := LightsparkClientLnurlInvoiceCreator{
		LightsparkClient: *lsClient,
		NodeId:           v.config.NodeUUID,
		ExpirySecs:       &expirySecs,
	}

	conversionRate := 1000.0
	decimals := 0
	if *payreq.ReceivingCurrencyCode == "USD" {
		conversionRate = MillisatoshiPerUsd
		decimals = 2
	}
	exchangeFees := int64(0)

	payreqResponse, err := uma.GetPayReqResponse(
		*payreq,
		invoiceCreator,
		metadata,
		payreq.ReceivingCurrencyCode,
		&decimals,
		&conversionRate,
		&exchangeFees,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	context.JSON(http.StatusOK, payreqResponse)
}

func (v *Vasp2) handleUmaPayreq(context *gin.Context) {
	uuid := context.Param("uuid")

	if uuid != v.config.UserID {
		context.JSON(http.StatusNotFound, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("User not found: %s", uuid),
		})
		return
	}

	requestBody, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid request body: %v", err),
		})
		return
	}
	request, err := uma.ParsePayRequest(requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid request body: %v", err),
		})
		return
	}
	if !request.IsUmaRequest() {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": "Invalid request body: not a UMA request.",
		})
		return
	}

	sendingVaspDomain, err := uma.GetVaspDomainFromUmaAddress(*request.PayerData.Identifier())
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid sender indentifier. UMA address required in the format $alice@vasp.com: %v", err),
		})
		return
	}
	addressValidationError := ValidateUmaAddress(*request.PayerData.Identifier())
	if addressValidationError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid sender indentifier. UMA address required in the format $alice@vasp.com: %v", addressValidationError),
		})
		return
	}

	pubKeys, err := uma.FetchPublicKeyForVasp(sendingVaspDomain, v.pubKeyCache)
	if err != nil || pubKeys == nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	if err := uma.VerifyPayReqSignature(request, *pubKeys, v.nonceCache); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	metadata, err := v.getMetadata()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	lsClient := services.NewLightsparkClient(v.config.ApiClientID, v.config.ApiClientSecret, v.config.ClientBaseURL)
	expirySecs := int32(600) // Expire in 10 minutes
	invoiceCreator := LightsparkClientUmaInvoiceCreator{
		LightsparkClient: *lsClient,
		NodeId:           v.config.NodeUUID,
		ExpirySecs:       &expirySecs,
	}

	conversionRate := 1000.0
	if *request.ReceivingCurrencyCode == "USD" {
		conversionRate = MillisatoshiPerUsd
	}
	exchangeFees := int64(100_000)
	txID := "1234" // In practice, you'd probably use some real transaction ID here.
	receiverUtxos, err := lsClient.GetNodeChannelUtxos(v.config.NodeUUID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Error getting pre-screening utxos: %v", err),
		})
		return
	}

	receiverNode, err := GetNode(lsClient, v.config.NodeUUID)
	if err != nil || receiverNode == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Error getting receiver node: %v", err),
		})
		return
	}

	decimals := 0
	if *request.ReceivingCurrencyCode == "USD" {
		decimals = 2
	}
	receiverUma := "$" + v.config.Username + "@" + v.getVaspDomain(context)
	signingKey, err := v.config.UmaSigningPrivKeyBytes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}
	payeeInfo := v.getPayeeInfo(request.RequestedPayeeData, context)
	utxoCallback := v.getUtxoCallback(context, txID)
	response, err := uma.GetPayReqResponse(
		*request,
		invoiceCreator,
		metadata,
		request.ReceivingCurrencyCode,
		&decimals,
		&conversionRate,
		&exchangeFees,
		&receiverUtxos,
		(*receiverNode).GetPublicKey(),
		&utxoCallback,
		&umaprotocol.PayeeData{
			umaprotocol.CounterPartyDataFieldIdentifier.String(): payeeInfo.Identifier,
			umaprotocol.CounterPartyDataFieldName.String():       payeeInfo.Name,
			umaprotocol.CounterPartyDataFieldEmail.String():      payeeInfo.Email,
		},
		&signingKey,
		&receiverUma,
		nil,
		nil,
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
	}

	context.JSON(http.StatusOK, response)
}

func (v *Vasp2) handlePubKeyRequest(context *gin.Context) {
	twoWeeksFromNow := time.Now().AddDate(0, 0, 14)
	twoWeeksFromNowSec := twoWeeksFromNow.Unix()
	response, err := uma.GetPubKeyResponse(v.config.UmaSigningCertChain, v.config.UmaEncryptionCertChain, &twoWeeksFromNowSec)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, response)
}

func (v *Vasp2) handleUtxoCallback(context *gin.Context) {
	txId := context.Query("txid")
	if txId == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": "Missing txid query parameter",
		})
		return
	}

	requestBody, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid request body: %v", err),
		})
		return
	}
	callbackData, err := uma.ParsePostTransactionCallback(requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid request body: %v", err),
		})
		return
	}

	log.Info("Received UTXO callback", "txId", txId, "callbackData", callbackData)

	context.Status(http.StatusOK)
}

func (v *Vasp2) handleCreateInvoice(context *gin.Context) {
	uuid := context.Param("uuid")
	if uuid != v.config.UserID {
		context.JSON(http.StatusNotFound, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("User not found: %s", uuid),
		})
		return
	}

	var requestBody struct {
		Amount       int64  `json:"amount"`
		CurrencyCode string `json:"currency_code"`
	}
	if err := context.BindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid request body: %v", err),
		})
		return
	}

	if requestBody.CurrencyCode == "" {
		requestBody.CurrencyCode = "SAT"
	}

	receiverCurrencies := []umaprotocol.Currency{}
	currencies := []umaprotocol.Currency{UsdCurrency, SatsCurrency}
	for _, currency := range currencies {
		if currency.Code == requestBody.CurrencyCode {
			receiverCurrencies = append(receiverCurrencies, currency)
		}
	}
	if len(receiverCurrencies) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("User does not support currency %s", requestBody.CurrencyCode),
		})
		return
	}
	currency := receiverCurrencies[0]

	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	callback := v.getLnurlpCallback(context)

	payerDataOptions := umaprotocol.CounterPartyDataOptions{
		umaprotocol.CounterPartyDataFieldName.String():       {Mandatory: false},
		umaprotocol.CounterPartyDataFieldEmail.String():      {Mandatory: false},
		umaprotocol.CounterPartyDataFieldIdentifier.String(): {Mandatory: true},
		umaprotocol.CounterPartyDataFieldCompliance.String(): {Mandatory: true},
	}

	privateKey, err := v.config.UmaSigningPrivKeyBytes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Failed to get UMA signing private key: %v", err),
		})
		return
	}

	invoice, err := uma.CreateUmaInvoice(
		"$" + v.config.Username+"@"+v.getVaspDomain(context),
		uint64(requestBody.Amount),
		protocol.InvoiceCurrency{
			Code:     currency.Code,
			Decimals: uint8(currency.Decimals),
			Symbol: currency.Symbol,
			Name: currency.Name,
		},
		uint64(twoDaysFromNow.Unix()),
		callback,
		true,
		&payerDataOptions,
		nil,
		nil,
		nil,
		nil,
		privateKey,
	)

	invoiceString, err := invoice.ToBech32String()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Failed to convert invoice to bech32 string: %v", err),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"invoice": invoiceString,
	})
}

func (v *Vasp2) getVaspDomain(context *gin.Context) string {
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

type PayeeInfo struct {
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Identifier string  `json:"identifier"`
}

func (v *Vasp2) getPayeeInfo(options *umaprotocol.CounterPartyDataOptions, context *gin.Context) PayeeInfo {
	var name string
	if options != nil && (*options)[umaprotocol.CounterPartyDataFieldName.String()].Mandatory {
		name = v.config.Username
	}
	var email string
	if options != nil && (*options)[umaprotocol.CounterPartyDataFieldEmail.String()].Mandatory {
		email = v.config.Username + "@" + v.getVaspDomain(context)
	}
	return PayeeInfo{
		Name:       &name,
		Email:      &email,
		Identifier: "$" + v.config.Username + "@" + v.getVaspDomain(context),
	}
}

// TODO(Jeremy): Switch back to lightsparkdev/go-sdk version once the UMA changes are merged.
type LightsparkClientUmaInvoiceCreator struct {
	LightsparkClient services.LightsparkClient
	// NodeId: the node ID of the receiver.
	NodeId string
	// ExpirySecs: the number of seconds until the invoice expires.
	ExpirySecs *int32
	// EnableUmaAnalytics: A flag indicating whether UMA analytics should be enabled. If `true`,
	// the receiver identifier will be hashed using a monthly-rotated seed and used for anonymized
	// analysis.
	EnableUmaAnalytics bool
	// SigningPrivateKey: Optional, the receiver's signing private key. Used to hash the receiver
	// identifier if UMA analytics is enabled.
	SigningPrivateKey *[]byte
}

func (l LightsparkClientUmaInvoiceCreator) CreateInvoice(amountMsats int64, metadata string, receiverIdentifier *string) (*string, error) {
	var invoice *objects.Invoice
	var err error
	if l.EnableUmaAnalytics && l.SigningPrivateKey != nil {
		invoice, err = l.LightsparkClient.CreateUmaInvoiceWithReceiverIdentifier(l.NodeId, amountMsats, metadata, l.ExpirySecs, l.SigningPrivateKey, receiverIdentifier)
	} else {
		invoice, err = l.LightsparkClient.CreateUmaInvoice(l.NodeId, amountMsats, metadata, l.ExpirySecs)
	}
	if err != nil {
		return nil, err
	}
	return &invoice.Data.EncodedPaymentRequest, nil
}

type LightsparkClientLnurlInvoiceCreator struct {
	LightsparkClient services.LightsparkClient
	// NodeId: the node ID of the receiver.
	NodeId string
	// ExpirySecs: the number of seconds until the invoice expires.
	ExpirySecs *int32
}

func (l LightsparkClientLnurlInvoiceCreator) CreateInvoice(amountMsats int64, metadata string, receiverIdentifier *string) (*string, error) {
	invoice, err := l.LightsparkClient.CreateLnurlInvoice(l.NodeId, amountMsats, metadata, l.ExpirySecs)
	if err != nil {
		return nil, err
	}
	return &invoice.Data.EncodedPaymentRequest, nil
}
