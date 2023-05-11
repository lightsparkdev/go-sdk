// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// The connection from an outgoing payment attempt to the list of sequential hops that define the path from sender node to recipient node.
type OutgoingPaymentAttemptToHopsConnection struct {

	// The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"outgoing_payment_attempt_to_hops_connection_count"`

	// The hops for the current page of this connection.
	Entities []Hop `json:"outgoing_payment_attempt_to_hops_connection_entities"`
}

const (
	OutgoingPaymentAttemptToHopsConnectionFragment = `
fragment OutgoingPaymentAttemptToHopsConnectionFragment on OutgoingPaymentAttemptToHopsConnection {
    __typename
    outgoing_payment_attempt_to_hops_connection_count: count
    outgoing_payment_attempt_to_hops_connection_entities: entities {
        id
    }
}
`
)
