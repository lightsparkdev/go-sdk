// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const DECLINE_TO_SIGN_MESSAGES_MUTATION = `
mutation DeclineToSignMessages($payload_ids: [ID!]!) {
	decline_to_sign_messages(input: {
		payload_ids: $payload_ids
	}) {
		...DeclineToSignMessagesOutputFragment
	}
}

` + objects.DeclineToSignMessagesOutputFragment
