package platform

import (
	"context"
	"encoding/json"
)

// CredentialsService handles communication with the Credentials related endpoints of the Platform API.
type CredentialsService struct {
	client *Client
}

// Credential represents a credential.
type Credential struct {
	Id          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Username    string                 `json:"username,omitempty"`
	AuthType    string                 `json:"auth_type,omitempty"`
	CreatedAt   string                 `json:"created_time,omitempty"`
	ModifiedAt  string                 `json:"modified_time,omitempty"`
	Permissions int                    `json:"permissions,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// CredentialsListResponse represents the response from listing credentials.
type CredentialsListResponse struct {
	Credentials []Credential `json:"credentials,omitempty"`
	Total       int          `json:"total,omitempty"`
}

// ListCredentials returns all credentials.
func (s *CredentialsService) ListCredentials(ctx context.Context) (*CredentialsListResponse, error) {
	resp, err := s.client.get(ctx, "/credentials")
	if err != nil {
		return nil, err
	}

	var result CredentialsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateCredential creates a new credential.
func (s *CredentialsService) CreateCredential(ctx context.Context, req *Credential) (*Credential, error) {
	resp, err := s.client.post(ctx, "/credentials", req)
	if err != nil {
		return nil, err
	}

	var result Credential
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetCredential returns a specific credential by ID.
func (s *CredentialsService) GetCredential(ctx context.Context, credentialId string) (*Credential, error) {
	resp, err := s.client.get(ctx, "/credentials/"+credentialId)
	if err != nil {
		return nil, err
	}

	var result Credential
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateCredential updates a credential.
func (s *CredentialsService) UpdateCredential(ctx context.Context, credentialId string, req *Credential) (*Credential, error) {
	resp, err := s.client.put(ctx, "/credentials/"+credentialId, req)
	if err != nil {
		return nil, err
	}

	var result Credential
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteCredential deletes a credential.
func (s *CredentialsService) DeleteCredential(ctx context.Context, credentialId string) error {
	_, err := s.client.delete(ctx, "/credentials/"+credentialId)
	return err
}

// TagsService handles communication with the Tags related endpoints of the Platform API.
type TagsService struct {
	client *Client
}

// Tag represents a tag.
type Tag struct {
	Id        string   `json:"id,omitempty"`
	Uuid      string   `json:"uuid,omitempty"`
	Key       string   `json:"key,omitempty"`
	Value     string   `json:"value,omitempty"`
	Sources   []string `json:"sources,omitempty"`
	CreatedAt string   `json:"created_time,omitempty"`
	UpdatedAt string   `json:"updated_time,omitempty"`
}

// TagsListResponse represents the response from listing tags.
type TagsListResponse struct {
	Tags  []Tag `json:"tags,omitempty"`
	Total int   `json:"total,omitempty"`
}

// ListTags returns all tags.
func (s *TagsService) ListTags(ctx context.Context) (*TagsListResponse, error) {
	resp, err := s.client.get(ctx, "/tags")
	if err != nil {
		return nil, err
	}

	var result TagsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateTag creates a new tag.
func (s *TagsService) CreateTag(ctx context.Context, req *Tag) (*Tag, error) {
	resp, err := s.client.post(ctx, "/tags", req)
	if err != nil {
		return nil, err
	}

	var result Tag
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTag returns a specific tag by ID.
func (s *TagsService) GetTag(ctx context.Context, tagId string) (*Tag, error) {
	resp, err := s.client.get(ctx, "/tags/"+tagId)
	if err != nil {
		return nil, err
	}

	var result Tag
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateTag updates a tag.
func (s *TagsService) UpdateTag(ctx context.Context, tagId string, req *Tag) (*Tag, error) {
	resp, err := s.client.put(ctx, "/tags/"+tagId, req)
	if err != nil {
		return nil, err
	}

	var result Tag
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteTag deletes a tag.
func (s *TagsService) DeleteTag(ctx context.Context, tagId string) error {
	_, err := s.client.delete(ctx, "/tags/"+tagId)
	return err
}

// NetworksService handles communication with the Networks related endpoints of the Platform API.
type NetworksService struct {
	client *Client
}

// Network represents a network.
type Network struct {
	Id          string `json:"id,omitempty"`
	Uuid        string `json:"uuid,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_time,omitempty"`
	ModifiedAt  string `json:"modified_time,omitempty"`
}

// NetworksListResponse represents the response from listing networks.
type NetworksListResponse struct {
	Networks []Network `json:"networks,omitempty"`
	Total    int       `json:"total,omitempty"`
}

// ListNetworks returns all networks.
func (s *NetworksService) ListNetworks(ctx context.Context) (*NetworksListResponse, error) {
	resp, err := s.client.get(ctx, "/networks")
	if err != nil {
		return nil, err
	}

	var result NetworksListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateNetwork creates a new network.
func (s *NetworksService) CreateNetwork(ctx context.Context, req *Network) (*Network, error) {
	resp, err := s.client.post(ctx, "/networks", req)
	if err != nil {
		return nil, err
	}

	var result Network
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetNetwork returns a specific network by ID.
func (s *NetworksService) GetNetwork(ctx context.Context, networkId string) (*Network, error) {
	resp, err := s.client.get(ctx, "/networks/"+networkId)
	if err != nil {
		return nil, err
	}

	var result Network
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateNetwork updates a network.
func (s *NetworksService) UpdateNetwork(ctx context.Context, networkId string, req *Network) (*Network, error) {
	resp, err := s.client.put(ctx, "/networks/"+networkId, req)
	if err != nil {
		return nil, err
	}

	var result Network
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteNetwork deletes a network.
func (s *NetworksService) DeleteNetwork(ctx context.Context, networkId string) error {
	_, err := s.client.delete(ctx, "/networks/"+networkId)
	return err
}
