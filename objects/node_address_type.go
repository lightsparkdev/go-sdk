// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

// An enum that enumerates all possible types of addresses of a node on the Lightning Network.
type NodeAddressType int

const (
	NodeAddressTypeUndefined NodeAddressType = iota

	NodeAddressTypeIpv4

	NodeAddressTypeIpv6

	NodeAddressTypeTor
)

func (a *NodeAddressType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = NodeAddressTypeUndefined
	case "IPV4":
		*a = NodeAddressTypeIpv4
	case "IPV6":
		*a = NodeAddressTypeIpv6
	case "TOR":
		*a = NodeAddressTypeTor

	}
	return nil
}

func (a NodeAddressType) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case NodeAddressTypeIpv4:
		s = "IPV4"
	case NodeAddressTypeIpv6:
		s = "IPV6"
	case NodeAddressTypeTor:
		s = "TOR"

	}
	return s
}

func (a NodeAddressType) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
