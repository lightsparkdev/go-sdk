// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

const API_ENDPOINT = "API_ENDPOINT"
const API_CLIENT_ID = "API_CLIENT_ID"
const API_CLIENT_SECRET = "API_CLIENT_SECRET"
const WEBHOOK_SECRET = "WEBHOOK_SECRET"
const MASTER_SEED_HEX = "MASTER_SEED_HEX"

type Config struct {
	ApiEndpoint     *string
	ApiClientId     string
	ApiClientSecret string
	WebhookSecret   string
	MasterSeed      []byte
}

func NewConfigFromEnv() (*Config, error) {
	var apiEndpoint *string
	apiEndpointStr := os.Getenv(API_ENDPOINT)
	if apiEndpointStr != "" {
		apiEndpoint = &apiEndpointStr
	}

	apiClientId, err := lookupEnv(API_CLIENT_ID)
	if err != nil {
		return nil, err
	}

	apiClientSecret, err := lookupEnv(API_CLIENT_SECRET)
	if err != nil {
		return nil, err
	}

	webhookSecret, err := lookupEnv(WEBHOOK_SECRET)
	if err != nil {
		return nil, err
	}

	masterSeedHex, err := lookupEnv(MASTER_SEED_HEX)
	if err != nil {
		return nil, err
	}

	masterSeed, err := hex.DecodeString(masterSeedHex)
	if err != nil {
		return nil, fmt.Errorf("invalid master seed: %s", err)
	}

	log.Print("Loaded configuration:")
	log.Printf("  - API_ENDPOINT: %s", showEmpty(apiEndpointStr))
	log.Printf("  - API_CLIENT_ID: %s", showEmpty(apiClientId))
	log.Printf("  - API_CLIENT_SECRET: %s", showEmpty(apiClientSecret))
	log.Printf("  - WEBHOOK_SECRET: %s", showEmpty(webhookSecret))
	log.Printf("  - MASTER_SEED_HEX: %s", showEmpty(masterSeedHex))

	return &Config{
		ApiEndpoint:     apiEndpoint,
		ApiClientId:     apiClientId,
		ApiClientSecret: apiClientSecret,
		WebhookSecret:   webhookSecret,
		MasterSeed:      masterSeed,
	}, nil
}

func showEmpty(str string) string {
	if str == "" {
		return "<empty>"
	}

	return str
}

func lookupEnv(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("missing configuration: %s", key)
	}

	return value, nil
}
