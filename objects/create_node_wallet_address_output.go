// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type CreateNodeWalletAddressOutput struct {
	Node types.EntityWrapper `json:"create_node_wallet_address_output_node"`

	WalletAddress string `json:"create_node_wallet_address_output_wallet_address"`

	// MultisigWalletAddressValidationParameters Vaildation parameters for the 2-of-2 multisig address. None if the address is not a 2-of-2 multisig address.
	MultisigWalletAddressValidationParameters *MultiSigAddressValidationParameters `json:"create_node_wallet_address_output_multisig_wallet_address_validation_parameters"`
}

const (
	CreateNodeWalletAddressOutputFragment = `
fragment CreateNodeWalletAddressOutputFragment on CreateNodeWalletAddressOutput {
    __typename
    create_node_wallet_address_output_node: node {
        id
    }
    create_node_wallet_address_output_wallet_address: wallet_address
    create_node_wallet_address_output_multisig_wallet_address_validation_parameters: multisig_wallet_address_validation_parameters {
        __typename
        multi_sig_address_validation_parameters_counterparty_funding_pubkey: counterparty_funding_pubkey
        multi_sig_address_validation_parameters_funding_pubkey_derivation_path: funding_pubkey_derivation_path
    }
}
`
)
