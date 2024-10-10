// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type LightsparkNodeToDailyLiquidityForecastsConnection struct {
	FromDate types.Date `json:"lightspark_node_to_daily_liquidity_forecasts_connection_from_date"`

	ToDate types.Date `json:"lightspark_node_to_daily_liquidity_forecasts_connection_to_date"`

	Direction LightningPaymentDirection `json:"lightspark_node_to_daily_liquidity_forecasts_connection_direction"`

	// Entities The daily liquidity forecasts for the current page of this connection.
	Entities []DailyLiquidityForecast `json:"lightspark_node_to_daily_liquidity_forecasts_connection_entities"`
}

const (
	LightsparkNodeToDailyLiquidityForecastsConnectionFragment = `
fragment LightsparkNodeToDailyLiquidityForecastsConnectionFragment on LightsparkNodeToDailyLiquidityForecastsConnection {
    __typename
    lightspark_node_to_daily_liquidity_forecasts_connection_from_date: from_date
    lightspark_node_to_daily_liquidity_forecasts_connection_to_date: to_date
    lightspark_node_to_daily_liquidity_forecasts_connection_direction: direction
    lightspark_node_to_daily_liquidity_forecasts_connection_entities: entities {
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
}
`
)
