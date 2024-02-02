// Copyright ©, 2024-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "github.com/lightsparkdev/go-sdk/types"

type LightsparkNodeToRipcordUpdatesConnection struct {
	FromDate types.Date `json:"lightspark_node_to_ripcord_updates_connection_from_date"`

	ToDate types.Date `json:"lightspark_node_to_ripcord_updates_connection_to_date"`

	// Entities The daily liquidity forecasts for the current page of this connection.
	Entities []RipcordUpdate `json:"lightspark_node_to_ripcord_updates_connection_entities"`
}

const (
	LightsparkNodeToRipcordUpdatesConnectionFragment = `
fragment LightsparkNodeToRipcordUpdatesConnectionFragment on LightsparkNodeToRipcordUpdatesConnection {
	__typename
	ripcord_update_commitment_number: commitment_number
	ripcord_update_ripcord_update_status: ripcord_update_status
	ripcord_update_data: data
	ripcord_update_channel: channel {
		__typename
		channel_id: id
		channel_created_at: created_at
		channel_updated_at: updated_at
		channel_funding_transaction: funding_transaction {
			id
		}
		channel_capacity: capacity {
			__typename
			currency_amount_original_value: original_value
			currency_amount_original_unit: original_unit
			currency_amount_preferred_currency_unit: preferred_currency_unit
			currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
			currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
		}
		channel_local_balance: local_balance {
			__typename
			currency_amount_original_value: original_value
			currency_amount_original_unit: original_unit
			currency_amount_preferred_currency_unit: preferred_currency_unit
			currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
			currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
		}
		channel_local_unsettled_balance: local_unsettled_balance {
			__typename
			currency_amount_original_value: original_value
			currency_amount_original_unit: original_unit
			currency_amount_preferred_currency_unit: preferred_currency_unit
			currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
			currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
		}
		channel_remote_balance: remote_balance {
			__typename
			currency_amount_original_value: original_value
			currency_amount_original_unit: original_unit
			currency_amount_preferred_currency_unit: preferred_currency_unit
			currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
			currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
		}
		channel_remote_unsettled_balance: remote_unsettled_balance {
			__typename
			currency_amount_original_value: original_value
			currency_amount_original_unit: original_unit
			currency_amount_preferred_currency_unit: preferred_currency_unit
			currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
			currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
		}
		channel_unsettled_balance: unsettled_balance {
			__typename
			currency_amount_original_value: original_value
			currency_amount_original_unit: original_unit
			currency_amount_preferred_currency_unit: preferred_currency_unit
			currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
			currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
		}
		channel_total_balance: total_balance {
			__typename
			currency_amount_original_value: original_value
			currency_amount_original_unit: original_unit
			currency_amount_preferred_currency_unit: preferred_currency_unit
			currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
			currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
		}
		channel_status: status
		channel_estimated_force_closure_wait_minutes: estimated_force_closure_wait_minutes
		channel_commit_fee: commit_fee {
			__typename
			currency_amount_original_value: original_value
			currency_amount_original_unit: original_unit
			currency_amount_preferred_currency_unit: preferred_currency_unit
			currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
			currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
		}
		channel_fees: fees {
			__typename
			channel_fees_base_fee: base_fee {
				__typename
				currency_amount_original_value: original_value
				currency_amount_original_unit: original_unit
				currency_amount_preferred_currency_unit: preferred_currency_unit
				currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
				currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
			}
			channel_fees_fee_rate_per_mil: fee_rate_per_mil
		}
		channel_remote_node: remote_node {
			id
		}
		channel_local_node: local_node {
			id
		}
		channel_short_channel_id: short_channel_id
		channel_channel_point: channel_point
	}
}
`
)
