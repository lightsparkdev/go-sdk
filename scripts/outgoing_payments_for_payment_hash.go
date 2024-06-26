package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const OUTGOING_PAYMENTS_FOR_PAYMENT_HASH_QUERY = `
query OutgoingPaymentsForPaymentHash(
	$payment_hash: Hash32!
	$statuses: [TransactionStatus!]
) {
	outgoing_payments_for_payment_hash(input: {
		payment_hash: $payment_hash
		statuses: $statuses
	}) {
		payments {
			...OutgoingPaymentFragment
		}
	}
}

` + objects.OutgoingPaymentFragment
