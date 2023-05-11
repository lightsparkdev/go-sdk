// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// The connection from incoming payment to all attempts.
type IncomingPaymentToAttemptsConnection struct {

	// The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"incoming_payment_to_attempts_connection_count"`

	// The incoming payment attempts for the current page of this connection.
	Entities []IncomingPaymentAttempt `json:"incoming_payment_to_attempts_connection_entities"`
}

const (
	IncomingPaymentToAttemptsConnectionFragment = `
fragment IncomingPaymentToAttemptsConnectionFragment on IncomingPaymentToAttemptsConnection {
    __typename
    incoming_payment_to_attempts_connection_count: count
    incoming_payment_to_attempts_connection_entities: entities {
        id
    }
}
`
)
