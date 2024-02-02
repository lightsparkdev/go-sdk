// Copyright ©, 2024-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// RipcordUpdate This is an object representing the state of a channel which can be combined into broadcastable transactions for customers to claim their funds.
type RipcordUpdate struct {

	// CommitmentNumber The decreasing number that represents the state of a channel.
	CommitmentNumber *int `json:"ripcord_update_commitment_number"`

	// RipcordUpdateStatus The status of the ripcord update that signifies db creation and s3 sending.
	RipcordUpdateStatus *string `json:"ripcord_update_ripcord_update_status"`

	// Data The json blob that has all of the relevant state information.
	Data *string `json:"ripcord_update_data"`

	// Channel The ent channel that this ripcord update is regarding.
	Channel *Channel `json:"ripcord_update_channel"`
}

const (
	RipcordUpdateFragment = `
fragment RipcordUpdateFragment on RipcordUpdate {
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
