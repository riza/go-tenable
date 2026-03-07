
package sc

import (
	"encoding/json"
	"fmt"
)

// SolutionsService handles solutions operations.
type SolutionsService struct {
	client *Client
}

// Solutions represents a solutions resource.
type Solutions struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// SolutionsCreateInput represents the request body for creating a solutions.
type SolutionsCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// Create creates a new solutions.
func (s *SolutionsService) Create(input *SolutionsCreateInput) (*Solutions, error) {
	resp, err := s.client.post("/solutions", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create solutions: %w", err)
	}

	var result Solutions
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal solutions response: %w", err)
	}

	return &result, nil
}

// CreateByID performs the  action on the solutions with the given ID.
func (s *SolutionsService) CreateByID(id string) (*Solutions, error) {
	resp, err := s.client.post("/solutions" + "/" + id, nil)
	if err != nil {
		return nil, fmt.Errorf("sc: createByID solutions %s: %w", id, err)
	}

	var result Solutions
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal solutions createByID response: %w", err)
	}

	return &result, nil
}

// Vuln performs the vuln action on the solutions with the given ID.
func (s *SolutionsService) Vuln(id string) (*Solutions, error) {
	resp, err := s.client.post("/solutions" + "/" + id + "/vuln", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: vuln solutions %s: %w", id, err)
	}

	var result Solutions
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal solutions vuln response: %w", err)
	}

	return &result, nil
}

// Asset performs the asset action on the solutions with the given ID.
func (s *SolutionsService) Asset(id string) (*Solutions, error) {
	resp, err := s.client.post("/solutions" + "/" + id + "/asset", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: asset solutions %s: %w", id, err)
	}

	var result Solutions
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal solutions asset response: %w", err)
	}

	return &result, nil
}

