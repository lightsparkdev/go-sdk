// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

// WithdrawalMode This is an enum of the potential modes that your Bitcoin withdrawal can take.
type WithdrawalMode int

const (
	WithdrawalModeUndefined WithdrawalMode = iota

	WithdrawalModeWalletOnly

	WithdrawalModeWalletThenChannels
)

func (a *WithdrawalMode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = WithdrawalModeUndefined
	case "WALLET_ONLY":
		*a = WithdrawalModeWalletOnly
	case "WALLET_THEN_CHANNELS":
		*a = WithdrawalModeWalletThenChannels

	}
	return nil
}

func (a WithdrawalMode) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case WithdrawalModeWalletOnly:
		s = "WALLET_ONLY"
	case WithdrawalModeWalletThenChannels:
		s = "WALLET_THEN_CHANNELS"

	}
	return s
}

func (a WithdrawalMode) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
