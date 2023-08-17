// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const RELEASE_PAYMENT_PREIMAGE_MUTATION = `
mutation ReleasePaymentPreimage(
  $invoice_id: ID!
  $payment_preimage: Hash32!
) {
    release_payment_preimage(input: {
		invoice_id: $invoice_id
		payment_preimage: $payment_preimage
	}) {
        ...ReleasePaymentPreimageOutputFragment
    }
}

` + objects.ReleasePaymentPreimageOutputFragment
