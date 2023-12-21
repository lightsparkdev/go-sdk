// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package requester

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"time"

	lightspark "github.com/lightsparkdev/go-sdk"
)

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
	isWhitelistedLocalHost := parsedUrl.Hostname() == "localhost" ||
		parsedUrl.Hostname() == "app.minikube.local" ||
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

	var nonce uint32
	if signingKey != nil {
		randomBigInt, err := rand.Int(rand.Reader, big.NewInt(0xFFFFFFFF))
		if err != nil {
			return nil, err
		}
		nonce = uint32(randomBigInt.Uint64())
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

	var serverUrl string
	if r.BaseUrl == nil {
		serverUrl = DEFAULT_BASE_URL
	} else {
		serverUrl = *r.BaseUrl
	}
	if err := ValidateBaseUrl(serverUrl); err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", serverUrl, bytes.NewBuffer(encodedPayload))
	request.SetBasicAuth(r.ApiTokenClientId, r.ApiTokenClientSecret)
	request.Header.Add("Content-Type", "application/json")
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
		return nil, errors.New("lightspark request failed: " + response.Status)
	}

	data, err := ioutil.ReadAll(response.Body)
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
			return nil, errors.New(errorMessage)
		}
		extensions := errMap["extensions"].(map[string]interface{})
		if extensions["error_name"] == nil {
			return nil, errors.New(errorMessage)
		}
		errorName := extensions["error_name"].(string)
		return nil, errors.New(errorName + " - " + errorMessage)
	}

	return result["data"].(map[string]interface{}), nil
}

func (r *Requester) getUserAgent() string {
	return "lightspark-go-sdk/" + lightspark.VERSION + " go/" + runtime.Version()
}
