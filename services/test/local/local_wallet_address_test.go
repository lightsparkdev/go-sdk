//go:build integration
// +build integration

package local_test

import (
	"encoding/hex"
	"testing"

	"github.com/lightsparkdev/go-sdk/crypto"
	"github.com/lightsparkdev/go-sdk/services"
	servicestest "github.com/lightsparkdev/go-sdk/services/test"
	"github.com/lightsparkdev/go-sdk/utils"
	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
	"github.com/stretchr/testify/require"
	"github.com/tyler-smith/go-bip32"
)

func TestGetNodeWalletAddressWithKeys(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	address, err := client.CreateNodeWalletAddressWithKeys(env.NodeID)
	require.NoError(t, err)
	t.Log(address.WalletAddress)
	t.Log(address.MultisigWalletAddressValidationParameters.CounterpartyFundingPubkey)
	t.Log(address.MultisigWalletAddressValidationParameters.FundingPubkeyDerivationPath)

	seed := env.MasterSeedHex
	ourPubkey, err := crypto.DerivePublicKey(seed, lightspark_crypto.Regtest, address.MultisigWalletAddressValidationParameters.FundingPubkeyDerivationPath)
	require.NoError(t, err)
	t.Log(ourPubkey)

	pubkey1, _ := hex.DecodeString(address.MultisigWalletAddressValidationParameters.CounterpartyFundingPubkey)
	xpub, _ := bip32.B58Deserialize(ourPubkey)
	pubkey2 := xpub.Key

	generatedAddress, _ := utils.GenerateMultiSigAddress(lightspark_crypto.Regtest, pubkey1, pubkey2)
	t.Log(generatedAddress)
	require.Equal(t, address.WalletAddress, generatedAddress)
}
