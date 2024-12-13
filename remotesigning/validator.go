// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package remotesigning

import (
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/lightsparkdev/go-sdk/webhooks"
)

// Validator an interface which decides whether to sign or reject a remote signing webhook event.
type Validator interface {
	ShouldSign(webhook webhooks.WebhookEvent) bool
}

type PositiveValidator struct{}

func (v PositiveValidator) ShouldSign(webhook webhooks.WebhookEvent) bool {
	return true
}

type MultiValidator struct{
	validators []Validator
}

func NewMultiValidator(validators ...Validator) MultiValidator {
	return MultiValidator{validators: validators}
}

func (v MultiValidator) ShouldSign(webhookEvent webhooks.WebhookEvent) bool {
	for _, validator := range v.validators {
		if !validator.ShouldSign(webhookEvent) {
			return false
		}
	}
	return true
}

type HashValidator struct{}

func (v HashValidator) ShouldSign(webhookEvent webhooks.WebhookEvent) bool {
	request, err := ParseDeriveAndSignRequest(webhookEvent)
	if err != nil {
		// Only validate DeriveAndSignRequest events
		return true
	}
	for _, signing := range request.SigningJobs {
		if !ValidateWitnessHash(&signing) {
			return false
		}
	}
	return true
}

func ValidateWitnessHash(signing *SigningJob) bool {
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

// A validator that checks that the outputs of a transaction pass restrictions
// for where we allow sending funds. This is used to ensure that when signing
// transactions spending L1 wallet funds, we are only sending funds to certain
// addresses.
type DestinationValidator struct{
	masterSeed []byte
}

func NewDestinationValidator(masterSeed []byte) DestinationValidator {
	return DestinationValidator{masterSeed: masterSeed}
}

func (v DestinationValidator) ShouldSign(webhookEvent webhooks.WebhookEvent) bool {
	request, err := ParseDeriveAndSignRequest(webhookEvent)
	if err != nil {
		// Only validate DeriveAndSignRequest events
		return true
	}
	for _, signing := range request.SigningJobs {
		if strings.HasPrefix(signing.DerivationPath, "m/84") {
			publicKey, err := DerivePublicKey(v.masterSeed, signing.DerivationPath, &chaincfg.MainNetParams)
			if err != nil {
				return false
			}
			validScript, err := ValidateScript(&signing, publicKey)
			if err != nil || !validScript {
				return false
			}
		}
	}
	return true
}
