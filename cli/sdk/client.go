// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package sdk

import (
	"encoding/hex"
	"strings"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/spf13/viper"
)

func GetClient() *services.LightsparkClient {
	var baseUrl *string

	if viperBaseUrl := viper.GetString("LIGHTSPARK_API_ENDPOINT"); viperBaseUrl != "" {
		baseUrl = &viperBaseUrl
	}

	client := services.NewLightsparkClient(
		viper.GetString("LIGHTSPARK_API_TOKEN_CLIENT_ID"),
		viper.GetString("LIGHTSPARK_API_TOKEN_CLIENT_SECRET"),
		baseUrl,
	)

	println("Created client with:")
	println("  - Client ID: " + client.Requester.ApiTokenClientId)
	println("  - Client Secret: " + client.Requester.ApiTokenClientSecret)
	println("  - Base URL: " + *client.Requester.BaseUrl)

	return client
}

func GetClientWithSigningKey() (*services.LightsparkClient, error) {
	client := GetClient()

	masterSeed := viper.GetString("LIGHTSPARK_MASTER_SEED_HEX")
	masterSeedBytes, err := hex.DecodeString(masterSeed)
	if err != nil {
		return nil, err
	}

	network := viper.GetString("LIGHTSPARK_BITCOIN_NETWORK")
	var bitcoin_network objects.BitcoinNetwork

	switch strings.ToUpper(network) {
	case "MAINNET":
		bitcoin_network = objects.BitcoinNetworkMainnet
	case "TESTNET":
		bitcoin_network = objects.BitcoinNetworkTestnet
	case "REGTEST":
		bitcoin_network = objects.BitcoinNetworkRegtest
	default:
		bitcoin_network = objects.BitcoinNetworkUndefined
	}

	client.LoadNodeSigningKey(viper.GetString("LIGHTSPARK_NODE_ID"), *services.NewSigningKeyLoaderFromSignerMasterSeed(masterSeedBytes, bitcoin_network))

	return client, nil
}
