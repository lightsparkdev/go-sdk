
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
    "encoding/json"
    "strings"
)

// TransactionType This is an enum of the potential types of transactions that can be associated with your Lightspark Node.
type TransactionType int
const(
    TransactionTypeUndefined TransactionType = iota

    // TransactionTypeOutgoingPayment Transactions initiated from a Lightspark node on Lightning Network.
    TransactionTypeOutgoingPayment
    // TransactionTypeIncomingPayment Transactions received by a Lightspark node on Lightning Network.
    TransactionTypeIncomingPayment
    // TransactionTypeRouted Transactions that forwarded payments through Lightspark nodes on Lightning Network.
    TransactionTypeRouted
    // TransactionTypeL1Withdraw Transactions on the Bitcoin blockchain to withdraw funds from a Lightspark node to a Bitcoin wallet.
    TransactionTypeL1Withdraw
    // TransactionTypeL1Deposit Transactions on Bitcoin blockchain to fund a Lightspark node's wallet.
    TransactionTypeL1Deposit
    // TransactionTypeChannelOpen Transactions on Bitcoin blockchain to open a channel on Lightning Network funded by the local Lightspark node.
    TransactionTypeChannelOpen
    // TransactionTypeChannelClose Transactions on Bitcoin blockchain to close a channel on Lightning Network where the balances are allocated back to local and remote nodes.
    TransactionTypeChannelClose
    // TransactionTypePayment Transactions initiated from a Lightspark node on Lightning Network.
    TransactionTypePayment
    // TransactionTypePaymentRequest Payment requests from a Lightspark node on Lightning Network
    TransactionTypePaymentRequest
    // TransactionTypeRoute Transactions that forwarded payments through Lightspark nodes on Lightning Network.
    TransactionTypeRoute

)

func (a *TransactionType) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    switch s {
    default:
        *a = TransactionTypeUndefined
    case "OUTGOING_PAYMENT":
        *a = TransactionTypeOutgoingPayment
    case "INCOMING_PAYMENT":
        *a = TransactionTypeIncomingPayment
    case "ROUTED":
        *a = TransactionTypeRouted
    case "L1_WITHDRAW":
        *a = TransactionTypeL1Withdraw
    case "L1_DEPOSIT":
        *a = TransactionTypeL1Deposit
    case "CHANNEL_OPEN":
        *a = TransactionTypeChannelOpen
    case "CHANNEL_CLOSE":
        *a = TransactionTypeChannelClose
    case "PAYMENT":
        *a = TransactionTypePayment
    case "PAYMENT_REQUEST":
        *a = TransactionTypePaymentRequest
    case "ROUTE":
        *a = TransactionTypeRoute

    }
    return nil
}

func (a TransactionType) StringValue() string {
    var s string
    switch a {
    default:
        s = "undefined"
    case TransactionTypeOutgoingPayment:
        s = "OUTGOING_PAYMENT"
    case TransactionTypeIncomingPayment:
        s = "INCOMING_PAYMENT"
    case TransactionTypeRouted:
        s = "ROUTED"
    case TransactionTypeL1Withdraw:
        s = "L1_WITHDRAW"
    case TransactionTypeL1Deposit:
        s = "L1_DEPOSIT"
    case TransactionTypeChannelOpen:
        s = "CHANNEL_OPEN"
    case TransactionTypeChannelClose:
        s = "CHANNEL_CLOSE"
    case TransactionTypePayment:
        s = "PAYMENT"
    case TransactionTypePaymentRequest:
        s = "PAYMENT_REQUEST"
    case TransactionTypeRoute:
        s = "ROUTE"

    }
    return s
}

func (a TransactionType) MarshalJSON() ([]byte, error) {
    s := a.StringValue()
    return json.Marshal(s)
}
