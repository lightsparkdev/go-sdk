// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const CREATE_TEST_MODE_PAYMENT_MUTATION = `
mutation CreateTestModePayment(
    $local_node_id: ID!
    $encoded_invoice: String!
    $amount_msats: Long
) {
    create_test_mode_payment(input: {
        local_node_id: $local_node_id
        encoded_invoice: $encoded_invoice
        amount_msats: $amount_msats
    }) {
		incoming_payment {
			...IncomingPaymentFragment
		}
    }
}

` + objects.IncomingPaymentFragment
