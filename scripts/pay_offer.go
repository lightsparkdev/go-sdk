// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const PAY_OFFER_MUTATION = `
mutation PayOffer(
    $node_id: ID!
    $encoded_offer: String!
    $timeout_secs: Int!
    $maximum_fees_msats: Long!
    $amount_msats: Long
	$idempotency_key: String
) {
    pay_offer(input: {
        node_id: $node_id
        encoded_offer: $encoded_offer
        timeout_secs: $timeout_secs
        maximum_fees_msats: $maximum_fees_msats
        amount_msats: $amount_msats
		idempotency_key: $idempotency_key
    }) {
        payment {
            ...OutgoingPaymentFragment
        }
    }
}

` + objects.OutgoingPaymentFragment
