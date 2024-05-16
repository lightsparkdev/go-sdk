package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const OUTGOING_PAYMENT_FOR_IDEMPOTENCY_KEY_QUERY = `
query OutgoingPaymentForIdempotencyKey(
	$idempotency_key: String!
) {
	outgoing_payment_for_idempotency_key(input: {
		idempotency_key: $idempotency_key
	}) {
		...OutgoingPaymentForIdempotencyKeyOutputFragment
	}
}

` + objects.OutgoingPaymentForIdempotencyKeyOutputFragment
