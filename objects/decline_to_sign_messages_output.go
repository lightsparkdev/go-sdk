// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type DeclineToSignMessagesOutput struct {
	DeclinedPayloads []SignablePayload `json:"decline_to_sign_messages_output_declined_payloads"`
}

const (
	DeclineToSignMessagesOutputFragment = `
fragment DeclineToSignMessagesOutputFragment on DeclineToSignMessagesOutput {
    __typename
    decline_to_sign_messages_output_declined_payloads: declined_payloads {
        id
    }
}
`
)
