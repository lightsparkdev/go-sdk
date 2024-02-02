// Copyright ©, 2024-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

const SET_RIPCORD_ADDRESS_SPARKNODE_MUTATION = `
mutation SetRipcordAddressSparknode(
    $node_id: ID!
	$ripcord_address: String!
) {
    set_ripcord_address_sparknode(input: {
        node_id: $node_id
		ripcord_address: $ripcord_address
    })
}
`