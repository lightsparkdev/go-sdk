// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type LightningFeeEstimateForNodeInput struct {

	// NodeId The node from where you want to send the payment.
	NodeId string `json:"lightning_fee_estimate_for_node_input_node_id"`

	// DestinationNodePublicKey The public key of the node that you want to pay.
	DestinationNodePublicKey string `json:"lightning_fee_estimate_for_node_input_destination_node_public_key"`

	// AmountMsats The payment amount expressed in msats.
	AmountMsats int64 `json:"lightning_fee_estimate_for_node_input_amount_msats"`
}
