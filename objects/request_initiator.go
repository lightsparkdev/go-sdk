// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"strings"
)

type RequestInitiator int

const (
	RequestInitiatorUndefined RequestInitiator = iota

	RequestInitiatorCustomer

	RequestInitiatorLightspark
)

func (a *RequestInitiator) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = RequestInitiatorUndefined
	case "CUSTOMER":
		*a = RequestInitiatorCustomer
	case "LIGHTSPARK":
		*a = RequestInitiatorLightspark

	}
	return nil
}

func (a RequestInitiator) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case RequestInitiatorCustomer:
		s = "CUSTOMER"
	case RequestInitiatorLightspark:
		s = "LIGHTSPARK"

	}
	return s
}

func (a RequestInitiator) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
