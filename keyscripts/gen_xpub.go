// Copyright ©, 2024-present, Lightspark Group, Inc. - All Rights Reserved
package utils

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
)

func getBitcoinNetwork(network string) (uint32, *chaincfg.Params) {
	switch network {
	case "mainnet":
		return 0, &chaincfg.MainNetParams
	case "testnet":
		return 1, &chaincfg.TestNet3Params
	case "regtest":
		return 1, &chaincfg.RegressionNetParams
	default:
		log.Fatalf("Network not supported: %s", network)
		return 0, nil
	}
}

func GenHardenedXPub(masterSeedHex string, derivationPath []uint32, bitcoinNetwork string) (string, error) {
	_, params := getBitcoinNetwork(bitcoinNetwork)
	masterSeed, err := hex.DecodeString(masterSeedHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode master seed from hex: %v", err)
	}

	masterKey, err := hdkeychain.NewMaster(masterSeed, params)
	if err != nil {
		return "", fmt.Errorf("failed to create master key: %v", err)
	}

	key := masterKey
	for _, index := range derivationPath {
		key, err = key.Derive(index)
		if err != nil {
			return "", fmt.Errorf("failed to derive key: %v", err)
		}
	}

	xpub, err := key.Neuter()
	if err != nil {
		return "", fmt.Errorf("failed to neuter key: %v", err)
	}
	return xpub.String(), nil
}

func DeriveChildPubKeyFromExistingXPub(xpubStr string, remainingPath []uint32) ([]byte, error) {
	extKey, err := hdkeychain.NewKeyFromString(xpubStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse xpub: %v", err)
	}

	key := extKey
	for _, index := range remainingPath {
		if index >= 0x80000000 {
			return nil, fmt.Errorf("cannot do hardened derivation from xpub")
		}
		key, err = key.Derive(index)
		if err != nil {
			return nil, err
		}
	}

	ecPubKey, err := key.ECPubKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get EC pubkey: %v", err)
	}

	return ecPubKey.SerializeCompressed(), nil
}
