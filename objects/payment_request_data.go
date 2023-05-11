// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
)

type PaymentRequestData interface {
	GetEncodedPaymentRequest() string

	GetBitcoinNetwork() BitcoinNetwork
}

type PaymentRequestDataUnmarshaler struct {
	Object PaymentRequestData
}

func (unmarshaler *PaymentRequestDataUnmarshaler) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var t PaymentRequestData
	switch string(raw["__typename"]) {
	case `"InvoiceData"`:
		var invoiceData InvoiceData
		if err := json.Unmarshal(data, &invoiceData); err != nil {
			return err
		}
		t = &invoiceData

	default:
		return fmt.Errorf("unknown PaymentRequestData type: %s", raw["__typename"])
	}

	unmarshaler.Object = t
	return nil
}
