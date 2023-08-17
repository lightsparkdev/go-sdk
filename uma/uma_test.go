package uma

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	eciesgo "github.com/ecies/go/v2"
	"github.com/lightsparkdev/go-sdk/objects"
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
	expectedQuery := LnurlpRequest{
		ReceiverAddress: "bob@vasp2.com",
		Signature:       "signature",
		TrStatus:        true,
		Nonce:           "12345",
		Timestamp:       expectedTime,
		VaspDomain:      "vasp1.com",
	}
	urlString := "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&trStatus=true&timestamp=" + strconv.FormatInt(timeSec, 10)
	urlObj, _ := url.Parse(urlString)
	query, err := ParseLnurlpRequest(*urlObj)
	if err != nil || query == nil {
		t.Fatalf("Parse(%s) failed: %s", urlObj, err)
	}
	assert.ObjectsAreEqual(expectedQuery, *query)
}

func TestIsUmaQueryValid(t *testing.T) {
	urlString := "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&trStatus=true&timestamp=12345678"
	urlObj, _ := url.Parse(urlString)
	assert.True(t, IsUmaLnurlpQuery(*urlObj))
}

func TestIsUmaQueryMissingParams(t *testing.T) {
	urlString := "https://vasp2.com/.well-known/lnurlp/bob?nonce=12345&vaspDomain=vasp1.com&trStatus=true&timestamp=12345678"
	urlObj, _ := url.Parse(urlString)
	assert.False(t, IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&vaspDomain=vasp1.com&trStatus=true&timestamp=12345678"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&nonce=12345&trStatus=true&timestamp=12345678"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&timestamp=12345678"
	urlObj, _ = url.Parse(urlString)
	// TrStatus is optional
	assert.True(t, IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&trStatus=true"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/.well-known/lnurlp/bob"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaLnurlpQuery(*urlObj))
}

func TestIsUmaQueryInvalidPath(t *testing.T) {
	urlString := "https://vasp2.com/.well-known/lnurla/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&trStatus=true&timestamp=12345678"
	urlObj, _ := url.Parse(urlString)
	assert.False(t, IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/bob?signature=signature&nonce=12345&vaspDomain=vasp1.com&trStatus=true&timestamp=12345678"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaLnurlpQuery(*urlObj))

	urlString = "https://vasp2.com/?signature=signature&nonce=12345&vaspDomain=vasp1.com&trStatus=true&timestamp=12345678"
	urlObj, _ = url.Parse(urlString)
	assert.False(t, IsUmaLnurlpQuery(*urlObj))
}

func TestSignAndVerifyLnurlpRequest(t *testing.T) {
	privateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)
	queryUrl, err := GetSignedLnurlpRequestUrl(privateKey.Serialize(), "$bob@vasp2.com", "vasp1.com", true)
	require.NoError(t, err)
	query, err := ParseLnurlpRequest(*queryUrl)
	require.NoError(t, err)
	err = VerifyUmaLnurlpQuerySignature(query, privateKey.PubKey().SerializeUncompressed())
	require.NoError(t, err)
}

func TestSignAndVerifyLnurlpRequestInvalidSignature(t *testing.T) {
	privateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)
	queryUrl, err := GetSignedLnurlpRequestUrl(privateKey.Serialize(), "$bob@vasp2.com", "vasp1.com", true)
	require.NoError(t, err)
	query, err := ParseLnurlpRequest(*queryUrl)
	require.NoError(t, err)
	err = VerifyUmaLnurlpQuerySignature(query, []byte("invalid pub key"))
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
	response, err := GetLnurlpResponse(
		request,
		receiverSigningPrivateKey.Serialize(),
		true,
		"https://vasp2.com/api/lnurl/payreq/$bob",
		metadata,
		1,
		10_000_000,
		PayerDataOptions{
			NameRequired:       false,
			EmailRequired:      false,
			ComplianceRequired: true,
		},
		[]Currency{
			{
				Code:                "USD",
				Name:                "US Dollar",
				Symbol:              "$",
				MillisatoshiPerUnit: 34_150,
				MinSendable:         1,
				MaxSendable:         10_000_000,
			},
		},
		true,
	)
	require.NoError(t, err)
	responseJson, err := json.Marshal(response)
	require.NoError(t, err)

	response, err = ParseLnurlpResponse(responseJson)
	require.NoError(t, err)
	err = VerifyUmaLnurlpResponseSignature(response, receiverSigningPrivateKey.PubKey().SerializeUncompressed())
	require.NoError(t, err)
}

