package uma_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	eciesgo "github.com/ecies/go/v2"
	"github.com/lightsparkdev/go-sdk/uma"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	expectedTime, _ := time.Parse(time.RFC3339, "2023-07-27T22:46:08Z")
	timeSec := expectedTime.Unix()
	expectedQuery := uma.LnurlpRequest{
		ReceiverAddress:       "bob@vasp2.com",
		Signature:             "signature",
		IsSubjectToTravelRule: true,
		Nonce:                 "12345",
		Timestamp:             expectedTime,
		VaspDomain:            "vasp1.com",
	}
	urlString := "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&isSubjectToTravelRule=true&timestamp=" + strconv.FormatInt(timeSec, 10)
	urlObj, _ := url.Parse(urlString)
	query, err := uma.ParseLnurlpRequest(*urlObj)
	if err != nil || query == nil {
		t.Fatalf("Parse(%s) failed: %s", urlObj, err)
	}
	assert.ObjectsAreEqual(expectedQuery, *query)
}

func TestIsUmaQueryValid(t *testing.T) {
	urlString := "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&isSubjectToTravelRule=true&timestamp=12345678"
	urlObj, _ := url.Parse(urlString)
	assert.True(t, uma.IsUmaLnurlpQuery(*urlObj))
}

func TestIsUmaQueryMissingParams(t *testing.T) {
	urlString := "https://vasp2.com/.well-known/lnurlp/bob?nonce=12345&vaspDomain=vasp1.com&isSubjectToTravelRule=true&timestamp=12345678"
	urlObj, _ := url.Parse(urlString)
	assert.False(t, uma.IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&vaspDomain=vasp1.com&isSubjectToTravelRule=true&timestamp=12345678"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, uma.IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&nonce=12345&isSubjectToTravelRule=true&timestamp=12345678"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, uma.IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&timestamp=12345678"
	urlObj, _ = url.Parse(urlString)
	// IsSubjectToTravelRule is optional
	assert.True(t, uma.IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&isSubjectToTravelRule=true"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, uma.IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, uma.IsUmaLnurlpQuery(*urlObj))
}

func TestIsUmaQueryInvalidPath(t *testing.T) {
	urlString := "https://vasp2.com/.well-known/lnurla/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&isSubjectToTravelRule=true&timestamp=12345678"
	urlObj, _ := url.Parse(urlString)
	assert.False(t, uma.IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&isSubjectToTravelRule=true&timestamp=12345678"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, uma.IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/?signature=signature&nonce=12345&vaspDomain=vasp1.com&isSubjectToTravelRule=true&timestamp=12345678"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, uma.IsUmaLnurlpQuery(*urlObj))
}

func TestSignAndVerifyLnurlpRequest(t *testing.T) {
	privateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)
	queryUrl, err := uma.GetSignedLnurlpRequestUrl(privateKey.Serialize(), "$bob@vasp2.com", "vasp1.com", true)
	require.NoError(t, err)
	query, err := uma.ParseLnurlpRequest(*queryUrl)
	require.NoError(t, err)
	err = uma.VerifyUmaLnurlpQuerySignature(query, privateKey.PubKey().SerializeUncompressed())
	require.NoError(t, err)
}

func TestSignAndVerifyLnurlpRequestInvalidSignature(t *testing.T) {
	privateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)
	queryUrl, err := uma.GetSignedLnurlpRequestUrl(privateKey.Serialize(), "$bob@vasp2.com", "vasp1.com", true)
	require.NoError(t, err)
	query, err := uma.ParseLnurlpRequest(*queryUrl)
	require.NoError(t, err)
	err = uma.VerifyUmaLnurlpQuerySignature(query, []byte("invalid pub key"))
	require.Error(t, err)
}

func TestSignAndVerifyLnurlpResponse(t *testing.T) {
	senderSigningPrivateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)
	receiverSigningPrivateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)
	request := createLnurlpRequest(t, senderSigningPrivateKey.Serialize())
	metadata, err := createMetadataForBob()
	require.NoError(t, err)
	response, err := uma.GetLnurlpResponse(
		request,
		receiverSigningPrivateKey.Serialize(),
		true,
		"https://vasp2.com/api/lnurl/payreq/$bob",
		metadata,
		1,
		10_000_000,
		uma.PayerDataOptions{
			NameRequired:       false,
			EmailRequired:      false,
			ComplianceRequired: true,
		},
		[]uma.Currency{
			{
				Code:                "USD",
				Name:                "US Dollar",
				Symbol:              "$",
				MillisatoshiPerUnit: 34_150,
				MinSendable:         1,
				MaxSendable:         10_000_000,
			},
		},
		uma.KycStatusVerified,
	)
	require.NoError(t, err)
	responseJson, err := json.Marshal(response)
	require.NoError(t, err)

	response, err = uma.ParseLnurlpResponse(responseJson)
	require.NoError(t, err)
	err = uma.VerifyUmaLnurlpResponseSignature(response, receiverSigningPrivateKey.PubKey().SerializeUncompressed())
	require.NoError(t, err)
}

