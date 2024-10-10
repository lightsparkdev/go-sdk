
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type SetInvoicePaymentHashInput struct {

    // InvoiceId The invoice that needs to be updated.
    InvoiceId string `json:"set_invoice_payment_hash_input_invoice_id"`

    // PaymentHash The 32-byte hash of the payment preimage.
    PaymentHash string `json:"set_invoice_payment_hash_input_payment_hash"`

    // PreimageNonce The 32-byte nonce used to generate the invoice preimage if applicable. It will later be included in RELEASE_PAYMENT_PREIMAGE webhook to help recover the raw preimage.
    PreimageNonce *string `json:"set_invoice_payment_hash_input_preimage_nonce"`

}








