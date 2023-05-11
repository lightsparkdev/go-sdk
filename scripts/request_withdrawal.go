package scripts

import "lightspark/objects"


const REQUEST_WITHDRAWAL_MUTATION = `
mutation RequestWithdrawal(
    $node_id: ID!
    $amount_sats: Long!
    $bitcoin_address: String!
    $withdrawal_mode: WithdrawalMode!
) {
    request_withdrawal(input: {
        node_id: $node_id
        amount_sats: $amount_sats
        bitcoin_address: $bitcoin_address
        withdrawal_mode: $withdrawal_mode
    }) {
        request {
            ...WithdrawalRequestFragment
        }
    }
}

` + objects.WithdrawalRequestFragment