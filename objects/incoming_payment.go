// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"time"

	"github.com/lightsparkdev/go-sdk/requester"
	"github.com/lightsparkdev/go-sdk/types"
)

// IncomingPayment This object represents any payment sent to a Lightspark node on the Lightning Network. You can retrieve this object to receive payment related information about a specific payment received by a Lightspark node.
type IncomingPayment struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"incoming_payment_id"`

	// CreatedAt The date and time when this transaction was initiated.
	CreatedAt time.Time `json:"incoming_payment_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"incoming_payment_updated_at"`

	// Status The current status of this transaction.
	Status TransactionStatus `json:"incoming_payment_status"`

	// ResolvedAt The date and time when this transaction was completed or failed.
	ResolvedAt *time.Time `json:"incoming_payment_resolved_at"`

	// Amount The amount of money involved in this transaction.
	Amount CurrencyAmount `json:"incoming_payment_amount"`

	// TransactionHash The hash of this transaction, so it can be uniquely identified on the Lightning Network.
	TransactionHash *string `json:"incoming_payment_transaction_hash"`

	// IsUma Whether this payment is an UMA payment or not. NOTE: this field is only set if the invoice that is being paid has been created using the recommended `create_uma_invoice` function.
	IsUma bool `json:"incoming_payment_is_uma"`

	// Destination The recipient Lightspark node this payment was sent to.
	Destination types.EntityWrapper `json:"incoming_payment_destination"`

	// PaymentRequest The optional payment request for this incoming payment, which will be null if the payment is sent through keysend.
	PaymentRequest *types.EntityWrapper `json:"incoming_payment_payment_request"`

	// UmaPostTransactionData The post transaction data which can be used in KYT payment registration.
	UmaPostTransactionData *[]PostTransactionData `json:"incoming_payment_uma_post_transaction_data"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
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
    incoming_payment_is_uma: is_uma
    incoming_payment_destination: destination {
        id
    }
    incoming_payment_payment_request: payment_request {
        id
    }
    incoming_payment_uma_post_transaction_data: uma_post_transaction_data {
        __typename
        post_transaction_data_utxo: utxo
        post_transaction_data_amount: amount {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
    }
}
`
)

// GetStatus The current status of this transaction.
func (obj IncomingPayment) GetStatus() TransactionStatus {
	return obj.Status
}

// GetResolvedAt The date and time when this transaction was completed or failed.
func (obj IncomingPayment) GetResolvedAt() *time.Time {
	return obj.ResolvedAt
}

// GetAmount The amount of money involved in this transaction.
func (obj IncomingPayment) GetAmount() CurrencyAmount {
	return obj.Amount
}

// GetTransactionHash The hash of this transaction, so it can be uniquely identified on the Lightning Network.
func (obj IncomingPayment) GetTransactionHash() *string {
	return obj.TransactionHash
}

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj IncomingPayment) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj IncomingPayment) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj IncomingPayment) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj IncomingPayment) GetTypename() string {
	return obj.Typename
}

func (obj IncomingPayment) GetAttempts(requester *requester.Requester, first *int64, statuses *[]IncomingPaymentAttemptStatus, after *string) (*IncomingPaymentToAttemptsConnection, error) {
	query := `query FetchIncomingPaymentToAttemptsConnection($entity_id: ID!, $first: Int, $statuses: [IncomingPaymentAttemptStatus!], $after: String) {
    entity(id: $entity_id) {
        ... on IncomingPayment {
            attempts(, first: $first, statuses: $statuses, after: $after) {
                __typename
                incoming_payment_to_attempts_connection_count: count
                incoming_payment_to_attempts_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
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
		"after":     after,
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
