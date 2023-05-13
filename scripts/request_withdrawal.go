// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

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
