// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
	"time"
)

// Transaction This object represents a payment transaction. The transaction can occur either on a Bitcoin Network, or over the Lightning Network. You can retrieve this object to receive specific information about a particular transaction tied to your Lightspark Node.
type Transaction interface {
	Entity

	// GetStatus The current status of this transaction.
	GetStatus() TransactionStatus

	// GetResolvedAt The date and time when this transaction was completed or failed.
	GetResolvedAt() *time.Time

	// GetAmount The amount of money involved in this transaction.
	GetAmount() CurrencyAmount

	// GetTransactionHash The hash of this transaction, so it can be uniquely identified on the Lightning Network.
	GetTransactionHash() *string
}

func TransactionUnmarshal(data map[string]interface{}) (Transaction, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "ChannelClosingTransaction":
		var channelClosingTransaction ChannelClosingTransaction
		if err := json.Unmarshal(dataJSON, &channelClosingTransaction); err != nil {
			return nil, err
		}
		return channelClosingTransaction, nil
	case "ChannelOpeningTransaction":
		var channelOpeningTransaction ChannelOpeningTransaction
		if err := json.Unmarshal(dataJSON, &channelOpeningTransaction); err != nil {
			return nil, err
		}
		return channelOpeningTransaction, nil
	case "Deposit":
		var deposit Deposit
		if err := json.Unmarshal(dataJSON, &deposit); err != nil {
			return nil, err
		}
		return deposit, nil
	case "IncomingPayment":
		var incomingPayment IncomingPayment
		if err := json.Unmarshal(dataJSON, &incomingPayment); err != nil {
			return nil, err
		}
		return incomingPayment, nil
	case "OutgoingPayment":
		var outgoingPayment OutgoingPayment
		if err := json.Unmarshal(dataJSON, &outgoingPayment); err != nil {
			return nil, err
		}
		return outgoingPayment, nil
	case "RoutingTransaction":
		var routingTransaction RoutingTransaction
		if err := json.Unmarshal(dataJSON, &routingTransaction); err != nil {
			return nil, err
		}
		return routingTransaction, nil
	case "Withdrawal":
		var withdrawal Withdrawal
		if err := json.Unmarshal(dataJSON, &withdrawal); err != nil {
			return nil, err
		}
		return withdrawal, nil

	default:
		return nil, fmt.Errorf("unknown Transaction type: %s", data["__typename"])
	}
}
