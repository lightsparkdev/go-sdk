// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type CreateInvoiceOutput struct {
	Invoice types.EntityWrapper `json:"create_invoice_output_invoice"`
}

const (
	CreateInvoiceOutputFragment = `
fragment CreateInvoiceOutputFragment on CreateInvoiceOutput {
    __typename
    create_invoice_output_invoice: invoice {
        id
    }
}
`
)
