// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// CurrencyAmount This object represents the value and unit for an amount of currency.
type CurrencyAmount struct {

	// OriginalValue The original numeric value for this CurrencyAmount.
	OriginalValue int64 `json:"currency_amount_original_value"`

	// OriginalUnit The original unit of currency for this CurrencyAmount.
	OriginalUnit CurrencyUnit `json:"currency_amount_original_unit"`

	// PreferredCurrencyUnit The unit of user's preferred currency.
	PreferredCurrencyUnit CurrencyUnit `json:"currency_amount_preferred_currency_unit"`

	// PreferredCurrencyValueRounded The rounded numeric value for this CurrencyAmount in the very base level of user's preferred currency. For example, for USD, the value will be in cents.
	PreferredCurrencyValueRounded int64 `json:"currency_amount_preferred_currency_value_rounded"`

	// PreferredCurrencyValueApprox The approximate float value for this CurrencyAmount in the very base level of user's preferred currency. For example, for USD, the value will be in cents.
	PreferredCurrencyValueApprox float64 `json:"currency_amount_preferred_currency_value_approx"`
}

const (
	CurrencyAmountFragment = `
fragment CurrencyAmountFragment on CurrencyAmount {
    __typename
    currency_amount_original_value: original_value
    currency_amount_original_unit: original_unit
    currency_amount_preferred_currency_unit: preferred_currency_unit
    currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
    currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
}
`
)
