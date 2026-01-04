package resources

import (
	"net/http"

	"github.com/tigusigalpa/yandex-cloud-client-go/auth"
	"github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

// FolderResource handles folder-related operations
type FolderResource struct {
	*AbstractResource
}

// NewFolderResource creates a new folder resource
func NewFolderResource(httpClient *http.Client, authManager *auth.IAMTokenManager, baseURI string) *FolderResource {
	return &FolderResource{
		AbstractResource: NewAbstractResource(httpClient, authManager, baseURI),
	}
}

// List gets list of folders
func (r *FolderResource) List(cloudID string, pageSize *int, pageToken *string) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	params["cloudId"] = cloudID

	if pageSize != nil {
		params["pageSize"] = *pageSize
	}
	if pageToken != nil {
		params["pageToken"] = *pageToken
	}

	query := r.BuildQueryString(params)
	return r.MakeRequest("GET", "resource-manager/v1/folders"+query, nil)
}

// Get gets folder details
func (r *FolderResource) Get(folderID string) (map[string]interface{}, error) {
	if folderID == "" {
		return nil, errors.NewValidationError("Folder ID cannot be empty")
	}

	return r.MakeRequest("GET", "resource-manager/v1/folders/"+folderID, nil)
}

// Create creates a new folder
func (r *FolderResource) Create(cloudID, name string, description *string, labels map[string]string) (map[string]interface{}, error) {
	if cloudID == "" {
		return nil, errors.NewValidationError("Cloud ID cannot be empty")
	}

	if name == "" {
		return nil, errors.NewValidationError("Folder name cannot be empty")
	}

	data := map[string]interface{}{
		"cloudId": cloudID,
		"name":    name,
	}

	if description != nil {
		data["description"] = *description
	}

	if labels != nil {
		data["labels"] = labels
	}

	return r.MakeRequest("POST", "resource-manager/v1/folders", data)
}

// Update updates folder
func (r *FolderResource) Update(folderID string, data map[string]interface{}) (map[string]interface{}, error) {
	if folderID == "" {
		return nil, errors.NewValidationError("Folder ID cannot be empty")
	}

	if len(data) == 0 {
		return nil, errors.NewValidationError("Update data cannot be empty")
	}

	return r.MakeRequest("PATCH", "resource-manager/v1/folders/"+folderID, data)
}

// Delete deletes folder
func (r *FolderResource) Delete(folderID string) (map[string]interface{}, error) {
	if folderID == "" {
		return nil, errors.NewValidationError("Folder ID cannot be empty")
	}

	return r.MakeRequest("DELETE", "resource-manager/v1/folders/"+folderID, nil)
}

// ListOperations lists operations for folder
func (r *FolderResource) ListOperations(folderID string, pageSize *int, pageToken *string) (map[string]interface{}, error) {
	if folderID == "" {
		return nil, errors.NewValidationError("Folder ID cannot be empty")
	}

	params := make(map[string]interface{})
	if pageSize != nil {
		params["pageSize"] = *pageSize
	}
	if pageToken != nil {
		params["pageToken"] = *pageToken
	}

	query := r.BuildQueryString(params)
	return r.MakeRequest("GET", "resource-manager/v1/folders/"+folderID+"/operations"+query, nil)
}

// ListAccessBindings lists access bindings for folder
func (r *FolderResource) ListAccessBindings(folderID string, pageSize *int, pageToken *string) (map[string]interface{}, error) {
	if folderID == "" {
		return nil, errors.NewValidationError("Folder ID cannot be empty")
	}

	params := make(map[string]interface{})
	if pageSize != nil {
		params["pageSize"] = *pageSize
	}
	if pageToken != nil {
		params["pageToken"] = *pageToken
	}

	query := r.BuildQueryString(params)
	return r.MakeRequest("GET", "resource-manager/v1/folders/"+folderID+":listAccessBindings"+query, nil)
}

// UpdateAccessBindings updates access bindings for folder
func (r *FolderResource) UpdateAccessBindings(folderID string, accessBindingDeltas []map[string]interface{}) (map[string]interface{}, error) {
	if folderID == "" {
		return nil, errors.NewValidationError("Folder ID cannot be empty")
	}

	if len(accessBindingDeltas) == 0 {
		return nil, errors.NewValidationError("Access binding deltas cannot be empty")
	}

	body := map[string]interface{}{
		"accessBindingDeltas": accessBindingDeltas,
	}

	return r.MakeRequest("POST", "resource-manager/v1/folders/"+folderID+":updateAccessBindings", body)
}

// AddRole adds a role to folder (helper method)
func (r *FolderResource) AddRole(folderID, subjectID, roleID, subjectType string) (map[string]interface{}, error) {
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

	return r.UpdateAccessBindings(folderID, deltas)
}

// RemoveRole removes a role from folder (helper method)
func (r *FolderResource) RemoveRole(folderID, subjectID, roleID, subjectType string) (map[string]interface{}, error) {
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

	return r.UpdateAccessBindings(folderID, deltas)
}
