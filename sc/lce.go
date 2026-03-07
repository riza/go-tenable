
package sc

import (
	"encoding/json"
	"fmt"
)

// LCEService handles lCE operations.
type LCEService struct {
	client *Client
}

// LCE represents a lCE resource.
type LCE struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// LCEListResponse represents the response from listing lCEs.
type LCEListResponse struct {
	Usable     []LCE `json:"usable"`
	Manageable []LCE `json:"manageable"`
}

// LCECreateInput represents the request body for creating a lCE.
type LCECreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// LCEUpdateInput represents the request body for updating a lCE.
type LCEUpdateInput = LCECreateInput

// List returns all lCEs.
func (s *LCEService) List() (*LCEListResponse, error) {
	resp, err := s.client.get("/lce")
	if err != nil {
		return nil, fmt.Errorf("sc: list lCEs: %w", err)
	}

	var result LCEListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lCE list response: %w", err)
	}

	return &result, nil
}

// Create creates a new lCE.
func (s *LCEService) Create(input *LCECreateInput) (*LCE, error) {
	resp, err := s.client.post("/lce", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create lCE: %w", err)
	}

	var result LCE
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lCE response: %w", err)
	}

	return &result, nil
}

// Authorize performs the authorize action on the lCE.
func (s *LCEService) Authorize() (*LCE, error) {
	resp, err := s.client.post("/lce/authorize", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: authorize lCE: %w", err)
	}

	var result LCE
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lCE authorize response: %w", err)
	}

	return &result, nil
}

// Get returns the lCE with the given ID.
func (s *LCEService) Get(id string) (*LCE, error) {
	resp, err := s.client.get("/lce" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get lCE %s: %w", id, err)
	}

	var result LCE
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lCE response: %w", err)
	}

	return &result, nil
}

// Update updates the lCE with the given ID.
func (s *LCEService) Update(id string, input *LCEUpdateInput) (*LCE, error) {
	resp, err := s.client.patch("/lce" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update lCE %s: %w", id, err)
	}

	var result LCE
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lCE response: %w", err)
	}

	return &result, nil
}

// Delete deletes the lCE with the given ID.
func (s *LCEService) Delete(id string) error {
	_, err := s.client.delete("/lce" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete lCE %s: %w", id, err)
	}

	return nil
}

// EventTypes performs the eventTypes action on the lCE.
func (s *LCEService) EventTypes() (*LCE, error) {
	resp, err := s.client.get("/lce/eventTypes")
	if err != nil {
		return nil, fmt.Errorf("sc: eventTypes lCE: %w", err)
	}

	var result LCE
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lCE eventTypes response: %w", err)
	}

	return &result, nil
}

