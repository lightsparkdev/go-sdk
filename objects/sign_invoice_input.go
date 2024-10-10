
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type SignInvoiceInput struct {

    // InvoiceId The unique identifier of the invoice to be signed.
    InvoiceId string `json:"sign_invoice_input_invoice_id"`

    // Signature The cryptographic signature for the invoice.
    Signature string `json:"sign_invoice_input_signature"`

    // RecoveryId The recovery identifier for the signature.
    RecoveryId int64 `json:"sign_invoice_input_recovery_id"`

}








