// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type RegisterPaymentInput struct {

	// Provider The compliance provider that is going to screen the node. You need to be a customer of the selected provider and store the API key on the Lightspark account setting page.
	Provider ComplianceProvider `json:"register_payment_input_provider"`

	// PaymentId The Lightspark ID of the lightning payment you want to register. It can be the id of either an OutgoingPayment or an IncomingPayment.
	PaymentId string `json:"register_payment_input_payment_id"`

	// NodePubkey The public key of the counterparty lightning node, which would be the public key of the recipient node if it is to register an outgoing payment, or the public key of the sender node if it is to register an incoming payment.
	NodePubkey string `json:"register_payment_input_node_pubkey"`

	// Direction Indicates whether this payment is an OutgoingPayment or an IncomingPayment.
	Direction PaymentDirection `json:"register_payment_input_direction"`
}
