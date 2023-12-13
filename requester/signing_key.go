package requester

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"errors"

	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
)

type SigningKey interface {
	Sign(payload []byte) ([]byte, error)
}

type Secp256k1SigningKey struct {
	MasterSeedBytes []byte
	Network         lightspark_crypto.BitcoinNetwork
}

func (s *Secp256k1SigningKey) Sign(payload []byte) ([]byte, error) {
	derivationPath := "m/5"
	key, error := lightspark_crypto.DerivePrivateKey(s.MasterSeedBytes, s.Network, derivationPath)
	if error != nil {
		return nil, error
	}
	keyBytes, error := hex.DecodeString(key)
	if error != nil {
		return nil, error
	}
	return lightspark_crypto.SignEcdsa(keyBytes, payload)
}

type RsaSigningKey struct {
	PrivateKey []byte
}

func (s *RsaSigningKey) Sign(payload []byte) ([]byte, error) {
	privateKey, err := x509.ParsePKCS8PrivateKey(s.PrivateKey)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is not an RSA key")
	}

	hashed := sha256.Sum256(payload)
	signature, err := rsa.SignPSS(rand.Reader, rsaKey, crypto.SHA256, hashed[:], nil)

	if err != nil {
		return nil, err
	}

	return signature, nil
}
