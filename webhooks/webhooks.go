// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package webhooks

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
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
	Data      *map[string]interface{}
}

// VerifyAndParse Verifies the signature and parses the message into a WebhookEvent object.
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

	headerBytes, err := hex.DecodeString(hexdigest)
	if err != nil {
		return nil, errors.New("invalid message signature format")
	}
	if !hmac.Equal(result, headerBytes) {
		return nil, errors.New("webhook message hash does not match signature")
	}

	return Parse(data)
}

// Parse Parses the message into a WebhookEvent object.
//
// Args:
//
//	data: the POST message body received by the webhook.
func Parse(data []byte) (*WebhookEvent, error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()

	var eventJSON map[string]interface{}
	if err := dec.Decode(&eventJSON); err != nil {
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

	var additionalData *map[string]interface{} = nil
	if eventJSON["data"] != nil {
		additionalData = new(map[string]interface{})
		*additionalData = eventJSON["data"].(map[string]interface{})
	}

	return &WebhookEvent{
		EventType: eventType,
		EventId:   eventJSON["event_id"].(string),
		Timestamp: timestamp,
		EntityId:  eventJSON["entity_id"].(string),
		WalletId:  walletId,
		Data:      additionalData,
	}, nil
}
