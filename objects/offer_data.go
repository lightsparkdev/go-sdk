// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "time"

// OfferData This object represents the data associated with a BOLT #12 offer. You can retrieve this object to receive the relevant data associated with a specific offer.
type OfferData struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"offer_data_id"`

	// CreatedAt The date and time when the entity was first created.
	CreatedAt time.Time `json:"offer_data_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"offer_data_updated_at"`

	// Amount The requested amount in this invoice. If it is equal to 0, the sender should choose the amount to send.
	Amount *CurrencyAmount `json:"offer_data_amount"`

	// EncodedOffer The Bech32 encoded offer.
	EncodedOffer string `json:"offer_data_encoded_offer"`

	// BitcoinNetworks The bitcoin networks supported by the offer.
	BitcoinNetworks []BitcoinNetwork `json:"offer_data_bitcoin_networks"`

	// ExpiresAt The date and time when this invoice will expire.
	ExpiresAt *time.Time `json:"offer_data_expires_at"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
}

const (
	OfferDataFragment = `
fragment OfferDataFragment on OfferData {
    __typename
    offer_data_id: id
    offer_data_created_at: created_at
    offer_data_updated_at: updated_at
    offer_data_amount: amount {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    offer_data_encoded_offer: encoded_offer
    offer_data_bitcoin_networks: bitcoin_networks
    offer_data_expires_at: expires_at
}
`
)

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj OfferData) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj OfferData) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj OfferData) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj OfferData) GetTypename() string {
	return obj.Typename
}
