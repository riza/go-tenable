package platform

import (
	"context"
	"encoding/json"
)

// ProfilesService handles communication with the Profiles related endpoints of the Platform API.
type ProfilesService struct {
	client *Client
}

// Profile represents a scan profile.
type Profile struct {
	Id          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Settings    map[string]interface{} `json:"settings,omitempty"`
	CreatedAt   string                 `json:"created_time,omitempty"`
	UpdatedAt   string                 `json:"updated_time,omitempty"`
}

// ProfilesListResponse represents the response from listing profiles.
type ProfilesListResponse struct {
	Profiles []Profile `json:"profiles,omitempty"`
	Total    int       `json:"total,omitempty"`
}

// ListProfiles returns all profiles.
func (s *ProfilesService) ListProfiles(ctx context.Context) (*ProfilesListResponse, error) {
	resp, err := s.client.get(ctx, "/profiles")
	if err != nil {
		return nil, err
	}

	var result ProfilesListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
