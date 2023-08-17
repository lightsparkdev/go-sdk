// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type ReleaseChannelPerCommitmentSecretInput struct {
	ChannelId string `json:"release_channel_per_commitment_secret_input_channel_id"`

	PerCommitmentSecret string `json:"release_channel_per_commitment_secret_input_per_commitment_secret"`
}
