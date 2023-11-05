// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"time"

	"github.com/lightsparkdev/go-sdk/requester"
	"github.com/lightsparkdev/go-sdk/types"
)

// LightsparkNodeWithRemoteSigning This is a Lightspark node with remote signing.
type LightsparkNodeWithRemoteSigning struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"lightspark_node_with_remote_signing_id"`

	// CreatedAt The date and time when the entity was first created.
	CreatedAt time.Time `json:"lightspark_node_with_remote_signing_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"lightspark_node_with_remote_signing_updated_at"`

	// Alias A name that identifies the node. It has no importance in terms of operating the node, it is just a way to identify and search for commercial services or popular nodes. This alias can be changed at any time by the node operator.
	Alias *string `json:"lightspark_node_with_remote_signing_alias"`

	// BitcoinNetwork The Bitcoin Network this node is deployed in.
	BitcoinNetwork BitcoinNetwork `json:"lightspark_node_with_remote_signing_bitcoin_network"`

	// Color A hexadecimal string that describes a color. For example "#000000" is black, "#FFFFFF" is white. It has no importance in terms of operating the node, it is just a way to visually differentiate nodes. That color can be changed at any time by the node operator.
	Color *string `json:"lightspark_node_with_remote_signing_color"`

	// Conductivity A summary metric used to capture how well positioned a node is to send, receive, or route transactions efficiently. Maximizing a node's conductivity helps a node’s transactions to be capital efficient. The value is an integer ranging between 0 and 10 (bounds included).
	Conductivity *int64 `json:"lightspark_node_with_remote_signing_conductivity"`

	// DisplayName The name of this node in the network. It will be the most human-readable option possible, depending on the data available for this node.
	DisplayName string `json:"lightspark_node_with_remote_signing_display_name"`

	// PublicKey The public key of this node. It acts as a unique identifier of this node in the Lightning Network.
	PublicKey *string `json:"lightspark_node_with_remote_signing_public_key"`

	// Owner The owner of this LightsparkNode.
	Owner types.EntityWrapper `json:"lightspark_node_with_remote_signing_owner"`

	// Status The current status of this node.
	Status *LightsparkNodeStatus `json:"lightspark_node_with_remote_signing_status"`

	// TotalBalance The sum of the balance on the Bitcoin Network, channel balances, and commit fees on this node.
	// Deprecated: Use `balances` instead.
	TotalBalance *CurrencyAmount `json:"lightspark_node_with_remote_signing_total_balance"`

	// TotalLocalBalance The total sum of the channel balances (online and offline) on this node.
	// Deprecated: Use `balances` instead.
	TotalLocalBalance *CurrencyAmount `json:"lightspark_node_with_remote_signing_total_local_balance"`

	// LocalBalance The sum of the channel balances (online only) that are available to send on this node.
	// Deprecated: Use `balances` instead.
	LocalBalance *CurrencyAmount `json:"lightspark_node_with_remote_signing_local_balance"`

	// RemoteBalance The sum of the channel balances that are available to receive on this node.
	// Deprecated: Use `balances` instead.
	RemoteBalance *CurrencyAmount `json:"lightspark_node_with_remote_signing_remote_balance"`

	// BlockchainBalance The details of the balance of this node on the Bitcoin Network.
	// Deprecated: Use `balances` instead.
	BlockchainBalance *BlockchainBalance `json:"lightspark_node_with_remote_signing_blockchain_balance"`

	// UmaPrescreeningUtxos The utxos of the channels that are connected to this node. This is used in uma flow for pre-screening.
	UmaPrescreeningUtxos []string `json:"lightspark_node_with_remote_signing_uma_prescreening_utxos"`

	// Balances The balances that describe the funds in this node.
	Balances *Balances `json:"lightspark_node_with_remote_signing_balances"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
}

const (
	LightsparkNodeWithRemoteSigningFragment = `
