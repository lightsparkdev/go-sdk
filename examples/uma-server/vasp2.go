package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/services"
	lsuma "github.com/lightsparkdev/go-sdk/uma"
	"github.com/uma-universal-money-address/uma-go-sdk/uma"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Vasp2 is an implementation of the receiving VASP in the UMA protocol.
type Vasp2 struct {
	config      *UmaConfig
	pubKeyCache uma.PublicKeyCache
}

func (v *Vasp2) getLnurlpCallback(context *gin.Context) string {
	scheme := "https://"
	if strings.HasPrefix(context.Request.Host, "localhost:") {
		scheme = "http://"
	}
	return fmt.Sprintf("%s%s/api/uma/payreq/%s", scheme, context.Request.Host, v.config.UserID)
}

func (v *Vasp2) getUtxoCallback(context *gin.Context, txId string) string {
	scheme := "https://"
	if strings.HasPrefix(context.Request.Host, "localhost:") {
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

	// Fallback to regular LNURL if the request is not a UMA request.
	if !uma.IsUmaLnurlpQuery(*requestUrl) {
		v.handleNonUmaLnurlRequest(context)
		return
	}

	responseJson, hadError := v.parseUmaQueryData(context)
	if hadError {
		return
	}
	context.Data(http.StatusOK, "application/json", responseJson)
	return
}

func (v *Vasp2) handleNonUmaLnurlRequest(context *gin.Context) {
	callback := v.getLnurlpCallback(context)
	metadata, err := v.getMetadata()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"callback":    callback,
		"maxSendable": 10_000_000,
		"minSendable": 1_000,
		"metadata":    metadata,
		"tag":         "payRequest",
	})
}

func (v *Vasp2) parseUmaQueryData(context *gin.Context) ([]byte, bool) {
	requestUrl := context.Request.URL
	requestUrl.Host = context.Request.Host
	query, err := uma.ParseLnurlpRequest(*requestUrl)
	if err != nil {
		var unsupportedVersionErr *uma.UnsupportedVersionError
		if errors.As(err, &unsupportedVersionErr) {
			context.JSON(http.StatusPreconditionFailed, gin.H{
				"status":                 "ERROR",
				"reason":                 fmt.Sprintf("Unsupported version: %s", unsupportedVersionErr.UnsupportedVersion),
				"supportedMajorVersions": unsupportedVersionErr.SupportedMajorVersions,
				"unsupportedVersion":     unsupportedVersionErr.UnsupportedVersion,
			})
			return nil, true
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}

	pubKeys, err := uma.FetchPublicKeyForVasp(query.VaspDomain, v.pubKeyCache)
	if err != nil || pubKeys == nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}

	sendingVaspSigningPubKey, err := pubKeys.SigningPubKey()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}
	if err := uma.VerifyUmaLnurlpQuerySignature(query, sendingVaspSigningPubKey); err != nil {
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

	signedResponse, err := uma.GetLnurlpResponse(
		query,
		umaPrivateKey,
		true,
		v.getLnurlpCallback(context),
		metadata,
		1,
		10_000_000,
		uma.PayerDataOptions{
			NameRequired:       false,
			EmailRequired:      false,
			ComplianceRequired: true,
		},
		[]uma.Currency{
			{
				Code:                "USD",
				Name:                "US Dollar",
				Symbol:              "$",
				MillisatoshiPerUnit: 34_150,
				MinSendable:         1,
				MaxSendable:         10_000_000,
			},
		},
		uma.KycStatusVerified,
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}

	responseJson, err := json.Marshal(signedResponse)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}
	return responseJson, false
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

	amountParam := context.Query("amount")
	amountMsats, err := strconv.Atoi(amountParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid amount %s: %v", amountParam, err),
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

	lsClient := services.NewLightsparkClient(v.config.ApiClientID, v.config.ApiClientSecret, &v.config.ClientBaseURL)
	lsInvoice, err := lsClient.CreateLnurlInvoice(v.config.NodeUUID, int64(amountMsats), metadata, nil)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"pr":     lsInvoice.Data.EncodedPaymentRequest,
		"routes": []string{},
	})
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

	sendingVaspDomain, err := uma.GetVaspDomainFromUmaAddress(request.PayerData.Identifier)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid sender indentifier. UMA address required in the format $alice@vasp.com: %v", err),
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

	sendingVaspSigningPubKey, err := pubKeys.SigningPubKey()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}
	if err := uma.VerifyPayReqSignature(request, sendingVaspSigningPubKey); err != nil {
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

	lsClient := services.NewLightsparkClient(v.config.ApiClientID, v.config.ApiClientSecret, &v.config.ClientBaseURL)
	expirySecs := int32(600) // Expire in 10 minutes
	invoiceCreator := lsuma.LightsparkClientUmaInvoiceCreator{
		LightsparkClient: *lsClient,
		NodeId:           v.config.NodeUUID,
		ExpirySecs:       &expirySecs,
	}

	conversionRate := int64(34_150)
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

	response, err := uma.GetPayReqResponse(
		request,
		invoiceCreator,
		metadata,
		"USD",
		conversionRate,
		exchangeFees,
		receiverUtxos,
		nil,
		v.getUtxoCallback(context, txID),
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
		SigningPubKeyHex:    hex.EncodeToString(signingPubKeyBytes),
		EncryptionPubKeyHex: hex.EncodeToString(encryptionPubKeyBytes),
		ExpirationTimestamp: &twoWeeksFromNowSec,
	}

	context.JSON(http.StatusOK, response)
}
