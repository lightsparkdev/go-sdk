// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// The connection from outgoing payment to all attempts.
type OutgoingPaymentToAttemptsConnection struct {

	// The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"outgoing_payment_to_attempts_connection_count"`

	// The attempts for the current page of this connection.
	Entities []OutgoingPaymentAttempt `json:"outgoing_payment_to_attempts_connection_entities"`
}

const (
	OutgoingPaymentToAttemptsConnectionFragment = `
fragment OutgoingPaymentToAttemptsConnectionFragment on OutgoingPaymentToAttemptsConnection {
    __typename
    outgoing_payment_to_attempts_connection_count: count
    outgoing_payment_to_attempts_connection_entities: entities {
        id
    }
}
`
)
