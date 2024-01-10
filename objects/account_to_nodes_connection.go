
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects



// AccountToNodesConnection A connection between an account and the nodes it manages.
type AccountToNodesConnection struct {

    // Count The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
    Count int64 `json:"account_to_nodes_connection_count"`

    // PageInfo An object that holds pagination information about the objects in this connection.
    PageInfo PageInfo `json:"account_to_nodes_connection_page_info"`

    // Entities The nodes for the current page of this connection.
    Entities []LightsparkNode `json:"account_to_nodes_connection_entities"`

    // Typename The typename of the object
    Typename string `json:"__typename"`

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
    account_to_nodes_connection_entities: entities {
        id
    }
}
`
)




// GetCount The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
func (obj AccountToNodesConnection) GetCount() int64 {
    return obj.Count
}

// GetPageInfo An object that holds pagination information about the objects in this connection.
func (obj AccountToNodesConnection) GetPageInfo() PageInfo {
    return obj.PageInfo
}


    func (obj AccountToNodesConnection) GetTypename() string {
        return obj.Typename
    }





type AccountToNodesConnectionJSON struct {

    // Count The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
    Count int64 `json:"account_to_nodes_connection_count"`

    // PageInfo An object that holds pagination information about the objects in this connection.
    PageInfo PageInfo `json:"account_to_nodes_connection_page_info"`

    // Entities The nodes for the current page of this connection.
    Entities []map[string]interface{} `json:"account_to_nodes_connection_entities"`

    // Typename The typename of the object
    Typename string `json:"__typename"`

}


func (data *AccountToNodesConnection) UnmarshalJSON(dataBytes []byte) error {
    var temp AccountToNodesConnectionJSON
	if err := json.Unmarshal(dataBytes, &temp); err != nil {
		return err
	}

	
    data.Count = temp.Count


    data.PageInfo = temp.PageInfo


    if temp.Entities != nil {
        var entities []LightsparkNode
        for _, json := range temp.Entities {
            entity, err := LightsparkNodeUnmarshal(json)
            if err != nil {
                return err
            }
            entities = append(entities, entity)
        }
        data.Entities = entities
    }


    data.Typename = temp.Typename


    return nil
}


