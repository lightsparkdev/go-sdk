// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
)

type PaymentRequest interface {
	Entity

	// The details of the payment request.
	GetData() PaymentRequestData

	// The status of the payment request.
	GetStatus() PaymentRequestStatus
}

type PaymentRequestUnmarshaler struct {
	Object PaymentRequest
}

func (unmarshaler *PaymentRequestUnmarshaler) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var t PaymentRequest
	switch string(raw["__typename"]) {
	case `"Invoice"`:
		var invoice Invoice
		if err := json.Unmarshal(data, &invoice); err != nil {
			return err
		}
		t = &invoice

	default:
		return fmt.Errorf("unknown PaymentRequest type: %s", raw["__typename"])
	}

	unmarshaler.Object = t
	return nil
}
