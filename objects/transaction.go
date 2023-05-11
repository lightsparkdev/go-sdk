// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
	"time"
)

type Transaction interface {
	Entity

	// The current status of this transaction.
	GetStatus() TransactionStatus

	// The date and time when this transaction was completed or failed.
	GetResolvedAt() *time.Time

	// The amount of money involved in this transaction.
	GetAmount() CurrencyAmount

	// The hash of this transaction, so it can be uniquely identified on the Lightning Network.
	GetTransactionHash() *string
}

type TransactionUnmarshaler struct {
	Object Transaction
}

func (unmarshaler *TransactionUnmarshaler) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var t Transaction
	switch string(raw["__typename"]) {
	case `"ChannelClosingTransaction"`:
		var channelClosingTransaction ChannelClosingTransaction
		if err := json.Unmarshal(data, &channelClosingTransaction); err != nil {
			return err
		}
		t = &channelClosingTransaction
	case `"ChannelOpeningTransaction"`:
		var channelOpeningTransaction ChannelOpeningTransaction
		if err := json.Unmarshal(data, &channelOpeningTransaction); err != nil {
			return err
		}
		t = &channelOpeningTransaction
	case `"Deposit"`:
		var deposit Deposit
		if err := json.Unmarshal(data, &deposit); err != nil {
			return err
		}
		t = &deposit
	case `"IncomingPayment"`:
		var incomingPayment IncomingPayment
		if err := json.Unmarshal(data, &incomingPayment); err != nil {
			return err
		}
		t = &incomingPayment
	case `"OutgoingPayment"`:
		var outgoingPayment OutgoingPayment
		if err := json.Unmarshal(data, &outgoingPayment); err != nil {
			return err
		}
		t = &outgoingPayment
	case `"RoutingTransaction"`:
		var routingTransaction RoutingTransaction
		if err := json.Unmarshal(data, &routingTransaction); err != nil {
			return err
		}
		t = &routingTransaction
	case `"Withdrawal"`:
		var withdrawal Withdrawal
		if err := json.Unmarshal(data, &withdrawal); err != nil {
			return err
		}
		t = &withdrawal

	default:
		return fmt.Errorf("unknown Transaction type: %s", raw["__typename"])
	}

	unmarshaler.Object = t
	return nil
}
