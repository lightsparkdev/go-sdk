// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// IncomingPaymentToAttemptsConnection The connection from incoming payment to all attempts.
type IncomingPaymentToAttemptsConnection struct {

	// Count The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"incoming_payment_to_attempts_connection_count"`

	// PageInfo An object that holds pagination information about the objects in this connection.
	PageInfo PageInfo `json:"incoming_payment_to_attempts_connection_page_info"`

	// Entities The incoming payment attempts for the current page of this connection.
	Entities []IncomingPaymentAttempt `json:"incoming_payment_to_attempts_connection_entities"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
}

const (
	IncomingPaymentToAttemptsConnectionFragment = `
fragment IncomingPaymentToAttemptsConnectionFragment on IncomingPaymentToAttemptsConnection {
    __typename
    incoming_payment_to_attempts_connection_count: count
    incoming_payment_to_attempts_connection_page_info: page_info {
        __typename
        page_info_has_next_page: has_next_page
        page_info_has_previous_page: has_previous_page
        page_info_start_cursor: start_cursor
        page_info_end_cursor: end_cursor
    }
    incoming_payment_to_attempts_connection_entities: entities {
        id
    }
}
`
)

// GetCount The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
func (obj IncomingPaymentToAttemptsConnection) GetCount() int64 {
	return obj.Count
}

// GetPageInfo An object that holds pagination information about the objects in this connection.
func (obj IncomingPaymentToAttemptsConnection) GetPageInfo() PageInfo {
	return obj.PageInfo
}

func (obj IncomingPaymentToAttemptsConnection) GetTypename() string {
	return obj.Typename
}
