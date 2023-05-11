// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// An object that represents the address of a node on the Lightning Network.
type NodeAddress struct {

	// The string representation of the address.
	Address string `json:"node_address_address"`

	// The type, or protocol, of this address.
	Typex NodeAddressType `json:"node_address_type"`
}

const (
	NodeAddressFragment = `
fragment NodeAddressFragment on NodeAddress {
    __typename
    node_address_address: address
    node_address_type: type
}
`
)
