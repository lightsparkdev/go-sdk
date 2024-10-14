package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

type UmaConfig struct {
	ApiClientID                    string
	ApiClientSecret                string
	CookieSecret                   string
	NodeUUID                       string
	Username                       string
	HashedUserPassword             string
	UserID                         string
	UmaEncryptionCertChain         string
	UmaEncryptionPubKeyHex         string
	UmaEncryptionPrivKeyHex        string
	UmaSigningCertChain            string
	UmaSigningPubKeyHex            string
	UmaSigningPrivKeyHex           string
	RemoteSigningNodeMasterSeedHex string
	OskNodeSigningKeyPassword      string
	ClientBaseURL                  *string
	OwnVaspDomain                  string
}

func (c *UmaConfig) UmaEncryptionPrivKeyBytes() ([]byte, error) {
	return hex.DecodeString(c.UmaEncryptionPrivKeyHex)
}

func (c *UmaConfig) UmaSigningPrivKeyBytes() ([]byte, error) {
	return hex.DecodeString(c.UmaSigningPrivKeyHex)
}

func (c *UmaConfig) NodeMasterSeedBytes() ([]byte, error) {
	return hex.DecodeString(c.RemoteSigningNodeMasterSeedHex)
}

func (c *UmaConfig) GetVaspDomain(context *gin.Context) string {
	envVaspDomain := c.OwnVaspDomain
	if envVaspDomain != "" {
		return envVaspDomain
	}
	requestHost := context.Request.Host
	requestHostWithoutPort := strings.Split(requestHost, ":")[0]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	if port != "80" && port != "443" {
		return fmt.Sprintf("%s:%s", requestHostWithoutPort, port)
	}
	return requestHostWithoutPort
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
	password := os.Getenv("LIGHTSPARK_UMA_RECEIVER_PASSWORD")
	if password == "" {
		password = "pa$$w0rd" // Super secure, right?
	}
	// Hash the password using sha256 for storage
	hashedPassword := sha256.New()
	hashedPassword.Write([]byte(password))
	hashedPasswordStr := hex.EncodeToString(hashedPassword.Sum(nil))

	baseUrlEnv := os.Getenv("LIGHTSPARK_EXAMPLE_BASE_URL")
	baseUrlStr := fmt.Sprintf("https://%s/graphql/server/rc", baseUrlEnv)
	baseUrl := &baseUrlStr
	if baseUrlEnv == "" {
		baseUrl = nil
	}

	return UmaConfig{
		ApiClientID:        os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_ID"),
		ApiClientSecret:    os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_SECRET"),
		NodeUUID:           os.Getenv("LIGHTSPARK_UMA_NODE_ID"),
		CookieSecret:       os.Getenv("LIGHTSPARK_UMA_COOKIE_SECRET"),
		Username:           username,
		HashedUserPassword: hashedPasswordStr,
		// Static UUID so that callback URLs are always the same.
		UserID:                         "4b41ae03-01b8-4974-8d26-26a35d28851b",
		UmaEncryptionCertChain:         os.Getenv("LIGHTSPARK_UMA_ENCRYPTION_CERT_CHAIN"),
		UmaEncryptionPubKeyHex:         os.Getenv("LIGHTSPARK_UMA_ENCRYPTION_PUBKEY"),
		UmaEncryptionPrivKeyHex:        os.Getenv("LIGHTSPARK_UMA_ENCRYPTION_PRIVKEY"),
		UmaSigningCertChain:            os.Getenv("LIGHTSPARK_UMA_SIGNING_CERT_CHAIN"),
		UmaSigningPubKeyHex:            os.Getenv("LIGHTSPARK_UMA_SIGNING_PUBKEY"),
		UmaSigningPrivKeyHex:           os.Getenv("LIGHTSPARK_UMA_SIGNING_PRIVKEY"),
		RemoteSigningNodeMasterSeedHex: os.Getenv("LIGHTSPARK_UMA_REMOTE_SIGNING_NODE_MASTER_SEED"),
		OskNodeSigningKeyPassword:      os.Getenv("LIGHTSPARK_UMA_OSK_NODE_SIGNING_KEY_PASSWORD"),
		ClientBaseURL:                  baseUrl,
		OwnVaspDomain:                  os.Getenv("LIGHTSPARK_UMA_VASP_DOMAIN"),
	}
}
