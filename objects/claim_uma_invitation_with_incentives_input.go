// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type ClaimUmaInvitationWithIncentivesInput struct {

	// InvitationCode The unique code that identifies this invitation and was shared by the inviter.
	InvitationCode string `json:"claim_uma_invitation_with_incentives_input_invitation_code"`

	// InviteeUma The UMA of the user claiming the invitation. It will be sent to the inviter so that they can start transacting with the invitee.
	InviteeUma string `json:"claim_uma_invitation_with_incentives_input_invitee_uma"`

	// InviteePhoneHash The phone hash of the user getting the invitation.
	InviteePhoneHash string `json:"claim_uma_invitation_with_incentives_input_invitee_phone_hash"`

	// InviteeRegion The region of the user getting the invitation.
	InviteeRegion RegionCode `json:"claim_uma_invitation_with_incentives_input_invitee_region"`
}
