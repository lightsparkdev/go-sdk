package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const CREATE_UMA_INVITATION_MUTATION = `
mutation CreateUmaInvitation(
    $inviter_uma: String!
) {
    create_uma_invitation(input: {
        inviter_uma: $inviter_uma
    }) {
        invitation {
            ...UmaInvitationFragment
        }
    }
}

` + objects.UmaInvitationFragment

const CREATE_UMA_INVITATION_WITH_INCENTIVES_MUTATION = `
mutation CreateUmaInvitationWithIncentives(
    $inviter_uma: String!
    $inviter_phone_hash: String!
    $inviter_region: RegionCode!
) {
    create_uma_invitation_with_incentives(input: {
        inviter_uma: $inviter_uma
        inviter_phone_hash: $inviter_phone_hash
        inviter_region: $inviter_region
    }) {
        invitation {
            ...UmaInvitationFragment
        }
    }
}

` + objects.UmaInvitationFragment
