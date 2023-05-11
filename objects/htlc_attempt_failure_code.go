// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

type HtlcAttemptFailureCode int

const (
	HtlcAttemptFailureCodeUndefined HtlcAttemptFailureCode = iota

	HtlcAttemptFailureCodeIncorrectOrUnknownPaymentDetails

	HtlcAttemptFailureCodeIncorrectPaymentAmount

	HtlcAttemptFailureCodeFinalIncorrectCltvExpiry

	HtlcAttemptFailureCodeFinalIncorrectHtlcAmount

	HtlcAttemptFailureCodeFinalExpiryTooSoon

	HtlcAttemptFailureCodeInvalidRealm

	HtlcAttemptFailureCodeExpiryTooSoon

	HtlcAttemptFailureCodeInvalidOnionVersion

	HtlcAttemptFailureCodeInvalidOnionHmac

	HtlcAttemptFailureCodeInvalidOnionKey

	HtlcAttemptFailureCodeAmountBelowMinimum

	HtlcAttemptFailureCodeFeeInsufficient

	HtlcAttemptFailureCodeIncorrectCltvExpiry

	HtlcAttemptFailureCodeChannelDisabled

	HtlcAttemptFailureCodeTemporaryChannelFailure

	HtlcAttemptFailureCodeRequiredNodeFeatureMissing

	HtlcAttemptFailureCodeRequiredChannelFeatureMissing

	HtlcAttemptFailureCodeUnknownNextPeer

	HtlcAttemptFailureCodeTemporaryNodeFailure

	HtlcAttemptFailureCodePermanentNodeFailure

	HtlcAttemptFailureCodePermanentChannelFailure

	HtlcAttemptFailureCodeExpiryTooFar

	HtlcAttemptFailureCodeMppTimeout

	HtlcAttemptFailureCodeInvalidOnionPayload

	HtlcAttemptFailureCodeInternalFailure

	HtlcAttemptFailureCodeUnknownFailure

	HtlcAttemptFailureCodeUnreadableFailure
)

func (a *HtlcAttemptFailureCode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = HtlcAttemptFailureCodeUndefined
	case "INCORRECT_OR_UNKNOWN_PAYMENT_DETAILS":
		*a = HtlcAttemptFailureCodeIncorrectOrUnknownPaymentDetails
	case "INCORRECT_PAYMENT_AMOUNT":
		*a = HtlcAttemptFailureCodeIncorrectPaymentAmount
	case "FINAL_INCORRECT_CLTV_EXPIRY":
		*a = HtlcAttemptFailureCodeFinalIncorrectCltvExpiry
	case "FINAL_INCORRECT_HTLC_AMOUNT":
		*a = HtlcAttemptFailureCodeFinalIncorrectHtlcAmount
	case "FINAL_EXPIRY_TOO_SOON":
		*a = HtlcAttemptFailureCodeFinalExpiryTooSoon
	case "INVALID_REALM":
		*a = HtlcAttemptFailureCodeInvalidRealm
	case "EXPIRY_TOO_SOON":
		*a = HtlcAttemptFailureCodeExpiryTooSoon
	case "INVALID_ONION_VERSION":
		*a = HtlcAttemptFailureCodeInvalidOnionVersion
	case "INVALID_ONION_HMAC":
		*a = HtlcAttemptFailureCodeInvalidOnionHmac
	case "INVALID_ONION_KEY":
		*a = HtlcAttemptFailureCodeInvalidOnionKey
	case "AMOUNT_BELOW_MINIMUM":
		*a = HtlcAttemptFailureCodeAmountBelowMinimum
	case "FEE_INSUFFICIENT":
		*a = HtlcAttemptFailureCodeFeeInsufficient
	case "INCORRECT_CLTV_EXPIRY":
		*a = HtlcAttemptFailureCodeIncorrectCltvExpiry
	case "CHANNEL_DISABLED":
		*a = HtlcAttemptFailureCodeChannelDisabled
	case "TEMPORARY_CHANNEL_FAILURE":
		*a = HtlcAttemptFailureCodeTemporaryChannelFailure
	case "REQUIRED_NODE_FEATURE_MISSING":
		*a = HtlcAttemptFailureCodeRequiredNodeFeatureMissing
	case "REQUIRED_CHANNEL_FEATURE_MISSING":
		*a = HtlcAttemptFailureCodeRequiredChannelFeatureMissing
	case "UNKNOWN_NEXT_PEER":
		*a = HtlcAttemptFailureCodeUnknownNextPeer
	case "TEMPORARY_NODE_FAILURE":
		*a = HtlcAttemptFailureCodeTemporaryNodeFailure
	case "PERMANENT_NODE_FAILURE":
		*a = HtlcAttemptFailureCodePermanentNodeFailure
	case "PERMANENT_CHANNEL_FAILURE":
		*a = HtlcAttemptFailureCodePermanentChannelFailure
	case "EXPIRY_TOO_FAR":
		*a = HtlcAttemptFailureCodeExpiryTooFar
	case "MPP_TIMEOUT":
		*a = HtlcAttemptFailureCodeMppTimeout
	case "INVALID_ONION_PAYLOAD":
		*a = HtlcAttemptFailureCodeInvalidOnionPayload
	case "INTERNAL_FAILURE":
		*a = HtlcAttemptFailureCodeInternalFailure
	case "UNKNOWN_FAILURE":
		*a = HtlcAttemptFailureCodeUnknownFailure
	case "UNREADABLE_FAILURE":
		*a = HtlcAttemptFailureCodeUnreadableFailure

	}
	return nil
}

