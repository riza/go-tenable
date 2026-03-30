package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// ScanZoneService handles scanZone operations.
type ScanZoneService struct {
	client *Client
}

// ScanZone represents a scanZone resource.
type ScanZone struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ScanZoneListResponse represents the response from listing scanZones.
type ScanZoneListResponse struct {
	Usable     []ScanZone `json:"usable"`
	Manageable []ScanZone `json:"manageable"`
}

// ScanZoneCreateInput represents the request body for creating a scanZone.
type ScanZoneCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// ScanZoneUpdateInput represents the request body for updating a scanZone.
type ScanZoneUpdateInput = ScanZoneCreateInput

// List returns all scanZones.
func (s *ScanZoneService) List(ctx context.Context) (*ScanZoneListResponse, error) {
	resp, err := s.client.get(ctx, "/zone")
	if err != nil {
		return nil, fmt.Errorf("sc: list scanZones: %w", err)
	}

	var result ScanZoneListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanZone list response: %w", err)
	}

	return &result, nil
}

// Create creates a new scanZone.
func (s *ScanZoneService) Create(ctx context.Context, input *ScanZoneCreateInput) (*ScanZone, error) {
	resp, err := s.client.post(ctx, "/zone", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create scanZone: %w", err)
	}

	var result ScanZone
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanZone response: %w", err)
	}

	return &result, nil
}

// Get returns the scanZone with the given ID.
func (s *ScanZoneService) Get(ctx context.Context, id string) (*ScanZone, error) {
	resp, err := s.client.get(ctx, "/zone"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get scanZone %s: %w", id, err)
	}

	var result ScanZone
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanZone response: %w", err)
	}

	return &result, nil
}

// Update updates the scanZone with the given ID.
func (s *ScanZoneService) Update(ctx context.Context, id string, input *ScanZoneUpdateInput) (*ScanZone, error) {
	resp, err := s.client.patch(ctx, "/zone"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update scanZone %s: %w", id, err)
	}

	var result ScanZone
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanZone response: %w", err)
	}

	return &result, nil
}

// Delete deletes the scanZone with the given ID.
func (s *ScanZoneService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/zone"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete scanZone %s: %w", id, err)
	}

	return nil
}
