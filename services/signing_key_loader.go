package services

import (
	"errors"

	"github.com/lightsparkdev/go-sdk/crypto"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/requester"
	"github.com/lightsparkdev/go-sdk/scripts"
)

type SigningKeyLoader struct {
	cachedSigningKey     requester.SigningKey
	masterSeedAndNetwork *masterSeedAndNetwork
	idPasswordPair       *idPasswordPair
}

// NewSigningKeyLoaderFromNodeIdAndPassword creates a new SigningKeyLoader from a node ID and password.
// This cannot be used if you are using remote signing. It is used to recover an RSA operation signing key using
// the password you chose when setting up your node. For REGTEST nodes, the password is "1234!@#$".
func NewSigningKeyLoaderFromNodeIdAndPassword(nodeId string, password string) *SigningKeyLoader {
	return &SigningKeyLoader{idPasswordPair: &idPasswordPair{nodeId: nodeId, password: password}}
}

// NewSigningKeyLoaderFromRsaPrivateKey creates a new SigningKeyLoader from a raw RSA private key.
func NewSigningKeyLoaderFromRsaPrivateKey(rsaPrivateKeyBytes []byte) *SigningKeyLoader {
	return &SigningKeyLoader{cachedSigningKey: &requester.RsaSigningKey{PrivateKey: rsaPrivateKeyBytes}}
}

// NewSigningKeyLoaderFromSignerMasterSeed creates a new SigningKeyLoader from a master seed and network.
// This should be used if you are using remote signing, rather than an RSA operation signing key.
func NewSigningKeyLoaderFromSignerMasterSeed(masterSeedBytes []byte, network objects.BitcoinNetwork) *SigningKeyLoader {
	return &SigningKeyLoader{
		masterSeedAndNetwork: &masterSeedAndNetwork{masterSeed: masterSeedBytes, network: network},
	}
}

func (s *SigningKeyLoader) LoadSigningKey(req requester.Requester) (requester.SigningKey, error) {
	if s.cachedSigningKey != nil {
		return s.cachedSigningKey, nil
	}

	if s.masterSeedAndNetwork != nil {
		key, err := s.loadSigningKeyFromMasterSeed()
		if err != nil {
			return nil, err
		}
		s.cachedSigningKey = key
		return key, nil
	}

	if s.idPasswordPair != nil {
		key, err := s.loadSigningKeyFromIdAndPassword(req)
		if err != nil {
			return nil, err
		}
		s.cachedSigningKey = key
		return key, nil
	}

	return nil, errors.New("invalid signing key loader")
}

func (s *SigningKeyLoader) loadSigningKeyFromMasterSeed() (requester.SigningKey, error) {
	if s.masterSeedAndNetwork == nil {
		return nil, errors.New("invalid signing key loader")
	}

	derivationPath := "m/5"
	key, error := crypto.DeriveXpriv(s.masterSeedAndNetwork.masterSeed, derivationPath)
	if error != nil {
		return nil, error
	}
	keyBytes := key.Key
	if error != nil {
		return nil, error
	}
	return &requester.Secp256k1SigningKey{PrivateKey: keyBytes}, nil
}

func (s *SigningKeyLoader) loadSigningKeyFromIdAndPassword(req requester.Requester) (requester.SigningKey, error) {
	if s.idPasswordPair == nil {
		return nil, errors.New("invalid signing key loader")
	}
	variables := map[string]interface{}{
		"node_id": s.idPasswordPair.nodeId,
	}
	response, err := req.ExecuteGraphql(scripts.RECOVER_NODE_SIGNING_KEY_QUERY, variables, nil)
	if err != nil {
		return nil, err
	}

	output := response["entity"].(map[string]interface{})
	encryptedKeyOutput := output["encrypted_signing_private_key"].(map[string]interface{})
	encryptedKey := encryptedKeyOutput["encrypted_value"].(string)
	cipher := encryptedKeyOutput["cipher"].(string)

	signingKey, err := crypto.DecryptPrivateKey(cipher, encryptedKey, s.idPasswordPair.password)
	if err != nil {
		return nil, err
	}

	return &requester.RsaSigningKey{PrivateKey: signingKey}, nil
}

type idPasswordPair struct {
	nodeId   string
	password string
}

type masterSeedAndNetwork struct {
	masterSeed []byte
	network    objects.BitcoinNetwork
}
