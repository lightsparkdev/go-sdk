// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type PayTestModeInvoiceInput struct {

	// NodeId The node from where you want to send the payment.
	NodeId string `json:"pay_test_mode_invoice_input_node_id"`

	// EncodedInvoice The invoice you want to pay (as defined by the BOLT11 standard).
	EncodedInvoice string `json:"pay_test_mode_invoice_input_encoded_invoice"`

	// TimeoutSecs The timeout in seconds that we will try to make the payment.
	TimeoutSecs int64 `json:"pay_test_mode_invoice_input_timeout_secs"`

	// MaximumFeesMsats The maximum amount of fees that you want to pay for this payment to be sent, expressed in msats.
	MaximumFeesMsats int64 `json:"pay_test_mode_invoice_input_maximum_fees_msats"`

	// FailureReason The failure reason to trigger for the payment. If not set, pay_invoice will be called.
	FailureReason *PaymentFailureReason `json:"pay_test_mode_invoice_input_failure_reason"`

	// AmountMsats The amount you will pay for this invoice, expressed in msats. It should ONLY be set when the invoice amount is zero.
	AmountMsats *int64 `json:"pay_test_mode_invoice_input_amount_msats"`

	// IdempotencyKey The idempotency key of the request. The same result will be returned for the same idempotency key.
	IdempotencyKey *string `json:"pay_test_mode_invoice_input_idempotency_key"`
}
