
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type CreateTestModeInvoiceInput struct {

    
    LocalNodeId string `json:"create_test_mode_invoice_input_local_node_id"`

    
    AmountMsats int64 `json:"create_test_mode_invoice_input_amount_msats"`

    
    Memo *string `json:"create_test_mode_invoice_input_memo"`

    
    InvoiceType *InvoiceType `json:"create_test_mode_invoice_input_invoice_type"`

}








