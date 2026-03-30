package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// SoftwareUpdateService handles softwareUpdate operations.
type SoftwareUpdateService struct {
	client *Client
}

// SoftwareUpdate represents a softwareUpdate resource.
type SoftwareUpdate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// SoftwareUpdateListResponse represents the response from listing softwareUpdates.
type SoftwareUpdateListResponse struct {
	Usable     []SoftwareUpdate `json:"usable"`
	Manageable []SoftwareUpdate `json:"manageable"`
}

// SoftwareUpdateUpdateInput represents the request body for updating a softwareUpdate.
type SoftwareUpdateUpdateInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// List returns all softwareUpdates.
func (s *SoftwareUpdateService) List(ctx context.Context) (*SoftwareUpdateListResponse, error) {
	resp, err := s.client.get(ctx, "/softwareUpdate")
	if err != nil {
		return nil, fmt.Errorf("sc: list softwareUpdates: %w", err)
	}

	var result SoftwareUpdateListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal softwareUpdate list response: %w", err)
	}

	return &result, nil
}

// Update updates the softwareUpdate.
func (s *SoftwareUpdateService) Update(ctx context.Context, input *SoftwareUpdateUpdateInput) (*SoftwareUpdate, error) {
	resp, err := s.client.patch(ctx, "/softwareUpdate", input)
	if err != nil {
		return nil, fmt.Errorf("sc: update softwareUpdate: %w", err)
	}

	var result SoftwareUpdate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal softwareUpdate response: %w", err)
	}

	return &result, nil
}
