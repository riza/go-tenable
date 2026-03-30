package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DirectorRepositoryService handles directorRepository operations.
type DirectorRepositoryService struct {
	client *Client
}

// DirectorRepository represents a directorRepository resource.
type DirectorRepository struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DirectorRepositoryListResponse represents the response from listing directorRepositorys.
type DirectorRepositoryListResponse struct {
	Usable     []DirectorRepository `json:"usable"`
	Manageable []DirectorRepository `json:"manageable"`
}

// List returns all directorRepositorys.
func (s *DirectorRepositoryService) List(ctx context.Context) (*DirectorRepositoryListResponse, error) {
	resp, err := s.client.get(ctx, "/mgmt/repository")
	if err != nil {
		return nil, fmt.Errorf("sc: list directorRepositorys: %w", err)
	}

	var result DirectorRepositoryListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorRepository list response: %w", err)
	}

	return &result, nil
}

// Get returns the directorRepository with the given ID.
func (s *DirectorRepositoryService) Get(ctx context.Context, id string) (*DirectorRepository, error) {
	resp, err := s.client.get(ctx, "/mgmt/repository"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get directorRepository %s: %w", id, err)
	}

	var result DirectorRepository
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorRepository response: %w", err)
	}

	return &result, nil
}
