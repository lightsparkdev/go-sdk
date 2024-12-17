
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects



// TransactionFailures This object represents payment failures associated with your Lightspark Node.
type TransactionFailures struct {

    
    PaymentFailures *[]PaymentFailureReason `json:"transaction_failures_payment_failures"`

    
    RoutingTransactionFailures *[]RoutingTransactionFailureReason `json:"transaction_failures_routing_transaction_failures"`

}








