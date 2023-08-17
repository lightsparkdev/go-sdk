// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
)

type Connection interface {

	// GetCount The total count of objects in this connection, using the current filters. It is different from the number of objects returned in the current page (in the `entities` field).
	GetCount() int64

	// GetPageInfo An object that holds pagination information about the objects in this connection.
	GetPageInfo() PageInfo
}

func ConnectionUnmarshal(data map[string]interface{}) (Connection, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "AccountToApiTokensConnection":
		var accountToApiTokensConnection AccountToApiTokensConnection
		if err := json.Unmarshal(dataJSON, &accountToApiTokensConnection); err != nil {
			return nil, err
		}
		return accountToApiTokensConnection, nil
	case "AccountToNodesConnection":
		var accountToNodesConnection AccountToNodesConnection
		if err := json.Unmarshal(dataJSON, &accountToNodesConnection); err != nil {
			return nil, err
		}
		return accountToNodesConnection, nil
	case "AccountToPaymentRequestsConnection":
		var accountToPaymentRequestsConnection AccountToPaymentRequestsConnection
		if err := json.Unmarshal(dataJSON, &accountToPaymentRequestsConnection); err != nil {
			return nil, err
		}
		return accountToPaymentRequestsConnection, nil
	case "AccountToTransactionsConnection":
		var accountToTransactionsConnection AccountToTransactionsConnection
		if err := json.Unmarshal(dataJSON, &accountToTransactionsConnection); err != nil {
			return nil, err
		}
		return accountToTransactionsConnection, nil
	case "AccountToWalletsConnection":
		var accountToWalletsConnection AccountToWalletsConnection
		if err := json.Unmarshal(dataJSON, &accountToWalletsConnection); err != nil {
			return nil, err
		}
		return accountToWalletsConnection, nil
	case "IncomingPaymentToAttemptsConnection":
		var incomingPaymentToAttemptsConnection IncomingPaymentToAttemptsConnection
		if err := json.Unmarshal(dataJSON, &incomingPaymentToAttemptsConnection); err != nil {
			return nil, err
		}
		return incomingPaymentToAttemptsConnection, nil
	case "LightsparkNodeToChannelsConnection":
		var lightsparkNodeToChannelsConnection LightsparkNodeToChannelsConnection
		if err := json.Unmarshal(dataJSON, &lightsparkNodeToChannelsConnection); err != nil {
			return nil, err
		}
		return lightsparkNodeToChannelsConnection, nil
	case "OutgoingPaymentAttemptToHopsConnection":
		var outgoingPaymentAttemptToHopsConnection OutgoingPaymentAttemptToHopsConnection
		if err := json.Unmarshal(dataJSON, &outgoingPaymentAttemptToHopsConnection); err != nil {
			return nil, err
		}
		return outgoingPaymentAttemptToHopsConnection, nil
	case "OutgoingPaymentToAttemptsConnection":
		var outgoingPaymentToAttemptsConnection OutgoingPaymentToAttemptsConnection
		if err := json.Unmarshal(dataJSON, &outgoingPaymentToAttemptsConnection); err != nil {
			return nil, err
		}
		return outgoingPaymentToAttemptsConnection, nil

	default:
		return nil, fmt.Errorf("unknown Connection type: %s", data["__typename"])
	}
}
