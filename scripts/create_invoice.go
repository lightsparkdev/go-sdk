// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const CREATE_INVOICE_MUTATION = `
mutation CreateInvoice(
    $node_id: ID!
    $amount_msats: Long!
    $memo: String
    $invoice_type: InvoiceType
	$expiry_secs: Int
) {
    create_invoice(input: {
        node_id: $node_id
        amount_msats: $amount_msats
        memo: $memo
        invoice_type: $invoice_type
		expiry_secs: $expiry_secs
    }) {
        invoice {
            ...InvoiceFragment
        }
    }
}

` + objects.InvoiceFragment
