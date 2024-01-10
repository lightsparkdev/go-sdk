
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type CancelInvoiceOutput struct {

    
    Invoice types.EntityWrapper `json:"cancel_invoice_output_invoice"`

}

const (
    CancelInvoiceOutputFragment = `
fragment CancelInvoiceOutputFragment on CancelInvoiceOutput {
    __typename
    cancel_invoice_output_invoice: invoice {
        id
    }
}
`
)







