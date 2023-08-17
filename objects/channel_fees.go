// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// This represents the fee policies set for a channel on the Lightning Network.
type ChannelFees struct {
	BaseFee *CurrencyAmount `json:"channel_fees_base_fee"`

	FeeRatePerMil *int64 `json:"channel_fees_fee_rate_per_mil"`
}

const (
	ChannelFeesFragment = `
fragment ChannelFeesFragment on ChannelFees {
    __typename
    channel_fees_base_fee: base_fee {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    channel_fees_fee_rate_per_mil: fee_rate_per_mil
}
`
)
