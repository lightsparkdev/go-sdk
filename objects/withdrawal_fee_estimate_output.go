
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type WithdrawalFeeEstimateOutput struct {

    // FeeEstimate The estimated fee for the withdrawal.
    FeeEstimate CurrencyAmount `json:"withdrawal_fee_estimate_output_fee_estimate"`

}

const (
    WithdrawalFeeEstimateOutputFragment = `
fragment WithdrawalFeeEstimateOutputFragment on WithdrawalFeeEstimateOutput {
    __typename
    withdrawal_fee_estimate_output_fee_estimate: fee_estimate {
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







