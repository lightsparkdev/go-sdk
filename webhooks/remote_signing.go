package webhooks

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/lightsparkdev/go-sdk/utils"
	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/scripts"
	"github.com/lightsparkdev/go-sdk/services"
)

// HandleRemoteSigningWebhook handles a webhook event that is related to remote signing.
//
// This method should only be called with a webhook event that has the event_type `WebhookEventTypeRemoteSigning`.
// The method will call the appropriate handler for the sub_event_type of the webhook.
//
// Args:
//
//	webhook: The webhook event that you want to handle.
//	seedBytes: The bytes of the master seed that you want to use to sign messages or derive keys.
func HandleRemoteSigningWebhook(client *services.LightsparkClient, webhook WebhookEvent, seedBytes []byte) (string, error) {
	if webhook.EventType != objects.WebhookEventTypeRemoteSigning {
		return "", errors.New("webhook event is not for remote signing")
	}
	if webhook.Data == nil {
		return "", errors.New("webhook data is missing")
	}
	var subtype objects.RemoteSigningSubEventType
	subEventTypeStr := (*webhook.Data)["sub_event_type"].(string)
	log.Printf("Received remote signing webhook with sub_event_type %s", subEventTypeStr)
	err := subtype.UnmarshalJSON([]byte(`"` + subEventTypeStr + `"`))
	if err != nil {
		return "", errors.New("webhook data is missing")
	}

	switch subtype {
	case objects.RemoteSigningSubEventTypeEcdh:
		return HandleEcdhWebhook(client, webhook, seedBytes)
	case objects.RemoteSigningSubEventTypeGetPerCommitmentPoint:
		return HandleGetPerCommitmentPointWebhook(client, webhook, seedBytes)
	case objects.RemoteSigningSubEventTypeReleasePerCommitmentSecret:
		return HandleReleasePerCommitmentSecretWebhook(client, webhook, seedBytes)
	case objects.RemoteSigningSubEventTypeDeriveKeyAndSign:
		return HandleDeriveKeyAndSignWebhook(client, webhook, seedBytes)
	case objects.RemoteSigningSubEventTypeSignInvoice:
		return HandleSignInvoiceWebhook(client, webhook, seedBytes)
	case objects.RemoteSigningSubEventTypeReleasePaymentPreimage:
		return HandleReleaseInvoicePreimageWebhook(client, webhook, seedBytes)
	default:
		return "", errors.New("webhook event is not for remote signing")
	}
}

