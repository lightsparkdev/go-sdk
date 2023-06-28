package main

import "os"

type LnurlConfig struct {
	ApiClientID     string
	ApiClientSecret string
	NodeUUID        string
	Username        string
	UserID          string
}

/**
 * This object defines the configuration for the LNURL server. First, initialize your client ID and
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
 * export LIGHTSPARK_LNURL_NODE_UUID=0187c4d6-704b-f96b-0000-a2e8145bc1f9
*/

func NewConfig() LnurlConfig {
	username := os.Getenv("LIGHTSPARK_LNURL_USERNAME")
	if username == "" {
		username = "ls_test"
	}

	return LnurlConfig{
		ApiClientID:     os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_ID"),
		ApiClientSecret: os.Getenv("LIGHTSPARK_API_TOKEN_CLIENT_SECRET"),
		NodeUUID:        os.Getenv("LIGHTSPARK_LNURL_NODE_UUID"),
		Username:        username,
		// Static UUID so that callback URLs are always the same.
		UserID: "4b41ae03-01b8-4974-8d26-26a35d28851b",
	}
}
