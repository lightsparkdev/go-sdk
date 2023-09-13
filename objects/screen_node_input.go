// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type ScreenNodeInput struct {
	Provider ComplianceProvider `json:"screen_node_input_provider"`

	NodePubkey string `json:"screen_node_input_node_pubkey"`
}
