
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type ClaimUmaInvitationWithIncentivesOutput struct {

    
    Invitation types.EntityWrapper `json:"claim_uma_invitation_with_incentives_output_invitation"`

}

const (
    ClaimUmaInvitationWithIncentivesOutputFragment = `
fragment ClaimUmaInvitationWithIncentivesOutputFragment on ClaimUmaInvitationWithIncentivesOutput {
    __typename
    claim_uma_invitation_with_incentives_output_invitation: invitation {
        id
    }
}
`
)







