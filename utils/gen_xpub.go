// Copyright ©, 2024-present, Lightspark Group, Inc. - All Rights Reserved
package utils

import (
	"encoding/hex"
	"log"

	// "flag"

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

func GenHardenedXPub(masterSeedHex string, bitcoinNetwork string) (string, error) {
	networkID, params := getBitcoinNetwork(bitcoinNetwork)
	masterSeed, err := hex.DecodeString(masterSeedHex)
	if err != nil {
		log.Fatalf("Failed to decode master seed from hex: %v", err)
	}

	masterKey, err := hdkeychain.NewMaster(masterSeed, params)
	if err != nil {
		log.Fatalf("Failed to create master key: %v", err)
	}

	accountKey, err := masterKey.Derive(hdkeychain.HardenedKeyStart + 84)
	if err != nil {
		log.Fatalf("Failed to derive key: %v", err)
	}

	purposeKey, err := accountKey.Derive(hdkeychain.HardenedKeyStart + networkID)
	if err != nil {
		log.Fatalf("Failed to derive key: %v", err)
	}

	changeKey, err := purposeKey.Derive(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		log.Fatalf("Failed to derive key: %v", err)
	}

	xpub, err := changeKey.Neuter()
	if err != nil {
		log.Fatalf("Failed to neuter key: %v", err)
	}
	return xpub.String(), nil
}

// Uncomment to run as an executable script for customers

// func main() {
// 	// Flags, both default to empty string
// 	// Ex: go run gen_xpub.go --masterseed="[seed]" --network "[mainnet]"
// 	masterSeed := flag.String("masterseed", "", "Master seed per account")
// 	bitcoinNetwork := flag.String("network", "", "Bitcoin network of account")
// 	flag.Parse()
// 	if *masterSeed == "" {
// 		log.Fatal("Please provide master seed using --masterseed!")
// 	}
// 	if *bitcoinNetwork == "" {
// 		log.Fatal("Please provide bitcoin network using --network! (ex. mainnet, testnet, regtest)")
// 	}

// 	// Generate hardened xpub
// 	xpub, err := GenHardenedXPub(*masterSeed, *bitcoinNetwork)
// 	if err != nil {
// 		log.Fatalf("Error generating hardened xpub: %v", err)
// 	}
// 	log.Printf("Extended Public Key: %s\n", xpub)
// }
