package main

import (
	"bytes"
	"encoding/json"
	stderrors "errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/uma-universal-money-address/uma-go-sdk/uma"
	"github.com/uma-universal-money-address/uma-go-sdk/uma/errors"
	"github.com/uma-universal-money-address/uma-go-sdk/uma/generated"
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
	return fmt.Sprintf("%s%s/uma/payreq/%s", scheme, context.Request.Host, v.config.UserID)
}

func (v *Vasp2) getUtxoCallback(context *gin.Context, txId string) string {
	scheme := "https://"
	if umautils.IsDomainLocalhost(context.Request.Host) {
		scheme = "http://"
	}
	return fmt.Sprintf("%s%s/uma/utxocallback?txid=%s", scheme, context.Request.Host, txId)
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
		umaError := &errors.UmaError{
			Reason:    fmt.Sprintf("User not found: %s", username),
			ErrorCode: generated.UserNotFound,
		}
		context.AbortWithError(umaError.ToHttpStatusCode(), umaError)
		return
	}

	requestUrl := context.Request.URL
	requestUrl.Host = context.Request.Host

	lnurlpRequest, err := uma.ParseLnurlpRequest(*requestUrl)
	if err != nil {
		var unsupportedVersionErr *uma.UnsupportedVersionError
		if stderrors.As(err, &unsupportedVersionErr) {
			context.Error(err)
		} else {
			context.Error(&errors.UmaError{
				Reason:    err.Error(),
				ErrorCode: generated.ParseLnurlpRequestError,
			})
		}
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
		context.Error(&errors.UmaError{
			Reason:    err.Error(),
			ErrorCode: generated.InternalError,
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

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, response)
}

func (v *Vasp2) handleUmaQueryData(context *gin.Context, lnurlpRequest umaprotocol.UmaLnurlpRequest) (*umaprotocol.LnurlpResponse, bool) {
	vaspDomainValidationErr := ValidateDomain(lnurlpRequest.VaspDomain)
	if vaspDomainValidationErr != nil {
		context.Error(vaspDomainValidationErr)
		return nil, true
	}
	pubKeys, err := uma.FetchPublicKeyForVasp(lnurlpRequest.VaspDomain, v.pubKeyCache)
	if err != nil || pubKeys == nil {
		context.Error(&errors.UmaError{
			Reason:    err.Error(),
			ErrorCode: generated.CounterpartyPubkeyFetchError,
		})
		return nil, true
	}

	if err := uma.VerifyUmaLnurlpQuerySignature(lnurlpRequest, *pubKeys, v.nonceCache); err != nil {
		context.Error(err)
		return nil, true
	}

	metadata, err := v.getMetadata()
	if err != nil {
		context.Error(err)
		return nil, true
	}

	umaPrivateKey, err := v.config.UmaSigningPrivKeyBytes()
	if err != nil {
		context.Error(err)
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
		context.Error(err)
		return nil, true
	}
	return signedResponse, false
}

// This is the handler for regular (non-UMA) LNURL payreq requests when the request is a GET.
func (v *Vasp2) handleLnurlPayreq(context *gin.Context) {
	uuid := context.Param("uuid")

	if uuid != v.config.UserID {
		context.Error(&errors.UmaError{
			Reason:    fmt.Sprintf("User not found: %s", uuid),
			ErrorCode: generated.UserNotFound,
		})
		return
	}

	payreq, err := umaprotocol.ParsePayRequestFromQueryParams(context.Request.URL.Query())
	if err != nil {
		var umaErr *errors.UmaError
		if stderrors.As(err, &umaErr) {
			context.Error(err)
		} else {
			context.Error(&errors.UmaError{
				Reason:    err.Error(),
				ErrorCode: generated.ParsePayreqRequestError,
			})
		}
		return
	}

	metadata, err := v.getMetadata()
	if err != nil {
		context.Error(&errors.UmaError{
			Reason:    err.Error(),
			ErrorCode: generated.InternalError,
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

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, payreqResponse)
}

func (v *Vasp2) handleUmaPayreq(context *gin.Context) {
	uuid := context.Param("uuid")

	if uuid != v.config.UserID {
		context.Error(&errors.UmaError{
			Reason:    fmt.Sprintf("User not found: %s", uuid),
			ErrorCode: generated.UserNotFound,
		})
		return
	}

	requestBody, err := context.GetRawData()
	if err != nil {
		context.Error(&errors.UmaError{
			Reason:    fmt.Sprintf("Invalid request body: %v", err),
			ErrorCode: generated.ParsePayreqRequestError,
		})
		return
	}
	request, err := uma.ParsePayRequest(requestBody)
	if err != nil {
		var umaErr *errors.UmaError
		if stderrors.As(err, &umaErr) {
			context.Error(err)
		} else {
			context.Error(&errors.UmaError{
				Reason:    err.Error(),
				ErrorCode: generated.ParsePayreqRequestError,
			})
		}
		return
	}
	if !request.IsUmaRequest() {
		context.Error(&errors.UmaError{
			Reason:    "Invalid request body: not a UMA request.",
			ErrorCode: generated.InternalError,
		})
		return
	}

	sendingVaspDomain, err := uma.GetVaspDomainFromUmaAddress(*request.PayerData.Identifier())
	if err != nil {
		context.Error(err)
		return
	}
	addressValidationError := ValidateUmaAddress(*request.PayerData.Identifier())
	if addressValidationError != nil {
		context.Error(addressValidationError)
		return
	}

	pubKeys, err := uma.FetchPublicKeyForVasp(sendingVaspDomain, v.pubKeyCache)
	if err != nil || pubKeys == nil {
		context.Error(&errors.UmaError{
			Reason:    err.Error(),
			ErrorCode: generated.CounterpartyPubkeyFetchError,
		})
		return
	}

	if err := uma.VerifyPayReqSignature(request, *pubKeys, v.nonceCache); err != nil {
		context.Error(err)
		return
	}

	metadata, err := v.getMetadata()
	if err != nil {
		context.Error(&errors.UmaError{
			Reason:    err.Error(),
			ErrorCode: generated.InternalError,
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
		context.Error(&errors.UmaError{
			Reason:    fmt.Sprintf("Error getting pre-screening utxos: %v", err),
			ErrorCode: generated.InternalError,
		})
		return
	}

	receiverNode, err := GetNode(lsClient, v.config.NodeUUID)
	if err != nil || receiverNode == nil {
		context.Error(&errors.UmaError{
			Reason:    fmt.Sprintf("Error getting receiver node: %v", err),
			ErrorCode: generated.InternalError,
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
		context.Error(err)
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
		context.Error(err)
	}

	context.JSON(http.StatusOK, response)
}

func (v *Vasp2) handlePubKeyRequest(context *gin.Context) {
	twoWeeksFromNow := time.Now().AddDate(0, 0, 14)
	twoWeeksFromNowSec := twoWeeksFromNow.Unix()
	response, err := uma.GetPubKeyResponse(v.config.UmaSigningCertChain, v.config.UmaEncryptionCertChain, &twoWeeksFromNowSec)
	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, response)
}

func (v *Vasp2) handleUtxoCallback(context *gin.Context) {
	txId := context.Query("txid")
	if txId == "" {
		context.Error(&errors.UmaError{
			Reason:    "Missing txid query parameter",
			ErrorCode: generated.ParseUtxoCallbackError,
		})
		return
	}

	requestBody, err := context.GetRawData()
	if err != nil {
		context.Error(&errors.UmaError{
			Reason:    fmt.Sprintf("Invalid request body: %v", err),
			ErrorCode: generated.ParseUtxoCallbackError,
		})
		return
	}
	callbackData, err := uma.ParsePostTransactionCallback(requestBody)
	if err != nil {
		var umaErr *errors.UmaError
		if stderrors.As(err, &umaErr) {
			context.Error(err)
		} else {
			context.Error(&errors.UmaError{
				Reason:    err.Error(),
				ErrorCode: generated.ParseUtxoCallbackError,
			})
		}
		return
	}

	log.Info("Received UTXO callback", "txId", txId, "callbackData", callbackData)

	context.Status(http.StatusOK)
}

func (v *Vasp2) handleCreateInvoice(context *gin.Context) {
	invoice, err := v.createInvoice(context, false)
	if err != nil {
		context.Error(err)
		return
	}

	invoiceString, err := invoice.ToBech32String()
	if err != nil {
		context.Error(&errors.UmaError{
			Reason:    fmt.Sprintf("Failed to convert invoice to bech32 string: %v", err),
			ErrorCode: generated.InternalError,
		})
		return
	}

	context.JSON(http.StatusOK, invoiceString)
}

func (v *Vasp2) handleCreateAndSendInvoice(context *gin.Context) {
	invoice, err := v.createInvoice(context, true)
	if err != nil {
		context.Error(err)
		return
	}

	// Send the invoice to the sender.

	// Step 1: Get the sender's domain from sender's UMA address.
	senderUma := *invoice.SenderUma
	senderVaspDomain, err := uma.GetVaspDomainFromUmaAddress(senderUma)
	if strings.Contains(senderVaspDomain, "local") {
		senderVaspDomain = "http://" + senderVaspDomain
	} else {
		senderVaspDomain = "https://" + senderVaspDomain
	}
	if err != nil {
		context.Error(err)
		return
	}

	// Step 2: Query sender's domain /.well-known/uma-configruration to get sender's request URL.
	// Make a GET request to the sender's /.well-known/uma-configuration endpoint.
	url := senderVaspDomain + "/.well-known/uma-configuration"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching sending VASP configuration:", err)
		return
	}
	defer resp.Body.Close()

	type UmaConfig struct {
		UmaRequestEndpoint string `json:"uma_request_endpoint"`
	}

	var config UmaConfig

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parse the JSON response
	err = json.Unmarshal(body, &config)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return
	}

	requestEndpoint := config.UmaRequestEndpoint

	// Step 3: Send the invoice to the sender's request URL.
	// Make a POST request to the sender's request URL.
	// The invoice is sent in the request body json "invoice" field.

	invoiceString, err := invoice.ToBech32String()
	if err != nil {
		context.Error(&errors.UmaError{
			Reason:    fmt.Sprintf("Failed to convert invoice to bech32 string: %v", err),
			ErrorCode: generated.InternalError,
		})
		return
	}

	requestBody, err := json.Marshal(map[string]string{
		"invoice": invoiceString,
	})

	if err != nil {
		context.Error(&errors.UmaError{
			Reason:    fmt.Sprintf("Error marshalling request body: %v", err),
			ErrorCode: generated.InternalError,
		})
		return
	}

	resp, err = http.Post(requestEndpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		context.Error(&errors.UmaError{
			Reason:    fmt.Sprintf("Error making POST request: %v", err),
			ErrorCode: generated.InternalError,
		})
		return
	}
	defer resp.Body.Close()

	context.Status(http.StatusOK)
}

func (v *Vasp2) createInvoice(context *gin.Context, request bool) (*umaprotocol.UmaInvoice, error) {
	uuid := context.Param("uuid")
	if uuid != v.config.UserID {
		return nil, &errors.UmaError{
			Reason:    fmt.Sprintf("user not found: %s", uuid),
			ErrorCode: generated.UserNotFound,
		}
	}

	var requestBody struct {
		Amount       int64   `json:"amount"`
		CurrencyCode string  `json:"currency_code"`
		SenderUma    *string `json:"sender_uma"`
	}
	if err := context.BindJSON(&requestBody); err != nil {
		return nil, &errors.UmaError{
			Reason:    fmt.Sprintf("failed to bind request body: %v", err),
			ErrorCode: generated.InvalidInput,
		}
	}

	if requestBody.SenderUma == nil && request {
		return nil, &errors.UmaError{
			Reason:    "sender_uma is required",
			ErrorCode: generated.InvalidInput,
		}
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
		return nil, &errors.UmaError{
			Reason:    fmt.Sprintf("user does not support currency %s", requestBody.CurrencyCode),
			ErrorCode: generated.InvalidCurrency,
		}
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
		return nil, err
	}

	invoice, err := uma.CreateUmaInvoice(
		"$"+v.config.Username+"@"+v.getVaspDomain(context),
		uint64(requestBody.Amount),
		umaprotocol.InvoiceCurrency{
			Code:     currency.Code,
			Decimals: uint8(currency.Decimals),
			Symbol:   currency.Symbol,
			Name:     currency.Name,
		},
		uint64(twoDaysFromNow.Unix()),
		callback,
		true,
		&payerDataOptions,
		nil,
		nil,
		nil,
		requestBody.SenderUma,
		privateKey,
	)

	if err != nil {
		return nil, err
	}

	return invoice, nil
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