func TestPayReqCreationAndParsing(t *testing.T) {
	senderSigningPrivateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)
	receiverEncryptionPrivateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)

	trInfo := "some TR info for VASP2"
	payreq, err := GetPayRequest(
		receiverEncryptionPrivateKey.PubKey().SerializeUncompressed(),
		senderSigningPrivateKey.Serialize(),
		"USD",
		1000,
		"$alice@vasp1.com",
		nil,
		nil,
		&trInfo,
		true,
		nil,
		"/api/lnurl/utxocallback?txid=1234",
	)
	require.NoError(t, err)

	payreqJson, err := json.Marshal(payreq)
	require.NoError(t, err)

	payreq, err = ParsePayRequest(payreqJson)
	require.NoError(t, err)

	err = VerifyPayReqSignature(payreq, senderSigningPrivateKey.PubKey().SerializeUncompressed())
	require.NoError(t, err)

	encryptedTrInfo := payreq.PayerData.Compliance.TrInfo
	require.NotNil(t, encryptedTrInfo)

	encryptedTrInfoBytes, err := hex.DecodeString(*encryptedTrInfo)
	require.NoError(t, err)
	eciesPrivKey := eciesgo.NewPrivateKeyFromBytes(receiverEncryptionPrivateKey.Serialize())
	decryptedTrInfo, err := eciesgo.Decrypt(eciesPrivKey, encryptedTrInfoBytes)
	require.NoError(t, err)
	assert.Equal(t, trInfo, string(decryptedTrInfo))
}

type FakeInvoiceCreator struct{}

func (f *FakeInvoiceCreator) CreateLnurlInvoice(string, *[]byte, int64, string, *int32) (*objects.Invoice, error) {
	return &objects.Invoice{
		Id:        "1234",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    objects.PaymentRequestStatusOpen,
		Data: objects.InvoiceData{
			EncodedPaymentRequest: "lntb100n1p0z9j",
		},
	}, nil
}

func TestPayReqResponseAndParsing(t *testing.T) {
	senderSigningPrivateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)
	receiverEncryptionPrivateKey, err := secp256k1.GeneratePrivateKey()
	require.NoError(t, err)

	trInfo := "some TR info for VASP2"
	payreq, err := GetPayRequest(
		receiverEncryptionPrivateKey.PubKey().SerializeUncompressed(),
		senderSigningPrivateKey.Serialize(),
		"USD",
		1000,
		"$alice@vasp1.com",
		nil,
		nil,
		&trInfo,
		true,
		nil,
		"/api/lnurl/utxocallback?txid=1234",
	)
	require.NoError(t, err)
	client := &FakeInvoiceCreator{}
	nodeId := "nodeId"
	metadata, err := createMetadataForBob()
	require.NoError(t, err)
	payreqResponse, err := GetPayReqResponse(
		payreq,
		client,
		nodeId,
		nil,
		metadata,
		"USD",
		34_150,
		int32(time.Hour.Seconds()),
		[]string{"abcdef12345"},
		"/api/lnurl/utxocallback?txid=1234",
	)
	require.NoError(t, err)

	payreqResponseJson, err := json.Marshal(payreqResponse)
	require.NoError(t, err)

	payreqResponse, err = ParsePayReqResponse(payreqResponseJson)
	require.NoError(t, err)
}

func createLnurlpRequest(t *testing.T, signingPrivateKey []byte) *LnurlpRequest {
	queryUrl, err := GetSignedLnurlpRequestUrl(signingPrivateKey, "$bob@vasp2.com", "vasp1.com", true)
	require.NoError(t, err)
	query, err := ParseLnurlpRequest(*queryUrl)
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
