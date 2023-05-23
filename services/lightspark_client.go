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
	nodeKeys  map[string][]byte
}

func NewLightsparkClient(apiTokenClientId string, apiTokenClientSecret string,
	baseUrl *string) *LightsparkClient {

	requester := &requester.Requester{
		ApiTokenClientId:     apiTokenClientId,
		ApiTokenClientSecret: apiTokenClientSecret,
		BaseUrl:              baseUrl,
	}
	return &LightsparkClient{Requester: requester, nodeKeys: map[string][]byte{}}
}

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
	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_API_TOKEN_MUTATION, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["create_api_token"].(map[string]interface{})
	var apiToken objects.ApiToken
	apiTokenJson, err := json.Marshal(output["api_token"].(map[string]interface{}))
	json.Unmarshal(apiTokenJson, &apiToken)
	return &scripts.CreateApiTokenOutput{ApiToken: &apiToken, ClientSecret: output["client_secret"].(string)}, nil
}

func (client *LightsparkClient) CreateInvoice(nodeId string, amountMsats int64,
	memo *string, invoiceType *objects.InvoiceType) (*objects.Invoice, error) {

	variables := map[string]interface{}{
		"amount_msats": amountMsats,
		"node_id":      nodeId,
		"memo":         memo,
		"invoice_type": invoiceType,
	}
	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_INVOICE_MUTATION, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["create_invoice"].(map[string]interface{})
	var invoice objects.Invoice
	invoiceJson, err := json.Marshal(output["invoice"].(map[string]interface{}))
	json.Unmarshal(invoiceJson, &invoice)
	return &invoice, nil
}

func (client *LightsparkClient) CreateNodeWalletAddress(nodeId string) (string, error) {
	variables := map[string]interface{}{
		"node_id": nodeId,
	}
	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_NODE_WALLET_ADDRESS_MUTATION, variables, nil)
	if err != nil {
		return "", err
	}
	output := response["create_node_wallet_address"].(map[string]interface{})
	walletAddress := output["wallet_address"].(string)
	return walletAddress, nil
}

func (client *LightsparkClient) CreateTestModeInvoice(localNodeId string, amountMsats int64,
	memo *string, invoiceType *objects.InvoiceType) (*string, error) {

	variables := map[string]interface{}{
		"amount_msats": amountMsats,
		"local_node_id":      localNodeId,
		"memo":         memo,
		"invoice_type": invoiceType,
	}
	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_TEST_MODE_INVOICE_MUTATION, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["create_test_mode_invoice"].(map[string]interface{})
	encodedInvoice := output["encoded_payment_request"].(string)
	return &encodedInvoice, nil
}

