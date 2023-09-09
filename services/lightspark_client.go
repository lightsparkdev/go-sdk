// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package services

import (
	"encoding/json"
	"errors"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/requester"
	"github.com/lightsparkdev/go-sdk/scripts"
	"github.com/lightsparkdev/go-sdk/utils"
)

type LightsparkClient struct {
	Requester *requester.Requester
}

// NewLightsparkClient creates a new LightsparkClient instance
//
// Args:
//
//	apiTokenClientId: the client id of the API token
//	apiTokenClientSecret: the client secret of the API token
//	baseUrl: the base url of the Lightspark API
func NewLightsparkClient(apiTokenClientId string, apiTokenClientSecret string,
	baseUrl *string) *LightsparkClient {

	gqlRequester := &requester.Requester{
		ApiTokenClientId:     apiTokenClientId,
		ApiTokenClientSecret: apiTokenClientSecret,
		BaseUrl:              baseUrl,
	}
	return &LightsparkClient{Requester: gqlRequester}
}

// CreateApiToken creates a new API token that can be used to authenticate requests
// for this account when using the Lightspark APIs and SDKs.
//
// Args:
//
//	name: the name of the API token
//	transact: whether the API token should be able to used to initiate transactions
//	testMode: whether the API token should be created for test mode or Mainnet mode
func (client *LightsparkClient) CreateApiToken(name string, transact bool,
	testMode bool) (*scripts.CreateApiTokenOutput, error) {

	permissions := []objects.Permission{}
	if transact && testMode {
		permissions = append(permissions, objects.PermissionRegtestView)
		permissions = append(permissions, objects.PermissionRegtestTransact)
	} else if transact && !testMode {
		permissions = append(permissions, objects.PermissionMainnetView)
		permissions = append(permissions, objects.PermissionMainnetTransact)
	} else if !transact && testMode {
		permissions = append(permissions, objects.PermissionRegtestView)
	} else {
		permissions = append(permissions, objects.PermissionMainnetView)
	}

	variables := map[string]interface{}{
		"name":        name,
		"permissions": permissions,
	}
	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_API_TOKEN_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["create_api_token"].(map[string]interface{})
	var apiToken objects.ApiToken
	apiTokenJson, err := json.Marshal(output["api_token"].(map[string]interface{}))
	if err != nil {
		return nil, errors.New("error parsing api token")
	}
	json.Unmarshal(apiTokenJson, &apiToken)
	return &scripts.CreateApiTokenOutput{ApiToken: &apiToken, ClientSecret: output["client_secret"].(string)}, nil
}

// CreateInvoice generates a Lightning Invoice (follows the Bolt 11 specification)
// to request a payment from another Lightning Node.
//
// Args:
//
//	nodeId: the id of the node that should be paid
//	amountMsats: the amount of the invoice in millisatoshis
//	memo: the memo of the invoice
//	invoiceType: the type of the invoice
//	expirySecs: the expiry of the invoice in seconds. Default value is 86400 (1 day).
func (client *LightsparkClient) CreateInvoice(nodeId string, amountMsats int64,
	memo *string, invoiceType *objects.InvoiceType, expirySecs *int32) (*objects.Invoice, error) {

	variables := map[string]interface{}{
		"amount_msats": amountMsats,
		"node_id":      nodeId,
		"memo":         memo,
		"invoice_type": invoiceType,
	}
	if expirySecs != nil {
		variables["expiry_secs"] = expirySecs
	}
	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_INVOICE_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["create_invoice"].(map[string]interface{})
	var invoice objects.Invoice
	invoiceJson, err := json.Marshal(output["invoice"].(map[string]interface{}))
	if err != nil {
		return nil, errors.New("error parsing invoice")
	}
	json.Unmarshal(invoiceJson, &invoice)
	return &invoice, nil
}

