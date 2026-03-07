
package sc

import (
	"encoding/json"
	"fmt"
)

// ScanPolicyService handles scanPolicy operations.
type ScanPolicyService struct {
	client *Client
}

// ScanPolicy represents a scanPolicy resource.
type ScanPolicy struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ScanPolicyListResponse represents the response from listing scanPolicys.
type ScanPolicyListResponse struct {
	Usable     []ScanPolicy `json:"usable"`
	Manageable []ScanPolicy `json:"manageable"`
}

// ScanPolicyCreateInput represents the request body for creating a scanPolicy.
type ScanPolicyCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// ScanPolicyUpdateInput represents the request body for updating a scanPolicy.
type ScanPolicyUpdateInput = ScanPolicyCreateInput

// List returns all scanPolicys.
func (s *ScanPolicyService) List() (*ScanPolicyListResponse, error) {
	resp, err := s.client.get("/policy")
	if err != nil {
		return nil, fmt.Errorf("sc: list scanPolicys: %w", err)
	}

	var result ScanPolicyListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanPolicy list response: %w", err)
	}

	return &result, nil
}

// Create creates a new scanPolicy.
func (s *ScanPolicyService) Create(input *ScanPolicyCreateInput) (*ScanPolicy, error) {
	resp, err := s.client.post("/policy", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create scanPolicy: %w", err)
	}

	var result ScanPolicy
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanPolicy response: %w", err)
	}

	return &result, nil
}

// Get returns the scanPolicy with the given ID.
func (s *ScanPolicyService) Get(id string) (*ScanPolicy, error) {
	resp, err := s.client.get("/policy" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get scanPolicy %s: %w", id, err)
	}

	var result ScanPolicy
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanPolicy response: %w", err)
	}

	return &result, nil
}

// Update updates the scanPolicy with the given ID.
func (s *ScanPolicyService) Update(id string, input *ScanPolicyUpdateInput) (*ScanPolicy, error) {
	resp, err := s.client.patch("/policy" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update scanPolicy %s: %w", id, err)
	}

	var result ScanPolicy
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanPolicy response: %w", err)
	}

	return &result, nil
}

// Delete deletes the scanPolicy with the given ID.
func (s *ScanPolicyService) Delete(id string) error {
	_, err := s.client.delete("/policy" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete scanPolicy %s: %w", id, err)
	}

	return nil
}

// Copy performs the copy action on the scanPolicy with the given ID.
func (s *ScanPolicyService) Copy(id string) (*ScanPolicy, error) {
	resp, err := s.client.post("/policy" + "/" + id + "/copy", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: copy scanPolicy %s: %w", id, err)
	}

	var result ScanPolicy
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanPolicy copy response: %w", err)
	}

	return &result, nil
}

// Export performs the export action on the scanPolicy with the given ID.
func (s *ScanPolicyService) Export(id string) (*ScanPolicy, error) {
	resp, err := s.client.post("/policy" + "/" + id + "/export", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: export scanPolicy %s: %w", id, err)
	}

	var result ScanPolicy
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanPolicy export response: %w", err)
	}

	return &result, nil
}

// Share performs the share action on the scanPolicy with the given ID.
func (s *ScanPolicyService) Share(id string) (*ScanPolicy, error) {
	resp, err := s.client.post("/policy" + "/" + id + "/share", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: share scanPolicy %s: %w", id, err)
	}

	var result ScanPolicy
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanPolicy share response: %w", err)
	}

	return &result, nil
}

// Import performs the import action on the scanPolicy.
func (s *ScanPolicyService) Import() (*ScanPolicy, error) {
	resp, err := s.client.post("/policy/import", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: import scanPolicy: %w", err)
	}

	var result ScanPolicy
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanPolicy import response: %w", err)
	}

	return &result, nil
}

// Tag performs the tag action on the scanPolicy.
func (s *ScanPolicyService) Tag() (*ScanPolicy, error) {
	resp, err := s.client.get("/policy/tag")
	if err != nil {
		return nil, fmt.Errorf("sc: tag scanPolicy: %w", err)
	}

	var result ScanPolicy
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanPolicy tag response: %w", err)
	}

	return &result, nil
}

