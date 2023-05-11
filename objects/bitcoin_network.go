// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

type BitcoinNetwork int

const (
	BitcoinNetworkUndefined BitcoinNetwork = iota

	// The production version of the Bitcoin Blockchain.
	BitcoinNetworkMainnet
	// A test version of the Bitcoin Blockchain, maintained by Lightspark.
	BitcoinNetworkRegtest
	// A test version of the Bitcoin Blockchain, maintained by a centralized organization. Not in use at Lightspark.
	BitcoinNetworkSignet
	// A test version of the Bitcoin Blockchain, publically available.
	BitcoinNetworkTestnet
)

func (a *BitcoinNetwork) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = BitcoinNetworkUndefined
	case "MAINNET":
		*a = BitcoinNetworkMainnet
	case "REGTEST":
		*a = BitcoinNetworkRegtest
	case "SIGNET":
		*a = BitcoinNetworkSignet
	case "TESTNET":
		*a = BitcoinNetworkTestnet

	}
	return nil
}

func (a BitcoinNetwork) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case BitcoinNetworkMainnet:
		s = "MAINNET"
	case BitcoinNetworkRegtest:
		s = "REGTEST"
	case BitcoinNetworkSignet:
		s = "SIGNET"
	case BitcoinNetworkTestnet:
		s = "TESTNET"

	}
	return s
}

func (a BitcoinNetwork) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