func (client *LightsparkClient) CreateTestModePayment(localNodeId string, 
	encodedInvoice string, amountMsats *int64) (*objects.OutgoingPayment, error) {

	variables := map[string]interface{}{
		"local_node_id":	localNodeId,
		"encoded_invoice":	encodedInvoice,
	}
	if amountMsats != nil {
		variables["amount_msats"] = amountMsats
	}

	response, err := client.Requester.ExecuteGraphql(scripts.CREATE_TEST_MODE_PAYMENT_MUTATION, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["create_test_mode_payment"].(map[string]interface{})
	var payment objects.OutgoingPayment
	paymentJson, err := json.Marshal(output["payment"].(map[string]interface{}))
	json.Unmarshal(paymentJson, &payment)
	return &payment, nil
}

func (client *LightsparkClient) DecodePaymentRequest(encodedPaymentRequest string) (*objects.PaymentRequestData, error) {
	variables := map[string]interface{}{"encoded_payment_request": encodedPaymentRequest}
	response, err := client.Requester.ExecuteGraphql(scripts.DECODE_PAYMENT_REQUEST_QUERY, variables, nil)
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

func (client *LightsparkClient) DeleteApiToken(apiTokenId string) error {
	variables := map[string]interface{}{
		"api_token_id": apiTokenId,
	}
	_, err := client.Requester.ExecuteGraphql(scripts.DELETE_API_TOKEN_MUTATION, variables, nil)
	return err
}

func (client *LightsparkClient) FundNode(nodeId string, amountSats int64) (
	*objects.CurrencyAmount, error) {

	variables := map[string]interface{}{
		"node_id":     nodeId,
		"amount_sats": amountSats,
	}

	signingKey, err := client.getNodeSigningKey(nodeId)
	if err != nil {
		return nil, err
	}

	response, err := client.Requester.ExecuteGraphql(scripts.FUND_NODE_MUTATION,
		variables, signingKey)
	if err != nil {
		return nil, err
	}

	output := response["fund_node"].(map[string]interface{})
	var amount objects.CurrencyAmount
	amountJson, err := json.Marshal(output["amount"].(map[string]interface{}))
	json.Unmarshal(amountJson, &amount)
	return &amount, nil
}

func (client *LightsparkClient) GetBitcoinFeeEstimate(
	bitcoinNetwork objects.BitcoinNetwork) (*objects.FeeEstimate, error) {

	variables := map[string]interface{}{"bitcoin_network": bitcoinNetwork}
	response, err := client.Requester.ExecuteGraphql(scripts.BITCOIN_FEE_ESTIMATE_QUERY, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["bitcoin_fee_estimate"].(map[string]interface{})
	var feeEstimate objects.FeeEstimate
	feeEstimateJson, err := json.Marshal(output)
	json.Unmarshal(feeEstimateJson, &feeEstimate)
	return &feeEstimate, nil
}

func (client *LightsparkClient) GetCurrentAccount() (*objects.Account, error) {
	variables := map[string]interface{}{}
	response, err := client.Requester.ExecuteGraphql(scripts.CURRENT_ACCOUNT_QUERY, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["current_account"].(map[string]interface{})
	var account objects.Account
	accountJson, err := json.Marshal(output)
	json.Unmarshal(accountJson, &account)
	return &account, nil
}

func (client *LightsparkClient) GetLightningFeeEstimateForInvoice(nodeId string,
	encodedInvoice string, amountMsats *int64) (*objects.LightningFeeEstimateOutput, error) {

	variables := map[string]interface{}{
		"node_id":                 nodeId,
		"encoded_payment_request": encodedInvoice,
		"amount_msats":            amountMsats,
	}
	response, err := client.Requester.ExecuteGraphql(scripts.LIGHTNING_FEE_ESTIMATE_FOR_INVOICE_QUERY, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["lightning_fee_estimate_for_invoice"].(map[string]interface{})
	var feeEstimate objects.LightningFeeEstimateOutput
	feeEstimateJson, err := json.Marshal(output)
	json.Unmarshal(feeEstimateJson, &feeEstimate)
	return &feeEstimate, nil
}

func (client *LightsparkClient) GetLightningFeeEstimateForNode(nodeId string,
	destinationNodePublicKey string, amountMsats int64) (*objects.LightningFeeEstimateOutput, error) {

	variables := map[string]interface{}{
		"node_id":                     nodeId,
		"destination_node_public_key": destinationNodePublicKey,
		"amount_msats":                amountMsats,
	}
	response, err := client.Requester.ExecuteGraphql(scripts.LIGHTNING_FEE_ESTIMATE_FOR_NODE_QUERY, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["lightning_fee_estimate_for_node"].(map[string]interface{})
	var feeEstimate objects.LightningFeeEstimateOutput
	feeEstimateJson, err := json.Marshal(output)
	json.Unmarshal(feeEstimateJson, &feeEstimate)
	return &feeEstimate, nil
}

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

	signingKey, err := client.getNodeSigningKey(nodeId)
	if err != nil {
		return nil, err
	}

	response, err := client.Requester.ExecuteGraphql(scripts.PAY_INVOICE_MUTATION,
		variables, signingKey)
	if err != nil {
		return nil, err
	}

	output := response["pay_invoice"].(map[string]interface{})
	var payment objects.OutgoingPayment
	paymentJson, err := json.Marshal(output["payment"].(map[string]interface{}))
	json.Unmarshal(paymentJson, &payment)
	return &payment, nil
}

func (client *LightsparkClient) RecoverNodeSigningKey(nodeId string,
	nodePassword string) ([]byte, error) {

	variables := map[string]interface{}{
		"node_id": nodeId,
	}
	response, err := client.Requester.ExecuteGraphql(scripts.RECOVER_NODE_SIGNING_KEY_QUERY, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})
	encryptedKeyOutput := output["encrypted_signing_private_key"].(map[string]interface{})
	encryptedKey := encryptedKeyOutput["encrypted_value"].(string)
	cipher := encryptedKeyOutput["cipher"].(string)

	signingKey, err := utils.DecryptPrivateKey(cipher, encryptedKey, nodePassword)
	client.LoadNodeSigningKey(nodeId, signingKey)

	return signingKey, nil
}

func (client *LightsparkClient) RequestWithdrawal(nodeId string, amountSats int64,
	bitcoinAddress string, withdrawalMode objects.WithdrawalMode) (*objects.WithdrawalRequest, error) {

	variables := map[string]interface{}{
		"node_id":         nodeId,
		"amount_sats":     amountSats,
		"bitcoin_address": bitcoinAddress,
		"withdrawal_mode": withdrawalMode,
	}

	signingKey, err := client.getNodeSigningKey(nodeId)
	if err != nil {
		return nil, err
	}

	response, err := client.Requester.ExecuteGraphql(scripts.REQUEST_WITHDRAWAL_MUTATION,
		variables, signingKey)
	if err != nil {
		return nil, err
	}

	output := response["request_withdrawal"].(map[string]interface{})
	var withdrawalRequest objects.WithdrawalRequest
	withdrawalRequestJson, err := json.Marshal(output["request"].(map[string]interface{}))
	json.Unmarshal(withdrawalRequestJson, &withdrawalRequest)
	return &withdrawalRequest, nil
}

func (client *LightsparkClient) SendPayment(nodeId string, destinationPublicKey string,
	amountMsats int64, timeoutSecs int, maximumFeesMsats int64) (*objects.OutgoingPayment, error) {

	variables := map[string]interface{}{
		"node_id":                nodeId,
		"destination_public_key": destinationPublicKey,
		"amount_msats":           amountMsats,
		"timeout_secs":           timeoutSecs,
		"maximum_fees_msats":     maximumFeesMsats,
	}

	signingKey, err := client.getNodeSigningKey(nodeId)
	if err != nil {
		return nil, err
	}

	response, err := client.Requester.ExecuteGraphql(scripts.SEND_PAYMENT_MUTATION,
		variables, signingKey)
	if err != nil {
		return nil, err
	}

	output := response["send_payment"].(map[string]interface{})
	var payment objects.OutgoingPayment
	paymentJson, err := json.Marshal(output["payment"].(map[string]interface{}))
	json.Unmarshal(paymentJson, &payment)
	return &payment, nil
}

func (client *LightsparkClient) GetEntity(id string) (*objects.Entity, error) {
	variables := map[string]interface{}{
		"id": id,
	}
	response, err := client.Requester.ExecuteGraphql(objects.GetEntityQuery, variables, nil)
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

func (client *LightsparkClient) ExecuteGraphqlRequest(document string,
	variables map[string]interface{}) (map[string]interface{}, error) {

	return client.Requester.ExecuteGraphql(document, variables, nil)
}

func (client *LightsparkClient) LoadNodeSigningKey(nodeId string, signingKey []byte) {
	client.nodeKeys[nodeId] = signingKey
}

func (client *LightsparkClient) getNodeSigningKey(nodeId string) ([]byte, error) {
	nodeKey, ok := client.nodeKeys[nodeId]
	if !ok {
		return nil, errors.New("we did not find the signing key for node. Please call recover_node_signing_key")
	}
	return nodeKey, nil
}
