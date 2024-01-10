
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type WalletToWithdrawalRequestsConnection struct {

    // Count The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
    Count int64 `json:"wallet_to_withdrawal_requests_connection_count"`

    // PageInfo An object that holds pagination information about the objects in this connection.
    PageInfo PageInfo `json:"wallet_to_withdrawal_requests_connection_page_info"`

    // Entities The withdrawal requests for the current page of this connection.
    Entities []WithdrawalRequest `json:"wallet_to_withdrawal_requests_connection_entities"`

    // Typename The typename of the object
    Typename string `json:"__typename"`

}

const (
    WalletToWithdrawalRequestsConnectionFragment = `
fragment WalletToWithdrawalRequestsConnectionFragment on WalletToWithdrawalRequestsConnection {
    __typename
    wallet_to_withdrawal_requests_connection_count: count
    wallet_to_withdrawal_requests_connection_page_info: page_info {
        __typename
        page_info_has_next_page: has_next_page
        page_info_has_previous_page: has_previous_page
        page_info_start_cursor: start_cursor
        page_info_end_cursor: end_cursor
    }
    wallet_to_withdrawal_requests_connection_entities: entities {
        id
    }
}
`
)




// GetCount The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
func (obj WalletToWithdrawalRequestsConnection) GetCount() int64 {
    return obj.Count
}

// GetPageInfo An object that holds pagination information about the objects in this connection.
func (obj WalletToWithdrawalRequestsConnection) GetPageInfo() PageInfo {
    return obj.PageInfo
}


    func (obj WalletToWithdrawalRequestsConnection) GetTypename() string {
        return obj.Typename
    }





