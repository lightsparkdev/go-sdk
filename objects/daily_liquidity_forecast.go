// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type DailyLiquidityForecast struct {

	// Date The date for which this forecast was generated.
	Date types.Date `json:"daily_liquidity_forecast_date"`

	// Direction The direction for which this forecast was generated.
	Direction LightningPaymentDirection `json:"daily_liquidity_forecast_direction"`

	// Amount The value of the forecast. It represents the amount of msats that we think will be moved for that specified direction, for that node, on that date.
	Amount CurrencyAmount `json:"daily_liquidity_forecast_amount"`
}

const (
	DailyLiquidityForecastFragment = `
fragment DailyLiquidityForecastFragment on DailyLiquidityForecast {
    __typename
    daily_liquidity_forecast_date: date
    daily_liquidity_forecast_direction: direction
    daily_liquidity_forecast_amount: amount {
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
