package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DirectorScanZoneService handles directorScanZone operations.
type DirectorScanZoneService struct {
	client *Client
}

// DirectorScanZone represents a directorScanZone resource.
type DirectorScanZone struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DirectorScanZoneListResponse represents the response from listing directorScanZones.
type DirectorScanZoneListResponse struct {
	Usable     []DirectorScanZone `json:"usable"`
	Manageable []DirectorScanZone `json:"manageable"`
}

// DirectorScanZoneCreateInput represents the request body for creating a directorScanZone.
type DirectorScanZoneCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// DirectorScanZoneUpdateInput represents the request body for updating a directorScanZone.
type DirectorScanZoneUpdateInput = DirectorScanZoneCreateInput

// List returns all directorScanZones.
func (s *DirectorScanZoneService) List(ctx context.Context) (*DirectorScanZoneListResponse, error) {
	resp, err := s.client.get(ctx, "/mgmt/zone")
	if err != nil {
		return nil, fmt.Errorf("sc: list directorScanZones: %w", err)
	}

	var result DirectorScanZoneListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanZone list response: %w", err)
	}

	return &result, nil
}

// Create creates a new directorScanZone.
func (s *DirectorScanZoneService) Create(ctx context.Context, input *DirectorScanZoneCreateInput) (*DirectorScanZone, error) {
	resp, err := s.client.post(ctx, "/mgmt/zone", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create directorScanZone: %w", err)
	}

	var result DirectorScanZone
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanZone response: %w", err)
	}

	return &result, nil
}

// Update updates the directorScanZone with the given ID.
func (s *DirectorScanZoneService) Update(ctx context.Context, id string, input *DirectorScanZoneUpdateInput) (*DirectorScanZone, error) {
	resp, err := s.client.patch(ctx, "/mgmt/zone"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update directorScanZone %s: %w", id, err)
	}

	var result DirectorScanZone
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanZone response: %w", err)
	}

	return &result, nil
}

// Delete deletes the directorScanZone with the given ID.
func (s *DirectorScanZoneService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/mgmt/zone"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete directorScanZone %s: %w", id, err)
	}

	return nil
}

// Get returns the directorScanZone with the given ID.
func (s *DirectorScanZoneService) Get(ctx context.Context, id string) (*DirectorScanZone, error) {
	resp, err := s.client.get(ctx, "/mgmt/zone"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get directorScanZone %s: %w", id, err)
	}

	var result DirectorScanZone
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanZone response: %w", err)
	}

	return &result, nil
}
