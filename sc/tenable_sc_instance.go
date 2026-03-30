package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// TenableSCInstanceService handles tenableSCInstance operations.
type TenableSCInstanceService struct {
	client *Client
}

// TenableSCInstance represents a tenableSCInstance resource.
type TenableSCInstance struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// TenableSCInstanceListResponse represents the response from listing tenableSCInstances.
type TenableSCInstanceListResponse struct {
	Usable     []TenableSCInstance `json:"usable"`
	Manageable []TenableSCInstance `json:"manageable"`
}

// TenableSCInstanceCreateInput represents the request body for creating a tenableSCInstance.
type TenableSCInstanceCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// TenableSCInstanceUpdateInput represents the request body for updating a tenableSCInstance.
type TenableSCInstanceUpdateInput = TenableSCInstanceCreateInput

// List returns all tenableSCInstances.
func (s *TenableSCInstanceService) List(ctx context.Context) (*TenableSCInstanceListResponse, error) {
	resp, err := s.client.get(ctx, "/sci")
	if err != nil {
		return nil, fmt.Errorf("sc: list tenableSCInstances: %w", err)
	}

	var result TenableSCInstanceListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal tenableSCInstance list response: %w", err)
	}

	return &result, nil
}

// Create creates a new tenableSCInstance.
func (s *TenableSCInstanceService) Create(ctx context.Context, input *TenableSCInstanceCreateInput) (*TenableSCInstance, error) {
	resp, err := s.client.post(ctx, "/sci", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create tenableSCInstance: %w", err)
	}

	var result TenableSCInstance
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal tenableSCInstance response: %w", err)
	}

	return &result, nil
}

// Get returns the tenableSCInstance with the given ID.
func (s *TenableSCInstanceService) Get(ctx context.Context, id string) (*TenableSCInstance, error) {
	resp, err := s.client.get(ctx, "/sci"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get tenableSCInstance %s: %w", id, err)
	}

	var result TenableSCInstance
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal tenableSCInstance response: %w", err)
	}

	return &result, nil
}

// Update updates the tenableSCInstance with the given ID.
func (s *TenableSCInstanceService) Update(ctx context.Context, id string, input *TenableSCInstanceUpdateInput) (*TenableSCInstance, error) {
	resp, err := s.client.patch(ctx, "/sci"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update tenableSCInstance %s: %w", id, err)
	}

	var result TenableSCInstance
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal tenableSCInstance response: %w", err)
	}

	return &result, nil
}

// Delete deletes the tenableSCInstance with the given ID.
func (s *TenableSCInstanceService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/sci"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete tenableSCInstance %s: %w", id, err)
	}

	return nil
}
