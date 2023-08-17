// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package uma

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Query struct {
	CurrencyCode    string
	Signature       string
	Utxos           []string
	SenderAddress   string
	ReceiverAddress string
	nonce           string
	expiry          time.Time
}

func (q *Query) signablePayload() []byte {
	payloadString := strings.Join([]string{q.ReceiverAddress, q.SenderAddress, q.nonce, q.expiry.String()}, "|")
	return []byte(payloadString)
}

// FetchPublicKeyForVasp fetches the public key for another VASP.
//
// If the public key is not in the cache, it will be fetched from the VASP's domain.
// The public key will be cached for future use.
//
// Args:
//
//	vaspDomain: the domain of the VASP.
//	cache: the PublicKeyCache cache to use. You can use the InMemoryPublicKeyCache struct, or implement your own persistent cache with any storage type.
func FetchPublicKeyForVasp(vaspDomain string, cache PublicKeyCache) (*rsa.PublicKey, error) {
	publicKey := cache.FetchPublicKeyForVasp(vaspDomain)
	if publicKey != nil {
		return publicKey, nil
	}

	// TODO: Scheme?
	resp, err := http.Get(vaspDomain + "/.well-known/lnurlpubkey")
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return nil, errors.New("invalid response from VASP")
	}

	publicKeyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	publicKey, err = x509.ParsePKCS1PublicKey(publicKeyBytes)
	if err != nil {
		return nil, err
	}

	cache.AddPublicKeyForVasp(vaspDomain, publicKey)
	return publicKey, nil
}

// VerifyAndParseUmaQuery Verifies the signature and parses the message into a UmaQuery object.
//
// Args:
//
//	url: the full URL of the uma request.
//	otherVaspPubKey: the base64-encoded public key of the VASP making this request.
func VerifyAndParseUmaQuery(url url.URL, otherVaspPubKey rsa.PublicKey) (*Query, error) {
	query, err := Parse(url)
	if err != nil {
		return nil, err
	}
	hashedPayload := sha256.Sum256(query.signablePayload())
	err = verifySignature(hashedPayload, query.Signature, otherVaspPubKey)
	if err != nil {
		return nil, err
	}
	return query, nil
}

// verifySignature Verifies the signature of the uma request.
//
// Args:
//
//	hashedPayload: the sha256 hash of the payload.
//	signature: the base64-encoded signature.
//	otherVaspPubKey: the base64-encoded public key of the VASP making this request.
func verifySignature(hashedPayload [32]byte, signature string, otherVaspPubKey rsa.PublicKey) error {
	decodedSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	err = rsa.VerifyPSS(&otherVaspPubKey, crypto.SHA256, hashedPayload[:], decodedSignature, nil)
	if err != nil {
		return errors.New("invalid uma signature")
	}

	return nil
}

// IsUmaQuery Checks if the given URL is a valid UMA request.
func IsUmaQuery(url url.URL) bool {
	query, err := Parse(url)
	return err == nil && query != nil
}

// Parse Parses the message into a Query object.
// Args:
//
//	url: the full URL of the uma request.
func Parse(url url.URL) (*Query, error) {
	query := url.Query()
	currencyCode := query.Get("currency")
	signature := query.Get("signature")
	utxos := query.Get("utxos")
	senderAddress := query.Get("sender")
	nonce := query.Get("nonce")
	expiry := query.Get("expiry")
	expiryDate, dateErr := time.Parse(time.RFC3339, expiry)

	if dateErr != nil {
		return nil, errors.New("invalid expiry")
	}

	if currencyCode == "" || signature == "" || utxos == "" || senderAddress == "" || expiry == "" || nonce == "" {
		return nil, errors.New("missing uma query parameters. Currency, signature, utxos, expiry, nonce, and sender are required")
	}

	pathParts := strings.Split(url.Path, "/")
	if len(pathParts) != 4 || pathParts[1] != ".well-known" || pathParts[2] != "lnurlp" {
		return nil, errors.New("invalid uma request path")
	}
	receiverAddress := pathParts[3] + "@" + url.Host

	return &Query{
		CurrencyCode:    currencyCode,
		Signature:       signature,
		Utxos:           strings.Split(utxos, ","),
		SenderAddress:   senderAddress,
		ReceiverAddress: receiverAddress,
		nonce:           nonce,
		expiry:          expiryDate,
	}, nil
}
