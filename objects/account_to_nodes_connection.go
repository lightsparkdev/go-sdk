// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// A connection between an account and the nodes it manages.
type AccountToNodesConnection struct {

	// The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"account_to_nodes_connection_count"`

	// An object that holds pagination information about the objects in this connection.
	PageInfo PageInfo `json:"account_to_nodes_connection_page_info"`

	// The main purpose for the selected set of nodes. It is automatically determined from the nodes that are selected in this connection and is used for optimization purposes, as well as to determine the variation of the UI that should be presented to the user.
	Purpose *LightsparkNodePurpose `json:"account_to_nodes_connection_purpose"`

	// The nodes for the current page of this connection.
	Entities []LightsparkNode `json:"account_to_nodes_connection_entities"`
}

const (
	AccountToNodesConnectionFragment = `
fragment AccountToNodesConnectionFragment on AccountToNodesConnection {
    __typename
    account_to_nodes_connection_count: count
    account_to_nodes_connection_page_info: page_info {
        __typename
        page_info_has_next_page: has_next_page
        page_info_has_previous_page: has_previous_page
        page_info_start_cursor: start_cursor
        page_info_end_cursor: end_cursor
    }
    account_to_nodes_connection_purpose: purpose
    account_to_nodes_connection_entities: entities {
        id
    }
}
`
)

// The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
func (obj AccountToNodesConnection) GetCount() int64 {
	return obj.Count
}

// An object that holds pagination information about the objects in this connection.
func (obj AccountToNodesConnection) GetPageInfo() PageInfo {
	return obj.PageInfo
}
