// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "lightspark/types"

type RequestWithdrawalOutput struct {

	// The request that is created for this withdrawal.
	Request types.EntityWrapper `json:"request_withdrawal_output_request"`
}

const (
	RequestWithdrawalOutputFragment = `
fragment RequestWithdrawalOutputFragment on RequestWithdrawalOutput {
    __typename
    request_withdrawal_output_request: request {
        id
    }
}
`
)
