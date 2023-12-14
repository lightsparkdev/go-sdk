package requester

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"errors"

	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
)

type SigningKey interface {
	Sign(payload []byte) ([]byte, error)
}

type Secp256k1SigningKey struct {
	PrivateKey []byte
}

func (s *Secp256k1SigningKey) Sign(payload []byte) ([]byte, error) {
	return lightspark_crypto.SignEcdsa(payload, s.PrivateKey)
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
