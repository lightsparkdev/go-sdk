
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type SignInvoiceOutput struct {

    // Invoice  The signed invoice object.
    Invoice types.EntityWrapper `json:"sign_invoice_output_invoice"`

}

const (
    SignInvoiceOutputFragment = `
fragment SignInvoiceOutputFragment on SignInvoiceOutput {
    __typename
    sign_invoice_output_invoice: invoice {
        id
    }
}
`
)







