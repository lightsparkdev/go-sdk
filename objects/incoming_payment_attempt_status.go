
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
    "encoding/json"
    "strings"
)

// IncomingPaymentAttemptStatus This is an enum that enumerates all potential statuses for an incoming payment attempt.
type IncomingPaymentAttemptStatus int
const(
    IncomingPaymentAttemptStatusUndefined IncomingPaymentAttemptStatus = iota


    IncomingPaymentAttemptStatusAccepted

    IncomingPaymentAttemptStatusSettled

    IncomingPaymentAttemptStatusCanceled

    IncomingPaymentAttemptStatusUnknown

)

func (a *IncomingPaymentAttemptStatus) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    switch s {
    default:
        *a = IncomingPaymentAttemptStatusUndefined
    case "ACCEPTED":
        *a = IncomingPaymentAttemptStatusAccepted
    case "SETTLED":
        *a = IncomingPaymentAttemptStatusSettled
    case "CANCELED":
        *a = IncomingPaymentAttemptStatusCanceled
    case "UNKNOWN":
        *a = IncomingPaymentAttemptStatusUnknown

    }
    return nil
}

func (a IncomingPaymentAttemptStatus) StringValue() string {
    var s string
    switch a {
    default:
        s = "undefined"
    case IncomingPaymentAttemptStatusAccepted:
        s = "ACCEPTED"
    case IncomingPaymentAttemptStatusSettled:
        s = "SETTLED"
    case IncomingPaymentAttemptStatusCanceled:
        s = "CANCELED"
    case IncomingPaymentAttemptStatusUnknown:
        s = "UNKNOWN"

    }
    return s
}

func (a IncomingPaymentAttemptStatus) MarshalJSON() ([]byte, error) {
    s := a.StringValue()
    return json.Marshal(s)
}
