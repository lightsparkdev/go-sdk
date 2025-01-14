// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"strings"
)

// PaymentRequestStatus This is an enum of the potential states that a payment request on the Lightning Network can take.
type PaymentRequestStatus int

const (
	PaymentRequestStatusUndefined PaymentRequestStatus = iota

	PaymentRequestStatusOpen

	PaymentRequestStatusClosed
)

func (a *PaymentRequestStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = PaymentRequestStatusUndefined
	case "OPEN":
		*a = PaymentRequestStatusOpen
	case "CLOSED":
		*a = PaymentRequestStatusClosed

	}
	return nil
}

func (a PaymentRequestStatus) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case PaymentRequestStatusOpen:
		s = "OPEN"
	case PaymentRequestStatusClosed:
		s = "CLOSED"

	}
	return s
}

func (a PaymentRequestStatus) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
