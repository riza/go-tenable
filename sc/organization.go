package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// OrganizationService handles organization operations.
type OrganizationService struct {
	client *Client
}

// Organization represents a organization resource.
type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// OrganizationListResponse represents the response from listing organizations.
type OrganizationListResponse struct {
	Usable     []Organization `json:"usable"`
	Manageable []Organization `json:"manageable"`
}

// OrganizationCreateInput represents the request body for creating a organization.
type OrganizationCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// OrganizationUpdateInput represents the request body for updating a organization.
type OrganizationUpdateInput = OrganizationCreateInput

// List returns all organizations.
func (s *OrganizationService) List(ctx context.Context) (*OrganizationListResponse, error) {
	resp, err := s.client.get(ctx, "/organization")
	if err != nil {
		return nil, fmt.Errorf("sc: list organizations: %w", err)
	}

	var result OrganizationListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organization list response: %w", err)
	}

	return &result, nil
}

// Create creates a new organization.
func (s *OrganizationService) Create(ctx context.Context, input *OrganizationCreateInput) (*Organization, error) {
	resp, err := s.client.post(ctx, "/organization", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create organization: %w", err)
	}

	var result Organization
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organization response: %w", err)
	}

	return &result, nil
}

// Get returns the organization with the given ID.
func (s *OrganizationService) Get(ctx context.Context, id string) (*Organization, error) {
	resp, err := s.client.get(ctx, "/organization"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get organization %s: %w", id, err)
	}

	var result Organization
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organization response: %w", err)
	}

	return &result, nil
}

// Update updates the organization with the given ID.
func (s *OrganizationService) Update(ctx context.Context, id string, input *OrganizationUpdateInput) (*Organization, error) {
	resp, err := s.client.patch(ctx, "/organization"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update organization %s: %w", id, err)
	}

	var result Organization
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organization response: %w", err)
	}

	return &result, nil
}

// Delete deletes the organization with the given ID.
func (s *OrganizationService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/organization"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete organization %s: %w", id, err)
	}

	return nil
}

// AcceptRiskRule performs the acceptRiskRule action on the organization with the given ID.
func (s *OrganizationService) AcceptRiskRule(ctx context.Context, id string) (*Organization, error) {
	resp, err := s.client.get(ctx, "/organization"+"/"+id+"/acceptRiskRule")
	if err != nil {
		return nil, fmt.Errorf("sc: acceptRiskRule organization %s: %w", id, err)
	}

	var result Organization
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organization acceptRiskRule response: %w", err)
	}

	return &result, nil
}

// RecastRiskRule performs the recastRiskRule action on the organization with the given ID.
func (s *OrganizationService) RecastRiskRule(ctx context.Context, id string) (*Organization, error) {
	resp, err := s.client.get(ctx, "/organization"+"/"+id+"/recastRiskRule")
	if err != nil {
		return nil, fmt.Errorf("sc: recastRiskRule organization %s: %w", id, err)
	}

	var result Organization
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal organization recastRiskRule response: %w", err)
	}

	return &result, nil
}