func TestPayReqCreationAndParsing(t *testing.T) {
	senderSigningPrivateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)
	receiverEncryptionPrivateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)

	trInfo := "some TR info for VASP2"
	payreq, err := uma.GetPayRequest(
		receiverEncryptionPrivateKey.PubKey().SerializeUncompressed(),
		senderSigningPrivateKey.Serialize(),
		"USD",
		1000,
		"$alice@vasp1.com",
		nil,
		nil,
		&trInfo,
		uma.KycStatusVerified,
		nil,
		nil,
		"/api/lnurl/utxocallback?txid=1234",
	)
	require.NoError(t, err)

	payreqJson, err := json.Marshal(payreq)
	require.NoError(t, err)

	payreq, err = uma.ParsePayRequest(payreqJson)
	require.NoError(t, err)

	err = uma.VerifyPayReqSignature(payreq, senderSigningPrivateKey.PubKey().SerializeUncompressed())
	require.NoError(t, err)

	encryptedTrInfo := payreq.PayerData.Compliance.EncryptedTravelRuleInfo
	require.NotNil(t, encryptedTrInfo)

	encryptedTrInfoBytes, err := hex.DecodeString(*encryptedTrInfo)
	require.NoError(t, err)
	eciesPrivKey := eciesgo.NewPrivateKeyFromBytes(receiverEncryptionPrivateKey.Serialize())
	decryptedTrInfo, err := eciesgo.Decrypt(eciesPrivKey, encryptedTrInfoBytes)
	require.NoError(t, err)
	assert.Equal(t, trInfo, string(decryptedTrInfo))
}

type FakeInvoiceCreator struct{}

func (f *FakeInvoiceCreator) CreateLnurlInvoice(int64, string) (*string, error) {
	encodedInvoice := "lnbcrt100n1p0z9j"
	return &encodedInvoice, nil
}

func TestPayReqResponseAndParsing(t *testing.T) {
	senderSigningPrivateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)
	receiverEncryptionPrivateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)

	trInfo := "some TR info for VASP2"
	payreq, err := uma.GetPayRequest(
		receiverEncryptionPrivateKey.PubKey().SerializeUncompressed(),
		senderSigningPrivateKey.Serialize(),
		"USD",
		1000,
		"$alice@vasp1.com",
		nil,
		nil,
		&trInfo,
		uma.KycStatusVerified,
		nil,
		nil,
		"/api/lnurl/utxocallback?txid=1234",
	)
	require.NoError(t, err)
	client := &FakeInvoiceCreator{}
	metadata, err := createMetadataForBob()
	require.NoError(t, err)
	payreqResponse, err := uma.GetPayReqResponse(
		payreq,
		client,
		metadata,
		"USD",
		34_150,
		100_000,
		[]string{"abcdef12345"},
		nil,
		"/api/lnurl/utxocallback?txid=1234",
	)
	require.NoError(t, err)

	payreqResponseJson, err := json.Marshal(payreqResponse)
	require.NoError(t, err)

	payreqResponse, err = uma.ParsePayReqResponse(payreqResponseJson)
	require.NoError(t, err)
}

func createLnurlpRequest(t *testing.T, signingPrivateKey []byte) *uma.LnurlpRequest {
	queryUrl, err := uma.GetSignedLnurlpRequestUrl(signingPrivateKey, "$bob@vasp2.com", "vasp1.com", true)
	require.NoError(t, err)
	query, err := uma.ParseLnurlpRequest(*queryUrl)
	require.NoError(t, err)
	return query
}

func createMetadataForBob() (string, error) {
	metadata := [][]string{
		{"text/plain", fmt.Sprintf("Pay to vasp2.com user $bob")},
		{"text/identifier", fmt.Sprintf("$bob@vasp2.com")},
	}

	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		return "", err
	}

	return string(jsonMetadata), nil
}
