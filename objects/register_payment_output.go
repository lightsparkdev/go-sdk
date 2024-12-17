
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type RegisterPaymentOutput struct {

    
    Payment types.EntityWrapper `json:"register_payment_output_payment"`

}

const (
    RegisterPaymentOutputFragment = `
fragment RegisterPaymentOutputFragment on RegisterPaymentOutput {
    __typename
    register_payment_output_payment: payment {
        id
    }
}
`
)







