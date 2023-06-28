package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/services"
)

/**
 * This is a simple Gin server (https://gin-gonic.com) that implements the LNURL payreq protocol
 * using the Lightspark SDK.
 *
 * By default, this server will run on port 8080. You can make a request to the API through curl
 * to make sure the server is working properly (replace ls_test with the username you have
 * configured):
 *
 * curl http://127.0.0.1:8080/.well-known/lnurlp/ls_test
 *
 * Configuration parameters (API keys, etc.) and information on how to set them can be found in
 * config.go.
 */

func getCallback(context *gin.Context, config *LnurlConfig) string {
	return fmt.Sprintf("%s/api/lnurl/payreq/%s", context.Request.Host, config.UserID)
}

func getMetadata(config *LnurlConfig) (string, error) {
	metadata := [][]string{
		{"text/plain", fmt.Sprintf("Pay to domain.org user %s", config.Username)},
		{"text/identifier", fmt.Sprintf("%s@domain.org", config.Username)},
	}

	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		return "", err
	}

	return string(jsonMetadata), nil
}

func handleWellKnownLnurlp(context *gin.Context, config *LnurlConfig) {
	username := context.Param("username")

	if username != config.Username {
		context.JSON(http.StatusNotFound, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("User not found: %s", username),
		})
		return
	}

	callback := getCallback(context, config)
	metadata, err := getMetadata(config)

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

func handleLnurlPayreq(context *gin.Context, config *LnurlConfig) {
	uuid := context.Param("uuid")

	if uuid != config.UserID {
		context.JSON(http.StatusNotFound, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("User not found: %s", uuid),
		})
		return
	}

	amount_param := context.Query("amount")
	amount_msats, err := strconv.Atoi(amount_param)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
			"reason": fmt.Sprintf("Invalid amount %s: %v", amount_param, err),
		})
		return
	}

	metadata, err := getMetadata(config)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "ERROR",
			"reason": err.Error(),
		})
		return
	}

	lsClient := services.NewLightsparkClient(config.ApiClientID, config.ApiClientSecret, nil)
	lsInvoice, err := lsClient.CreateLnurlInvoice(config.NodeUUID, int64(amount_msats), metadata)

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

func main() {
	config := NewConfig()

	engine := gin.Default()
	engine.GET("/.well-known/lnurlp/:username", func(c *gin.Context) {
		handleWellKnownLnurlp(c, &config)
	})
	engine.GET("/api/lnurl/payreq/:uuid", func(c *gin.Context) {
		handleLnurlPayreq(c, &config)
	})

	engine.Run()
}
