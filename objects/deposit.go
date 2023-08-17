// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"time"

	"github.com/lightsparkdev/go-sdk/types"
)

// This object represents a Deposit made to a Lightspark node wallet. This operation occurs for any L1 funding transaction to the wallet. You can retrieve this object to receive detailed information about the deposit.
type Deposit struct {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"deposit_id"`

	// The date and time when this transaction was initiated.
	CreatedAt time.Time `json:"deposit_created_at"`

	// The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"deposit_updated_at"`

	// The current status of this transaction.
	Status TransactionStatus `json:"deposit_status"`

	// The date and time when this transaction was completed or failed.
	ResolvedAt *time.Time `json:"deposit_resolved_at"`

	// The amount of money involved in this transaction.
	Amount CurrencyAmount `json:"deposit_amount"`

	// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
	TransactionHash *string `json:"deposit_transaction_hash"`

	// The fees that were paid by the wallet sending the transaction to commit it to the Bitcoin blockchain.
	Fees *CurrencyAmount `json:"deposit_fees"`

	// The hash of the block that included this transaction. This will be null for unconfirmed transactions.
	BlockHash *string `json:"deposit_block_hash"`

	// The height of the block that included this transaction. This will be zero for unconfirmed transactions.
	BlockHeight int64 `json:"deposit_block_height"`

	// The Bitcoin blockchain addresses this transaction was sent to.
	DestinationAddresses []string `json:"deposit_destination_addresses"`

	// The number of blockchain confirmations for this transaction in real time.
	NumConfirmations *int64 `json:"deposit_num_confirmations"`

	// The recipient Lightspark node this deposit was sent to.
	Destination types.EntityWrapper `json:"deposit_destination"`
}

const (
	DepositFragment = `
fragment DepositFragment on Deposit {
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
`
)

// The fees that were paid by the wallet sending the transaction to commit it to the Bitcoin blockchain.
func (obj Deposit) GetFees() *CurrencyAmount {
	return obj.Fees
}

// The hash of the block that included this transaction. This will be null for unconfirmed transactions.
func (obj Deposit) GetBlockHash() *string {
	return obj.BlockHash
}

// The height of the block that included this transaction. This will be zero for unconfirmed transactions.
func (obj Deposit) GetBlockHeight() int64 {
	return obj.BlockHeight
}

// The Bitcoin blockchain addresses this transaction was sent to.
func (obj Deposit) GetDestinationAddresses() []string {
	return obj.DestinationAddresses
}

// The number of blockchain confirmations for this transaction in real time.
func (obj Deposit) GetNumConfirmations() *int64 {
	return obj.NumConfirmations
}

// The current status of this transaction.
func (obj Deposit) GetStatus() TransactionStatus {
	return obj.Status
}

// The date and time when this transaction was completed or failed.
func (obj Deposit) GetResolvedAt() *time.Time {
	return obj.ResolvedAt
}

// The amount of money involved in this transaction.
func (obj Deposit) GetAmount() CurrencyAmount {
	return obj.Amount
}

// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
func (obj Deposit) GetTransactionHash() *string {
	return obj.TransactionHash
}

// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj Deposit) GetId() string {
	return obj.Id
}

// The date and time when the entity was first created.
func (obj Deposit) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// The date and time when the entity was last updated.
func (obj Deposit) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}
