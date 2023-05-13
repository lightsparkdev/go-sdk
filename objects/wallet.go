// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"lightspark/requester"
	"time"
)

type Wallet struct {

	// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
	Id string `json:"wallet_id"`

	// The date and time when the entity was first created.
	CreatedAt time.Time `json:"wallet_created_at"`

	// The date and time when the entity was last updated.
	UpdatedAt time.Time `json:"wallet_updated_at"`

	// The date and time when the wallet user last logged in.
	LastLoginAt *time.Time `json:"wallet_last_login_at"`

	// The balances that describe the funds in this wallet.
	Balances *Balances `json:"wallet_balances"`

	// The unique identifier of this wallet, as provided by the Lightspark Customer during login.
	ThirdPartyIdentifier string `json:"wallet_third_party_identifier"`
}

const (
	WalletFragment = `
fragment WalletFragment on Wallet {
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
}
`
)

// The unique identifier of this entity across all Lightspark systems. Should be treated as an opaque string.
func (obj Wallet) GetId() string {
	return obj.Id
}

// The date and time when the entity was first created.
func (obj Wallet) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

// The date and time when the entity was last updated.
func (obj Wallet) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}

func (obj Wallet) GetTotalAmountReceived(requester *requester.Requester, createdAfterDate *time.Time, createdBeforeDate *time.Time) (*CurrencyAmount, error) {
	query := `query FetchWalletTotalAmountReceived($entity_id: ID!, $created_after_date: DateTime, $created_before_date: DateTime) {
    entity(id: $entity_id) {
        ... on Wallet {
            total_amount_received(, created_after_date: $created_after_date, created_before_date: $created_before_date) {
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
		"entity_id":           obj.Id,
		"created_after_date":  createdAfterDate,
		"created_before_date": createdBeforeDate,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["total_amount_received"].(map[string]interface{})
	var result *CurrencyAmount
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}

func (obj Wallet) GetTotalAmountSent(requester *requester.Requester, createdAfterDate *time.Time, createdBeforeDate *time.Time) (*CurrencyAmount, error) {
	query := `query FetchWalletTotalAmountSent($entity_id: ID!, $created_after_date: DateTime, $created_before_date: DateTime) {
    entity(id: $entity_id) {
        ... on Wallet {
            total_amount_sent(, created_after_date: $created_after_date, created_before_date: $created_before_date) {
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
		"entity_id":           obj.Id,
		"created_after_date":  createdAfterDate,
		"created_before_date": createdBeforeDate,
	}

	response, err := requester.ExecuteGraphql(query, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})["total_amount_sent"].(map[string]interface{})
	var result *CurrencyAmount
	jsonString, err := json.Marshal(output)
	json.Unmarshal(jsonString, &result)
	return result, nil
}
