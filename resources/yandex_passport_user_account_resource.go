package resources

import (
	"net/http"

	"github.com/tigusigalpa/yandex-cloud-client-go/auth"
	"github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

// YandexPassportUserAccountResource handles Yandex Passport user account-related operations
type YandexPassportUserAccountResource struct {
	*AbstractResource
}

// NewYandexPassportUserAccountResource creates a new Yandex Passport user account resource
func NewYandexPassportUserAccountResource(httpClient *http.Client, authManager *auth.IAMTokenManager, baseURI string) *YandexPassportUserAccountResource {
	return &YandexPassportUserAccountResource{
		AbstractResource: NewAbstractResource(httpClient, authManager, baseURI),
	}
}

// GetByLogin gets user account by Yandex Passport login
func (r *YandexPassportUserAccountResource) GetByLogin(login string) (map[string]interface{}, error) {
	if login == "" {
		return nil, errors.NewValidationError("Login cannot be empty")
	}

	return r.MakeRequest("GET", "iam/v1/yandexPassportUserAccounts:byLogin?login="+login, nil)
}
