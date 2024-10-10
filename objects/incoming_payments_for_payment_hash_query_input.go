// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type IncomingPaymentsForPaymentHashQueryInput struct {

	// PaymentHash The 32-byte hash of the payment preimage for which to fetch payments
	PaymentHash string `json:"incoming_payments_for_payment_hash_query_input_payment_hash"`

	// Statuses An optional filter to only query incoming payments of given statuses.
	Statuses *[]TransactionStatus `json:"incoming_payments_for_payment_hash_query_input_statuses"`
}
