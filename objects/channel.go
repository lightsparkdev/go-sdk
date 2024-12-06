
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "time"

// Channel This is an object representing a channel on the Lightning Network. You can retrieve this object to get detailed information on a specific Lightning Network channel.
type Channel struct {

    // Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
    Id string `json:"channel_id"`

    // CreatedAt The date and time when the entity was first created.
    CreatedAt time.Time `json:"channel_created_at"`

    // UpdatedAt The date and time when the entity was last updated.
    UpdatedAt time.Time `json:"channel_updated_at"`

    // FundingTransaction The transaction that funded the channel upon channel opening.
    FundingTransaction *types.EntityWrapper `json:"channel_funding_transaction"`

    // Capacity The total amount of funds in this channel, including the channel balance on the local node, the channel balance on the remote node and the on-chain fees to close the channel.
    Capacity *CurrencyAmount `json:"channel_capacity"`

    // LocalBalance The channel balance on the local node.
    LocalBalance *CurrencyAmount `json:"channel_local_balance"`

    // LocalUnsettledBalance The channel balance on the local node that is currently allocated to in-progress payments.
    LocalUnsettledBalance *CurrencyAmount `json:"channel_local_unsettled_balance"`

    // RemoteBalance The channel balance on the remote node.
    RemoteBalance *CurrencyAmount `json:"channel_remote_balance"`

    // RemoteUnsettledBalance The channel balance on the remote node that is currently allocated to in-progress payments.
    RemoteUnsettledBalance *CurrencyAmount `json:"channel_remote_unsettled_balance"`

    // UnsettledBalance The channel balance that is currently allocated to in-progress payments.
    UnsettledBalance *CurrencyAmount `json:"channel_unsettled_balance"`

    // TotalBalance The total balance in this channel, including the channel balance on both local and remote nodes.
    TotalBalance *CurrencyAmount `json:"channel_total_balance"`

    // Status The current status of this channel.
    Status *ChannelStatus `json:"channel_status"`

    // EstimatedForceClosureWaitMinutes The estimated time to wait for the channel's hash timelock contract to expire when force closing the channel. It is in unit of minutes.
    EstimatedForceClosureWaitMinutes *int64 `json:"channel_estimated_force_closure_wait_minutes"`

    // CommitFee The amount to be paid in fees for the current set of commitment transactions.
    CommitFee *CurrencyAmount `json:"channel_commit_fee"`

    // Fees The fees charged for routing payments through this channel.
    // Deprecated: Customer nodes do not route payments anymore.
    Fees *ChannelFees `json:"channel_fees"`

    // RemoteNode If known, the remote node of the channel.
    RemoteNode *types.EntityWrapper `json:"channel_remote_node"`

    // LocalNode The local Lightspark node of the channel.
    LocalNode types.EntityWrapper `json:"channel_local_node"`

    // ShortChannelId The unique identifier of the channel on Lightning Network, which is the location in the chain that the channel was confirmed. The format is <block-height>:<tx-index>:<tx-output>.
    ShortChannelId *string `json:"channel_short_channel_id"`

    // Typename The typename of the object
    Typename string `json:"__typename"`

}

const (
    ChannelFragment = `
fragment ChannelFragment on Channel {
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
}
`
)




// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj Channel) GetId() string {
    return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj Channel) GetCreatedAt() time.Time {
    return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj Channel) GetUpdatedAt() time.Time {
    return obj.UpdatedAt
}


    func (obj Channel) GetTypename() string {
        return obj.Typename
    }



    func (obj Channel) GetUptimePercentage(requester *requester.Requester, afterDate *time.Time, beforeDate *time.Time) (*int64, error) {
        query := `query FetchChannelUptimePercentage($entity_id: ID!, $after_date: DateTime, $before_date: DateTime) {
    entity(id: $entity_id) {
        ... on Channel {
            uptime_percentage(, after_date: $after_date, before_date: $before_date)
        }
    }
}`
        variables := map[string]interface{} {
        "entity_id": obj.Id,
"after_date": afterDate,
"before_date": beforeDate,

        }
      
        response, err := requester.ExecuteGraphql(query, variables, nil)
    	if err != nil {
	    	return nil, err
    	}

        output := response["entity"].(map[string]interface{})["uptime_percentage"]
        var result *int64
    	jsonString, err := json.Marshal(output)
	    json.Unmarshal(jsonString, &result)
    	return result, nil
    }

    func (obj Channel) GetTransactions(requester *requester.Requester, types *[]TransactionType, afterDate *time.Time, beforeDate *time.Time) (*ChannelToTransactionsConnection, error) {
        query := `query FetchChannelToTransactionsConnection($entity_id: ID!, $types: [TransactionType!], $after_date: DateTime, $before_date: DateTime) {
    entity(id: $entity_id) {
        ... on Channel {
            transactions(, types: $types, after_date: $after_date, before_date: $before_date) {
                __typename
                channel_to_transactions_connection_count: count
                channel_to_transactions_connection_average_fee: average_fee {
                    __typename
                    currency_amount_original_value: original_value
                    currency_amount_original_unit: original_unit
                    currency_amount_preferred_currency_unit: preferred_currency_unit
                    currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                    currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                }
                channel_to_transactions_connection_total_amount_transacted: total_amount_transacted {
                    __typename
                    currency_amount_original_value: original_value
                    currency_amount_original_unit: original_unit
                    currency_amount_preferred_currency_unit: preferred_currency_unit
                    currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                    currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                }
                channel_to_transactions_connection_total_fees: total_fees {
                    __typename
                    currency_amount_original_value: original_value
                    currency_amount_original_unit: original_unit
                    currency_amount_preferred_currency_unit: preferred_currency_unit
                    currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                    currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                }
            }
        }
    }
}`
        variables := map[string]interface{} {
        "entity_id": obj.Id,
"types": types,
"after_date": afterDate,
"before_date": beforeDate,

        }
      
        response, err := requester.ExecuteGraphql(query, variables, nil)
    	if err != nil {
	    	return nil, err
    	}

        output := response["entity"].(map[string]interface{})["transactions"].(map[string]interface{})
        var result *ChannelToTransactionsConnection
    	jsonString, err := json.Marshal(output)
	    json.Unmarshal(jsonString, &result)
    	return result, nil
    }



