package uma

import "fmt"

type PayerDataOptions struct {
	NameRequired       bool
	EmailRequired      bool
	ComplianceRequired bool
}

func (p PayerDataOptions) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
		"identifier": { "mandatory": true },
		"name": { "mandatory": %t },
		"email": { "mandatory": %t },
		"compliance": { "mandatory": %t }
	}`, p.NameRequired, p.EmailRequired, p.ComplianceRequired)), nil
}

type PayerData struct {
	Name       *string              `json:"name"`
	Email      *string              `json:"email"`
	Identifier string               `json:"identifier"`
	Compliance *CompliancePayerData `json:"compliance"`
}

type CompliancePayerData struct {
	// Utxos is the list of UTXOs of the sender's channels that might be used to fund the payment.
	Utxos *[]string `json:"utxos"`
	// IsKYCd indicates whether VASP1 has KYC information about the sender.
	IsKYCd bool `json:"isKYCd"`
	// TrInfo is the travel rule information of the sender. This is encrypted with the receiver's public encryption key.
	TrInfo *string `json:"trInfo"`
	// Signature is the base64-encoded signature of sha256(ReceiverAddress|Nonce|Timestamp).
	Signature          string `json:"signature"`
	SignatureNonce     string `json:"signatureNonce"`
	SignatureTimestamp int64  `json:"signatureTimestamp"`
	// UtxoCallback is the URL that the receiver will call to send UTXOs of the channel that the receiver used to receive the payment once it completes.
	UtxoCallback string `json:"utxoCallback"`
}
