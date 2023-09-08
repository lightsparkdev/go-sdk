// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type RegisterPaymentInput struct {
	Provider ComplianceProvider `json:"register_payment_input_provider"`

	PaymentId string `json:"register_payment_input_payment_id"`

	NodePubkey string `json:"register_payment_input_node_pubkey"`

	Direction PaymentDirection `json:"register_payment_input_direction"`
}
