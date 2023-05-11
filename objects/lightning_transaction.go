// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
)

type LightningTransaction interface {
	Transaction
	Entity
}

type LightningTransactionUnmarshaler struct {
	Object LightningTransaction
}

func (unmarshaler *LightningTransactionUnmarshaler) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var t LightningTransaction
	switch string(raw["__typename"]) {
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

	default:
		return fmt.Errorf("unknown LightningTransaction type: %s", raw["__typename"])
	}

	unmarshaler.Object = t
	return nil
}
