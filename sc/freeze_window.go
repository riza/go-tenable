package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// FreezeWindowService handles freezeWindow operations.
type FreezeWindowService struct {
	client *Client
}

// FreezeWindow represents a freezeWindow resource.
type FreezeWindow struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// FreezeWindowListResponse represents the response from listing freezeWindows.
type FreezeWindowListResponse struct {
	Usable     []FreezeWindow `json:"usable"`
	Manageable []FreezeWindow `json:"manageable"`
}

// FreezeWindowCreateInput represents the request body for creating a freezeWindow.
type FreezeWindowCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// FreezeWindowUpdateInput represents the request body for updating a freezeWindow.
type FreezeWindowUpdateInput = FreezeWindowCreateInput

// List returns all freezeWindows.
func (s *FreezeWindowService) List(ctx context.Context) (*FreezeWindowListResponse, error) {
	resp, err := s.client.get(ctx, "/freeze")
	if err != nil {
		return nil, fmt.Errorf("sc: list freezeWindows: %w", err)
	}

	var result FreezeWindowListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal freezeWindow list response: %w", err)
	}

	return &result, nil
}

// Create creates a new freezeWindow.
func (s *FreezeWindowService) Create(ctx context.Context, input *FreezeWindowCreateInput) (*FreezeWindow, error) {
	resp, err := s.client.post(ctx, "/freeze", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create freezeWindow: %w", err)
	}

	var result FreezeWindow
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal freezeWindow response: %w", err)
	}

	return &result, nil
}

// Get returns the freezeWindow with the given ID.
func (s *FreezeWindowService) Get(ctx context.Context, id string) (*FreezeWindow, error) {
	resp, err := s.client.get(ctx, "/freeze"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get freezeWindow %s: %w", id, err)
	}

	var result FreezeWindow
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal freezeWindow response: %w", err)
	}

	return &result, nil
}

// Update updates the freezeWindow with the given ID.
func (s *FreezeWindowService) Update(ctx context.Context, id string, input *FreezeWindowUpdateInput) (*FreezeWindow, error) {
	resp, err := s.client.patch(ctx, "/freeze"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update freezeWindow %s: %w", id, err)
	}

	var result FreezeWindow
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal freezeWindow response: %w", err)
	}

	return &result, nil
}

// Delete deletes the freezeWindow with the given ID.
func (s *FreezeWindowService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/freeze"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete freezeWindow %s: %w", id, err)
	}

	return nil
}
