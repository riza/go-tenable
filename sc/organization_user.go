
package sc

import (
	"encoding/json"
	"fmt"
)

// OrganizationUserService handles organizationUser operations.
type OrganizationUserService struct {
	client *Client
}

// OrganizationUser represents a organizationUser resource.
type OrganizationUser struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// OrganizationUserListResponse represents the response from listing organizationUsers.
type OrganizationUserListResponse struct {
	Usable     []OrganizationUser `json:"usable"`
	Manageable []OrganizationUser `json:"manageable"`
}

// List returns all organizationUsers.
func (s *OrganizationUserService) List() (*OrganizationUserListResponse, error) {
	resp, err := s.client.get("/organization/{orgID}/user")
	if err != nil {
		return nil, fmt.Errorf("sc: list organizationUsers: %w", err)
	}

	var result OrganizationUserListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organizationUser list response: %w", err)
	}

	return &result, nil
}

// Get returns the organizationUser with the given ID.
func (s *OrganizationUserService) Get(id string) (*OrganizationUser, error) {
	resp, err := s.client.get("/organization/{orgID}/user" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get organizationUser %s: %w", id, err)
	}

	var result OrganizationUser
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organizationUser response: %w", err)
	}

	return &result, nil
}

