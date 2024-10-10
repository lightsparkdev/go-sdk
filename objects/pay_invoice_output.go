// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type PayInvoiceOutput struct {

	// Payment The payment that has been sent.
	Payment types.EntityWrapper `json:"pay_invoice_output_payment"`
}

const (
	PayInvoiceOutputFragment = `
fragment PayInvoiceOutputFragment on PayInvoiceOutput {
    __typename
    pay_invoice_output_payment: payment {
        id
    }
}
`
)
