// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type CreateLnurlInvoiceInput struct {

	// The node from which to create the invoice.
	NodeId string `json:"create_lnurl_invoice_input_node_id"`

	// The amount for which the invoice should be created, in millisatoshis.
	AmountMsats int64 `json:"create_lnurl_invoice_input_amount_msats"`

	// The SHA256 hash of the LNURL metadata payload. This will be present in the h-tag (SHA256 purpose of payment) of the resulting Bolt 11 invoice.
	MetadataHash string `json:"create_lnurl_invoice_input_metadata_hash"`
}
