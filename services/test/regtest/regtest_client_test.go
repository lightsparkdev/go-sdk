//go:build integration
// +build integration

package regtest_test

import (
	"encoding/hex"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	servicestest "github.com/lightsparkdev/go-sdk/services/test"
	"github.com/lightsparkdev/go-sdk/utils"
	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
	"github.com/stretchr/testify/require"
)

const (
	InvoiceAmount = 1_000_000
	OfferAmount   = 1_000_000
)

func TestCreateInvoice(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(t, client, env.NodeID)
	require.NoError(t, err)
	t.Log(invoice)
}

// Create invoice for node 1 and pay it from routing node. You'll need to run this a few
// times the first time you are funding a node to get enough funds in.
// Note: This will only work with REGTEST nodes.
func TestCreateTestPaymentNode1(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(t, client, env.NodeID)
	require.NoError(t, err)
	t.Logf("Created invoice %v", invoice)
	payment, err := client.CreateTestModePayment(env.NodeID, invoice.Data.EncodedPaymentRequest, nil)
	require.NoError(t, err)
	t.Logf("Created payment %v", payment)
	tx := *waitForPaymentCompletion(t, client, payment.Id)
	if tx != nil {
		require.Equal(t, objects.TransactionStatusSuccess, tx.GetStatus())
		t.Log(tx)
	}
}

// Create invoice for node 2 and pay it from routing node. You'll need to run this a few
// times the first time you are funding a node to get enough funds in.
// Note: This will only work with REGTEST nodes.
func TestCreateTestPaymentNode2(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(t, client, env.NodeID2)
	require.NoError(t, err)
	t.Logf("Created invoice %v", invoice)
	payment, err := client.CreateTestModePayment(env.NodeID2, invoice.Data.EncodedPaymentRequest, nil)
	require.NoError(t, err)
	t.Logf("Created payment %v", payment)
	tx := *waitForPaymentCompletion(t, client, payment.Id)
	if tx != nil {
		require.Equal(t, objects.TransactionStatusSuccess, tx.GetStatus())
		t.Log(tx)
	}
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
	if completedTransaction != nil {
		t.Log(completedTransaction)
	}
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
	if completedTransaction != nil {
		t.Log(completedTransaction)
	}
}

// Create an invoice from node 1, pay it from node 2
func TestPayInvoiceNode2ToNode1(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	invoice, err := createInvoiceForNode(t, client, env.NodeID)
	require.NoError(t, err)

	t.Log(invoice)
	client2 := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	ensureEnoughNodeFunds(t, client2, env.NodeID2, InvoiceAmount)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID2, env.MasterSeedHex2, objects.BitcoinNetworkRegtest, client2)
	payment, err := client2.PayInvoice(env.NodeID2, invoice.Data.EncodedPaymentRequest, 60, InvoiceAmount*0.16, nil)
	require.NoError(t, err)
	completedTransaction := *waitForPaymentCompletion(t, client2, payment.Id)
	if completedTransaction != nil {
		t.Log(completedTransaction)
	}
}

// Create an invoice from node 2 (providing payment hash), pay it from node 1
func TestPayInvoiceNode1ToNode2(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	invoice, err := createInvoiceWithPaymentHashForNode(t, client, env.NodeID2, env.MasterSeedHex2)
	require.NoError(t, err)

	t.Log(invoice)
	client2 := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	ensureEnoughNodeFunds(t, client2, env.NodeID, OfferAmount)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID, env.MasterSeedHex, objects.BitcoinNetworkRegtest, client2)
	payment, err := client2.PayInvoice(env.NodeID, invoice.Data.EncodedPaymentRequest, 60, InvoiceAmount*0.16, nil)
	require.NoError(t, err)

	completedTransaction := *waitForPaymentCompletion(t, client2, payment.Id)
	if completedTransaction != nil {
		t.Log(completedTransaction)
	}
}

func TestPayOfferNode2ToNode1(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	offer, err := createOfferForNode(client, env.NodeID)
	require.NoError(t, err)

	t.Log(offer)
	client2 := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	ensureEnoughNodeFunds(t, client2, env.NodeID2, OfferAmount)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID2, env.MasterSeedHex2, objects.BitcoinNetworkRegtest, client2)
	payment, err := client2.PayOffer(env.NodeID2, offer.EncodedOffer, 60, OfferAmount*0.16, nil, nil)
	require.NoError(t, err)
	completedTransaction := *waitForPaymentCompletion(t, client2, payment.Id)
	if completedTransaction != nil {
		t.Log(completedTransaction)
	}
}

func TestPayOfferNode1ToNode2(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID2, env.ApiClientSecret2, &env.ApiClientEndpoint)
	offer, err := createOfferForNode(client, env.NodeID2)
	require.NoError(t, err)

	t.Log(offer)
	client2 := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	ensureEnoughNodeFunds(t, client2, env.NodeID, OfferAmount)
	servicestest.LoadSeedAsSigningKey(t, env.NodeID, env.MasterSeedHex, objects.BitcoinNetworkRegtest, client2)
	payment, err := client2.PayOffer(env.NodeID, offer.EncodedOffer, 60, OfferAmount*0.16, nil, nil)
	require.NoError(t, err)
	completedTransaction := *waitForPaymentCompletion(t, client2, payment.Id)
	if completedTransaction != nil {
		t.Log(completedTransaction)
	}
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

