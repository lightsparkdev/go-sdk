
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type ClaimUmaInvitationInput struct {

    // InvitationCode The unique code that identifies this invitation and was shared by the inviter.
    InvitationCode string `json:"claim_uma_invitation_input_invitation_code"`

    // InviteeUma The UMA of the user claiming the invitation. It will be sent to the inviter so that they can start transacting with the invitee.
    InviteeUma string `json:"claim_uma_invitation_input_invitee_uma"`

}








