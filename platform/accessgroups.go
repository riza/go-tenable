package platform

import (
	"context"
	"encoding/json"
	"fmt"
)

// AccessGroupsService handles communication with the Access Groups related endpoints of the Platform API.
type AccessGroupsService struct {
	client *Client
}

// AccessGroup represents an access group.
type AccessGroup struct {
	Id          string            `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	CreatedAt   string            `json:"created_time,omitempty"`
	UpdatedAt   string            `json:"modified_time,omitempty"`
	Rules       []AccessGroupRule `json:"rules,omitempty"`
}

// AccessGroupRule represents an access group rule.
type AccessGroupRule struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
}

// AccessGroupsListResponse represents the response from listing access groups.
type AccessGroupsListResponse struct {
	AccessGroups []AccessGroup `json:"access_groups,omitempty"`
	Total        int           `json:"total,omitempty"`
}

// ListAccessGroups returns all access groups (v1).
func (s *AccessGroupsService) ListAccessGroups(ctx context.Context) (*AccessGroupsListResponse, error) {
	resp, err := s.client.get(ctx, "/access-groups")
	if err != nil {
		return nil, err
	}

	var result AccessGroupsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateAccessGroup creates a new access group (v1).
func (s *AccessGroupsService) CreateAccessGroup(ctx context.Context, req *AccessGroup) (*AccessGroup, error) {
	resp, err := s.client.post(ctx, "/access-groups", req)
	if err != nil {
		return nil, err
	}

	var result AccessGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAccessGroup returns a specific access group by ID (v1).
func (s *AccessGroupsService) GetAccessGroup(ctx context.Context, id string) (*AccessGroup, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/access-groups/%s", id))
	if err != nil {
		return nil, err
	}

	var result AccessGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateAccessGroup updates an access group (v1).
func (s *AccessGroupsService) UpdateAccessGroup(ctx context.Context, id string, req *AccessGroup) (*AccessGroup, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/access-groups/%s", id), req)
	if err != nil {
		return nil, err
	}

	var result AccessGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteAccessGroup deletes an access group (v1).
func (s *AccessGroupsService) DeleteAccessGroup(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/access-groups/%s", id))
	return err
}

// AccessGroupsFiltersResponse represents the response from getting access group filters.
type AccessGroupsFiltersResponse struct {
	Filters []map[string]interface{} `json:"filters,omitempty"`
}

// ListAccessGroupFilters returns available filters for access groups (v1).
func (s *AccessGroupsService) ListAccessGroupFilters(ctx context.Context) (*AccessGroupsFiltersResponse, error) {
	resp, err := s.client.get(ctx, "/access-groups/filters")
	if err != nil {
		return nil, err
	}

	var result AccessGroupsFiltersResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ListAccessGroupRuleFilters returns available filters for access group rules (v1).
func (s *AccessGroupsService) ListAccessGroupRuleFilters(ctx context.Context) (*AccessGroupsFiltersResponse, error) {
	resp, err := s.client.get(ctx, "/access-groups/rules/filters")
	if err != nil {
		return nil, err
	}

	var result AccessGroupsFiltersResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ListAccessGroupsV2 returns all access groups (v2).
func (s *AccessGroupsService) ListAccessGroupsV2(ctx context.Context) (*AccessGroupsListResponse, error) {
	resp, err := s.client.get(ctx, "/v2/access-groups")
	if err != nil {
		return nil, err
	}

	var result AccessGroupsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateAccessGroupV2 creates a new access group (v2).
func (s *AccessGroupsService) CreateAccessGroupV2(ctx context.Context, req *AccessGroup) (*AccessGroup, error) {
	resp, err := s.client.post(ctx, "/v2/access-groups", req)
	if err != nil {
		return nil, err
	}

	var result AccessGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAccessGroupV2 returns a specific access group by ID (v2).
func (s *AccessGroupsService) GetAccessGroupV2(ctx context.Context, id string) (*AccessGroup, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/v2/access-groups/%s", id))
	if err != nil {
		return nil, err
	}

	var result AccessGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateAccessGroupV2 updates an access group (v2).
func (s *AccessGroupsService) UpdateAccessGroupV2(ctx context.Context, id string, req *AccessGroup) (*AccessGroup, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/v2/access-groups/%s", id), req)
	if err != nil {
		return nil, err
	}

	var result AccessGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteAccessGroupV2 deletes an access group (v2).
func (s *AccessGroupsService) DeleteAccessGroupV2(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/v2/access-groups/%s", id))
	return err
}

// ListAccessGroupFiltersV2 returns available filters for access groups (v2).
func (s *AccessGroupsService) ListAccessGroupFiltersV2(ctx context.Context) (*AccessGroupsFiltersResponse, error) {
	resp, err := s.client.get(ctx, "/v2/access-groups/filters")
	if err != nil {
		return nil, err
	}

	var result AccessGroupsFiltersResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ListAccessGroupRuleFiltersV2 returns available filters for access group rules (v2).
func (s *AccessGroupsService) ListAccessGroupRuleFiltersV2(ctx context.Context) (*AccessGroupsFiltersResponse, error) {
	resp, err := s.client.get(ctx, "/v2/access-groups/rules/filters")
	if err != nil {
		return nil, err
	}

	var result AccessGroupsFiltersResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
