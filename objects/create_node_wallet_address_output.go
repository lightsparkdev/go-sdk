// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "lightspark/types"

type CreateNodeWalletAddressOutput struct {
	Node types.EntityWrapper `json:"create_node_wallet_address_output_node"`

	WalletAddress string `json:"create_node_wallet_address_output_wallet_address"`
}

const (
	CreateNodeWalletAddressOutputFragment = `
fragment CreateNodeWalletAddressOutputFragment on CreateNodeWalletAddressOutput {
    __typename
    create_node_wallet_address_output_node: node {
        id
    }
    create_node_wallet_address_output_wallet_address: wallet_address
}
`
)
