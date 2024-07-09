package remotesigning_test

import (
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/webhooks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/lightsparkdev/go-sdk/remotesigning"
)

func TestGetPaymentHashFromScript(t *testing.T) {
	scriptHex := "76a9148610bc927c248b7b1542e02fca750934375dcb0f8763ac672103cdf1a3db51f4de23eb1a2a254605957525d8aca3787163afb1db55ded526a4dc7c820120876475527c2103b72a4c31af6cf9601bb25355522e8fa4f6904d28ff6f6e4e552998349acd92ef52ae67a9140b6ae907cf1ed5e44a49ebadf31c623faa057b5888ac6851b27568"
	paymentHash, err := remotesigning.GetPaymentHashFromScript(scriptHex)
	if err != nil {
		t.Fatalf("GetPaymentHashFromScript() failed: %s", err)
	}
	if *paymentHash != "0b6ae907cf1ed5e44a49ebadf31c623faa057b58" {
		t.Fatalf("payment hash not equal: %v vs. %v", *paymentHash, "0b6ae907cf1ed5e44a49ebadf31c623faa057b58")
	}
}

func TestParseReleasePaymentPreimage(t *testing.T) {
	webhookEvent := webhooks.WebhookEvent{
		EventType: objects.WebhookEventTypeRemoteSigning,
		EventId:   "event-id",
		Timestamp: time.Now(),
		EntityId:  "Node:node-id",
		WalletId:  nil,
		Data: &map[string]interface{}{
			"sub_event_type":  objects.RemoteSigningSubEventTypeReleasePaymentPreimage.StringValue(),
			"invoice_id":      "invoice-id",
			"bitcoin_network": "MAINNET",
			"is_uma":          true,
		},
	}

	parsedRequest, err := remotesigning.ParseReleasePaymentPreimageRequest(webhookEvent)
	assert.NoError(t, err)
	assert.Equal(t, "invoice-id", parsedRequest.InvoiceId)
	assert.True(t, parsedRequest.IsUma)
	assert.False(t, parsedRequest.IsLnurl)
}
