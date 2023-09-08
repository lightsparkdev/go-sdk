// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type OutgoingPaymentsForInvoiceQueryOutput struct {
	Payments []OutgoingPayment `json:"outgoing_payments_for_invoice_query_output_payments"`
}

const (
	OutgoingPaymentsForInvoiceQueryOutputFragment = `
fragment OutgoingPaymentsForInvoiceQueryOutputFragment on OutgoingPaymentsForInvoiceQueryOutput {
    __typename
    outgoing_payments_for_invoice_query_output_payments: payments {
        id
    }
}
`
)