func createInvoiceForNode(t *testing.T, client *services.LightsparkClient, nodeID string) (*objects.Invoice, error) {
	startTime := time.Now()

	invoice, err := client.CreateInvoice(nodeID, InvoiceAmount, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	t.Logf("Created invoice %s in %s", invoice.Id, time.Since(startTime))

	return invoice, nil
}

func createInvoiceWithPaymentHashForNode(t *testing.T, client *services.LightsparkClient, nodeID string, seedHex string) (*objects.Invoice, error) {
	startTime := time.Now()

	seedBytes, err := hex.DecodeString(seedHex)
	if err != nil {
		return nil, err
	}

	nonce, err := lightspark_crypto.GeneratePreimageNonce(seedBytes)
	if err != nil {
		return nil, err
	}

	paymentHash, err := lightspark_crypto.GeneratePreimageHash(seedBytes, nonce)
	if err != nil {
		return nil, err
	}

	paymentHashHex := hex.EncodeToString(paymentHash)

	invoice, err := client.CreateInvoiceWithPaymentHash(
		nodeID,
		InvoiceAmount,
		nil,
		nil,
		nil,
		&paymentHashHex,
		&nonce,
	)
	if err != nil {
		return nil, err
	}

	t.Logf("Created invoice %s in %s", invoice.Id, time.Since(startTime))

	return invoice, nil
}

func createOfferForNode(client *services.LightsparkClient, nodeId string) (*objects.Offer, error) {
	var offerAmount int64 = OfferAmount

	offer, err := client.CreateOffer(nodeId, &offerAmount, nil)
	if err != nil {
		return nil, err
	}
	return offer, nil
}

func waitForPaymentCompletion(
	t *testing.T, client *services.LightsparkClient, paymentId string) *objects.Transaction {
	payment, err := client.GetEntity(paymentId)
	require.NoError(t, err)
	castPayment, didCast := (*payment).(objects.Transaction)
	require.True(t, didCast)
	startTime := time.Now()
	for castPayment.GetStatus() != objects.TransactionStatusSuccess && castPayment.GetStatus() != objects.TransactionStatusFailed {
		t.Logf("Payment status: %s; Sleeping for 5 seconds before refetching...", castPayment.GetStatus().StringValue())
		time.Sleep(5 * time.Second)
		if time.Since(startTime) > time.Minute*3 {
			t.Fatalf("Payment timed out: %s", paymentId)
			return nil
		}
		payment, err = client.GetEntity(paymentId)
		require.NoError(t, err)
		castPayment, didCast = (*payment).(objects.LightningTransaction)
		require.True(t, didCast)
	}
	if castPayment.GetStatus() == objects.TransactionStatusFailed {
		if reflect.TypeOf(castPayment) == reflect.TypeOf(objects.OutgoingPayment{}) {
			outgoingPayment, ok := castPayment.(objects.OutgoingPayment)
			require.True(t, ok)
			if outgoingPayment.FailureReason != nil {
				t.Errorf("Payment failed due to: %s", outgoingPayment.FailureReason.StringValue())
			} else {
				t.Error("Payment failed with failure reason unavailable.")
			}
		} else {
			t.Error("Payment failed with failure reason unavailable.")
		}
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
	balanceMilliSats := int64(0)
	balances := castNode.GetBalances()
	if balances != nil {
		balanceMilliSats, err = utils.ValueMilliSatoshi(balances.AvailableToSendBalance)
		require.NoError(t, err)
	}
	log.Printf("Check if node id, %s, has enough local balance as of, %d msats, to send, %d msats, before funding.", nodeId, balanceMilliSats, amountMillisatoshis)

	// Requiring 10x the amount to be sent to ensure we have enough to pay the invoice. Note that we really shouldn't
	// need this much of a buffer, but even 2x is proving to be insufficient in some cases. This means our
	// availableToSendBalance is not actually returning what's available to send. See
	// https://linear.app/lightsparkdev/issue/LIG-4098 for more info.
	balanceBufferFactor := int64(10)
	if balanceMilliSats < amountMillisatoshis*balanceBufferFactor {
		invoice, err := client.CreateInvoice(nodeId, amountMillisatoshis*balanceBufferFactor, nil, nil, nil)
		require.NoError(t, err)
		payment, err := client.CreateTestModePayment(nodeId, invoice.Data.EncodedPaymentRequest, nil)
		require.NoError(t, err)
		transaction := (*waitForPaymentCompletion(t, client, payment.Id)).(objects.IncomingPayment)
		require.Equal(t, objects.TransactionStatusSuccess, transaction.Status)

		startTime := time.Now()
		for {
			nodeEntity, err := client.GetEntity(nodeId)
			require.NoError(t, err)
			castNode, didCast := (*nodeEntity).(objects.LightsparkNode)
			require.True(t, didCast)
			balanceMilliSats = int64(0)
			balances := castNode.GetBalances()
			if balances != nil {
				balanceMilliSats, err = utils.ValueMilliSatoshi(balances.AvailableToSendBalance)
				require.NoError(t, err)
			}
			log.Printf("Check if node id, %s, has enough local balance as of, %d msats, to send, %d msats, after funding.", nodeId, balanceMilliSats, amountMillisatoshis)
			if balanceMilliSats >= amountMillisatoshis {
				break
			}
			if time.Since(startTime) > time.Minute*3 {
				t.Fatalf("Funding node id, %s, failed, it still has %d msats balance.", nodeId, balanceMilliSats)
			}
		}
	}
}
