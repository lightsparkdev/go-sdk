// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

const CREATE_NODE_WALLET_ADDRESS_MUTATION = `
mutation CreateNodeWalletAddress(
    $node_id: ID!
) {
    create_node_wallet_address(input: {
        node_id: $node_id
    }) {
        wallet_address
    }
}
`
