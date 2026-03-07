
package sc

import (
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

// StatusListResponse represents the response from listing statuss.
type StatusListResponse struct {
	Usable     []Status `json:"usable"`
	Manageable []Status `json:"manageable"`
}

// List returns all statuss.
func (s *StatusService) List() (*StatusListResponse, error) {
	resp, err := s.client.get("/status")
	if err != nil {
		return nil, fmt.Errorf("sc: list statuss: %w", err)
	}

	var result StatusListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal status list response: %w", err)
	}

	return &result, nil
}

