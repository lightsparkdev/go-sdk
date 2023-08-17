// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const UPDATE_CHANNEL_PER_COMMITMENT_POINT_MUTATION = `
mutation UpdateChannelPerCommitmentPoint(
  $channel_id : ID!
  $per_commitment_point : PublicKey!
  $per_commitment_point_index : Long!
) {
    update_channel_per_commitment_point(input: {
		channel_id: $channel_id
    per_commitment_point_index: $per_commitment_point_index
		per_commitment_point: $per_commitment_point
	}) {
        ...UpdateChannelPerCommitmentPointOutputFragment
    }
}

` + objects.UpdateChannelPerCommitmentPointOutputFragment
