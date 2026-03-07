
package sc

import (
	"encoding/json"
	"fmt"
)

// ScanPolicyTemplateService handles scanPolicyTemplate operations.
type ScanPolicyTemplateService struct {
	client *Client
}

// ScanPolicyTemplate represents a scanPolicyTemplate resource.
type ScanPolicyTemplate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ScanPolicyTemplateListResponse represents the response from listing scanPolicyTemplates.
type ScanPolicyTemplateListResponse struct {
	Usable     []ScanPolicyTemplate `json:"usable"`
	Manageable []ScanPolicyTemplate `json:"manageable"`
}

// List returns all scanPolicyTemplates.
func (s *ScanPolicyTemplateService) List() (*ScanPolicyTemplateListResponse, error) {
	resp, err := s.client.get("/policyTemplate")
	if err != nil {
		return nil, fmt.Errorf("sc: list scanPolicyTemplates: %w", err)
	}

	var result ScanPolicyTemplateListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanPolicyTemplate list response: %w", err)
	}

	return &result, nil
}

// Get returns the scanPolicyTemplate with the given ID.
func (s *ScanPolicyTemplateService) Get(id string) (*ScanPolicyTemplate, error) {
	resp, err := s.client.get("/policyTemplate" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get scanPolicyTemplate %s: %w", id, err)
	}

	var result ScanPolicyTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanPolicyTemplate response: %w", err)
	}

	return &result, nil
}

