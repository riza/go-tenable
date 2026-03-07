
package sc

import (
	"encoding/json"
	"fmt"
)

// DirectorScanPolicyService handles directorScanPolicy operations.
type DirectorScanPolicyService struct {
	client *Client
}

// DirectorScanPolicy represents a directorScanPolicy resource.
type DirectorScanPolicy struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DirectorScanPolicyListResponse represents the response from listing directorScanPolicys.
type DirectorScanPolicyListResponse struct {
	Usable     []DirectorScanPolicy `json:"usable"`
	Manageable []DirectorScanPolicy `json:"manageable"`
}

// DirectorScanPolicyCreateInput represents the request body for creating a directorScanPolicy.
type DirectorScanPolicyCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// List returns all directorScanPolicys.
func (s *DirectorScanPolicyService) List() (*DirectorScanPolicyListResponse, error) {
	resp, err := s.client.get("/mgmt/policy")
	if err != nil {
		return nil, fmt.Errorf("sc: list directorScanPolicys: %w", err)
	}

	var result DirectorScanPolicyListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanPolicy list response: %w", err)
	}

	return &result, nil
}

// Create creates a new directorScanPolicy.
func (s *DirectorScanPolicyService) Create(input *DirectorScanPolicyCreateInput) (*DirectorScanPolicy, error) {
	resp, err := s.client.post("/mgmt/policy", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create directorScanPolicy: %w", err)
	}

	var result DirectorScanPolicy
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanPolicy response: %w", err)
	}

	return &result, nil
}

// Get returns the directorScanPolicy with the given ID.
func (s *DirectorScanPolicyService) Get(id string) (*DirectorScanPolicy, error) {
	resp, err := s.client.get("/mgmt/policy" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get directorScanPolicy %s: %w", id, err)
	}

	var result DirectorScanPolicy
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanPolicy response: %w", err)
	}

	return &result, nil
}

// Delete deletes the directorScanPolicy with the given ID.
func (s *DirectorScanPolicyService) Delete(id string) error {
	_, err := s.client.delete("/mgmt/policy" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete directorScanPolicy %s: %w", id, err)
	}

	return nil
}

