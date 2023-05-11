// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type CreateInvoiceInput struct {
	NodeId string `json:"create_invoice_input_node_id"`

	AmountMsats int64 `json:"create_invoice_input_amount_msats"`

	Memo *string `json:"create_invoice_input_memo"`

	InvoiceType *InvoiceType `json:"create_invoice_input_invoice_type"`
}