// CreateLnurlInvoice creates a new LNURL invoice. The metadata is hashed and included in the invoice.
// This API generates a Lightning Invoice (follows the Bolt 11 specification) to request a payment
// from another Lightning Node. This should only be used for generating invoices
// for LNURLs, with `create_invoice` preferred in the general case.
//
// Args:
//
//		nodeId: the id of the node that should be paid
//		amountMsats: the amount of the invoice in millisatoshis
//		metadata: the metadata to include with the invoice
//	 expirySecs: the expiry of the invoice in seconds. Default value is 86400 (1 day)
func (client *LightsparkClient) CreateLnurlInvoice(nodeId string, amountMsats int64,
	metadata string, expirySecs *int32) (*objects.Invoice, error) {

	variables := map[string]interface{}{
		"amount_msats":  amountMsats,
		"node_id":       nodeId,
		"metadata_hash": utils.Sha256HexString(metadata),
	}
	if expirySecs != nil {
		variables["expiry_secs"] = expirySecs
	}
	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_LNURL_INVOICE_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["create_lnurl_invoice"].(map[string]interface{})
	var invoice objects.Invoice
	invoiceJson, err := json.Marshal(output["invoice"].(map[string]interface{}))
	if err != nil {
		return nil, errors.New("error parsing invoice")
	}
	json.Unmarshal(invoiceJson, &invoice)
	return &invoice, nil
}

// CreateUmaInvoice creates a new invoice for the UMA protocol. The metadata is hashed and included in the invoice.
// This API generates a Lightning Invoice (follows the Bolt 11 specification) to request a payment
// from another Lightning Node. This should only be used for generating invoices for UMA, with `create_invoice`
// preferred in the general case.
//
// Args:
//
//		nodeId: the id of the node that should be paid
//		amountMsats: the amount of the invoice in millisatoshis
//		metadata: the metadata to include with the invoice
//	 	expirySecs: the expiry of the invoice in seconds. Default value is 86400 (1 day)
func (client *LightsparkClient) CreateUmaInvoice(nodeId string, amountMsats int64,
	metadata string, expirySecs *int32) (*objects.Invoice, error) {

	variables := map[string]interface{}{
		"amount_msats":  amountMsats,
		"node_id":       nodeId,
		"metadata_hash": utils.Sha256HexString(metadata),
	}
	if expirySecs != nil {
		variables["expiry_secs"] = expirySecs
	}
	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_UMA_INVOICE_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["create_uma_invoice"].(map[string]interface{})
	var invoice objects.Invoice
	invoiceJson, err := json.Marshal(output["invoice"].(map[string]interface{}))
	if err != nil {
		return nil, errors.New("error parsing invoice")
	}
	json.Unmarshal(invoiceJson, &invoice)
	return &invoice, nil
}

// CreateNodeWalletAddress creates a Bitcoin address for the wallet associated with
// your Lightning Node. You can use this address to send funds to your node. It is
// a best practice to generate a new wallet address every time you need to send money.
// You can generate as many wallet addresses as you want.
func (client *LightsparkClient) CreateNodeWalletAddress(nodeId string) (string, error) {
	variables := map[string]interface{}{
		"node_id": nodeId,
	}
	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_NODE_WALLET_ADDRESS_MUTATION, variables)
	if err != nil {
		return "", err
	}
	output := response["create_node_wallet_address"].(map[string]interface{})
	walletAddress := output["wallet_address"].(string)
	return walletAddress, nil
}

// In test mode, CreateTestModeInvoice generates a Lightning Invoice which can be paid by a local node.
// This is useful for testing your integration with Lightspark.
//
// Args:
//
//	localNodeId: the id of the node that should be paid
//	amountMsats: the amount of the invoice in millisatoshis
//	memo: the memo of the invoice
//	invoiceType: the type of the invoice
func (client *LightsparkClient) CreateTestModeInvoice(localNodeId string, amountMsats int64,
	memo *string, invoiceType *objects.InvoiceType) (*string, error) {

	variables := map[string]interface{}{
		"amount_msats":  amountMsats,
		"local_node_id": localNodeId,
		"memo":          memo,
		"invoice_type":  invoiceType,
	}
	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_TEST_MODE_INVOICE_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["create_test_mode_invoice"].(map[string]interface{})
	encodedInvoice := output["encoded_payment_request"].(string)
	return &encodedInvoice, nil
}

