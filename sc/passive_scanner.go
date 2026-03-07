
package sc

import (
	"encoding/json"
	"fmt"
)

// PassiveScannerService handles passiveScanner operations.
type PassiveScannerService struct {
	client *Client
}

// PassiveScanner represents a passiveScanner resource.
type PassiveScanner struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PassiveScannerListResponse represents the response from listing passiveScanners.
type PassiveScannerListResponse struct {
	Usable     []PassiveScanner `json:"usable"`
	Manageable []PassiveScanner `json:"manageable"`
}

// PassiveScannerCreateInput represents the request body for creating a passiveScanner.
type PassiveScannerCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// PassiveScannerUpdateInput represents the request body for updating a passiveScanner.
type PassiveScannerUpdateInput = PassiveScannerCreateInput

// List returns all passiveScanners.
func (s *PassiveScannerService) List() (*PassiveScannerListResponse, error) {
	resp, err := s.client.get("/passivescanner")
	if err != nil {
		return nil, fmt.Errorf("sc: list passiveScanners: %w", err)
	}

	var result PassiveScannerListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal passiveScanner list response: %w", err)
	}

	return &result, nil
}

// Create creates a new passiveScanner.
func (s *PassiveScannerService) Create(input *PassiveScannerCreateInput) (*PassiveScanner, error) {
	resp, err := s.client.post("/passivescanner", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create passiveScanner: %w", err)
	}

	var result PassiveScanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal passiveScanner response: %w", err)
	}

	return &result, nil
}

// Get returns the passiveScanner with the given ID.
func (s *PassiveScannerService) Get(id string) (*PassiveScanner, error) {
	resp, err := s.client.get("/passivescanner" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get passiveScanner %s: %w", id, err)
	}

	var result PassiveScanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal passiveScanner response: %w", err)
	}

	return &result, nil
}

// Update updates the passiveScanner with the given ID.
func (s *PassiveScannerService) Update(id string, input *PassiveScannerUpdateInput) (*PassiveScanner, error) {
	resp, err := s.client.patch("/passivescanner" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update passiveScanner %s: %w", id, err)
	}

	var result PassiveScanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal passiveScanner response: %w", err)
	}

	return &result, nil
}

// Delete deletes the passiveScanner with the given ID.
func (s *PassiveScannerService) Delete(id string) error {
	_, err := s.client.delete("/passivescanner" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete passiveScanner %s: %w", id, err)
	}

	return nil
}

// UpdateStatus performs the updateStatus action on the passiveScanner.
func (s *PassiveScannerService) UpdateStatus() (*PassiveScanner, error) {
	resp, err := s.client.post("/passivescanner/updateStatus", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: updateStatus passiveScanner: %w", err)
	}

	var result PassiveScanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal passiveScanner updateStatus response: %w", err)
	}

	return &result, nil
}

