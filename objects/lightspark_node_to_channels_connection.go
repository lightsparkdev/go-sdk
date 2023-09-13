// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type LightsparkNodeToChannelsConnection struct {

	// Count The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"lightspark_node_to_channels_connection_count"`

	// PageInfo An object that holds pagination information about the objects in this connection.
	PageInfo PageInfo `json:"lightspark_node_to_channels_connection_page_info"`

	// Entities The channels for the current page of this connection.
	Entities []Channel `json:"lightspark_node_to_channels_connection_entities"`
}

const (
	LightsparkNodeToChannelsConnectionFragment = `
fragment LightsparkNodeToChannelsConnectionFragment on LightsparkNodeToChannelsConnection {
    __typename
    lightspark_node_to_channels_connection_count: count
    lightspark_node_to_channels_connection_page_info: page_info {
        __typename
        page_info_has_next_page: has_next_page
        page_info_has_previous_page: has_previous_page
        page_info_start_cursor: start_cursor
        page_info_end_cursor: end_cursor
    }
    lightspark_node_to_channels_connection_entities: entities {
        id
    }
}
`
)

// GetCount The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
func (obj LightsparkNodeToChannelsConnection) GetCount() int64 {
	return obj.Count
}

// GetPageInfo An object that holds pagination information about the objects in this connection.
func (obj LightsparkNodeToChannelsConnection) GetPageInfo() PageInfo {
	return obj.PageInfo
}
