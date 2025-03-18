// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "time"

// Offer This object represents a BOLT #12 offer (https://github.com/lightning/bolts/blob/master/12-offer-encoding.md) created by a Lightspark Node.
type Offer struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"offer_id"`

	// CreatedAt The date and time when the entity was first created.
	CreatedAt time.Time `json:"offer_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"offer_updated_at"`

	// Data The details of the offer.
	Data types.EntityWrapper `json:"offer_data"`

	// EncodedOffer The BOLT12 encoded offer. Starts with 'lno'.
	EncodedOffer string `json:"offer_encoded_offer"`

	// Amount The amount of the offer. If null, the payer chooses the amount.
	Amount *CurrencyAmount `json:"offer_amount"`

	// Description The description of the offer.
	Description *string `json:"offer_description"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
}

const (
	OfferFragment = `
fragment OfferFragment on Offer {
    __typename
    offer_id: id
    offer_created_at: created_at
    offer_updated_at: updated_at
    offer_data: data {
        id
    }
    offer_encoded_offer: encoded_offer
    offer_amount: amount {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    offer_description: description
}
`
)

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj Offer) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj Offer) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj Offer) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj Offer) GetTypename() string {
	return obj.Typename
}
