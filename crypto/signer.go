package crypto

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/tyler-smith/go-bip32"
)

func GeneratePreimageNonce() ([]byte, error) {
	nonce := make([]byte, 32)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

func DeriveXpriv(seed []byte, path string) (*bip32.Key, error) {
	masterKey, err := bip32.NewMasterKey(seed)
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

		currentKey, err = currentKey.NewChildKey(uint32(childIndex) + baseIndex)
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

	return xpriv.Key, nil
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

func GetPerCommitmentPoint(seedBytes []byte, derivationPath string, perCommitmentPointIdx uint64) ([]byte, error) {
	secret, err := ReleasePerCommitmentSecret(seedBytes, derivationPath, perCommitmentPointIdx)
	if err != nil {
		return nil, err
	}

	privKey, _ := btcec.PrivKeyFromBytes(secret)
	pubKey := privKey.PubKey()
	return pubKey.SerializeCompressed(), nil
}

func ReleasePerCommitmentSecret(seedBytes []byte, derivationPath string, perCommitmentPointIdx uint64) ([]byte, error) {
	xpriv, err := DeriveXpriv(seedBytes, derivationPath)
	if err != nil {
		return nil, err
	}

	privKeyHash := sha256.Sum256(xpriv.Key)
	channelSeed := privKeyHash[:]
	commitmentSeed := buildCommitmentSeed(channelSeed)
	commitmentSecret := buildCommitmentSecret(commitmentSeed[:], perCommitmentPointIdx)
	return commitmentSecret, nil
}

func buildCommitmentSeed(channelSeed []byte) [32]byte {
	combined := append(channelSeed, []byte("commitment seed")...)
	return sha256.Sum256(combined)
}

func buildCommitmentSecret(seed []byte, idx uint64) []byte {
	res := make([]byte, len(seed))
	copy(res, seed)

	for i := range 48 {
		bitpos := 47 - i
		if (idx & (1 << bitpos)) == (1 << bitpos) {
			res[bitpos/8] ^= 1 << (bitpos & 7)
			hash := sha256.Sum256(res)
			res = hash[:]
		}
	}
	return res
}
