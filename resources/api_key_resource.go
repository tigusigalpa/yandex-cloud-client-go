package resources

import (
	"net/http"

	"github.com/tigusigalpa/yandex-cloud-client-go/auth"
	"github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

// APIKeyResource handles API key-related operations
type APIKeyResource struct {
	*AbstractResource
}

// NewAPIKeyResource creates a new API key resource
func NewAPIKeyResource(httpClient *http.Client, authManager *auth.IAMTokenManager, baseURI string) *APIKeyResource {
	return &APIKeyResource{
		AbstractResource: NewAbstractResource(httpClient, authManager, baseURI),
	}
}

// List gets list of API keys for service account
func (r *APIKeyResource) List(serviceAccountID string, pageSize *int, pageToken *string) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	params["serviceAccountId"] = serviceAccountID

	if pageSize != nil {
		params["pageSize"] = *pageSize
	}
	if pageToken != nil {
		params["pageToken"] = *pageToken
	}

	query := r.BuildQueryString(params)
	return r.MakeRequest("GET", "iam/v1/apiKeys"+query, nil)
}

// Get gets API key details
func (r *APIKeyResource) Get(apiKeyID string) (map[string]interface{}, error) {
	if apiKeyID == "" {
		return nil, errors.NewValidationError("API key ID cannot be empty")
	}

	return r.MakeRequest("GET", "iam/v1/apiKeys/"+apiKeyID, nil)
}

// Create creates a new API key
func (r *APIKeyResource) Create(serviceAccountID string, description *string) (map[string]interface{}, error) {
	if serviceAccountID == "" {
		return nil, errors.NewValidationError("Service account ID cannot be empty")
	}

	data := map[string]interface{}{
		"serviceAccountId": serviceAccountID,
	}

	if description != nil {
		data["description"] = *description
	}

	return r.MakeRequest("POST", "iam/v1/apiKeys", data)
}

// Update updates API key
func (r *APIKeyResource) Update(apiKeyID string, data map[string]interface{}) (map[string]interface{}, error) {
	if apiKeyID == "" {
		return nil, errors.NewValidationError("API key ID cannot be empty")
	}

	if len(data) == 0 {
		return nil, errors.NewValidationError("Update data cannot be empty")
	}

	return r.MakeRequest("PATCH", "iam/v1/apiKeys/"+apiKeyID, data)
}

// Delete deletes API key
func (r *APIKeyResource) Delete(apiKeyID string) (map[string]interface{}, error) {
	if apiKeyID == "" {
		return nil, errors.NewValidationError("API key ID cannot be empty")
	}

	return r.MakeRequest("DELETE", "iam/v1/apiKeys/"+apiKeyID, nil)
}
