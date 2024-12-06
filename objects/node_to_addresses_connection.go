
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects



// NodeToAddressesConnection A connection between a node and the addresses it has announced for itself on Lightning Network.
type NodeToAddressesConnection struct {

    // Count The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
    Count int64 `json:"node_to_addresses_connection_count"`

    // Entities The addresses for the current page of this connection.
    Entities []NodeAddress `json:"node_to_addresses_connection_entities"`

}

const (
    NodeToAddressesConnectionFragment = `
fragment NodeToAddressesConnectionFragment on NodeToAddressesConnection {
    __typename
    node_to_addresses_connection_count: count
    node_to_addresses_connection_entities: entities {
        __typename
        node_address_address: address
        node_address_type: type
    }
}
`
)







