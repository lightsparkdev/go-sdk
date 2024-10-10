// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"time"

	"github.com/lightsparkdev/go-sdk/types"
)

// ChannelOpeningTransaction This is an object representing a transaction which opens a channel on the Lightning Network. This object occurs only for channels funded by the local Lightspark node.
type ChannelOpeningTransaction struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"channel_opening_transaction_id"`

	// CreatedAt The date and time when this transaction was initiated.
	CreatedAt time.Time `json:"channel_opening_transaction_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"channel_opening_transaction_updated_at"`

	// Status The current status of this transaction.
	Status TransactionStatus `json:"channel_opening_transaction_status"`

	// ResolvedAt The date and time when this transaction was completed or failed.
	ResolvedAt *time.Time `json:"channel_opening_transaction_resolved_at"`

	// Amount The amount of money involved in this transaction.
	Amount CurrencyAmount `json:"channel_opening_transaction_amount"`

	// TransactionHash The hash of this transaction, so it can be uniquely identified on the Lightning Network.
	TransactionHash *string `json:"channel_opening_transaction_transaction_hash"`

	// Fees The fees that were paid by the node for this transaction.
	Fees *CurrencyAmount `json:"channel_opening_transaction_fees"`

	// BlockHash The hash of the block that included this transaction. This will be null for unconfirmed transactions.
	BlockHash *string `json:"channel_opening_transaction_block_hash"`

	// BlockHeight The height of the block that included this transaction. This will be zero for unconfirmed transactions.
	BlockHeight int64 `json:"channel_opening_transaction_block_height"`

	// DestinationAddresses The Bitcoin blockchain addresses this transaction was sent to.
	DestinationAddresses []string `json:"channel_opening_transaction_destination_addresses"`

	// NumConfirmations The number of blockchain confirmations for this transaction in real time.
	NumConfirmations *int64 `json:"channel_opening_transaction_num_confirmations"`

	// Channel If known, the channel this transaction is opening.
	Channel *types.EntityWrapper `json:"channel_opening_transaction_channel"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
}

const (
	ChannelOpeningTransactionFragment = `
fragment ChannelOpeningTransactionFragment on ChannelOpeningTransaction {
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
`
)

// GetFees The fees that were paid by the node for this transaction.
func (obj ChannelOpeningTransaction) GetFees() *CurrencyAmount {
	return obj.Fees
}

// GetBlockHash The hash of the block that included this transaction. This will be null for unconfirmed transactions.
func (obj ChannelOpeningTransaction) GetBlockHash() *string {
	return obj.BlockHash
}

// GetBlockHeight The height of the block that included this transaction. This will be zero for unconfirmed transactions.
func (obj ChannelOpeningTransaction) GetBlockHeight() int64 {
	return obj.BlockHeight
}

// GetDestinationAddresses The Bitcoin blockchain addresses this transaction was sent to.
func (obj ChannelOpeningTransaction) GetDestinationAddresses() []string {
	return obj.DestinationAddresses
}

// GetNumConfirmations The number of blockchain confirmations for this transaction in real time.
func (obj ChannelOpeningTransaction) GetNumConfirmations() *int64 {
	return obj.NumConfirmations
}

// GetStatus The current status of this transaction.
func (obj ChannelOpeningTransaction) GetStatus() TransactionStatus {
	return obj.Status
}

// GetResolvedAt The date and time when this transaction was completed or failed.
func (obj ChannelOpeningTransaction) GetResolvedAt() *time.Time {
	return obj.ResolvedAt
}

// GetAmount The amount of money involved in this transaction.
func (obj ChannelOpeningTransaction) GetAmount() CurrencyAmount {
	return obj.Amount
}

// GetTransactionHash The hash of this transaction, so it can be uniquely identified on the Lightning Network.
func (obj ChannelOpeningTransaction) GetTransactionHash() *string {
	return obj.TransactionHash
}

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj ChannelOpeningTransaction) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj ChannelOpeningTransaction) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj ChannelOpeningTransaction) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj ChannelOpeningTransaction) GetTypename() string {
	return obj.Typename
}
