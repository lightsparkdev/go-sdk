// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type CreateInvoiceInput struct {

	// NodeId The node from which to create the invoice.
	NodeId string `json:"create_invoice_input_node_id"`

	// AmountMsats The amount for which the invoice should be created, in millisatoshis. Setting the amount to 0 will allow the payer to specify an amount.
	AmountMsats int64 `json:"create_invoice_input_amount_msats"`

	Memo *string `json:"create_invoice_input_memo"`

	InvoiceType *InvoiceType `json:"create_invoice_input_invoice_type"`

	// ExpirySecs The expiry of the invoice in seconds. Default value is 86400 (1 day) for AMP invoice, or 3600 (1 hour) for STANDARD invoice.
	ExpirySecs *int64 `json:"create_invoice_input_expiry_secs"`

	// PaymentHash The payment hash of the invoice. It should only be set if your node is a remote signing node. If not set, it will be requested through REMOTE_SIGNING webhooks with sub event type REQUEST_INVOICE_PAYMENT_HASH.
	PaymentHash *string `json:"create_invoice_input_payment_hash"`

	// PreimageNonce The 32-byte nonce used to generate the invoice preimage if applicable. It will later be included in RELEASE_PAYMENT_PREIMAGE webhook to help recover the raw preimage. This can only be specified when `payment_hash` is specified.
	PreimageNonce *string `json:"create_invoice_input_preimage_nonce"`
}
