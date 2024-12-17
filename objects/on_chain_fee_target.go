
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
    "encoding/json"
    "strings"
)


type OnChainFeeTarget int
const(
    OnChainFeeTargetUndefined OnChainFeeTarget = iota

    // OnChainFeeTargetHigh Transaction expected to be confirmed within 2 blocks.
    OnChainFeeTargetHigh
    // OnChainFeeTargetMedium Transaction expected to be confirmed within 6 blocks.
    OnChainFeeTargetMedium
    // OnChainFeeTargetLow Transaction expected to be confirmed within 18 blocks.
    OnChainFeeTargetLow
    // OnChainFeeTargetBackground Transaction expected to be confirmed within 50 blocks.
    OnChainFeeTargetBackground

)

func (a *OnChainFeeTarget) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    switch s {
    default:
        *a = OnChainFeeTargetUndefined
    case "HIGH":
        *a = OnChainFeeTargetHigh
    case "MEDIUM":
        *a = OnChainFeeTargetMedium
    case "LOW":
        *a = OnChainFeeTargetLow
    case "BACKGROUND":
        *a = OnChainFeeTargetBackground

    }
    return nil
}

func (a OnChainFeeTarget) StringValue() string {
    var s string
    switch a {
    default:
        s = "undefined"
    case OnChainFeeTargetHigh:
        s = "HIGH"
    case OnChainFeeTargetMedium:
        s = "MEDIUM"
    case OnChainFeeTargetLow:
        s = "LOW"
    case OnChainFeeTargetBackground:
        s = "BACKGROUND"

    }
    return s
}

func (a OnChainFeeTarget) MarshalJSON() ([]byte, error) {
    s := a.StringValue()
    return json.Marshal(s)
}
