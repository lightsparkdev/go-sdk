package remotesigning_test

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
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
func TestDerivationPath(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedPath   []uint32
		expectedErrMsg string
	}{
		{
			name:         "valid path with hardened and non-hardened components",
			path:         "m/84'/0'/0'/0/1",
			expectedPath: []uint32{84 + 0x80000000, 0x80000000, 0x80000000, 0, 1},
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
			path, err := remotesigning.DerivationPathFromString(tt.path)

			if tt.expectedErrMsg != "" {
				assert.EqualError(t, err, tt.expectedErrMsg)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedPath, path)
		})
	}
}

func TestL1WalletDerivationPath(t *testing.T) {
	masterSeed, err := hex.DecodeString("69f580170954f411bbabc60118c0a3a0e483381d196d4087d32b78bdfee4a114")
	assert.NoError(t, err)

	tests := []struct {
		name           string
		network        chaincfg.Params
		expectedPath   string
		expectedKey    string
		expectedErrMsg string
	}{
		{
			name:         "mainnet",
			network:      chaincfg.MainNetParams,
			expectedPath: "m/84'/0'/0'",
			expectedKey:  "xpub6D87MELKGzkxrhQcQVNddwHLC3td1ftP4sS6zBaz5JU4MimSwKqK3XcEnrE38mgfydBmsedXc35tETx5HALSzBfre3ZXo26nS2kBmPwvJ3n",
		},
		{
			name:         "testnet",
			network:      chaincfg.TestNet3Params,
			expectedPath: "m/84'/1'/0'",
			expectedKey:  "tpubDDPLaqgr8bawQ6chGQ7ZkSChHBjVihms7dXdcZF8XvDNirxsctoK17brjJt7eMb7ZgppHz86uQNT1ksw89svDiAqNeMKsHfYfs8K8F1kq8m",
		},
		{
			name:         "regtest",
			network:      chaincfg.RegressionNetParams,
			expectedPath: "m/84'/2'/0'",
			expectedKey:  "tpubDC7fZywQZQ45q3ebc3HC2CiCWUe1p3g5ZrM4uh7GBXRqnBzpG5qLD8swyYqUThvmNksGLEHYcjtChXUXGUWXuH1FqTF6rwwr2ErEoZJQnE3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := remotesigning.L1WalletDerivationPrefix(&tt.network)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedPath, path)

			key, err := remotesigning.DeriveL1WalletHardenedXpub(masterSeed, &tt.network)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedKey, key.String())
		})
	}
}

