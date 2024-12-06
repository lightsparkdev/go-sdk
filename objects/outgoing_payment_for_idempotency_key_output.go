
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type OutgoingPaymentForIdempotencyKeyOutput struct {

    
    Payment *types.EntityWrapper `json:"outgoing_payment_for_idempotency_key_output_payment"`

}

const (
    OutgoingPaymentForIdempotencyKeyOutputFragment = `
fragment OutgoingPaymentForIdempotencyKeyOutputFragment on OutgoingPaymentForIdempotencyKeyOutput {
    __typename
    outgoing_payment_for_idempotency_key_output_payment: payment {
        id
    }
}
`
)







