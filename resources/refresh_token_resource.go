package resources

import (
	"net/http"

	"github.com/tigusigalpa/yandex-cloud-client-go/auth"
	"github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

// RefreshTokenResource handles refresh token-related operations
type RefreshTokenResource struct {
	*AbstractResource
}

// NewRefreshTokenResource creates a new refresh token resource
func NewRefreshTokenResource(httpClient *http.Client, authManager *auth.IAMTokenManager, baseURI string) *RefreshTokenResource {
	return &RefreshTokenResource{
		AbstractResource: NewAbstractResource(httpClient, authManager, baseURI),
	}
}

// List gets list of refresh tokens
func (r *RefreshTokenResource) List(pageSize *int, pageToken *string) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	if pageSize != nil {
		params["pageSize"] = *pageSize
	}
	if pageToken != nil {
		params["pageToken"] = *pageToken
	}

	query := r.BuildQueryString(params)
	return r.MakeRequest("GET", "iam/v1/refreshTokens"+query, nil)
}

// Revoke revokes a refresh token
func (r *RefreshTokenResource) Revoke(tokenID string) (map[string]interface{}, error) {
	if tokenID == "" {
		return nil, errors.NewValidationError("Token ID cannot be empty")
	}

	return r.MakeRequest("DELETE", "iam/v1/refreshTokens/"+tokenID, nil)
}
