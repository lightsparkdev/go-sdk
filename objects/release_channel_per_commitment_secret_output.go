// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type ReleaseChannelPerCommitmentSecretOutput struct {

	// Channel The channel object after the per-commitment secret release operation.
	Channel types.EntityWrapper `json:"release_channel_per_commitment_secret_output_channel"`
}

const (
	ReleaseChannelPerCommitmentSecretOutputFragment = `
fragment ReleaseChannelPerCommitmentSecretOutputFragment on ReleaseChannelPerCommitmentSecretOutput {
    __typename
    release_channel_per_commitment_secret_output_channel: channel {
        id
    }
}
`
)
