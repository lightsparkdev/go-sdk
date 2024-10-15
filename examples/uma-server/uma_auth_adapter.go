package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/lightsparkdev/go-sdk/utils"
	"github.com/uma-universal-money-address/uma-auth-api/codegen/go/umaauth"
)

type UmaAuthAdapter struct {
	sendingVasp *Vasp1
	config      *UmaConfig
	client      *services.LightsparkClient
	userService UserService
}

func NewUmaAuthAdapter(config *UmaConfig, sendingVasp *Vasp1, userService UserService) *UmaAuthAdapter {
	return &UmaAuthAdapter{
		sendingVasp: sendingVasp,
		config:      config,
		client:      services.NewLightsparkClient(config.ApiClientID, config.ApiClientSecret, config.ClientBaseURL),
		userService: userService,
	}
}

func (a *UmaAuthAdapter) GetBalance(currencyCode string) (*umaauth.GetBalanceResponse, error) {
	if currencyCode != "" && currencyCode != "SAT" {
		return nil, fmt.Errorf("unsupported currency code: %s", currencyCode)
	}
	// TODO: If we had multiple users and wanted to get the balance for a specific user, we would need to tweak this to
	// only read the active user's balance.
	node, err := GetNode(a.client, a.config.NodeUUID)
	if err != nil {
		return nil, fmt.Errorf("error getting balance: %s", err)
	}

	msats, err := utils.ValueMilliSatoshi((*node).GetBalances().OwnedBalance)
	if err != nil {
		return nil, fmt.Errorf("error getting balance: %s", err)
	}

	return &umaauth.GetBalanceResponse{
		Balance: float32(msats),
	}, nil
}

func (a *UmaAuthAdapter) HandleTokenExchange(context *gin.Context) string {
	return a.config.GetVaspDomain(context)
}

func (a *UmaAuthAdapter) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/umanwc/balance", func(context *gin.Context) {
		currencyCode := context.Query("currency_code")
		response, err := a.GetBalance(currencyCode)
		if err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}
		context.JSON(200, response)
	})

	engine.POST("/umanwc/token", func(context *gin.Context) {
		user, err := a.userService.GetUserFromContext(context)
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		body, err := context.GetRawData()
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var bodyJson map[string]interface{}
		err = json.Unmarshal(body, &bodyJson)
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		requestedExpirySecs, ok := bodyJson["expiration"].(int)
		if !ok {
			context.JSON(400, gin.H{"error": "expiration must be an integer"})
			return
		}

		umaAuthJwt := NewUmaAuthJwtForUser(user, a.config, context, requestedExpirySecs)
		token, err := umaAuthJwt.Sign(a.config)
		if err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}
		context.JSON(200, gin.H{"token": token})
	})

	// TODO: Implement the rest
}
