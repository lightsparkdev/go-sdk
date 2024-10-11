// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type IncomingPaymentsForPaymentHashQueryOutput struct {
	Payments []IncomingPayment `json:"incoming_payments_for_payment_hash_query_output_payments"`
}

const (
	IncomingPaymentsForPaymentHashQueryOutputFragment = `
fragment IncomingPaymentsForPaymentHashQueryOutputFragment on IncomingPaymentsForPaymentHashQueryOutput {
    __typename
    incoming_payments_for_payment_hash_query_output_payments: payments {
        id
    }
}
`
)
