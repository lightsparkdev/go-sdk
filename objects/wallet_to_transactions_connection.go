// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "encoding/json"

type WalletToTransactionsConnection struct {

	// Count The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"wallet_to_transactions_connection_count"`

	// PageInfo An object that holds pagination information about the objects in this connection.
	PageInfo PageInfo `json:"wallet_to_transactions_connection_page_info"`

	// Entities The transactions for the current page of this connection.
	Entities []Transaction `json:"wallet_to_transactions_connection_entities"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
}

const (
	WalletToTransactionsConnectionFragment = `
fragment WalletToTransactionsConnectionFragment on WalletToTransactionsConnection {
    __typename
    wallet_to_transactions_connection_count: count
    wallet_to_transactions_connection_page_info: page_info {
        __typename
        page_info_has_next_page: has_next_page
        page_info_has_previous_page: has_previous_page
        page_info_start_cursor: start_cursor
        page_info_end_cursor: end_cursor
    }
    wallet_to_transactions_connection_entities: entities {
        id
    }
}
`
)

// GetCount The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
func (obj WalletToTransactionsConnection) GetCount() int64 {
	return obj.Count
}

// GetPageInfo An object that holds pagination information about the objects in this connection.
func (obj WalletToTransactionsConnection) GetPageInfo() PageInfo {
	return obj.PageInfo
}

func (obj WalletToTransactionsConnection) GetTypename() string {
	return obj.Typename
}

type WalletToTransactionsConnectionJSON struct {

	// Count The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"wallet_to_transactions_connection_count"`

	// PageInfo An object that holds pagination information about the objects in this connection.
	PageInfo PageInfo `json:"wallet_to_transactions_connection_page_info"`

	// Entities The transactions for the current page of this connection.
	Entities []map[string]interface{} `json:"wallet_to_transactions_connection_entities"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
}

func (data *WalletToTransactionsConnection) UnmarshalJSON(dataBytes []byte) error {
	var temp WalletToTransactionsConnectionJSON
	if err := json.Unmarshal(dataBytes, &temp); err != nil {
		return err
	}

	data.Count = temp.Count

	data.PageInfo = temp.PageInfo

	if temp.Entities != nil {
		var entities []Transaction
		for _, json := range temp.Entities {
			entity, err := TransactionUnmarshal(json)
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
