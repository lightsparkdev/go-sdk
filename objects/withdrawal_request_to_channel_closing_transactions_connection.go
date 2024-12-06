
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type WithdrawalRequestToChannelClosingTransactionsConnection struct {

    // Count The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
    Count int64 `json:"withdrawal_request_to_channel_closing_transactions_connection_count"`

    // PageInfo An object that holds pagination information about the objects in this connection.
    PageInfo PageInfo `json:"withdrawal_request_to_channel_closing_transactions_connection_page_info"`

    // Entities The channel closing transactions for the current page of this connection.
    Entities []ChannelClosingTransaction `json:"withdrawal_request_to_channel_closing_transactions_connection_entities"`

    // Typename The typename of the object
    Typename string `json:"__typename"`

}

const (
    WithdrawalRequestToChannelClosingTransactionsConnectionFragment = `
fragment WithdrawalRequestToChannelClosingTransactionsConnectionFragment on WithdrawalRequestToChannelClosingTransactionsConnection {
    __typename
    withdrawal_request_to_channel_closing_transactions_connection_count: count
    withdrawal_request_to_channel_closing_transactions_connection_page_info: page_info {
        __typename
        page_info_has_next_page: has_next_page
        page_info_has_previous_page: has_previous_page
        page_info_start_cursor: start_cursor
        page_info_end_cursor: end_cursor
    }
    withdrawal_request_to_channel_closing_transactions_connection_entities: entities {
        id
    }
}
`
)




// GetCount The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
func (obj WithdrawalRequestToChannelClosingTransactionsConnection) GetCount() int64 {
    return obj.Count
}

// GetPageInfo An object that holds pagination information about the objects in this connection.
func (obj WithdrawalRequestToChannelClosingTransactionsConnection) GetPageInfo() PageInfo {
    return obj.PageInfo
}


    func (obj WithdrawalRequestToChannelClosingTransactionsConnection) GetTypename() string {
        return obj.Typename
    }





