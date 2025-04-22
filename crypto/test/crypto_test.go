// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package crypto_test

import (
	"encoding/hex"
	"github.com/lightsparkdev/go-sdk/crypto"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/stretchr/testify/require"
)

func TestDerivePublicKey(t *testing.T) {
	privateKeySeed := "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542"
	derivationPath := "m/0/2147483647'/1"

	publicKey, err := crypto.DerivePublicKey(privateKeySeed, derivationPath)
	require.NoError(t, err)
	require.Equal(t, "xpub6DF8uhdarytz3FWdA8TvFSvvAh8dP3283MY7p2V4SeE2wyWmG5mg5EwVvmdMVCQcoNJxGoWaU9DCWh89LojfZ537wTfunKau47EL2dhHKon", publicKey)
}

func TestCommitment(t *testing.T) {
	seedHexString := "000102030405060708090a0b0c0d0e0f"
	seedBytes, err := hex.DecodeString(seedHexString)
	require.NoError(t, err)

	derivationPath := "m/3/2104864975"
	commitmentPointIdx := uint64(281474976710654)
	commitmentPoint, err := crypto.GetPerCommitmentPoint(seedBytes, derivationPath, commitmentPointIdx)
	require.NoError(t, err)

	commitmentSecret, err := crypto.ReleasePerCommitmentSecret(seedBytes, derivationPath, commitmentPointIdx)
	require.NoError(t, err)

	privKey, _ := btcec.PrivKeyFromBytes(commitmentSecret)
	pubKey := privKey.PubKey()
	serializedPubKey := pubKey.SerializeCompressed()
	require.Equal(t, commitmentPoint, serializedPubKey)
}