fragment LightsparkNodeWithRemoteSigningFragment on LightsparkNodeWithRemoteSigning {
    __typename
    lightspark_node_with_remote_signing_id: id
    lightspark_node_with_remote_signing_created_at: created_at
    lightspark_node_with_remote_signing_updated_at: updated_at
    lightspark_node_with_remote_signing_alias: alias
    lightspark_node_with_remote_signing_bitcoin_network: bitcoin_network
    lightspark_node_with_remote_signing_color: color
    lightspark_node_with_remote_signing_conductivity: conductivity
    lightspark_node_with_remote_signing_display_name: display_name
    lightspark_node_with_remote_signing_public_key: public_key
    lightspark_node_with_remote_signing_owner: owner {
        id
    }
    lightspark_node_with_remote_signing_status: status
    lightspark_node_with_remote_signing_total_balance: total_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    lightspark_node_with_remote_signing_total_local_balance: total_local_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    lightspark_node_with_remote_signing_local_balance: local_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    lightspark_node_with_remote_signing_remote_balance: remote_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    lightspark_node_with_remote_signing_blockchain_balance: blockchain_balance {
        __typename
        blockchain_balance_total_balance: total_balance {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
        blockchain_balance_confirmed_balance: confirmed_balance {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
        blockchain_balance_unconfirmed_balance: unconfirmed_balance {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
        blockchain_balance_locked_balance: locked_balance {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
        blockchain_balance_required_reserve: required_reserve {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
        blockchain_balance_available_balance: available_balance {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
    }
    lightspark_node_with_remote_signing_uma_prescreening_utxos: uma_prescreening_utxos
    lightspark_node_with_remote_signing_balances: balances {
        __typename
        balances_owned_balance: owned_balance {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
        balances_available_to_send_balance: available_to_send_balance {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
        balances_available_to_withdraw_balance: available_to_withdraw_balance {
            __typename
            currency_amount_original_value: original_value
            currency_amount_original_unit: original_unit
            currency_amount_preferred_currency_unit: preferred_currency_unit
            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
        }
    }
}
`
)

// GetOwnerId The owner of this LightsparkNode.
func (obj LightsparkNodeWithRemoteSigning) GetOwnerId() types.EntityWrapper {
	return obj.Owner
}

// GetStatus The current status of this node.
func (obj LightsparkNodeWithRemoteSigning) GetStatus() *LightsparkNodeStatus {
	return obj.Status
}

// GetTotalBalance The sum of the balance on the Bitcoin Network, channel balances, and commit fees on this node.
// Deprecated: Use `balances` instead.
func (obj LightsparkNodeWithRemoteSigning) GetTotalBalance() *CurrencyAmount {
	return obj.TotalBalance
}

// GetTotalLocalBalance The total sum of the channel balances (online and offline) on this node.
// Deprecated: Use `balances` instead.
func (obj LightsparkNodeWithRemoteSigning) GetTotalLocalBalance() *CurrencyAmount {
	return obj.TotalLocalBalance
}

// GetLocalBalance The sum of the channel balances (online only) that are available to send on this node.
// Deprecated: Use `balances` instead.
func (obj LightsparkNodeWithRemoteSigning) GetLocalBalance() *CurrencyAmount {
	return obj.LocalBalance
}

// GetRemoteBalance The sum of the channel balances that are available to receive on this node.
// Deprecated: Use `balances` instead.
func (obj LightsparkNodeWithRemoteSigning) GetRemoteBalance() *CurrencyAmount {
	return obj.RemoteBalance
}

// GetBlockchainBalance The details of the balance of this node on the Bitcoin Network.
// Deprecated: Use `balances` instead.
func (obj LightsparkNodeWithRemoteSigning) GetBlockchainBalance() *BlockchainBalance {
	return obj.BlockchainBalance
}

// GetUmaPrescreeningUtxos The utxos of the channels that are connected to this node. This is used in uma flow for pre-screening.
func (obj LightsparkNodeWithRemoteSigning) GetUmaPrescreeningUtxos() []string {
	return obj.UmaPrescreeningUtxos
}

// GetBalances The balances that describe the funds in this node.
func (obj LightsparkNodeWithRemoteSigning) GetBalances() *Balances {
	return obj.Balances
}

// GetAlias A name that identifies the node. It has no importance in terms of operating the node, it is just a way to identify and search for commercial services or popular nodes. This alias can be changed at any time by the node operator.
func (obj LightsparkNodeWithRemoteSigning) GetAlias() *string {
	return obj.Alias
}

// GetBitcoinNetwork The Bitcoin Network this node is deployed in.
func (obj LightsparkNodeWithRemoteSigning) GetBitcoinNetwork() BitcoinNetwork {
	return obj.BitcoinNetwork
}

// GetColor A hexadecimal string that describes a color. For example "#000000" is black, "#FFFFFF" is white. It has no importance in terms of operating the node, it is just a way to visually differentiate nodes. That color can be changed at any time by the node operator.
func (obj LightsparkNodeWithRemoteSigning) GetColor() *string {
	return obj.Color
}

// GetConductivity A summary metric used to capture how well positioned a node is to send, receive, or route transactions efficiently. Maximizing a node's conductivity helps a node’s transactions to be capital efficient. The value is an integer ranging between 0 and 10 (bounds included).
func (obj LightsparkNodeWithRemoteSigning) GetConductivity() *int64 {
	return obj.Conductivity
}

// GetDisplayName The name of this node in the network. It will be the most human-readable option possible, depending on the data available for this node.
func (obj LightsparkNodeWithRemoteSigning) GetDisplayName() string {
	return obj.DisplayName
}

// GetPublicKey The public key of this node. It acts as a unique identifier of this node in the Lightning Network.
func (obj LightsparkNodeWithRemoteSigning) GetPublicKey() *string {
	return obj.PublicKey
}

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj LightsparkNodeWithRemoteSigning) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj LightsparkNodeWithRemoteSigning) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj LightsparkNodeWithRemoteSigning) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj LightsparkNodeWithRemoteSigning) GetTypename() string {
	return obj.Typename
}

func (obj LightsparkNodeWithRemoteSigning) GetAddresses(requester *requester.Requester, first *int64, types *[]NodeAddressType) (*NodeToAddressesConnection, error) {
	query := `query FetchNodeToAddressesConnection($entity_id: ID!, $first: Int, $types: [NodeAddressType!]) {
    entity(id: $entity_id) {
        ... on LightsparkNodeWithRemoteSigning {
            addresses(, first: $first, types: $types) {
                __typename
                node_to_addresses_connection_count: count
                node_to_addresses_connection_entities: entities {
                    __typename
                    node_address_address: address
                    node_address_type: type
                }
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id": obj.Id,
		"first":     first,
		"types":     types,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["addresses"].(map[string]interface{})
	var result *NodeToAddressesConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj LightsparkNodeWithRemoteSigning) GetChannels(requester *requester.Requester, first *int64, statuses *[]ChannelStatus, after *string) (*LightsparkNodeToChannelsConnection, error) {
	query := `query FetchLightsparkNodeToChannelsConnection($entity_id: ID!, $first: Int, $statuses: [ChannelStatus!], $after: String) {
    entity(id: $entity_id) {
        ... on LightsparkNodeWithRemoteSigning {
            channels(, first: $first, statuses: $statuses, after: $after) {
                __typename
                lightspark_node_to_channels_connection_count: count
                lightspark_node_to_channels_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                lightspark_node_to_channels_connection_entities: entities {
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
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id": obj.Id,
		"first":     first,
		"statuses":  statuses,
		"after":     after,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["channels"].(map[string]interface{})
	var result *LightsparkNodeToChannelsConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}
