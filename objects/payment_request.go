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

func PaymentRequestUnmarshal(data map[string]interface{}) (PaymentRequest, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "Invoice":
		var invoice Invoice
		if err := json.Unmarshal(dataJSON, &invoice); err != nil {
			return nil, err
		}
		return invoice, nil

	default:
		return nil, fmt.Errorf("unknown PaymentRequest type: %s", data["__typename"])
	}
}
