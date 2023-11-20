package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const FETCH_UMA_INVITATION_QUERY = `
query FetchUmaInvitation(
    $invitation_code: String!
) {
    uma_invitation_by_code(code: $invitation_code) {
        ...UmaInvitationFragment
    }
}

` + objects.UmaInvitationFragment
