package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DirectorUserService handles directorUser operations.
type DirectorUserService struct {
	client *Client
}

// DirectorUser represents a directorUser resource.
type DirectorUser struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Get returns the directorUser with the given ID.
func (s *DirectorUserService) Get(ctx context.Context, id string) (*DirectorUser, error) {
	resp, err := s.client.get(ctx, "/mgmt/user"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get directorUser %s: %w", id, err)
	}

	var result DirectorUser
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorUser response: %w", err)
	}

	return &result, nil
}
