// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package requester

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"github.com/DataDog/zstd"

	lightspark "github.com/lightsparkdev/go-sdk"
)

// RequestError indicates that a request to the Lightspark API failed.
// It could be due to a service outage or a network error.
// The request should be retried if RequestError is returned with server errors (500-599).
type RequestError struct {
	Message string
	StatusCode int
}

func (e RequestError) Error() string {
	return "lightspark request failed: " + strconv.Itoa(e.StatusCode) + ": " + e.Message
}

// GraphQLInternalError indicates there's a failure in the Lightspark API.
// It could be due to a bug on Ligthspark's side. 
// The request can be retried, because the error might be transient.
type GraphQLInternalError struct {
	Message string
}

func (e GraphQLInternalError) Error() string {
	return "lightspark request failed: " + e.Message
}

// GraphQLError indicates the GraphQL request succeeded, but there's a user error.
// The request should not be retried, because the error is due to the user's input.
type GraphQLError struct {
	Message string
	Type string
}

func (e GraphQLError) Error() string {
	return  e.Type + ": " + e.Message
}

type Requester struct {
	ApiTokenClientId string

	ApiTokenClientSecret string

	BaseUrl *string

	HTTPClient *http.Client
}

func NewRequester(apiTokenClientId string, apiTokenClientSecret string) *Requester {
	return &Requester{
		ApiTokenClientId:     apiTokenClientId,
		ApiTokenClientSecret: apiTokenClientSecret,
	}
}

func NewRequesterWithBaseUrl(apiTokenClientId string, apiTokenClientSecret string, baseUrl *string) *Requester {
	if baseUrl == nil {
		return NewRequester(apiTokenClientId, apiTokenClientSecret)
	}
	if err := ValidateBaseUrl(*baseUrl); err != nil {
		panic(err)
	}
	return &Requester{
		ApiTokenClientId:     apiTokenClientId,
		ApiTokenClientSecret: apiTokenClientSecret,
		BaseUrl:              baseUrl,
	}
}

func ValidateBaseUrl(baseUrl string) error {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		return errors.New("invalid base url. Not a valid URL")
	}
	hostNameParts := strings.Split(parsedUrl.Hostname(), ".")
	hostNameTld := hostNameParts[len(hostNameParts)-1]
	isWhitelistedLocalHost := parsedUrl.Hostname() == "localhost" ||
		hostNameTld == "local" ||
		hostNameTld == "internal" ||
		parsedUrl.Hostname() == "127.0.0.1"
	if parsedUrl.Scheme != "https" && !isWhitelistedLocalHost {
		return errors.New("invalid base url. Must be https:// if not targeting localhost")
	}
	return nil
}

const DEFAULT_BASE_URL = "https://api.lightspark.com/graphql/server/2023-09-13"

func (r *Requester) ExecuteGraphql(query string, variables map[string]interface{},
	signingKey SigningKey,
) (map[string]interface{}, error) {
	re := regexp.MustCompile(`(?i)\s*(?:query|mutation)\s+(?P<OperationName>\w+)`)
	matches := re.FindStringSubmatch(query)
	index := re.SubexpIndex("OperationName")
	if len(matches) <= index {
		return nil, errors.New("invalid query payload")
	}
	operationName := matches[index]

	var nonce uint64
	if signingKey != nil {
		randomBigInt, err := rand.Int(rand.Reader, big.NewInt(0x7FFFFFFFFFFFFFFF))
		if err != nil {
			return nil, err
		}
		nonce = randomBigInt.Uint64()
	}

	var expiresAt string
	if signingKey != nil {
		expiresAt = time.Now().UTC().Add(time.Hour).Format(time.RFC3339)
	}

	payload := map[string]interface{}{
		"operationName": operationName,
		"query":         query,
		"variables":     variables,
		"nonce":         nonce,
		"expires_at":    expiresAt,
	}

	encodedPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.New("error when encoding payload")
	}

	body := encodedPayload;
	compressed := false
	if len(encodedPayload) > 1024 {
		compressed = true
		body, err = zstd.Compress(nil, encodedPayload)
		if err != nil {
			return nil, err
		}
	}

	var serverUrl string
	if r.BaseUrl == nil {
		serverUrl = DEFAULT_BASE_URL
	} else {
		serverUrl = *r.BaseUrl
	}
	if err := ValidateBaseUrl(serverUrl); err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", serverUrl, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(r.ApiTokenClientId, r.ApiTokenClientSecret)
	request.Header.Add("Content-Type", "application/json")
	if compressed {
		request.Header.Add("Content-Encoding", "zstd")
	}
	request.Header.Add("Accept-Encoding", "zstd")
	request.Header.Add("X-GraphQL-Operation", operationName)
	request.Header.Add("User-Agent", r.getUserAgent())
	request.Header.Add("X-Lightspark-SDK", r.getUserAgent())

	if signingKey != nil {
		signature, err := signingKey.Sign(encodedPayload)
		if err != nil {
			return nil, err
		}
		signaturePayloadBytes, err := json.Marshal(map[string]interface{}{
			"v":         1,
			"signature": base64.StdEncoding.EncodeToString(signature),
		})
		if err != nil {
			return nil, err
		}
		request.Header.Add("X-Lightspark-Signing", bytes.NewBuffer(signaturePayloadBytes).String())
	}

	httpClient := r.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, RequestError { Message: response.Status, StatusCode: response.StatusCode }
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.Header.Get("Content-Encoding") == "zstd" {
		data, err = zstd.Decompress(nil, data)
		if err != nil {
			return nil, err
		}
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	if errs, ok := result["errors"]; ok {
		err := errs.([]interface{})[0]
		errMap := err.(map[string]interface{})
		errorMessage := errMap["message"].(string)
		if errMap["extensions"] == nil {
			return nil, GraphQLInternalError{Message: errorMessage}
		}
		extensions := errMap["extensions"].(map[string]interface{})
		if extensions["error_name"] == nil {
			return nil, GraphQLInternalError{Message: errorMessage}
		}
		errorName := extensions["error_name"].(string)
		return nil, GraphQLError{Message: errorMessage, Type: errorName}
	}

	return result["data"].(map[string]interface{}), nil
}

func (r *Requester) getUserAgent() string {
	return "lightspark-go-sdk/" + lightspark.VERSION + " go/" + runtime.Version()
}
