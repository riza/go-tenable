
package sc

import (
	"encoding/json"
	"fmt"
)

// ConfigSectionService handles configSection operations.
type ConfigSectionService struct {
	client *Client
}

// ConfigSection represents a configSection resource.
type ConfigSection struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ConfigSectionListResponse represents the response from listing configSections.
type ConfigSectionListResponse struct {
	Usable     []ConfigSection `json:"usable"`
	Manageable []ConfigSection `json:"manageable"`
}

// ConfigSectionUpdateInput represents the request body for updating a configSection.
type ConfigSectionUpdateInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// List returns all configSections.
func (s *ConfigSectionService) List() (*ConfigSectionListResponse, error) {
	resp, err := s.client.get("/configSection")
	if err != nil {
		return nil, fmt.Errorf("sc: list configSections: %w", err)
	}

	var result ConfigSectionListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal configSection list response: %w", err)
	}

	return &result, nil
}

// Get returns the configSection with the given ID.
func (s *ConfigSectionService) Get(id string) (*ConfigSection, error) {
	resp, err := s.client.get("/configSection" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get configSection %s: %w", id, err)
	}

	var result ConfigSection
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal configSection response: %w", err)
	}

	return &result, nil
}

// Update updates the configSection with the given ID.
func (s *ConfigSectionService) Update(id string, input *ConfigSectionUpdateInput) (*ConfigSection, error) {
	resp, err := s.client.patch("/configSection" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update configSection %s: %w", id, err)
	}

	var result ConfigSection
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal configSection response: %w", err)
	}

	return &result, nil
}

