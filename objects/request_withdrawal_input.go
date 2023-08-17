// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type RequestWithdrawalInput struct {

	// NodeId The node from which you'd like to make the withdrawal.
	NodeId string `json:"request_withdrawal_input_node_id"`

	// BitcoinAddress The bitcoin address where the withdrawal should be sent.
	BitcoinAddress string `json:"request_withdrawal_input_bitcoin_address"`

	// AmountSats The amount you want to withdraw from this node in Satoshis. Use the special value -1 to withdrawal all funds from this node.
	AmountSats int64 `json:"request_withdrawal_input_amount_sats"`

	// WithdrawalMode The strategy that should be used to withdraw the funds from this node.
	WithdrawalMode WithdrawalMode `json:"request_withdrawal_input_withdrawal_mode"`
}
