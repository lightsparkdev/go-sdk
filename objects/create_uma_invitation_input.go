// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type CreateUmaInvitationInput struct {

	// InviterUma The UMA of the user creating the invitation. It will be used to identify the inviter when receiving the invitation.
	InviterUma string `json:"create_uma_invitation_input_inviter_uma"`
}
