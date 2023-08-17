// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// This is an object representing a detailed breakdown of the balance for a Lightspark Node.
type BlockchainBalance struct {

	// The total wallet balance, including unconfirmed UTXOs.
	TotalBalance *CurrencyAmount `json:"blockchain_balance_total_balance"`

	// The balance of confirmed UTXOs in the wallet.
	ConfirmedBalance *CurrencyAmount `json:"blockchain_balance_confirmed_balance"`

	// The balance of unconfirmed UTXOs in the wallet.
	UnconfirmedBalance *CurrencyAmount `json:"blockchain_balance_unconfirmed_balance"`

	// The balance that's locked by an on-chain transaction.
	LockedBalance *CurrencyAmount `json:"blockchain_balance_locked_balance"`

	// Funds required to be held in reserve for channel bumping.
	RequiredReserve *CurrencyAmount `json:"blockchain_balance_required_reserve"`

	// Funds available for creating channels or withdrawing.
	AvailableBalance *CurrencyAmount `json:"blockchain_balance_available_balance"`
}

const (
	BlockchainBalanceFragment = `
fragment BlockchainBalanceFragment on BlockchainBalance {
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
`
)
