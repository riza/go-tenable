package sc

import (
	"context"
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
func (s *PassiveScannerService) List(ctx context.Context) (*PassiveScannerListResponse, error) {
	resp, err := s.client.get(ctx, "/passivescanner")
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
func (s *PassiveScannerService) Create(ctx context.Context, input *PassiveScannerCreateInput) (*PassiveScanner, error) {
	resp, err := s.client.post(ctx, "/passivescanner", input)
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
func (s *PassiveScannerService) Get(ctx context.Context, id string) (*PassiveScanner, error) {
	resp, err := s.client.get(ctx, "/passivescanner"+"/"+id)
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
func (s *PassiveScannerService) Update(ctx context.Context, id string, input *PassiveScannerUpdateInput) (*PassiveScanner, error) {
	resp, err := s.client.patch(ctx, "/passivescanner"+"/"+id, input)
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
func (s *PassiveScannerService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/passivescanner"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete passiveScanner %s: %w", id, err)
	}

	return nil
}

// UpdateStatus performs the updateStatus action on the passiveScanner.
func (s *PassiveScannerService) UpdateStatus(ctx context.Context) (*PassiveScanner, error) {
	resp, err := s.client.post(ctx, "/passivescanner/updateStatus", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: updateStatus passiveScanner: %w", err)
	}

	var result PassiveScanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal passiveScanner updateStatus response: %w", err)
	}

	return &result, nil
}
