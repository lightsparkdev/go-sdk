// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type ChannelToTransactionsConnection struct {

	// The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"channel_to_transactions_connection_count"`

	// The average fee for the transactions that transited through this channel, according to the filters and constraints of the connection.
	AverageFee *CurrencyAmount `json:"channel_to_transactions_connection_average_fee"`

	// The total amount transacted for the transactions that transited through this channel, according to the filters and constraints of the connection.
	TotalAmountTransacted *CurrencyAmount `json:"channel_to_transactions_connection_total_amount_transacted"`

	// The total amount of fees for the transactions that transited through this channel, according to the filters and constraints of the connection.
	TotalFees *CurrencyAmount `json:"channel_to_transactions_connection_total_fees"`
}

const (
	ChannelToTransactionsConnectionFragment = `
fragment ChannelToTransactionsConnectionFragment on ChannelToTransactionsConnection {
    __typename
    channel_to_transactions_connection_count: count
    channel_to_transactions_connection_average_fee: average_fee {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    channel_to_transactions_connection_total_amount_transacted: total_amount_transacted {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    channel_to_transactions_connection_total_fees: total_fees {
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
