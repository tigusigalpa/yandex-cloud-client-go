package resources

import (
	"net/http"

	"github.com/tigusigalpa/yandex-cloud-client-go/auth"
	"github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

// OrganizationResource handles organization-related operations
type OrganizationResource struct {
	*AbstractResource
}

// NewOrganizationResource creates a new organization resource
func NewOrganizationResource(httpClient *http.Client, authManager *auth.IAMTokenManager, baseURI string) *OrganizationResource {
	return &OrganizationResource{
		AbstractResource: NewAbstractResource(httpClient, authManager, baseURI),
	}
}

// List gets list of organizations
func (r *OrganizationResource) List(pageSize *int, pageToken *string) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	if pageSize != nil {
		params["pageSize"] = *pageSize
	}
	if pageToken != nil {
		params["pageToken"] = *pageToken
	}

	query := r.BuildQueryString(params)
	return r.MakeRequest("GET", "organization-manager/v1/organizations"+query, nil)
}

// Get gets organization details
func (r *OrganizationResource) Get(organizationID string) (map[string]interface{}, error) {
	if organizationID == "" {
		return nil, errors.NewValidationError("Organization ID cannot be empty")
	}

	return r.MakeRequest("GET", "organization-manager/v1/organizations/"+organizationID, nil)
}

// Update updates organization
func (r *OrganizationResource) Update(organizationID string, data map[string]interface{}) (map[string]interface{}, error) {
	if organizationID == "" {
		return nil, errors.NewValidationError("Organization ID cannot be empty")
	}

	if len(data) == 0 {
		return nil, errors.NewValidationError("Update data cannot be empty")
	}

	return r.MakeRequest("PATCH", "organization-manager/v1/organizations/"+organizationID, data)
}

// ListAccessBindings lists access bindings for organization
func (r *OrganizationResource) ListAccessBindings(organizationID string, pageSize *int, pageToken *string) (map[string]interface{}, error) {
	if organizationID == "" {
		return nil, errors.NewValidationError("Organization ID cannot be empty")
	}

	params := make(map[string]interface{})
	if pageSize != nil {
		params["pageSize"] = *pageSize
	}
	if pageToken != nil {
		params["pageToken"] = *pageToken
	}

	query := r.BuildQueryString(params)
	return r.MakeRequest("GET", "organization-manager/v1/organizations/"+organizationID+":listAccessBindings"+query, nil)
}

// UpdateAccessBindings updates access bindings for organization
func (r *OrganizationResource) UpdateAccessBindings(organizationID string, accessBindingDeltas []map[string]interface{}) (map[string]interface{}, error) {
	if organizationID == "" {
		return nil, errors.NewValidationError("Organization ID cannot be empty")
	}

	if len(accessBindingDeltas) == 0 {
		return nil, errors.NewValidationError("Access binding deltas cannot be empty")
	}

	body := map[string]interface{}{
		"accessBindingDeltas": accessBindingDeltas,
	}

	return r.MakeRequest("POST", "organization-manager/v1/organizations/"+organizationID+":updateAccessBindings", body)
}

// AddRole adds a role to organization (helper method)
func (r *OrganizationResource) AddRole(organizationID, subjectID, roleID, subjectType string) (map[string]interface{}, error) {
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

	return r.UpdateAccessBindings(organizationID, deltas)
}

// RemoveRole removes a role from organization (helper method)
func (r *OrganizationResource) RemoveRole(organizationID, subjectID, roleID, subjectType string) (map[string]interface{}, error) {
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

	return r.UpdateAccessBindings(organizationID, deltas)
}
