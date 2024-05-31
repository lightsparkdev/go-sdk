// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const REQUEST_WITHDRAWAL_MUTATION = `
mutation RequestWithdrawal(
    $node_id: ID!
    $amount_sats: Long!
    $bitcoin_address: String!
    $withdrawal_mode: WithdrawalMode!
	$idempotency_key: String
	$fee_target: OnChainFeeTarget
	$sats_per_vbyte: Int
) {
    request_withdrawal(input: {
        node_id: $node_id
        amount_sats: $amount_sats
        bitcoin_address: $bitcoin_address
        withdrawal_mode: $withdrawal_mode
		idempotency_key: $idempotency_key
		fee_target: $fee_target
		sats_per_vbyte: $sats_per_vbyte
    }) {
        request {
            ...WithdrawalRequestFragment
        }
    }
}

` + objects.WithdrawalRequestFragment
