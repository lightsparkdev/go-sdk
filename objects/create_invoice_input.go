// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type CreateInvoiceInput struct {

	// The node from which to create the invoice.
	NodeId string `json:"create_invoice_input_node_id"`

	// The amount for which the invoice should be created, in millisatoshis.
	AmountMsats int64 `json:"create_invoice_input_amount_msats"`

	Memo *string `json:"create_invoice_input_memo"`

	InvoiceType *InvoiceType `json:"create_invoice_input_invoice_type"`

	// The expiry of the invoice in seconds. Default value is 86400 (1 day).
	ExpirySecs *int64 `json:"create_invoice_input_expiry_secs"`

	PaymentHash *string `json:"create_invoice_input_payment_hash"`

	PreimageNonce *string `json:"create_invoice_input_preimage_nonce"`
}
