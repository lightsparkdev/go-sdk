// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const PAY_UMA_INVOICE_MUTATION = `
mutation PayUmaInvoice(
    $node_id: ID!
    $encoded_invoice: String!
    $timeout_secs: Int!
    $maximum_fees_msats: Long!
    $amount_msats: Long
) {
    pay_uma_invoice(input: {
        node_id: $node_id
        encoded_invoice: $encoded_invoice
        timeout_secs: $timeout_secs
        maximum_fees_msats: $maximum_fees_msats
        amount_msats: $amount_msats
    }) {
        payment {
            ...OutgoingPaymentFragment
        }
    }
}

` + objects.OutgoingPaymentFragment
