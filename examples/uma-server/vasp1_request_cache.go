package main

import (
	"github.com/google/uuid"
	"github.com/lightsparkdev/go-sdk/objects"
	umaprotocol "github.com/uma-universal-money-address/uma-go-sdk/uma/protocol"
)

type Vasp1InitialRequestData struct {
	lnurlpResponse umaprotocol.LnurlpResponse
	receiverId     string
	vasp2Domain    string
}

type Vasp1PayReqData struct {
	encodedInvoice string
	utxoCallback   *string
	invoiceData    *objects.InvoiceData
}

type Vasp1RequestCache struct {
	// This is a map of the UMA request UUID to the LnurlpResponse from that initial Lnurlp request.
	// This is used to cache the LnurlpResponse so that we can use it to generate the UMA payreq without the client
	// having to make another Lnurlp request or remember lots of details.
	// NOTE: In production, this should be stored in a database or other persistent storage.
	umaRequestCache map[string]Vasp1InitialRequestData

	// This is a map of the UMA request UUID to the payreq data that we generated for that request.
	// This is used to cache the payreq data so that we can pay the invoice when the user confirms
	// NOTE: In production, this should be stored in a database or other persistent storage.
	payReqCache map[string]Vasp1PayReqData
}

func NewVasp1RequestCache() *Vasp1RequestCache {
	return &Vasp1RequestCache{
		umaRequestCache: make(map[string]Vasp1InitialRequestData),
		payReqCache:     make(map[string]Vasp1PayReqData),
	}
}

func (c *Vasp1RequestCache) GetLnurlpResponseData(uuid string) (Vasp1InitialRequestData, bool) {
	lnurlpResponse, ok := c.umaRequestCache[uuid]
	return lnurlpResponse, ok
}

func (c *Vasp1RequestCache) SaveLnurlpResponseData(lnurlpResponse umaprotocol.LnurlpResponse, receiverId string, vasp2Domain string) string {
	// Generate a UUID for this request
	requestUUID := uuid.New().String()
	c.umaRequestCache[requestUUID] = Vasp1InitialRequestData{
		lnurlpResponse: lnurlpResponse,
		receiverId:     receiverId,
		vasp2Domain:    vasp2Domain,
	}
	return requestUUID
}

func (c *Vasp1RequestCache) DeleteLnurlpResponse(uuid string) {
	delete(c.umaRequestCache, uuid)
}

func (c *Vasp1RequestCache) GetPayReqData(uuid string) (Vasp1PayReqData, bool) {
	payReqData, ok := c.payReqCache[uuid]
	return payReqData, ok
}

func (c *Vasp1RequestCache) SavePayReqData(
	requestUUID string,
	encodedInvoice string,
	utxoCallback *string,
	invoiceData *objects.InvoiceData,
) string {
	c.payReqCache[requestUUID] = Vasp1PayReqData{
		encodedInvoice: encodedInvoice,
		utxoCallback:   utxoCallback,
		invoiceData:    invoiceData,
	}
	return requestUUID
}

func (c *Vasp1RequestCache) DeletePayReqData(uuid string) {
	delete(c.payReqCache, uuid)
}
