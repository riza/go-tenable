package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// StyleService handles style operations.
type StyleService struct {
	client *Client
}

// Style represents a style resource.
type Style struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// StyleListResponse represents the response from listing styles.
type StyleListResponse struct {
	Usable     []Style `json:"usable"`
	Manageable []Style `json:"manageable"`
}

// List returns all styles.
func (s *StyleService) List(ctx context.Context) (*StyleListResponse, error) {
	resp, err := s.client.get(ctx, "/style")
	if err != nil {
		return nil, fmt.Errorf("sc: list styles: %w", err)
	}

	var result StyleListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal style list response: %w", err)
	}

	return &result, nil
}

// Get returns the style with the given ID.
func (s *StyleService) Get(ctx context.Context, id string) (*Style, error) {
	resp, err := s.client.get(ctx, "/style"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get style %s: %w", id, err)
	}

	var result Style
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal style response: %w", err)
	}

	return &result, nil
}
