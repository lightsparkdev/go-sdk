// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package remotesigning

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/webhooks"
)

func ParseRemoteSigningRequest(webhook webhooks.WebhookEvent) (SigningRequest, error) {
	if webhook.EventType != objects.WebhookEventTypeRemoteSigning {
		return nil, errors.New("webhook event is not for remote signing")
	}

	var subtype objects.RemoteSigningSubEventType
	subEventTypeStr := (*webhook.Data)["sub_event_type"].(string)
	err := subtype.UnmarshalJSON([]byte(`"` + subEventTypeStr + `"`))
	if err != nil {
		return nil, errors.New("invalid remote signing sub_event_type")
	}

	switch subtype {
	case objects.RemoteSigningSubEventTypeEcdh:
		return ParseECDHRequest(webhook)
	case objects.RemoteSigningSubEventTypeGetPerCommitmentPoint:
		return ParseGetPerCommitmentPointRequest(webhook)
	case objects.RemoteSigningSubEventTypeReleasePerCommitmentSecret:
		return ParseReleasePerCommitmentSecretRequest(webhook)
	case objects.RemoteSigningSubEventTypeDeriveKeyAndSign:
		return ParseDeriveAndSignRequest(webhook)
	case objects.RemoteSigningSubEventTypeRequestInvoicePaymentHash:
		return ParseRequestInvoicePaymentHashRequest(webhook)
	case objects.RemoteSigningSubEventTypeSignInvoice:
		return ParseSignInvoiceRequest(webhook)
	case objects.RemoteSigningSubEventTypeReleasePaymentPreimage:
		return ParseReleasePaymentPreimageRequest(webhook)
	case objects.RemoteSigningSubEventTypeRevealCounterpartyPerCommitmentSecret:
		return ParseReleaseCounterpartyPerCommitmentSecretRequest(webhook)
	default:
		return nil, errors.New("invalid remote signing sub_event_type")
	}
}

func ParseECDHRequest(webhook webhooks.WebhookEvent) (*ECDHRequest, error) {
	log.Println("Parsing ECDH webhook")
	if webhook.Data == nil {
		return nil, errors.New("webhook data is missing")
	}

	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeEcdh.StringValue() {
		return nil, errors.New("sub_event_type is not ECDH")
	}

	publicKey := (*webhook.Data)["peer_public_key"]
	if publicKey == nil {
		return nil, errors.New("missing peer_public_key in webhook")
	}

	network, err := bitcoinNetworkFromWebhookData(*webhook.Data)
	if err != nil {
		return nil, err
	}

	request := ECDHRequest{
		NodeId:         webhook.EntityId,
		PeerPubKeyHex:  publicKey.(string),
		BitcoinNetwork: network,
	}

	return &request, nil
}

