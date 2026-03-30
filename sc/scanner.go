package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// ScannerService handles scanner operations.
type ScannerService struct {
	client *Client
}

// Scanner represents a scanner resource.
type Scanner struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ScannerListResponse represents the response from listing scanners.
type ScannerListResponse struct {
	Usable     []Scanner `json:"usable"`
	Manageable []Scanner `json:"manageable"`
}

// ScannerCreateInput represents the request body for creating a scanner.
type ScannerCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// ScannerUpdateInput represents the request body for updating a scanner.
type ScannerUpdateInput = ScannerCreateInput

// List returns all scanners.
func (s *ScannerService) List(ctx context.Context) (*ScannerListResponse, error) {
	resp, err := s.client.get(ctx, "/scanner")
	if err != nil {
		return nil, fmt.Errorf("sc: list scanners: %w", err)
	}

	var result ScannerListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanner list response: %w", err)
	}

	return &result, nil
}

// Create creates a new scanner.
func (s *ScannerService) Create(ctx context.Context, input *ScannerCreateInput) (*Scanner, error) {
	resp, err := s.client.post(ctx, "/scanner", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create scanner: %w", err)
	}

	var result Scanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanner response: %w", err)
	}

	return &result, nil
}

// Get returns the scanner with the given ID.
func (s *ScannerService) Get(ctx context.Context, id string) (*Scanner, error) {
	resp, err := s.client.get(ctx, "/scanner"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get scanner %s: %w", id, err)
	}

	var result Scanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanner response: %w", err)
	}

	return &result, nil
}

// Update updates the scanner with the given ID.
func (s *ScannerService) Update(ctx context.Context, id string, input *ScannerUpdateInput) (*Scanner, error) {
	resp, err := s.client.patch(ctx, "/scanner"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update scanner %s: %w", id, err)
	}

	var result Scanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanner response: %w", err)
	}

	return &result, nil
}

// Delete deletes the scanner with the given ID.
func (s *ScannerService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/scanner"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete scanner %s: %w", id, err)
	}

	return nil
}

// TestScansQuery performs the testScansQuery action on the scanner with the given ID.
func (s *ScannerService) TestScansQuery(ctx context.Context, id string) (*Scanner, error) {
	resp, err := s.client.post(ctx, "/scanner"+"/"+id+"/testScansQuery", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: testScansQuery scanner %s: %w", id, err)
	}

	var result Scanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanner testScansQuery response: %w", err)
	}

	return &result, nil
}

// BugReport performs the bug-report action on the scanner with the given ID.
func (s *ScannerService) BugReport(ctx context.Context, id string) (*Scanner, error) {
	resp, err := s.client.post(ctx, "/scanner"+"/"+id+"/bug-report", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: bugReport scanner %s: %w", id, err)
	}

	var result Scanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanner bugReport response: %w", err)
	}

	return &result, nil
}

// Health performs the health action on the scanner with the given ID.
func (s *ScannerService) Health(ctx context.Context, id string) (*Scanner, error) {
	resp, err := s.client.get(ctx, "/scanner"+"/"+id+"/health")
	if err != nil {
		return nil, fmt.Errorf("sc: health scanner %s: %w", id, err)
	}

	var result Scanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanner health response: %w", err)
	}

	return &result, nil
}

// UpdateStatus performs the updateStatus action on the scanner.
func (s *ScannerService) UpdateStatus(ctx context.Context) (*Scanner, error) {
	resp, err := s.client.post(ctx, "/scanner/updateStatus", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: updateStatus scanner: %w", err)
	}

	var result Scanner
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scanner updateStatus response: %w", err)
	}

	return &result, nil
}
