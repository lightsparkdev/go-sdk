// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

const UPDATE_NODE_SHARED_SECRET_MUTATION = `
mutation UpdateNodeSharedSecret(
  $node_id : ID!
  $shared_secret : Hash32!
) {
    update_node_shared_secret(input: {
		node_id: $node_id
		shared_secret: $shared_secret
	}) {
        ...UpdateNodeSharedSecretOutputFragment
    }
}

` + objects.UpdateNodeSharedSecretOutputFragment
