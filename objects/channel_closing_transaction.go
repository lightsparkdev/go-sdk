// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"lightspark/types"
	"time"
)

// The transaction on Bitcoin blockchain to close a channel on Lightning Network where the balances are allocated back to local and remote nodes.
type ChannelClosingTransaction struct {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"channel_closing_transaction_id"`

	// The date and time when this transaction was initiated.
	CreatedAt time.Time `json:"channel_closing_transaction_created_at"`

	// The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"channel_closing_transaction_updated_at"`

	// The current status of this transaction.
	Status TransactionStatus `json:"channel_closing_transaction_status"`

	// The date and time when this transaction was completed or failed.
	ResolvedAt *time.Time `json:"channel_closing_transaction_resolved_at"`

	// The amount of money involved in this transaction.
	Amount CurrencyAmount `json:"channel_closing_transaction_amount"`

	// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
	TransactionHash *string `json:"channel_closing_transaction_transaction_hash"`

	// The fees that were paid by the wallet sending the transaction to commit it to the Bitcoin blockchain.
	Fees *CurrencyAmount `json:"channel_closing_transaction_fees"`

	// The hash of the block that included this transaction. This will be null for unconfirmed transactions.
	BlockHash *string `json:"channel_closing_transaction_block_hash"`

	// The height of the block that included this transaction. This will be zero for unconfirmed transactions.
	BlockHeight int64 `json:"channel_closing_transaction_block_height"`

	// The Bitcoin blockchain addresses this transaction was sent to.
	DestinationAddresses []string `json:"channel_closing_transaction_destination_addresses"`

	// The number of blockchain confirmations for this transaction in real time.
	NumConfirmations *int64 `json:"channel_closing_transaction_num_confirmations"`

	// If known, the channel this transaction is closing.
	Channel *types.EntityWrapper `json:"channel_closing_transaction_channel"`
}

const (
	ChannelClosingTransactionFragment = `
fragment ChannelClosingTransactionFragment on ChannelClosingTransaction {
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
`
)

// The fees that were paid by the wallet sending the transaction to commit it to the Bitcoin blockchain.
func (obj ChannelClosingTransaction) GetFees() *CurrencyAmount {
	return obj.Fees
}

// The hash of the block that included this transaction. This will be null for unconfirmed transactions.
func (obj ChannelClosingTransaction) GetBlockHash() *string {
	return obj.BlockHash
}

// The height of the block that included this transaction. This will be zero for unconfirmed transactions.
func (obj ChannelClosingTransaction) GetBlockHeight() int64 {
	return obj.BlockHeight
}

// The Bitcoin blockchain addresses this transaction was sent to.
func (obj ChannelClosingTransaction) GetDestinationAddresses() []string {
	return obj.DestinationAddresses
}

// The number of blockchain confirmations for this transaction in real time.
func (obj ChannelClosingTransaction) GetNumConfirmations() *int64 {
	return obj.NumConfirmations
}

// The current status of this transaction.
func (obj ChannelClosingTransaction) GetStatus() TransactionStatus {
	return obj.Status
}

// The date and time when this transaction was completed or failed.
func (obj ChannelClosingTransaction) GetResolvedAt() *time.Time {
	return obj.ResolvedAt
}

// The amount of money involved in this transaction.
func (obj ChannelClosingTransaction) GetAmount() CurrencyAmount {
	return obj.Amount
}

// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
func (obj ChannelClosingTransaction) GetTransactionHash() *string {
	return obj.TransactionHash
}

// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj ChannelClosingTransaction) GetId() string {
	return obj.Id
}

// The date and time when the entity was first created.
func (obj ChannelClosingTransaction) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// The date and time when the entity was last updated.
func (obj ChannelClosingTransaction) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}
