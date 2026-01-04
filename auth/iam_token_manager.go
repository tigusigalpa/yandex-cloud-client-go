package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

const (
	iamTokenEndpoint   = "https://iam.api.cloud.yandex.net/iam/v1/tokens"
	tokenLifetime      = 12 * time.Hour
	tokenRefreshMargin = 5 * time.Minute
)

// IAMTokenManager manages IAM token lifecycle including caching and auto-refresh
type IAMTokenManager struct {
	oauthToken     string
	httpClient     *http.Client
	iamToken       string
	iamTokenExpiry time.Time
	mu             sync.RWMutex
}

// NewIAMTokenManager creates a new IAM token manager
func NewIAMTokenManager(oauthToken string, httpClient *http.Client) (*IAMTokenManager, error) {
	if oauthToken == "" {
		return nil, errors.NewAuthenticationError("OAuth token cannot be empty", nil)
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return &IAMTokenManager{
		oauthToken: oauthToken,
		httpClient: httpClient,
	}, nil
}

// GetValidIAMToken returns a valid IAM token (with auto-refresh)
func (m *IAMTokenManager) GetValidIAMToken() (string, error) {
	m.mu.RLock()
	needsRefresh := m.iamToken == "" || time.Now().After(m.iamTokenExpiry)
	m.mu.RUnlock()

	if needsRefresh {
		if err := m.refreshIAMToken(); err != nil {
			return "", err
		}
	}

	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.iamToken, nil
}

// GetIAMToken gets a new IAM token using the OAuth token
func (m *IAMTokenManager) GetIAMToken() (string, error) {
	requestBody := map[string]string{
		"yandexPassportOauthToken": m.oauthToken,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", errors.NewAuthenticationError("Failed to marshal request", err)
	}

	req, err := http.NewRequest("POST", iamTokenEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", errors.NewAuthenticationError("Failed to create request", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return "", errors.NewAuthenticationError("Failed to get IAM token", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.NewAuthenticationError("Failed to read response", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorData map[string]interface{}
		json.Unmarshal(body, &errorData)
		errorMessage := "Unknown error"
		if msg, ok := errorData["message"].(string); ok {
			errorMessage = msg
		}
		return "", errors.NewAuthenticationError(
			fmt.Sprintf("Failed to get IAM token (HTTP %d): %s", resp.StatusCode, errorMessage),
			nil,
		)
	}

	var responseData struct {
		IAMToken string `json:"iamToken"`
	}

	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", errors.NewAuthenticationError("Failed to parse response", err)
	}

	if responseData.IAMToken == "" {
		return "", errors.NewAuthenticationError("IAM token not found in response", nil)
	}

	return responseData.IAMToken, nil
}

// refreshIAMToken refreshes the cached IAM token
func (m *IAMTokenManager) refreshIAMToken() error {
	token, err := m.GetIAMToken()
	if err != nil {
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.iamToken = token
	m.iamTokenExpiry = time.Now().Add(tokenLifetime - tokenRefreshMargin)

	return nil
}

// ClearCache clears the cached IAM token (force refresh on next request)
func (m *IAMTokenManager) ClearCache() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.iamToken = ""
	m.iamTokenExpiry = time.Time{}
}

// HasValidCachedToken checks if cached IAM token is still valid
func (m *IAMTokenManager) HasValidCachedToken() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.iamToken != "" && time.Now().Before(m.iamTokenExpiry)
}

// GetOAuthToken returns the OAuth token (for debugging purposes)
func (m *IAMTokenManager) GetOAuthToken() string {
	return m.oauthToken
}
