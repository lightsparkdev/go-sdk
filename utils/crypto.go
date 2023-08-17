// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package utils

import (
	"crypto/sha256"
	"encoding/hex"

	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
)

func Sha256HexString(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func DerivePublicKey(seedHexString string, network lightspark_crypto.BitcoinNetwork, derivationPath string) (string, error) {
	seedBytes, err := hex.DecodeString(seedHexString)
	if err != nil {
		return "", err
	}

	return lightspark_crypto.DerivePublicKey(seedBytes, network, derivationPath)
}

func ECDH(seedBytes []byte, network lightspark_crypto.BitcoinNetwork, otherPubKey string) (string, error) {
	otherPubKeyBytes, err := hex.DecodeString(otherPubKey)
	if err != nil {
		return "", err
	}

	secretBytes, err := lightspark_crypto.Ecdh(seedBytes, network, otherPubKeyBytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(secretBytes), nil
}
