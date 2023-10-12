// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package remotesigning

import (
	"github.com/lightsparkdev/go-sdk/scripts"
)

type SigningResponse interface {
	GraphqlResponse() *GraphQLResponse
}

type ECDHResponse struct {
	NodeId          string
	SharedSecretHex string
}

func (r ECDHResponse) GraphqlResponse() *GraphQLResponse {
	return &GraphQLResponse{
		Query: scripts.UPDATE_NODE_SHARED_SECRET_MUTATION,
		Variables: map[string]interface{}{
			"node_id":       r.NodeId,
			"shared_secret": r.SharedSecretHex,
		},
		OutputField: "update_node_shared_secret",
	}
}

type GetPerCommitmentPointResponse struct {
	ChannelId             string
	PerCommitmentPointIdx uint64
	PerCommitmentPointHex string
}

func (r GetPerCommitmentPointResponse) GraphqlResponse() *GraphQLResponse {
	return &GraphQLResponse{
		Query: scripts.UPDATE_CHANNEL_PER_COMMITMENT_POINT_MUTATION,
		Variables: map[string]interface{}{
			"channel_id":                 r.ChannelId,
			"per_commitment_point":       r.PerCommitmentPointHex,
			"per_commitment_point_index": r.PerCommitmentPointIdx,
		},
		OutputField: "update_channel_per_commitment_point",
	}
}

type ReleasePerCommitmentSecretResponse struct {
	ChannelId             string
	PerCommitmentPointIdx uint64
	PerCommitmentSecret   string
}

func (r ReleasePerCommitmentSecretResponse) GraphqlResponse() *GraphQLResponse {
	return &GraphQLResponse{
		Query: scripts.RELEASE_CHANNEL_PER_COMMITMENT_SECRET_MUTATION,
		Variables: map[string]interface{}{
			"channel_id":            r.ChannelId,
			"per_commitment_secret": r.PerCommitmentSecret,
			"per_commitment_index":  r.PerCommitmentPointIdx,
		},
		OutputField: "release_channel_per_commitment_secret",
	}
}

// SignatureResponse A separate type is required for the response because the json field names are different from objects.Signature.
type SignatureResponse struct {
	Id        string `json:"id"`
	Signature string `json:"signature"`
}

type DeriveKeyAndSignResponse struct {
	Signatures []SignatureResponse
}

func (r DeriveKeyAndSignResponse) GraphqlResponse() *GraphQLResponse {
	return &GraphQLResponse{
		Query: scripts.SIGN_MESSAGES_MUTATION,
		Variables: map[string]interface{}{
			"signatures": r.Signatures,
		},
		OutputField: "sign_messages",
	}
}

type InvoicePaymentHashResponse struct {
	InvoiceId      string
	PaymentHashHex string
	Nonce          *string
}

func (r InvoicePaymentHashResponse) GraphqlResponse() *GraphQLResponse {
	variables := map[string]interface{}{
		"invoice_id":   r.InvoiceId,
		"payment_hash": r.PaymentHashHex,
	}

	if r.Nonce != nil {
		variables["preimage_nonce"] = *r.Nonce
	}

	return &GraphQLResponse{
		Query:       scripts.SET_INVOICE_PAYMENT_HASH,
		Variables:   variables,
		OutputField: "set_invoice_payment_hash",
	}
}

type SignInvoiceResponse struct {
	InvoiceId  string
	Signature  string
	RecoveryId int32
}

func (r SignInvoiceResponse) GraphqlResponse() *GraphQLResponse {
	return &GraphQLResponse{
		Query: scripts.SIGN_INVOICE_MUTATION,
		Variables: map[string]interface{}{
			"invoice_id":  r.InvoiceId,
			"signature":   r.Signature,
			"recovery_id": r.RecoveryId,
		},
		OutputField: "sign_invoice",
	}
}

type ReleasePaymentPreimageResponse struct {
	InvoiceId       string
	PaymentPreimage string
}

func (r ReleasePaymentPreimageResponse) GraphqlResponse() *GraphQLResponse {
	return &GraphQLResponse{
		Query: scripts.RELEASE_PAYMENT_PREIMAGE_MUTATION,
		Variables: map[string]interface{}{
			"invoice_id":       r.InvoiceId,
			"payment_preimage": r.PaymentPreimage,
		},
		OutputField: "release_payment_preimage",
	}
}
