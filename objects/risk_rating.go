// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

// RiskRating This is an enum of the potential risk ratings related to a transaction made over the Lightning Network. These risk ratings are returned from the CryptoSanctionScreeningProvider.
type RiskRating int

const (
	RiskRatingUndefined RiskRating = iota

	RiskRatingHighRisk

	RiskRatingLowRisk

	RiskRatingUnknown
)

func (a *RiskRating) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = RiskRatingUndefined
	case "HIGH_RISK":
		*a = RiskRatingHighRisk
	case "LOW_RISK":
		*a = RiskRatingLowRisk
	case "UNKNOWN":
		*a = RiskRatingUnknown

	}
	return nil
}

func (a RiskRating) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case RiskRatingHighRisk:
		s = "HIGH_RISK"
	case RiskRatingLowRisk:
		s = "LOW_RISK"
	case RiskRatingUnknown:
		s = "UNKNOWN"

	}
	return s
}

func (a RiskRating) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
