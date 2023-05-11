// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "lightspark/types"

type DeleteApiTokenOutput struct {
	Account types.EntityWrapper `json:"delete_api_token_output_account"`
}

const (
	DeleteApiTokenOutputFragment = `
fragment DeleteApiTokenOutputFragment on DeleteApiTokenOutput {
    __typename
    delete_api_token_output_account: account {
        id
    }
}
`
)
