// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type FundNodeInput struct {
	NodeId string `json:"fund_node_input_node_id"`

	AmountSats *int64 `json:"fund_node_input_amount_sats"`
}
