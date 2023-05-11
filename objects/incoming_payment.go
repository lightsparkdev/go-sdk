// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"lightspark/requester"
	"lightspark/types"
	"time"
)

// A transaction that was sent to a Lightspark node on the Lightning Network.
type IncomingPayment struct {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"incoming_payment_id"`

	// The date and time when this transaction was initiated.
	CreatedAt time.Time `json:"incoming_payment_created_at"`

	// The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"incoming_payment_updated_at"`

	// The current status of this transaction.
	Status TransactionStatus `json:"incoming_payment_status"`

	// The date and time when this transaction was completed or failed.
	ResolvedAt *time.Time `json:"incoming_payment_resolved_at"`

	// The amount of money involved in this transaction.
	Amount CurrencyAmount `json:"incoming_payment_amount"`

	// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
	TransactionHash *string `json:"incoming_payment_transaction_hash"`

	// If known, the Lightspark node this payment originated from.
	Origin *types.EntityWrapper `json:"incoming_payment_origin"`

	// The recipient Lightspark node this payment was sent to.
	Destination types.EntityWrapper `json:"incoming_payment_destination"`

	// The optional payment request for this incoming payment, which will be null if the payment is sent through keysend.
	PaymentRequest *types.EntityWrapper `json:"incoming_payment_payment_request"`
}

const (
	IncomingPaymentFragment = `
fragment IncomingPaymentFragment on IncomingPayment {
    __typename
    incoming_payment_id: id
    incoming_payment_created_at: created_at
    incoming_payment_updated_at: updated_at
    incoming_payment_status: status
    incoming_payment_resolved_at: resolved_at
    incoming_payment_amount: amount {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    incoming_payment_transaction_hash: transaction_hash
    incoming_payment_origin: origin {
        id
    }
    incoming_payment_destination: destination {
        id
    }
    incoming_payment_payment_request: payment_request {
        id
    }
}
`
)

// The current status of this transaction.
func (obj IncomingPayment) GetStatus() TransactionStatus {
	return obj.Status
}

// The date and time when this transaction was completed or failed.
func (obj IncomingPayment) GetResolvedAt() *time.Time {
	return obj.ResolvedAt
}

// The amount of money involved in this transaction.
func (obj IncomingPayment) GetAmount() CurrencyAmount {
	return obj.Amount
}

// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
func (obj IncomingPayment) GetTransactionHash() *string {
	return obj.TransactionHash
}

// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj IncomingPayment) GetId() string {
	return obj.Id
}

// The date and time when the entity was first created.
func (obj IncomingPayment) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// The date and time when the entity was last updated.
func (obj IncomingPayment) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj IncomingPayment) GetAttempts(requester *requester.Requester, first *int64, statuses *[]IncomingPaymentAttemptStatus) (*IncomingPaymentToAttemptsConnection, error) {
	query := `query FetchIncomingPaymentToAttemptsConnection($entity_id: ID!, $first: Int, $statuses: [IncomingPaymentAttemptStatus!]) {
    entity(id: $entity_id) {
        ... on IncomingPayment {
            attempts(, first: $first, statuses: $statuses) {
                __typename
                incoming_payment_to_attempts_connection_count: count
                incoming_payment_to_attempts_connection_entities: entities {
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
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id": obj.Id,
		"first":     first,
		"statuses":  statuses,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["attempts"].(map[string]interface{})
	var result *IncomingPaymentToAttemptsConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}
