// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// CreateTestModePaymentoutput This is an object identifying the output of a test mode payment. This object can be used to retrieve the associated payment made from a Test Mode Payment call.
type CreateTestModePaymentoutput struct {

	// Payment The payment that has been sent.
	// Deprecated: Use incoming_payment instead.
	Payment types.EntityWrapper `json:"create_test_mode_paymentoutput_payment"`

	// IncomingPayment The payment that has been received.
	IncomingPayment types.EntityWrapper `json:"create_test_mode_paymentoutput_incoming_payment"`
}

const (
	CreateTestModePaymentoutputFragment = `
fragment CreateTestModePaymentoutputFragment on CreateTestModePaymentoutput {
    __typename
    create_test_mode_paymentoutput_payment: payment {
        id
    }
    create_test_mode_paymentoutput_incoming_payment: incoming_payment {
        id
    }
}
`
)
