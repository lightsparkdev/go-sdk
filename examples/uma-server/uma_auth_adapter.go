package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lightsparkdev/go-sdk/objects"
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

func (a *UmaAuthAdapter) ExecuteQuote(ctx context.Context, paymentHash string, request umaauth.ExecuteQuoteRequest) (umaauth.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) FetchQuoteForLud16(ctx context.Context, sendingCurrencyCode string, receivingCurrencyCode string, lockedCurrencyAmount int64, lockedCurrencySide umaauth.LockedCurrencySide, receiverAddress string) (umaauth.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) GetBalance(ctx context.Context, currencyCode string) (umaauth.ImplResponse, error) {
	if currencyCode != "" && currencyCode != "SAT" {
		return umaauth.ImplResponse{Code: 400}, fmt.Errorf("unsupported currency code: %s", currencyCode)
	}
	// TODO: If we had multiple users and wanted to get the balance for a specific user, we would need to tweak this to
	// only read the active user's balance.
	node, err := GetNode(a.client, a.config.NodeUUID)
	if err != nil {
		return umaauth.ImplResponse{Code: 500}, fmt.Errorf("error getting balance: %s", err)
	}

	msats, err := utils.ValueMilliSatoshi((*node).GetBalances().OwnedBalance)
	if err != nil {
		return umaauth.ImplResponse{Code: 500}, fmt.Errorf("error getting balance: %s", err)
	}

	return umaauth.ImplResponse{
		Code: 200,
		Body: &umaauth.GetBalanceResponse{
			Balance: float32(msats),
		},
	}, nil
}

func (a *UmaAuthAdapter) GetBudgetEstimate(ctx context.Context, sendingCurrencyCode string, sendingCurrencyAmount int64, budgetCurrencyCode string) (umaauth.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) GetInfo(ctx context.Context) (umaauth.ImplResponse, error) {
	node, err := GetNode(a.client, a.config.NodeUUID)
	if err != nil {
		return umaauth.ImplResponse{Code: 500}, fmt.Errorf("error getting node: %s", err)
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

	return umaauth.ImplResponse{
		Code: 200,
		Body: &umaauth.GetInfoResponse{
			Alias:   "Golang Demo VASP",
			Color:   *color,
			Pubkey:  *pubkey,
			Network: network,
			Methods: a.config.SupportedNwcCommands,
			Lud16:
		},
	}, nil
}

func (a *UmaAuthAdapter) ListTransactions(ctx context.Context, from int64, until int64, limit int32, offset int32, unpaid bool, transactionType umaauth.TransactionType) (umaauth.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) LookupInvoice(ctx context.Context, paymentHash string) (umaauth.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) LookupUserByLud16(ctx context.Context, receiverAddress string, baseSendingCurrencyCode string) (umaauth.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) MakeInvoice(ctx context.Context, request umaauth.MakeInvoiceRequest) (umaauth.ImplResponse, error) {
	invoiceType := objects.InvoiceTypeStandard
	invoice, err := a.client.CreateInvoice(a.config.NodeUUID, request.Amount, &request.Description, &invoiceType, &request.Expiry)
	if err != nil {
		return umaauth.ImplResponse{Code: 500}, fmt.Errorf("error creating invoice: %s", err)
	}
	expiresAt := invoice.Data.ExpiresAt.Unix()
	return umaauth.ImplResponse{
		Code: 200,
		Body: &umaauth.Transaction{
			Type:        umaauth.INCOMING,
			Amount:      request.Amount,
			Invoice:     &invoice.Data.EncodedPaymentRequest,
			Description: invoice.Data.Memo,
			ExpiresAt:   &expiresAt,
			PaymentHash: invoice.Data.PaymentHash,
			CreatedAt:   invoice.CreatedAt.Unix(),
		},
	}, nil
}

func (a *UmaAuthAdapter) PayInvoice(ctx context.Context, request umaauth.PayInvoiceRequest) (umaauth.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) PayKeysend(ctx context.Context, request umaauth.PayKeysendRequest) (umaauth.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *UmaAuthAdapter) PayToLud16Address(ctx context.Context, request umaauth.PayToAddressRequest) (umaauth.ImplResponse, error) {
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
	umaAuthController := umaauth.NewUmaAuthAPIController(a)
	routes := umaAuthController.Routes()
	for _, route := range routes {
		engine.Handle(route.Method, route.Pattern, func(context *gin.Context) {
			route.HandlerFunc(context.Writer, context.Request)
		})
	}

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
