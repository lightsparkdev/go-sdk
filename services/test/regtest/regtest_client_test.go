package regtest_test

import (
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/scripts"
	"github.com/lightsparkdev/go-sdk/services"
	servicestest "github.com/lightsparkdev/go-sdk/services/test"
	"github.com/lightsparkdev/go-sdk/utils"
	"github.com/stretchr/testify/require"
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
	invoice, err := createInvoiceForNode(client, env.NodeID2)
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
	invoice, err := createInvoiceForNode(client, env.NodeID)
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

// TestRipcordMultipleUpdates makes multiple payments, and confirms the ripcord state updates makes sense for each payment
func TestRipcordMultipleUpdates(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)

	variables := map[string]interface{}{
		"node_id": env.NodeID,
		"ripcord_address": "bcrt1qdvqqntu9kxq996ur477un3z4y8z4yky0ex3yye", // random test regtest bitcoin wallet address
	}
	_, err := client.Requester.ExecuteGraphql(scripts.SET_RIPCORD_ADDRESS_SPARKNODE_MUTATION, variables, nil)
	if err != nil {
		t.Fatalf("Failed to set ripcord address on the sparknode. %v\n", err)
	}

	// We initially get a snapshot of the state of the updates table because this can vary depending on how and when this test runs.
	// The main things we care about are the channel point and commitment number which can uniquely identify succesful state updates
	nodeEntity, err := client.GetEntity(env.NodeID)
	require.NoError(t, err)
	castNode, didCast := (*nodeEntity).(objects.LightsparkNodeWithRemoteSigning)
	require.True(t, didCast)
	limit := 1
	updates, err := castNode.GetRipcordUpdates(client.Requester, &limit, nil)
	require.NoError(t, err)
	initialUpdateLength := len(updates.Entities)
	var initialChannelPoint string
	var initialCommitmentNumber int
	if initialUpdateLength > 0 {
		initialChannelPoint = *updates.Entities[0].Channel.ChannelPoint
		initialCommitmentNumber = *updates.Entities[0].CommitmentNumber
	}

	// Payment 1
	invoice, err := createInvoiceForNode(client, env.NodeID)
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

	var lastUpdate objects.RipcordUpdate
	compareRipcordUpdates := func(initialChannelPoint string, initialCommitmentNumber int) {
		// If the initial snapshot of the ripcord state is not empty, which is likely to be the case, we continuously poll
		// for the latest updates until either the channel point is different or the commitment number is different. Right
		// now we cannot force the payment through a specific channel easily. As such, if the channel points are not the
		// same, we know that we do not need to compare the two updates, we can simply do a status check and return indicating
		// that the ripcord update was successful. If the channel point stays the same, we break out of the below loop when
		// the commitment number changes, indicating that a new commitment was received. We then compare the commitment numbers
		// of the previous two states and check if they are related in the correct way. If no new update happens, this section
		// will timeout the test, causing it to fail. This polling section seems sturdier than sleeping a default amount.
		updates, err = castNode.GetRipcordUpdates(client.Requester, &limit, nil)
		require.NoError(t, err)
		lastUpdate = updates.Entities[0]
		if lastUpdate.Channel.ChannelPoint == nil {
			t.Fatalf("Channel point should be set.")
		}
		if lastUpdate.CommitmentNumber == nil {
			t.Fatalf("Commitment number should be set.")
		}
		for initialChannelPoint == *lastUpdate.Channel.ChannelPoint && initialCommitmentNumber == *lastUpdate.CommitmentNumber {
			updates, err = castNode.GetRipcordUpdates(client.Requester, &limit, nil)
			require.NoError(t, err)
			lastUpdate = updates.Entities[0]
		}
		if initialChannelPoint == *lastUpdate.Channel.ChannelPoint {
			if *lastUpdate.CommitmentNumber >= initialCommitmentNumber {
				t.Fatalf("Commitment Numbers for newer states should be strictly decreasing.")
			}
			if initialCommitmentNumber - *lastUpdate.CommitmentNumber > 2 {
				t.Fatalf("These should not differ by more than 2. There will be two new updates every payment, one for htlc and one for the main payment.")
			}
		}
		if *lastUpdate.RipcordUpdateStatus != "SUCCESS" {
			t.Fatalf("Ripcord Update status should be successful to indicate proper storing in the db.")
		}
	}
	if initialUpdateLength == 0 {
		// If the initial snapshot of the ripcord state for this node is nonexistent, we just wait and poll until there
		// exists updates. When there is an update that exists, we wait to see if it is the status we expect.
		updates, err = castNode.GetRipcordUpdates(client.Requester, &limit, nil)
		require.NoError(t, err)
		
		for len(updates.Entities) == 0 {
			nodeEntity, err := client.GetEntity(env.NodeID)
			require.NoError(t, err)
			castNode, didCast = (*nodeEntity).(objects.LightsparkNodeWithRemoteSigning)
			require.True(t, didCast)
			updates, err = castNode.GetRipcordUpdates(client.Requester, &limit, nil)
			require.NoError(t, err)
			time.Sleep(100 * time.Millisecond)
		}
		lastUpdate = updates.Entities[0]
		if lastUpdate.RipcordUpdateStatus == nil {
			t.Fatalf("Ripcord Update Status should be set.")
		}
		if *lastUpdate.RipcordUpdateStatus != "SUCCESS" {
			t.Fatalf("Ripcord Update status should be successful to indicate proper storing in the db.")
		}
	} else {
		compareRipcordUpdates(initialChannelPoint, initialCommitmentNumber)
	}

	// Once again, after the first round of payments we get the intermediate state information to compare with the second round
	// of payments.
	initialCommitmentNumber = *lastUpdate.CommitmentNumber
	initialChannelPoint = *lastUpdate.Channel.ChannelPoint

	// Payment 2
	invoice, err = createInvoiceForNode(client, env.NodeID)
	require.NoError(t, err)
	t.Logf("Created invoice %v", invoice)
	payment, err = client.CreateTestModePayment(env.NodeID, invoice.Data.EncodedPaymentRequest, nil)
	require.NoError(t, err)
	t.Logf("Created payment %v", payment)
	tx = *waitForPaymentCompletion(t, client, payment.Id)
	if tx != nil {
		require.Equal(t, objects.TransactionStatusSuccess, tx.GetStatus())
		t.Log(tx)
	}

	compareRipcordUpdates(initialChannelPoint, initialCommitmentNumber)
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
