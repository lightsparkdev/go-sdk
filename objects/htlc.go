// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// Htlc This object represents an HTLC of a SUCCEEDED payment that could be used to register the payment for KYT.
type Htlc struct {

	// Utxo The utxo of the channel over which the htlc went through in the format of <transaction_hash>:<output_index>.
	Utxo string `json:"htlc_utxo"`

	// Amount The amount of funds transferred in the htlc.
	Amount CurrencyAmount `json:"htlc_amount"`
}

const (
	HtlcFragment = `
fragment HtlcFragment on Htlc {
    __typename
    htlc_utxo: utxo
    htlc_amount: amount {
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
