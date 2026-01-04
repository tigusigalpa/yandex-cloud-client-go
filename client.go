package yandexcloud

import (
	"net/http"

	"github.com/tigusigalpa/yandex-cloud-client-go/auth"
	"github.com/tigusigalpa/yandex-cloud-client-go/errors"
	"github.com/tigusigalpa/yandex-cloud-client-go/resources"
)

const (
	iamBaseURI             = "https://iam.api.cloud.yandex.net/"
	organizationBaseURI    = "https://organization-manager.api.cloud.yandex.net/"
	resourceManagerBaseURI = "https://resource-manager.api.cloud.yandex.net/"
)

// Client is the main client for Yandex Cloud API
type Client struct {
	httpClient  *http.Client
	authManager *auth.IAMTokenManager
}

// NewClient creates a new Yandex Cloud client
func NewClient(oauthToken string, httpClient *http.Client) (*Client, error) {
	if oauthToken == "" {
		return nil, errors.NewAuthenticationError("OAuth token cannot be empty", nil)
	}

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	authManager, err := auth.NewIAMTokenManager(oauthToken, httpClient)
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient:  httpClient,
		authManager: authManager,
	}, nil
}

// Organizations returns the organization resource
func (c *Client) Organizations() *resources.OrganizationResource {
	return resources.NewOrganizationResource(c.httpClient, c.authManager, organizationBaseURI)
}

// Clouds returns the cloud resource
func (c *Client) Clouds() *resources.CloudResource {
	return resources.NewCloudResource(c.httpClient, c.authManager, resourceManagerBaseURI)
}

// Folders returns the folder resource
func (c *Client) Folders() *resources.FolderResource {
	return resources.NewFolderResource(c.httpClient, c.authManager, resourceManagerBaseURI)
}

// RefreshTokens returns the refresh token resource
func (c *Client) RefreshTokens() *resources.RefreshTokenResource {
	return resources.NewRefreshTokenResource(c.httpClient, c.authManager, iamBaseURI)
}

// ServiceAccounts returns the service account resource
func (c *Client) ServiceAccounts() *resources.ServiceAccountResource {
	return resources.NewServiceAccountResource(c.httpClient, c.authManager, iamBaseURI)
}

// UserAccounts returns the user account resource
func (c *Client) UserAccounts() *resources.UserAccountResource {
	return resources.NewUserAccountResource(c.httpClient, c.authManager, iamBaseURI)
}

// YandexPassportUserAccounts returns the Yandex Passport user account resource
func (c *Client) YandexPassportUserAccounts() *resources.YandexPassportUserAccountResource {
	return resources.NewYandexPassportUserAccountResource(c.httpClient, c.authManager, iamBaseURI)
}

// APIKeys returns the API key resource
func (c *Client) APIKeys() *resources.APIKeyResource {
	return resources.NewAPIKeyResource(c.httpClient, c.authManager, iamBaseURI)
}

// GetHTTPClient returns the HTTP client
func (c *Client) GetHTTPClient() *http.Client {
	return c.httpClient
}

// GetAuthManager returns the authentication manager
func (c *Client) GetAuthManager() *auth.IAMTokenManager {
	return c.authManager
}

// GetOAuthToken returns the OAuth token
func (c *Client) GetOAuthToken() string {
	return c.authManager.GetOAuthToken()
}
