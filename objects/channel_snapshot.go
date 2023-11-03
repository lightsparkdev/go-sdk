// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"time"

	"github.com/lightsparkdev/go-sdk/types"
)

type ChannelSnapshot struct {
	Channel types.EntityWrapper `json:"channel_snapshot_channel"`

	Timestamp time.Time `json:"channel_snapshot_timestamp"`

	LocalBalance *CurrencyAmount `json:"channel_snapshot_local_balance"`

	LocalUnsettledBalance *CurrencyAmount `json:"channel_snapshot_local_unsettled_balance"`

	LocalChannelReserve *CurrencyAmount `json:"channel_snapshot_local_channel_reserve"`

	RemoteBalance *CurrencyAmount `json:"channel_snapshot_remote_balance"`

	RemoteUnsettledBalance *CurrencyAmount `json:"channel_snapshot_remote_unsettled_balance"`
}

const (
	ChannelSnapshotFragment = `
fragment ChannelSnapshotFragment on ChannelSnapshot {
    __typename
    channel_snapshot_channel: channel {
        id
    }
    channel_snapshot_timestamp: timestamp
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
    channel_snapshot_local_channel_reserve: local_channel_reserve {
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
}
`
)
