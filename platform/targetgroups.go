package platform

import (
	"context"
	"encoding/json"
)

// TargetGroupsService handles communication with the Target Groups related endpoints of the Platform API.
type TargetGroupsService struct {
	client *Client
}

// TargetGroup represents a target group.
type TargetGroup struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type,omitempty"`
	Targets     []string `json:"targets,omitempty"`
	CreatedAt   string   `json:"created_time,omitempty"`
	UpdatedAt   string   `json:"updated_time,omitempty"`
}

// TargetGroupsListResponse represents the response from listing target groups.
type TargetGroupsListResponse struct {
	TargetGroups []TargetGroup `json:"target_groups,omitempty"`
	Total        int           `json:"total,omitempty"`
}

// ListTargetGroups returns all target groups.
func (s *TargetGroupsService) ListTargetGroups(ctx context.Context) (*TargetGroupsListResponse, error) {
	resp, err := s.client.get(ctx, "/target-groups")
	if err != nil {
		return nil, err
	}

	var result TargetGroupsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateTargetGroup creates a new target group.
func (s *TargetGroupsService) CreateTargetGroup(ctx context.Context, req *TargetGroup) (*TargetGroup, error) {
	resp, err := s.client.post(ctx, "/target-groups", req)
	if err != nil {
		return nil, err
	}

	var result TargetGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTargetGroup returns a specific target group by ID.
func (s *TargetGroupsService) GetTargetGroup(ctx context.Context, targetGroupId string) (*TargetGroup, error) {
	resp, err := s.client.get(ctx, "/target-groups/"+targetGroupId)
	if err != nil {
		return nil, err
	}

	var result TargetGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateTargetGroup updates a target group.
func (s *TargetGroupsService) UpdateTargetGroup(ctx context.Context, targetGroupId string, req *TargetGroup) (*TargetGroup, error) {
	resp, err := s.client.put(ctx, "/target-groups/"+targetGroupId, req)
	if err != nil {
		return nil, err
	}

	var result TargetGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteTargetGroup deletes a target group.
func (s *TargetGroupsService) DeleteTargetGroup(ctx context.Context, targetGroupId string) error {
	_, err := s.client.delete(ctx, "/target-groups/"+targetGroupId)
	return err
}
