// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type CreateApiTokenOutput struct {

	// ApiToken The API Token that has been created.
	ApiToken ApiToken `json:"create_api_token_output_api_token"`

	// ClientSecret The secret that should be used to authenticate against our API.
	// This secret is not stored and will never be available again after this. You must keep this secret secure as it grants access to your account.
	ClientSecret string `json:"create_api_token_output_client_secret"`
}

const (
	CreateApiTokenOutputFragment = `
fragment CreateApiTokenOutputFragment on CreateApiTokenOutput {
    __typename
    create_api_token_output_api_token: api_token {
        __typename
        api_token_id: id
        api_token_created_at: created_at
        api_token_updated_at: updated_at
        api_token_client_id: client_id
        api_token_name: name
        api_token_permissions: permissions
    }
    create_api_token_output_client_secret: client_secret
}
`
)
