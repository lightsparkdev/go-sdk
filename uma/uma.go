// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package uma

import (
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
)

// LightsparkClientUmaInvoiceCreator is a wrapper around the LightsparkClient that implements the UmaInvoiceCreator
// interface.
// See github.com/uma-universal-money-address/uma-go-sdk for the interface and its documentation.
type LightsparkClientUmaInvoiceCreator struct {
	LightsparkClient services.LightsparkClient
	// NodeId: the node ID of the receiver.
	NodeId string
	// ExpirySecs: the number of seconds until the invoice expires.
	ExpirySecs *int32
	// EnableUmaAnalytics: A flag indicating whether UMA analytics should be enabled. If `true`,
	// the receiver identifier will be hashed using a monthly-rotated seed and used for anonymized
	// analysis.
	EnableUmaAnalytics bool
	// SigningPrivateKey: Optional, the receiver's signing private key. Used to hash the receiver
	// identifier if UMA analytics is enabled.
	SigningPrivateKey *[]byte
}

func (l LightsparkClientUmaInvoiceCreator) CreateInvoice(amountMsats int64, metadata string, receiverIdentifier *string) (*string, error) {
	var invoice *objects.Invoice
	var err error
	if l.EnableUmaAnalytics && l.SigningPrivateKey != nil {
		invoice, err = l.LightsparkClient.CreateUmaInvoiceWithReceiverIdentifier(l.NodeId, amountMsats, metadata, l.ExpirySecs, l.SigningPrivateKey, receiverIdentifier)
	} else {
		invoice, err = l.LightsparkClient.CreateUmaInvoice(l.NodeId, amountMsats, metadata, l.ExpirySecs)
	}
	if err != nil {
		return nil, err
	}
	return &invoice.Data.EncodedPaymentRequest, nil
}
