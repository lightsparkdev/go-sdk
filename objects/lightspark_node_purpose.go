// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

type LightsparkNodePurpose int

const (
	LightsparkNodePurposeUndefined LightsparkNodePurpose = iota

	LightsparkNodePurposeSend

	LightsparkNodePurposeReceive

	LightsparkNodePurposeRouting
)

func (a *LightsparkNodePurpose) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = LightsparkNodePurposeUndefined
	case "SEND":
		*a = LightsparkNodePurposeSend
	case "RECEIVE":
		*a = LightsparkNodePurposeReceive
	case "ROUTING":
		*a = LightsparkNodePurposeRouting

	}
	return nil
}

func (a LightsparkNodePurpose) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case LightsparkNodePurposeSend:
		s = "SEND"
	case LightsparkNodePurposeReceive:
		s = "RECEIVE"
	case LightsparkNodePurposeRouting:
		s = "ROUTING"

	}
	return s
}

func (a LightsparkNodePurpose) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
