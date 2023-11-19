// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type ClaimUmaInvitationOutput struct {
	Invitation types.EntityWrapper `json:"claim_uma_invitation_output_invitation"`
}

const (
	ClaimUmaInvitationOutputFragment = `
fragment ClaimUmaInvitationOutputFragment on ClaimUmaInvitationOutput {
    __typename
    claim_uma_invitation_output_invitation: invitation {
        id
    }
}
`
)
