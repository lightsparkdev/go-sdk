
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
    "encoding/json"
    "strings"
)

// IncentivesIneligibilityReason Describes the reason for an invitation to not be eligible for incentives.
type IncentivesIneligibilityReason int
const(
    IncentivesIneligibilityReasonUndefined IncentivesIneligibilityReason = iota

    // IncentivesIneligibilityReasonDisabled This invitation is not eligible for incentives because it has been created outside of the incentives flow.
    IncentivesIneligibilityReasonDisabled
    // IncentivesIneligibilityReasonSenderNotEligible This invitation is not eligible for incentives because the sender is not eligible.
    IncentivesIneligibilityReasonSenderNotEligible
    // IncentivesIneligibilityReasonReceiverNotEligible This invitation is not eligible for incentives because the receiver is not eligible.
    IncentivesIneligibilityReasonReceiverNotEligible
    // IncentivesIneligibilityReasonSendingVaspNotEligible This invitation is not eligible for incentives because the sending VASP is not part of the incentives program.
    IncentivesIneligibilityReasonSendingVaspNotEligible
    // IncentivesIneligibilityReasonReceivingVaspNotEligible This invitation is not eligible for incentives because the receiving VASP is not part of the incentives program.
    IncentivesIneligibilityReasonReceivingVaspNotEligible
    // IncentivesIneligibilityReasonNotCrossBorder This invitation is not eligible for incentives because the sender and receiver are in the same region.
    IncentivesIneligibilityReasonNotCrossBorder

)

func (a *IncentivesIneligibilityReason) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    switch s {
    default:
        *a = IncentivesIneligibilityReasonUndefined
    case "DISABLED":
        *a = IncentivesIneligibilityReasonDisabled
    case "SENDER_NOT_ELIGIBLE":
        *a = IncentivesIneligibilityReasonSenderNotEligible
    case "RECEIVER_NOT_ELIGIBLE":
        *a = IncentivesIneligibilityReasonReceiverNotEligible
    case "SENDING_VASP_NOT_ELIGIBLE":
        *a = IncentivesIneligibilityReasonSendingVaspNotEligible
    case "RECEIVING_VASP_NOT_ELIGIBLE":
        *a = IncentivesIneligibilityReasonReceivingVaspNotEligible
    case "NOT_CROSS_BORDER":
        *a = IncentivesIneligibilityReasonNotCrossBorder

    }
    return nil
}

func (a IncentivesIneligibilityReason) StringValue() string {
    var s string
    switch a {
    default:
        s = "undefined"
    case IncentivesIneligibilityReasonDisabled:
        s = "DISABLED"
    case IncentivesIneligibilityReasonSenderNotEligible:
        s = "SENDER_NOT_ELIGIBLE"
    case IncentivesIneligibilityReasonReceiverNotEligible:
        s = "RECEIVER_NOT_ELIGIBLE"
    case IncentivesIneligibilityReasonSendingVaspNotEligible:
        s = "SENDING_VASP_NOT_ELIGIBLE"
    case IncentivesIneligibilityReasonReceivingVaspNotEligible:
        s = "RECEIVING_VASP_NOT_ELIGIBLE"
    case IncentivesIneligibilityReasonNotCrossBorder:
        s = "NOT_CROSS_BORDER"

    }
    return s
}

func (a IncentivesIneligibilityReason) MarshalJSON() ([]byte, error) {
    s := a.StringValue()
    return json.Marshal(s)
}
