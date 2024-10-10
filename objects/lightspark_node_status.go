
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
    "encoding/json"
    "strings"
)


type LightsparkNodeStatus int
const(
    LightsparkNodeStatusUndefined LightsparkNodeStatus = iota


    LightsparkNodeStatusCreated

    LightsparkNodeStatusDeployed

    LightsparkNodeStatusStarted

    LightsparkNodeStatusSyncing

    LightsparkNodeStatusReady

    LightsparkNodeStatusStopped

    LightsparkNodeStatusTerminated

    LightsparkNodeStatusTerminating

    LightsparkNodeStatusWalletLocked

    LightsparkNodeStatusFailedToDeploy

)

func (a *LightsparkNodeStatus) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    switch s {
    default:
        *a = LightsparkNodeStatusUndefined
    case "CREATED":
        *a = LightsparkNodeStatusCreated
    case "DEPLOYED":
        *a = LightsparkNodeStatusDeployed
    case "STARTED":
        *a = LightsparkNodeStatusStarted
    case "SYNCING":
        *a = LightsparkNodeStatusSyncing
    case "READY":
        *a = LightsparkNodeStatusReady
    case "STOPPED":
        *a = LightsparkNodeStatusStopped
    case "TERMINATED":
        *a = LightsparkNodeStatusTerminated
    case "TERMINATING":
        *a = LightsparkNodeStatusTerminating
    case "WALLET_LOCKED":
        *a = LightsparkNodeStatusWalletLocked
    case "FAILED_TO_DEPLOY":
        *a = LightsparkNodeStatusFailedToDeploy

    }
    return nil
}

func (a LightsparkNodeStatus) StringValue() string {
    var s string
    switch a {
    default:
        s = "undefined"
    case LightsparkNodeStatusCreated:
        s = "CREATED"
    case LightsparkNodeStatusDeployed:
        s = "DEPLOYED"
    case LightsparkNodeStatusStarted:
        s = "STARTED"
    case LightsparkNodeStatusSyncing:
        s = "SYNCING"
    case LightsparkNodeStatusReady:
        s = "READY"
    case LightsparkNodeStatusStopped:
        s = "STOPPED"
    case LightsparkNodeStatusTerminated:
        s = "TERMINATED"
    case LightsparkNodeStatusTerminating:
        s = "TERMINATING"
    case LightsparkNodeStatusWalletLocked:
        s = "WALLET_LOCKED"
    case LightsparkNodeStatusFailedToDeploy:
        s = "FAILED_TO_DEPLOY"

    }
    return s
}

func (a LightsparkNodeStatus) MarshalJSON() ([]byte, error) {
    s := a.StringValue()
    return json.Marshal(s)
}
