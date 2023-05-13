// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const LIGHTNING_FEE_ESTIMATE_FOR_NODE_QUERY = `
query LightningFeeEstimateForNode(
    $node_id: ID!
    $destination_node_public_key: String!
    $amount_msats: Long!
  ) {
    lightning_fee_estimate_for_node(input: {
      node_id: $node_id,
      destination_node_public_key: $destination_node_public_key,
      amount_msats: $amount_msats
    }) {
      ...LightningFeeEstimateOutputFragment
    }
  }

` + objects.LightningFeeEstimateOutputFragment
