package webhooks_test

import (
	"testing"
	"time"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/webhooks"
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
	if event.Timestamp != timeStamp {
		t.Fatalf("timestamp not equal: %v vs. %v", event.Timestamp, timeStamp)
	}
}