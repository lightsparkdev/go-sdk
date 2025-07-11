package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/lightsparkdev/go-sdk/objects"
)

func TestVerifyAndParse(t *testing.T) {
	tests := []struct {
		name          string
		data          string
		webhookSecret string
		want          *WebhookEvent
	}{
		{
			name: "payment finished",
			data: `{
				"event_type": "PAYMENT_FINISHED",
				"event_id": "test-event-123",
				"timestamp": "2025-01-01T12:00:00Z",
				"entity_id": "invoice-entity-456",
				"wallet_id": "wallet-789",
				"data": {
					"amount": 1000,
					"currency": "USD"
				}
			}`,
			webhookSecret: "test-secret-key",
			want: &WebhookEvent{
				EventType: objects.WebhookEventTypePaymentFinished,
				EventId:   "test-event-123",
				Timestamp: time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
				EntityId:  "invoice-entity-456",
				WalletId:  stringPtr("wallet-789"),
				Data:      &map[string]any{"amount": json.Number("1000"), "currency": "USD"},
			},
		},
		{
			name: "no wallet_id",
			data: `{
				"event_type": "WALLET_OUTGOING_PAYMENT_FINISHED",
				"event_id": "payment-event-456",
				"timestamp": "2025-01-02T15:30:00Z",
				"entity_id": "payment-entity-789",
				"data": {
					"status": "COMPLETED"
				}
			}`,
			webhookSecret: "test-secret-key",
			want: &WebhookEvent{
				EventType: objects.WebhookEventTypeWalletOutgoingPaymentFinished,
				EventId:   "payment-event-456",
				Timestamp: time.Date(2025, 1, 2, 15, 30, 0, 0, time.UTC),
				EntityId:  "payment-entity-789",
				WalletId:  nil,
				Data:      &map[string]any{"status": "COMPLETED"},
			},
		},
		{
			name: "empty data",
			data: `{
				"event_type": "NODE_STATUS",
				"event_id": "node-event-789",
				"timestamp": "2025-01-03T09:15:00Z",
				"entity_id": "node-entity-123",
				"wallet_id": "wallet-456"
			}`,
			webhookSecret: "test-secret-key",
			want: &WebhookEvent{
				EventType: objects.WebhookEventTypeNodeStatus,
				EventId:   "node-event-789",
				Timestamp: time.Date(2025, 1, 3, 9, 15, 0, 0, time.UTC),
				EntityId:  "node-entity-123",
				WalletId:  stringPtr("wallet-456"),
				Data:      nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hexDigest := hexHMAC(tt.webhookSecret, tt.data)

			result, err := VerifyAndParse([]byte(tt.data), hexDigest, tt.webhookSecret)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if diff := cmp.Diff(tt.want, result); diff != "" {
				t.Fatalf("WebhookEvent mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestVerifyAndParse_InvalidSignature_Errors(t *testing.T) {
	tests := []struct {
		name          string
		data          string
		hexdigest     string
		webhookSecret string
		wantErr       string
	}{
		{
			name: "invalid signature",
			data: `{
				"event_type": "PAYMENT_FINISHED",
				"event_id": "test-event-123",
				"timestamp": "2025-01-01T12:00:00Z",
				"entity_id": "invoice-entity-456"
			}`,
			hexdigest:     "a1b2c3d4e5f6",
			webhookSecret: "test-secret-key",
			wantErr:       "webhook message hash does not match signature",
		},
		{
			name: "malformed hex signature",
			data: `{
				"event_type": "PAYMENT_FINISHED",
				"event_id": "test-event-123",
				"timestamp": "2025-01-01T12:00:00Z",
				"entity_id": "invoice-entity-456"
			}`,
			hexdigest:     "not-a-valid-hex-string",
			webhookSecret: "test-secret-key",
			wantErr:       "invalid message signature format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyAndParse([]byte(tt.data), tt.hexdigest, tt.webhookSecret)

			if got != nil {
				t.Errorf("VerifyAndParse() got = %v, want nil", got)
			}
			if err == nil {
				t.Fatalf("Expected error but got none")
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("Expected error to contain '%s', but got '%s'", tt.wantErr, err.Error())
			}
		})
	}
}

func hexHMAC(secret, data string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func stringPtr(s string) *string {
	return &s
}