func TestValidTransaction(t *testing.T) {
	tests := []struct {
		name        string
		signingJob  *remotesigning.SigningJob
		expectValid bool
	}{
		{
			name: "valid transaction",
			signingJob: &remotesigning.SigningJob{
				Transaction: ptr("02000000017ab44ffadf03b57ce0eb63074c541b3aea0b57497764a6790611332c441b989d0100000000ffffffff02a086010000000000160014aff40d81f6ffd5a98e358af465b1e1bf3fe9c012a086010000000000160014dd71b57f94e6876380850d0fbbaedb52d698b9e000000000"),
			},
			expectValid: true,
		}, {
			name: "invalid transaction",
			signingJob: &remotesigning.SigningJob{
				Transaction: ptr("abcd"),
			},
			expectValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.signingJob.BitcoinTx()
			if tt.expectValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestValidateChangeScript(t *testing.T) {
	testMasterSeed, err := hex.DecodeString("370eb72fc3dd38c74f933b477378d51c3b5e6db126ad64fa3244b16e2b8a1bd37f6454dbdb2d58b98f3f6fa2d1c9232216b67616eb2c61bf8c2abe8f67edf252")
	assert.NoError(t, err)

	tests := []struct {
		name        string
		signingJob  *remotesigning.SigningJob
		masterSeed  []byte
		expectValid bool
	}{
		{
			name: "valid transaction and script",
			signingJob: &remotesigning.SigningJob{
				DerivationPath:            "m/84'/1'/0'/0/0",
				DestinationDerivationPath: "m/1/77",
				Transaction:               ptr("02000000017ab44ffadf03b57ce0eb63074c541b3aea0b57497764a6790611332c441b989d0100000000ffffffff02a086010000000000160014aff40d81f6ffd5a98e358af465b1e1bf3fe9c012a086010000000000160014dd71b57f94e6876380850d0fbbaedb52d698b9e000000000"),
			},
			masterSeed:  testMasterSeed,
			expectValid: true,
		}, {
			name: "invalid derivation path",
			signingJob: &remotesigning.SigningJob{
				DerivationPath:            "m/84'/0/0",
				DestinationDerivationPath: "m/1/299",
				Transaction:               ptr("02000000017ab44ffadf03b57ce0eb63074c541b3aea0b57497764a6790611332c441b989d0100000000ffffffff02a086010000000000160014aff40d81f6ffd5a98e358af465b1e1bf3fe9c012a086010000000000160014dd71b57f94e6876380850d0fbbaedb52d698b9e000000000"),
			},
			masterSeed:  testMasterSeed,
			expectValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publicKey, err := remotesigning.DerivePublicKey(tt.masterSeed, tt.signingJob.DestinationDerivationPath, &chaincfg.MainNetParams)
			assert.NoError(t, err)
			script, err := remotesigning.GenerateP2WPKHFromPubkey(publicKey.SerializeCompressed())
			assert.NoError(t, err)
			tx, err := tt.signingJob.BitcoinTx()
			assert.NoError(t, err)
			isValid, err := remotesigning.ValidateChangeScript(tx, script)
			if tt.expectValid {
				assert.NoError(t, err)
				assert.True(t, isValid)
			} else {
				assert.False(t, isValid)
			}
		})
	}
}

func TestValidateOutputScript(t *testing.T) {
	testMasterSeed, err := hex.DecodeString("370eb72fc3dd38c74f933b477378d51c3b5e6db126ad64fa3244b16e2b8a1bd37f6454dbdb2d58b98f3f6fa2d1c9232216b67616eb2c61bf8c2abe8f67edf252")
	assert.NoError(t, err)

	tests := []struct {
		name        string
		signingJob  *remotesigning.SigningJob
		masterSeed  []byte
		expectValid bool
	}{
		{
			name: "valid transaction and script",
			signingJob: &remotesigning.SigningJob{
				DerivationPath:            "m/84'/1'/0'/0/0",
				DestinationDerivationPath: "m/1/12",
				Transaction:               ptr("02000000017ab44ffadf03b57ce0eb63074c541b3aea0b57497764a6790611332c441b989d0100000000ffffffff02a0860100000000001600147a0770afcc2afeba0cc9f01c52acc77b5b82f387a086010000000000160014aff40d81f6ffd5a98e358af465b1e1bf3fe9c01200000000"),
			},
			masterSeed:  testMasterSeed,
			expectValid: true,
		}, {
			name: "invalid derivation path",
			signingJob: &remotesigning.SigningJob{
				DerivationPath:            "m/84'/0/0",
				DestinationDerivationPath: "m/1/299",
				Transaction:               ptr("02000000017ab44ffadf03b57ce0eb63074c541b3aea0b57497764a6790611332c441b989d0100000000ffffffff02a0860100000000001600147a0770afcc2afeba0cc9f01c52acc77b5b82f387a086010000000000160014aff40d81f6ffd5a98e358af465b1e1bf3fe9c01200000000"),
			},
			masterSeed:  testMasterSeed,
			expectValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publicKey, err := remotesigning.DerivePublicKey(tt.masterSeed, tt.signingJob.DestinationDerivationPath, &chaincfg.MainNetParams)
			assert.NoError(t, err)
			script, err := remotesigning.GenerateP2WPKHFromPubkey(publicKey.SerializeCompressed())
			assert.NoError(t, err)
			tx, err := tt.signingJob.BitcoinTx()
			assert.NoError(t, err)
			isValid, err := remotesigning.ValidateOutputScript(tx, script)
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
