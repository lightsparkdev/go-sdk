// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type MultiSigAddressValidationParameters struct {

	// CounterpartyFundingPubkey The counterparty funding public key used to create the 2-of-2 multisig for the address.
	CounterpartyFundingPubkey string `json:"multi_sig_address_validation_parameters_counterparty_funding_pubkey"`

	// FundingPubkeyDerivationPath The derivation path used to derive the funding public key for the 2-of-2 multisig address.
	FundingPubkeyDerivationPath string `json:"multi_sig_address_validation_parameters_funding_pubkey_derivation_path"`
}

const (
	MultiSigAddressValidationParametersFragment = `
fragment MultiSigAddressValidationParametersFragment on MultiSigAddressValidationParameters {
    __typename
    multi_sig_address_validation_parameters_counterparty_funding_pubkey: counterparty_funding_pubkey
    multi_sig_address_validation_parameters_funding_pubkey_derivation_path: funding_pubkey_derivation_path
}
`
)
