// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const SIGN_MESSAGES_MUTATION = `
mutation SignMessages(
  $signatures: [IdAndSignature!]!
) {
    sign_messages(input: {
		signatures: $signatures
	}) {
        ...SignMessagesOutputFragment
    }
}

` + objects.SignMessagesOutputFragment
