
package sc

import (
	"encoding/json"
	"fmt"
)

// PublishingSiteService handles publishingSite operations.
type PublishingSiteService struct {
	client *Client
}

// PublishingSite represents a publishingSite resource.
type PublishingSite struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PublishingSiteListResponse represents the response from listing publishingSites.
type PublishingSiteListResponse struct {
	Usable     []PublishingSite `json:"usable"`
	Manageable []PublishingSite `json:"manageable"`
}

// PublishingSiteCreateInput represents the request body for creating a publishingSite.
type PublishingSiteCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// PublishingSiteUpdateInput represents the request body for updating a publishingSite.
type PublishingSiteUpdateInput = PublishingSiteCreateInput

// List returns all publishingSites.
func (s *PublishingSiteService) List() (*PublishingSiteListResponse, error) {
	resp, err := s.client.get("/pubSite")
	if err != nil {
		return nil, fmt.Errorf("sc: list publishingSites: %w", err)
	}

	var result PublishingSiteListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal publishingSite list response: %w", err)
	}

	return &result, nil
}

// Create creates a new publishingSite.
func (s *PublishingSiteService) Create(input *PublishingSiteCreateInput) (*PublishingSite, error) {
	resp, err := s.client.post("/pubSite", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create publishingSite: %w", err)
	}

	var result PublishingSite
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal publishingSite response: %w", err)
	}

	return &result, nil
}

// Get returns the publishingSite with the given ID.
func (s *PublishingSiteService) Get(id string) (*PublishingSite, error) {
	resp, err := s.client.get("/pubSite" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get publishingSite %s: %w", id, err)
	}

	var result PublishingSite
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal publishingSite response: %w", err)
	}

	return &result, nil
}

// Update updates the publishingSite with the given ID.
func (s *PublishingSiteService) Update(id string, input *PublishingSiteUpdateInput) (*PublishingSite, error) {
	resp, err := s.client.patch("/pubSite" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update publishingSite %s: %w", id, err)
	}

	var result PublishingSite
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal publishingSite response: %w", err)
	}

	return &result, nil
}

// Delete deletes the publishingSite with the given ID.
func (s *PublishingSiteService) Delete(id string) error {
	_, err := s.client.delete("/pubSite" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete publishingSite %s: %w", id, err)
	}

	return nil
}

