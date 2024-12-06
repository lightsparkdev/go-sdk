
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
    "encoding/json"
    "strings"
)

// TransactionStatus This is an enum of the potential statuses a transaction associated with your Lightspark Node can take.
type TransactionStatus int
const(
    TransactionStatusUndefined TransactionStatus = iota

    // TransactionStatusSuccess Transaction succeeded.
    TransactionStatusSuccess
    // TransactionStatusFailed Transaction failed.
    TransactionStatusFailed
    // TransactionStatusPending Transaction has been initiated and is currently in-flight.
    TransactionStatusPending
    // TransactionStatusNotStarted For transaction type PAYMENT_REQUEST only. No payments have been made to a payment request.
    TransactionStatusNotStarted
    // TransactionStatusExpired For transaction type PAYMENT_REQUEST only. A payment request has expired.
    TransactionStatusExpired
    // TransactionStatusCancelled For transaction type PAYMENT_REQUEST only.
    TransactionStatusCancelled

)

func (a *TransactionStatus) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    switch s {
    default:
        *a = TransactionStatusUndefined
    case "SUCCESS":
        *a = TransactionStatusSuccess
    case "FAILED":
        *a = TransactionStatusFailed
    case "PENDING":
        *a = TransactionStatusPending
    case "NOT_STARTED":
        *a = TransactionStatusNotStarted
    case "EXPIRED":
        *a = TransactionStatusExpired
    case "CANCELLED":
        *a = TransactionStatusCancelled

    }
    return nil
}

func (a TransactionStatus) StringValue() string {
    var s string
    switch a {
    default:
        s = "undefined"
    case TransactionStatusSuccess:
        s = "SUCCESS"
    case TransactionStatusFailed:
        s = "FAILED"
    case TransactionStatusPending:
        s = "PENDING"
    case TransactionStatusNotStarted:
        s = "NOT_STARTED"
    case TransactionStatusExpired:
        s = "EXPIRED"
    case TransactionStatusCancelled:
        s = "CANCELLED"

    }
    return s
}

func (a TransactionStatus) MarshalJSON() ([]byte, error) {
    s := a.StringValue()
    return json.Marshal(s)
}
