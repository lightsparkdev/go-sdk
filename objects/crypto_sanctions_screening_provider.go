// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

// CryptoSanctionsScreeningProvider This is an enum identifying a type of crypto sanctions screening provider.
type CryptoSanctionsScreeningProvider int

const (
	CryptoSanctionsScreeningProviderUndefined CryptoSanctionsScreeningProvider = iota

	CryptoSanctionsScreeningProviderChainalysis
)

func (a *CryptoSanctionsScreeningProvider) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = CryptoSanctionsScreeningProviderUndefined
	case "CHAINALYSIS":
		*a = CryptoSanctionsScreeningProviderChainalysis

	}
	return nil
}

func (a CryptoSanctionsScreeningProvider) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case CryptoSanctionsScreeningProviderChainalysis:
		s = "CHAINALYSIS"

	}
	return s
}

func (a CryptoSanctionsScreeningProvider) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
