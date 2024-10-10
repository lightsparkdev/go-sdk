// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type ReleaseChannelPerCommitmentSecretInput struct {

	// ChannelId The unique identifier of the channel.
	ChannelId string `json:"release_channel_per_commitment_secret_input_channel_id"`

	// PerCommitmentSecret The per-commitment secret to be released.
	PerCommitmentSecret string `json:"release_channel_per_commitment_secret_input_per_commitment_secret"`

	// PerCommitmentIndex The index associated with the per-commitment secret.
	PerCommitmentIndex int64 `json:"release_channel_per_commitment_secret_input_per_commitment_index"`
}
