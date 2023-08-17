// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "time"

// This is an object representing a Lightspark API token, that can be used to authenticate this account when making API calls or using our SDKs. See the “Authentication” section of our API docs for more details on its usage.
type ApiToken struct {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"api_token_id"`

	// The date and time when the entity was first created.
	CreatedAt time.Time `json:"api_token_created_at"`

	// The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"api_token_updated_at"`

	// An opaque identifier that should be used as a client_id (or username) in the HTTP Basic Authentication scheme when issuing requests against the Lightspark API.
	ClientId string `json:"api_token_client_id"`

	// An arbitrary name chosen by the creator of the token to help identify the token in the list of tokens that have been created for the account.
	Name string `json:"api_token_name"`

	// A list of permissions granted to the token.
	Permissions []Permission `json:"api_token_permissions"`
}

const (
	ApiTokenFragment = `
fragment ApiTokenFragment on ApiToken {
    __typename
    api_token_id: id
    api_token_created_at: created_at
    api_token_updated_at: updated_at
    api_token_client_id: client_id
    api_token_name: name
    api_token_permissions: permissions
}
`
)

// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj ApiToken) GetId() string {
	return obj.Id
}

// The date and time when the entity was first created.
func (obj ApiToken) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// The date and time when the entity was last updated.
func (obj ApiToken) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}
