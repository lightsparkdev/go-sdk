// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type OutgoingPaymentsForPaymentHashQueryOutput struct {
	Payments []OutgoingPayment `json:"outgoing_payments_for_payment_hash_query_output_payments"`
}

const (
	OutgoingPaymentsForPaymentHashQueryOutputFragment = `
fragment OutgoingPaymentsForPaymentHashQueryOutputFragment on OutgoingPaymentsForPaymentHashQueryOutput {
    __typename
    outgoing_payments_for_payment_hash_query_output_payments: payments {
        id
    }
}
`
)
