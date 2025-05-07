package crypto

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
)

func GeneratePreimageNonce() ([]byte, error) {
	nonce := make([]byte, 32)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

func DeriveXpriv(seed []byte, path string) (*hdkeychain.ExtendedKey, error) {
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	currentKey := masterKey

	paths := strings.Split(path, "/")
	for i, path := range paths {
		if i == 0 {
			if path == "m" {
				continue
			} else {
				return nil, fmt.Errorf("invalid path: %s", path)
			}
		}

		var baseIndex uint32
		if strings.HasSuffix(path, "'") {
			baseIndex = 0x80000000
			path = path[:len(path)-1]
		}

		childIndex, err := strconv.ParseUint(path, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid path: %s", path)
		}

		if childIndex > 0x80000000 {
			return nil, fmt.Errorf("invalid path: %s", path)
		}

		currentKey, err = currentKey.Derive(uint32(childIndex) + baseIndex)
		if err != nil {
			return nil, err
		}
	}

	return currentKey, nil
}

func DerivePreimageBaseKey(seed []byte) ([]byte, error) {
	xpriv, err := DeriveXpriv(seed, "m/4'")
	if err != nil {
		return nil, err
	}

	privKey, err := xpriv.ECPrivKey()
	if err != nil {
		return nil, err
	}

	return privKey.Serialize(), nil
}

func GeneratePreimageAndPaymentHash(key []byte, nonce []byte) ([]byte, []byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write([]byte("invoice preimage"))
	if err != nil {
		return nil, nil, err
	}
	_, err = h.Write(nonce)
	if err != nil {
		return nil, nil, err
	}
	preimage := h.Sum(nil)[:32]
	paymentHash := sha256.Sum256(preimage)

	return preimage, paymentHash[:], nil
}
