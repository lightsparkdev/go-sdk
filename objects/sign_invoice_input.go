// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type SignInvoiceInput struct {
	InvoiceId string `json:"sign_invoice_input_invoice_id"`

	Signature string `json:"sign_invoice_input_signature"`

	RecoveryId int64 `json:"sign_invoice_input_recovery_id"`
}
