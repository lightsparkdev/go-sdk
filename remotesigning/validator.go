// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package remotesigning

import "github.com/lightsparkdev/go-sdk/webhooks"

// Validator an interface which decides whether to sign or reject a remote signing webhook event.
type Validator interface {
	ShouldSign(webhookEvent webhooks.WebhookEvent) bool
}

type PositiveValidator struct{}

func (v PositiveValidator) ShouldSign(webhooks.WebhookEvent) bool {
	return true
}
