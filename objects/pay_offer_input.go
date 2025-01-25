// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type PayOfferInput struct {

	// NodeId The ID of the node that will be sending the payment.
	NodeId string `json:"pay_offer_input_node_id"`

	// EncodedOffer The Bech32 offer you want to pay (as defined by the BOLT12 standard).
	EncodedOffer string `json:"pay_offer_input_encoded_offer"`

	// TimeoutSecs The timeout in seconds that we will try to make the payment.
	TimeoutSecs int64 `json:"pay_offer_input_timeout_secs"`

	// MaximumFeesMsats The maximum amount of fees that you want to pay for this payment to be sent, expressed in msats.
	MaximumFeesMsats int64 `json:"pay_offer_input_maximum_fees_msats"`

	// AmountMsats The amount you will pay for this offer, expressed in msats. It should ONLY be set when the offer amount is zero.
	AmountMsats *int64 `json:"pay_offer_input_amount_msats"`

	// IdempotencyKey An idempotency key for this payment. If provided, it will be used to create a payment with the same idempotency key. If not provided, a new idempotency key will be generated.
	IdempotencyKey *string `json:"pay_offer_input_idempotency_key"`
}
