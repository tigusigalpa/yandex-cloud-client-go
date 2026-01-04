package resources

import (
	"net/http"

	"github.com/tigusigalpa/yandex-cloud-client-go/auth"
	"github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

// UserAccountResource handles user account-related operations
type UserAccountResource struct {
	*AbstractResource
}

// NewUserAccountResource creates a new user account resource
func NewUserAccountResource(httpClient *http.Client, authManager *auth.IAMTokenManager, baseURI string) *UserAccountResource {
	return &UserAccountResource{
		AbstractResource: NewAbstractResource(httpClient, authManager, baseURI),
	}
}

// Get gets user account details by ID
func (r *UserAccountResource) Get(userAccountID string) (map[string]interface{}, error) {
	if userAccountID == "" {
		return nil, errors.NewValidationError("User account ID cannot be empty")
	}

	return r.MakeRequest("GET", "iam/v1/userAccounts/"+userAccountID, nil)
}
