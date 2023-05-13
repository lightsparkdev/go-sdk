// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const BITCOIN_FEE_ESTIMATE_QUERY = `
query BitcoinFeeEstimate(
    $bitcoin_network: BitcoinNetwork!
) {
    bitcoin_fee_estimate(network: $bitcoin_network) {
        ...FeeEstimateFragment
    }
}

` + objects.FeeEstimateFragment
