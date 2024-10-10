// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type DeclineToSignMessagesInput struct {

	// PayloadIds List of payload ids to decline to sign because validation failed.
	PayloadIds []string `json:"decline_to_sign_messages_input_payload_ids"`
}
