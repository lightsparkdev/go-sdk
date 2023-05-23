// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"time"

	"github.com/lightsparkdev/go-sdk/requester"
	"github.com/lightsparkdev/go-sdk/types"
)

// A transaction that was sent from a Lightspark node on the Lightning Network.
type OutgoingPayment struct {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"outgoing_payment_id"`

	// The date and time when this transaction was initiated.
	CreatedAt time.Time `json:"outgoing_payment_created_at"`

	// The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"outgoing_payment_updated_at"`

	// The current status of this transaction.
	Status TransactionStatus `json:"outgoing_payment_status"`

	// The date and time when this transaction was completed or failed.
	ResolvedAt *time.Time `json:"outgoing_payment_resolved_at"`

	// The amount of money involved in this transaction.
	Amount CurrencyAmount `json:"outgoing_payment_amount"`

	// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
	TransactionHash *string `json:"outgoing_payment_transaction_hash"`

	// The Lightspark node this payment originated from.
	Origin types.EntityWrapper `json:"outgoing_payment_origin"`

	// If known, the final recipient node this payment was sent to.
	Destination *types.EntityWrapper `json:"outgoing_payment_destination"`

	// The fees paid by the sender node to send the payment.
	Fees *CurrencyAmount `json:"outgoing_payment_fees"`

	// The data of the payment request that was paid by this transaction, if known.
	PaymentRequestData *PaymentRequestData `json:"outgoing_payment_payment_request_data"`

	// If applicable, the reason why the payment failed.
	FailureReason *PaymentFailureReason `json:"outgoing_payment_failure_reason"`

	// If applicable, user-facing error message describing why the payment failed.
	FailureMessage *RichText `json:"outgoing_payment_failure_message"`
}

const (
	OutgoingPaymentFragment = `
fragment OutgoingPaymentFragment on OutgoingPayment {
    __typename
    outgoing_payment_id: id
    outgoing_payment_created_at: created_at
    outgoing_payment_updated_at: updated_at
    outgoing_payment_status: status
    outgoing_payment_resolved_at: resolved_at
    outgoing_payment_amount: amount {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    outgoing_payment_transaction_hash: transaction_hash
    outgoing_payment_origin: origin {
        id
    }
    outgoing_payment_destination: destination {
        id
    }
    outgoing_payment_fees: fees {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    outgoing_payment_payment_request_data: payment_request_data {
        __typename
        ... on InvoiceData {
            __typename
            invoice_data_encoded_payment_request: encoded_payment_request
            invoice_data_bitcoin_network: bitcoin_network
            invoice_data_payment_hash: payment_hash
            invoice_data_amount: amount {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
            invoice_data_created_at: created_at
            invoice_data_expires_at: expires_at
            invoice_data_memo: memo
            invoice_data_destination: destination {
                __typename
                ... on GraphNode {
                    __typename
                    graph_node_id: id
                    graph_node_created_at: created_at
                    graph_node_updated_at: updated_at
                    graph_node_alias: alias
                    graph_node_bitcoin_network: bitcoin_network
                    graph_node_color: color
                    graph_node_conductivity: conductivity
                    graph_node_display_name: display_name
                    graph_node_public_key: public_key
                }
                ... on LightsparkNode {
                    __typename
                    lightspark_node_id: id
                    lightspark_node_created_at: created_at
                    lightspark_node_updated_at: updated_at
                    lightspark_node_alias: alias
                    lightspark_node_bitcoin_network: bitcoin_network
                    lightspark_node_color: color
                    lightspark_node_conductivity: conductivity
                    lightspark_node_display_name: display_name
                    lightspark_node_public_key: public_key
                    lightspark_node_account: account {
                        id
                    }
                    lightspark_node_blockchain_balance: blockchain_balance {
                        __typename
                        blockchain_balance_total_balance: total_balance {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        blockchain_balance_confirmed_balance: confirmed_balance {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        blockchain_balance_unconfirmed_balance: unconfirmed_balance {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        blockchain_balance_locked_balance: locked_balance {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        blockchain_balance_required_reserve: required_reserve {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        blockchain_balance_available_balance: available_balance {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                    }
                    lightspark_node_encrypted_signing_private_key: encrypted_signing_private_key {
                        __typename
                        secret_encrypted_value: encrypted_value
                        secret_cipher: cipher
                    }
                    lightspark_node_total_balance: total_balance {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    lightspark_node_total_local_balance: total_local_balance {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    lightspark_node_local_balance: local_balance {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    lightspark_node_purpose: purpose
                    lightspark_node_remote_balance: remote_balance {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    lightspark_node_status: status
                }
            }
        }
    }
    outgoing_payment_failure_reason: failure_reason
    outgoing_payment_failure_message: failure_message {
        __typename
        rich_text_text: text
    }
}
`
)

