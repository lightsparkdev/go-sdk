package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

type UmaConfig struct {
	ApiClientID                    string
	ApiClientSecret                string
	NodeUUID                       string
	Username                       string
	UserID                         string
	UmaEncryptionPubKeyHex         string
	UmaEncryptionPrivKeyHex        string
	UmaSigningPubKeyHex            string
	UmaSigningPrivKeyHex           string
	RemoteSigningNodeMasterSeedHex string
	OskNodeSigningKeyPassword      string
	ClientBaseURL                  string
	SenderVaspDomain               string
}

func (c *UmaConfig) UmaEncryptionPubKeyBytes() ([]byte, error) {
	return hex.DecodeString(c.UmaEncryptionPubKeyHex)
}

func (c *UmaConfig) UmaEncryptionPrivKeyBytes() ([]byte, error) {
	return hex.DecodeString(c.UmaEncryptionPrivKeyHex)
}

func (c *UmaConfig) UmaSigningPubKeyBytes() ([]byte, error) {
	return hex.DecodeString(c.UmaSigningPubKeyHex)
}

func (c *UmaConfig) UmaSigningPrivKeyBytes() ([]byte, error) {
	return hex.DecodeString(c.UmaSigningPrivKeyHex)
}

func (c *UmaConfig) NodeMasterSeedBytes() ([]byte, error) {
	return hex.DecodeString(c.RemoteSigningNodeMasterSeedHex)
}

/**
 * This object defines the configuration for the UMA server. First, initialize your client ID and
 * client secret. Those are available in your account at https://app.lightspark.com/api_config.
 *
 * Those variables can either be set through environmental variables:
*
 * export LIGHTSPARK_API_TOKEN_CLIENT_ID=<client_id>
 * export LIGHTSPARK_API_TOKEN_CLIENT_SECRET=<client_secret>
 *
 * Or simply edited inline.
 *
 * This example also assumes you already know your node UUID. Generally, an LNURL API would serve
 * many different usernames while maintaining some internal mapping from username to node UUID. For
 * simplicity, this example works with a single username and node UUID.
 *
 * export LIGHTSPARK_UMA_NODE_ID=0187c4d6-704b-f96b-0000-a2e8145bc1f9
*/

func NewConfig() UmaConfig {
	username := os.Getenv("LIGHTSPARK_UMA_RECEIVER_USER")
	if username == "" {
		username = "ls_test"
	}

	return UmaConfig{
		ApiClientID:     os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_ID"),
		ApiClientSecret: os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_SECRET"),
		NodeUUID:        os.Getenv("LIGHTSPARK_UMA_NODE_ID"),
		Username:        username,
		// Static UUID so that callback URLs are always the same.
		UserID:                         "4b41ae03-01b8-4974-8d26-26a35d28851b",
		UmaEncryptionPubKeyHex:         os.Getenv("LIGHTSPARK_UMA_ENCRYPTION_PUBKEY"),
		UmaEncryptionPrivKeyHex:        os.Getenv("LIGHTSPARK_UMA_ENCRYPTION_PRIVKEY"),
		UmaSigningPubKeyHex:            os.Getenv("LIGHTSPARK_UMA_SIGNING_PUBKEY"),
		UmaSigningPrivKeyHex:           os.Getenv("LIGHTSPARK_UMA_SIGNING_PRIVKEY"),
		RemoteSigningNodeMasterSeedHex: os.Getenv("LIGHTSPARK_UMA_REMOTE_SIGNING_NODE_MASTER_SEED"),
		OskNodeSigningKeyPassword:      os.Getenv("LIGHTSPARK_UMA_OSK_NODE_SIGNING_KEY_PASSWORD"),
		ClientBaseURL:                  fmt.Sprintf("https://%s/graphql/server/rc", os.Getenv("LIGHTSPARK_EXAMPLE_BASE_URL")),
		SenderVaspDomain:               os.Getenv("LIGHTSPARK_UMA_VASP_DOMAIN"),
	}
}
