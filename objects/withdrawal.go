// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"time"

	"github.com/lightsparkdev/go-sdk/types"
)

// Withdrawal This object represents an L1 withdrawal from your Lightspark Node to any Bitcoin wallet. You can retrieve this object to receive detailed information about any L1 withdrawal associated with your Lightspark Node or account.
type Withdrawal struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"withdrawal_id"`

	// CreatedAt The date and time when this transaction was initiated.
	CreatedAt time.Time `json:"withdrawal_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"withdrawal_updated_at"`

	// Status The current status of this transaction.
	Status TransactionStatus `json:"withdrawal_status"`

	// ResolvedAt The date and time when this transaction was completed or failed.
	ResolvedAt *time.Time `json:"withdrawal_resolved_at"`

	// Amount The amount of money involved in this transaction.
	Amount CurrencyAmount `json:"withdrawal_amount"`

	// TransactionHash The hash of this transaction, so it can be uniquely identified on the Lightning Network.
	TransactionHash *string `json:"withdrawal_transaction_hash"`

	// Fees The fees that were paid by the wallet sending the transaction to commit it to the Bitcoin blockchain.
	Fees *CurrencyAmount `json:"withdrawal_fees"`

	// BlockHash The hash of the block that included this transaction. This will be null for unconfirmed transactions.
	BlockHash *string `json:"withdrawal_block_hash"`

	// BlockHeight The height of the block that included this transaction. This will be zero for unconfirmed transactions.
	BlockHeight int64 `json:"withdrawal_block_height"`

	// DestinationAddresses The Bitcoin blockchain addresses this transaction was sent to.
	DestinationAddresses []string `json:"withdrawal_destination_addresses"`

	// NumConfirmations The number of blockchain confirmations for this transaction in real time.
	NumConfirmations *int64 `json:"withdrawal_num_confirmations"`

	// Origin The Lightspark node this withdrawal originated from.
	Origin types.EntityWrapper `json:"withdrawal_origin"`
}

const (
	WithdrawalFragment = `
fragment WithdrawalFragment on Withdrawal {
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
`
)

// GetFees The fees that were paid by the wallet sending the transaction to commit it to the Bitcoin blockchain.
func (obj Withdrawal) GetFees() *CurrencyAmount {
	return obj.Fees
}

// GetBlockHash The hash of the block that included this transaction. This will be null for unconfirmed transactions.
func (obj Withdrawal) GetBlockHash() *string {
	return obj.BlockHash
}

// GetBlockHeight The height of the block that included this transaction. This will be zero for unconfirmed transactions.
func (obj Withdrawal) GetBlockHeight() int64 {
	return obj.BlockHeight
}

// GetDestinationAddresses The Bitcoin blockchain addresses this transaction was sent to.
func (obj Withdrawal) GetDestinationAddresses() []string {
	return obj.DestinationAddresses
}

// GetNumConfirmations The number of blockchain confirmations for this transaction in real time.
func (obj Withdrawal) GetNumConfirmations() *int64 {
	return obj.NumConfirmations
}

// GetStatus The current status of this transaction.
func (obj Withdrawal) GetStatus() TransactionStatus {
	return obj.Status
}

// GetResolvedAt The date and time when this transaction was completed or failed.
func (obj Withdrawal) GetResolvedAt() *time.Time {
	return obj.ResolvedAt
}

// GetAmount The amount of money involved in this transaction.
func (obj Withdrawal) GetAmount() CurrencyAmount {
	return obj.Amount
}

// GetTransactionHash The hash of this transaction, so it can be uniquely identified on the Lightning Network.
func (obj Withdrawal) GetTransactionHash() *string {
	return obj.TransactionHash
}

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj Withdrawal) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj Withdrawal) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj Withdrawal) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}
