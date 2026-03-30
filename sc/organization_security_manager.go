package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// OrganizationSecurityManagerService handles organizationSecurityManager operations.
type OrganizationSecurityManagerService struct {
	client *Client
}

// OrganizationSecurityManager represents a organizationSecurityManager resource.
type OrganizationSecurityManager struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// OrganizationSecurityManagerListResponse represents the response from listing organizationSecurityManagers.
type OrganizationSecurityManagerListResponse struct {
	Usable     []OrganizationSecurityManager `json:"usable"`
	Manageable []OrganizationSecurityManager `json:"manageable"`
}

// OrganizationSecurityManagerCreateInput represents the request body for creating a organizationSecurityManager.
type OrganizationSecurityManagerCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// OrganizationSecurityManagerUpdateInput represents the request body for updating a organizationSecurityManager.
type OrganizationSecurityManagerUpdateInput = OrganizationSecurityManagerCreateInput

// List returns all organizationSecurityManagers.
func (s *OrganizationSecurityManagerService) List(ctx context.Context) (*OrganizationSecurityManagerListResponse, error) {
	resp, err := s.client.get(ctx, "/organization/{orgID}/securityManager")
	if err != nil {
		return nil, fmt.Errorf("sc: list organizationSecurityManagers: %w", err)
	}

	var result OrganizationSecurityManagerListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organizationSecurityManager list response: %w", err)
	}

	return &result, nil
}

// Create creates a new organizationSecurityManager.
func (s *OrganizationSecurityManagerService) Create(ctx context.Context, input *OrganizationSecurityManagerCreateInput) (*OrganizationSecurityManager, error) {
	resp, err := s.client.post(ctx, "/organization/{orgID}/securityManager", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create organizationSecurityManager: %w", err)
	}

	var result OrganizationSecurityManager
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organizationSecurityManager response: %w", err)
	}

	return &result, nil
}

// Get returns the organizationSecurityManager with the given ID.
func (s *OrganizationSecurityManagerService) Get(ctx context.Context, id string) (*OrganizationSecurityManager, error) {
	resp, err := s.client.get(ctx, "/organization/{orgID}/securityManager"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get organizationSecurityManager %s: %w", id, err)
	}

	var result OrganizationSecurityManager
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organizationSecurityManager response: %w", err)
	}

	return &result, nil
}

// Update updates the organizationSecurityManager with the given ID.
func (s *OrganizationSecurityManagerService) Update(ctx context.Context, id string, input *OrganizationSecurityManagerUpdateInput) (*OrganizationSecurityManager, error) {
	resp, err := s.client.patch(ctx, "/organization/{orgID}/securityManager"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update organizationSecurityManager %s: %w", id, err)
	}

	var result OrganizationSecurityManager
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organizationSecurityManager response: %w", err)
	}

	return &result, nil
}

// Delete deletes the organizationSecurityManager with the given ID.
func (s *OrganizationSecurityManagerService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/organization/{orgID}/securityManager"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete organizationSecurityManager %s: %w", id, err)
	}

	return nil
}
