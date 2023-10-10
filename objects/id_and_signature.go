// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type IdAndSignature struct {

	// Id The id of the message.
	Id string `json:"id_and_signature_id"`

	// Signature The signature of the message.
	Signature string `json:"id_and_signature_signature"`
}
