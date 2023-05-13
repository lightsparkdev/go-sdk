// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
)

type OnChainTransaction interface {
	Transaction
	Entity

	// The fees that were paid by the wallet sending the transaction to commit it to the Bitcoin blockchain.
	GetFees() *CurrencyAmount

	// The hash of the block that included this transaction. This will be null for unconfirmed transactions.
	GetBlockHash() *string

	// The height of the block that included this transaction. This will be zero for unconfirmed transactions.
	GetBlockHeight() int64

	// The Bitcoin blockchain addresses this transaction was sent to.
	GetDestinationAddresses() []string

	// The number of blockchain confirmations for this transaction in real time.
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
