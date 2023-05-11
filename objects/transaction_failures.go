// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type TransactionFailures struct {
	PaymentFailures *[]PaymentFailureReason `json:"transaction_failures_payment_failures"`

	RoutingTransactionFailures *[]RoutingTransactionFailureReason `json:"transaction_failures_routing_transaction_failures"`
}
