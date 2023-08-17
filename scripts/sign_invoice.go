// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const SIGN_INVOICE_MUTATION = `
mutation SignInvoice(
  $invoice_id : ID!
  $signature : Signature!
  $recovery_id : Int!
) {
    sign_invoice(input: {
		invoice_id: $invoice_id
		signature: $signature
		recovery_id: $recovery_id
	}) {
        ...SignInvoiceOutputFragment
    }
}

` + objects.SignInvoiceOutputFragment
