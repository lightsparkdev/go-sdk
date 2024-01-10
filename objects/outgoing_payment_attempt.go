
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "time"

// OutgoingPaymentAttempt This object represents an attempted Lightning Network payment sent from a Lightspark Node. You can retrieve this object to receive payment related information about any payment attempt sent from your Lightspark Node on the Lightning Network, including any potential reasons the payment may have failed.
type OutgoingPaymentAttempt struct {

    // Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
    Id string `json:"outgoing_payment_attempt_id"`

    // CreatedAt The date and time when the entity was first created.
    CreatedAt time.Time `json:"outgoing_payment_attempt_created_at"`

    // UpdatedAt The date and time when the entity was last updated.
    UpdatedAt time.Time `json:"outgoing_payment_attempt_updated_at"`

    // Status The status of an outgoing payment attempt.
    Status OutgoingPaymentAttemptStatus `json:"outgoing_payment_attempt_status"`

    // FailureCode If the payment attempt failed, then this contains the Bolt #4 failure code.
    FailureCode *HtlcAttemptFailureCode `json:"outgoing_payment_attempt_failure_code"`

    // FailureSourceIndex If the payment attempt failed, then this contains the index of the hop at which the problem occurred.
    FailureSourceIndex *int64 `json:"outgoing_payment_attempt_failure_source_index"`

    // AttemptedAt The date and time when the attempt was initiated.
    AttemptedAt time.Time `json:"outgoing_payment_attempt_attempted_at"`

    // ResolvedAt The time the outgoing payment attempt failed or succeeded.
    ResolvedAt *time.Time `json:"outgoing_payment_attempt_resolved_at"`

    // Amount The total amount of funds required to complete a payment over this route. This value includes the cumulative fees for each hop. As a result, the attempt extended to the first-hop in the route will need to have at least this much value, otherwise the route will fail at an intermediate node due to an insufficient amount.
    Amount *CurrencyAmount `json:"outgoing_payment_attempt_amount"`

    // Fees The sum of the fees paid at each hop within the route of this attempt. In the case of a one-hop payment, this value will be zero as we don't need to pay a fee to ourselves.
    Fees *CurrencyAmount `json:"outgoing_payment_attempt_fees"`

    // OutgoingPayment The outgoing payment for this attempt.
    OutgoingPayment types.EntityWrapper `json:"outgoing_payment_attempt_outgoing_payment"`

    // ChannelSnapshot The channel snapshot at the time the outgoing payment attempt was made.
    ChannelSnapshot *types.EntityWrapper `json:"outgoing_payment_attempt_channel_snapshot"`

    // Typename The typename of the object
    Typename string `json:"__typename"`

}

const (
    OutgoingPaymentAttemptFragment = `
fragment OutgoingPaymentAttemptFragment on OutgoingPaymentAttempt {
    __typename
    outgoing_payment_attempt_id: id
    outgoing_payment_attempt_created_at: created_at
    outgoing_payment_attempt_updated_at: updated_at
    outgoing_payment_attempt_status: status
    outgoing_payment_attempt_failure_code: failure_code
    outgoing_payment_attempt_failure_source_index: failure_source_index
    outgoing_payment_attempt_attempted_at: attempted_at
    outgoing_payment_attempt_resolved_at: resolved_at
    outgoing_payment_attempt_amount: amount {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    outgoing_payment_attempt_fees: fees {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    outgoing_payment_attempt_outgoing_payment: outgoing_payment {
        id
    }
    outgoing_payment_attempt_channel_snapshot: channel_snapshot {
        id
    }
}
`
)




// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj OutgoingPaymentAttempt) GetId() string {
    return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj OutgoingPaymentAttempt) GetCreatedAt() time.Time {
    return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj OutgoingPaymentAttempt) GetUpdatedAt() time.Time {
    return obj.UpdatedAt
}


    func (obj OutgoingPaymentAttempt) GetTypename() string {
        return obj.Typename
    }



    func (obj OutgoingPaymentAttempt) GetHops(requester *requester.Requester, first *int64, after *string) (*OutgoingPaymentAttemptToHopsConnection, error) {
        query := `query FetchOutgoingPaymentAttemptToHopsConnection($entity_id: ID!, $first: Int, $after: String) {
    entity(id: $entity_id) {
        ... on OutgoingPaymentAttempt {
            hops(, first: $first, after: $after) {
                __typename
                outgoing_payment_attempt_to_hops_connection_count: count
                outgoing_payment_attempt_to_hops_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                outgoing_payment_attempt_to_hops_connection_entities: entities {
                    __typename
                    hop_id: id
                    hop_created_at: created_at
                    hop_updated_at: updated_at
                    hop_destination: destination {
                        id
                    }
                    hop_index: index
                    hop_public_key: public_key
                    hop_amount_to_forward: amount_to_forward {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    hop_fee: fee {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    hop_expiry_block_height: expiry_block_height
                }
            }
        }
    }
}`
        variables := map[string]interface{} {
        "entity_id": obj.Id,
"first": first,
"after": after,

        }
      
        response, err := requester.ExecuteGraphql(query, variables, nil)
    	if err != nil {
	    	return nil, err
    	}

        output := response["entity"].(map[string]interface{})["hops"].(map[string]interface{})
        var result *OutgoingPaymentAttemptToHopsConnection
    	jsonString, err := json.Marshal(output)
	    json.Unmarshal(jsonString, &result)
    	return result, nil
    }



