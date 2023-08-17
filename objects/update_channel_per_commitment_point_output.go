// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type UpdateChannelPerCommitmentPointOutput struct {
	Channel types.EntityWrapper `json:"update_channel_per_commitment_point_output_channel"`
}

const (
	UpdateChannelPerCommitmentPointOutputFragment = `
fragment UpdateChannelPerCommitmentPointOutputFragment on UpdateChannelPerCommitmentPointOutput {
    __typename
    update_channel_per_commitment_point_output_channel: channel {
        id
    }
}
`
)
