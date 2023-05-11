// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

type WebhookEventType int

const (
	WebhookEventTypeUndefined WebhookEventType = iota

	WebhookEventTypePaymentFinished

	WebhookEventTypeNodeStatus

	WebhookEventTypeWalletStatus

	WebhookEventTypeWalletOutgoingPaymentFinished

	WebhookEventTypeWalletIncomingPaymentFinished

	WebhookEventTypeWalletWithdrawalFinished

	WebhookEventTypeWalletFundsReceived
)

func (a *WebhookEventType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = WebhookEventTypeUndefined
	case "PAYMENT_FINISHED":
		*a = WebhookEventTypePaymentFinished
	case "NODE_STATUS":
		*a = WebhookEventTypeNodeStatus
	case "WALLET_STATUS":
		*a = WebhookEventTypeWalletStatus
	case "WALLET_OUTGOING_PAYMENT_FINISHED":
		*a = WebhookEventTypeWalletOutgoingPaymentFinished
	case "WALLET_INCOMING_PAYMENT_FINISHED":
		*a = WebhookEventTypeWalletIncomingPaymentFinished
	case "WALLET_WITHDRAWAL_FINISHED":
		*a = WebhookEventTypeWalletWithdrawalFinished
	case "WALLET_FUNDS_RECEIVED":
		*a = WebhookEventTypeWalletFundsReceived

	}
	return nil
}

func (a WebhookEventType) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case WebhookEventTypePaymentFinished:
		s = "PAYMENT_FINISHED"
	case WebhookEventTypeNodeStatus:
		s = "NODE_STATUS"
	case WebhookEventTypeWalletStatus:
		s = "WALLET_STATUS"
	case WebhookEventTypeWalletOutgoingPaymentFinished:
		s = "WALLET_OUTGOING_PAYMENT_FINISHED"
	case WebhookEventTypeWalletIncomingPaymentFinished:
		s = "WALLET_INCOMING_PAYMENT_FINISHED"
	case WebhookEventTypeWalletWithdrawalFinished:
		s = "WALLET_WITHDRAWAL_FINISHED"
	case WebhookEventTypeWalletFundsReceived:
		s = "WALLET_FUNDS_RECEIVED"

	}
	return s
}

func (a WebhookEventType) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
