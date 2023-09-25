// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package uma

import (
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
}

func (l LightsparkClientUmaInvoiceCreator) CreateUmaInvoice(amountMsats int64, metadata string) (*string, error) {
	invoice, err := l.LightsparkClient.CreateUmaInvoice(l.NodeId, amountMsats, metadata, l.ExpirySecs)
	if err != nil {
		return nil, err
	}
	return &invoice.Data.EncodedPaymentRequest, nil
}
