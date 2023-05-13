// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "github.com/lightsparkdev/go-sdk/objects"

type CreateApiTokenOutput struct {
	// ApiToken is the API Token that has been created
	ApiToken *objects.ApiToken

	// ClientSecret should be used to authenticate against Lightspark API. This
	// secret is not stored and will never be available again after this. You
	// must keep this secret secure as it grants access to your account.
	ClientSecret string
}

const CREATE_API_TOKEN_MUTATION = `
mutation CreateApiToken(
    $name: String!
    $permissions: [Permission!]!
) {
    create_api_token(input: {
        name: $name
        permissions: $permissions
    }) {
        api_token {
            ...ApiTokenFragment
        }
        client_secret
    }
}

` + objects.ApiTokenFragment
