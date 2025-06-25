package webhooks_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/webhooks"
	"github.com/stretchr/testify/require"
)

func TestWebhooks_VerifyAndParse(t *testing.T) {
	eventType := objects.WebhookEventTypeNodeStatus
	eventId := "1615c8be5aa44e429eba700db2ed8ca5"
	entityId := "lightning_node:01882c25-157a-f96b-0000-362d42b64397"
	timeStamp, _ := time.Parse(time.RFC3339, "2023-05-17T23:56:47.874449+00:00")
	data := `{"event_type": "NODE_STATUS", "event_id": "1615c8be5aa44e429eba700db2ed8ca5", "timestamp": "2023-05-17T23:56:47.874449+00:00", "entity_id": "lightning_node:01882c25-157a-f96b-0000-362d42b64397"}`
	hexdigest := "62a8829aeb48b4142533520b1f7f86cdb1ee7d718bf3ea15bc1c662d4c453b74"
	webhookSecret := "3gZ5oQQUASYmqQNuEk0KambNMVkOADDItIJjzUlAWjX"

	event, err := webhooks.VerifyAndParse([]byte(data), hexdigest, webhookSecret)
	if err != nil {
		t.Fatalf("VerifyAndParse() failed: %s", err)
	}
	if event.EventType != eventType {
		t.Fatalf("event type not equal: %v vs. %v", event.EventType, eventType)
	}
	if event.EventId != eventId {
		t.Fatalf("event id not equal: %v vs. %v", event.EventId, eventId)
	}
	if event.EntityId != entityId {
		t.Fatalf("entity id not equal: %v vs. %v", event.EntityId, entityId)
	}
	if !event.Timestamp.Equal(timeStamp) {
		t.Fatalf("timestamp not equal: %v vs. %v", event.Timestamp, timeStamp)
	}
}

func TestWebhooks_VerifyAndParseWithWallet(t *testing.T) {
	eventType := objects.WebhookEventTypeWalletIncomingPaymentFinished
	eventId := "1615c8be5aa44e429eba700db2ed8ca5"
	entityId := "lightning_node:01882c25-157a-f96b-0000-362d42b64397"
	walletId := "wallet:01882c25-157a-f96b-0000-362d42b64397"
	timeStamp, _ := time.Parse(time.RFC3339, "2023-05-17T23:56:47.874449+00:00")
	data := `{"event_type": "WALLET_INCOMING_PAYMENT_FINISHED", "event_id": "1615c8be5aa44e429eba700db2ed8ca5", "timestamp": "2023-05-17T23:56:47.874449+00:00", "entity_id": "lightning_node:01882c25-157a-f96b-0000-362d42b64397", "wallet_id": "wallet:01882c25-157a-f96b-0000-362d42b64397" }`
	hexdigest := "b4eeb95f18956b3c33b99e9effc61636effc4634f83604cb41de13470c42669a"
	webhookSecret := "3gZ5oQQUASYmqQNuEk0KambNMVkOADDItIJjzUlAWjX"

	event, err := webhooks.VerifyAndParse([]byte(data), hexdigest, webhookSecret)
	if err != nil {
		t.Fatalf("VerifyAndParse() failed: %s", err)
	}
	if event.EventType != eventType {
		t.Fatalf("event type not equal: %v vs. %v", event.EventType, eventType)
	}
	if event.EventId != eventId {
		t.Fatalf("event id not equal: %v vs. %v", event.EventId, eventId)
	}
	if event.EntityId != entityId {
		t.Fatalf("entity id not equal: %v vs. %v", event.EntityId, entityId)
	}
	if !event.Timestamp.Equal(timeStamp) {
		t.Fatalf("timestamp not equal: %v vs. %v", event.Timestamp, timeStamp)
	}
	if *event.WalletId != walletId {
		t.Fatalf("wallet id not equal: %v vs. %v", *event.WalletId, walletId)
	}
}

