package remotesigning_test

import (
	"testing"
	"time"

	utils "github.com/lightsparkdev/go-sdk/keyscripts"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/remotesigning"
	"github.com/lightsparkdev/go-sdk/webhooks"
	"github.com/stretchr/testify/assert"
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
func TestSplitDerivationPath(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		wantHardened   []uint32
		wantRemaining  []uint32
		expectedErrMsg string
	}{
		{
			name:          "valid path with hardened and non-hardened components",
			path:          "m/84'/0'/0'/0/1",
			wantHardened:  []uint32{84 + 0x80000000, 0x80000000, 0x80000000},
			wantRemaining: []uint32{0, 1},
		},
		{
			name:           "path with empty component 1",
			path:           "m/",
			expectedErrMsg: "invalid derivation path: empty component",
		},
		{
			name:           "path with empty component 2",
			path:           "m//1/2",
			expectedErrMsg: "invalid derivation path: empty component",
		},
		{
			name:           "invalid number",
			path:           "m/84'/abc/0",
			expectedErrMsg: "invalid path: abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hardened, remaining, err := remotesigning.SplitDerivationPath(tt.path)

			if tt.expectedErrMsg != "" {
				assert.EqualError(t, err, tt.expectedErrMsg)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantHardened, hardened)
			assert.Equal(t, tt.wantRemaining, remaining)
		})
	}
}

func TestDeriveChildPubKeyFromExistingXPub(t *testing.T) {
	testXPub := "xpub6BosfCnifzxcFwrSzQiqu2DBVTshkCXacvNsWGYJVVhhawA7d4R5WSWGFNbi8Aw6ZRc1brxMyWMzG3DSSSSoekkudhUd9yLb6qx39T9nMdj"

	tests := []struct {
		name        string
		path        []uint32
		expectedLen int
	}{
		{
			name:        "valid derivation",
			path:        []uint32{0, 0},
			expectedLen: 33,
		},
		{
			name:        "empty path",
			path:        []uint32{},
			expectedLen: 33,
		},
		{
			name: "hardened index should fail",
			path: []uint32{0x80000000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pubkey, err := utils.DeriveChildPubKeyFromExistingXPub(testXPub, tt.path)

			if tt.expectedLen > 0 {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedLen, len(pubkey))
			} else {
				assert.Error(t, err)
			}
		})
	}
}
func TestValidateScript(t *testing.T) {
	// THIS IS A TEST XPUB - DO NOT USE IN PRODUCTION!
	testXPub := "xpub6CrnwQT4n7fEqLPG6A4KZcNXRctojRQGtvUztN5aKUqEMDU3ai5N9SvPnA56y5kwATN6CCHzmA7ccTwXbKtU7kZALRCVs1YY88987Ghv4jy"

	tests := []struct {
		name        string
		signingJob  *remotesigning.SigningJob
		xpub        string
		expectValid bool
	}{
		{
			name: "valid transaction and script",
			signingJob: &remotesigning.SigningJob{
				DerivationPath:            "m/84'/1'/0'/0/0",
				DestinationDerivationPath: "m/1/77",
				Transaction:               ptr("020000000001017ab44ffadf03b57ce0eb63074c541b3aea0b57497764a6790611332c441b989d0100000000ffffffff02f40100000000000016001427d703f9a06364bd45d122da1baea1e517a9ff1810500e0000000000160014b8a75a216b0a957b259d1b27049fdbaba42f950c020021021c902f59731d64721914f4826bc1868d2b3e5df40edbb3a8d7c4b21b95affb0400000000"),
			},
			xpub:        testXPub,
			expectValid: true,
		}, {
			name: "invalid transaction",
			signingJob: &remotesigning.SigningJob{
				DerivationPath:            "m/84'/1'/0'/0/0",
				DestinationDerivationPath: "m/1/77",
				Transaction:               ptr("abcd"),
			},
			xpub:        testXPub,
			expectValid: false,
		}, {
			name: "invalid derivation path",
			signingJob: &remotesigning.SigningJob{
				DerivationPath:            "m/84'/0/0",
				DestinationDerivationPath: "m/1/299",
				Transaction:               ptr("020000000001017ab44ffadf03b57ce0eb63074c541b3aea0b57497764a6790611332c441b989d0100000000ffffffff02f40100000000000016001427d703f9a06364bd45d122da1baea1e517a9ff1810500e0000000000160014b8a75a216b0a957b259d1b27049fdbaba42f950c020021021c902f59731d64721914f4826bc1868d2b3e5df40edbb3a8d7c4b21b95affb0400000000"),
			},
			xpub:        testXPub,
			expectValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid, err := remotesigning.ValidateScript(tt.signingJob, tt.xpub)
			if tt.expectValid {
				assert.NoError(t, err)
				assert.True(t, isValid)
			} else {
				assert.False(t, isValid)
			}
		})
	}
}

// Helper function to create string pointer
func ptr(s string) *string {
	return &s
}
