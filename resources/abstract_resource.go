package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tigusigalpa/yandex-cloud-client-go/auth"
	"github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

// AbstractResource provides common functionality for all resources
type AbstractResource struct {
	httpClient  *http.Client
	authManager *auth.IAMTokenManager
	baseURI     string
}

// NewAbstractResource creates a new abstract resource
func NewAbstractResource(httpClient *http.Client, authManager *auth.IAMTokenManager, baseURI string) *AbstractResource {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &AbstractResource{
		httpClient:  httpClient,
		authManager: authManager,
		baseURI:     baseURI,
	}
}

// MakeRequest makes an HTTP request to Yandex Cloud API
func (r *AbstractResource) MakeRequest(method, uri string, body interface{}) (map[string]interface{}, error) {
	// Get valid IAM token
	iamToken, err := r.authManager.GetValidIAMToken()
	if err != nil {
		return nil, err
	}

	// Prepare request body
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, errors.NewAPIError("Failed to marshal request body", 0, err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	// Create request
	fullURL := r.baseURI + uri
	req, err := http.NewRequest(method, fullURL, reqBody)
	if err != nil {
		return nil, errors.NewAPIError("Failed to create request", 0, err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+iamToken)

	// Execute request
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewAPIError("HTTP request failed", 0, err)
	}
	defer resp.Body.Close()

	return r.parseResponse(resp)
}

// parseResponse parses HTTP response
func (r *AbstractResource) parseResponse(resp *http.Response) (map[string]interface{}, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewAPIError("Failed to read response body", resp.StatusCode, err)
	}

	if resp.StatusCode >= 400 {
		return nil, errors.NewAPIError(
			fmt.Sprintf("API request failed with status %d: %s", resp.StatusCode, string(body)),
			resp.StatusCode,
			nil,
		)
	}

	// Handle empty responses
	if len(body) == 0 {
		return make(map[string]interface{}), nil
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, errors.NewAPIError("Failed to parse JSON response", resp.StatusCode, err)
	}

	return data, nil
}

// BuildQueryString builds query string from parameters
func (r *AbstractResource) BuildQueryString(params map[string]interface{}) string {
	values := url.Values{}
	for key, value := range params {
		if value != nil {
			values.Add(key, fmt.Sprintf("%v", value))
		}
	}
	if len(values) == 0 {
		return ""
	}
	return "?" + values.Encode()
}
