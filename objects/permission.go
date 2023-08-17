// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

// Permission This is an enum of the potential permissions that a Lightspark user can have in regards to account management.
type Permission int

const (
	PermissionUndefined Permission = iota

	PermissionAll

	PermissionMainnetView

	PermissionMainnetTransact

	PermissionMainnetManage

	PermissionTestnetView

	PermissionTestnetTransact

	PermissionTestnetManage

	PermissionRegtestView

	PermissionRegtestTransact

	PermissionRegtestManage

	PermissionUserView

	PermissionUserManage

	PermissionAccountView

	PermissionAccountManage
)

func (a *Permission) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = PermissionUndefined
	case "ALL":
		*a = PermissionAll
	case "MAINNET_VIEW":
		*a = PermissionMainnetView
	case "MAINNET_TRANSACT":
		*a = PermissionMainnetTransact
	case "MAINNET_MANAGE":
		*a = PermissionMainnetManage
	case "TESTNET_VIEW":
		*a = PermissionTestnetView
	case "TESTNET_TRANSACT":
		*a = PermissionTestnetTransact
	case "TESTNET_MANAGE":
		*a = PermissionTestnetManage
	case "REGTEST_VIEW":
		*a = PermissionRegtestView
	case "REGTEST_TRANSACT":
		*a = PermissionRegtestTransact
	case "REGTEST_MANAGE":
		*a = PermissionRegtestManage
	case "USER_VIEW":
		*a = PermissionUserView
	case "USER_MANAGE":
		*a = PermissionUserManage
	case "ACCOUNT_VIEW":
		*a = PermissionAccountView
	case "ACCOUNT_MANAGE":
		*a = PermissionAccountManage

	}
	return nil
}

func (a Permission) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case PermissionAll:
		s = "ALL"
	case PermissionMainnetView:
		s = "MAINNET_VIEW"
	case PermissionMainnetTransact:
		s = "MAINNET_TRANSACT"
	case PermissionMainnetManage:
		s = "MAINNET_MANAGE"
	case PermissionTestnetView:
		s = "TESTNET_VIEW"
	case PermissionTestnetTransact:
		s = "TESTNET_TRANSACT"
	case PermissionTestnetManage:
		s = "TESTNET_MANAGE"
	case PermissionRegtestView:
		s = "REGTEST_VIEW"
	case PermissionRegtestTransact:
		s = "REGTEST_TRANSACT"
	case PermissionRegtestManage:
		s = "REGTEST_MANAGE"
	case PermissionUserView:
		s = "USER_VIEW"
	case PermissionUserManage:
		s = "USER_MANAGE"
	case PermissionAccountView:
		s = "ACCOUNT_VIEW"
	case PermissionAccountManage:
		s = "ACCOUNT_MANAGE"

	}
	return s
}

func (a Permission) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
