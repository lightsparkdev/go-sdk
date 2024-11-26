// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package remotesigning

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/lightsparkdev/go-sdk/crypto"
	utils "github.com/lightsparkdev/go-sdk/keyscripts"

	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/scripts"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/lightsparkdev/go-sdk/webhooks"
)

// HandleRemoteSigningWebhook handles a webhook event that is related to remote signing.
//
// This method should only be called with a webhook event that has the event_type `WebhookEventTypeRemoteSigning`.
// The method will call the appropriate handler for the sub_event_type of the webhook.
//
// Args:
//
//		client: The LightsparkClient used to respond to webhook events.
//	    validator: A validator for deciding whether to sign events.
//		webhook: The webhook event that you want to handle.
//		seedBytes: The bytes of the master seed that you want to use to sign messages or derive keys.
func HandleRemoteSigningWebhook(
	client *services.LightsparkClient,
	validator Validator,
	webhook webhooks.WebhookEvent,
	seedBytes []byte,
) (string, error) {
	response, err := GraphQLResponseForRemoteSigningWebhook(validator, webhook, seedBytes)

	if err != nil {
		if err.Error() == "declined to sign messages" {
			DeclineToSignMessages(client, webhook)
		}
		return "", err
	}

	if response == nil {
		// No response is required for this event type.
		return "", nil
	}

	return HandleSigningResponse(client, response)
}

