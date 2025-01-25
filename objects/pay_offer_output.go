// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type PayOfferOutput struct {

	// Payment The payment that has been sent.
	Payment types.EntityWrapper `json:"pay_offer_output_payment"`
}

const (
	PayOfferOutputFragment = `
fragment PayOfferOutputFragment on PayOfferOutput {
    __typename
    pay_offer_output_payment: payment {
        id
    }
}
`
)
