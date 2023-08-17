// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"time"

	"github.com/lightsparkdev/go-sdk/types"
)

// Hop This object represents a specific node that existed on a particular payment route. You can retrieve this object to get information about a node on a particular payment path and all payment-relevant information for that node.
type Hop struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"hop_id"`

	// CreatedAt The date and time when the entity was first created.
	CreatedAt time.Time `json:"hop_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"hop_updated_at"`

	// Destination The destination node of the hop.
	Destination *types.EntityWrapper `json:"hop_destination"`

	// Index The zero-based index position of this hop in the path
	Index int64 `json:"hop_index"`

	// PublicKey The public key of the node to which the hop is bound.
	PublicKey *string `json:"hop_public_key"`

	// AmountToForward The amount that is to be forwarded to the destination node.
	AmountToForward *CurrencyAmount `json:"hop_amount_to_forward"`

	// Fee The fees to be collected by the source node for forwarding the payment over the hop.
	Fee *CurrencyAmount `json:"hop_fee"`

	// ExpiryBlockHeight The block height at which an unsettled HTLC is considered expired.
	ExpiryBlockHeight *int64 `json:"hop_expiry_block_height"`
}

const (
	HopFragment = `
fragment HopFragment on Hop {
    __typename
    hop_id: id
    hop_created_at: created_at
    hop_updated_at: updated_at
    hop_destination: destination {
        id
    }
    hop_index: index
    hop_public_key: public_key
    hop_amount_to_forward: amount_to_forward {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    hop_fee: fee {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    hop_expiry_block_height: expiry_block_height
}
`
)

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj Hop) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj Hop) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj Hop) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}
