// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"strings"
)

// LightningPaymentDirection This is an enum identifying the payment direction.
type LightningPaymentDirection int

const (
	LightningPaymentDirectionUndefined LightningPaymentDirection = iota

	// LightningPaymentDirectionIncoming A payment that is received by the node.
	LightningPaymentDirectionIncoming
	// LightningPaymentDirectionOutgoing A payment that is sent by the node.
	LightningPaymentDirectionOutgoing
)

func (a *LightningPaymentDirection) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = LightningPaymentDirectionUndefined
	case "INCOMING":
		*a = LightningPaymentDirectionIncoming
	case "OUTGOING":
		*a = LightningPaymentDirectionOutgoing

	}
	return nil
}

func (a LightningPaymentDirection) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case LightningPaymentDirectionIncoming:
		s = "INCOMING"
	case LightningPaymentDirectionOutgoing:
		s = "OUTGOING"

	}
	return s
}

func (a LightningPaymentDirection) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
