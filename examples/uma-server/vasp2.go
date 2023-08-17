package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/lightsparkdev/go-sdk/uma"
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

func (v *Vasp2) getCallback(context *gin.Context) string {
	scheme := "https://"
	if strings.HasPrefix(context.Request.Host, "localhost:") {
		scheme = "http://"
	}
	return fmt.Sprintf("%s%s/api/uma/payreq/%s", scheme, context.Request.Host, v.config.UserID)
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

	if username != v.config.Username {
		context.JSON(http.StatusNotFound, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("User not found: %s", username),
		})
		return
	}

	requestUrl := context.Request.URL
	requestUrl.Host = context.Request.Host
	if uma.IsUmaLnurlpQuery(*requestUrl) {
		responseJson, hadError := v.parseUmaQueryData(context)
		if hadError {
			return
		}
		context.Data(http.StatusOK, "application/json", responseJson)
		return
	}

	callback := v.getCallback(context)
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
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}

	pubKey, err := uma.FetchPublicKeyForVasp(query.VaspDomain, v.pubKeyCache)
	if err != nil || pubKey == nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return nil, true
	}

	if err := uma.VerifyUmaLnurlpQuerySignature(query, pubKey.SigningPubKey); err != nil {
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
		v.getCallback(context),
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
		true,
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

	masterSeedBytes, err := v.config.NodeMasterSeedBytes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Invalid master seed",
		})
	}
	lsClient := services.NewLightsparkClient(v.config.ApiClientID, v.config.ApiClientSecret, &v.config.ClientBaseURL)
	lsInvoice, err := lsClient.CreateLnurlInvoice(v.config.NodeUUID, &masterSeedBytes, int64(amountMsats), metadata, nil)

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

	metadata, err := v.getMetadata()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	conversionRate := 34_150
	lsClient := services.NewLightsparkClient(v.config.ApiClientID, v.config.ApiClientSecret, &v.config.ClientBaseURL)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	masterSeedBytes, err := v.config.NodeMasterSeedBytes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": "Invalid master seed",
		})
	}
	response, err := uma.GetPayReqResponse(
		request,
		lsClient,
		v.config.NodeUUID,
		&masterSeedBytes,
		metadata,
		"USD",
		int64(conversionRate),
		int32(time.Hour.Seconds()),
		// TODO: Actually get the UTXOs from the request.
		[]string{"abcdef12345"},
		"/api/lnurl/utxocallback?txid=1234",
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
		SigningPubKey:       signingPubKeyBytes,
		EncryptionPubKey:    encryptionPubKeyBytes,
		ExpirationTimestamp: &twoWeeksFromNowSec,
	}

	context.JSON(http.StatusOK, response)
}
