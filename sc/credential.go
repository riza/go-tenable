package sc

import (
	"encoding/json"
	"fmt"
)

// CredentialService handles communication with the credential-related endpoints of the SC API.
type CredentialService struct {
	client *Client
}

// Credential represents a credential resource from the API.
type Credential struct {
	ID           string          `json:"id"`
	Type         string          `json:"type"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Tags         string          `json:"tags"`
	CreatedTime  string          `json:"createdTime"`
	ModifiedTime string          `json:"modifiedTime"`
	TypeFields   json.RawMessage `json:"typeFields,omitempty"`
	Groups       []IDRef         `json:"groups,omitempty"`
	CanUse       string          `json:"canUse"`
	CanManage    string          `json:"canManage"`
	Creator      *IDRef          `json:"creator,omitempty"`
	Owner        *IDRef          `json:"owner,omitempty"`
	OwnerGroup   *IDRef          `json:"ownerGroup,omitempty"`
	TargetGroup  *IDRef          `json:"targetGroup,omitempty"`
	UUID         string          `json:"uuid"`
}

// CredentialListResponse holds the usable and manageable credential arrays returned by the list endpoint.
type CredentialListResponse struct {
	Usable     []Credential `json:"usable"`
	Manageable []Credential `json:"manageable"`
}

// CredentialCreateInput contains the fields for creating a new credential.
type CredentialCreateInput struct {
	Name        string          `json:"name"`
	Type        string          `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	Tags        string          `json:"tags,omitempty"`
	TypeFields  json.RawMessage `json:"typeFields,omitempty"`
	Groups      []IDRef         `json:"groups,omitempty"`
}

// CredentialUpdateInput contains the fields for updating an existing credential.
type CredentialUpdateInput = CredentialCreateInput

// CredentialShareInput contains the fields for sharing a credential with groups.
type CredentialShareInput struct {
	Groups []IDRef `json:"groups"`
}

// List returns the list of credentials (usable and manageable).
func (s *CredentialService) List() (*CredentialListResponse, error) {
	resp, err := s.client.get("/credential")
	if err != nil {
		return nil, fmt.Errorf("sc: list credentials: %w", err)
	}

	var result CredentialListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal credential list response: %w", err)
	}

	return &result, nil
}

// Get returns the credential with the given ID.
func (s *CredentialService) Get(id string) (*Credential, error) {
	resp, err := s.client.get("/credential/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get credential %s: %w", id, err)
	}

	var cred Credential
	if err := json.Unmarshal(resp.Response, &cred); err != nil {
		return nil, fmt.Errorf("sc: unmarshal credential response: %w", err)
	}

	return &cred, nil
}

// Create creates a new credential with the given input.
func (s *CredentialService) Create(input *CredentialCreateInput) (*Credential, error) {
	resp, err := s.client.post("/credential", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create credential: %w", err)
	}

	var cred Credential
	if err := json.Unmarshal(resp.Response, &cred); err != nil {
		return nil, fmt.Errorf("sc: unmarshal credential response: %w", err)
	}

	return &cred, nil
}

// Update updates an existing credential with the given input.
func (s *CredentialService) Update(id string, input *CredentialUpdateInput) (*Credential, error) {
	resp, err := s.client.patch("/credential/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update credential %s: %w", id, err)
	}

	var cred Credential
	if err := json.Unmarshal(resp.Response, &cred); err != nil {
		return nil, fmt.Errorf("sc: unmarshal credential response: %w", err)
	}

	return &cred, nil
}

// Delete deletes the credential with the given ID.
func (s *CredentialService) Delete(id string) error {
	_, err := s.client.delete("/credential/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete credential %s: %w", id, err)
	}

	return nil
}

// Share shares the credential with the given ID to the specified groups.
func (s *CredentialService) Share(id string, input *CredentialShareInput) (*Credential, error) {
	resp, err := s.client.post("/credential/"+id+"/share", input)
	if err != nil {
		return nil, fmt.Errorf("sc: share credential %s: %w", id, err)
	}

	var cred Credential
	if err := json.Unmarshal(resp.Response, &cred); err != nil {
		return nil, fmt.Errorf("sc: unmarshal credential share response: %w", err)
	}

	return &cred, nil
}

// Tags returns the list of credential tags.
func (s *CredentialService) Tags() ([]string, error) {
	resp, err := s.client.get("/credential/tag")
	if err != nil {
		return nil, fmt.Errorf("sc: list credential tags: %w", err)
	}

	var tags []string
	if err := json.Unmarshal(resp.Response, &tags); err != nil {
		return nil, fmt.Errorf("sc: unmarshal credential tags response: %w", err)
	}

	return tags, nil
}
