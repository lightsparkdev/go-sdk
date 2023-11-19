package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const CLAIM_UMA_INVITATION_MUTATION = `
mutation ClaimUmaInvitation(
    $invitation_code: String!
    $invitee_uma: String!
) {
    claim_uma_invitation(input: {
        invitation_code: $invitation_code
        invitee_uma: $invitee_uma
    }) {
        invitation {
            ...UmaInvitationFragment
        }
    }
}

` + objects.UmaInvitationFragment

const CLAIM_UMA_INVITATION_WITH_INCENTIVES_MUTATION = `
mutation ClaimUmaInvitationWithIncentives(
    $invitation_code: String!
    $invitee_uma: String!
    $invitee_phone_hash: String!
    $invitee_region: RegionCode!
) {
    claim_uma_invitation_with_incentives(input: {
        invitation_code: $invitation_code
        invitee_uma: $invitee_uma
        invitee_phone_hash: $invitee_phone_hash
        invitee_region: $invitee_region
    }) {
        invitation {
            ...UmaInvitationFragment
        }
    }
}

` + objects.UmaInvitationFragment
