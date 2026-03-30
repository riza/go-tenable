package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// WASScanService handles wASScan operations.
type WASScanService struct {
	client *Client
}

// WASScan represents a wASScan resource.
type WASScan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// WASScanListResponse represents the response from listing wASScans.
type WASScanListResponse struct {
	Usable     []WASScan `json:"usable"`
	Manageable []WASScan `json:"manageable"`
}

// WASScanCreateInput represents the request body for creating a wASScan.
type WASScanCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// WASScanUpdateInput represents the request body for updating a wASScan.
type WASScanUpdateInput = WASScanCreateInput

// List returns all wASScans.
func (s *WASScanService) List(ctx context.Context) (*WASScanListResponse, error) {
	resp, err := s.client.get(ctx, "/wasScan")
	if err != nil {
		return nil, fmt.Errorf("sc: list wASScans: %w", err)
	}

	var result WASScanListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal wASScan list response: %w", err)
	}

	return &result, nil
}

// Create creates a new wASScan.
func (s *WASScanService) Create(ctx context.Context, input *WASScanCreateInput) (*WASScan, error) {
	resp, err := s.client.post(ctx, "/wasScan", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create wASScan: %w", err)
	}

	var result WASScan
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal wASScan response: %w", err)
	}

	return &result, nil
}

// Get returns the wASScan with the given ID.
func (s *WASScanService) Get(ctx context.Context, id string) (*WASScan, error) {
	resp, err := s.client.get(ctx, "/wasScan"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get wASScan %s: %w", id, err)
	}

	var result WASScan
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal wASScan response: %w", err)
	}

	return &result, nil
}

// Update updates the wASScan with the given ID.
func (s *WASScanService) Update(ctx context.Context, id string, input *WASScanUpdateInput) (*WASScan, error) {
	resp, err := s.client.patch(ctx, "/wasScan"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update wASScan %s: %w", id, err)
	}

	var result WASScan
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal wASScan response: %w", err)
	}

	return &result, nil
}

// Delete deletes the wASScan with the given ID.
func (s *WASScanService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/wasScan"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete wASScan %s: %w", id, err)
	}

	return nil
}

// Copy performs the copy action on the wASScan with the given ID.
func (s *WASScanService) Copy(ctx context.Context, id string) (*WASScan, error) {
	resp, err := s.client.post(ctx, "/wasScan"+"/"+id+"/copy", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: copy wASScan %s: %w", id, err)
	}

	var result WASScan
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal wASScan copy response: %w", err)
	}

	return &result, nil
}

// Launch performs the launch action on the wASScan with the given ID.
func (s *WASScanService) Launch(ctx context.Context, id string) (*WASScan, error) {
	resp, err := s.client.post(ctx, "/wasScan"+"/"+id+"/launch", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: launch wASScan %s: %w", id, err)
	}

	var result WASScan
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal wASScan launch response: %w", err)
	}

	return &result, nil
}