func ParseGetPerCommitmentPointRequest(webhook webhooks.WebhookEvent) (*GetPerCommitmentPointRequest, error) {
	log.Println("Parsing GET_PER_COMMITMENT_POINT webhook")
	if webhook.Data == nil {
		return nil, errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeGetPerCommitmentPoint.StringValue() {
		return nil, errors.New("sub_event_type is not GET_PER_COMMITMENT_POINT")
	}

	perCommitmentPointIdx := (*webhook.Data)["per_commitment_point_idx"]
	if perCommitmentPointIdx == nil {
		return nil, errors.New("missing per_commitment_point_idx in webhook")
	}

	perCommitmentPointIdxInt, err := perCommitmentPointIdx.(json.Number).Int64()
	if err != nil {
		return nil, fmt.Errorf("invalid per_commitment_point_idx in webhook (%v)", perCommitmentPointIdx)
	}

	derivationPath := (*webhook.Data)["derivation_path"]
	if derivationPath == nil {
		return nil, errors.New("missing derivation_path in webhook")
	}

	channelId := webhook.EntityId

	nodeID := (*webhook.Data)["node_id"]
	if nodeID == nil {
		return nil, errors.New("missing node_id in webhook")
	}

	network, err := bitcoinNetworkFromWebhookData(*webhook.Data)
	if err != nil {
		return nil, err
	}

	request := GetPerCommitmentPointRequest{
		ChannelId:             channelId,
		DerivationPath:        derivationPath.(string),
		PerCommitmentPointIdx: uint64(perCommitmentPointIdxInt),
		NodeId:                nodeID.(string),
		BitcoinNetwork:        network,
	}

	return &request, nil
}

func ParseReleasePerCommitmentSecretRequest(webhook webhooks.WebhookEvent) (*ReleasePerCommitmentSecretRequest, error) {
	log.Println("Parsing RELEASE_PER_COMMITMENT_SECRET webhook")
	if webhook.Data == nil {
		return nil, errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeReleasePerCommitmentSecret.StringValue() {
		return nil, errors.New("sub_event_type is not RELEASE_PER_COMMITMENT_POINT")
	}

	perCommitmentPointIdxNumber := (*webhook.Data)["per_commitment_point_idx"]
	if perCommitmentPointIdxNumber == nil {
		return nil, errors.New("missing per_commitment_point_idx in webhook")
	}

	perCommitmentPointIdx, err := perCommitmentPointIdxNumber.(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	derivationPath := (*webhook.Data)["derivation_path"]
	if derivationPath == nil {
		return nil, errors.New("missing derivation_path in webhook")
	}

	channelId := webhook.EntityId

	nodeID := (*webhook.Data)["node_id"]
	if nodeID == nil {
		return nil, errors.New("missing node_id in webhook")
	}

	network, err := bitcoinNetworkFromWebhookData(*webhook.Data)
	if err != nil {
		return nil, err
	}

	request := ReleasePerCommitmentSecretRequest{
		ChannelId:             channelId,
		DerivationPath:        derivationPath.(string),
		PerCommitmentPointIdx: uint64(perCommitmentPointIdx),
		NodeId:                nodeID.(string),
		BitcoinNetwork:        network,
	}

	return &request, nil
}

func ParseRequestInvoicePaymentHashRequest(webhook webhooks.WebhookEvent) (*InvoicePaymentHashRequest, error) {
	log.Println("Parsing REQUEST_INVOICE_PAYMENT_HASH webhook")
	if webhook.Data == nil {
		return nil, errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeRequestInvoicePaymentHash.StringValue() {
		return nil, errors.New("sub_event_type is not REQUEST_INVOICE_PAYMENT_HASH")
	}

	invoiceId := (*webhook.Data)["invoice_id"]
	if invoiceId == nil {
		return nil, errors.New("missing invoice_id in webhook")
	}
	log.Printf("invoiceId: %v", invoiceId)

	network, err := bitcoinNetworkFromWebhookData(*webhook.Data)
	if err != nil {
		return nil, err
	}

	request := InvoicePaymentHashRequest{
		InvoiceId:      invoiceId.(string),
		BitcoinNetwork: network,
	}
	return &request, nil
}

func ParseDeriveAndSignRequest(webhook webhooks.WebhookEvent) (*DeriveKeyAndSignRequest, error) {
	log.Println("Parsing DERIVE_KEY_AND_SIGN webhook")
	if webhook.Data == nil {
		return nil, errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeDeriveKeyAndSign.StringValue() {
		return nil, errors.New("sub_event_type is not DERIVE_KEY_AND_SIGN")
	}

	signingJobsJson := (*webhook.Data)["signing_jobs"]
	if signingJobsJson == nil {
		return nil, errors.New("missing signing_jobs in webhook")
	}
	signingJobsJsonString, err := json.Marshal(signingJobsJson.([]interface{}))
	if err != nil {
		return nil, err
	}

	var signingJobs []SigningJob
	err = json.Unmarshal(signingJobsJsonString, &signingJobs)
	if err != nil {
		return nil, err
	}

	network, err := bitcoinNetworkFromWebhookData(*webhook.Data)
	if err != nil {
		return nil, err
	}

	request := DeriveKeyAndSignRequest{
		SigningJobs:    signingJobs,
		BitcoinNetwork: network,
	}

	return &request, nil
}

func ParseSignInvoiceRequest(webhook webhooks.WebhookEvent) (*SignInvoiceRequest, error) {
	log.Println("Parsing SIGN_INVOICE webhook")
	if webhook.Data == nil {
		return nil, errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeSignInvoice.StringValue() {
		return nil, errors.New("sub_event_type is not SIGN_INVOICE")
	}

	invoiceId := (*webhook.Data)["invoice_id"]
	if invoiceId == nil {
		return nil, errors.New("missing invoice_id in webhook")
	}

	paymentRequestHash := (*webhook.Data)["payreq_hash"]
	if paymentRequestHash == nil {
		return nil, errors.New("missing payreq_hash in webhook")
	}

	network, err := bitcoinNetworkFromWebhookData(*webhook.Data)
	if err != nil {
		return nil, err
	}

	request := SignInvoiceRequest{
		InvoiceId:          invoiceId.(string),
		PaymentRequestHash: paymentRequestHash.(string),
		BitcoinNetwork:     network,
	}

	return &request, nil
}

func ParseReleasePaymentPreimageRequest(webhook webhooks.WebhookEvent) (*ReleasePaymentPreimageRequest, error) {
	log.Println("Parsing RELEASE_PAYMENT_PREIMAGE webhook")
	if webhook.Data == nil {
		return nil, errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeReleasePaymentPreimage.StringValue() {
		return nil, errors.New("sub_event_type is not RELEASE_PAYMENT_PREIMAGE")
	}

	invoiceId := (*webhook.Data)["invoice_id"]
	if invoiceId == nil {
		return nil, errors.New("missing invoice_id in webhook")
	}

	nonce := (*webhook.Data)["preimage_nonce"]

	network, err := bitcoinNetworkFromWebhookData(*webhook.Data)
	if err != nil {
		return nil, err
	}

	request := ReleasePaymentPreimageRequest{
		InvoiceId:      invoiceId.(string),
		BitcoinNetwork: network,
	}

	if nonce != nil {
		nonceStr := nonce.(string)
		request.Nonce = &nonceStr
	}

	return &request, nil
}

func ParseReleaseCounterpartyPerCommitmentSecretRequest(webhook webhooks.WebhookEvent) (*ReleaseCounterpartyPerCommitmentSecretRequest, error) {
	log.Println("Parsing RELEASE_COUNTERPARTY_PER_COMMITMENT_SECRET webhook")
	if webhook.Data == nil {
		return nil, errors.New("missing data in webhook")
	}
	if (*webhook.Data)["sub_event_type"] != objects.RemoteSigningSubEventTypeRevealCounterpartyPerCommitmentSecret.StringValue() {
		return nil, errors.New("sub_event_type is not RELEASE_COUNTERPARTY_PER_COMMITMENT_SECRET")
	}

	perCommitmentSecretIdx := (*webhook.Data)["per_commitment_secret_idx"]
	if perCommitmentSecretIdx == nil {
		return nil, errors.New("missing per_commitment_secret_idx in webhook")
	}

	perCommitmentSecretIdxInt, err := perCommitmentSecretIdx.(json.Number).Int64()
	if err != nil {
		return nil, fmt.Errorf("invalid per_commitment_point_idx in webhook (%v)", perCommitmentSecretIdx)
	}

	perCommitmentSecret := (*webhook.Data)["per_commitment_secret"]
	if perCommitmentSecret == nil {
		return nil, errors.New("missing per_commitment_secret in webhook")
	}

	nodeID := (*webhook.Data)["node_id"]
	if nodeID == nil {
		return nil, errors.New("missing node_id in webhook")
	}

	channelId := webhook.EntityId

	request := ReleaseCounterpartyPerCommitmentSecretRequest{
		ChannelId:              channelId,
		PerCommitmentSecretIdx: uint64(perCommitmentSecretIdxInt),
		PerCommitmentSecret:    perCommitmentSecret.(string),
		NodeId:                 nodeID.(string),
	}

	return &request, nil
}

func bitcoinNetworkFromWebhookData(data map[string]interface{}) (objects.BitcoinNetwork, error) {
	network := data["bitcoin_network"].(string)
	switch network {
	case "MAINNET":
		return objects.BitcoinNetworkMainnet, nil
	case "TESTNET":
		return objects.BitcoinNetworkTestnet, nil
	case "REGTEST":
		return objects.BitcoinNetworkRegtest, nil
	default:
		return objects.BitcoinNetworkUndefined, errors.New("invalid network")
	}
}

type SigningRequest interface {
	Type() objects.RemoteSigningSubEventType
}

// A signing request asking for a shared secret to be computed using Eliptic Curve Diffie-Hellman.
// The shared secret is computed using your node's identity private key and the peer's public key.
// This request is only for nodes created before 10/06/2023.
type ECDHRequest struct {
	NodeId         string
	PeerPubKeyHex  string
	BitcoinNetwork objects.BitcoinNetwork
}

func (r *ECDHRequest) Type() objects.RemoteSigningSubEventType {
	return objects.RemoteSigningSubEventTypeEcdh
}

// A signing request asking for a per-commitment point for a particular channel.
// The per-commitment point is the point on the secp256k1 curve for the commitment secret described
// in bolt 3.
//
// [Bolt 3]: https://github.com/lightning/bolts/blob/master/03-transactions.md#per-commitment-secret-requirements
type GetPerCommitmentPointRequest struct {
	ChannelId             string
	DerivationPath        string
	PerCommitmentPointIdx uint64
	NodeId                string
	BitcoinNetwork        objects.BitcoinNetwork
}

func (r *GetPerCommitmentPointRequest) Type() objects.RemoteSigningSubEventType {
	return objects.RemoteSigningSubEventTypeGetPerCommitmentPoint
}

// A signing request asking for a per-commitment secret to be released for a particular channel.
// The per-commitment secret is the secret described in bolt 3.
//
// [Bolt 3]: https://github.com/lightning/bolts/blob/master/03-transactions.md#per-commitment-secret-requirements
type ReleasePerCommitmentSecretRequest struct {
	ChannelId             string
	DerivationPath        string
	PerCommitmentPointIdx uint64
	NodeId                string
	BitcoinNetwork        objects.BitcoinNetwork
}

func (r *ReleasePerCommitmentSecretRequest) Type() objects.RemoteSigningSubEventType {
	return objects.RemoteSigningSubEventTypeReleasePerCommitmentSecret
}

// A signing request asking for a key to be derived and used to sign a message.
type DeriveKeyAndSignRequest struct {
	SigningJobs    []SigningJob
	BitcoinNetwork objects.BitcoinNetwork
}

func (r *DeriveKeyAndSignRequest) Type() objects.RemoteSigningSubEventType {
	return objects.RemoteSigningSubEventTypeDeriveKeyAndSign
}

// A signing request asking for a payment hash.
// A payment hash is the hash of a payment preimage. The payment preimage is the secret that is
// revealed when a payment is made on the Lightning Network.
type InvoicePaymentHashRequest struct {
	InvoiceId      string
	BitcoinNetwork objects.BitcoinNetwork
}

func (r *InvoicePaymentHashRequest) Type() objects.RemoteSigningSubEventType {
	return objects.RemoteSigningSubEventTypeRequestInvoicePaymentHash
}

// A signing request asking for an invoice to be signed.
// The invoice is signed using the node's identity private key.
// This request is only for nodes created before 10/06/2023.
type SignInvoiceRequest struct {
	InvoiceId          string
	PaymentRequestHash string
	BitcoinNetwork     objects.BitcoinNetwork
}

func (r *SignInvoiceRequest) Type() objects.RemoteSigningSubEventType {
	return objects.RemoteSigningSubEventTypeSignInvoice
}

// A signing request asking for a payment preimage to be released.
type ReleasePaymentPreimageRequest struct {
	InvoiceId      string
	Nonce          *string
	BitcoinNetwork objects.BitcoinNetwork
}

func (r *ReleasePaymentPreimageRequest) Type() objects.RemoteSigningSubEventType {
	return objects.RemoteSigningSubEventTypeReleasePaymentPreimage
}

// A informational request that reveals counterparty per-commitment secret.
type ReleaseCounterpartyPerCommitmentSecretRequest struct {
	ChannelId              string
	PerCommitmentSecretIdx uint64
	PerCommitmentSecret    string
	NodeId                string
}

func (r *ReleaseCounterpartyPerCommitmentSecretRequest) Type() objects.RemoteSigningSubEventType {
	return objects.RemoteSigningSubEventTypeRevealCounterpartyPerCommitmentSecret
}

// A signing job is a request to sign a message with a particular key.
// The signig key is computed using the node's master key and the parameter.
// DerivationPath is the bip32 derivation path to get the key from the master key `k`.
// Then apply MulTweak * k + AddTweak to get the final signing key.
type SigningJob struct {
	Id             string  `json:"id"`
	DerivationPath string  `json:"derivation_path"`
	Message        string  `json:"message"`
	AddTweak       *string `json:"add_tweak"`
	MulTweak       *string `json:"mul_tweak"`
	Script         *string `json:"script"`
	Transaction    *string `json:"transaction"`
	Amount 	       *int64 `json:"amount"`
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
