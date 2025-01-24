// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type CreateOfferInput struct {

	// NodeId The node from which to create the offer.
	NodeId string `json:"create_offer_input_node_id"`

	// AmountMsats The amount for which the offer should be created, in millisatoshis. Setting the amount to 0 will allow the payer to specify an amount.
	AmountMsats *int64 `json:"create_offer_input_amount_msats"`

	// Description A short description of the offer.
	Description *string `json:"create_offer_input_description"`
}
