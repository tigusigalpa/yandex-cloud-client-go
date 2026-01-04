package resources

import (
	"net/http"

	"github.com/tigusigalpa/yandex-cloud-client-go/auth"
	"github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

// CloudResource handles cloud-related operations
type CloudResource struct {
	*AbstractResource
}

// NewCloudResource creates a new cloud resource
func NewCloudResource(httpClient *http.Client, authManager *auth.IAMTokenManager, baseURI string) *CloudResource {
	return &CloudResource{
		AbstractResource: NewAbstractResource(httpClient, authManager, baseURI),
	}
}

// List gets list of clouds
func (r *CloudResource) List(organizationID *string, pageSize *int, pageToken *string) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	if organizationID != nil {
		params["organizationId"] = *organizationID
	}
	if pageSize != nil {
		params["pageSize"] = *pageSize
	}
	if pageToken != nil {
		params["pageToken"] = *pageToken
	}

	query := r.BuildQueryString(params)
	return r.MakeRequest("GET", "resource-manager/v1/clouds"+query, nil)
}

// Get gets cloud details
func (r *CloudResource) Get(cloudID string) (map[string]interface{}, error) {
	if cloudID == "" {
		return nil, errors.NewValidationError("Cloud ID cannot be empty")
	}

	return r.MakeRequest("GET", "resource-manager/v1/clouds/"+cloudID, nil)
}

// Create creates a new cloud
func (r *CloudResource) Create(organizationID, name string, description *string, labels map[string]string) (map[string]interface{}, error) {
	if organizationID == "" {
		return nil, errors.NewValidationError("Organization ID cannot be empty")
	}

	if name == "" {
		return nil, errors.NewValidationError("Cloud name cannot be empty")
	}

	data := map[string]interface{}{
		"organizationId": organizationID,
		"name":           name,
	}

	if description != nil {
		data["description"] = *description
	}

	if labels != nil {
		data["labels"] = labels
	}

	return r.MakeRequest("POST", "resource-manager/v1/clouds", data)
}

// Update updates cloud
func (r *CloudResource) Update(cloudID string, data map[string]interface{}) (map[string]interface{}, error) {
	if cloudID == "" {
		return nil, errors.NewValidationError("Cloud ID cannot be empty")
	}

	if len(data) == 0 {
		return nil, errors.NewValidationError("Update data cannot be empty")
	}

	return r.MakeRequest("PATCH", "resource-manager/v1/clouds/"+cloudID, data)
}

// Delete deletes cloud
func (r *CloudResource) Delete(cloudID string) (map[string]interface{}, error) {
	if cloudID == "" {
		return nil, errors.NewValidationError("Cloud ID cannot be empty")
	}

	return r.MakeRequest("DELETE", "resource-manager/v1/clouds/"+cloudID, nil)
}

// SetAccessBindings sets access bindings for cloud
func (r *CloudResource) SetAccessBindings(cloudID string, accessBindings []map[string]interface{}) (map[string]interface{}, error) {
	if cloudID == "" {
		return nil, errors.NewValidationError("Cloud ID cannot be empty")
	}

	if len(accessBindings) == 0 {
		return nil, errors.NewValidationError("Access bindings cannot be empty")
	}

	body := map[string]interface{}{
		"accessBindings": accessBindings,
	}

	return r.MakeRequest("POST", "resource-manager/v1/clouds/"+cloudID+":setAccessBindings", body)
}

// ListAccessBindings lists access bindings for cloud
func (r *CloudResource) ListAccessBindings(cloudID string, pageSize *int, pageToken *string) (map[string]interface{}, error) {
	if cloudID == "" {
		return nil, errors.NewValidationError("Cloud ID cannot be empty")
	}

	params := make(map[string]interface{})
	if pageSize != nil {
		params["pageSize"] = *pageSize
	}
	if pageToken != nil {
		params["pageToken"] = *pageToken
	}

	query := r.BuildQueryString(params)
	return r.MakeRequest("GET", "resource-manager/v1/clouds/"+cloudID+":listAccessBindings"+query, nil)
}

// UpdateAccessBindings updates access bindings for cloud
func (r *CloudResource) UpdateAccessBindings(cloudID string, accessBindingDeltas []map[string]interface{}) (map[string]interface{}, error) {
	if cloudID == "" {
		return nil, errors.NewValidationError("Cloud ID cannot be empty")
	}

	if len(accessBindingDeltas) == 0 {
		return nil, errors.NewValidationError("Access binding deltas cannot be empty")
	}

	body := map[string]interface{}{
		"accessBindingDeltas": accessBindingDeltas,
	}

	return r.MakeRequest("POST", "resource-manager/v1/clouds/"+cloudID+":updateAccessBindings", body)
}

// AddRole adds a role to cloud (helper method)
func (r *CloudResource) AddRole(cloudID, subjectID, roleID, subjectType string) (map[string]interface{}, error) {
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

	return r.UpdateAccessBindings(cloudID, deltas)
}

// RemoveRole removes a role from cloud (helper method)
func (r *CloudResource) RemoveRole(cloudID, subjectID, roleID, subjectType string) (map[string]interface{}, error) {
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

	return r.UpdateAccessBindings(cloudID, deltas)
}
