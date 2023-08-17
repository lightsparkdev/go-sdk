// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type SignMessagesOutput struct {
	SignedPayloads []SignablePayload `json:"sign_messages_output_signed_payloads"`
}

const (
	SignMessagesOutputFragment = `
fragment SignMessagesOutputFragment on SignMessagesOutput {
    __typename
    sign_messages_output_signed_payloads: signed_payloads {
        id
    }
}
`
)
