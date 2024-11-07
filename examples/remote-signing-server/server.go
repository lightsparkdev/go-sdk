// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/remotesigning"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/lightsparkdev/go-sdk/webhooks"
)

func getValidator(config *Config) remotesigning.Validator {
	if config.ValidationEnabled {
		return &remotesigning.HashValidator{}
	} else {
		return &remotesigning.PositiveValidator{}
	}
}

/**
 * This is a simple Gin server (https://gin-gonic.com) that implements a simple remote-signer using
 * the Lightspark SDK.
 *
 * By default, this server will run on port 8080. You can make a request to the API through curl
 * to make sure the server is working properly (replace ls_test with the username you have
 * configured):
 *
 * curl 127.0.0.1:8080/ping
 *
 */

func main() {
	config, err := NewConfigFromEnv()
	if err != nil {
		log.Fatalf("Invalid config: %s", err)
	}

	lsClient := services.NewLightsparkClient(config.ApiClientId, config.ApiClientSecret, config.ApiEndpoint)
	validator := getValidator(config)

	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	engine.POST("/ln/webhooks", func(c *gin.Context) {
		signature := c.Request.Header.Get(webhooks.SIGNATURE_HEADER)
		if signature == "" {
			log.Print("ERROR: Signature was not present")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		data, err := c.GetRawData()
		if err != nil {
			log.Printf("ERROR: Couldn't get data: %s", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		event, err := webhooks.VerifyAndParse(data, signature, config.WebhookSecret)
		if err != nil {
			log.Printf("ERROR: Couldn't parse webhook data: %s", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		log.Printf("Received %s", event.EventType.StringValue())

		switch event.EventType {
		case objects.WebhookEventTypeRemoteSigning:
			if config.RespondDirectly {
				resp, err := remotesigning.GraphQLResponseForRemoteSigningWebhook(
					remotesigning.HashValidator{}, *event, config.MasterSeed)
				if err != nil {
					log.Printf("ERROR: Unable to handle remote signing webhook: %s", err)
					c.AbortWithStatus(http.StatusInternalServerError)
					return
				}

				if resp != nil {
					c.JSON(http.StatusOK, resp.GraphqlResponse().Variables)
				} else {
					c.Status(http.StatusNoContent)
				}
			} else {
				resp, err := remotesigning.HandleRemoteSigningWebhook(
					lsClient, validator, *event, config.MasterSeed)
				if err != nil {
					log.Printf("ERROR: Unable to handle remote signing webhook: %s", err)
					c.AbortWithStatus(http.StatusInternalServerError)
					return
				}

				if resp != "" {
					log.Printf("Webhook complete with response: %s", resp)
				} else {
					log.Printf("Webhook complete")
				}

				c.Status(http.StatusNoContent)
			}
		default:
			c.Status(http.StatusNoContent)
		}
	})

	engine.Run()
}
