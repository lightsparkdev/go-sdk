// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"time"

	"github.com/lightsparkdev/go-sdk/types"
)

// An attempt for a payment over a route from sender node to recipient node.
type IncomingPaymentAttempt struct {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"incoming_payment_attempt_id"`

	// The date and time when the entity was first created.
	CreatedAt time.Time `json:"incoming_payment_attempt_created_at"`

	// The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"incoming_payment_attempt_updated_at"`

	// The status of the incoming payment attempt.
	Status IncomingPaymentAttemptStatus `json:"incoming_payment_attempt_status"`

	// The time the incoming payment attempt failed or succeeded.
	ResolvedAt *time.Time `json:"incoming_payment_attempt_resolved_at"`

	// The total amount of that was attempted to send.
	Amount CurrencyAmount `json:"incoming_payment_attempt_amount"`

	// The channel this attempt was made on.
	Channel types.EntityWrapper `json:"incoming_payment_attempt_channel"`
}

const (
	IncomingPaymentAttemptFragment = `
fragment IncomingPaymentAttemptFragment on IncomingPaymentAttempt {
    __typename
    incoming_payment_attempt_id: id
    incoming_payment_attempt_created_at: created_at
    incoming_payment_attempt_updated_at: updated_at
    incoming_payment_attempt_status: status
    incoming_payment_attempt_resolved_at: resolved_at
    incoming_payment_attempt_amount: amount {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    incoming_payment_attempt_channel: channel {
        id
    }
}
`
)

// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj IncomingPaymentAttempt) GetId() string {
	return obj.Id
}

// The date and time when the entity was first created.
func (obj IncomingPaymentAttempt) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// The date and time when the entity was last updated.
func (obj IncomingPaymentAttempt) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}
