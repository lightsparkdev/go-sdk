// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

type TransactionStatus int

const (
	TransactionStatusUndefined TransactionStatus = iota

	// Transaction succeeded..
	TransactionStatusSuccess
	// Transaction failed.
	TransactionStatusFailed
	// Transaction has been initiated and is currently in-flight.
	TransactionStatusPending
	// For transaction type PAYMENT_REQUEST only. No payments have been made to a payment request.
	TransactionStatusNotStarted
	// For transaction type PAYMENT_REQUEST only. A payment request has expired.
	TransactionStatusExpired
	// For transaction type PAYMENT_REQUEST only.
	TransactionStatusCancelled
)

func (a *TransactionStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = TransactionStatusUndefined
	case "SUCCESS":
		*a = TransactionStatusSuccess
	case "FAILED":
		*a = TransactionStatusFailed
	case "PENDING":
		*a = TransactionStatusPending
	case "NOT_STARTED":
		*a = TransactionStatusNotStarted
	case "EXPIRED":
		*a = TransactionStatusExpired
	case "CANCELLED":
		*a = TransactionStatusCancelled

	}
	return nil
}

func (a TransactionStatus) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case TransactionStatusSuccess:
		s = "SUCCESS"
	case TransactionStatusFailed:
		s = "FAILED"
	case TransactionStatusPending:
		s = "PENDING"
	case TransactionStatusNotStarted:
		s = "NOT_STARTED"
	case TransactionStatusExpired:
		s = "EXPIRED"
	case TransactionStatusCancelled:
		s = "CANCELLED"

	}
	return s
}

func (a TransactionStatus) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
