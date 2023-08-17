// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type SignInvoiceOutput struct {
	Invoice types.EntityWrapper `json:"sign_invoice_output_invoice"`
}

const (
	SignInvoiceOutputFragment = `
fragment SignInvoiceOutputFragment on SignInvoiceOutput {
    __typename
    sign_invoice_output_invoice: invoice {
        id
    }
}
`
)
