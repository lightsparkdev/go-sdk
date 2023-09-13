// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

// OutgoingPaymentAttemptStatus This is an enum of all potential statuses of a payment attempt made from a Lightspark Node.
type OutgoingPaymentAttemptStatus int

const (
	OutgoingPaymentAttemptStatusUndefined OutgoingPaymentAttemptStatus = iota

	OutgoingPaymentAttemptStatusInFlight

	OutgoingPaymentAttemptStatusSucceeded

	OutgoingPaymentAttemptStatusFailed
)

func (a *OutgoingPaymentAttemptStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = OutgoingPaymentAttemptStatusUndefined
	case "IN_FLIGHT":
		*a = OutgoingPaymentAttemptStatusInFlight
	case "SUCCEEDED":
		*a = OutgoingPaymentAttemptStatusSucceeded
	case "FAILED":
		*a = OutgoingPaymentAttemptStatusFailed

	}
	return nil
}

func (a OutgoingPaymentAttemptStatus) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case OutgoingPaymentAttemptStatusInFlight:
		s = "IN_FLIGHT"
	case OutgoingPaymentAttemptStatusSucceeded:
		s = "SUCCEEDED"
	case OutgoingPaymentAttemptStatusFailed:
		s = "FAILED"

	}
	return s
}

func (a OutgoingPaymentAttemptStatus) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
