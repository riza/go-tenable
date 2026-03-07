
package sc

import (
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
func (s *FreezeWindowService) List() (*FreezeWindowListResponse, error) {
	resp, err := s.client.get("/freeze")
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
func (s *FreezeWindowService) Create(input *FreezeWindowCreateInput) (*FreezeWindow, error) {
	resp, err := s.client.post("/freeze", input)
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
func (s *FreezeWindowService) Get(id string) (*FreezeWindow, error) {
	resp, err := s.client.get("/freeze" + "/" + id)
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
func (s *FreezeWindowService) Update(id string, input *FreezeWindowUpdateInput) (*FreezeWindow, error) {
	resp, err := s.client.patch("/freeze" + "/" + id, input)
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
func (s *FreezeWindowService) Delete(id string) error {
	_, err := s.client.delete("/freeze" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete freezeWindow %s: %w", id, err)
	}

	return nil
}

