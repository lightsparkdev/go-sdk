// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type ClaimUmaInvitationInput struct {
	InvitationCode string `json:"claim_uma_invitation_input_invitation_code"`

	InviteeUma string `json:"claim_uma_invitation_input_invitee_uma"`
}
