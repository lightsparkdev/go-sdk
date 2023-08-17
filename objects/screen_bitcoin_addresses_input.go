// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type ScreenBitcoinAddressesInput struct {
	Provider CryptoSanctionsScreeningProvider `json:"screen_bitcoin_addresses_input_provider"`

	Addresses []string `json:"screen_bitcoin_addresses_input_addresses"`
}
