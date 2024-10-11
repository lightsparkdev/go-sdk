package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const WITHDRAWAL_FEE_ESTIMATE_QUERY = `
query WithdrawalFeeEstimate(
	$node_id: ID!
	$amount_sats: Long!
	$withdrawal_mode: WithdrawalMode!
) {
	withdrawal_fee_estimate(input: {
		node_id: $node_id
		amount_sats: $amount_sats
		withdrawal_mode: $withdrawal_mode
	}) {
		...WithdrawalFeeEstimateOutputFragment
	}
}

` + objects.WithdrawalFeeEstimateOutputFragment