// The current status of this transaction.
func (obj OutgoingPayment) GetStatus() TransactionStatus {
	return obj.Status
}

// The date and time when this transaction was completed or failed.
func (obj OutgoingPayment) GetResolvedAt() *time.Time {
	return obj.ResolvedAt
}

// The amount of money involved in this transaction.
func (obj OutgoingPayment) GetAmount() CurrencyAmount {
	return obj.Amount
}

// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
func (obj OutgoingPayment) GetTransactionHash() *string {
	return obj.TransactionHash
}

// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj OutgoingPayment) GetId() string {
	return obj.Id
}

// The date and time when the entity was first created.
func (obj OutgoingPayment) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// The date and time when the entity was last updated.
func (obj OutgoingPayment) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj OutgoingPayment) GetAttempts(requester *requester.Requester, first *int64) (*OutgoingPaymentToAttemptsConnection, error) {
	query := `query FetchOutgoingPaymentToAttemptsConnection($entity_id: ID!, $first: Int) {
    entity(id: $entity_id) {
        ... on OutgoingPayment {
            attempts(, first: $first) {
                __typename
                outgoing_payment_to_attempts_connection_count: count
                outgoing_payment_to_attempts_connection_entities: entities {
                    __typename
                    outgoing_payment_attempt_id: id
                    outgoing_payment_attempt_created_at: created_at
                    outgoing_payment_attempt_updated_at: updated_at
                    outgoing_payment_attempt_status: status
                    outgoing_payment_attempt_failure_code: failure_code
                    outgoing_payment_attempt_failure_source_index: failure_source_index
                    outgoing_payment_attempt_resolved_at: resolved_at
                    outgoing_payment_attempt_amount: amount {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    outgoing_payment_attempt_fees: fees {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    outgoing_payment_attempt_outgoing_payment: outgoing_payment {
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
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["attempts"].(map[string]interface{})
	var result *OutgoingPaymentToAttemptsConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

type OutgoingPaymentJSON struct {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"outgoing_payment_id"`

	// The date and time when this transaction was initiated.
	CreatedAt time.Time `json:"outgoing_payment_created_at"`

	// The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"outgoing_payment_updated_at"`

	// The current status of this transaction.
	Status TransactionStatus `json:"outgoing_payment_status"`

	// The date and time when this transaction was completed or failed.
	ResolvedAt *time.Time `json:"outgoing_payment_resolved_at"`

	// The amount of money involved in this transaction.
	Amount CurrencyAmount `json:"outgoing_payment_amount"`

	// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
	TransactionHash *string `json:"outgoing_payment_transaction_hash"`

	// The Lightspark node this payment originated from.
	Origin types.EntityWrapper `json:"outgoing_payment_origin"`

	// If known, the final recipient node this payment was sent to.
	Destination *types.EntityWrapper `json:"outgoing_payment_destination"`

	// The fees paid by the sender node to send the payment.
	Fees *CurrencyAmount `json:"outgoing_payment_fees"`

	// The data of the payment request that was paid by this transaction, if known.
	PaymentRequestData map[string]interface{} `json:"outgoing_payment_payment_request_data"`

	// If applicable, the reason why the payment failed.
	FailureReason *PaymentFailureReason `json:"outgoing_payment_failure_reason"`

	// If applicable, user-facing error message describing why the payment failed.
	FailureMessage *RichText `json:"outgoing_payment_failure_message"`
}

func (data *OutgoingPayment) UnmarshalJSON(dataBytes []byte) error {
	var temp OutgoingPaymentJSON
	if err := json.Unmarshal(dataBytes, &temp); err != nil {
		return err
	}

	data.Id = temp.Id

	data.CreatedAt = temp.CreatedAt

	data.UpdatedAt = temp.UpdatedAt

	data.Status = temp.Status

	data.ResolvedAt = temp.ResolvedAt

	data.Amount = temp.Amount

	data.TransactionHash = temp.TransactionHash

	data.Origin = temp.Origin

	data.Destination = temp.Destination

	data.Fees = temp.Fees

	PaymentRequestData, err := PaymentRequestDataUnmarshal(temp.PaymentRequestData)
	if err != nil {
		return err
	}
	data.PaymentRequestData = &PaymentRequestData

	data.FailureReason = temp.FailureReason

	data.FailureMessage = temp.FailureMessage

	return nil
}
