// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// OutgoingPaymentAttemptToHopsConnection The connection from an outgoing payment attempt to the list of sequential hops that define the path from sender node to recipient node.
type OutgoingPaymentAttemptToHopsConnection struct {

	// Count The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"outgoing_payment_attempt_to_hops_connection_count"`

	// PageInfo An object that holds pagination information about the objects in this connection.
	PageInfo PageInfo `json:"outgoing_payment_attempt_to_hops_connection_page_info"`

	// Entities The hops for the current page of this connection.
	Entities []Hop `json:"outgoing_payment_attempt_to_hops_connection_entities"`
}

const (
	OutgoingPaymentAttemptToHopsConnectionFragment = `
fragment OutgoingPaymentAttemptToHopsConnectionFragment on OutgoingPaymentAttemptToHopsConnection {
    __typename
    outgoing_payment_attempt_to_hops_connection_count: count
    outgoing_payment_attempt_to_hops_connection_page_info: page_info {
        __typename
        page_info_has_next_page: has_next_page
        page_info_has_previous_page: has_previous_page
        page_info_start_cursor: start_cursor
        page_info_end_cursor: end_cursor
    }
    outgoing_payment_attempt_to_hops_connection_entities: entities {
        id
    }
}
`
)

// GetCount The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
func (obj OutgoingPaymentAttemptToHopsConnection) GetCount() int64 {
	return obj.Count
}

// GetPageInfo An object that holds pagination information about the objects in this connection.
func (obj OutgoingPaymentAttemptToHopsConnection) GetPageInfo() PageInfo {
	return obj.PageInfo
}
