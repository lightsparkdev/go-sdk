
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
    "encoding/json"
    "strings"
)

// WebhookEventType This is an enum of the potential event types that can be associated with your Lightspark wallets.
type WebhookEventType int
const(
    WebhookEventTypeUndefined WebhookEventType = iota


    WebhookEventTypePaymentFinished

    WebhookEventTypeForceClosure

    WebhookEventTypeWithdrawalFinished

    WebhookEventTypeFundsReceived

    WebhookEventTypeNodeStatus

    WebhookEventTypeUmaInvitationClaimed

    WebhookEventTypeWalletStatus

    WebhookEventTypeWalletOutgoingPaymentFinished

    WebhookEventTypeWalletIncomingPaymentFinished

    WebhookEventTypeWalletWithdrawalFinished

    WebhookEventTypeWalletFundsReceived

    WebhookEventTypeRemoteSigning

    WebhookEventTypeLowBalance

    WebhookEventTypeHighBalance

    WebhookEventTypeChannelOpeningFees

)

func (a *WebhookEventType) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    switch s {
    default:
        *a = WebhookEventTypeUndefined
    case "PAYMENT_FINISHED":
        *a = WebhookEventTypePaymentFinished
    case "FORCE_CLOSURE":
        *a = WebhookEventTypeForceClosure
    case "WITHDRAWAL_FINISHED":
        *a = WebhookEventTypeWithdrawalFinished
    case "FUNDS_RECEIVED":
        *a = WebhookEventTypeFundsReceived
    case "NODE_STATUS":
        *a = WebhookEventTypeNodeStatus
    case "UMA_INVITATION_CLAIMED":
        *a = WebhookEventTypeUmaInvitationClaimed
    case "WALLET_STATUS":
        *a = WebhookEventTypeWalletStatus
    case "WALLET_OUTGOING_PAYMENT_FINISHED":
        *a = WebhookEventTypeWalletOutgoingPaymentFinished
    case "WALLET_INCOMING_PAYMENT_FINISHED":
        *a = WebhookEventTypeWalletIncomingPaymentFinished
    case "WALLET_WITHDRAWAL_FINISHED":
        *a = WebhookEventTypeWalletWithdrawalFinished
    case "WALLET_FUNDS_RECEIVED":
        *a = WebhookEventTypeWalletFundsReceived
    case "REMOTE_SIGNING":
        *a = WebhookEventTypeRemoteSigning
    case "LOW_BALANCE":
        *a = WebhookEventTypeLowBalance
    case "HIGH_BALANCE":
        *a = WebhookEventTypeHighBalance
    case "CHANNEL_OPENING_FEES":
        *a = WebhookEventTypeChannelOpeningFees

    }
    return nil
}

func (a WebhookEventType) StringValue() string {
    var s string
    switch a {
    default:
        s = "undefined"
    case WebhookEventTypePaymentFinished:
        s = "PAYMENT_FINISHED"
    case WebhookEventTypeForceClosure:
        s = "FORCE_CLOSURE"
    case WebhookEventTypeWithdrawalFinished:
        s = "WITHDRAWAL_FINISHED"
    case WebhookEventTypeFundsReceived:
        s = "FUNDS_RECEIVED"
    case WebhookEventTypeNodeStatus:
        s = "NODE_STATUS"
    case WebhookEventTypeUmaInvitationClaimed:
        s = "UMA_INVITATION_CLAIMED"
    case WebhookEventTypeWalletStatus:
        s = "WALLET_STATUS"
    case WebhookEventTypeWalletOutgoingPaymentFinished:
        s = "WALLET_OUTGOING_PAYMENT_FINISHED"
    case WebhookEventTypeWalletIncomingPaymentFinished:
        s = "WALLET_INCOMING_PAYMENT_FINISHED"
    case WebhookEventTypeWalletWithdrawalFinished:
        s = "WALLET_WITHDRAWAL_FINISHED"
    case WebhookEventTypeWalletFundsReceived:
        s = "WALLET_FUNDS_RECEIVED"
    case WebhookEventTypeRemoteSigning:
        s = "REMOTE_SIGNING"
    case WebhookEventTypeLowBalance:
        s = "LOW_BALANCE"
    case WebhookEventTypeHighBalance:
        s = "HIGH_BALANCE"
    case WebhookEventTypeChannelOpeningFees:
        s = "CHANNEL_OPENING_FEES"

    }
    return s
}

func (a WebhookEventType) MarshalJSON() ([]byte, error) {
    s := a.StringValue()
    return json.Marshal(s)
}
