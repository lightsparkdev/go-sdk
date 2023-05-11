// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type AccountToApiTokensConnection struct {

	// An object that holds pagination information about the objects in this connection.
	PageInfo PageInfo `json:"account_to_api_tokens_connection_page_info"`

	// The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"account_to_api_tokens_connection_count"`

	// The API tokens for the current page of this connection.
	Entities []ApiToken `json:"account_to_api_tokens_connection_entities"`
}

const (
	AccountToApiTokensConnectionFragment = `
fragment AccountToApiTokensConnectionFragment on AccountToApiTokensConnection {
    __typename
    account_to_api_tokens_connection_page_info: page_info {
        __typename
        page_info_has_next_page: has_next_page
        page_info_has_previous_page: has_previous_page
        page_info_start_cursor: start_cursor
        page_info_end_cursor: end_cursor
    }
    account_to_api_tokens_connection_count: count
    account_to_api_tokens_connection_entities: entities {
        id
    }
}
`
)
