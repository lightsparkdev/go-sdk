
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type LightningFeeEstimateForInvoiceInput struct {

    // NodeId The node from where you want to send the payment.
    NodeId string `json:"lightning_fee_estimate_for_invoice_input_node_id"`

    // EncodedPaymentRequest The invoice you want to pay (as defined by the BOLT11 standard).
    EncodedPaymentRequest string `json:"lightning_fee_estimate_for_invoice_input_encoded_payment_request"`

    // AmountMsats If the invoice does not specify a payment amount, then the amount that you wish to pay, expressed in msats.
    AmountMsats *int64 `json:"lightning_fee_estimate_for_invoice_input_amount_msats"`

}








