// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"time"
)

// InvoiceData This object represents the data associated with a BOLT #11 invoice. You can retrieve this object to receive the relevant data associated with a specific invoice.
type InvoiceData struct {
	EncodedPaymentRequest string `json:"invoice_data_encoded_payment_request"`

	BitcoinNetwork BitcoinNetwork `json:"invoice_data_bitcoin_network"`

	// PaymentHash The payment hash of this invoice.
	PaymentHash string `json:"invoice_data_payment_hash"`

	// Amount The requested amount in this invoice. If it is equal to 0, the sender should choose the amount to send.
	Amount CurrencyAmount `json:"invoice_data_amount"`

	// CreatedAt The date and time when this invoice was created.
	CreatedAt time.Time `json:"invoice_data_created_at"`

	// ExpiresAt The date and time when this invoice will expire.
	ExpiresAt time.Time `json:"invoice_data_expires_at"`

	// Memo A short, UTF-8 encoded, description of the purpose of this invoice.
	Memo *string `json:"invoice_data_memo"`

	// Destination The lightning node that will be paid when fulfilling this invoice.
	Destination Node `json:"invoice_data_destination"`
}

const (
	InvoiceDataFragment = `
fragment InvoiceDataFragment on InvoiceData {
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
        ... on LightsparkNodeWithOSK {
            __typename
            lightspark_node_with_o_s_k_id: id
            lightspark_node_with_o_s_k_created_at: created_at
            lightspark_node_with_o_s_k_updated_at: updated_at
            lightspark_node_with_o_s_k_alias: alias
            lightspark_node_with_o_s_k_bitcoin_network: bitcoin_network
            lightspark_node_with_o_s_k_color: color
            lightspark_node_with_o_s_k_conductivity: conductivity
            lightspark_node_with_o_s_k_display_name: display_name
            lightspark_node_with_o_s_k_public_key: public_key
            lightspark_node_with_o_s_k_owner: owner {
                id
            }
            lightspark_node_with_o_s_k_status: status
            lightspark_node_with_o_s_k_total_balance: total_balance {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
            lightspark_node_with_o_s_k_total_local_balance: total_local_balance {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
            lightspark_node_with_o_s_k_local_balance: local_balance {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
            lightspark_node_with_o_s_k_remote_balance: remote_balance {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
            lightspark_node_with_o_s_k_blockchain_balance: blockchain_balance {
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
            lightspark_node_with_o_s_k_uma_prescreening_utxos: uma_prescreening_utxos
            lightspark_node_with_o_s_k_encrypted_signing_private_key: encrypted_signing_private_key {
                __typename
                secret_encrypted_value: encrypted_value
                secret_cipher: cipher
            }
        }
        ... on LightsparkNodeWithRemoteSigning {
            __typename
            lightspark_node_with_remote_signing_id: id
            lightspark_node_with_remote_signing_created_at: created_at
            lightspark_node_with_remote_signing_updated_at: updated_at
            lightspark_node_with_remote_signing_alias: alias
            lightspark_node_with_remote_signing_bitcoin_network: bitcoin_network
            lightspark_node_with_remote_signing_color: color
            lightspark_node_with_remote_signing_conductivity: conductivity
            lightspark_node_with_remote_signing_display_name: display_name
            lightspark_node_with_remote_signing_public_key: public_key
            lightspark_node_with_remote_signing_owner: owner {
                id
            }
            lightspark_node_with_remote_signing_status: status
            lightspark_node_with_remote_signing_total_balance: total_balance {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
            lightspark_node_with_remote_signing_total_local_balance: total_local_balance {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
            lightspark_node_with_remote_signing_local_balance: local_balance {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
            lightspark_node_with_remote_signing_remote_balance: remote_balance {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
            lightspark_node_with_remote_signing_blockchain_balance: blockchain_balance {
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
            lightspark_node_with_remote_signing_uma_prescreening_utxos: uma_prescreening_utxos
        }
    }
}
`
)

func (obj InvoiceData) GetEncodedPaymentRequest() string {
	return obj.EncodedPaymentRequest
}

func (obj InvoiceData) GetBitcoinNetwork() BitcoinNetwork {
	return obj.BitcoinNetwork
}

type InvoiceDataJSON struct {
	EncodedPaymentRequest string `json:"invoice_data_encoded_payment_request"`

	BitcoinNetwork BitcoinNetwork `json:"invoice_data_bitcoin_network"`

	// PaymentHash The payment hash of this invoice.
	PaymentHash string `json:"invoice_data_payment_hash"`

	// Amount The requested amount in this invoice. If it is equal to 0, the sender should choose the amount to send.
	Amount CurrencyAmount `json:"invoice_data_amount"`

	// CreatedAt The date and time when this invoice was created.
	CreatedAt time.Time `json:"invoice_data_created_at"`

	// ExpiresAt The date and time when this invoice will expire.
	ExpiresAt time.Time `json:"invoice_data_expires_at"`

	// Memo A short, UTF-8 encoded, description of the purpose of this invoice.
	Memo *string `json:"invoice_data_memo"`

	// Destination The lightning node that will be paid when fulfilling this invoice.
	Destination map[string]interface{} `json:"invoice_data_destination"`
}

func (data *InvoiceData) UnmarshalJSON(dataBytes []byte) error {
	var temp InvoiceDataJSON
	if err := json.Unmarshal(dataBytes, &temp); err != nil {
		return err
	}

	data.EncodedPaymentRequest = temp.EncodedPaymentRequest

	data.BitcoinNetwork = temp.BitcoinNetwork

	data.PaymentHash = temp.PaymentHash

	data.Amount = temp.Amount

	data.CreatedAt = temp.CreatedAt

	data.ExpiresAt = temp.ExpiresAt

	data.Memo = temp.Memo

	Destination, err := NodeUnmarshal(temp.Destination)
	if err != nil {
		return err
	}
	data.Destination = Destination

	return nil
}
