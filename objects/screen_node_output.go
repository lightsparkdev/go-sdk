
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type ScreenNodeOutput struct {

    
    Rating RiskRating `json:"screen_node_output_rating"`

}

const (
    ScreenNodeOutputFragment = `
fragment ScreenNodeOutputFragment on ScreenNodeOutput {
    __typename
    screen_node_output_rating: rating
}
`
)







