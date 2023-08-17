// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

type SignablePayloadStatus int

const (
	SignablePayloadStatusUndefined SignablePayloadStatus = iota

	SignablePayloadStatusCreated

	SignablePayloadStatusSigned
)

func (a *SignablePayloadStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = SignablePayloadStatusUndefined
	case "CREATED":
		*a = SignablePayloadStatusCreated
	case "SIGNED":
		*a = SignablePayloadStatusSigned

	}
	return nil
}

func (a SignablePayloadStatus) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case SignablePayloadStatusCreated:
		s = "CREATED"
	case SignablePayloadStatusSigned:
		s = "SIGNED"

	}
	return s
}

func (a SignablePayloadStatus) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
