// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"

	"golang.org/x/crypto/pbkdf2"

	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
)

const KEY_LEN = 32

func DecryptPrivateKey(cipherVersion string, encryptedValue string,
	password string) ([]byte, error) {

	decoded, err := base64.StdEncoding.DecodeString(encryptedValue)
	if err != nil {
		return nil, err
	}

	var header map[string]interface{}
	if cipherVersion == "AES_256_CBC_PBKDF2_5000_SHA256" {
		header = map[string]interface{}{"v": 0, "i": 5000}
		decoded = decoded[8:]
	} else {
		err = json.Unmarshal([]byte(cipherVersion), &header)
		if err != nil {
			return nil, err
		}
		if lvs, ok := header["lsv"]; ok {
			if lvs.(int) == 2 {
				header["v"] = 3
			}
		}
	}

	version := int(header["v"].(float64))
	if version < 0 || version > 4 {
		return nil, errors.New("unknown version ")
	}

	iteration := int(header["i"].(float64))

	if version == 3 {
		salt := decoded[len(decoded)-8:]
		nonce := decoded[0:12]
		ciphertext := decoded[12 : len(decoded)-8]
		key := deriveKey([]byte(password), salt, iteration)
		return decryptGcm(ciphertext, key, nonce)
	}

	var saltLen, ivLen int
	if version < 4 {
		saltLen = 8
		ivLen = 16
	} else {
		saltLen = 16
		ivLen = 12
	}

	salt := decoded[:saltLen]
	ciphertext := decoded[saltLen:]

	key, iv := deriveKeyIv([]byte(password), salt, iteration, KEY_LEN+ivLen)

	if version < 2 {
		return decryptCbc(ciphertext, key, iv)
	} else {
		return decryptGcm(ciphertext, key, iv)
	}
}

func Sha256HexString(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func DerivePublicKey(seedHexString string, derivationPath string) (string, error) {
	seedBytes, err := hex.DecodeString(seedHexString)
	if err != nil {
		return "", err
	}

	xpriv, err := DeriveXpriv(seedBytes, derivationPath)
	if err != nil {
		return "", err
	}

	xpub := xpriv.PublicKey()
	return xpub.String(), nil
}

func ECDH(seedBytes []byte, network lightspark_crypto.BitcoinNetwork, otherPubKey string) (string, error) {
	otherPubKeyBytes, err := hex.DecodeString(otherPubKey)
	if err != nil {
		return "", err
	}

	secretBytes, err := lightspark_crypto.Ecdh(seedBytes, network, otherPubKeyBytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(secretBytes), nil
}

func deriveKey(password []byte, salt []byte, iterations int) []byte {
	return pbkdf2.Key(password, salt, iterations, 32, sha256.New)
}

func deriveKeyIv(password []byte, salt []byte, iterations int, length int) ([]byte, []byte) {
	derived := pbkdf2.Key(password, salt, iterations, length, sha256.New)
	return derived[:KEY_LEN], derived[KEY_LEN:]
}

func decryptGcm(ciphertext []byte, key []byte, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func decryptCbc(ciphertext []byte, key []byte, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, nonce)
	decryptedData := make([]byte, len(ciphertext)-aes.BlockSize)
	mode.CryptBlocks(decryptedData, ciphertext[aes.BlockSize:])

	// Remove PKCS7 padding from the decrypted data
	unpaddedData, err := pkcs7Unpad(decryptedData)
	if err != nil {
		return nil, err
	}

	return unpaddedData, nil
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("cannot unpad empty data")
	}
	paddingLength := int(data[len(data)-1])
	if paddingLength > len(data) {
		return nil, errors.New("invalid padding length")
	}
	padding := data[len(data)-paddingLength:]
	for _, b := range padding {
		if int(b) != paddingLength {
			return nil, errors.New("invalid padding")
		}
	}
	return data[:len(data)-paddingLength], nil
}
