
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
    "encoding/json"
    "strings"
)

// PaymentDirection This is an enum indicating the direction of the payment.
type PaymentDirection int
const(
    PaymentDirectionUndefined PaymentDirection = iota


    PaymentDirectionSent

    PaymentDirectionReceived

)

func (a *PaymentDirection) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    switch s {
    default:
        *a = PaymentDirectionUndefined
    case "SENT":
        *a = PaymentDirectionSent
    case "RECEIVED":
        *a = PaymentDirectionReceived

    }
    return nil
}

func (a PaymentDirection) StringValue() string {
    var s string
    switch a {
    default:
        s = "undefined"
    case PaymentDirectionSent:
        s = "SENT"
    case PaymentDirectionReceived:
        s = "RECEIVED"

    }
    return s
}

func (a PaymentDirection) MarshalJSON() ([]byte, error) {
    s := a.StringValue()
    return json.Marshal(s)
}