// In test mode, CreateTestModePayment simulates a payment from other node to an invoice.
// This is useful for testing your integration with Lightspark.
//
// Args:
//
//	localNodeId: The node to where you want to send the payment.
//	encodedInvoice: The invoice you want to be paid (as defined by the BOLT11 standard).
//	amountMsats: The amount you will be paid for this invoice, expressed in msats.
//		It should ONLY be set when the invoice amount is zero.
func (client *LightsparkClient) CreateTestModePayment(localNodeId string,
	encodedInvoice string, amountMsats *int64) (*objects.OutgoingPayment, error) {

	variables := map[string]interface{}{
		"local_node_id":   localNodeId,
		"encoded_invoice": encodedInvoice,
	}
	if amountMsats != nil {
		variables["amount_msats"] = amountMsats
	}

	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_TEST_MODE_PAYMENT_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["create_test_mode_payment"].(map[string]interface{})
	var payment objects.OutgoingPayment
	paymentJson, err := json.Marshal(output["payment"].(map[string]interface{}))
	if err != nil {
		return nil, errors.New("error parsing payment")
	}
	json.Unmarshal(paymentJson, &payment)
	return &payment, nil
}

// DecodePaymentRequest decodes the content of an encoded payment request into
// structured data that can be used by the client.
//
// Args:
//
//	encodedPaymentRequest: The encoded payment request.
func (client *LightsparkClient) DecodePaymentRequest(encodedPaymentRequest string) (*objects.PaymentRequestData, error) {
	variables := map[string]interface{}{"encoded_payment_request": encodedPaymentRequest}
	response, err := client.Requester.ExecuteGraphql(scripts.DECODE_PAYMENT_REQUEST_QUERY, variables)
	if err != nil {
		return nil, err
	}

	output := response["decoded_payment_request"].(map[string]interface{})
	paymentRequest, err := objects.PaymentRequestDataUnmarshal(output)
	if err != nil {
		return nil, err
	}
	return &paymentRequest, nil
}

// DeleteApiToken deletes an existing API token from this account.
//
// Args:
//
//	apiTokenId: The id of the API token to delete.
func (client *LightsparkClient) DeleteApiToken(apiTokenId string) error {
	variables := map[string]interface{}{
		"api_token_id": apiTokenId,
	}
	_, err := client.Requester.ExecuteGraphql(scripts.DELETE_API_TOKEN_MUTATION, variables)
	return err
}

