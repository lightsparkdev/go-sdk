// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// LightsparkNode This is an object representing a node managed by Lightspark and owned by the current connected account. This object contains information about the node’s configuration, state, and metadata.
type LightsparkNode interface {
	Node
	Entity

	// GetOwnerId The owner of this LightsparkNode.
	GetOwnerId() types.EntityWrapper

	// GetStatus The current status of this node.
	GetStatus() *LightsparkNodeStatus

	// GetTotalBalance The sum of the balance on the Bitcoin Network, channel balances, and commit fees on this node.
	// Deprecated: Use `balances` instead.
	GetTotalBalance() *CurrencyAmount

	// GetTotalLocalBalance The total sum of the channel balances (online and offline) on this node.
	// Deprecated: Use `balances` instead.
	GetTotalLocalBalance() *CurrencyAmount

	// GetLocalBalance The sum of the channel balances (online only) that are available to send on this node.
	// Deprecated: Use `balances` instead.
	GetLocalBalance() *CurrencyAmount

	// GetRemoteBalance The sum of the channel balances that are available to receive on this node.
	// Deprecated: Use `balances` instead.
	GetRemoteBalance() *CurrencyAmount

	// GetBlockchainBalance The details of the balance of this node on the Bitcoin Network.
	// Deprecated: Use `balances` instead.
	GetBlockchainBalance() *BlockchainBalance

	// GetUmaPrescreeningUtxos The utxos of the channels that are connected to this node. This is used in uma flow for pre-screening.
	GetUmaPrescreeningUtxos() []string

	// GetBalances The balances that describe the funds in this node.
	GetBalances() *Balances
}

func LightsparkNodeUnmarshal(data map[string]interface{}) (LightsparkNode, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "LightsparkNodeWithOSK":
		var lightsparkNodeWithOSK LightsparkNodeWithOSK
		if err := json.Unmarshal(dataJSON, &lightsparkNodeWithOSK); err != nil {
			return nil, err
		}
		return lightsparkNodeWithOSK, nil
	case "LightsparkNodeWithRemoteSigning":
		var lightsparkNodeWithRemoteSigning LightsparkNodeWithRemoteSigning
		if err := json.Unmarshal(dataJSON, &lightsparkNodeWithRemoteSigning); err != nil {
			return nil, err
		}
		return lightsparkNodeWithRemoteSigning, nil

	default:
		return nil, fmt.Errorf("unknown LightsparkNode type: %s", data["__typename"])
	}
}
