package services_test

import "os"


type TestConfig struct {
	ApiClientEndpoint string
	ApiClientID       string
	ApiClientSecret   string
	NodeID            string
	ApiClientID2      string
	ApiClientSecret2  string
	NodeID2           string
	MasterSeedHex     string
	MasterSeedHex2    string
	UmaVaspDomain     string
}

func NewConfig() TestConfig {
	var endpoint string
	if endpoint = os.Getenv("LIGHTSPARK_API_ENDPOINT"); endpoint == "" {
		endpoint = "https://api.dev.dev.sparkinfra.net/graphql/server/rc"
	}

	config := TestConfig{
		ApiClientEndpoint: endpoint,
		ApiClientID:       os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_ID"),
		ApiClientSecret:   os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_SECRET"),
		NodeID:            os.Getenv("LIGHTSPARK_RS_NODE_ID"),
		NodeID2:           os.Getenv("LIGHTSPARK_RS_NODE_ID_2"),
		ApiClientID2:      os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_ID_2"),
		ApiClientSecret2:  os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_SECRET_2"),
		MasterSeedHex:     os.Getenv("LIGHTSPARK_MASTER_SEED_HEX"),
		MasterSeedHex2:    os.Getenv("LIGHTSPARK_MASTER_SEED_HEX_2"),
		UmaVaspDomain:     os.Getenv("LIGHTSPARK_UMA_VASP_DOMAIN"),
	}

	if len(config.MasterSeedHex) == 0 && len(config.MasterSeedHex2) == 0 {
		panic("Missing master seed. Did you setup LIGHTSPARK_MASTER_SEED_HEX and LIGHTSPARK_MASTER_SEED_HEX_2 correctly in your environment?")
	}

	return config
}
