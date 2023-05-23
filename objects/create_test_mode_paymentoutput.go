// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

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
