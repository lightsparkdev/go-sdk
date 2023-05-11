package scripts

import "lightspark/objects"


const SEND_PAYMENT_MUTATION = `
mutation SendPayment(
    $node_id: ID!
    $destination_public_key: String!
    $amount_msats: Long!
    $timeout_secs: Int!
    $maximum_fees_msats: Long!
) {
    send_payment(input: {
        node_id: $node_id
        destination_public_key: $destination_public_key
        amount_msats: $amount_msats
        timeout_secs: $timeout_secs
        maximum_fees_msats: $maximum_fees_msats
    }) {
        payment {
            ...OutgoingPaymentFragment
        }
    }
}

` + objects.OutgoingPaymentFragment