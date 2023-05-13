// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
	"time"
)

type Entity interface {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	GetId() string

	// The date and time when the entity was first created.
	GetCreatedAt() time.Time

	// The date and time when the entity was last updated.
	GetUpdatedAt() time.Time
}

func EntityUnmarshal(data map[string]interface{}) (Entity, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "Account":
		var account Account
		if err := json.Unmarshal(dataJSON, &account); err != nil {
			return nil, err
		}
		return account, nil
	case "ApiToken":
		var apiToken ApiToken
		if err := json.Unmarshal(dataJSON, &apiToken); err != nil {
			return nil, err
		}
		return apiToken, nil
	case "Channel":
		var channel Channel
		if err := json.Unmarshal(dataJSON, &channel); err != nil {
			return nil, err
		}
		return channel, nil
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
	case "GraphNode":
		var graphNode GraphNode
		if err := json.Unmarshal(dataJSON, &graphNode); err != nil {
			return nil, err
		}
		return graphNode, nil
	case "Hop":
		var hop Hop
		if err := json.Unmarshal(dataJSON, &hop); err != nil {
			return nil, err
		}
		return hop, nil
	case "IncomingPayment":
		var incomingPayment IncomingPayment
		if err := json.Unmarshal(dataJSON, &incomingPayment); err != nil {
			return nil, err
		}
		return incomingPayment, nil
	case "IncomingPaymentAttempt":
		var incomingPaymentAttempt IncomingPaymentAttempt
		if err := json.Unmarshal(dataJSON, &incomingPaymentAttempt); err != nil {
			return nil, err
		}
		return incomingPaymentAttempt, nil
	case "Invoice":
		var invoice Invoice
		if err := json.Unmarshal(dataJSON, &invoice); err != nil {
			return nil, err
		}
		return invoice, nil
	case "LightsparkNode":
		var lightsparkNode LightsparkNode
		if err := json.Unmarshal(dataJSON, &lightsparkNode); err != nil {
			return nil, err
		}
		return lightsparkNode, nil
	case "OutgoingPayment":
		var outgoingPayment OutgoingPayment
		if err := json.Unmarshal(dataJSON, &outgoingPayment); err != nil {
			return nil, err
		}
		return outgoingPayment, nil
	case "OutgoingPaymentAttempt":
		var outgoingPaymentAttempt OutgoingPaymentAttempt
		if err := json.Unmarshal(dataJSON, &outgoingPaymentAttempt); err != nil {
			return nil, err
		}
		return outgoingPaymentAttempt, nil
	case "RoutingTransaction":
		var routingTransaction RoutingTransaction
		if err := json.Unmarshal(dataJSON, &routingTransaction); err != nil {
			return nil, err
		}
		return routingTransaction, nil
	case "Wallet":
		var wallet Wallet
		if err := json.Unmarshal(dataJSON, &wallet); err != nil {
			return nil, err
		}
		return wallet, nil
	case "Withdrawal":
		var withdrawal Withdrawal
		if err := json.Unmarshal(dataJSON, &withdrawal); err != nil {
			return nil, err
		}
		return withdrawal, nil
	case "WithdrawalRequest":
		var withdrawalRequest WithdrawalRequest
		if err := json.Unmarshal(dataJSON, &withdrawalRequest); err != nil {
			return nil, err
		}
		return withdrawalRequest, nil

	default:
		return nil, fmt.Errorf("unknown Entity type: %s", data["__typename"])
	}
}
