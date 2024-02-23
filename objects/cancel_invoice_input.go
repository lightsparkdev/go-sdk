// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// CancelInvoiceInput The unique identifier of the Invoice that should be cancelled. The invoice is supposed to be open, not settled and not expired.
type CancelInvoiceInput struct {
	InvoiceId string `json:"cancel_invoice_input_invoice_id"`
}
