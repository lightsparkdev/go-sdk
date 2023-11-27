// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const CANCEL_INVOICE_MUTATION = `
mutation CancelInvoice(
    $invoice_id: ID!
) {
    cancel_invoice(input: {
        invoice_id: $invoice_id
    }) {
        invoice {
            ...InvoiceFragment
        }
    }
}

` + objects.InvoiceFragment
