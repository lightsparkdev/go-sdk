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

type OnChainTransactionUnmarshaler struct {
	Object OnChainTransaction
}

func (unmarshaler *OnChainTransactionUnmarshaler) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var t OnChainTransaction
	switch string(raw["__typename"]) {
	case `"ChannelClosingTransaction"`:
		var channelClosingTransaction ChannelClosingTransaction
		if err := json.Unmarshal(data, &channelClosingTransaction); err != nil {
			return err
		}
		t = &channelClosingTransaction
	case `"ChannelOpeningTransaction"`:
		var channelOpeningTransaction ChannelOpeningTransaction
		if err := json.Unmarshal(data, &channelOpeningTransaction); err != nil {
			return err
		}
		t = &channelOpeningTransaction
	case `"Deposit"`:
		var deposit Deposit
		if err := json.Unmarshal(data, &deposit); err != nil {
			return err
		}
		t = &deposit
	case `"Withdrawal"`:
		var withdrawal Withdrawal
		if err := json.Unmarshal(data, &withdrawal); err != nil {
			return err
		}
		t = &withdrawal

	default:
		return fmt.Errorf("unknown OnChainTransaction type: %s", raw["__typename"])
	}

	unmarshaler.Object = t
	return nil
}
