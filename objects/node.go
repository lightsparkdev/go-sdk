// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
)

type Node interface {
	Entity

	// A name that identifies the node. It has no importance in terms of operating the node, it is just a way to identify and search for commercial services or popular nodes. This alias can be changed at any time by the node operator.
	GetAlias() *string

	// The Bitcoin Network this node is deployed in.
	GetBitcoinNetwork() BitcoinNetwork

	// A hexadecimal string that describes a color. For example "#000000" is black, "#FFFFFF" is white. It has no importance in terms of operating the node, it is just a way to visually differentiate nodes. That color can be changed at any time by the node operator.
	GetColor() *string

	// A summary metric used to capture how well positioned a node is to send, receive, or route transactions efficiently. Maximizing a node's conductivity helps a node’s transactions to be capital efficient. The value is an integer ranging between 0 and 10 (bounds included).
	GetConductivity() *int64

	// The name of this node in the network. It will be the most human-readable option possible, depending on the data available for this node.
	GetDisplayName() string

	// The public key of this node. It acts as a unique identifier of this node in the Lightning Network.
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
	case "LightsparkNode":
		var lightsparkNode LightsparkNode
		if err := json.Unmarshal(dataJSON, &lightsparkNode); err != nil {
			return nil, err
		}
		return lightsparkNode, nil

	default:
		return nil, fmt.Errorf("unknown Node type: %s", data["__typename"])
	}
}
