// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"time"

	"github.com/lightsparkdev/go-sdk/requester"
	"github.com/lightsparkdev/go-sdk/types"
)

// This is a node that is managed by Lightspark and is managed within the current connected account. It contains many details about the node configuration, state, and metadata.
type LightsparkNode struct {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"lightspark_node_id"`

	// The date and time when the entity was first created.
	CreatedAt time.Time `json:"lightspark_node_created_at"`

	// The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"lightspark_node_updated_at"`

	// A name that identifies the node. It has no importance in terms of operating the node, it is just a way to identify and search for commercial services or popular nodes. This alias can be changed at any time by the node operator.
	Alias *string `json:"lightspark_node_alias"`

	// The Bitcoin Network this node is deployed in.
	BitcoinNetwork BitcoinNetwork `json:"lightspark_node_bitcoin_network"`

	// A hexadecimal string that describes a color. For example "#000000" is black, "#FFFFFF" is white. It has no importance in terms of operating the node, it is just a way to visually differentiate nodes. That color can be changed at any time by the node operator.
	Color *string `json:"lightspark_node_color"`

	// A summary metric used to capture how well positioned a node is to send, receive, or route transactions efficiently. Maximizing a node's conductivity helps a node’s transactions to be capital efficient. The value is an integer ranging between 0 and 10 (bounds included).
	Conductivity *int64 `json:"lightspark_node_conductivity"`

	// The name of this node in the network. It will be the most human-readable option possible, depending on the data available for this node.
	DisplayName string `json:"lightspark_node_display_name"`

	// The public key of this node. It acts as a unique identifier of this node in the Lightning Network.
	PublicKey *string `json:"lightspark_node_public_key"`

	// The account that owns this LightsparkNode.
	Account types.EntityWrapper `json:"lightspark_node_account"`

	// The owner of this LightsparkNode.
	Owner types.EntityWrapper `json:"lightspark_node_owner"`

	// The details of the balance of this node on the Bitcoin Network.
	BlockchainBalance *BlockchainBalance `json:"lightspark_node_blockchain_balance"`

	// The private key client is using to sign a GraphQL request which will be verified at LND.
	EncryptedSigningPrivateKey *Secret `json:"lightspark_node_encrypted_signing_private_key"`

	// The sum of the balance on the Bitcoin Network, channel balances, and commit fees on this node.
	TotalBalance *CurrencyAmount `json:"lightspark_node_total_balance"`

	// The total sum of the channel balances (online and offline) on this node.
	TotalLocalBalance *CurrencyAmount `json:"lightspark_node_total_local_balance"`

	// The sum of the channel balances (online only) that are available to send on this node.
	LocalBalance *CurrencyAmount `json:"lightspark_node_local_balance"`

	// The main purpose of this node. It is used by Lightspark for optimizations on the node's channels.
	Purpose *LightsparkNodePurpose `json:"lightspark_node_purpose"`

	// The sum of the channel balances that are available to receive on this node.
	RemoteBalance *CurrencyAmount `json:"lightspark_node_remote_balance"`

	// The current status of this node.
	Status *LightsparkNodeStatus `json:"lightspark_node_status"`
}

const (
	LightsparkNodeFragment = `
fragment LightsparkNodeFragment on LightsparkNode {
    __typename
    lightspark_node_id: id
    lightspark_node_created_at: created_at
    lightspark_node_updated_at: updated_at
    lightspark_node_alias: alias
    lightspark_node_bitcoin_network: bitcoin_network
    lightspark_node_color: color
    lightspark_node_conductivity: conductivity
    lightspark_node_display_name: display_name
    lightspark_node_public_key: public_key
    lightspark_node_account: account {
        id
    }
    lightspark_node_owner: owner {
        id
    }
    lightspark_node_blockchain_balance: blockchain_balance {
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
    lightspark_node_encrypted_signing_private_key: encrypted_signing_private_key {
        __typename
        secret_encrypted_value: encrypted_value
        secret_cipher: cipher
    }
    lightspark_node_total_balance: total_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    lightspark_node_total_local_balance: total_local_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    lightspark_node_local_balance: local_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    lightspark_node_purpose: purpose
    lightspark_node_remote_balance: remote_balance {
        __typename
        currency_amount_original_value: original_value
        currency_amount_original_unit: original_unit
        currency_amount_preferred_currency_unit: preferred_currency_unit
        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
    }
    lightspark_node_status: status
}
`
)

// A name that identifies the node. It has no importance in terms of operating the node, it is just a way to identify and search for commercial services or popular nodes. This alias can be changed at any time by the node operator.
func (obj LightsparkNode) GetAlias() *string {
	return obj.Alias
}

// The Bitcoin Network this node is deployed in.
func (obj LightsparkNode) GetBitcoinNetwork() BitcoinNetwork {
	return obj.BitcoinNetwork
}

// A hexadecimal string that describes a color. For example "#000000" is black, "#FFFFFF" is white. It has no importance in terms of operating the node, it is just a way to visually differentiate nodes. That color can be changed at any time by the node operator.
func (obj LightsparkNode) GetColor() *string {
	return obj.Color
}

// A summary metric used to capture how well positioned a node is to send, receive, or route transactions efficiently. Maximizing a node's conductivity helps a node’s transactions to be capital efficient. The value is an integer ranging between 0 and 10 (bounds included).
func (obj LightsparkNode) GetConductivity() *int64 {
	return obj.Conductivity
}

// The name of this node in the network. It will be the most human-readable option possible, depending on the data available for this node.
func (obj LightsparkNode) GetDisplayName() string {
	return obj.DisplayName
}

// The public key of this node. It acts as a unique identifier of this node in the Lightning Network.
func (obj LightsparkNode) GetPublicKey() *string {
	return obj.PublicKey
}

// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj LightsparkNode) GetId() string {
	return obj.Id
}

// The date and time when the entity was first created.
func (obj LightsparkNode) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// The date and time when the entity was last updated.
func (obj LightsparkNode) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj LightsparkNode) GetAddresses(requester *requester.Requester, first *int64, types *[]NodeAddressType) (*NodeToAddressesConnection, error) {
	query := `query FetchNodeToAddressesConnection($entity_id: ID!, $first: Int, $types: [NodeAddressType!]) {
    entity(id: $entity_id) {
        ... on LightsparkNode {
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

func (obj LightsparkNode) GetChannels(requester *requester.Requester, first *int64, statuses *[]ChannelStatus) (*LightsparkNodeToChannelsConnection, error) {
	query := `query FetchLightsparkNodeToChannelsConnection($entity_id: ID!, $first: Int, $statuses: [ChannelStatus!]) {
    entity(id: $entity_id) {
        ... on LightsparkNode {
            channels(, first: $first, statuses: $statuses) {
                __typename
                lightspark_node_to_channels_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                lightspark_node_to_channels_connection_count: count
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
