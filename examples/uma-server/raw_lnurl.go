package main

type NonUmaLnurlpResponse struct {
	Tag         string `json:"tag"`
	Callback    string `json:"callback"`
	MinSendable int64  `json:"minSendable"`
	MaxSendable int64  `json:"maxSendable"`
	Metadata    string `json:"metadata"`
}

type NonUmaPayReqResponse struct {
	// EncodedInvoice is the BOLT11 invoice that the sender will pay.
	EncodedInvoice string `json:"pr"`
}
