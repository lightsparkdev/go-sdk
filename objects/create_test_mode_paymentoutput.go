// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

// This is an object identifying the output of a test mode payment. This object can be used to retrieve the associated payment made from a Test Mode Payment call.
type CreateTestModePaymentoutput struct {

	// The payment that has been sent.
	Payment types.EntityWrapper `json:"create_test_mode_paymentoutput_payment"`
}

const (
	CreateTestModePaymentoutputFragment = `
fragment CreateTestModePaymentoutputFragment on CreateTestModePaymentoutput {
    __typename
    create_test_mode_paymentoutput_payment: payment {
        id
    }
}
`
)
