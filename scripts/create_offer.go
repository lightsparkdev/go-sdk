// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const CREATE_OFFER_MUTATION = `
mutation CreateOffer(
    $node_id: ID!
    $amount_msats: Long
    $description: String
) {
    create_offer(input: {
        node_id: $node_id
        amount_msats: $amount_msats
        description: $description
    }) {
        offer {
            ...OfferFragment
        }
    }
}

` + objects.OfferFragment
