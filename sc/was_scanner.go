
package sc

import (
	"encoding/json"
	"fmt"
)

// WASScannerService handles wASScanner operations.
type WASScannerService struct {
	client *Client
}

// WASScanner represents a wASScanner resource.
type WASScanner struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// WASScannerListResponse represents the response from listing wASScanners.
type WASScannerListResponse struct {
	Usable     []WASScanner `json:"usable"`
	Manageable []WASScanner `json:"manageable"`
}

// WASScannerUpdateInput represents the request body for updating a wASScanner.
type WASScannerUpdateInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// List returns all wASScanners.
func (s *WASScannerService) List() (*WASScannerListResponse, error) {
	resp, err := s.client.get("/wasScanner")
	if err != nil {
		return nil, fmt.Errorf("sc: list wASScanners: %w", err)
	}

	var result WASScannerListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal wASScanner list response: %w", err)
	}

	return &result, nil
}

// Get returns the wASScanner with the given ID.
func (s *WASScannerService) Get(id string) (*WASScanner, error) {
	resp, err := s.client.get("/wasScanner" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get wASScanner %s: %w", id, err)
	}

	var result WASScanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal wASScanner response: %w", err)
	}

	return &result, nil
}

// Update updates the wASScanner with the given ID.
func (s *WASScannerService) Update(id string, input *WASScannerUpdateInput) (*WASScanner, error) {
	resp, err := s.client.patch("/wasScanner" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update wASScanner %s: %w", id, err)
	}

	var result WASScanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal wASScanner response: %w", err)
	}

	return &result, nil
}

// Delete deletes the wASScanner with the given ID.
func (s *WASScannerService) Delete(id string) error {
	_, err := s.client.delete("/wasScanner" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete wASScanner %s: %w", id, err)
	}

	return nil
}

