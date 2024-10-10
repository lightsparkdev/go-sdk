// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type FailHtlcsOutput struct {
	Invoice types.EntityWrapper `json:"fail_htlcs_output_invoice"`
}

const (
	FailHtlcsOutputFragment = `
fragment FailHtlcsOutputFragment on FailHtlcsOutput {
    __typename
    fail_htlcs_output_invoice: invoice {
        id
    }
}
`
)
