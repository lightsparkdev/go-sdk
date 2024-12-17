
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "time"


type ChannelSnapshot struct {

    // Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
    Id string `json:"channel_snapshot_id"`

    // CreatedAt The date and time when the entity was first created.
    CreatedAt time.Time `json:"channel_snapshot_created_at"`

    // UpdatedAt The date and time when the entity was last updated.
    UpdatedAt time.Time `json:"channel_snapshot_updated_at"`

    
    LocalBalance *CurrencyAmount `json:"channel_snapshot_local_balance"`

    
    LocalUnsettledBalance *CurrencyAmount `json:"channel_snapshot_local_unsettled_balance"`

    
    RemoteBalance *CurrencyAmount `json:"channel_snapshot_remote_balance"`

    
    RemoteUnsettledBalance *CurrencyAmount `json:"channel_snapshot_remote_unsettled_balance"`

    
    Status *string `json:"channel_snapshot_status"`

    
    Channel types.EntityWrapper `json:"channel_snapshot_channel"`

    
    // Deprecated: Use channel.local_channel_reserve instead.
    LocalChannelReserve *CurrencyAmount `json:"channel_snapshot_local_channel_reserve"`

    // Timestamp The timestamp that was used to query the snapshot of the channel
    Timestamp time.Time `json:"channel_snapshot_timestamp"`

    // Typename The typename of the object
    Typename string `json:"__typename"`

}

const (
    ChannelSnapshotFragment = `
fragment ChannelSnapshotFragment on ChannelSnapshot {
    __typename
    channel_snapshot_id: id
    channel_snapshot_created_at: created_at
    channel_snapshot_updated_at: updated_at
    channel_snapshot_local_balance: local_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    channel_snapshot_local_unsettled_balance: local_unsettled_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    channel_snapshot_remote_balance: remote_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    channel_snapshot_remote_unsettled_balance: remote_unsettled_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    channel_snapshot_status: status
    channel_snapshot_channel: channel {
        id
    }
    channel_snapshot_local_channel_reserve: local_channel_reserve {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    channel_snapshot_timestamp: timestamp
}
`
)




// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj ChannelSnapshot) GetId() string {
    return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj ChannelSnapshot) GetCreatedAt() time.Time {
    return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj ChannelSnapshot) GetUpdatedAt() time.Time {
    return obj.UpdatedAt
}


    func (obj ChannelSnapshot) GetTypename() string {
        return obj.Typename
    }





