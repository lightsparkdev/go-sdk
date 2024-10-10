
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type OutgoingPaymentsForInvoiceQueryInput struct {

    // EncodedInvoice The encoded invoice that the outgoing payments paid to.
    EncodedInvoice string `json:"outgoing_payments_for_invoice_query_input_encoded_invoice"`

    // Statuses An optional filter to only query outgoing payments of given statuses.
    Statuses *[]TransactionStatus `json:"outgoing_payments_for_invoice_query_input_statuses"`

}








