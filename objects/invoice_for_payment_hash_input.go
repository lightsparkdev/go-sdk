// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type InvoiceForPaymentHashInput struct {

	// PaymentHash The 32-byte hash of the payment preimage for which to fetch an invoice.
	PaymentHash string `json:"invoice_for_payment_hash_input_payment_hash"`
}
