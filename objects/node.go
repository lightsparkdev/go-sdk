// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
)

// Node This object is an interface representing a Lightning Node on the Lightning Network, and could either be a Lightspark node or a node managed by a third party.
type Node interface {
	Entity

	// GetAlias A name that identifies the node. It has no importance in terms of operating the node, it is just a way to identify and search for commercial services or popular nodes. This alias can be changed at any time by the node operator.
	GetAlias() *string

	// GetBitcoinNetwork The Bitcoin Network this node is deployed in.
	GetBitcoinNetwork() BitcoinNetwork

	// GetColor A hexadecimal string that describes a color. For example "#000000" is black, "#FFFFFF" is white. It has no importance in terms of operating the node, it is just a way to visually differentiate nodes. That color can be changed at any time by the node operator.
	GetColor() *string

	// GetConductivity A summary metric used to capture how well positioned a node is to send, receive, or route transactions efficiently. Maximizing a node's conductivity helps a node’s transactions to be capital efficient. The value is an integer ranging between 0 and 10 (bounds included).
	// Deprecated: Not supported.
	GetConductivity() *int64

	// GetDisplayName The name of this node in the network. It will be the most human-readable option possible, depending on the data available for this node.
	GetDisplayName() string

	// GetPublicKey The public key of this node. It acts as a unique identifier of this node in the Lightning Network.
	GetPublicKey() *string
}

func NodeUnmarshal(data map[string]interface{}) (Node, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "GraphNode":
		var graphNode GraphNode
		if err := json.Unmarshal(dataJSON, &graphNode); err != nil {
			return nil, err
		}
		return graphNode, nil
	case "LightsparkNodeWithOSK":
		var lightsparkNodeWithOSK LightsparkNodeWithOSK
		if err := json.Unmarshal(dataJSON, &lightsparkNodeWithOSK); err != nil {
			return nil, err
		}
		return lightsparkNodeWithOSK, nil
	case "LightsparkNodeWithRemoteSigning":
		var lightsparkNodeWithRemoteSigning LightsparkNodeWithRemoteSigning
		if err := json.Unmarshal(dataJSON, &lightsparkNodeWithRemoteSigning); err != nil {
			return nil, err
		}
		return lightsparkNodeWithRemoteSigning, nil

	default:
		return nil, fmt.Errorf("unknown Node type: %s", data["__typename"])
	}
}
