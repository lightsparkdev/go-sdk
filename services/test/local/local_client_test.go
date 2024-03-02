//go:build integration
// +build integration

package local_test

import (
	"testing"
	"time"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	servicestest "github.com/lightsparkdev/go-sdk/services/test"
	"github.com/stretchr/testify/require"
)

const TEST_NETWORK = objects.BitcoinNetworkRegtest

func TestCreateInvoice(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(client, env.NodeID)
	require.NoError(t, err)
	t.Log(invoice)
}

// Create invoice for node 1 and pay it from routing node. You'll need to run this a few
// times the first time you are funding a node to get enough funds in.
// Note: This will only work with REGTEST nodes.
func TestCreateTestPaymentNode1(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(client, env.NodeID)
	require.NoError(t, err)
	payment, err := client.CreateTestModePayment(env.NodeID, invoice.Data.EncodedPaymentRequest, nil)
	require.NoError(t, err)
	t.Log(payment)
	// Add a delay to ensure time for the channel to be setup and funded.
	time.Sleep(3 * time.Second)
}

// Create invoice for node 2 and pay it from routing node. You'll need to run this a few
// times the first time you are funding a node to get enough funds in.
// Note: This will only work with REGTEST nodes.
func TestCreateTestPaymentNode2(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(client, env.NodeID2)
	require.NoError(t, err)
	payment, err := client.CreateTestModePayment(env.NodeID2, invoice.Data.EncodedPaymentRequest, nil)
	require.NoError(t, err)
	t.Log(payment)
	// Add a delay to ensure time for the channel to be setup and funded.
	time.Sleep(3 * time.Second)
}

// Create test invoice from routing node and pay it from node 1.
// Note: This will only work with REGTEST nodes.
func TestCreateTestInvoiceNode1(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	invoice, err := client.CreateTestModeInvoice(env.NodeID, 50_000, nil, nil)
	require.NoError(t, err)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID, env.MasterSeedHex, TEST_NETWORK, client)
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

// Create test invoice from routing node and pay it from node 2.
// Note: This will only work with REGTEST nodes.
func TestCreateTestInvoiceNode2(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	invoice, err := client.CreateTestModeInvoice(env.NodeID2, 50_000, nil, nil)
	require.NoError(t, err)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID2, env.MasterSeedHex2, TEST_NETWORK, client)
	payment, err := client.PayInvoice(env.NodeID2, *invoice, 60, 10_000_000, nil)
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

// Create the invoice from node 1, pay it from node 2
func TestPayInvoice(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(client, env.NodeID)
	require.NoError(t, err)

	t.Log(invoice)
	client_2 := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID2, env.MasterSeedHex2, TEST_NETWORK, client_2)
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

func TestGetNodeWalletAddressWithKeys(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	address, err := client.CreateNodeWalletAddress(env.NodeID)
	require.NoError(t, err)
	t.Log(address)
}

func TestGetChannelUtxos(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	utxos, err := client.GetNodeChannelUtxos(env.NodeID)
	require.NoError(t, err)
	t.Log(utxos)
}

func TestGetFundingAddress(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID2, env.MasterSeedHex2, TEST_NETWORK, client)
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
