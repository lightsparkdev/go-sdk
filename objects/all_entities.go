// Copyright ©, 2022-present, Lightspark Group, Inc. - All Rights Reserved
package objects


const (
    GetEntityQuery = `query GetEntity($id: ID!) {
    entity(id: $id) {

        ... on Account {
            ...AccountFragment
        }
        ... on ApiToken {
            ...ApiTokenFragment
        }
        ... on Channel {
            ...ChannelFragment
        }
        ... on ChannelClosingTransaction {
            ...ChannelClosingTransactionFragment
        }
        ... on ChannelOpeningTransaction {
            ...ChannelOpeningTransactionFragment
        }
        ... on ChannelSnapshot {
            ...ChannelSnapshotFragment
        }
        ... on Deposit {
            ...DepositFragment
        }
        ... on GraphNode {
            ...GraphNodeFragment
        }
        ... on Hop {
            ...HopFragment
        }
        ... on IncomingPayment {
            ...IncomingPaymentFragment
        }
        ... on IncomingPaymentAttempt {
            ...IncomingPaymentAttemptFragment
        }
        ... on Invoice {
            ...InvoiceFragment
        }
        ... on LightsparkNodeWithOSK {
            ...LightsparkNodeWithOSKFragment
        }
        ... on LightsparkNodeWithRemoteSigning {
            ...LightsparkNodeWithRemoteSigningFragment
        }
        ... on OutgoingPayment {
            ...OutgoingPaymentFragment
        }
        ... on OutgoingPaymentAttempt {
            ...OutgoingPaymentAttemptFragment
        }
        ... on RoutingTransaction {
            ...RoutingTransactionFragment
        }
        ... on Signable {
            ...SignableFragment
        }
        ... on SignablePayload {
            ...SignablePayloadFragment
        }
        ... on UmaInvitation {
            ...UmaInvitationFragment
        }
        ... on Wallet {
            ...WalletFragment
        }
        ... on Withdrawal {
            ...WithdrawalFragment
        }
        ... on WithdrawalRequest {
            ...WithdrawalRequestFragment
        }
    }
}`  + 
    AccountFragment + 
    ApiTokenFragment + 
    ChannelFragment + 
    ChannelClosingTransactionFragment + 
    ChannelOpeningTransactionFragment + 
    ChannelSnapshotFragment + 
    DepositFragment + 
    GraphNodeFragment + 
    HopFragment + 
    IncomingPaymentFragment + 
    IncomingPaymentAttemptFragment + 
    InvoiceFragment + 
    LightsparkNodeWithOSKFragment + 
    LightsparkNodeWithRemoteSigningFragment + 
    OutgoingPaymentFragment + 
    OutgoingPaymentAttemptFragment + 
    RoutingTransactionFragment + 
    SignableFragment + 
    SignablePayloadFragment + 
    UmaInvitationFragment + 
    WalletFragment + 
    WithdrawalFragment + 
    WithdrawalRequestFragment
)
