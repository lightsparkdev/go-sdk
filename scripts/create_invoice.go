// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "lightspark/objects"

const CREATE_INVOICE_MUTATION = `
mutation CreateInvoice(
    $node_id: ID!
    $amount_msats: Long!
    $memo: String
    $invoice_type: InvoiceType
) {
    create_invoice(input: {
        node_id: $node_id
        amount_msats: $amount_msats
        memo: $memo
        invoice_type: $invoice_type
    }) {
        invoice {
            ...InvoiceFragment
        }
    }
}

` + objects.InvoiceFragment
