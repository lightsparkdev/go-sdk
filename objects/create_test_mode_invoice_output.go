
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type CreateTestModeInvoiceOutput struct {

    
    EncodedPaymentRequest string `json:"create_test_mode_invoice_output_encoded_payment_request"`

}

const (
    CreateTestModeInvoiceOutputFragment = `
fragment CreateTestModeInvoiceOutputFragment on CreateTestModeInvoiceOutput {
    __typename
    create_test_mode_invoice_output_encoded_payment_request: encoded_payment_request
}
`
)







