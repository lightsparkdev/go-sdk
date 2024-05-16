// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type PayUmaInvoiceInput struct {
	NodeId string `json:"pay_uma_invoice_input_node_id"`

	EncodedInvoice string `json:"pay_uma_invoice_input_encoded_invoice"`

	TimeoutSecs int64 `json:"pay_uma_invoice_input_timeout_secs"`

	MaximumFeesMsats int64 `json:"pay_uma_invoice_input_maximum_fees_msats"`

	AmountMsats *int64 `json:"pay_uma_invoice_input_amount_msats"`

	IdempotencyKey *string `json:"pay_uma_invoice_input_idempotency_key"`
}
