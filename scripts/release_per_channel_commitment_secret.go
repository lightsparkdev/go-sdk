// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const RELEASE_CHANNEL_PER_COMMITMENT_SECRET_MUTATION = `
mutation ReleaseChannelPerCommitmentSecret(
  $channel_id: ID!
  $per_commitment_secret: Hash32!
) {
    release_channel_per_commitment_secret(input: {
		channel_id: $channel_id
		per_commitment_secret: $per_commitment_secret
	}) {
        ...ReleaseChannelPerCommitmentSecretOutputFragment
    }
}

` + objects.ReleaseChannelPerCommitmentSecretOutputFragment
