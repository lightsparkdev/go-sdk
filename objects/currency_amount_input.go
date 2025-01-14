// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type CurrencyAmountInput struct {
	Value int64 `json:"currency_amount_input_value"`

	Unit CurrencyUnit `json:"currency_amount_input_unit"`
}
