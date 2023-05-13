// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "encoding/json"

type AccountToPaymentRequestsConnection struct {

	// The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count *int64 `json:"account_to_payment_requests_connection_count"`

	// The payment requests for the current page of this connection.
	Entities []PaymentRequest `json:"account_to_payment_requests_connection_entities"`

	// An object that holds pagination information about the objects in this connection.
	PageInfo PageInfo `json:"account_to_payment_requests_connection_page_info"`
}

const (
	AccountToPaymentRequestsConnectionFragment = `
fragment AccountToPaymentRequestsConnectionFragment on AccountToPaymentRequestsConnection {
    __typename
    account_to_payment_requests_connection_count: count
    account_to_payment_requests_connection_entities: entities {
        id
    }
    account_to_payment_requests_connection_page_info: page_info {
        __typename
        page_info_has_next_page: has_next_page
        page_info_has_previous_page: has_previous_page
        page_info_start_cursor: start_cursor
        page_info_end_cursor: end_cursor
    }
}
`
)

type AccountToPaymentRequestsConnectionJSON struct {

	// The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count *int64 `json:"account_to_payment_requests_connection_count"`

	// The payment requests for the current page of this connection.
	Entities []map[string]interface{} `json:"account_to_payment_requests_connection_entities"`

	// An object that holds pagination information about the objects in this connection.
	PageInfo PageInfo `json:"account_to_payment_requests_connection_page_info"`
}

func (data *AccountToPaymentRequestsConnection) UnmarshalJSON(dataBytes []byte) error {
	var temp AccountToPaymentRequestsConnectionJSON
	if err := json.Unmarshal(dataBytes, &temp); err != nil {
		return err
	}

	data.Count = temp.Count

	if temp.Entities != nil {
		var entities []PaymentRequest
		for _, json := range temp.Entities {
			entity, err := PaymentRequestUnmarshal(json)
			if err != nil {
				return err
			}
			entities = append(entities, entity)
		}
		data.Entities = entities
	}

	data.PageInfo = temp.PageInfo

	return nil
}
