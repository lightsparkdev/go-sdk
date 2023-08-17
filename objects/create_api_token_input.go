// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type CreateApiTokenInput struct {

	// Name An arbitrary name that the user can choose to identify the API token in a list.
	Name string `json:"create_api_token_input_name"`

	// Permissions List of permissions to grant to the API token
	Permissions []Permission `json:"create_api_token_input_permissions"`
}