func GraphQLResponseForRemoteSigningWebhook(
	validator Validator,
	webhook webhooks.WebhookEvent,
	seedBytes []byte,
) (SigningResponse, error) {
	// Calculate the xpub for each L1 signing job
	signingJobs, hasSigningJobs := (*webhook.Data)["signing_jobs"].([]interface{})
	if !hasSigningJobs {
		return nil, fmt.Errorf("signing_jobs not found or invalid type")
	}

	var xpubs []string
	for _, job := range signingJobs {
		jobMap, isValidJobMap := job.(map[string]interface{})
		if !isValidJobMap {
			return nil, fmt.Errorf("invalid signing job format")
		}
		derivationPath := jobMap["derivation_path"].(string)

		hardenedPath, _, err := SplitDerivationPath(derivationPath)
		if err != nil {
			return nil, err
		}

		masterSeedHex := hex.EncodeToString(seedBytes)
		xpub, err := utils.GenHardenedXPub(masterSeedHex, hardenedPath, "mainnet")
		if err != nil {
			return nil, err
		}
		xpubs = append(xpubs, xpub)
	}
	if !validator.ShouldSign(webhook, xpubs) {
		return nil, errors.New("declined to sign messages")
	}
	if webhook.EventType != objects.WebhookEventTypeRemoteSigning {
		return nil, errors.New("webhook event is not for remote signing")
	}
	if webhook.Data == nil {
		return nil, errors.New("webhook data is missing")
	}
	var subtype objects.RemoteSigningSubEventType
	subEventTypeStr := (*webhook.Data)["sub_event_type"].(string)
	log.Printf("Received remote signing webhook with sub_event_type %s", subEventTypeStr)
	err := subtype.UnmarshalJSON([]byte(`"` + subEventTypeStr + `"`))
	if err != nil {
		return nil, errors.New("invalid remote signing sub_event_type")
	}

	request, err := ParseRemoteSigningRequest(webhook)
	if err != nil {
		return nil, err
	}

	response, err := HandleSigningRequest(request, seedBytes)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func HandleSigningRequest(request SigningRequest, seedBytes []byte) (SigningResponse, error) {
	var response SigningResponse
	var err error
	switch request.Type() {
	case objects.RemoteSigningSubEventTypeEcdh:
		response, err = HandleEcdhRequest(request.(*ECDHRequest), seedBytes)
	case objects.RemoteSigningSubEventTypeGetPerCommitmentPoint:
		response, err = HandleGetPerCommitmentPointRequest(request.(*GetPerCommitmentPointRequest), seedBytes)
	case objects.RemoteSigningSubEventTypeReleasePerCommitmentSecret:
		response, err = HandleReleasePerCommitmentSecretRequest(request.(*ReleasePerCommitmentSecretRequest), seedBytes)
	case objects.RemoteSigningSubEventTypeDeriveKeyAndSign:
		response, err = HandleDeriveKeyAndSignRequest(request.(*DeriveKeyAndSignRequest), seedBytes)
	case objects.RemoteSigningSubEventTypeRequestInvoicePaymentHash:
		response, err = HandleInvoicePaymentHashRequest(request.(*InvoicePaymentHashRequest), seedBytes)
	case objects.RemoteSigningSubEventTypeSignInvoice:
		response, err = HandleSignInvoiceRequest(request.(*SignInvoiceRequest), seedBytes)
	case objects.RemoteSigningSubEventTypeReleasePaymentPreimage:
		response, err = HandleReleaseInvoicePreimageRequest(request.(*ReleasePaymentPreimageRequest), seedBytes)
	case objects.RemoteSigningSubEventTypeRevealCounterpartyPerCommitmentSecret:
		// No op for this event type.
		return nil, nil
	default:
		return nil, errors.New("webhook event is not for remote signing")
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

func HandleSigningResponse(client *services.LightsparkClient, response SigningResponse) (string, error) {
	graphql := response.GraphqlResponse()

	result, err := client.Requester.ExecuteGraphql(graphql.Query, graphql.Variables, nil)
	if err != nil {
		return "", err
	}

	output := result[graphql.OutputField].(map[string]interface{})
	var responseObj objects.UpdateNodeSharedSecretOutput
	outputJson, err := json.Marshal(output)
	if err != nil {
		return "", err
	}

	// This is just to validate the response.
	err = json.Unmarshal(outputJson, &responseObj)
	if err != nil {
		return "", err
	}

	return string(outputJson), nil
}

func DeclineToSignMessages(client *services.LightsparkClient, event webhooks.WebhookEvent) (string, error) {
	signingJobsJson := (*event.Data)["signing_jobs"]
	if signingJobsJson == nil {
		return "", errors.New("missing signing_jobs in webhook")
	}
	signingJobsJsonString, err := json.Marshal(signingJobsJson.([]interface{}))
	if err != nil {
		return "", err
	}

	var signingJobs []SigningJob
	err = json.Unmarshal(signingJobsJsonString, &signingJobs)
	if err != nil {
		return "", err
	}

	var payloadIds []string
	for _, signingJob := range signingJobs {
		payloadIds = append(payloadIds, signingJob.Id)
	}

	variables := map[string]interface{}{
		"payload_ids": payloadIds,
	}

	response, err := client.Requester.ExecuteGraphql(scripts.DECLINE_TO_SIGN_MESSAGES_MUTATION, variables, nil)
	if err != nil {
		return "", err
	}

	output := response["decline_to_sign_messages"].(map[string]interface{})
	var responseObj objects.DeclineToSignMessagesOutput
	outputJson, err := json.Marshal(output)
	if err != nil {
		return "", err
	}

	// This is just to validate the response.
	err = json.Unmarshal(outputJson, &responseObj)
	if err != nil {
		return "", err
	}

	return "rejected signing", nil
}

func HandleEcdhRequest(request *ECDHRequest, seedBytes []byte) (*ECDHResponse, error) {
	log.Println("Handling ECDH webhook")
	bitcoinNetwork, err := bitcoinNetworkConversion(request.BitcoinNetwork)
	if err != nil {
		return nil, err
	}

	sharedSecret, err := crypto.ECDH(seedBytes, bitcoinNetwork, request.PeerPubKeyHex)
	if err != nil {
		return nil, err
	}

	response := ECDHResponse{
		NodeId:          request.NodeId,
		SharedSecretHex: sharedSecret,
	}

	return &response, nil
}

func HandleGetPerCommitmentPointRequest(request *GetPerCommitmentPointRequest, seedBytes []byte) (*GetPerCommitmentPointResponse, error) {
	log.Println("Handling GET_PER_COMMITMENT_POINT webhook")
	bitcoinNetwork, err := bitcoinNetworkConversion(request.BitcoinNetwork)
	if err != nil {
		return nil, err
	}

	perCommitmentPoint, err := lightspark_crypto.GetPerCommitmentPoint(
		seedBytes,
		bitcoinNetwork,
		request.DerivationPath,
		request.PerCommitmentPointIdx)
	if err != nil {
		return nil, err
	}

	response := GetPerCommitmentPointResponse{
		ChannelId:             request.ChannelId,
		PerCommitmentPointIdx: request.PerCommitmentPointIdx,
		PerCommitmentPointHex: hex.EncodeToString(perCommitmentPoint),
	}

	return &response, nil
}

func HandleReleasePerCommitmentSecretRequest(request *ReleasePerCommitmentSecretRequest, seedBytes []byte) (*ReleasePerCommitmentSecretResponse, error) {
	log.Println("Handling RELEASE_PER_COMMITMENT_SECRET webhook")
	bitcoinNetwork, err := bitcoinNetworkConversion(request.BitcoinNetwork)
	if err != nil {
		return nil, err
	}

	perCommitmentSecret, err := lightspark_crypto.ReleasePerCommitmentSecret(
		seedBytes,
		bitcoinNetwork,
		request.DerivationPath,
		request.PerCommitmentPointIdx)
	if err != nil {
		return nil, err
	}

	response := ReleasePerCommitmentSecretResponse{
		ChannelId:             request.ChannelId,
		PerCommitmentPointIdx: request.PerCommitmentPointIdx,
		PerCommitmentSecret:   hex.EncodeToString(perCommitmentSecret),
	}

	return &response, nil
}

func HandleInvoicePaymentHashRequest(request *InvoicePaymentHashRequest, seedBytes []byte) (*InvoicePaymentHashResponse, error) {
	log.Println("Handling REQUEST_INVOICE_PAYMENT_HASH webhook")
	nonce, err := lightspark_crypto.GeneratePreimageNonce(seedBytes)
	if err != nil {
		return nil, err
	}
	paymentHash, err := lightspark_crypto.GeneratePreimageHash(seedBytes, nonce)
	if err != nil {
		return nil, err
	}

	nonce_str := hex.EncodeToString(nonce)

	response := InvoicePaymentHashResponse{
		InvoiceId:      request.InvoiceId,
		PaymentHashHex: hex.EncodeToString(paymentHash),
		Nonce:          &nonce_str,
	}

	return &response, nil
}

func HandleSignInvoiceRequest(request *SignInvoiceRequest, seedBytes []byte) (*SignInvoiceResponse, error) {
	log.Println("Handling SIGN_INVOICE webhook")
	bitcoinNetwork, err := bitcoinNetworkConversion(request.BitcoinNetwork)
	if err != nil {
		return nil, err
	}

	hash, err := hex.DecodeString(request.PaymentRequestHash)
	if err != nil {
		return nil, err
	}

	signedInvoice, err := lightspark_crypto.SignInvoiceHash(seedBytes, bitcoinNetwork, hash)
	if err != nil {
		log.Fatalf("Error signing invoice: %v", err)
		return nil, err
	}

	response := SignInvoiceResponse{
		InvoiceId:  request.InvoiceId,
		Signature:  hex.EncodeToString(signedInvoice.Signature),
		RecoveryId: signedInvoice.RecoveryId,
	}

	return &response, nil
}

func HandleReleaseInvoicePreimageRequest(request *ReleasePaymentPreimageRequest, seedBytes []byte) (*ReleasePaymentPreimageResponse, error) {
	log.Println("Handling RELEASE_PAYMENT_PREIMAGE webhook")
	nonce := request.Nonce
	if nonce == nil {
		return nil, errors.New("missing preimage_nonce in webhook")
	}
	nonceBytes, err := hex.DecodeString(*nonce)
	if err != nil {
		return nil, err
	}

	preimage, err := lightspark_crypto.GeneratePreimage(seedBytes, nonceBytes)
	if err != nil {
		return nil, err
	}

	response := ReleasePaymentPreimageResponse{
		InvoiceId:       request.InvoiceId,
		PaymentPreimage: hex.EncodeToString(preimage),
	}

	return &response, nil
}

func HandleDeriveKeyAndSignRequest(request *DeriveKeyAndSignRequest, seedBytes []byte) (*DeriveKeyAndSignResponse, error) {
	log.Println("Handling DERIVE_KEY_AND_SIGN webhook")
	bitcoinNetwork, err := bitcoinNetworkConversion(request.BitcoinNetwork)
	if err != nil {
		return nil, err
	}

	var signatures []SignatureResponse
	for _, signingJob := range request.SigningJobs {
		signature, err := signSigningJob(signingJob, seedBytes, bitcoinNetwork)
		if err != nil {
			return nil, err
		}
		signatures = append(signatures, SignatureResponse{
			Id:        signingJob.Id,
			Signature: signature.Signature,
		})
	}

	response := DeriveKeyAndSignResponse{
		Signatures: signatures,
	}

	return &response, nil
}

func bitcoinNetworkConversion(network objects.BitcoinNetwork) (lightspark_crypto.BitcoinNetwork, error) {
	switch network {
	case objects.BitcoinNetworkMainnet:
		return lightspark_crypto.Mainnet, nil
	case objects.BitcoinNetworkTestnet:
		return lightspark_crypto.Testnet, nil
	case objects.BitcoinNetworkRegtest:
		return lightspark_crypto.Regtest, nil
	default:
		return lightspark_crypto.BitcoinNetwork(0), errors.New("invalid network")
	}
}

func signSigningJob(signingJob SigningJob, seedBytes []byte, network lightspark_crypto.BitcoinNetwork) (objects.IdAndSignature, error) {
	addTweakBytes, err := signingJob.AddTweakBytes()
	if err != nil {
		return objects.IdAndSignature{}, err
	}
	mulTweakBytes, err := signingJob.MulTweakBytes()
	if err != nil {
		return objects.IdAndSignature{}, err
	}
	messageBytes, err := signingJob.MessageBytes()
	if err != nil {
		return objects.IdAndSignature{}, err
	}

	signatureBytes, err := lightspark_crypto.DeriveKeyAndSign(
		seedBytes,
		network,
		messageBytes,
		signingJob.DerivationPath,
		true,
		&addTweakBytes,
		&mulTweakBytes)
	if err != nil {
		return objects.IdAndSignature{}, err
	}
	signature := objects.IdAndSignature{
		Id:        signingJob.Id,
		Signature: hex.EncodeToString(signatureBytes),
	}
	return signature, nil
}
