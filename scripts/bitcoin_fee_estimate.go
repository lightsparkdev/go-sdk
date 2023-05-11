package scripts

import "lightspark/objects"


const BITCOIN_FEE_ESTIMATE_QUERY = `
query BitcoinFeeEstimate(
    $bitcoin_network: BitcoinNetwork!
) {
    bitcoin_fee_estimate(network: $bitcoin_network) {
        ...FeeEstimateFragment
    }
}

` + objects.FeeEstimateFragment