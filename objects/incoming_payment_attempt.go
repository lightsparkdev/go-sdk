// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "time"

// IncomingPaymentAttempt This object represents any attempted payment sent to a Lightspark node on the Lightning Network. You can retrieve this object to receive payment related information about a specific incoming payment attempt.
type IncomingPaymentAttempt struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"incoming_payment_attempt_id"`

	// CreatedAt The date and time when the entity was first created.
	CreatedAt time.Time `json:"incoming_payment_attempt_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"incoming_payment_attempt_updated_at"`

	// Status The status of the incoming payment attempt.
	Status IncomingPaymentAttemptStatus `json:"incoming_payment_attempt_status"`

	// ResolvedAt The time the incoming payment attempt failed or succeeded.
	ResolvedAt *time.Time `json:"incoming_payment_attempt_resolved_at"`

	// Amount The total amount of that was attempted to send.
	Amount CurrencyAmount `json:"incoming_payment_attempt_amount"`

	// Channel The channel this attempt was made on.
	Channel types.EntityWrapper `json:"incoming_payment_attempt_channel"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
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

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj IncomingPaymentAttempt) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj IncomingPaymentAttempt) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj IncomingPaymentAttempt) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj IncomingPaymentAttempt) GetTypename() string {
	return obj.Typename
}
