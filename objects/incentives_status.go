
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
    "encoding/json"
    "strings"
)

// IncentivesStatus Describes the status of the incentives for this invitation.
type IncentivesStatus int
const(
    IncentivesStatusUndefined IncentivesStatus = iota

    // IncentivesStatusPending The invitation is eligible for incentives in its current state. When it is claimed, we will reassess.
    IncentivesStatusPending
    // IncentivesStatusValidated The incentives have been validated.
    IncentivesStatusValidated
    // IncentivesStatusIneligible This invitation is not eligible for incentives. A more detailed reason can be found in the `incentives_ineligibility_reason` field.
    IncentivesStatusIneligible

)

func (a *IncentivesStatus) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    switch s {
    default:
        *a = IncentivesStatusUndefined
    case "PENDING":
        *a = IncentivesStatusPending
    case "VALIDATED":
        *a = IncentivesStatusValidated
    case "INELIGIBLE":
        *a = IncentivesStatusIneligible

    }
    return nil
}

func (a IncentivesStatus) StringValue() string {
    var s string
    switch a {
    default:
        s = "undefined"
    case IncentivesStatusPending:
        s = "PENDING"
    case IncentivesStatusValidated:
        s = "VALIDATED"
    case IncentivesStatusIneligible:
        s = "INELIGIBLE"

    }
    return s
}

func (a IncentivesStatus) MarshalJSON() ([]byte, error) {
    s := a.StringValue()
    return json.Marshal(s)
}
