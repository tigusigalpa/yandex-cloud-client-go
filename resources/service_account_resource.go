package resources

import (
	"net/http"

	"github.com/tigusigalpa/yandex-cloud-client-go/auth"
	"github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

// ServiceAccountResource handles service account-related operations
type ServiceAccountResource struct {
	*AbstractResource
}

// NewServiceAccountResource creates a new service account resource
func NewServiceAccountResource(httpClient *http.Client, authManager *auth.IAMTokenManager, baseURI string) *ServiceAccountResource {
	return &ServiceAccountResource{
		AbstractResource: NewAbstractResource(httpClient, authManager, baseURI),
	}
}

// List gets list of service accounts in folder
func (r *ServiceAccountResource) List(folderID string, pageSize *int, pageToken *string) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	params["folderId"] = folderID

	if pageSize != nil {
		params["pageSize"] = *pageSize
	}
	if pageToken != nil {
		params["pageToken"] = *pageToken
	}

	query := r.BuildQueryString(params)
	return r.MakeRequest("GET", "iam/v1/serviceAccounts"+query, nil)
}

// Get gets service account details
func (r *ServiceAccountResource) Get(serviceAccountID string) (map[string]interface{}, error) {
	if serviceAccountID == "" {
		return nil, errors.NewValidationError("Service account ID cannot be empty")
	}

	return r.MakeRequest("GET", "iam/v1/serviceAccounts/"+serviceAccountID, nil)
}

// Create creates a new service account
func (r *ServiceAccountResource) Create(folderID, name string, description *string) (map[string]interface{}, error) {
	if folderID == "" {
		return nil, errors.NewValidationError("Folder ID cannot be empty")
	}

	if name == "" {
		return nil, errors.NewValidationError("Service account name cannot be empty")
	}

	data := map[string]interface{}{
		"folderId": folderID,
		"name":     name,
	}

	if description != nil {
		data["description"] = *description
	}

	return r.MakeRequest("POST", "iam/v1/serviceAccounts", data)
}

// Update updates service account
func (r *ServiceAccountResource) Update(serviceAccountID string, data map[string]interface{}) (map[string]interface{}, error) {
	if serviceAccountID == "" {
		return nil, errors.NewValidationError("Service account ID cannot be empty")
	}

	if len(data) == 0 {
		return nil, errors.NewValidationError("Update data cannot be empty")
	}

	return r.MakeRequest("PATCH", "iam/v1/serviceAccounts/"+serviceAccountID, data)
}

// Delete deletes service account
func (r *ServiceAccountResource) Delete(serviceAccountID string) (map[string]interface{}, error) {
	if serviceAccountID == "" {
		return nil, errors.NewValidationError("Service account ID cannot be empty")
	}

	return r.MakeRequest("DELETE", "iam/v1/serviceAccounts/"+serviceAccountID, nil)
}

// ListAccessBindings lists access bindings for service account
func (r *ServiceAccountResource) ListAccessBindings(serviceAccountID string, pageSize *int, pageToken *string) (map[string]interface{}, error) {
	if serviceAccountID == "" {
		return nil, errors.NewValidationError("Service account ID cannot be empty")
	}

	params := make(map[string]interface{})
	if pageSize != nil {
		params["pageSize"] = *pageSize
	}
	if pageToken != nil {
		params["pageToken"] = *pageToken
	}

	query := r.BuildQueryString(params)
	return r.MakeRequest("GET", "iam/v1/serviceAccounts/"+serviceAccountID+":listAccessBindings"+query, nil)
}

// UpdateAccessBindings updates access bindings for service account
func (r *ServiceAccountResource) UpdateAccessBindings(serviceAccountID string, accessBindingDeltas []map[string]interface{}) (map[string]interface{}, error) {
	if serviceAccountID == "" {
		return nil, errors.NewValidationError("Service account ID cannot be empty")
	}

	if len(accessBindingDeltas) == 0 {
		return nil, errors.NewValidationError("Access binding deltas cannot be empty")
	}

	body := map[string]interface{}{
		"accessBindingDeltas": accessBindingDeltas,
	}

	return r.MakeRequest("POST", "iam/v1/serviceAccounts/"+serviceAccountID+":updateAccessBindings", body)
}

// AddRole adds a role to service account (helper method)
func (r *ServiceAccountResource) AddRole(serviceAccountID, subjectID, roleID, subjectType string) (map[string]interface{}, error) {
	if subjectType == "" {
		subjectType = "userAccount"
	}

	deltas := []map[string]interface{}{
		{
			"action": "ADD",
			"accessBinding": map[string]interface{}{
				"roleId": roleID,
				"subject": map[string]interface{}{
					"id":   subjectID,
					"type": subjectType,
				},
			},
		},
	}

	return r.UpdateAccessBindings(serviceAccountID, deltas)
}

// RemoveRole removes a role from service account (helper method)
func (r *ServiceAccountResource) RemoveRole(serviceAccountID, subjectID, roleID, subjectType string) (map[string]interface{}, error) {
	if subjectType == "" {
		subjectType = "userAccount"
	}

	deltas := []map[string]interface{}{
		{
			"action": "REMOVE",
			"accessBinding": map[string]interface{}{
				"roleId": roleID,
				"subject": map[string]interface{}{
					"id":   subjectID,
					"type": subjectType,
				},
			},
		},
	}

	return r.UpdateAccessBindings(serviceAccountID, deltas)
}
