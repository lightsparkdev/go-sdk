// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const SET_INVOICE_PAYMENT_HASH = `
mutation SetInvoicePaymentHash(
  $invoice_id: ID!
  $payment_hash: Hash32!
  $preimage_nonce: Hash32!
) {
    set_invoice_payment_hash(input: {
		invoice_id: $invoice_id
		payment_hash: $payment_hash
		preimage_nonce: $preimage_nonce
	}) {
        ...SetInvoicePaymentHashOutputFragment
    }
}

` + objects.SetInvoicePaymentHashOutputFragment
