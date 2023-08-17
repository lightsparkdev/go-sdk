// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package utils

import (
	"testing"

	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
	"github.com/stretchr/testify/require"
)

func TestDerivePublicKey(t *testing.T) {
	privateKeySeed := "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542"
	derivationPath := "m/0/2147483647'/1"

	publicKey, err := DerivePublicKey(privateKeySeed, lightspark_crypto.Mainnet, derivationPath)
	require.NoError(t, err)
	require.Equal(t, "xpub6DF8uhdarytz3FWdA8TvFSvvAh8dP3283MY7p2V4SeE2wyWmG5mg5EwVvmdMVCQcoNJxGoWaU9DCWh89LojfZ537wTfunKau47EL2dhHKon", publicKey)
}
