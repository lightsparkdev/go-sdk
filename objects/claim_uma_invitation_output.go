// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type ClaimUmaInvitationOutput struct {

	// Invitation An UMA.ME invitation object.
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
