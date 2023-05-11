// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"lightspark/types"
	"time"
)

// A transaction that was forwarded through a Lightspark node on the Lightning Network.
type RoutingTransaction struct {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"routing_transaction_id"`

	// The date and time when this transaction was initiated.
	CreatedAt time.Time `json:"routing_transaction_created_at"`

	// The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"routing_transaction_updated_at"`

	// The current status of this transaction.
	Status TransactionStatus `json:"routing_transaction_status"`

	// The date and time when this transaction was completed or failed.
	ResolvedAt *time.Time `json:"routing_transaction_resolved_at"`

	// The amount of money involved in this transaction.
	Amount CurrencyAmount `json:"routing_transaction_amount"`

	// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
	TransactionHash *string `json:"routing_transaction_transaction_hash"`

	// If known, the channel this transaction was received from.
	IncomingChannel *types.EntityWrapper `json:"routing_transaction_incoming_channel"`

	// If known, the channel this transaction was forwarded to.
	OutgoingChannel *types.EntityWrapper `json:"routing_transaction_outgoing_channel"`

	// The fees collected by the node when routing this transaction. We subtract the outgoing amount to the incoming amount to determine how much fees were collected.
	Fees *CurrencyAmount `json:"routing_transaction_fees"`

	// If applicable, user-facing error message describing why the routing failed.
	FailureMessage *RichText `json:"routing_transaction_failure_message"`

	// If applicable, the reason why the routing failed.
	FailureReason *RoutingTransactionFailureReason `json:"routing_transaction_failure_reason"`
}

const (
	RoutingTransactionFragment = `
fragment RoutingTransactionFragment on RoutingTransaction {
    __typename
    routing_transaction_id: id
    routing_transaction_created_at: created_at
    routing_transaction_updated_at: updated_at
    routing_transaction_status: status
    routing_transaction_resolved_at: resolved_at
    routing_transaction_amount: amount {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    routing_transaction_transaction_hash: transaction_hash
    routing_transaction_incoming_channel: incoming_channel {
        id
    }
    routing_transaction_outgoing_channel: outgoing_channel {
        id
    }
    routing_transaction_fees: fees {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    routing_transaction_failure_message: failure_message {
        __typename
        rich_text_text: text
    }
    routing_transaction_failure_reason: failure_reason
}
`
)

// The current status of this transaction.
func (obj RoutingTransaction) GetStatus() TransactionStatus {
	return obj.Status
}

// The date and time when this transaction was completed or failed.
func (obj RoutingTransaction) GetResolvedAt() *time.Time {
	return obj.ResolvedAt
}

// The amount of money involved in this transaction.
func (obj RoutingTransaction) GetAmount() CurrencyAmount {
	return obj.Amount
}

// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
func (obj RoutingTransaction) GetTransactionHash() *string {
	return obj.TransactionHash
}

// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj RoutingTransaction) GetId() string {
	return obj.Id
}

// The date and time when the entity was first created.
func (obj RoutingTransaction) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// The date and time when the entity was last updated.
func (obj RoutingTransaction) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}
