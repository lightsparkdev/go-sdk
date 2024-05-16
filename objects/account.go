// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"time"

	"github.com/lightsparkdev/go-sdk/requester"
)

// Account This is an object representing the connected Lightspark account. You can retrieve this object to see your account information and objects tied to your account.
type Account struct {

	// Id The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"account_id"`

	// CreatedAt The date and time when the entity was first created.
	CreatedAt time.Time `json:"account_created_at"`

	// UpdatedAt The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"account_updated_at"`

	// Name The name of this account.
	Name *string `json:"account_name"`

	// Typename The typename of the object
	Typename string `json:"__typename"`
}

const (
	AccountFragment = `
fragment AccountFragment on Account {
    __typename
    account_id: id
    account_created_at: created_at
    account_updated_at: updated_at
    account_name: name
}
`
)

// GetId The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj Account) GetId() string {
	return obj.Id
}

// GetCreatedAt The date and time when the entity was first created.
func (obj Account) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// GetUpdatedAt The date and time when the entity was last updated.
func (obj Account) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj Account) GetTypename() string {
	return obj.Typename
}

func (obj Account) GetApiTokens(requester *requester.Requester, first *int64, after *string) (*AccountToApiTokensConnection, error) {
	query := `query FetchAccountToApiTokensConnection($entity_id: ID!, $first: Int, $after: String) {
    entity(id: $entity_id) {
        ... on Account {
            api_tokens(, first: $first, after: $after) {
                __typename
                account_to_api_tokens_connection_count: count
                account_to_api_tokens_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                account_to_api_tokens_connection_entities: entities {
                    __typename
                    api_token_id: id
                    api_token_created_at: created_at
                    api_token_updated_at: updated_at
                    api_token_client_id: client_id
                    api_token_name: name
                    api_token_permissions: permissions
                    api_token_is_deleted: is_deleted
                }
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id": obj.Id,
		"first":     first,
		"after":     after,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["api_tokens"].(map[string]interface{})
	var result *AccountToApiTokensConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Account) GetBlockchainBalance(requester *requester.Requester, bitcoinNetworks *[]BitcoinNetwork, nodeIds *[]string) (*BlockchainBalance, error) {
	query := `query FetchAccountBlockchainBalance($entity_id: ID!, $bitcoin_networks: [BitcoinNetwork!], $node_ids: [ID!]) {
    entity(id: $entity_id) {
        ... on Account {
            blockchain_balance(, bitcoin_networks: $bitcoin_networks, node_ids: $node_ids) {
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
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":        obj.Id,
		"bitcoin_networks": bitcoinNetworks,
		"node_ids":         nodeIds,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["blockchain_balance"].(map[string]interface{})
	var result *BlockchainBalance
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Account) GetConductivity(requester *requester.Requester, bitcoinNetworks *[]BitcoinNetwork, nodeIds *[]string) (*int64, error) {
	query := `query FetchAccountConductivity($entity_id: ID!, $bitcoin_networks: [BitcoinNetwork!], $node_ids: [ID!]) {
    entity(id: $entity_id) {
        ... on Account {
            conductivity(, bitcoin_networks: $bitcoin_networks, node_ids: $node_ids)
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":        obj.Id,
		"bitcoin_networks": bitcoinNetworks,
		"node_ids":         nodeIds,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["conductivity"]
	var result *int64
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Account) GetLocalBalance(requester *requester.Requester, bitcoinNetworks *[]BitcoinNetwork, nodeIds *[]string) (*CurrencyAmount, error) {
	query := `query FetchAccountLocalBalance($entity_id: ID!, $bitcoin_networks: [BitcoinNetwork!], $node_ids: [ID!]) {
    entity(id: $entity_id) {
        ... on Account {
            local_balance(, bitcoin_networks: $bitcoin_networks, node_ids: $node_ids) {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":        obj.Id,
		"bitcoin_networks": bitcoinNetworks,
		"node_ids":         nodeIds,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["local_balance"].(map[string]interface{})
	var result *CurrencyAmount
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Account) GetNodes(requester *requester.Requester, first *int64, bitcoinNetworks *[]BitcoinNetwork, nodeIds *[]string, after *string) (*AccountToNodesConnection, error) {
	query := `query FetchAccountToNodesConnection($entity_id: ID!, $first: Int, $bitcoin_networks: [BitcoinNetwork!], $node_ids: [ID!], $after: String) {
    entity(id: $entity_id) {
        ... on Account {
            nodes(, first: $first, bitcoin_networks: $bitcoin_networks, node_ids: $node_ids, after: $after) {
                __typename
                account_to_nodes_connection_count: count
                account_to_nodes_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                account_to_nodes_connection_entities: entities {
                    __typename
                    ... on LightsparkNodeWithOSK {
                        __typename
                        lightspark_node_with_o_s_k_id: id
                        lightspark_node_with_o_s_k_created_at: created_at
                        lightspark_node_with_o_s_k_updated_at: updated_at
                        lightspark_node_with_o_s_k_alias: alias
                        lightspark_node_with_o_s_k_bitcoin_network: bitcoin_network
                        lightspark_node_with_o_s_k_color: color
                        lightspark_node_with_o_s_k_conductivity: conductivity
                        lightspark_node_with_o_s_k_display_name: display_name
                        lightspark_node_with_o_s_k_public_key: public_key
                        lightspark_node_with_o_s_k_owner: owner {
                            id
                        }
                        lightspark_node_with_o_s_k_status: status
                        lightspark_node_with_o_s_k_total_balance: total_balance {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        lightspark_node_with_o_s_k_total_local_balance: total_local_balance {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        lightspark_node_with_o_s_k_local_balance: local_balance {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        lightspark_node_with_o_s_k_remote_balance: remote_balance {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        lightspark_node_with_o_s_k_blockchain_balance: blockchain_balance {
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
                        lightspark_node_with_o_s_k_uma_prescreening_utxos: uma_prescreening_utxos
                        lightspark_node_with_o_s_k_balances: balances {
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
                        lightspark_node_with_o_s_k_encrypted_signing_private_key: encrypted_signing_private_key {
                            __typename
                            secret_encrypted_value: encrypted_value
                            secret_cipher: cipher
                        }
                    }
                    ... on LightsparkNodeWithRemoteSigning {
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
                }
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":        obj.Id,
		"first":            first,
		"bitcoin_networks": bitcoinNetworks,
		"node_ids":         nodeIds,
		"after":            after,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["nodes"].(map[string]interface{})
	var result *AccountToNodesConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Account) GetRemoteBalance(requester *requester.Requester, bitcoinNetworks *[]BitcoinNetwork, nodeIds *[]string) (*CurrencyAmount, error) {
	query := `query FetchAccountRemoteBalance($entity_id: ID!, $bitcoin_networks: [BitcoinNetwork!], $node_ids: [ID!]) {
    entity(id: $entity_id) {
        ... on Account {
            remote_balance(, bitcoin_networks: $bitcoin_networks, node_ids: $node_ids) {
                __typename
                currency_amount_original_value: original_value
                currency_amount_original_unit: original_unit
                currency_amount_preferred_currency_unit: preferred_currency_unit
                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":        obj.Id,
		"bitcoin_networks": bitcoinNetworks,
		"node_ids":         nodeIds,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["remote_balance"].(map[string]interface{})
	var result *CurrencyAmount
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Account) GetUptimePercentage(requester *requester.Requester, afterDate *time.Time, beforeDate *time.Time, bitcoinNetworks *[]BitcoinNetwork, nodeIds *[]string) (*int64, error) {
	query := `query FetchAccountUptimePercentage($entity_id: ID!, $after_date: DateTime, $before_date: DateTime, $bitcoin_networks: [BitcoinNetwork!], $node_ids: [ID!]) {
    entity(id: $entity_id) {
        ... on Account {
            uptime_percentage(, after_date: $after_date, before_date: $before_date, bitcoin_networks: $bitcoin_networks, node_ids: $node_ids)
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":        obj.Id,
		"after_date":       afterDate,
		"before_date":      beforeDate,
		"bitcoin_networks": bitcoinNetworks,
		"node_ids":         nodeIds,
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

func (obj Account) GetChannels(requester *requester.Requester, bitcoinNetwork BitcoinNetwork, lightningNodeId *string, afterDate *time.Time, beforeDate *time.Time, first *int64, after *string) (*AccountToChannelsConnection, error) {
	query := `query FetchAccountToChannelsConnection($entity_id: ID!, $bitcoin_network: BitcoinNetwork!, $lightning_node_id: ID, $after_date: DateTime, $before_date: DateTime, $first: Int, $after: String) {
    entity(id: $entity_id) {
        ... on Account {
            channels(, bitcoin_network: $bitcoin_network, lightning_node_id: $lightning_node_id, after_date: $after_date, before_date: $before_date, first: $first, after: $after) {
                __typename
                account_to_channels_connection_count: count
                account_to_channels_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                account_to_channels_connection_entities: entities {
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
		"entity_id":         obj.Id,
		"bitcoin_network":   bitcoinNetwork,
		"lightning_node_id": lightningNodeId,
		"after_date":        afterDate,
		"before_date":       beforeDate,
		"first":             first,
		"after":             after,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["channels"].(map[string]interface{})
	var result *AccountToChannelsConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Account) GetTransactions(requester *requester.Requester, first *int64, after *string, types *[]TransactionType, afterDate *time.Time, beforeDate *time.Time, bitcoinNetwork *BitcoinNetwork, lightningNodeId *string, statuses *[]TransactionStatus, excludeFailures *TransactionFailures) (*AccountToTransactionsConnection, error) {
	query := `query FetchAccountToTransactionsConnection($entity_id: ID!, $first: Int, $after: String, $types: [TransactionType!], $after_date: DateTime, $before_date: DateTime, $bitcoin_network: BitcoinNetwork, $lightning_node_id: ID, $statuses: [TransactionStatus!], $exclude_failures: TransactionFailures) {
    entity(id: $entity_id) {
        ... on Account {
            transactions(, first: $first, after: $after, types: $types, after_date: $after_date, before_date: $before_date, bitcoin_network: $bitcoin_network, lightning_node_id: $lightning_node_id, statuses: $statuses, exclude_failures: $exclude_failures) {
                __typename
                account_to_transactions_connection_count: count
                account_to_transactions_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                account_to_transactions_connection_profit_loss: profit_loss {
                    __typename
                    currency_amount_original_value: original_value
                    currency_amount_original_unit: original_unit
                    currency_amount_preferred_currency_unit: preferred_currency_unit
                    currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                    currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                }
                account_to_transactions_connection_average_fee_earned: average_fee_earned {
                    __typename
                    currency_amount_original_value: original_value
                    currency_amount_original_unit: original_unit
                    currency_amount_preferred_currency_unit: preferred_currency_unit
                    currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                    currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                }
                account_to_transactions_connection_total_amount_transacted: total_amount_transacted {
                    __typename
                    currency_amount_original_value: original_value
                    currency_amount_original_unit: original_unit
                    currency_amount_preferred_currency_unit: preferred_currency_unit
                    currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                    currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                }
                account_to_transactions_connection_entities: entities {
                    __typename
                    ... on ChannelClosingTransaction {
                        __typename
                        channel_closing_transaction_id: id
                        channel_closing_transaction_created_at: created_at
                        channel_closing_transaction_updated_at: updated_at
                        channel_closing_transaction_status: status
                        channel_closing_transaction_resolved_at: resolved_at
                        channel_closing_transaction_amount: amount {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        channel_closing_transaction_transaction_hash: transaction_hash
                        channel_closing_transaction_fees: fees {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        channel_closing_transaction_block_hash: block_hash
                        channel_closing_transaction_block_height: block_height
                        channel_closing_transaction_destination_addresses: destination_addresses
                        channel_closing_transaction_num_confirmations: num_confirmations
                        channel_closing_transaction_channel: channel {
                            id
                        }
                    }
                    ... on ChannelOpeningTransaction {
                        __typename
                        channel_opening_transaction_id: id
                        channel_opening_transaction_created_at: created_at
                        channel_opening_transaction_updated_at: updated_at
                        channel_opening_transaction_status: status
                        channel_opening_transaction_resolved_at: resolved_at
                        channel_opening_transaction_amount: amount {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        channel_opening_transaction_transaction_hash: transaction_hash
                        channel_opening_transaction_fees: fees {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        channel_opening_transaction_block_hash: block_hash
                        channel_opening_transaction_block_height: block_height
                        channel_opening_transaction_destination_addresses: destination_addresses
                        channel_opening_transaction_num_confirmations: num_confirmations
                        channel_opening_transaction_channel: channel {
                            id
                        }
                    }
                    ... on Deposit {
                        __typename
                        deposit_id: id
                        deposit_created_at: created_at
                        deposit_updated_at: updated_at
                        deposit_status: status
                        deposit_resolved_at: resolved_at
                        deposit_amount: amount {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        deposit_transaction_hash: transaction_hash
                        deposit_fees: fees {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        deposit_block_hash: block_hash
                        deposit_block_height: block_height
                        deposit_destination_addresses: destination_addresses
                        deposit_num_confirmations: num_confirmations
                        deposit_destination: destination {
                            id
                        }
                    }
                    ... on IncomingPayment {
                        __typename
                        incoming_payment_id: id
                        incoming_payment_created_at: created_at
                        incoming_payment_updated_at: updated_at
                        incoming_payment_status: status
                        incoming_payment_resolved_at: resolved_at
                        incoming_payment_amount: amount {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        incoming_payment_transaction_hash: transaction_hash
                        incoming_payment_is_uma: is_uma
                        incoming_payment_destination: destination {
                            id
                        }
                        incoming_payment_payment_request: payment_request {
                            id
                        }
                        incoming_payment_uma_post_transaction_data: uma_post_transaction_data {
                            __typename
                            post_transaction_data_utxo: utxo
                            post_transaction_data_amount: amount {
                                __typename
                                currency_amount_original_value: original_value
                                currency_amount_original_unit: original_unit
                                currency_amount_preferred_currency_unit: preferred_currency_unit
                                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                            }
                        }
                        incoming_payment_is_internal_payment: is_internal_payment
                    }
                    ... on OutgoingPayment {
                        __typename
                        outgoing_payment_id: id
                        outgoing_payment_created_at: created_at
                        outgoing_payment_updated_at: updated_at
                        outgoing_payment_status: status
                        outgoing_payment_resolved_at: resolved_at
                        outgoing_payment_amount: amount {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        outgoing_payment_transaction_hash: transaction_hash
                        outgoing_payment_is_uma: is_uma
                        outgoing_payment_origin: origin {
                            id
                        }
                        outgoing_payment_destination: destination {
                            id
                        }
                        outgoing_payment_fees: fees {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        outgoing_payment_payment_request_data: payment_request_data {
                            __typename
                            ... on InvoiceData {
                                __typename
                                invoice_data_encoded_payment_request: encoded_payment_request
                                invoice_data_bitcoin_network: bitcoin_network
                                invoice_data_payment_hash: payment_hash
                                invoice_data_amount: amount {
                                    __typename
                                    currency_amount_original_value: original_value
                                    currency_amount_original_unit: original_unit
                                    currency_amount_preferred_currency_unit: preferred_currency_unit
                                    currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                    currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                }
                                invoice_data_created_at: created_at
                                invoice_data_expires_at: expires_at
                                invoice_data_memo: memo
                                invoice_data_destination: destination {
                                    __typename
                                    ... on GraphNode {
                                        __typename
                                        graph_node_id: id
                                        graph_node_created_at: created_at
                                        graph_node_updated_at: updated_at
                                        graph_node_alias: alias
                                        graph_node_bitcoin_network: bitcoin_network
                                        graph_node_color: color
                                        graph_node_conductivity: conductivity
                                        graph_node_display_name: display_name
                                        graph_node_public_key: public_key
                                    }
                                    ... on LightsparkNodeWithOSK {
                                        __typename
                                        lightspark_node_with_o_s_k_id: id
                                        lightspark_node_with_o_s_k_created_at: created_at
                                        lightspark_node_with_o_s_k_updated_at: updated_at
                                        lightspark_node_with_o_s_k_alias: alias
                                        lightspark_node_with_o_s_k_bitcoin_network: bitcoin_network
                                        lightspark_node_with_o_s_k_color: color
                                        lightspark_node_with_o_s_k_conductivity: conductivity
                                        lightspark_node_with_o_s_k_display_name: display_name
                                        lightspark_node_with_o_s_k_public_key: public_key
                                        lightspark_node_with_o_s_k_owner: owner {
                                            id
                                        }
                                        lightspark_node_with_o_s_k_status: status
                                        lightspark_node_with_o_s_k_total_balance: total_balance {
                                            __typename
                                            currency_amount_original_value: original_value
                                            currency_amount_original_unit: original_unit
                                            currency_amount_preferred_currency_unit: preferred_currency_unit
                                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                        }
                                        lightspark_node_with_o_s_k_total_local_balance: total_local_balance {
                                            __typename
                                            currency_amount_original_value: original_value
                                            currency_amount_original_unit: original_unit
                                            currency_amount_preferred_currency_unit: preferred_currency_unit
                                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                        }
                                        lightspark_node_with_o_s_k_local_balance: local_balance {
                                            __typename
                                            currency_amount_original_value: original_value
                                            currency_amount_original_unit: original_unit
                                            currency_amount_preferred_currency_unit: preferred_currency_unit
                                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                        }
                                        lightspark_node_with_o_s_k_remote_balance: remote_balance {
                                            __typename
                                            currency_amount_original_value: original_value
                                            currency_amount_original_unit: original_unit
                                            currency_amount_preferred_currency_unit: preferred_currency_unit
                                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                        }
                                        lightspark_node_with_o_s_k_blockchain_balance: blockchain_balance {
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
                                        lightspark_node_with_o_s_k_uma_prescreening_utxos: uma_prescreening_utxos
                                        lightspark_node_with_o_s_k_balances: balances {
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
                                        lightspark_node_with_o_s_k_encrypted_signing_private_key: encrypted_signing_private_key {
                                            __typename
                                            secret_encrypted_value: encrypted_value
                                            secret_cipher: cipher
                                        }
                                    }
                                    ... on LightsparkNodeWithRemoteSigning {
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
                                }
                            }
                        }
                        outgoing_payment_failure_reason: failure_reason
                        outgoing_payment_failure_message: failure_message {
                            __typename
                            rich_text_text: text
                        }
                        outgoing_payment_uma_post_transaction_data: uma_post_transaction_data {
                            __typename
                            post_transaction_data_utxo: utxo
                            post_transaction_data_amount: amount {
                                __typename
                                currency_amount_original_value: original_value
                                currency_amount_original_unit: original_unit
                                currency_amount_preferred_currency_unit: preferred_currency_unit
                                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                            }
                        }
                        outgoing_payment_payment_preimage: payment_preimage
                        outgoing_payment_is_internal_payment: is_internal_payment
                        outgoing_payment_idempotency_key: idempotency_key
                    }
                    ... on RoutingTransaction {
                        __typename
                        routing_transaction_id: id
                        routing_transaction_created_at: created_at
                        routing_transaction_updated_at: updated_at
                        routing_transaction_status: status
                        routing_transaction_resolved_at: resolved_at
                        routing_transaction_amount: amount {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        routing_transaction_transaction_hash: transaction_hash
                        routing_transaction_incoming_channel: incoming_channel {
                            id
                        }
                        routing_transaction_outgoing_channel: outgoing_channel {
                            id
                        }
                        routing_transaction_fees: fees {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        routing_transaction_failure_message: failure_message {
                            __typename
                            rich_text_text: text
                        }
                        routing_transaction_failure_reason: failure_reason
                    }
                    ... on Withdrawal {
                        __typename
                        withdrawal_id: id
                        withdrawal_created_at: created_at
                        withdrawal_updated_at: updated_at
                        withdrawal_status: status
                        withdrawal_resolved_at: resolved_at
                        withdrawal_amount: amount {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        withdrawal_transaction_hash: transaction_hash
                        withdrawal_fees: fees {
                            __typename
                            currency_amount_original_value: original_value
                            currency_amount_original_unit: original_unit
                            currency_amount_preferred_currency_unit: preferred_currency_unit
                            currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                            currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                        }
                        withdrawal_block_hash: block_hash
                        withdrawal_block_height: block_height
                        withdrawal_destination_addresses: destination_addresses
                        withdrawal_num_confirmations: num_confirmations
                        withdrawal_origin: origin {
                            id
                        }
                    }
                }
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":         obj.Id,
		"first":             first,
		"after":             after,
		"types":             types,
		"after_date":        afterDate,
		"before_date":       beforeDate,
		"bitcoin_network":   bitcoinNetwork,
		"lightning_node_id": lightningNodeId,
		"statuses":          statuses,
		"exclude_failures":  excludeFailures,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["transactions"].(map[string]interface{})
	var result *AccountToTransactionsConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Account) GetPaymentRequests(requester *requester.Requester, first *int64, after *string, afterDate *time.Time, beforeDate *time.Time, bitcoinNetwork *BitcoinNetwork, lightningNodeId *string) (*AccountToPaymentRequestsConnection, error) {
	query := `query FetchAccountToPaymentRequestsConnection($entity_id: ID!, $first: Int, $after: String, $after_date: DateTime, $before_date: DateTime, $bitcoin_network: BitcoinNetwork, $lightning_node_id: ID) {
    entity(id: $entity_id) {
        ... on Account {
            payment_requests(, first: $first, after: $after, after_date: $after_date, before_date: $before_date, bitcoin_network: $bitcoin_network, lightning_node_id: $lightning_node_id) {
                __typename
                account_to_payment_requests_connection_count: count
                account_to_payment_requests_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                account_to_payment_requests_connection_entities: entities {
                    __typename
                    ... on Invoice {
                        __typename
                        invoice_id: id
                        invoice_created_at: created_at
                        invoice_updated_at: updated_at
                        invoice_data: data {
                            __typename
                            invoice_data_encoded_payment_request: encoded_payment_request
                            invoice_data_bitcoin_network: bitcoin_network
                            invoice_data_payment_hash: payment_hash
                            invoice_data_amount: amount {
                                __typename
                                currency_amount_original_value: original_value
                                currency_amount_original_unit: original_unit
                                currency_amount_preferred_currency_unit: preferred_currency_unit
                                currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                            }
                            invoice_data_created_at: created_at
                            invoice_data_expires_at: expires_at
                            invoice_data_memo: memo
                            invoice_data_destination: destination {
                                __typename
                                ... on GraphNode {
                                    __typename
                                    graph_node_id: id
                                    graph_node_created_at: created_at
                                    graph_node_updated_at: updated_at
                                    graph_node_alias: alias
                                    graph_node_bitcoin_network: bitcoin_network
                                    graph_node_color: color
                                    graph_node_conductivity: conductivity
                                    graph_node_display_name: display_name
                                    graph_node_public_key: public_key
                                }
                                ... on LightsparkNodeWithOSK {
                                    __typename
                                    lightspark_node_with_o_s_k_id: id
                                    lightspark_node_with_o_s_k_created_at: created_at
                                    lightspark_node_with_o_s_k_updated_at: updated_at
                                    lightspark_node_with_o_s_k_alias: alias
                                    lightspark_node_with_o_s_k_bitcoin_network: bitcoin_network
                                    lightspark_node_with_o_s_k_color: color
                                    lightspark_node_with_o_s_k_conductivity: conductivity
                                    lightspark_node_with_o_s_k_display_name: display_name
                                    lightspark_node_with_o_s_k_public_key: public_key
                                    lightspark_node_with_o_s_k_owner: owner {
                                        id
                                    }
                                    lightspark_node_with_o_s_k_status: status
                                    lightspark_node_with_o_s_k_total_balance: total_balance {
                                        __typename
                                        currency_amount_original_value: original_value
                                        currency_amount_original_unit: original_unit
                                        currency_amount_preferred_currency_unit: preferred_currency_unit
                                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                    }
                                    lightspark_node_with_o_s_k_total_local_balance: total_local_balance {
                                        __typename
                                        currency_amount_original_value: original_value
                                        currency_amount_original_unit: original_unit
                                        currency_amount_preferred_currency_unit: preferred_currency_unit
                                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                    }
                                    lightspark_node_with_o_s_k_local_balance: local_balance {
                                        __typename
                                        currency_amount_original_value: original_value
                                        currency_amount_original_unit: original_unit
                                        currency_amount_preferred_currency_unit: preferred_currency_unit
                                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                    }
                                    lightspark_node_with_o_s_k_remote_balance: remote_balance {
                                        __typename
                                        currency_amount_original_value: original_value
                                        currency_amount_original_unit: original_unit
                                        currency_amount_preferred_currency_unit: preferred_currency_unit
                                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                                    }
                                    lightspark_node_with_o_s_k_blockchain_balance: blockchain_balance {
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
                                    lightspark_node_with_o_s_k_uma_prescreening_utxos: uma_prescreening_utxos
                                    lightspark_node_with_o_s_k_balances: balances {
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
                                    lightspark_node_with_o_s_k_encrypted_signing_private_key: encrypted_signing_private_key {
                                        __typename
                                        secret_encrypted_value: encrypted_value
                                        secret_cipher: cipher
                                    }
                                }
                                ... on LightsparkNodeWithRemoteSigning {
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
                            }
                        }
                        invoice_status: status
                        invoice_amount_paid: amount_paid {
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
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":         obj.Id,
		"first":             first,
		"after":             after,
		"after_date":        afterDate,
		"before_date":       beforeDate,
		"bitcoin_network":   bitcoinNetwork,
		"lightning_node_id": lightningNodeId,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["payment_requests"].(map[string]interface{})
	var result *AccountToPaymentRequestsConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Account) GetWithdrawalRequests(requester *requester.Requester, first *int64, after *string, bitcoinNetworks *[]BitcoinNetwork, statuses *[]WithdrawalRequestStatus, nodeIds *[]string, idempotencyKeys *[]string, afterDate *time.Time, beforeDate *time.Time) (*AccountToWithdrawalRequestsConnection, error) {
	query := `query FetchAccountToWithdrawalRequestsConnection($entity_id: ID!, $first: Int, $after: String, $bitcoin_networks: [BitcoinNetwork!], $statuses: [WithdrawalRequestStatus!], $node_ids: [ID!], $idempotency_keys: [String!], $after_date: DateTime, $before_date: DateTime) {
    entity(id: $entity_id) {
        ... on Account {
            withdrawal_requests(, first: $first, after: $after, bitcoin_networks: $bitcoin_networks, statuses: $statuses, node_ids: $node_ids, idempotency_keys: $idempotency_keys, after_date: $after_date, before_date: $before_date) {
                __typename
                account_to_withdrawal_requests_connection_count: count
                account_to_withdrawal_requests_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                account_to_withdrawal_requests_connection_entities: entities {
                    __typename
                    withdrawal_request_id: id
                    withdrawal_request_created_at: created_at
                    withdrawal_request_updated_at: updated_at
                    withdrawal_request_requested_amount: requested_amount {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    withdrawal_request_amount: amount {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    withdrawal_request_estimated_amount: estimated_amount {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    withdrawal_request_amount_withdrawn: amount_withdrawn {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    withdrawal_request_total_fees: total_fees {
                        __typename
                        currency_amount_original_value: original_value
                        currency_amount_original_unit: original_unit
                        currency_amount_preferred_currency_unit: preferred_currency_unit
                        currency_amount_preferred_currency_value_rounded: preferred_currency_value_rounded
                        currency_amount_preferred_currency_value_approx: preferred_currency_value_approx
                    }
                    withdrawal_request_bitcoin_address: bitcoin_address
                    withdrawal_request_withdrawal_mode: withdrawal_mode
                    withdrawal_request_status: status
                    withdrawal_request_completed_at: completed_at
                    withdrawal_request_withdrawal: withdrawal {
                        id
                    }
                    withdrawal_request_idempotency_key: idempotency_key
                }
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":        obj.Id,
		"first":            first,
		"after":            after,
		"bitcoin_networks": bitcoinNetworks,
		"statuses":         statuses,
		"node_ids":         nodeIds,
		"idempotency_keys": idempotencyKeys,
		"after_date":       afterDate,
		"before_date":      beforeDate,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["withdrawal_requests"].(map[string]interface{})
	var result *AccountToWithdrawalRequestsConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Account) GetWallets(requester *requester.Requester, first *int64, after *string, thirdPartyIds *[]string) (*AccountToWalletsConnection, error) {
	query := `query FetchAccountToWalletsConnection($entity_id: ID!, $first: Int, $after: String, $third_party_ids: [String!]) {
    entity(id: $entity_id) {
        ... on Account {
            wallets(, first: $first, after: $after, third_party_ids: $third_party_ids) {
                __typename
                account_to_wallets_connection_count: count
                account_to_wallets_connection_page_info: page_info {
                    __typename
                    page_info_has_next_page: has_next_page
                    page_info_has_previous_page: has_previous_page
                    page_info_start_cursor: start_cursor
                    page_info_end_cursor: end_cursor
                }
                account_to_wallets_connection_entities: entities {
                    __typename
                    wallet_id: id
                    wallet_created_at: created_at
                    wallet_updated_at: updated_at
                    wallet_last_login_at: last_login_at
                    wallet_balances: balances {
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
                    wallet_third_party_identifier: third_party_identifier
                    wallet_account: account {
                        id
                    }
                    wallet_status: status
                }
            }
        }
    }
}`
	variables := map[string]interface{}{
		"entity_id":       obj.Id,
		"first":           first,
		"after":           after,
		"third_party_ids": thirdPartyIds,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["wallets"].(map[string]interface{})
	var result *AccountToWalletsConnection
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}
