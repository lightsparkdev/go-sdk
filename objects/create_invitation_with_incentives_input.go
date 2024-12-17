
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type CreateInvitationWithIncentivesInput struct {

    // InviterUma The UMA of the user creating the invitation. It will be used to identify the inviter when receiving the invitation.
    InviterUma string `json:"create_invitation_with_incentives_input_inviter_uma"`

    // InviterPhoneHash The phone hash of the user creating the invitation.
    InviterPhoneHash string `json:"create_invitation_with_incentives_input_inviter_phone_hash"`

    // InviterRegion The region of the user creating the invitation.
    InviterRegion RegionCode `json:"create_invitation_with_incentives_input_inviter_region"`

}








