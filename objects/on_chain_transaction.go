// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// OnChainTransaction This object represents an L1 transaction that occurred on the Bitcoin Network. You can retrieve this object to receive information about a specific on-chain transaction made on the Lightning Network associated with your Lightspark Node.
type OnChainTransaction interface {
	Transaction
	Entity

	// GetFees The fees that were paid by the node for this transaction.
	GetFees() *CurrencyAmount

	// GetBlockHash The hash of the block that included this transaction. This will be null for unconfirmed transactions.
	GetBlockHash() *string

	// GetBlockHeight The height of the block that included this transaction. This will be zero for unconfirmed transactions.
	GetBlockHeight() int64

	// GetDestinationAddresses The Bitcoin blockchain addresses this transaction was sent to.
	GetDestinationAddresses() []string

	// GetNumConfirmations The number of blockchain confirmations for this transaction in real time.
	GetNumConfirmations() *int64
}

func OnChainTransactionUnmarshal(data map[string]interface{}) (OnChainTransaction, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "ChannelClosingTransaction":
		var channelClosingTransaction ChannelClosingTransaction
		if err := json.Unmarshal(dataJSON, &channelClosingTransaction); err != nil {
			return nil, err
		}
		return channelClosingTransaction, nil
	case "ChannelOpeningTransaction":
		var channelOpeningTransaction ChannelOpeningTransaction
		if err := json.Unmarshal(dataJSON, &channelOpeningTransaction); err != nil {
			return nil, err
		}
		return channelOpeningTransaction, nil
	case "Deposit":
		var deposit Deposit
		if err := json.Unmarshal(dataJSON, &deposit); err != nil {
			return nil, err
		}
		return deposit, nil
	case "Withdrawal":
		var withdrawal Withdrawal
		if err := json.Unmarshal(dataJSON, &withdrawal); err != nil {
			return nil, err
		}
		return withdrawal, nil

	default:
		return nil, fmt.Errorf("unknown OnChainTransaction type: %s", data["__typename"])
	}
}
