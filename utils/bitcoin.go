package utils

import (
	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
)

func GenerateMultiSigAddress(network lightspark_crypto.BitcoinNetwork, publicKey1 []byte, publicKey2 []byte) (string, error) {
	return lightspark_crypto.GenerateMultiSigAddress(network, publicKey1, publicKey2)
}
