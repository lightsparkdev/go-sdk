package uma

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	expectedTime, _ := time.Parse(time.RFC3339, "2023-07-27T22:46:08Z")
	expectedQuery := Query{
		CurrencyCode:    "USD",
		Signature:       "signature",
		Utxos:           []string{"abcde", "feed", "beef", "dead"},
		SenderAddress:   "alice@vasp1.com",
		ReceiverAddress: "bob@vasp2.com",
		nonce:           "12345",
		expiry:          expectedTime,
	}
	urlString := "https://vasp2.com/.well-known/lnurlp/bob?currency=USD&signature=signature&utxos=abcde,feed,beef,dead&sender=alice@vasp1.com&nonce=12345&expiry=2023-07-27T22:46:08Z"
	urlObj, _ := url.Parse(urlString)
	query, err := Parse(*urlObj)
	if err != nil || query == nil {
		t.Errorf("Parse(%s) failed: %s", urlObj, err)
	}
	assert.ObjectsAreEqual(expectedQuery, *query)
}

func TestIsUmaQueryValid(t *testing.T) {
	urlString := "https://vasp2.com/.well-known/lnurlp/bob?currency=USD&signature=signature&utxos=abcde,feed,beef,dead&sender=alice@vasp1.com&nonce=12345&expiry=2023-07-27T22:46:08Z"
	urlObj, _ := url.Parse(urlString)
	assert.True(t, IsUmaQuery(*urlObj))
}

func TestIsUmaQueryMissingParams(t *testing.T) {
	urlString := "https://vasp2.com/.well-known/lnurlp/bob?currency=USD&signature=signature&utxos=abcde,feed,beef,dead&sender=alice@vasp1.com&nonce=12345"
	urlObj, _ := url.Parse(urlString)
	assert.False(t, IsUmaQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?currency=USD&signature=signature&utxos=abcde,feed,beef,dead&sender=alice@vasp1.com&expiry=2023-07-27T22:46:08Z"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?currency=USD&signature=signature&utxos=abcde,feed,beef,dead&nonce=12345&expiry=2023-07-27T22:46:08Z"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?currency=USD&signature=signature&sender=alice@vasp1.com&nonce=12345&expiry=2023-07-27T22:46:08Z"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?currency=USD&utxos=abcde,feed,beef,dead&sender=alice@vasp1.com&nonce=12345&expiry=2023-07-27T22:46:08Z"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&utxos=abcde,feed,beef,dead&sender=alice@vasp1.com&nonce=12345&expiry=2023-07-27T22:46:08Z"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaQuery(*urlObj))
}

func TestIsUmaQueryInvalidPath(t *testing.T) {
	urlString := "https://vasp2.com/.well-known/lnurla/bob?currency=USD&signature=signature&utxos=abcde,feed,beef,dead&sender=alice@vasp1.com&nonce=12345&expiry=2023-07-27T22:46:08Z"
	urlObj, _ := url.Parse(urlString)
	assert.False(t, IsUmaQuery(*urlObj))

	urlString = "https://vasp2.com/bob?currency=USD&signature=signature&utxos=abcde,feed,beef,dead&sender=alice@vasp1.com&nonce=12345&expiry=2023-07-27T22:46:08Z"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaQuery(*urlObj))

	urlString = "https://vasp2.com/?currency=USD&signature=signature&utxos=abcde,feed,beef,dead&sender=alice@vasp1.com&nonce=12345&expiry=2023-07-27T22:46:08Z"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaQuery(*urlObj))
}
