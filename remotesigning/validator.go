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

type MultiValidator struct {
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

func isL1WalletSigningJob(job SigningJob) bool {
	return strings.HasPrefix(job.DerivationPath, "m/84")
}

func isForceClosureClaimSigningJob(job SigningJob) bool {
	return !isL1WalletSigningJob(job) &&
		(strings.HasSuffix(job.DerivationPath, "/2") ||
			strings.HasSuffix(job.DerivationPath, "/3"))
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
	if isForceClosureClaimSigningJob(*signing) {
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
type DestinationValidator struct {
	masterSeed                 []byte
	validateForceClosureClaims bool
}

func NewDestinationValidator(masterSeed []byte, validateForceClosureClaims bool) DestinationValidator {
	return DestinationValidator{masterSeed, validateForceClosureClaims}
}

func (v DestinationValidator) ShouldSign(webhookEvent webhooks.WebhookEvent) bool {
	request, err := ParseDeriveAndSignRequest(webhookEvent)
	if err != nil {
		// Only validate DeriveAndSignRequest events
		return true
	}
	for _, signing := range request.SigningJobs {
		l1SigningJob := isL1WalletSigningJob(signing)
		forceClosureClaim := v.validateForceClosureClaims &&
			isForceClosureClaimSigningJob(signing)
		if !l1SigningJob && !forceClosureClaim {
			continue
		}

		tx, err := signing.BitcoinTx()
		if err != nil {
			return false
		}
		publicKey, err := DerivePublicKey(v.masterSeed, signing.DestinationDerivationPath, &chaincfg.MainNetParams)
		if err != nil {
			return false
		}
		script, err := GenerateP2WPKHFromPubkey(publicKey.SerializeCompressed())
		if err != nil {
			return false
		}

		if isL1WalletSigningJob(signing) {
			validScript, err := ValidateChangeScript(tx, script)
			if err != nil || !validScript {
				return false
			}
		}
		if forceClosureClaim {
			validScript, err := ValidateOutputScript(tx, script)
			if err != nil || !validScript {
				return false
			}
		}
	}
	return true
}
