// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "lightspark/objects"

const LIGHTNING_FEE_ESTIMATE_FOR_INVOICE_QUERY = `
query LightningFeeEstimateForInvoice(
    $node_id: ID!
    $encoded_payment_request: String!
    $amount_msats: Long
  ) {
    lightning_fee_estimate_for_invoice(input: {
      node_id: $node_id,
      encoded_payment_request: $encoded_payment_request,
      amount_msats: $amount_msats
    }) {
      ...LightningFeeEstimateOutputFragment
    }
  }

` + objects.LightningFeeEstimateOutputFragment
