// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

type WalletStatus int

const (
	WalletStatusUndefined WalletStatus = iota

	// The wallet has not been set up yet and is ready to be deployed. This is the default status after the first login.
	WalletStatusNotSetup
	// The wallet is currently being deployed in the Lightspark infrastructure.
	WalletStatusDeploying
	// The wallet has been deployed in the Lightspark infrastructure and is ready to be initialized.
	WalletStatusDeployed
	// The wallet is currently being initialized.
	WalletStatusInitializing
	// The wallet is available and ready to be used.
	WalletStatusReady
	// The wallet is temporarily available, due to a transient issue or a scheduled maintenance.
	WalletStatusUnavailable
	// The wallet had an unrecoverable failure. This status is not expected to happend and will be investigated by the Lightspark team.
	WalletStatusFailed
	// The wallet is being terminated.
	WalletStatusTerminating
	// The wallet has been terminated and is not available in the Lightspark infrastructure anymore. It is not connected to the Lightning network and its funds can only be accessed using the Funds Recovery flow.
	WalletStatusTerminated
)

func (a *WalletStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = WalletStatusUndefined
	case "NOT_SETUP":
		*a = WalletStatusNotSetup
	case "DEPLOYING":
		*a = WalletStatusDeploying
	case "DEPLOYED":
		*a = WalletStatusDeployed
	case "INITIALIZING":
		*a = WalletStatusInitializing
	case "READY":
		*a = WalletStatusReady
	case "UNAVAILABLE":
		*a = WalletStatusUnavailable
	case "FAILED":
		*a = WalletStatusFailed
	case "TERMINATING":
		*a = WalletStatusTerminating
	case "TERMINATED":
		*a = WalletStatusTerminated

	}
	return nil
}

func (a WalletStatus) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case WalletStatusNotSetup:
		s = "NOT_SETUP"
	case WalletStatusDeploying:
		s = "DEPLOYING"
	case WalletStatusDeployed:
		s = "DEPLOYED"
	case WalletStatusInitializing:
		s = "INITIALIZING"
	case WalletStatusReady:
		s = "READY"
	case WalletStatusUnavailable:
		s = "UNAVAILABLE"
	case WalletStatusFailed:
		s = "FAILED"
	case WalletStatusTerminating:
		s = "TERMINATING"
	case WalletStatusTerminated:
		s = "TERMINATED"

	}
	return s
}

func (a WalletStatus) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
