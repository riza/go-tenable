package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DirectorScanService handles directorScan operations.
type DirectorScanService struct {
	client *Client
}

// DirectorScan represents a directorScan resource.
type DirectorScan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DirectorScanListResponse represents the response from listing directorScans.
type DirectorScanListResponse struct {
	Usable     []DirectorScan `json:"usable"`
	Manageable []DirectorScan `json:"manageable"`
}

// DirectorScanCreateInput represents the request body for creating a directorScan.
type DirectorScanCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// List returns all directorScans.
func (s *DirectorScanService) List(ctx context.Context) (*DirectorScanListResponse, error) {
	resp, err := s.client.get(ctx, "/mgmt/scan")
	if err != nil {
		return nil, fmt.Errorf("sc: list directorScans: %w", err)
	}

	var result DirectorScanListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScan list response: %w", err)
	}

	return &result, nil
}

// Create creates a new directorScan.
func (s *DirectorScanService) Create(ctx context.Context, input *DirectorScanCreateInput) (*DirectorScan, error) {
	resp, err := s.client.post(ctx, "/mgmt/scan", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create directorScan: %w", err)
	}

	var result DirectorScan
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScan response: %w", err)
	}

	return &result, nil
}

// Get returns the directorScan with the given ID.
func (s *DirectorScanService) Get(ctx context.Context, id string) (*DirectorScan, error) {
	resp, err := s.client.get(ctx, "/mgmt/scan"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get directorScan %s: %w", id, err)
	}

	var result DirectorScan
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScan response: %w", err)
	}

	return &result, nil
}

// Delete deletes the directorScan with the given ID.
func (s *DirectorScanService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/mgmt/scan"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete directorScan %s: %w", id, err)
	}

	return nil
}
