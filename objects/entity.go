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

type EntityUnmarshaler struct {
	Object Entity
}

func (unmarshaler *EntityUnmarshaler) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var t Entity
	switch string(raw["__typename"]) {
	case `"Account"`:
		var account Account
		if err := json.Unmarshal(data, &account); err != nil {
			return err
		}
		t = &account
	case `"ApiToken"`:
		var apiToken ApiToken
		if err := json.Unmarshal(data, &apiToken); err != nil {
			return err
		}
		t = &apiToken
	case `"Channel"`:
		var channel Channel
		if err := json.Unmarshal(data, &channel); err != nil {
			return err
		}
		t = &channel
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
	case `"GraphNode"`:
		var graphNode GraphNode
		if err := json.Unmarshal(data, &graphNode); err != nil {
			return err
		}
		t = &graphNode
	case `"Hop"`:
		var hop Hop
		if err := json.Unmarshal(data, &hop); err != nil {
			return err
		}
		t = &hop
	case `"IncomingPayment"`:
		var incomingPayment IncomingPayment
		if err := json.Unmarshal(data, &incomingPayment); err != nil {
			return err
		}
		t = &incomingPayment
	case `"IncomingPaymentAttempt"`:
		var incomingPaymentAttempt IncomingPaymentAttempt
		if err := json.Unmarshal(data, &incomingPaymentAttempt); err != nil {
			return err
		}
		t = &incomingPaymentAttempt
	case `"Invoice"`:
		var invoice Invoice
		if err := json.Unmarshal(data, &invoice); err != nil {
			return err
		}
		t = &invoice
	case `"LightsparkNode"`:
		var lightsparkNode LightsparkNode
		if err := json.Unmarshal(data, &lightsparkNode); err != nil {
			return err
		}
		t = &lightsparkNode
	case `"OutgoingPayment"`:
		var outgoingPayment OutgoingPayment
		if err := json.Unmarshal(data, &outgoingPayment); err != nil {
			return err
		}
		t = &outgoingPayment
	case `"OutgoingPaymentAttempt"`:
		var outgoingPaymentAttempt OutgoingPaymentAttempt
		if err := json.Unmarshal(data, &outgoingPaymentAttempt); err != nil {
			return err
		}
		t = &outgoingPaymentAttempt
	case `"RoutingTransaction"`:
		var routingTransaction RoutingTransaction
		if err := json.Unmarshal(data, &routingTransaction); err != nil {
			return err
		}
		t = &routingTransaction
	case `"Wallet"`:
		var wallet Wallet
		if err := json.Unmarshal(data, &wallet); err != nil {
			return err
		}
		t = &wallet
	case `"Withdrawal"`:
		var withdrawal Withdrawal
		if err := json.Unmarshal(data, &withdrawal); err != nil {
			return err
		}
		t = &withdrawal
	case `"WithdrawalRequest"`:
		var withdrawalRequest WithdrawalRequest
		if err := json.Unmarshal(data, &withdrawalRequest); err != nil {
			return err
		}
		t = &withdrawalRequest

	default:
		return fmt.Errorf("unknown Entity type: %s", raw["__typename"])
	}

	unmarshaler.Object = t
	return nil
}
