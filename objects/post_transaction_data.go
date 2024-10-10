// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// PostTransactionData This object represents post-transaction data that could be used to register payment for KYT.
type PostTransactionData struct {

	// Utxo The utxo of the channel over which the payment went through in the format of <transaction_hash>:<output_index>.
	Utxo string `json:"post_transaction_data_utxo"`

	// Amount The amount of funds transferred in the payment.
	Amount CurrencyAmount `json:"post_transaction_data_amount"`
}

const (
	PostTransactionDataFragment = `
fragment PostTransactionDataFragment on PostTransactionData {
    __typename
    post_transaction_data_utxo: utxo
    post_transaction_data_amount: amount {
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
