// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"strings"
)

// InvoiceType This is an enum for potential invoice types.
type InvoiceType int

const (
	InvoiceTypeUndefined InvoiceType = iota

	// InvoiceTypeStandard A standard Bolt 11 invoice.
	InvoiceTypeStandard
	// InvoiceTypeAmp An AMP (Atomic Multi-path Payment) invoice.
	InvoiceTypeAmp
)

func (a *InvoiceType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = InvoiceTypeUndefined
	case "STANDARD":
		*a = InvoiceTypeStandard
	case "AMP":
		*a = InvoiceTypeAmp

	}
	return nil
}

func (a InvoiceType) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case InvoiceTypeStandard:
		s = "STANDARD"
	case InvoiceTypeAmp:
		s = "AMP"

	}
	return s
}

func (a InvoiceType) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
