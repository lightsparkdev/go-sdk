// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type SetInvoicePaymentHashOutput struct {
	Invoice types.EntityWrapper `json:"set_invoice_payment_hash_output_invoice"`
}

const (
	SetInvoicePaymentHashOutputFragment = `
fragment SetInvoicePaymentHashOutputFragment on SetInvoicePaymentHashOutput {
    __typename
    set_invoice_payment_hash_output_invoice: invoice {
        id
    }
}
`
)
