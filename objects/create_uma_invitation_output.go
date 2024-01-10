
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type CreateUmaInvitationOutput struct {

    
    Invitation types.EntityWrapper `json:"create_uma_invitation_output_invitation"`

}

const (
    CreateUmaInvitationOutputFragment = `
fragment CreateUmaInvitationOutputFragment on CreateUmaInvitationOutput {
    __typename
    create_uma_invitation_output_invitation: invitation {
        id
    }
}
`
)