// FundNode adds funds to a Lightspark node on the REGTEST network.
// If the amount is not specified, 10,000,000 SATOSHI will be added.
// This API only functions for nodes created on the REGTEST network
// and will return an error when called for any non-REGTEST node.
//
// Args:
//
//	nodeId: The id of the node to fund.
//	amountSats: The amount of funds to add to the node, in SATOSHI.
func (client *LightsparkClient) FundNode(nodeId string, amountSats int64) (
	*objects.CurrencyAmount, error) {

	variables := map[string]interface{}{
		"node_id":     nodeId,
		"amount_sats": amountSats,
	}

	response, err := client.Requester.ExecuteGraphql(scripts.FUND_NODE_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["fund_node"].(map[string]interface{})
	var amount objects.CurrencyAmount
	amountJson, err := json.Marshal(output["amount"].(map[string]interface{}))
	if err != nil {
		return nil, errors.New("error parsing amount")
	}
	json.Unmarshal(amountJson, &amount)
	return &amount, nil
}

// GetBitcoinFeeEstimate returns an estimate of the fees of a transaction on the Bitcoin Network.
//
// Args:
//
//	bitcoinNetwork: The Bitcoin network to use for the estimate.
func (client *LightsparkClient) GetBitcoinFeeEstimate(
	bitcoinNetwork objects.BitcoinNetwork) (*objects.FeeEstimate, error) {

	variables := map[string]interface{}{"bitcoin_network": bitcoinNetwork}
	response, err := client.Requester.ExecuteGraphql(scripts.BITCOIN_FEE_ESTIMATE_QUERY, variables)
	if err != nil {
		return nil, err
	}

	output := response["bitcoin_fee_estimate"].(map[string]interface{})
	var feeEstimate objects.FeeEstimate
	feeEstimateJson, err := json.Marshal(output)
	if err != nil {
		return nil, errors.New("error parsing fee estimate")
	}
	json.Unmarshal(feeEstimateJson, &feeEstimate)
	return &feeEstimate, nil
}

// GetCurrentAccount returns the current connected account.
func (client *LightsparkClient) GetCurrentAccount() (*objects.Account, error) {
	variables := map[string]interface{}{}
	response, err := client.Requester.ExecuteGraphql(scripts.CURRENT_ACCOUNT_QUERY, variables)
	if err != nil {
		return nil, err
	}

	output := response["current_account"].(map[string]interface{})
	var account objects.Account
	accountJson, err := json.Marshal(output)
	if err != nil {
		return nil, errors.New("error parsing account")
	}
	json.Unmarshal(accountJson, &account)
	return &account, nil
}

// GetLightningFeeEstimateForInvoice returns an estimate of the fees that
// will be paid for a Lightning invoice.
//
// Args:
//
//	nodeId: The node from where you want to send the payment
//	encodedInvoice: The invoice you want to pay (as defined by the BOLT11 standard).
//	amountMsats: If the invoice does not specify a payment amount,
//		then the amount that you wish to pay, expressed in msats.
func (client *LightsparkClient) GetLightningFeeEstimateForInvoice(nodeId string,
	encodedInvoice string, amountMsats *int64) (*objects.LightningFeeEstimateOutput, error) {

	variables := map[string]interface{}{
		"node_id":                 nodeId,
		"encoded_payment_request": encodedInvoice,
		"amount_msats":            amountMsats,
	}
	response, err := client.Requester.ExecuteGraphql(scripts.LIGHTNING_FEE_ESTIMATE_FOR_INVOICE_QUERY, variables)
	if err != nil {
		return nil, err
	}

	output := response["lightning_fee_estimate_for_invoice"].(map[string]interface{})
	var feeEstimate objects.LightningFeeEstimateOutput
	feeEstimateJson, err := json.Marshal(output)
	if err != nil {
		return nil, errors.New("error parsing fee estimate")
	}
	json.Unmarshal(feeEstimateJson, &feeEstimate)
	return &feeEstimate, nil
}

// GetLightningFeeEstimateForNode returns an estimate of the fees that will be
// paid to send a payment to another Lightning node.
//
// Args:
//
//	nodeId: The node from where you want to send the payment.
//	destinationNodePublicKey: The public key of the node that you want to pay.
//	amountMsats: The amount that you wish to pay, expressed in msats.
func (client *LightsparkClient) GetLightningFeeEstimateForNode(nodeId string,
	destinationNodePublicKey string, amountMsats int64) (*objects.LightningFeeEstimateOutput, error) {

	variables := map[string]interface{}{
		"node_id":                     nodeId,
		"destination_node_public_key": destinationNodePublicKey,
		"amount_msats":                amountMsats,
	}
	response, err := client.Requester.ExecuteGraphql(scripts.LIGHTNING_FEE_ESTIMATE_FOR_NODE_QUERY, variables)
	if err != nil {
		return nil, err
	}

	output := response["lightning_fee_estimate_for_node"].(map[string]interface{})
	var feeEstimate objects.LightningFeeEstimateOutput
	feeEstimateJson, err := json.Marshal(output)
	if err != nil {
		return nil, errors.New("error parsing fee estimate")
	}
	json.Unmarshal(feeEstimateJson, &feeEstimate)
	return &feeEstimate, nil
}

// PayInvoice sends a payment to a node on the Lightning Network, based on the invoice
// (as defined by the BOLT11 specification) that you provide.
// If you are in test mode, the invoice has to be generated by create_test_mode_invoice mutation.
//
// Args:
//
//	nodeId: The node from where you want to send the payment.
//	encodedInvoice: The invoice you want to pay (as defined by the BOLT11 standard).
//	timeoutSecs: The number of seconds that you are willing to wait for the payment to complete.
//	maximumFeesMsats: The maximum amount of fees that you are willing to pay for this payment, expressed in mSATs.
//	amountMsats: The amount you will pay for this invoice, expressed in msats.
//		It should ONLY be set when the invoice amount is zero.
func (client *LightsparkClient) PayInvoice(nodeId string, encodedInvoice string,
	timeoutSecs int, maximumFeesMsats int64, amountMsats *int64) (*objects.OutgoingPayment, error) {

	variables := map[string]interface{}{
		"node_id":            nodeId,
		"encoded_invoice":    encodedInvoice,
		"timeout_secs":       timeoutSecs,
		"maximum_fees_msats": maximumFeesMsats,
	}
	if amountMsats != nil {
		variables["amount_msats"] = amountMsats
	}

	response, err := client.Requester.ExecuteGraphql(scripts.PAY_INVOICE_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["pay_invoice"].(map[string]interface{})
	var payment objects.OutgoingPayment
	paymentJson, err := json.Marshal(output["payment"].(map[string]interface{}))
	if err != nil {
		return nil, errors.New("error parsing payment")
	}
	json.Unmarshal(paymentJson, &payment)
	return &payment, nil
}

// PayUmaInvoice sends an UMA payment to a node on the Lightning Network, based on the invoice
// (as defined by the BOLT11 specification) that you provide.
// This should only be used for paying UMA invoices, with `pay_invoice` preferred in the general case.
//
// Args:
//
//	nodeId: The node from where you want to send the payment.
//	encodedInvoice: The invoice you want to pay (as defined by the BOLT11 standard).
//	timeoutSecs: The number of seconds that you are willing to wait for the payment to complete.
//	maximumFeesMsats: The maximum amount of fees that you are willing to pay for this payment, expressed in mSATs.
//	amountMsats: The amount you will pay for this invoice, expressed in msats.
//		It should ONLY be set when the invoice amount is zero.
func (client *LightsparkClient) PayUmaInvoice(nodeId string, encodedInvoice string,
	timeoutSecs int, maximumFeesMsats int64, amountMsats *int64) (*objects.OutgoingPayment, error) {

	variables := map[string]interface{}{
		"node_id":            nodeId,
		"encoded_invoice":    encodedInvoice,
		"timeout_secs":       timeoutSecs,
		"maximum_fees_msats": maximumFeesMsats,
	}
	if amountMsats != nil {
		variables["amount_msats"] = amountMsats
	}

	response, err := client.Requester.ExecuteGraphql(scripts.PAY_UMA_INVOICE_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["pay_uma_invoice"].(map[string]interface{})
	var payment objects.OutgoingPayment
	paymentJson, err := json.Marshal(output["payment"].(map[string]interface{}))
	if err != nil {
		return nil, errors.New("error parsing payment")
	}
	json.Unmarshal(paymentJson, &payment)
	return &payment, nil
}

// RequestWithdrawal withdraws funds from the account and sends it to the requested
// bitcoin address. Depending on the chosen mode, it will first take the funds from
// the wallet, and if applicable, close channels appropriately to recover enough
// funds.
// The process is asynchronous and may take up to a few minutes.
// You can check the progress by polling the `WithdrawalRequest`
// that is created, or by subscribing to a webhook.
//
// Args:
//
//	nodeId: The node from which you'd like to make the withdrawal.
//	amountSats: The amount you want to withdraw from this node in Satoshis.
//		Use the special value -1 to withdrawal all funds from this node.
//	bitcoinAddress: The bitcoin address where you want to receive the funds.
//	withdrawalMode: The mode that will be used to withdraw the funds.
func (client *LightsparkClient) RequestWithdrawal(nodeId string, amountSats int64,
	bitcoinAddress string, withdrawalMode objects.WithdrawalMode) (*objects.WithdrawalRequest, error) {

	variables := map[string]interface{}{
		"node_id":         nodeId,
		"amount_sats":     amountSats,
		"bitcoin_address": bitcoinAddress,
		"withdrawal_mode": withdrawalMode,
	}

	response, err := client.Requester.ExecuteGraphql(scripts.REQUEST_WITHDRAWAL_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["request_withdrawal"].(map[string]interface{})
	var withdrawalRequest objects.WithdrawalRequest
	withdrawalRequestJson, err := json.Marshal(output["request"].(map[string]interface{}))
	if err != nil {
		return nil, errors.New("error parsing withdrawal request")
	}
	json.Unmarshal(withdrawalRequestJson, &withdrawalRequest)
	return &withdrawalRequest, nil
}

// SendPayment sends a payment directly to a node on the Lightning Network
// through the public key of the node without an invoice.
//
// Args:
//
//	nodeId: The node from where you want to send the payment.
//	destinationPublicKey: The public key of the node that will receive the payment.
//	amountMsats: The amount you will pay for this invoice, expressed in msats.
//	timeoutSecs: The number of seconds that you are willing to wait for the payment to complete.
//	maximumFeesMsats: The maximum amount of fees that you are willing to pay for this payment, expressed in mSATs.
func (client *LightsparkClient) SendPayment(nodeId string, destinationPublicKey string,
	amountMsats int64, timeoutSecs int, maximumFeesMsats int64) (*objects.OutgoingPayment, error) {

	variables := map[string]interface{}{
		"node_id":                nodeId,
		"destination_public_key": destinationPublicKey,
		"amount_msats":           amountMsats,
		"timeout_secs":           timeoutSecs,
		"maximum_fees_msats":     maximumFeesMsats,
	}

	response, err := client.Requester.ExecuteGraphql(scripts.SEND_PAYMENT_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["send_payment"].(map[string]interface{})
	var payment objects.OutgoingPayment
	paymentJson, err := json.Marshal(output["payment"].(map[string]interface{}))
	if err != nil {
		return nil, errors.New("error parsing payment")
	}
	json.Unmarshal(paymentJson, &payment)
	return &payment, nil
}

// ScreenNode performs sanction screening on a lightning node against a given provider.
// It should only be called if you have a Chainalysis API Key in settings.
//
// Args:
//
//	provider: The provider that you want to use to perform the screening.
//	nodePubkey: The public key of the node that needs to be screened.
func (client *LightsparkClient) ScreenNode(
	provider objects.ComplianceProvider, nodePubkey string) (*objects.RiskRating, error) {

	variables := map[string]interface{}{"provider": provider, "node_pubkey": nodePubkey}
	response, err := client.Requester.ExecuteGraphql(scripts.SCREEN_NODE_MUTATION, variables)
	if err != nil {
		return nil, err
	}

	output := response["screen_node"].(map[string]interface{})
	ratingJson, err := json.Marshal(output["rating"].(string))
	if err != nil {
		return nil, errors.New("error parsing rating")
	}
	var rating objects.RiskRating
	json.Unmarshal(ratingJson, &rating)
	return &rating, nil
}

// RegisterPayment registers a succeeded payment with a compliance provider.
// It should only be called if you have a Chainalysis API Key in settings.
//
// Args:
//
//	provider: The provider that you want to use to register the payment.
//	paymentId: The unique id of the payment.
//	nodePubkey: The public key of the counterparty node which is the recipient
//	            node if the payment is an outgoing payment and the sender node
//	            if the payment is an incoming payment.
func (client *LightsparkClient) RegisterPayment(provider objects.ComplianceProvider,
	paymentId string, nodePubkey string, direction objects.PaymentDirection) error {

	variables := map[string]interface{}{
		"provider":    provider,
		"payment_id":  paymentId,
		"node_pubkey": nodePubkey,
		"direction":   direction,
	}
	_, err := client.Requester.ExecuteGraphql(scripts.REGISTER_PAYMENT_MUTATION, variables)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// GetEntity returns any `Entity`, identified by its unique ID.
//
// Args:
//
//	id: The unique ID of the entity.
func (client *LightsparkClient) GetEntity(id string) (*objects.Entity, error) {
	variables := map[string]interface{}{
		"id": id,
	}
	response, err := client.Requester.ExecuteGraphql(objects.GetEntityQuery, variables)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})
	entity, err := objects.EntityUnmarshal(output)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// ExecuteGraphqlRequest executes a GraphQL request.
//
// Args:
//
//	document: The GraphQL document that you want to execute.
//	variables: The variables that you want to pass to the GraphQL document.
func (client *LightsparkClient) ExecuteGraphqlRequest(document string,
	variables map[string]interface{}) (map[string]interface{}, error) {

	return client.Requester.ExecuteGraphql(document, variables)
}
