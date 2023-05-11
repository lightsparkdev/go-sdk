// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type AccountToChannelsConnection struct {

	// The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	Count int64 `json:"account_to_channels_connection_count"`

	// The channels for the current page of this connection.
	Entities []Channel `json:"account_to_channels_connection_entities"`
}

const (
	AccountToChannelsConnectionFragment = `
fragment AccountToChannelsConnectionFragment on AccountToChannelsConnection {
    __typename
    account_to_channels_connection_count: count
    account_to_channels_connection_entities: entities {
        id
    }
}
`
)
