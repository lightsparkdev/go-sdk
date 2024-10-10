// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type ScreenNodeInput struct {

	// Provider The compliance provider that is going to screen the node. You need to be a customer of the selected provider and store the API key on the Lightspark account setting page.
	Provider ComplianceProvider `json:"screen_node_input_provider"`

	// NodePubkey The public key of the lightning node that needs to be screened.
	NodePubkey string `json:"screen_node_input_node_pubkey"`
}
