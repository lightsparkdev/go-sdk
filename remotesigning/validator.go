// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package remotesigning

import (
	"strings"

	"github.com/lightsparkdev/go-sdk/webhooks"
)

// Validator an interface which decides whether to sign or reject a remote signing webhook event.
type Validator interface {
	ShouldSign(webhook webhooks.WebhookEvent, xpubs []string) bool
}

type PositiveValidator struct{}

func (v PositiveValidator) ShouldSign(webhook webhooks.WebhookEvent, xpubs []string) bool {
	return true
}

type HashValidator struct{}

func (v HashValidator) ShouldSign(webhookEvent webhooks.WebhookEvent, xpubs []string) bool {
	request, err := ParseDeriveAndSignRequest(webhookEvent)
	if err != nil {
		// Only validate DeriveAndSignRequest events
		return true
	}
	for i, signing := range request.SigningJobs {
		if strings.HasPrefix(signing.DerivationPath, "m/84") {
			if !ValidateL1Transaction(&signing, xpubs[i]) {
				return false
			}
		} else {
			if !ValidateLightningTransaction(&signing) {
				return false
			}
		}
	}
	return true
}

func ValidateLightningTransaction(signing *SigningJob) bool {
	if strings.HasSuffix(signing.DerivationPath, "/2") || strings.HasSuffix(signing.DerivationPath, "/3") {
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
	return true
}

func ValidateL1Transaction(signing *SigningJob, xpub string) bool {
	// 1. Address Validation
	isValid, err := ValidateScript(signing, xpub)
	if err != nil {
		return false
	}
	if !isValid {
		return false
	}

	// 2. Witness Hash Validation
	msg, err := CalculateWitnessHash(*signing.Amount, *signing.Script, *signing.Transaction)
	if err != nil {
		return false
	}
	if strings.Compare(*msg, signing.Message) != 0 {
		return false
	}
	return true
}
