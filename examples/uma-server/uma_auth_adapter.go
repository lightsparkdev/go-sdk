package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/services"
)

type UmaAuthAdapter struct {
	sendingVasp *Vasp1
	config      *UmaConfig
	client      *services.LightsparkClient
	userService *UserService
}

func NewUmaAuthAdapter(config *UmaConfig, sendingVasp *Vasp1, userService UserService) *UmaAuthAdapter {
	return &UmaAuthAdapter{
		sendingVasp: sendingVasp,
		config:      config,
		client:      services.NewLightsparkClient(config.ApiClientID, config.ApiClientSecret, config.ClientBaseURL),
		userService: &userService,
	}
}

func (a *UmaAuthAdapter) GetBalance() (int, error) {
	// TODO: Implement this
	return 0, nil
}
func (a *UmaAuthAdapter) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/umanwc/balance", func(context *gin.Context) {
		balance, err := a.GetBalance()
		if err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}
		context.JSON(200, gin.H{"balance": balance})
	})
	// TODO: Implement the rest
}
