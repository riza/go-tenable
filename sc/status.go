package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// StatusService handles status operations.
type StatusService struct {
	client *Client
}

// Status represents a status resource.
type Status struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// StatusListResponse represents the response from listing statuses.
type StatusListResponse struct {
	Usable     []Status `json:"usable"`
	Manageable []Status `json:"manageable"`
}

// List returns all statuses.
func (s *StatusService) List(ctx context.Context) (*StatusListResponse, error) {
	resp, err := s.client.get(ctx, "/status")
	if err != nil {
		return nil, fmt.Errorf("sc: list statuses: %w", err)
	}

	var result StatusListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal status list response: %w", err)
	}

	return &result, nil
}
