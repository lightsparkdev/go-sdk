// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type ReleasePaymentPreimageOutput struct {
	Invoice types.EntityWrapper `json:"release_payment_preimage_output_invoice"`
}

const (
	ReleasePaymentPreimageOutputFragment = `
fragment ReleasePaymentPreimageOutputFragment on ReleasePaymentPreimageOutput {
    __typename
    release_payment_preimage_output_invoice: invoice {
        id
    }
}
`
)
