
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type UpdateNodeSharedSecretOutput struct {

    
    Node types.EntityWrapper `json:"update_node_shared_secret_output_node"`

}

const (
    UpdateNodeSharedSecretOutputFragment = `
fragment UpdateNodeSharedSecretOutputFragment on UpdateNodeSharedSecretOutput {
    __typename
    update_node_shared_secret_output_node: node {
        id
    }
}
`
)







