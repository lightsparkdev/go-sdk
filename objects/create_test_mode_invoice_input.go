
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type CreateTestModeInvoiceInput struct {

    // LocalNodeId The local node from which to create the invoice.
    LocalNodeId string `json:"create_test_mode_invoice_input_local_node_id"`

    // AmountMsats The amount for which the invoice should be created, in millisatoshis. Setting the amount to 0 will allow the payer to specify an amount.
    AmountMsats int64 `json:"create_test_mode_invoice_input_amount_msats"`

    // Memo An optional memo to include in the invoice.
    Memo *string `json:"create_test_mode_invoice_input_memo"`

    // InvoiceType The type of invoice to create.
    InvoiceType *InvoiceType `json:"create_test_mode_invoice_input_invoice_type"`

}








