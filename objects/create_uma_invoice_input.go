
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type CreateUmaInvoiceInput struct {

    // NodeId The node from which to create the invoice.
    NodeId string `json:"create_uma_invoice_input_node_id"`

    // AmountMsats The amount for which the invoice should be created, in millisatoshis.
    AmountMsats int64 `json:"create_uma_invoice_input_amount_msats"`

    // MetadataHash The SHA256 hash of the UMA metadata payload. This will be present in the h-tag (SHA256 purpose of payment) of the resulting Bolt 11 invoice.
    MetadataHash string `json:"create_uma_invoice_input_metadata_hash"`

    // ExpirySecs The expiry of the invoice in seconds. Default value is 86400 (1 day).
    ExpirySecs *int64 `json:"create_uma_invoice_input_expiry_secs"`

    // ReceiverHash An optional, monthly-rotated, unique hashed identifier corresponding to the receiver of the payment.
    ReceiverHash *string `json:"create_uma_invoice_input_receiver_hash"`

}








