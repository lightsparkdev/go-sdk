// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package remotesigning

import (
	"strings"

	"github.com/lightsparkdev/go-sdk/webhooks"
)

// Validator an interface which decides whether to sign or reject a remote signing webhook event.
type Validator interface {
	ShouldSign(webhookEvent webhooks.WebhookEvent) bool
}

type PositiveValidator struct{}

func (v PositiveValidator) ShouldSign(webhooks.WebhookEvent) bool {
	return true
}

type HashValidator struct {} 

func (v HashValidator) ShouldSign(webhookEvent webhooks.WebhookEvent) bool {
	request, err := ParseDeriveAndSignRequest(webhookEvent)	
	if err != nil {
		// Only validate DeriveAndSignRequest events
		return true
	}

	for _, signing := range request.SigningJobs {
		if strings.HasSuffix(signing.DerivationPath, "/4") {
			msg, err := CalculateWitnessHashPSBT(*signing.Transaction)
			if err != nil {
				return false
			}
			if strings.Compare(*msg, signing.Message) != 0 {
				return false
			}
		} else {
			msg, err := CalculateWitnessHash(*signing.Amount, *signing.Script, *signing.Transaction)
			if err != nil {
				return false
			}
			if strings.Compare(*msg, signing.Message) != 0 {
				return false
			}
		}
	}

	return true
}