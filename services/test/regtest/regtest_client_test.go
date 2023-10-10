//go:build integration
// +build integration

package regtest_test

import (
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	servicestest "github.com/lightsparkdev/go-sdk/services/test"
	"github.com/lightsparkdev/go-sdk/utils"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
	"time"
)

const (
	InvoiceAmount = 1_000_000
)

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
	tx := *waitForPaymentCompletion(t, client, payment.Id)
	require.Equal(t, objects.TransactionStatusSuccess, tx.GetStatus())
	t.Log(tx)
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
	tx := *waitForPaymentCompletion(t, client, payment.Id)
	require.Equal(t, objects.TransactionStatusSuccess, tx.GetStatus())
	t.Log(tx)
}

// Create test invoice from routing node and pay it from node 1.
func TestCreateTestInvoiceAndPayFromNode1(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	ensureEnoughNodeFunds(t, client, env.NodeID, InvoiceAmount)
	invoice, err := client.CreateTestModeInvoice(env.NodeID, 50_000, nil, nil)
	require.NoError(t, err)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID, env.MasterSeedHex, objects.BitcoinNetworkRegtest, client)
	payment, err := client.PayInvoice(env.NodeID, *invoice, 60, 100_000, nil)
	require.NoError(t, err)

	completedTransaction := *waitForPaymentCompletion(t, client, payment.Id)
	if completedTransaction.GetStatus() == objects.TransactionStatusFailed {
		t.Errorf("Payment failed: %s", payment.FailureReason.StringValue())
	}
	t.Log(completedTransaction)
}

// Create test invoice from routing node and pay it from node 2.
func TestCreateTestInvoiceAndPayFromNode2(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	ensureEnoughNodeFunds(t, client, env.NodeID2, InvoiceAmount)
	invoice, err := client.CreateTestModeInvoice(env.NodeID2, 50_000, nil, nil)
	require.NoError(t, err)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID2, env.MasterSeedHex2, objects.BitcoinNetworkRegtest, client)
	payment, err := client.PayInvoice(env.NodeID2, *invoice, 60, 100_000, nil)
	require.NoError(t, err)

	completedTransaction := *waitForPaymentCompletion(t, client, payment.Id)
	if completedTransaction.GetStatus() == objects.TransactionStatusFailed {
		t.Errorf("Payment failed: %s", payment.FailureReason.StringValue())
	}
	t.Log(completedTransaction)
}

// Create an invoice from node 1, pay it from node 2
func TestPayInvoiceNode2ToNode1(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(client, env.NodeID)
	require.NoError(t, err)

	t.Log(invoice)
	client2 := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	ensureEnoughNodeFunds(t, client2, env.NodeID2, InvoiceAmount)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID2, env.MasterSeedHex2, objects.BitcoinNetworkRegtest, client2)
	payment, err := client2.PayInvoice(env.NodeID2, invoice.Data.EncodedPaymentRequest, 60, InvoiceAmount*0.16, nil)
	require.NoError(t, err)

	completedTransaction := *waitForPaymentCompletion(t, client2, payment.Id)
	if completedTransaction.GetStatus() == objects.TransactionStatusFailed {
		t.Errorf("Payment failed: %s", payment.FailureReason.StringValue())
	}
	t.Log(completedTransaction)
}

// Create an invoice from node 2, pay it from node 1
func TestPayInvoiceNode1ToNode2(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(client, env.NodeID2)
	require.NoError(t, err)

	t.Log(invoice)
	client2 := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	ensureEnoughNodeFunds(t, client2, env.NodeID, InvoiceAmount)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID, env.MasterSeedHex, objects.BitcoinNetworkRegtest, client2)
	payment, err := client2.PayInvoice(env.NodeID, invoice.Data.EncodedPaymentRequest, 60, InvoiceAmount*0.16, nil)
	require.NoError(t, err)

	completedTransaction := *waitForPaymentCompletion(t, client2, payment.Id)
	if completedTransaction.GetStatus() == objects.TransactionStatusFailed {
		t.Errorf("Payment failed: %s", payment.FailureReason.StringValue())
	}
	t.Log(completedTransaction)
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
	servicestest.LoadSeedAsSigningKey(t, env.NodeID2, env.MasterSeedHex2, objects.BitcoinNetworkRegtest, client)
	address, err := client.CreateNodeWalletAddress(env.NodeID2)
	require.NoError(t, err)
	t.Log(address)
}

func createInvoiceForNode(client *services.LightsparkClient, nodeID string) (*objects.Invoice, error) {
	invoice, err := client.CreateInvoice(nodeID, InvoiceAmount, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

func waitForPaymentCompletion(
	t *testing.T, client *services.LightsparkClient, paymentId string) *objects.Transaction {
	payment, err := client.GetEntity(paymentId)
	require.NoError(t, err)
	castPayment, didCast := (*payment).(objects.Transaction)
	require.True(t, didCast)
	startTime := time.Now()
	for castPayment.GetStatus() != objects.TransactionStatusSuccess && castPayment.GetStatus() != objects.TransactionStatusFailed {
		if time.Since(startTime) > time.Minute*3 {
			t.Fatalf("Payment timed out: %s", paymentId)
			return nil
		}
		payment, err = client.GetEntity(paymentId)
		require.NoError(t, err)
		castPayment, didCast = (*payment).(objects.LightningTransaction)
		require.True(t, didCast)
	}
	return &castPayment
}

// Ensure node has enough funds to pay the invoice.
// If not, create an invoice from the routing node and pay it from the node.
// This will only work with REGTEST nodes.
func ensureEnoughNodeFunds(
	t *testing.T, client *services.LightsparkClient, nodeId string, amountMillisatoshis int64) {
	nodeEntity, err := client.GetEntity(nodeId)
	require.NoError(t, err)
	castNode, didCast := (*nodeEntity).(objects.LightsparkNode)
	require.True(t, didCast)
	balanceMilliSats, err := utils.ValueMilliSatoshi(*castNode.GetLocalBalance())
	require.NoError(t, err)
	buffer := int64(math.Round(float64(amountMillisatoshis) * 0.25))
	if balanceMilliSats < amountMillisatoshis+buffer {
		invoice, err := client.CreateInvoice(nodeId, amountMillisatoshis*5, nil, nil, nil)
		require.NoError(t, err)
		payment, err := client.CreateTestModePayment(nodeId, invoice.Data.EncodedPaymentRequest, nil)
		require.NoError(t, err)
		transaction := (*waitForPaymentCompletion(t, client, payment.Id)).(objects.IncomingPayment)
		require.Equal(t, objects.TransactionStatusSuccess, transaction.Status)
	}
}
