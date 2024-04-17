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
)

func TestGetNodeWalletAddressWithKeys(t *testing.T) {
	env := servicestest.NewConfig()
	client := services.NewLightsparkClient(env.ApiClientID, env.ApiClientSecret, &env.ApiClientEndpoint)
	address, err := client.CreateNodeWalletAddressWithKeys(env.NodeID)
	require.NoError(t, err)

	// Get the master xpub
	seed := env.MasterSeedHex
	ourMasterPubkey, err := crypto.DerivePublicKey(seed, lightspark_crypto.Regtest, "m")
	require.NoError(t, err)

	pubkey1, err := hex.DecodeString(address.MultisigWalletAddressValidationParameters.CounterpartyFundingPubkey)
	require.NoError(t, err)

	// Get the secp256k1 pubkey from the derivation path
	pubkey2, err := lightspark_crypto.DeriveAndTweakPubkey(ourMasterPubkey, address.MultisigWalletAddressValidationParameters.FundingPubkeyDerivationPath, nil, nil)
	require.NoError(t, err)

	generatedAddress, err := utils.GenerateMultiSigAddress(lightspark_crypto.Regtest, pubkey1, pubkey2)
	require.NoError(t, err)
	require.Equal(t, address.WalletAddress, generatedAddress)
}
