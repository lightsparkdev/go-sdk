
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type CreateInvitationWithIncentivesOutput struct {

    
    Invitation types.EntityWrapper `json:"create_invitation_with_incentives_output_invitation"`

}

const (
    CreateInvitationWithIncentivesOutputFragment = `
fragment CreateInvitationWithIncentivesOutputFragment on CreateInvitationWithIncentivesOutput {
    __typename
    create_invitation_with_incentives_output_invitation: invitation {
        id
    }
}
`
)







