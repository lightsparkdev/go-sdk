// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"time"

	"github.com/lightsparkdev/go-sdk/requester"
	"github.com/lightsparkdev/go-sdk/types"
)

// Wallet This object represents a Lightspark Wallet, tied to your Lightspark account. Wallets can be used to send or receive funds over the Lightning Network. You can retrieve this object to receive information about a specific wallet tied to your Lightspark account.
type Wallet struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"wallet_id"`

	// CreatedAt The date and time when the entity was first created.
	CreatedAt time.Time `json:"wallet_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"wallet_updated_at"`

	// LastLoginAt The date and time when the wallet user last logged in.
	LastLoginAt *time.Time `json:"wallet_last_login_at"`

	// Balances The balances that describe the funds in this wallet.
	Balances *Balances `json:"wallet_balances"`

	// ThirdPartyIdentifier The unique identifier of this wallet, as provided by the Lightspark Customer during login.
	ThirdPartyIdentifier string `json:"wallet_third_party_identifier"`

	// Account The account this wallet belongs to.
	Account *types.EntityWrapper `json:"wallet_account"`

	// Status The status of this wallet.
	Status WalletStatus `json:"wallet_status"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
}

const (
	WalletFragment = `
fragment WalletFragment on Wallet {
    __typename
    wallet_id: id
    wallet_created_at: created_at
    wallet_updated_at: updated_at
    wallet_last_login_at: last_login_at
    wallet_balances: balances {
        __typename
        balances_owned_balance: owned_balance {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
        balances_available_to_send_balance: available_to_send_balance {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
        balances_available_to_withdraw_balance: available_to_withdraw_balance {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
    }
    wallet_third_party_identifier: third_party_identifier
    wallet_account: account {
        id
    }
    wallet_status: status
}
`
)

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj Wallet) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj Wallet) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj Wallet) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj Wallet) GetTypename() string {
	return obj.Typename
}

func (obj Wallet) GetTransactions(requester *requester.Requester, first *int64, after *string, createdAfterDate *time.Time, createdBeforeDate *time.Time, statuses *[]TransactionStatus, types *[]TransactionType) (*WalletToTransactionsConnection, error) {
	query := `query FetchWalletToTransactionsConnection($entity_id: ID!, $first: Int, $after: ID, $created_after_date: DateTime, $created_before_date: DateTime, $statuses: [TransactionStatus!], $types: [TransactionType!]) {
    entity(id: $entity_id) {
        ... on Wallet {
            transactions(, first: $first, after: $after, created_after_date: $created_after_date, created_before_date: $created_before_date, statuses: $statuses, types: $types) {
                __typename
                wallet_to_transactions_connection_count: count
                wallet_to_transactions_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                wallet_to_transactions_connection_entities: entities {
                    __typename
                    ... on ChannelClosingTransaction {
                        __typename
                        channel_closing_transaction_id: id
                        channel_closing_transaction_created_at: created_at
                        channel_closing_transaction_updated_at: updated_at
                        channel_closing_transaction_status: status
                        channel_closing_transaction_resolved_at: resolved_at
                        channel_closing_transaction_amount: amount {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        channel_closing_transaction_transaction_hash: transaction_hash
                        channel_closing_transaction_fees: fees {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        channel_closing_transaction_block_hash: block_hash
                        channel_closing_transaction_block_height: block_height
                        channel_closing_transaction_destination_addresses: destination_addresses
                        channel_closing_transaction_num_confirmations: num_confirmations
                        channel_closing_transaction_channel: channel {
                            id
                        }
                    }
                    ... on ChannelOpeningTransaction {
                        __typename
                        channel_opening_transaction_id: id
                        channel_opening_transaction_created_at: created_at
                        channel_opening_transaction_updated_at: updated_at
                        channel_opening_transaction_status: status
                        channel_opening_transaction_resolved_at: resolved_at
                        channel_opening_transaction_amount: amount {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        channel_opening_transaction_transaction_hash: transaction_hash
                        channel_opening_transaction_fees: fees {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        channel_opening_transaction_block_hash: block_hash
                        channel_opening_transaction_block_height: block_height
                        channel_opening_transaction_destination_addresses: destination_addresses
                        channel_opening_transaction_num_confirmations: num_confirmations
                        channel_opening_transaction_channel: channel {
                            id
                        }
                    }
                    ... on Deposit {
                        __typename
                        deposit_id: id
                        deposit_created_at: created_at
                        deposit_updated_at: updated_at
                        deposit_status: status
                        deposit_resolved_at: resolved_at
                        deposit_amount: amount {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        deposit_transaction_hash: transaction_hash
                        deposit_fees: fees {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        deposit_block_hash: block_hash
                        deposit_block_height: block_height
                        deposit_destination_addresses: destination_addresses
                        deposit_num_confirmations: num_confirmations
                        deposit_destination: destination {
                            id
                        }
                    }
                    ... on IncomingPayment {
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
                    ... on OutgoingPayment {
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
                        outgoing_payment_is_uma: is_uma
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
                                        lightspark_node_with_o_s_k_balances: balances {
                                            __typename
                                            balances_owned_balance: owned_balance {
                                                __typename
                                                currency_amount_original_value: original_value
                                                currency_amount_original_unit: original_unit
                                                currency_amount_preferred_currency_unit: preferred_currency_unit
                                                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                            }
                                            balances_available_to_send_balance: available_to_send_balance {
                                                __typename
                                                currency_amount_original_value: original_value
                                                currency_amount_original_unit: original_unit
                                                currency_amount_preferred_currency_unit: preferred_currency_unit
                                                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                            }
                                            balances_available_to_withdraw_balance: available_to_withdraw_balance {
                                                __typename
                                                currency_amount_original_value: original_value
                                                currency_amount_original_unit: original_unit
                                                currency_amount_preferred_currency_unit: preferred_currency_unit
                                                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                            }
                                        }
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
                                        lightspark_node_with_remote_signing_balances: balances {
                                            __typename
                                            balances_owned_balance: owned_balance {
                                                __typename
                                                currency_amount_original_value: original_value
                                                currency_amount_original_unit: original_unit
                                                currency_amount_preferred_currency_unit: preferred_currency_unit
                                                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                            }
                                            balances_available_to_send_balance: available_to_send_balance {
                                                __typename
                                                currency_amount_original_value: original_value
                                                currency_amount_original_unit: original_unit
                                                currency_amount_preferred_currency_unit: preferred_currency_unit
                                                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                            }
                                            balances_available_to_withdraw_balance: available_to_withdraw_balance {
                                                __typename
                                                currency_amount_original_value: original_value
                                                currency_amount_original_unit: original_unit
                                                currency_amount_preferred_currency_unit: preferred_currency_unit
                                                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        outgoing_payment_failure_reason: failure_reason
                        outgoing_payment_failure_message: failure_message {
                            __typename
                            rich_text_text: text
                        }
                        outgoing_payment_uma_post_transaction_data: uma_post_transaction_data {
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
                        outgoing_payment_payment_preimage: payment_preimage
                    }
                    ... on RoutingTransaction {
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
                    ... on Withdrawal {
                        __typename
                        withdrawal_id: id
                        withdrawal_created_at: created_at
                        withdrawal_updated_at: updated_at
                        withdrawal_status: status
                        withdrawal_resolved_at: resolved_at
                        withdrawal_amount: amount {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        withdrawal_transaction_hash: transaction_hash
                        withdrawal_fees: fees {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        withdrawal_block_hash: block_hash
                        withdrawal_block_height: block_height
                        withdrawal_destination_addresses: destination_addresses
                        withdrawal_num_confirmations: num_confirmations
                        withdrawal_origin: origin {
                            id
                        }
                    }
                }
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":           obj.Id,
		"first":               first,
		"after":               after,
		"created_after_date":  createdAfterDate,
		"created_before_date": createdBeforeDate,
		"statuses":            statuses,
		"types":               types,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["transactions"].(map[string]interface{})
	var result *WalletToTransactionsConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Wallet) GetPaymentRequests(requester *requester.Requester, first *int64, after *string, createdAfterDate *time.Time, createdBeforeDate *time.Time) (*WalletToPaymentRequestsConnection, error) {
	query := `query FetchWalletToPaymentRequestsConnection($entity_id: ID!, $first: Int, $after: ID, $created_after_date: DateTime, $created_before_date: DateTime) {
    entity(id: $entity_id) {
        ... on Wallet {
            payment_requests(, first: $first, after: $after, created_after_date: $created_after_date, created_before_date: $created_before_date) {
                __typename
                wallet_to_payment_requests_connection_count: count
                wallet_to_payment_requests_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                wallet_to_payment_requests_connection_entities: entities {
                    __typename
                    ... on Invoice {
                        __typename
                        invoice_id: id
                        invoice_created_at: created_at
                        invoice_updated_at: updated_at
                        invoice_data: data {
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
                                    lightspark_node_with_o_s_k_balances: balances {
                                        __typename
                                        balances_owned_balance: owned_balance {
                                            __typename
                                            currency_amount_original_value: original_value
                                            currency_amount_original_unit: original_unit
                                            currency_amount_preferred_currency_unit: preferred_currency_unit
                                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                        }
                                        balances_available_to_send_balance: available_to_send_balance {
                                            __typename
                                            currency_amount_original_value: original_value
                                            currency_amount_original_unit: original_unit
                                            currency_amount_preferred_currency_unit: preferred_currency_unit
                                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                        }
                                        balances_available_to_withdraw_balance: available_to_withdraw_balance {
                                            __typename
                                            currency_amount_original_value: original_value
                                            currency_amount_original_unit: original_unit
                                            currency_amount_preferred_currency_unit: preferred_currency_unit
                                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                        }
                                    }
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
                                    lightspark_node_with_remote_signing_balances: balances {
                                        __typename
                                        balances_owned_balance: owned_balance {
                                            __typename
                                            currency_amount_original_value: original_value
                                            currency_amount_original_unit: original_unit
                                            currency_amount_preferred_currency_unit: preferred_currency_unit
                                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                        }
                                        balances_available_to_send_balance: available_to_send_balance {
                                            __typename
                                            currency_amount_original_value: original_value
                                            currency_amount_original_unit: original_unit
                                            currency_amount_preferred_currency_unit: preferred_currency_unit
                                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                        }
                                        balances_available_to_withdraw_balance: available_to_withdraw_balance {
                                            __typename
                                            currency_amount_original_value: original_value
                                            currency_amount_original_unit: original_unit
                                            currency_amount_preferred_currency_unit: preferred_currency_unit
                                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                        }
                                    }
                                }
                            }
                        }
                        invoice_status: status
                        invoice_amount_paid: amount_paid {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                    }
                }
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":           obj.Id,
		"first":               first,
		"after":               after,
		"created_after_date":  createdAfterDate,
		"created_before_date": createdBeforeDate,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["payment_requests"].(map[string]interface{})
	var result *WalletToPaymentRequestsConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Wallet) GetTotalAmountReceived(requester *requester.Requester, createdAfterDate *time.Time, createdBeforeDate *time.Time) (*CurrencyAmount, error) {
	query := `query FetchWalletTotalAmountReceived($entity_id: ID!, $created_after_date: DateTime, $created_before_date: DateTime) {
    entity(id: $entity_id) {
        ... on Wallet {
            total_amount_received(, created_after_date: $created_after_date, created_before_date: $created_before_date) {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":           obj.Id,
		"created_after_date":  createdAfterDate,
		"created_before_date": createdBeforeDate,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["total_amount_received"].(map[string]interface{})
	var result *CurrencyAmount
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Wallet) GetWithdrawalRequests(requester *requester.Requester, first *int64, after *string, statuses *[]WithdrawalRequestStatus, createdAfterDate *time.Time, createdBeforeDate *time.Time) (*WalletToWithdrawalRequestsConnection, error) {
	query := `query FetchWalletToWithdrawalRequestsConnection($entity_id: ID!, $first: Int, $after: ID, $statuses: [WithdrawalRequestStatus!], $created_after_date: DateTime, $created_before_date: DateTime) {
    entity(id: $entity_id) {
        ... on Wallet {
            withdrawal_requests(, first: $first, after: $after, statuses: $statuses, created_after_date: $created_after_date, created_before_date: $created_before_date) {
                __typename
                wallet_to_withdrawal_requests_connection_count: count
                wallet_to_withdrawal_requests_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                wallet_to_withdrawal_requests_connection_entities: entities {
                    __typename
                    withdrawal_request_id: id
                    withdrawal_request_created_at: created_at
                    withdrawal_request_updated_at: updated_at
                    withdrawal_request_requested_amount: requested_amount {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    withdrawal_request_amount: amount {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    withdrawal_request_estimated_amount: estimated_amount {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    withdrawal_request_amount_withdrawn: amount_withdrawn {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    withdrawal_request_bitcoin_address: bitcoin_address
                    withdrawal_request_withdrawal_mode: withdrawal_mode
                    withdrawal_request_status: status
                    withdrawal_request_completed_at: completed_at
                    withdrawal_request_withdrawal: withdrawal {
                        id
                    }
                }
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":           obj.Id,
		"first":               first,
		"after":               after,
		"statuses":            statuses,
		"created_after_date":  createdAfterDate,
		"created_before_date": createdBeforeDate,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["withdrawal_requests"].(map[string]interface{})
	var result *WalletToWithdrawalRequestsConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Wallet) GetTotalAmountSent(requester *requester.Requester, createdAfterDate *time.Time, createdBeforeDate *time.Time) (*CurrencyAmount, error) {
	query := `query FetchWalletTotalAmountSent($entity_id: ID!, $created_after_date: DateTime, $created_before_date: DateTime) {
    entity(id: $entity_id) {
        ... on Wallet {
            total_amount_sent(, created_after_date: $created_after_date, created_before_date: $created_before_date) {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":           obj.Id,
		"created_after_date":  createdAfterDate,
		"created_before_date": createdBeforeDate,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["total_amount_sent"].(map[string]interface{})
	var result *CurrencyAmount
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}
