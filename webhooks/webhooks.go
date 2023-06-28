// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/lightsparkdev/go-sdk/objects"
)

const SIGNATURE_HEADER = "lightspark-signature"

type WebhookEvent struct {
	EventType objects.WebhookEventType
	EventId   string
	Timestamp time.Time
	EntityId  string
	WalletId  *string
}

// Verifies the signature and parses the message into a WebhookEvent object.
//
// Args:
//
//	data: the POST message body received by the webhook.
//	hexdigest: the message signature sent in the `lightspark-signature` header.
//	webhookSecret: the webhook secret configured at the Lightspark API configuration.
func VerifyAndParse(data []byte, hexdigest string, webhookSecret string) (*WebhookEvent, error) {
	hash := hmac.New(sha256.New, []byte(webhookSecret))
	hash.Write(data)
	result := hash.Sum(nil)
	if strings.ToLower(hex.EncodeToString(result)) != strings.ToLower(hexdigest) {
		return nil, errors.New("Webhook message hash does not match signature")
	}
	return Parse(data)
}

// Parses the message into a WebhookEvent object.
//
// Args:
//
//	data: the POST message body received by the webhook.
func Parse(data []byte) (*WebhookEvent, error) {
	var eventJSON map[string]interface{}
	err := json.Unmarshal(data, &eventJSON)
	if err != nil {
		return nil, err
	}

	eventBytes, err := json.Marshal(eventJSON["event_type"].(string))
	if err != nil {
		return nil, err
	}
	var eventType objects.WebhookEventType
	eventType.UnmarshalJSON(eventBytes)

	timestamp, err := time.Parse(time.RFC3339, eventJSON["timestamp"].(string))
	if err != nil {
		return nil, err
	}
	var walletId *string = nil
	if eventJSON["wallet_id"] != nil {
		walletId = new(string)
		*walletId = eventJSON["wallet_id"].(string)
	}

	return &WebhookEvent{
		EventType: eventType,
		EventId:   eventJSON["event_id"].(string),
		Timestamp: timestamp,
		EntityId:  eventJSON["entity_id"].(string),
		WalletId:  walletId,
	}, nil
}
