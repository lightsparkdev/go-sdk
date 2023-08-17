// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

// RemoteSigningSubEventType This is an enum of the potential sub-event types for Remote Signing webook events.
type RemoteSigningSubEventType int

const (
	RemoteSigningSubEventTypeUndefined RemoteSigningSubEventType = iota

	RemoteSigningSubEventTypeEcdh

	RemoteSigningSubEventTypeGetPerCommitmentPoint

	RemoteSigningSubEventTypeReleasePerCommitmentSecret

	RemoteSigningSubEventTypeSignInvoice

	RemoteSigningSubEventTypeDeriveKeyAndSign

	RemoteSigningSubEventTypeReleasePaymentPreimage
)

func (a *RemoteSigningSubEventType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = RemoteSigningSubEventTypeUndefined
	case "ECDH":
		*a = RemoteSigningSubEventTypeEcdh
	case "GET_PER_COMMITMENT_POINT":
		*a = RemoteSigningSubEventTypeGetPerCommitmentPoint
	case "RELEASE_PER_COMMITMENT_SECRET":
		*a = RemoteSigningSubEventTypeReleasePerCommitmentSecret
	case "SIGN_INVOICE":
		*a = RemoteSigningSubEventTypeSignInvoice
	case "DERIVE_KEY_AND_SIGN":
		*a = RemoteSigningSubEventTypeDeriveKeyAndSign
	case "RELEASE_PAYMENT_PREIMAGE":
		*a = RemoteSigningSubEventTypeReleasePaymentPreimage

	}
	return nil
}

func (a RemoteSigningSubEventType) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case RemoteSigningSubEventTypeEcdh:
		s = "ECDH"
	case RemoteSigningSubEventTypeGetPerCommitmentPoint:
		s = "GET_PER_COMMITMENT_POINT"
	case RemoteSigningSubEventTypeReleasePerCommitmentSecret:
		s = "RELEASE_PER_COMMITMENT_SECRET"
	case RemoteSigningSubEventTypeSignInvoice:
		s = "SIGN_INVOICE"
	case RemoteSigningSubEventTypeDeriveKeyAndSign:
		s = "DERIVE_KEY_AND_SIGN"
	case RemoteSigningSubEventTypeReleasePaymentPreimage:
		s = "RELEASE_PAYMENT_PREIMAGE"

	}
	return s
}

func (a RemoteSigningSubEventType) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