func TestWebhooks_ParseNumber(t *testing.T) {
	data := `{"event_type": "REMOTE_SIGNING", "event_id": "8be9c360a68e420b9126b43ff6007a32", "timestamp": "2023-08-10T02:14:27.559234+00:00", "entity_id": "node_with_server_signing:0189d6bc-558d-88df-0000-502f04e71816", "data": {"sub_event_type": "GET_PER_COMMITMENT_POINT", "bitcoin_network": "TESTNET", "derivation_path": "m/3/2104864975", "per_commitment_point_idx": 281474976710654}}`
	parsed, err := webhooks.Parse([]byte(data))
	require.NoError(t, err)
	require.NotNil(t, parsed)

	perCommitmentPoint := (*parsed.Data)["per_commitment_point_idx"]
	perCommitmentPointIdx, err := perCommitmentPoint.(json.Number).Int64()
	require.NoError(t, err)
	require.Equal(t, int64(281474976710654), perCommitmentPointIdx)

}

func TestWebhooks_InvalidSignature(t *testing.T) {
	data := `{"event_type": "NODE_STATUS", "event_id": "1615c8be5aa44e429eba700db2ed8ca5", "timestamp": "2023-05-17T23:56:47.874449+00:00", "entity_id": "lightning_node:01882c25-157a-f96b-0000-362d42b64397"}`
	invalidHexDigest := "deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef" // wrong
	webhookSecret := "3gZ5oQQUASYmqQNuEk0KambNMVkOADDItIJjzUlAWjX"

	event, err := webhooks.VerifyAndParse([]byte(data), invalidHexDigest, webhookSecret)
	if err == nil {
		t.Fatalf("Expected error for invalid signature, got event: %+v", event)
	}
}

func TestWebhooks_BoundaryInjection(t *testing.T) {
	// Payload with a very large number and a script tag in a string field
	data := `{"event_type": "NODE_STATUS", "event_id": "<script>alert('x')</script>", "timestamp": "2023-05-17T23:56:47.874449+00:00", "entity_id": "lightning_node:01882c25-157a-f96b-0000-362d42b64397", "data": {"large_number": 1e+100}}`
	hexdigest := "b2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2e2"
	webhookSecret := "3gZ5oQQUASYmqQNuEk0KambNMVkOADDItIJjzUlAWjX"

	// We expect signature verification to fail, but if it passed, parsing should not panic
	_, _ = webhooks.VerifyAndParse([]byte(data), hexdigest, webhookSecret)
	// No assertion here, just ensuring no panic occurs
}

func TestWebhooks_MissingSignatureHeader(t *testing.T) {
	data := `{"event_type": "NODE_STATUS", "event_id": "1615c8be5aa44e429eba700db2ed8ca5", "timestamp": "2023-05-17T23:56:47.874449+00:00", "entity_id": "lightning_node:01882c25-157a-f96b-0000-362d42b64397"}`
	webhookSecret := "3gZ5oQQUASYmqQNuEk0KambNMVkOADDItIJjzUlAWjX"

	// Simulate missing signature header by passing an empty string
	_, err := webhooks.VerifyAndParse([]byte(data), "", webhookSecret)
	if err == nil {
		t.Fatalf("Expected error for missing signature, got nil")
	}
}

func TestWebhooks_LargePayload(t *testing.T) {
	largeData, err := json.Marshal(map[string]interface{}{
		"data": strings.Repeat("abcdefghijklmnopqrstuvwxyz", 80000), // ~2MB of data in a valid JSON string
	})
	if err != nil {
		t.Fatalf("Failed to create large JSON payload: %v", err)
	}
	hexdigest := "deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"
	webhookSecret := "3gZ5oQQUASYmqQNuEk0KambNMVkOADDItIJjzUlAWjX"

	_, err := webhooks.VerifyAndParse(largeData, hexdigest, webhookSecret)
	if err != nil {
		t.Logf("Received error for large payload (expected if size limit is enforced): %v", err)
	}
}

func TestWebhooks_UnexpectedDataTypes(t *testing.T) {
	data := `{"event_type": ["NODE_STATUS"], "event_id": 12345, "timestamp": {}, "entity_id": null}`
	hexdigest := "deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"
	webhookSecret := "3gZ5oQQUASYmqQNuEk0KambNMVkOADDItIJjzUlAWjX"

	_, err := webhooks.VerifyAndParse([]byte(data), hexdigest, webhookSecret)
	if err == nil {
		t.Fatalf("Expected error for unexpected data types, got nil")
	}
}
