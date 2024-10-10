// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type IncomingPaymentsForInvoiceQueryInput struct {
	InvoiceId string `json:"incoming_payments_for_invoice_query_input_invoice_id"`

	// Statuses An optional filter to only query outgoing payments of given statuses.
	Statuses *[]TransactionStatus `json:"incoming_payments_for_invoice_query_input_statuses"`
}
