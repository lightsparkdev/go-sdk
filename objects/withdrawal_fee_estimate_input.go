
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type WithdrawalFeeEstimateInput struct {

    // NodeId The node from which you'd like to make the withdrawal.
    NodeId string `json:"withdrawal_fee_estimate_input_node_id"`

    // AmountSats The amount you want to withdraw from this node in Satoshis. Use the special value -1 to withdrawal all funds from this node.
    AmountSats int64 `json:"withdrawal_fee_estimate_input_amount_sats"`

    // WithdrawalMode The strategy that should be used to withdraw the funds from this node.
    WithdrawalMode WithdrawalMode `json:"withdrawal_fee_estimate_input_withdrawal_mode"`

}








