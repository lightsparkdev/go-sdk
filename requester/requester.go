// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package requester

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"regexp"
	"runtime"
	"time"

	lightspark "github.com/lightsparkdev/go-sdk"
	"github.com/lightsparkdev/go-sdk/utils"
)

type Requester struct {
	ApiTokenClientId string

	ApiTokenClientSecret string

	BaseUrl *string
}

const DEFAULT_BASE_URL = "https://api.lightspark.com/graphql/server/2023-04-04"

func (r *Requester) ExecuteGraphql(query string, variables map[string]interface{},
	signingKey []byte) (map[string]interface{}, error) {

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

	request, err := http.NewRequest("POST", serverUrl, bytes.NewBuffer(encodedPayload))
	request.SetBasicAuth(r.ApiTokenClientId, r.ApiTokenClientSecret)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-GraphQL-Operation", operationName)
	request.Header.Add("User-Agent", r.getUserAgent())
	request.Header.Add("X-Lightspark-SDK", r.getUserAgent())

	if signingKey != nil {
		signature, err := utils.SignPayload(encodedPayload, signingKey)
		if err != nil {
			return nil, err
		}
		signaturePayloadBytes, err := json.Marshal(map[string]interface{}{
			"v":         1,
			"signature": signature,
		})
		request.Header.Add("X-Lightspark-Signing", bytes.NewBuffer(signaturePayloadBytes).String())
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	if errs, ok := result["errors"]; ok {
		err := errs.([]interface{})[0]
		errMap := err.(map[string]interface{})
		return nil, errors.New(errMap["message"].(string))
	}

	return result["data"].(map[string]interface{}), nil
}

func (r *Requester) getUserAgent() string {
	return "lightspark-go-sdk/" + lightspark.VERSION + " go/" + runtime.Version()
}
