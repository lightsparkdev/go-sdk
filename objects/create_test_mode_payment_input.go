
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type CreateTestModePaymentInput struct {

    // LocalNodeId The node to where you want to send the payment.
    LocalNodeId string `json:"create_test_mode_payment_input_local_node_id"`

    // EncodedInvoice The invoice you want to be paid (as defined by the BOLT11 standard).
    EncodedInvoice string `json:"create_test_mode_payment_input_encoded_invoice"`

    // AmountMsats The amount you will be paid for this invoice, expressed in msats. It should ONLY be set when the invoice amount is zero.
    AmountMsats *int64 `json:"create_test_mode_payment_input_amount_msats"`

}