func (a HtlcAttemptFailureCode) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case HtlcAttemptFailureCodeIncorrectOrUnknownPaymentDetails:
		s = "INCORRECT_OR_UNKNOWN_PAYMENT_DETAILS"
	case HtlcAttemptFailureCodeIncorrectPaymentAmount:
		s = "INCORRECT_PAYMENT_AMOUNT"
	case HtlcAttemptFailureCodeFinalIncorrectCltvExpiry:
		s = "FINAL_INCORRECT_CLTV_EXPIRY"
	case HtlcAttemptFailureCodeFinalIncorrectHtlcAmount:
		s = "FINAL_INCORRECT_HTLC_AMOUNT"
	case HtlcAttemptFailureCodeFinalExpiryTooSoon:
		s = "FINAL_EXPIRY_TOO_SOON"
	case HtlcAttemptFailureCodeInvalidRealm:
		s = "INVALID_REALM"
	case HtlcAttemptFailureCodeExpiryTooSoon:
		s = "EXPIRY_TOO_SOON"
	case HtlcAttemptFailureCodeInvalidOnionVersion:
		s = "INVALID_ONION_VERSION"
	case HtlcAttemptFailureCodeInvalidOnionHmac:
		s = "INVALID_ONION_HMAC"
	case HtlcAttemptFailureCodeInvalidOnionKey:
		s = "INVALID_ONION_KEY"
	case HtlcAttemptFailureCodeAmountBelowMinimum:
		s = "AMOUNT_BELOW_MINIMUM"
	case HtlcAttemptFailureCodeFeeInsufficient:
		s = "FEE_INSUFFICIENT"
	case HtlcAttemptFailureCodeIncorrectCltvExpiry:
		s = "INCORRECT_CLTV_EXPIRY"
	case HtlcAttemptFailureCodeChannelDisabled:
		s = "CHANNEL_DISABLED"
	case HtlcAttemptFailureCodeTemporaryChannelFailure:
		s = "TEMPORARY_CHANNEL_FAILURE"
	case HtlcAttemptFailureCodeRequiredNodeFeatureMissing:
		s = "REQUIRED_NODE_FEATURE_MISSING"
	case HtlcAttemptFailureCodeRequiredChannelFeatureMissing:
		s = "REQUIRED_CHANNEL_FEATURE_MISSING"
	case HtlcAttemptFailureCodeUnknownNextPeer:
		s = "UNKNOWN_NEXT_PEER"
	case HtlcAttemptFailureCodeTemporaryNodeFailure:
		s = "TEMPORARY_NODE_FAILURE"
	case HtlcAttemptFailureCodePermanentNodeFailure:
		s = "PERMANENT_NODE_FAILURE"
	case HtlcAttemptFailureCodePermanentChannelFailure:
		s = "PERMANENT_CHANNEL_FAILURE"
	case HtlcAttemptFailureCodeExpiryTooFar:
		s = "EXPIRY_TOO_FAR"
	case HtlcAttemptFailureCodeMppTimeout:
		s = "MPP_TIMEOUT"
	case HtlcAttemptFailureCodeInvalidOnionPayload:
		s = "INVALID_ONION_PAYLOAD"
	case HtlcAttemptFailureCodeInternalFailure:
		s = "INTERNAL_FAILURE"
	case HtlcAttemptFailureCodeUnknownFailure:
		s = "UNKNOWN_FAILURE"
	case HtlcAttemptFailureCodeUnreadableFailure:
		s = "UNREADABLE_FAILURE"

	}
	return s
}

func (a HtlcAttemptFailureCode) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
