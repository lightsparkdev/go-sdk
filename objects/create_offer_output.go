// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type CreateOfferOutput struct {
	Offer types.EntityWrapper `json:"create_offer_output_offer"`
}

const (
	CreateOfferOutputFragment = `
fragment CreateOfferOutputFragment on CreateOfferOutput {
    __typename
    create_offer_output_offer: offer {
        id
    }
}
`
)
