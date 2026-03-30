package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// RepositoryService handles communication with the repository-related endpoints of the SC API.
type RepositoryService struct {
	client *Client
}

// Repository represents a repository resource from the API.
type Repository struct {
	ID             string                   `json:"id"`
	Name           string                   `json:"name"`
	Description    string                   `json:"description"`
	Type           string                   `json:"type"`
	DataFormat     string                   `json:"dataFormat"`
	RemoteID       string                   `json:"remoteID"`
	RemoteIP       string                   `json:"remoteIP"`
	Running        string                   `json:"running"`
	DownloadFormat string                   `json:"downloadFormat"`
	LastSyncTime   string                   `json:"lastSyncTime"`
	CreatedTime    string                   `json:"createdTime"`
	ModifiedTime   string                   `json:"modifiedTime"`
	Organizations  []RepositoryOrganization `json:"organizations,omitempty"`
	TypeFields     json.RawMessage          `json:"typeFields,omitempty"`
	LuminFields    json.RawMessage          `json:"luminFields,omitempty"`
	UUID           string                   `json:"uuid"`
}

// RepositoryOrganization represents an organization entry within a repository.
type RepositoryOrganization struct {
	ID          string `json:"id"`
	GroupAssign string `json:"groupAssign"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UUID        string `json:"uuid"`
}

// RepositoryCreateInput contains the fields for creating a new repository.
type RepositoryCreateInput struct {
	Name           string          `json:"name"`
	Type           string          `json:"type,omitempty"`
	Description    string          `json:"description,omitempty"`
	DataFormat     string          `json:"dataFormat,omitempty"`
	TypeFields     json.RawMessage `json:"typeFields,omitempty"`
	Organizations  []IDRef         `json:"organizations,omitempty"`
	RemoteID       string          `json:"remoteID,omitempty"`
	RemoteIP       string          `json:"remoteIP,omitempty"`
	DownloadFormat string          `json:"downloadFormat,omitempty"`
}

// RepositoryUpdateInput contains the fields for updating an existing repository.
type RepositoryUpdateInput = RepositoryCreateInput

// RepositoryImportInput contains the fields for importing data into a repository.
type RepositoryImportInput struct {
	File string `json:"file"`
}

// RepositoryAuthorizeInput contains the fields for authorizing access to a remote repository.
type RepositoryAuthorizeInput struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// List returns the list of repositories. Unlike most resources, the repository
// list endpoint returns a direct array rather than usable/manageable groups.
func (s *RepositoryService) List(ctx context.Context) ([]Repository, error) {
	resp, err := s.client.get(ctx, "/repository")
	if err != nil {
		return nil, fmt.Errorf("sc: list repositories: %w", err)
	}

	var result []Repository
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal repository list response: %w", err)
	}

	return result, nil
}

// Get returns the repository with the given ID.
func (s *RepositoryService) Get(ctx context.Context, id string) (*Repository, error) {
	resp, err := s.client.get(ctx, "/repository/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get repository %s: %w", id, err)
	}

	var repo Repository
	if err := json.Unmarshal(resp.Response, &repo); err != nil {
		return nil, fmt.Errorf("sc: unmarshal repository response: %w", err)
	}

	return &repo, nil
}

// Create creates a new repository with the given input.
func (s *RepositoryService) Create(ctx context.Context, input *RepositoryCreateInput) (*Repository, error) {
	resp, err := s.client.post(ctx, "/repository", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create repository: %w", err)
	}

	var repo Repository
	if err := json.Unmarshal(resp.Response, &repo); err != nil {
		return nil, fmt.Errorf("sc: unmarshal repository response: %w", err)
	}

	return &repo, nil
}

// Update updates an existing repository with the given input.
func (s *RepositoryService) Update(ctx context.Context, id string, input *RepositoryUpdateInput) (*Repository, error) {
	resp, err := s.client.patch(ctx, "/repository/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update repository %s: %w", id, err)
	}

	var repo Repository
	if err := json.Unmarshal(resp.Response, &repo); err != nil {
		return nil, fmt.Errorf("sc: unmarshal repository response: %w", err)
	}

	return &repo, nil
}

// Delete deletes the repository with the given ID.
func (s *RepositoryService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/repository/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete repository %s: %w", id, err)
	}

	return nil
}

// AcceptRiskRules returns the accept risk rules for the repository with the given ID.
func (s *RepositoryService) AcceptRiskRules(ctx context.Context, id string) (json.RawMessage, error) {
	resp, err := s.client.get(ctx, "/repository/"+id+"/acceptRiskRule")
	if err != nil {
		return nil, fmt.Errorf("sc: list accept risk rules for repository %s: %w", id, err)
	}

	return resp.Response, nil
}

// RecastRiskRules returns the recast risk rules for the repository with the given ID.
func (s *RepositoryService) RecastRiskRules(ctx context.Context, id string) (json.RawMessage, error) {
	resp, err := s.client.get(ctx, "/repository/"+id+"/recastRiskRule")
	if err != nil {
		return nil, fmt.Errorf("sc: list recast risk rules for repository %s: %w", id, err)
	}

	return resp.Response, nil
}

// Import imports data into the repository with the given ID.
func (s *RepositoryService) Import(ctx context.Context, id string, input *RepositoryImportInput) (json.RawMessage, error) {
	resp, err := s.client.post(ctx, "/repository/"+id+"/import", input)
	if err != nil {
		return nil, fmt.Errorf("sc: import repository %s: %w", id, err)
	}

	return resp.Response, nil
}

// Export exports data from the repository with the given ID. The response is binary data.
func (s *RepositoryService) Export(ctx context.Context, id string) (json.RawMessage, error) {
	resp, err := s.client.get(ctx, "/repository/"+id+"/export")
	if err != nil {
		return nil, fmt.Errorf("sc: export repository %s: %w", id, err)
	}

	return resp.Response, nil
}

// Sync triggers a sync of the repository with the given ID.
func (s *RepositoryService) Sync(ctx context.Context, id string) (json.RawMessage, error) {
	resp, err := s.client.post(ctx, "/repository/"+id+"/sync", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: sync repository %s: %w", id, err)
	}

	return resp.Response, nil
}

// UpdateMobileData triggers a mobile data update for the repository with the given ID.
func (s *RepositoryService) UpdateMobileData(ctx context.Context, id string) (json.RawMessage, error) {
	resp, err := s.client.post(ctx, "/repository/"+id+"/updateMobileData", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: update mobile data for repository %s: %w", id, err)
	}

	return resp.Response, nil
}

// DeviceInfo returns device information for the given host UUID within the repository.
func (s *RepositoryService) DeviceInfo(ctx context.Context, id string, hostUUID string) (json.RawMessage, error) {
	resp, err := s.client.get(ctx, "/repository/"+id+"/deviceInfo?uuid="+hostUUID)
	if err != nil {
		return nil, fmt.Errorf("sc: get device info for repository %s: %w", id, err)
	}

	return resp.Response, nil
}

// Authorize authorizes access to a remote repository.
func (s *RepositoryService) Authorize(ctx context.Context, input *RepositoryAuthorizeInput) (json.RawMessage, error) {
	resp, err := s.client.post(ctx, "/repository/authorize", input)
	if err != nil {
		return nil, fmt.Errorf("sc: authorize repository: %w", err)
	}

	return resp.Response, nil
}

// FetchRemote fetches available remote repositories from the given host.
func (s *RepositoryService) FetchRemote(ctx context.Context, host string) (json.RawMessage, error) {
	resp, err := s.client.get(ctx, "/repository/fetchRemote?host="+host)
	if err != nil {
		return nil, fmt.Errorf("sc: fetch remote repositories: %w", err)
	}

	return resp.Response, nil
}
