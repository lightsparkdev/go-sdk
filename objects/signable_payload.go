// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"time"

	"github.com/lightsparkdev/go-sdk/types"
)

type SignablePayload struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"signable_payload_id"`

	// CreatedAt The date and time when the entity was first created.
	CreatedAt time.Time `json:"signable_payload_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"signable_payload_updated_at"`

	// Payload The payload that needs to be signed.
	Payload string `json:"signable_payload_payload"`

	// DerivationPath The consistent method for generating the same set of accounts and wallets for a given private key
	DerivationPath string `json:"signable_payload_derivation_path"`

	// Status The status of the payload.
	Status SignablePayloadStatus `json:"signable_payload_status"`

	// AddTweak The tweak value to add.
	AddTweak *string `json:"signable_payload_add_tweak"`

	// MulTweak The tweak value to multiply.
	MulTweak *string `json:"signable_payload_mul_tweak"`

	// Signable The signable this payload belongs to.
	Signable types.EntityWrapper `json:"signable_payload_signable"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
}

const (
	SignablePayloadFragment = `
fragment SignablePayloadFragment on SignablePayload {
    __typename
    signable_payload_id: id
    signable_payload_created_at: created_at
    signable_payload_updated_at: updated_at
    signable_payload_payload: payload
    signable_payload_derivation_path: derivation_path
    signable_payload_status: status
    signable_payload_add_tweak: add_tweak
    signable_payload_mul_tweak: mul_tweak
    signable_payload_signable: signable {
        id
    }
}
`
)

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj SignablePayload) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj SignablePayload) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj SignablePayload) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj SignablePayload) GetTypename() string {
	return obj.Typename
}
