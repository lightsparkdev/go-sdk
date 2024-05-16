// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type SendPaymentInput struct {

	// NodeId The node from where you want to send the payment.
	NodeId string `json:"send_payment_input_node_id"`

	// DestinationPublicKey The public key of the destination node.
	DestinationPublicKey string `json:"send_payment_input_destination_public_key"`

	// TimeoutSecs The timeout in seconds that we will try to make the payment.
	TimeoutSecs int64 `json:"send_payment_input_timeout_secs"`

	// AmountMsats The amount you will send to the destination node, expressed in msats.
	AmountMsats int64 `json:"send_payment_input_amount_msats"`

	// MaximumFeesMsats The maximum amount of fees that you want to pay for this payment to be sent, expressed in msats.
	MaximumFeesMsats int64 `json:"send_payment_input_maximum_fees_msats"`

	// IdempotencyKey The idempotency key of the request. The same result will be returned for the same idempotency key.
	IdempotencyKey *string `json:"send_payment_input_idempotency_key"`
}
