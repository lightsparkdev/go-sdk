package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestPreimage(t *testing.T) {
	seedHex := "000102030405060708090a0b0c0d0e0f"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		t.Fatalf("failed to decode seed: %v", err)
	}

	nonceHex := "5c3c1200b86db0eacbf1cbdd40b86ee5d66482b1f214ef26624aa407829a1a5b"
	nonce, err := hex.DecodeString(nonceHex)
	if err != nil {
		t.Fatalf("failed to decode nonce: %v", err)
	}

	key, err := DerivePreimageBaseKey(seed)
	if err != nil {
		t.Fatalf("failed to derive preimage base key: %v", err)
	}

	preimage, paymentHash, err := GeneratePreimageAndPaymentHash(key, nonce)
	if err != nil {
		t.Fatalf("failed to generate preimage and payment hash: %v", err)
	}

	expectedPreimageHex := "a29ce018c174a1a8a68a0acb2acb1111aa1625f486ca1a36d9d91863d5f26a6a"
	expectedPreimage, err := hex.DecodeString(expectedPreimageHex)
	if err != nil {
		t.Fatalf("failed to decode preimage: %v", err)
	}

	if !bytes.Equal(preimage, expectedPreimage) {
		t.Fatalf("preimage does not match expected preimage")
	}

	expectedPaymentHashHex := "2e28491c129551c7bed5f1df4852d2a711965c0d6ee1948e43aa639ddcd154c3"
	expectedPaymentHash, err := hex.DecodeString(expectedPaymentHashHex)
	if err != nil {
		t.Fatalf("failed to decode payment hash: %v", err)
	}

	if !bytes.Equal(paymentHash, expectedPaymentHash) {
		t.Fatalf("payment hash does not match expected payment hash")
	}
}
