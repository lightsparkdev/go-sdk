// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

const REGISTER_PAYMENT_MUTATION = `
mutation RegisterPayment(
    $provider: ComplianceProvider!
    $payment_id: ID!
    $node_pubkey: String!
    $direction: PaymentDirection!
) {
    register_payment(input: {
        provider: $provider
        payment_id: $payment_id
        node_pubkey: $node_pubkey
        direction: $direction
    }) {
        payment {
            id
        }
    }
}

`
