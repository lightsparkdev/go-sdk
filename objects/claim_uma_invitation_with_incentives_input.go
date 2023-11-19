// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type ClaimUmaInvitationWithIncentivesInput struct {
	InvitationCode string `json:"claim_uma_invitation_with_incentives_input_invitation_code"`

	InviteeUma string `json:"claim_uma_invitation_with_incentives_input_invitee_uma"`

	InviteePhoneHash string `json:"claim_uma_invitation_with_incentives_input_invitee_phone_hash"`

	InviteeRegion RegionCode `json:"claim_uma_invitation_with_incentives_input_invitee_region"`
}
