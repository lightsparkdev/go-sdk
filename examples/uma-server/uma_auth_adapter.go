package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/lightsparkdev/go-sdk/utils"
	decodepay "github.com/nbd-wtf/ln-decodepay"
	"github.com/uma-universal-money-address/uma-auth-api/codegen/go/umaauth"
	"math"
)

type UmaAuthAdapter struct {
	sendingVasp *Vasp1
	config      *UmaConfig
	client      *services.LightsparkClient
	userService UserService
}

func (a *UmaAuthAdapter) ExecuteQuote(context *gin.Context, paymentHash string, request umaauth.ExecuteQuoteRequest) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) FetchQuoteForLud16(context *gin.Context, sendingCurrencyCode string, receivingCurrencyCode string, lockedCurrencyAmount int64, lockedCurrencySide umaauth.LockedCurrencySide, receiverAddress string) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) GetBalance(context *gin.Context) {
	currencyCode := context.Query("currency_code")
	if currencyCode != "" && currencyCode != "SAT" {
		context.JSON(400, gin.H{"error": fmt.Sprintf("unsupported currency code: %s", currencyCode)})
		return
	}
	// TODO: If we had multiple users and wanted to get the balance for a specific user, we would need to tweak this to
	// only read the active user's balance.
	node, err := GetNode(a.client, a.config.NodeUUID)
	if err != nil {
		context.JSON(500, gin.H{"error": fmt.Sprintf("error getting node: %s", err)})
		return
	}

	msats, err := utils.ValueMilliSatoshi((*node).GetBalances().OwnedBalance)
	if err != nil {
		context.JSON(500, gin.H{"error": fmt.Sprintf("error getting balance: %s", err)})
		return
	}

	context.JSON(200, &umaauth.GetBalanceResponse{
		Balance: float32(msats),
	})
}

func (a *UmaAuthAdapter) GetBudgetEstimate(context *gin.Context, sendingCurrencyCode string, sendingCurrencyAmount int64, budgetCurrencyCode string) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) GetInfo(context *gin.Context) {
	user, err := a.userService.GetUserFromContext(context)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	node, err := GetNode(a.client, a.config.NodeUUID)
	if err != nil {
		context.JSON(500, gin.H{"error": fmt.Sprintf("error getting node: %s", err)})
		return
	}

	network := "mainnet"
	if (*node).GetBitcoinNetwork() == objects.BitcoinNetworkTestnet {
		network = "testnet"
	} else if (*node).GetBitcoinNetwork() == objects.BitcoinNetworkRegtest {
		network = "regtest"
	} else if (*node).GetBitcoinNetwork() == objects.BitcoinNetworkSignet {
		network = "signet"
	}

	pubkey := (*node).GetPublicKey()
	if pubkey == nil {
		pubkey = new(string)
	}

	color := (*node).GetColor()
	if color == nil {
		color = new(string)
	}

	context.JSON(200, &umaauth.GetInfoResponse{
		Alias:   "Golang Demo VASP",
		Color:   *color,
		Pubkey:  *pubkey,
		Network: network,
		Methods: a.config.SupportedNwcCommands,
		Lud16:   user.GetUmaAddress(a.config, context),
	})
}

func (a *UmaAuthAdapter) ListTransactions(context *gin.Context, from int64, until int64, limit int32, offset int32, unpaid bool, transactionType umaauth.TransactionType) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) LookupInvoice(context *gin.Context, paymentHash string) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) LookupUserByLud16(context *gin.Context, receiverAddress string, baseSendingCurrencyCode string) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) MakeInvoice(context *gin.Context, request umaauth.MakeInvoiceRequest) {
	invoiceType := objects.InvoiceTypeStandard
	invoice, err := a.client.CreateInvoice(a.config.NodeUUID, request.Amount, &request.Description, &invoiceType, &request.Expiry)
	if err != nil {
		context.JSON(500, gin.H{"error": fmt.Sprintf("error creating invoice: %s", err)})
		return
	}
	expiresAt := invoice.Data.ExpiresAt.Unix()
	context.JSON(200, &umaauth.Transaction{
		Type:        umaauth.INCOMING,
		Amount:      request.Amount,
		Invoice:     &invoice.Data.EncodedPaymentRequest,
		Description: invoice.Data.Memo,
		ExpiresAt:   &expiresAt,
		PaymentHash: invoice.Data.PaymentHash,
		CreatedAt:   invoice.CreatedAt.Unix(),
	})
}

func (a *UmaAuthAdapter) PayInvoice(context *gin.Context, request umaauth.PayInvoiceRequest) {
	decodedInvoice, err := decodepay.Decodepay(request.Invoice)
	if err != nil {
		context.JSON(400, gin.H{"error": fmt.Sprintf("error decoding invoice: %s", err)})
		return
	}

	amountMsats := request.Amount
	if amountMsats == nil && decodedInvoice.MSatoshi == 0 {
		context.JSON(400, gin.H{"error": "amount must be specified for zero amount invoices"})
		return
	}
	if amountMsats != nil && *amountMsats != decodedInvoice.MSatoshi {
		context.JSON(400, gin.H{"error": "amount does not match invoice amount"})
		return
	}

	amountToPay := amountMsats
	if amountToPay == nil {
		amountToPay = &decodedInvoice.MSatoshi
	}
	maxFeeMsats := int64(max(1000, math.Round(float64(*amountToPay)*0.0016)))
	payment, err := a.client.PayInvoice(a.config.NodeUUID, request.Invoice, 60, maxFeeMsats, amountMsats)
	if err != nil {
		context.JSON(500, gin.H{"error": fmt.Sprintf("error paying invoice: %s", err)})
		return
	}

}

func (a *UmaAuthAdapter) PayKeysend(context *gin.Context, request umaauth.PayKeysendRequest) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) PayToLud16Address(context *gin.Context, request umaauth.PayToAddressRequest) {
	//TODO implement me
	panic("implement me")
}

func NewUmaAuthAdapter(config *UmaConfig, sendingVasp *Vasp1, userService UserService) *UmaAuthAdapter {
	return &UmaAuthAdapter{
		sendingVasp: sendingVasp,
		config:      config,
		client:      services.NewLightsparkClient(config.ApiClientID, config.ApiClientSecret, config.ClientBaseURL),
		userService: userService,
	}
}

func (a *UmaAuthAdapter) HandleTokenExchange(context *gin.Context) string {
	return a.config.GetVaspDomain(context)
}

func (a *UmaAuthAdapter) RegisterRoutes(engine *gin.Engine) {
	//umaAuthController := umaauth.NewUmaAuthAPIController(a)
	//routes := umaAuthController.Routes()
	//for _, route := range routes {
	//	engine.Handle(route.Method, route.Pattern, func(context *gin.Context) {
	//		route.HandlerFunc(context.Writer, context.Request)
	//	})
	//}

	engine.GET("/umanwc/balance", func(context *gin.Context) {
		a.GetBalance(context)
	})

	engine.GET("/umanwc/info", func(context *gin.Context) {
		a.GetInfo(context)
	})

	engine.POST("/umanwc/invoice", func(context *gin.Context) {
		body, err := context.GetRawData()
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var requestBody umaauth.MakeInvoiceRequest
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		a.MakeInvoice(context, requestBody)
	})

	engine.POST("/umanwc/payments/bolt11", func(context *gin.Context) {
		body, err := context.GetRawData()
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var requestBody umaauth.PayInvoiceRequest
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		a.PayInvoice(context, requestBody)
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
