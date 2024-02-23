// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type CreateInvitationWithIncentivesOutput struct {

	// Invitation The created invitation in the form of a string identifier.
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
