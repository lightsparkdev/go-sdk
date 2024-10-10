// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type WithdrawalRequestToWithdrawalsConnection struct {

	// Count The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"withdrawal_request_to_withdrawals_connection_count"`

	// Entities The withdrawals for the current page of this connection.
	Entities []Withdrawal `json:"withdrawal_request_to_withdrawals_connection_entities"`
}

const (
	WithdrawalRequestToWithdrawalsConnectionFragment = `
fragment WithdrawalRequestToWithdrawalsConnectionFragment on WithdrawalRequestToWithdrawalsConnection {
    __typename
    withdrawal_request_to_withdrawals_connection_count: count
    withdrawal_request_to_withdrawals_connection_entities: entities {
        id
    }
}
`
)
