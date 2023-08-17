// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type ScreenBitcoinAddressesOutput struct {
	Ratings []RiskRating `json:"screen_bitcoin_addresses_output_ratings"`
}

const (
	ScreenBitcoinAddressesOutputFragment = `
fragment ScreenBitcoinAddressesOutputFragment on ScreenBitcoinAddressesOutput {
    __typename
    screen_bitcoin_addresses_output_ratings: ratings
}
`
)
