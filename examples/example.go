// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
)

func main() {
	// MODIFY THOSE VARIABLES BEFORE RUNNING THE EXAMPLE
	//
	// We defined those variables as environment variables, but if you are just
	// running the example locally, feel free to just set the values in code.
	//
	// First, initialize your client ID and client secret. Those are available
	// in your account at https://app.lightspark.com/api_config
	apiClientID := os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_ID")
	apiToken := os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_SECRET")
	baseUrl := os.Getenv("LIGHTSPARK_EXAMPLE_BASE_URL")
	client := services.NewLightsparkClient(apiClientID, apiToken, &baseUrl)

	nodeId := os.Getenv("LIGHTSPARK_TEST_NODE_ID")
	nodePassword := os.Getenv("LIGHTSPARK_TEST_NODE_PASSWORD")

	// Get current account
	fmt.Println("Getting current account...")
	account, err := client.GetCurrentAccount()
	if err != nil {
		fmt.Printf("get current account failed: %v", err)
		return
	}
	fmt.Printf("Your account name is %v.\n", *account.Name)
	fmt.Println()

	// Check your account's conductivity on REGTEST
	networks := []objects.BitcoinNetwork{objects.BitcoinNetworkRegtest}
	conductivity, err := account.GetConductivity(client.Requester, &networks, nil)
	if err != nil {
		fmt.Printf("get account conductivity failed: %v", err)
		return
	}
	fmt.Printf("Your account's conductivity on REGTEST is %v.\n", *conductivity)
	fmt.Println()

	// Check your account's local and remote balances on REGTEST
	localBalance, err := account.GetLocalBalance(client.Requester, &networks, nil)
	if err != nil {
		fmt.Printf("get local balance failed: %v", err)
		return
	}
	fmt.Printf("Your local balance on REGTEST is %v.\n", localBalance.OriginalValue)
	remoteBalance, err := account.GetRemoteBalance(client.Requester, &networks, nil)
	if err != nil {
		fmt.Printf("get remote balance failed: %v", err)
		return
	}
	fmt.Printf("Your remote balance on REGTEST is %v.\n", remoteBalance.OriginalValue)
	fmt.Println()

	// Check your nodes on REGTEST
	var count int64 = 50
	nodesConnection, err := account.GetNodes(client.Requester, &count, &networks, nil, nil)
	if err != nil {
		fmt.Printf("get nodes failed: %v", err)
		return
	}
	fmt.Printf("You have %v nodes in total.\n", nodesConnection.Count)
	for i, node := range nodesConnection.Entities {
		fmt.Printf("#%v: %v with id %v\n", i, node.GetDisplayName(), node.GetId())
	}
	fmt.Println()

	// Check your transactions on REGTEST
	network := objects.BitcoinNetworkRegtest
	transactionsConnection, err := account.GetTransactions(
		client.Requester,
		&count,   // first
		nil,      //after
		nil,      // types
		nil,      // after_date
		nil,      // before_date
		&network, // bitcoin_network
		nil,      // lightning_node_id
		nil,      // statuses
		nil,      //exclude_failures
	)
	if err != nil {
		fmt.Printf("get transactions failed: %v", err)
		return
	}
	fmt.Printf("You have %v transactions in total.\n", transactionsConnection.Count)
	var transactionId string
	for _, transaction := range transactionsConnection.Entities {
		transactionId = transaction.GetId()
		fmt.Printf(
			"    - %v at %v: %v %v (%v)\n",
			transaction.GetId(),
			transaction.GetCreatedAt(),
			transaction.GetAmount().OriginalValue,
			transaction.GetAmount().OriginalUnit.StringValue(),
			transaction.GetStatus().StringValue(),
		)
	}
	fmt.Println()

	// Fetch a transaction by id
	fmt.Println("Getting entity...")
	entity, err := client.GetEntity(transactionId)
	if err != nil {
		fmt.Printf("get entity failed: %v", err)
		return
	}
	fmt.Printf("We fetched transaction %v, created at %v.\n", (*entity).GetId(), (*entity).GetCreatedAt())
	fmt.Println()

	// Fetch transactions on REGTEST using pagination
	var pageSize int64 = 10
	var iterations int64 = 0
	hasNext := true
	var after *string
	for hasNext && iterations < 30 {
		iterations = iterations + 1
		transactionsConnection, err = account.GetTransactions(
			client.Requester,
			&pageSize, // first
			after,     //after
			nil,       // types
			nil,       // after_date
			nil,       // before_date
			&network,  // bitcoin_network
			nil,       // lightning_node_id
			nil,       // statuses
			nil,       //exclude_failures
		)
		fmt.Printf("We got %v transactions for the page (iteration #%v)\n", transactionsConnection.Count, iterations)
		if *transactionsConnection.PageInfo.HasNextPage {
			hasNext = true
			after = transactionsConnection.PageInfo.EndCursor
			fmt.Println("  And we have another page!")
		} else {
			hasNext = false
			fmt.Println("  And we're done!")
		}
	}
	fmt.Println()

	// Get the transactions that happened in the past day on REGTEST
	afterDate := time.Now().UTC().Add(-time.Hour * 24)
	transactionsConnection, err = account.GetTransactions(
		client.Requester,
		&count,     // first
		nil,        //after
		nil,        // types
		&afterDate, // after_date
		nil,        // before_date
		&network,   // bitcoin_network
		nil,        // lightning_node_id
		nil,        // statuses
		nil,        //exclude_failures
	)
	fmt.Printf("We had %v transactions in the past 24 hours.", transactionsConnection.Count)
	fmt.Println()

	apiTokenConnection, err := account.GetApiTokens(client.Requester, nil, nil)
	if err != nil {
		fmt.Printf("get api tokens failed: %v", err)
		return
	}
	fmt.Printf("You initially have %v active API tokens.\n", apiTokenConnection.Count)
	fmt.Println()

	// Create api token
	fmt.Println("Creating API token...")
	apiTokenOutput, err := client.CreateApiToken("Test", false, true)
	if err != nil {
		fmt.Printf("create api token failed: %v", err)
		return
	}
	fmt.Println("Your API token:")
	fmt.Printf("    client id: %v\n", apiTokenOutput.ApiToken.ClientId)
	fmt.Printf("    name: %v\n", apiTokenOutput.ApiToken.ClientId)
	fmt.Printf("    permissions: %v\n", apiTokenOutput.ApiToken.Permissions)
	fmt.Printf("Your API secret: %v\n", apiTokenOutput.ClientSecret)
	fmt.Println()

	apiTokenConnection, err = account.GetApiTokens(client.Requester, nil, nil)
	if err != nil {
		fmt.Printf("get api tokens failed: %v", err)
		return
	}
	fmt.Printf("You now have %v active API tokens.\n", apiTokenConnection.Count)
	fmt.Println()

	// Delete api token
	fmt.Println("Deleting API token...")
	err = client.DeleteApiToken(apiTokenOutput.ApiToken.Id)
	if err != nil {
		fmt.Printf("delete api token failed: %v", err)
		return
	}
	fmt.Println("API token deleted.")
	fmt.Println()

	apiTokenConnection, err = account.GetApiTokens(client.Requester, nil, nil)
	if err != nil {
		fmt.Printf("get api tokens failed: %v", err)
		return
	}
	fmt.Printf("You now have %v active API tokens.\n", apiTokenConnection.Count)
	fmt.Println()

	// Get some fee estimates for L1 transactions
	fmt.Println("Getting L1 transaction fee estimates...")
	feeEstimate, err := client.GetBitcoinFeeEstimate(objects.BitcoinNetworkMainnet)
	if err != nil {
		fmt.Printf("get bitcoin fee estimate failed: %v", err)
		return
	}
	fmt.Printf("Fees for a fast transaction %v %v.\n", feeEstimate.FeeFast.OriginalValue, feeEstimate.FeeFast.OriginalUnit.StringValue())
	fmt.Printf("Fees for a cheap transaction %v %v.\n", feeEstimate.FeeMin.OriginalValue, feeEstimate.FeeMin.OriginalUnit.StringValue())
	fmt.Println()

	// Create an L1 address
	fmt.Println("Creating an L1 address...")
	address, err := client.CreateNodeWalletAddress(nodeId)
	if err != nil {
		fmt.Printf("get node wallet failed: %v", err)
		return
	}
	fmt.Printf("Node wallet address created: %v\n", address)
	fmt.Println()

	client.LoadNodeSigningKey(nodeId, *services.NewSigningKeyLoaderFromNodeIdAndPassword(nodeId, nodePassword))

	// Pay an invoice
	fmt.Println("Paying an invoice...")
	encodedInvoice := "<your encoded invoice>"
	// When testing paying invoice in test mode, a test invoice can be generated
	fmt.Println("Creating a test mode invoice...")
	testInvoice, err := client.CreateTestModeInvoice(nodeId, 100000, nil, nil)
	if err != nil {
		fmt.Printf("create test invoice failed: %v", err)
		return
	}
	encodedInvoice = *testInvoice
	fmt.Printf("Test invoice created: %v\n", *testInvoice)
	fmt.Println()

	outgoingPayment, err := client.PayInvoice(nodeId, encodedInvoice, 1000, 60, nil)
	if err != nil {
		fmt.Printf("pay invoice failed: %v", err)
		return
	}
	fmt.Printf("Invoice paid with payment id: %v\n", outgoingPayment.Id)
	fmt.Println()

	// Decode an encoded invoice
	fmt.Println("Decoding an encoded invoice...")
	decodedPaymentRequest, err := client.DecodePaymentRequest(encodedInvoice)
	if err != nil {
		fmt.Printf("decode invoice failed: %v", err)
		return
	}
	decodedInvoice, ok := (*decodedPaymentRequest).(objects.InvoiceData)
	if !ok {
		fmt.Printf("casting payment request to invoice failed")
		return
	}
	destinationNodePublicKey := *(decodedInvoice.Destination.GetPublicKey())
	fmt.Println("Decoded invoice...")
	fmt.Printf("    destination public key: %v\n", destinationNodePublicKey)
	fmt.Printf("    amount: %v %v\n", decodedInvoice.Amount.OriginalValue, decodedInvoice.Amount.OriginalUnit.StringValue())
	fmt.Println("")

	// Send a payment
	fmt.Println("Sending a payment...")
	outgoingPayment, err = client.SendPayment(nodeId, destinationNodePublicKey, 100, 1000, 60)
	if err != nil {
		fmt.Printf("send payment failed: %v", err)
		return
	}
	fmt.Printf("Payment sent with payment id: %v\n", outgoingPayment.Id)
	fmt.Println()

	// Create an invoice
	fmt.Println("Creating an invoice...")
	invoice, err := client.CreateInvoice(nodeId, 100000, nil, nil, nil)
	encodedInvoice = invoice.Data.EncodedPaymentRequest
	if err != nil {
		fmt.Printf("create invoice failed: %v", err)
		return
	}
	fmt.Printf("Invoice created: %v\n", invoice.Data.EncodedPaymentRequest)
	fmt.Println()

	// Get fee estimate for a node
	fmt.Println("Getting fee estimate for a node...")
	nodeFeeEstimate, err := client.GetLightningFeeEstimateForNode(nodeId, destinationNodePublicKey, 1000)
	if err != nil {
		fmt.Printf("getting fee estimate for node failed: %v", err)
		return
	}
	fmt.Printf("Estimated fee for the node: %v %v\n", nodeFeeEstimate.FeeEstimate.OriginalValue, nodeFeeEstimate.FeeEstimate.OriginalUnit.StringValue())
	fmt.Println()

	// Get fee estimate for an invoice
	fmt.Println("Getting fee estimate for an invoice...")
	invoiceFeeEstimate, err := client.GetLightningFeeEstimateForInvoice(nodeId, encodedInvoice, nil)
	if err != nil {
		fmt.Printf("getting fee estimate for invoice failed: %v", err)
		return
	}
	fmt.Printf("Estimated fee for the invoice: %v %v\n", invoiceFeeEstimate.FeeEstimate.OriginalValue, invoiceFeeEstimate.FeeEstimate.OriginalUnit.StringValue())
	fmt.Println()

	// If the node is in test mode, CreateTestModePayment can simulate a payment to the invoice
	testPayment, err := client.CreateTestModePayment(nodeId, encodedInvoice, nil)
	if err != nil {
		fmt.Printf("simulating a test mode payment failed: %v", err)
		return
	}
	fmt.Printf("Invoice paid with a simulated payment %v\n", testPayment.Id)
	fmt.Println()

	// Create and cancel an invoice
	fmt.Println("Creating an invoice...")
	invoiceToCancel, err := client.CreateInvoice(nodeId, 100000, nil, nil, nil)
	if err != nil {
		fmt.Printf("create invoice failed: %v", err)
		return
	}
	fmt.Printf("Cancelling Invoice: %v\n", invoice.Data.EncodedPaymentRequest)
	canceledInvoice, err := client.CancelInvoice(invoiceToCancel.Id)
	if err != nil {
		fmt.Printf("cancel invoice failed: %v", err)
		return
	}
	fmt.Printf("Invoice canceled: %v\n", canceledInvoice.Id)
	fmt.Println()

	// Withdraw funds
	fmt.Println("Withdraw funds...")
	bitcoinAddress := "<your bitcoin address for withdrawal>"
	withdrawalRequest, err := client.RequestWithdrawal(nodeId, 100000, bitcoinAddress, objects.WithdrawalModeWalletThenChannels)
	if err != nil {
		fmt.Printf("withdraw failed: %v", err)
		return
	}
	fmt.Printf("Withdrawal initiated with request id: %v\n", withdrawalRequest.Id)
	fmt.Println()

	// Fund a node
	fmt.Println("Funding a node...")
	amountFunded, err := client.FundNode(nodeId, 1000000)
	if err != nil {
		fmt.Printf("fund node failed: %v", err)
		return
	}
	fmt.Printf("Amount funded: %v %v\n", amountFunded.OriginalValue, amountFunded.OriginalUnit.StringValue())
	fmt.Println()

	// Screen a lightning node
	fmt.Println("Screening lightning node...")
	provider := objects.ComplianceProviderChainalysis
	nodePubkey := "bc1qj4mfcgej3wxp8eundzq7sq8f80wps02kk38sgadrer39mr5l7ncqrgmp89"
	rating, err := client.ScreenNode(provider, nodePubkey)
	if err != nil {
		fmt.Printf("screening lightning node failed: %v", err)
		return
	}
	fmt.Printf("The risk rating is %v.\n", *rating)
	fmt.Println()

	// Register a successful payment
	fmt.Println("Registering a successful payment...")
	provider = objects.ComplianceProviderChainalysis
	nodePubkey = "bc1qj4mfcgej3wxp8eundzq7sq8f80wps02kk38sgadrer39mr5l7ncqrgmp89"
	paymentId := "<Your outgoing payment id>"
	direction := objects.PaymentDirectionSent
	err = client.RegisterPayment(provider, paymentId, nodePubkey, direction)
	if err != nil {
		fmt.Printf("Registering payment failed: %v", err)
		return
	}
	fmt.Printf("Payment was successfully registered.\n")

	// Run a custom query
	fmt.Println("Run a custom query...")
	response, err := client.ExecuteGraphqlRequest(
		`query MyCustomQuery($network: BitcoinNetwork!) {
			current_account {
			  id
			  conductivity(bitcoin_networks: [$network])
			}
		}`,
		map[string]interface{}{"network": objects.BitcoinNetworkRegtest},
	)
	if err != nil {
		fmt.Printf("execute graphql request failed: %v", err)
		return
	}
	accountMap := response["current_account"].(map[string]interface{})
	conductivityValue := int(accountMap["conductivity"].(float64))
	fmt.Printf("Your account conductivity is %v.\n", conductivityValue)
	fmt.Println()
}
