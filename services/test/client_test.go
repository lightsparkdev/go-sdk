//go:build integration
// +build integration

package services_test

import (
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type TestConfig struct {
	ApiClientEndpoint string
	ApiClientID       string
	ApiClientSecret   string
	NodeID            string
	ApiClientID2      string
	ApiClientSecret2  string
	NodeID2           string
}

func NewConfig() TestConfig {
	var endpoint string
	if endpoint = os.Getenv("LIGHTSPARK_API_ENDPOINT"); endpoint == "" {
		endpoint = "https://api.dev.dev.sparkinfra.net/graphql/server/rc"
	}

	return TestConfig{
		ApiClientEndpoint: endpoint,
		ApiClientID:       os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_ID"),
		ApiClientSecret:   os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_SECRET"),
		NodeID:            os.Getenv("LIGHTSPARK_RS_NODE_ID"),
		NodeID2:           os.Getenv("LIGHTSPARK_RS_NODE_ID_2"),
		ApiClientID2:      os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_ID_2"),
		ApiClientSecret2:  os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_SECRET_2"),
	}
}

func TestCreateInvoice(t *testing.T) {
	env := NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(client, env.NodeID)
	require.NoError(t, err)
	t.Log(invoice)
}

// Note: This will only work with REGTEST nodes.
func TestAddTestFunds(t *testing.T) {
	env := NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	amount, err := client.FundNode(env.NodeID2, 10_000_000)
	require.NoError(t, err)
	t.Logf("Added %d %s", amount.OriginalValue, amount.OriginalUnit.StringValue())
}

// Create invoice for node 2 and pay it from routing node since AddFunds above
// has some delay before the funds are available.
// Note: This will only work with REGTEST nodes.
func TestCreateTestPaymentNode2(t *testing.T) {
	env := NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	// Create and pay invoice 10 times to ensure we get enough funds on node 1
	// to pay an invoice in the next test.
	for i := 0; i < 10; i++ {
		invoice, err := createInvoiceForNode(client, env.NodeID2)
		require.NoError(t, err)
		payment, err := client.CreateTestModePayment(env.NodeID2, invoice.Data.EncodedPaymentRequest, nil)
		require.NoError(t, err)
		t.Log(payment)
	}
}

// Create the invoice from node 1, pay it from node 2
func TestPayInvoice(t *testing.T) {
	env := NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(client, env.NodeID)
	require.NoError(t, err)

	t.Log(invoice)
	client_2 := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	payment, err := client_2.PayInvoice(env.NodeID2, invoice.Data.EncodedPaymentRequest, 60, 1000000, nil)
	require.NoError(t, err)

	for payment.Status != objects.TransactionStatusSuccess && payment.Status != objects.TransactionStatusFailed {
		entity, err := client_2.GetEntity(payment.Id)
		require.NoError(t, err)
		castPayment, didCast := (*entity).(objects.OutgoingPayment)
		require.True(t, didCast)
		payment = &castPayment
	}
	if payment.Status == objects.TransactionStatusFailed {
		t.Errorf("Payment failed: %s", payment.FailureReason.StringValue())
	}
	t.Log(payment)
}

// Create invoice for node 1 and pay it from routing node.
// Note: This will only work with REGTEST nodes.
func TestCreateTestPaymentNode1(t *testing.T) {
	env := NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	// Create and pay invoice 10 times to ensure we get enough funds on node 1
	// to pay an invoice in the next test.
	for i := 0; i < 10; i++ {
		invoice, err := createInvoiceForNode(client, env.NodeID)
		require.NoError(t, err)
		payment, err := client.CreateTestModePayment(env.NodeID, invoice.Data.EncodedPaymentRequest, nil)
		require.NoError(t, err)
		t.Log(payment)
	}
}

// Note: This will only work with REGTEST nodes.
func TestCreateTestInvoice(t *testing.T) {
	env := NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	invoice, err := client.CreateTestModeInvoice(env.NodeID, 50_000, nil, nil)
	require.NoError(t, err)
	payment, err := client.PayInvoice(env.NodeID, *invoice, 60, 10_000_000, nil)
	require.NoError(t, err)
	for payment.Status != objects.TransactionStatusSuccess && payment.Status != objects.TransactionStatusFailed {
		entity, err := client.GetEntity(payment.Id)
		require.NoError(t, err)
		castPayment, didCast := (*entity).(objects.OutgoingPayment)
		require.True(t, didCast)
		payment = &castPayment
	}
	if payment.Status == objects.TransactionStatusFailed {
		t.Error("Payment failed")
	}
	t.Log(payment)
}

func TestGetChannelUtxos(t *testing.T) {
	env := NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	utxos, err := client.GetNodeChannelUtxos(env.NodeID)
	require.NoError(t, err)
	t.Log(utxos)
}

// NOTE: This will only work with MAINNET nodes.
func TestGetFundingAddress(t *testing.T) {
	env := NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	address, err := client.CreateNodeWalletAddress(env.NodeID2)
	require.NoError(t, err)
	t.Log(address)
}

func createInvoiceForNode(client *services.LightsparkClient, nodeID string) (*objects.Invoice, error) {
	invoice, err := client.CreateInvoice(nodeID, 10_000_000, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}
