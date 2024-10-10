// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

// ComplianceProvider This is an enum identifying a type of compliance provider.
type ComplianceProvider int

const (
	ComplianceProviderUndefined ComplianceProvider = iota

	ComplianceProviderChainalysis
)

func (a *ComplianceProvider) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = ComplianceProviderUndefined
	case "CHAINALYSIS":
		*a = ComplianceProviderChainalysis

	}
	return nil
}

func (a ComplianceProvider) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case ComplianceProviderChainalysis:
		s = "CHAINALYSIS"

	}
	return s
}

func (a ComplianceProvider) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
