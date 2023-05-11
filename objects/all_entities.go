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
        ... on LightsparkNode {
            ...LightsparkNodeFragment
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
}` +
		AccountFragment +
		ApiTokenFragment +
		ChannelFragment +
		ChannelClosingTransactionFragment +
		ChannelOpeningTransactionFragment +
		DepositFragment +
		GraphNodeFragment +
		HopFragment +
		IncomingPaymentFragment +
		IncomingPaymentAttemptFragment +
		InvoiceFragment +
		LightsparkNodeFragment +
		OutgoingPaymentFragment +
		OutgoingPaymentAttemptFragment +
		RoutingTransactionFragment +
		WalletFragment +
		WithdrawalFragment +
		WithdrawalRequestFragment
)
