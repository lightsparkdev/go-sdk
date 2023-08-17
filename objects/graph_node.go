// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"time"

	"github.com/lightsparkdev/go-sdk/requester"
)

// This object represents a node that exists on the Lightning Network, including nodes not managed by Lightspark. You can retrieve this object to get publicly available information about any node on the Lightning Network.
type GraphNode struct {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"graph_node_id"`

	// The date and time when the entity was first created.
	CreatedAt time.Time `json:"graph_node_created_at"`

	// The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"graph_node_updated_at"`

	// A name that identifies the node. It has no importance in terms of operating the node, it is just a way to identify and search for commercial services or popular nodes. This alias can be changed at any time by the node operator.
	Alias *string `json:"graph_node_alias"`

	// The Bitcoin Network this node is deployed in.
	BitcoinNetwork BitcoinNetwork `json:"graph_node_bitcoin_network"`

	// A hexadecimal string that describes a color. For example "#000000" is black, "#FFFFFF" is white. It has no importance in terms of operating the node, it is just a way to visually differentiate nodes. That color can be changed at any time by the node operator.
	Color *string `json:"graph_node_color"`

	// A summary metric used to capture how well positioned a node is to send, receive, or route transactions efficiently. Maximizing a node's conductivity helps a node’s transactions to be capital efficient. The value is an integer ranging between 0 and 10 (bounds included).
	Conductivity *int64 `json:"graph_node_conductivity"`

	// The name of this node in the network. It will be the most human-readable option possible, depending on the data available for this node.
	DisplayName string `json:"graph_node_display_name"`

	// The public key of this node. It acts as a unique identifier of this node in the Lightning Network.
	PublicKey *string `json:"graph_node_public_key"`
}

const (
	GraphNodeFragment = `
fragment GraphNodeFragment on GraphNode {
    __typename
    graph_node_id: id
    graph_node_created_at: created_at
    graph_node_updated_at: updated_at
    graph_node_alias: alias
    graph_node_bitcoin_network: bitcoin_network
    graph_node_color: color
    graph_node_conductivity: conductivity
    graph_node_display_name: display_name
    graph_node_public_key: public_key
}
`
)

// A name that identifies the node. It has no importance in terms of operating the node, it is just a way to identify and search for commercial services or popular nodes. This alias can be changed at any time by the node operator.
func (obj GraphNode) GetAlias() *string {
	return obj.Alias
}

// The Bitcoin Network this node is deployed in.
func (obj GraphNode) GetBitcoinNetwork() BitcoinNetwork {
	return obj.BitcoinNetwork
}

// A hexadecimal string that describes a color. For example "#000000" is black, "#FFFFFF" is white. It has no importance in terms of operating the node, it is just a way to visually differentiate nodes. That color can be changed at any time by the node operator.
func (obj GraphNode) GetColor() *string {
	return obj.Color
}

// A summary metric used to capture how well positioned a node is to send, receive, or route transactions efficiently. Maximizing a node's conductivity helps a node’s transactions to be capital efficient. The value is an integer ranging between 0 and 10 (bounds included).
func (obj GraphNode) GetConductivity() *int64 {
	return obj.Conductivity
}

// The name of this node in the network. It will be the most human-readable option possible, depending on the data available for this node.
func (obj GraphNode) GetDisplayName() string {
	return obj.DisplayName
}

// The public key of this node. It acts as a unique identifier of this node in the Lightning Network.
func (obj GraphNode) GetPublicKey() *string {
	return obj.PublicKey
}

// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj GraphNode) GetId() string {
	return obj.Id
}

// The date and time when the entity was first created.
func (obj GraphNode) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// The date and time when the entity was last updated.
func (obj GraphNode) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj GraphNode) GetAddresses(requester *requester.Requester, first *int64, types *[]NodeAddressType) (*NodeToAddressesConnection, error) {
	query := `query FetchNodeToAddressesConnection($entity_id: ID!, $first: Int, $types: [NodeAddressType!]) {
    entity(id: $entity_id) {
        ... on GraphNode {
            addresses(, first: $first, types: $types) {
                __typename
                node_to_addresses_connection_count: count
                node_to_addresses_connection_entities: entities {
                    __typename
                    node_address_address: address
                    node_address_type: type
                }
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id": obj.Id,
		"first":     first,
		"types":     types,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["addresses"].(map[string]interface{})
	var result *NodeToAddressesConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}
