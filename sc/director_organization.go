package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DirectorOrganizationService handles directorOrganization operations.
type DirectorOrganizationService struct {
	client *Client
}

// DirectorOrganization represents a directorOrganization resource.
type DirectorOrganization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DirectorOrganizationListResponse represents the response from listing directorOrganizations.
type DirectorOrganizationListResponse struct {
	Usable     []DirectorOrganization `json:"usable"`
	Manageable []DirectorOrganization `json:"manageable"`
}

// List returns all directorOrganizations.
func (s *DirectorOrganizationService) List(ctx context.Context) (*DirectorOrganizationListResponse, error) {
	resp, err := s.client.get(ctx, "/mgmt/organization")
	if err != nil {
		return nil, fmt.Errorf("sc: list directorOrganizations: %w", err)
	}

	var result DirectorOrganizationListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorOrganization list response: %w", err)
	}

	return &result, nil
}

// Get returns the directorOrganization with the given ID.
func (s *DirectorOrganizationService) Get(ctx context.Context, id string) (*DirectorOrganization, error) {
	resp, err := s.client.get(ctx, "/mgmt/organization"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get directorOrganization %s: %w", id, err)
	}

	var result DirectorOrganization
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorOrganization response: %w", err)
	}

	return &result, nil
}
