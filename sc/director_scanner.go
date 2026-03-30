package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DirectorScannerService handles directorScanner operations.
type DirectorScannerService struct {
	client *Client
}

// DirectorScanner represents a directorScanner resource.
type DirectorScanner struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DirectorScannerListResponse represents the response from listing directorScanners.
type DirectorScannerListResponse struct {
	Usable     []DirectorScanner `json:"usable"`
	Manageable []DirectorScanner `json:"manageable"`
}

// DirectorScannerCreateInput represents the request body for creating a directorScanner.
type DirectorScannerCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// List returns all directorScanners.
func (s *DirectorScannerService) List(ctx context.Context) (*DirectorScannerListResponse, error) {
	resp, err := s.client.get(ctx, "/mgmt/scanner")
	if err != nil {
		return nil, fmt.Errorf("sc: list directorScanners: %w", err)
	}

	var result DirectorScannerListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanner list response: %w", err)
	}

	return &result, nil
}

// Create creates a new directorScanner.
func (s *DirectorScannerService) Create(ctx context.Context, input *DirectorScannerCreateInput) (*DirectorScanner, error) {
	resp, err := s.client.post(ctx, "/mgmt/scanner", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create directorScanner: %w", err)
	}

	var result DirectorScanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanner response: %w", err)
	}

	return &result, nil
}
