// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// CancelInvoiceOutput The Invoice that was cancelled. If the invoice was already cancelled, the same invoice is returned.
type CancelInvoiceOutput struct {
	Invoice types.EntityWrapper `json:"cancel_invoice_output_invoice"`
}

const (
	CancelInvoiceOutputFragment = `
fragment CancelInvoiceOutputFragment on CancelInvoiceOutput {
    __typename
    cancel_invoice_output_invoice: invoice {
        id
    }
}
`
)
