
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
    "encoding/json"
    "strings"
)

// WithdrawalRequestStatus This is an enum of the potential statuses that a Withdrawal can take.
type WithdrawalRequestStatus int
const(
    WithdrawalRequestStatusUndefined WithdrawalRequestStatus = iota


    WithdrawalRequestStatusCreated

    WithdrawalRequestStatusFailed

    WithdrawalRequestStatusInProgress

    WithdrawalRequestStatusSuccessful

)

func (a *WithdrawalRequestStatus) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    switch s {
    default:
        *a = WithdrawalRequestStatusUndefined
    case "CREATED":
        *a = WithdrawalRequestStatusCreated
    case "FAILED":
        *a = WithdrawalRequestStatusFailed
    case "IN_PROGRESS":
        *a = WithdrawalRequestStatusInProgress
    case "SUCCESSFUL":
        *a = WithdrawalRequestStatusSuccessful

    }
    return nil
}

func (a WithdrawalRequestStatus) StringValue() string {
    var s string
    switch a {
    default:
        s = "undefined"
    case WithdrawalRequestStatusCreated:
        s = "CREATED"
    case WithdrawalRequestStatusFailed:
        s = "FAILED"
    case WithdrawalRequestStatusInProgress:
        s = "IN_PROGRESS"
    case WithdrawalRequestStatusSuccessful:
        s = "SUCCESSFUL"

    }
    return s
}

func (a WithdrawalRequestStatus) MarshalJSON() ([]byte, error) {
    s := a.StringValue()
    return json.Marshal(s)
}