func HandleEcdhWebhook(client *services.LightsparkClient, webhook WebhookEvent, seedBytes []byte) (string, error) {
	log.Println("Handling ECDH webhook")
	if webhook.Data == nil {
		return "", errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeEcdh.StringValue() {
		return "", errors.New("sub_event_type is not ECDH")
	}

	publicKey := (*webhook.Data)["peer_public_key"]
	if publicKey == nil {
		return "", errors.New("missing peer_public_key in webhook")
	}

	bitcoinNetwork, err := bitcoinNetworkFromData(*webhook.Data)
	if err != nil {
		return "", err
	}

	sharedSecret, err := utils.ECDH(seedBytes, bitcoinNetwork, publicKey.(string))
	if err != nil {
		return "", err
	}

	variables := map[string]interface{}{
		"node_id":       webhook.EntityId,
		"shared_secret": sharedSecret,
	}

	response, err := client.Requester.ExecuteGraphql(scripts.UPDATE_NODE_SHARED_SECRET_MUTATION, variables)
	if err != nil {
		return "", err
	}

	output := response["update_node_shared_secret"].(map[string]interface{})
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

func HandleGetPerCommitmentPointWebhook(client *services.LightsparkClient, webhook WebhookEvent, seedBytes []byte) (string, error) {
	log.Println("Handling GET_PER_COMMITMENT_POINT webhook")
	if webhook.Data == nil {
		return "", errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeGetPerCommitmentPoint.StringValue() {
		return "", errors.New("sub_event_type is not GET_PER_COMMITMENT_POINT")
	}

	perCommitmentPointIdx := (*webhook.Data)["per_commitment_point_idx"]
	if perCommitmentPointIdx == nil {
		return "", errors.New("missing per_commitment_point_idx in webhook")
	}

	perCommitmentPointIdxInt, err := perCommitmentPointIdx.(json.Number).Int64()
	if err != nil {
		return "", fmt.Errorf("invalid per_commitment_point_idx in webhook (%v)", perCommitmentPointIdx)
	}

	derivationPath := (*webhook.Data)["derivation_path"]
	if derivationPath == nil {
		return "", errors.New("missing derivation_path in webhook")
	}

	bitcoinNetwork, err := bitcoinNetworkFromData(*webhook.Data)
	if err != nil {
		return "", err
	}

	channelId := webhook.EntityId
	perCommitmentPoint, err := lightspark_crypto.GetPerCommitmentPoint(
		seedBytes,
		bitcoinNetwork,
		derivationPath.(string),
		uint64(perCommitmentPointIdxInt))
	if err != nil {
		return "", err
	}
	variables := map[string]interface{}{
		"channel_id":                 channelId,
		"per_commitment_point":       hex.EncodeToString(perCommitmentPoint),
		"per_commitment_point_index": uint64(perCommitmentPointIdxInt),
	}

	response, err := client.Requester.ExecuteGraphql(scripts.UPDATE_CHANNEL_PER_COMMITMENT_POINT_MUTATION, variables)
	if err != nil {
		return "", err
	}

	output := response["update_channel_per_commitment_point"].(map[string]interface{})
	var responseObj objects.UpdateChannelPerCommitmentPointOutput
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

func HandleReleasePerCommitmentSecretWebhook(client *services.LightsparkClient, webhook WebhookEvent, seedBytes []byte) (string, error) {
	log.Println("Handling RELEASE_PER_COMMITMENT_SECRET webhook")
	if webhook.Data == nil {
		return "", errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeReleasePerCommitmentSecret.StringValue() {
		return "", errors.New("sub_event_type is not RELEASE_PER_COMMITMENT_POINT")
	}

	perCommitmentPointIdxNumber := (*webhook.Data)["per_commitment_point_idx"]
	if perCommitmentPointIdxNumber == nil {
		return "", errors.New("missing per_commitment_point_idx in webhook")
	}

	perCommitmentPointIdx, err := perCommitmentPointIdxNumber.(json.Number).Int64()
	if err != nil {
		return "", err
	}

	derivationPath := (*webhook.Data)["derivation_path"]
	if derivationPath == nil {
		return "", errors.New("missing derivation_path in webhook")
	}

	bitcoinNetwork, err := bitcoinNetworkFromData(*webhook.Data)
	if err != nil {
		return "", err
	}

	channelId := webhook.EntityId
	perCommitmentSecret, err := lightspark_crypto.ReleasePerCommitmentSecret(
		seedBytes,
		bitcoinNetwork,
		derivationPath.(string),
		uint64(perCommitmentPointIdx))
	if err != nil {
		return "", err
	}
	variables := map[string]interface{}{
		"channel_id":            channelId,
		"per_commitment_secret": hex.EncodeToString(perCommitmentSecret),
	}

	response, err := client.Requester.ExecuteGraphql(scripts.RELEASE_CHANNEL_PER_COMMITMENT_SECRET_MUTATION, variables)
	if err != nil {
		return "", err
	}

	output := response["release_channel_per_commitment_secret"].(map[string]interface{})
	var responseObj objects.ReleaseChannelPerCommitmentSecretOutput
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

func HandleSignInvoiceWebhook(client *services.LightsparkClient, webhook WebhookEvent, seedBytes []byte) (string, error) {
	log.Println("Handling SIGN_INVOICE webhook")
	if webhook.Data == nil {
		return "", errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeSignInvoice.StringValue() {
		return "", errors.New("sub_event_type is not SIGN_INVOICE")
	}

	invoiceId := (*webhook.Data)["invoice_id"]
	if invoiceId == nil {
		return "", errors.New("missing invoice_id in webhook")
	}
	log.Printf("invoiceId: %v", invoiceId)

	payReqHash := (*webhook.Data)["payreq_hash"]
	if payReqHash == nil {
		return "", errors.New("missing payreq_hash in webhook")
	}
	log.Printf("payReqHash: %v", payReqHash)
	payReqHashBytes, err := hex.DecodeString(payReqHash.(string))
	if err != nil {
		log.Fatalf("Error decoding payreq_hash: %v", err)
		return "", err
	}

	log.Printf("payReqHashBytes: %v", payReqHashBytes)
	bitcoinNetwork, err := bitcoinNetworkFromData(*webhook.Data)
	log.Printf("bitcoinNetwork: %v", bitcoinNetwork)
	if err != nil {
		log.Fatalf("Error getting bitcoin network: %v", err)
		return "", err
	}

	signedInvoice, err := lightspark_crypto.SignInvoiceHash(seedBytes, bitcoinNetwork, payReqHashBytes)
	log.Printf("signedInvoice: %v", signedInvoice)
	if err != nil {
		log.Fatalf("Error signing invoice: %v", err)
		return "", err
	}
	variables := map[string]interface{}{
		"invoice_id":  invoiceId,
		"signature":   hex.EncodeToString(signedInvoice.Signature),
		"recovery_id": signedInvoice.RecoveryId,
	}

	response, err := client.Requester.ExecuteGraphql(scripts.SIGN_INVOICE_MUTATION, variables)
	if err != nil {
		return "", err
	}

	output := response["sign_invoice"].(map[string]interface{})
	var responseObj objects.SignInvoiceOutput
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

func HandleReleaseInvoicePreimageWebhook(client *services.LightsparkClient, webhook WebhookEvent, seedBytes []byte) (string, error) {
	log.Println("Handling RELEASE_PAYMENT_PREIMAGE webhook")
	if webhook.Data == nil {
		return "", errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeReleasePaymentPreimage.StringValue() {
		return "", errors.New("sub_event_type is not RELEASE_PAYMENT_PREIMAGE")
	}

	nonce := (*webhook.Data)["preimage_nonce"]
	if nonce == nil {
		return "", errors.New("missing preimage_nonce in webhook")
	}
	nonceBytes, err := hex.DecodeString(nonce.(string))
	if err != nil {
		return "", err
	}

	invoiceId := (*webhook.Data)["invoice_id"]
	if invoiceId == nil {
		return "", errors.New("missing invoice_id in webhook")
	}

	preimage, err := lightspark_crypto.GeneratePreimage(seedBytes, nonceBytes)
	if err != nil {
		return "", err
	}
	variables := map[string]interface{}{
		"invoice_id":       invoiceId.(string),
		"payment_preimage": hex.EncodeToString(preimage),
	}

	response, err := client.Requester.ExecuteGraphql(scripts.RELEASE_PAYMENT_PREIMAGE_MUTATION, variables)
	if err != nil {
		return "", err
	}

	output := response["release_payment_preimage"].(map[string]interface{})
	var responseObj objects.ReleasePaymentPreimageOutput
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

type SigningJob struct {
	Id             string  `json:"id"`
	DerivationPath string  `json:"derivation_path"`
	Message        string  `json:"message"`
	AddTweak       *string `json:"add_tweak"`
	MulTweak       *string `json:"mul_tweak"`
	IsRaw          bool    `json:"is_raw"`
}

func (j *SigningJob) MulTweakBytes() ([]byte, error) {
	if j.MulTweak == nil {
		return nil, nil
	}
	return hex.DecodeString(*j.MulTweak)
}

func (j *SigningJob) AddTweakBytes() ([]byte, error) {
	if j.AddTweak == nil {
		return nil, nil
	}
	return hex.DecodeString(*j.AddTweak)
}

func (j *SigningJob) MessageBytes() ([]byte, error) {
	return hex.DecodeString(j.Message)
}

// signatureResponse A separate type is required for the response because the json field names are different from objects.Signature.
type signatureResponse struct {
	Id        string `json:"id"`
	Signature string `json:"signature"`
}

func HandleDeriveKeyAndSignWebhook(client *services.LightsparkClient, webhook WebhookEvent, seedBytes []byte) (string, error) {
	log.Println("Handling DERIVE_KEY_AND_SIGN webhook")
	if webhook.Data == nil {
		return "", errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeDeriveKeyAndSign.StringValue() {
		return "", errors.New("sub_event_type is not DERIVE_KEY_AND_SIGN")
	}

	signingJobsJson := (*webhook.Data)["signing_jobs"]
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

	bitcoinNetwork, err := bitcoinNetworkFromData(*webhook.Data)
	if err != nil {
		return "", err
	}

	var signatures []signatureResponse
	for _, signingJob := range signingJobs {
		signature, err := signSigningJob(signingJob, seedBytes, bitcoinNetwork)
		if err != nil {
			return "", err
		}
		signatures = append(signatures, signatureResponse{
			Id:        signingJob.Id,
			Signature: signature.Signature,
		})
	}
	variables := map[string]interface{}{
		"signatures": signatures,
	}

	response, err := client.Requester.ExecuteGraphql(scripts.SIGN_MESSAGES_MUTATION, variables)
	if err != nil {
		return "", err
	}

	output := response["sign_messages"].(map[string]interface{})
	var responseObj objects.SignMessagesOutput
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

func bitcoinNetworkFromData(data map[string]interface{}) (lightspark_crypto.BitcoinNetwork, error) {
	network := data["bitcoin_network"].(string)
	switch network {
	case "MAINNET":
		return lightspark_crypto.Mainnet, nil
	case "TESTNET":
		return lightspark_crypto.Testnet, nil
	case "REGTEST":
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
		signingJob.IsRaw,
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
