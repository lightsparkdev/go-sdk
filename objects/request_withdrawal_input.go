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

	// IdempotencyKey The idempotency key of the request. The same result will be returned for the same idempotency key.
	IdempotencyKey *string `json:"request_withdrawal_input_idempotency_key"`

	// FeeTarget The target of the fee that should be used when crafting the L1 transaction. You should only set `fee_target` or `sats_per_vbyte`. If neither of them is set, default value of MEDIUM will be used as `fee_target`.
	FeeTarget *OnChainFeeTarget `json:"request_withdrawal_input_fee_target"`

	// SatsPerVbyte A manual fee rate set in sat/vbyte that should be used when crafting the L1 transaction. You should only set `fee_target` or `sats_per_vbyte`
	SatsPerVbyte *int64 `json:"request_withdrawal_input_sats_per_vbyte"`
}
