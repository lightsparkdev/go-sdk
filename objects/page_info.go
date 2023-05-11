// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type PageInfo struct {
	HasNextPage *bool `json:"page_info_has_next_page"`

	HasPreviousPage *bool `json:"page_info_has_previous_page"`

	StartCursor *string `json:"page_info_start_cursor"`

	EndCursor *string `json:"page_info_end_cursor"`
}

const (
	PageInfoFragment = `
fragment PageInfoFragment on PageInfo {
    __typename
    page_info_has_next_page: has_next_page
    page_info_has_previous_page: has_previous_page
    page_info_start_cursor: start_cursor
    page_info_end_cursor: end_cursor
}
`
)
