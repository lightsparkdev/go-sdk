
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type FailHtlcsInput struct {

    // InvoiceId The id of invoice which the pending HTLCs that need to be failed are paying for.
    InvoiceId string `json:"fail_htlcs_input_invoice_id"`

    // CancelInvoice Whether the invoice needs to be canceled after failing the htlcs. If yes, the invoice cannot be paid anymore.
    CancelInvoice bool `json:"fail_htlcs_input_cancel_invoice"`

}








