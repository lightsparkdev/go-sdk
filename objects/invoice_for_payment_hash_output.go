// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type InvoiceForPaymentHashOutput struct {
	Invoice *types.EntityWrapper `json:"invoice_for_payment_hash_output_invoice"`
}

const (
	InvoiceForPaymentHashOutputFragment = `
fragment InvoiceForPaymentHashOutputFragment on InvoiceForPaymentHashOutput {
    __typename
    invoice_for_payment_hash_output_invoice: invoice {
        id
    }
}
`
)
