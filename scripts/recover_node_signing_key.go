package scripts


const RECOVER_NODE_SIGNING_KEY_QUERY = `
query RecoverNodeSigningKey(
    $node_id: ID!
) {
    entity(id: $node_id) {
        ... on LightsparkNode {
            encrypted_signing_private_key {
                encrypted_value
                cipher
            }
        }
    }
}
`