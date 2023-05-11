// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

type RoutingTransactionFailureReason int

const (
	RoutingTransactionFailureReasonUndefined RoutingTransactionFailureReason = iota

	RoutingTransactionFailureReasonIncomingLinkFailure

	RoutingTransactionFailureReasonOutgoingLinkFailure

	RoutingTransactionFailureReasonForwardingFailure
)

func (a *RoutingTransactionFailureReason) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = RoutingTransactionFailureReasonUndefined
	case "INCOMING_LINK_FAILURE":
		*a = RoutingTransactionFailureReasonIncomingLinkFailure
	case "OUTGOING_LINK_FAILURE":
		*a = RoutingTransactionFailureReasonOutgoingLinkFailure
	case "FORWARDING_FAILURE":
		*a = RoutingTransactionFailureReasonForwardingFailure

	}
	return nil
}

func (a RoutingTransactionFailureReason) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case RoutingTransactionFailureReasonIncomingLinkFailure:
		s = "INCOMING_LINK_FAILURE"
	case RoutingTransactionFailureReasonOutgoingLinkFailure:
		s = "OUTGOING_LINK_FAILURE"
	case RoutingTransactionFailureReasonForwardingFailure:
		s = "FORWARDING_FAILURE"

	}
	return s
}

func (a RoutingTransactionFailureReason) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
