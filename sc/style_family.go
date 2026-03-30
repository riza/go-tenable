package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// StyleFamilyService handles styleFamily operations.
type StyleFamilyService struct {
	client *Client
}

// StyleFamily represents a styleFamily resource.
type StyleFamily struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// StyleFamilyListResponse represents the response from listing styleFamilys.
type StyleFamilyListResponse struct {
	Usable     []StyleFamily `json:"usable"`
	Manageable []StyleFamily `json:"manageable"`
}

// List returns all styleFamilys.
func (s *StyleFamilyService) List(ctx context.Context) (*StyleFamilyListResponse, error) {
	resp, err := s.client.get(ctx, "/styleFamily")
	if err != nil {
		return nil, fmt.Errorf("sc: list styleFamilys: %w", err)
	}

	var result StyleFamilyListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal styleFamily list response: %w", err)
	}

	return &result, nil
}

// Get returns the styleFamily with the given ID.
func (s *StyleFamilyService) Get(ctx context.Context, id string) (*StyleFamily, error) {
	resp, err := s.client.get(ctx, "/styleFamily"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get styleFamily %s: %w", id, err)
	}

	var result StyleFamily
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal styleFamily response: %w", err)
	}

	return &result, nil
}
