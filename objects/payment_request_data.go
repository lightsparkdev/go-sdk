// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
)

// PaymentRequestData This object is an interface of a payment request on the Lightning Network (i.e., a Lightning Invoice). It contains data related to parsing the payment details of a Lightning Invoice.
type PaymentRequestData interface {
	GetEncodedPaymentRequest() string

	GetBitcoinNetwork() BitcoinNetwork
}

func PaymentRequestDataUnmarshal(data map[string]interface{}) (PaymentRequestData, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "InvoiceData":
		var invoiceData InvoiceData
		if err := json.Unmarshal(dataJSON, &invoiceData); err != nil {
			return nil, err
		}
		return invoiceData, nil

	default:
		return nil, fmt.Errorf("unknown PaymentRequestData type: %s", data["__typename"])
	}
}
