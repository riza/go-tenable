package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// RecastRiskRuleService handles recastRiskRule operations.
type RecastRiskRuleService struct {
	client *Client
}

// RecastRiskRule represents a recastRiskRule resource.
type RecastRiskRule struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// RecastRiskRuleListResponse represents the response from listing recastRiskRules.
type RecastRiskRuleListResponse struct {
	Usable     []RecastRiskRule `json:"usable"`
	Manageable []RecastRiskRule `json:"manageable"`
}

// RecastRiskRuleCreateInput represents the request body for creating a recastRiskRule.
type RecastRiskRuleCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// RecastRiskRuleUpdateInput represents the request body for updating a recastRiskRule.
type RecastRiskRuleUpdateInput = RecastRiskRuleCreateInput

// List returns all recastRiskRules.
func (s *RecastRiskRuleService) List(ctx context.Context) (*RecastRiskRuleListResponse, error) {
	resp, err := s.client.get(ctx, "/recastRiskRule")
	if err != nil {
		return nil, fmt.Errorf("sc: list recastRiskRules: %w", err)
	}

	var result RecastRiskRuleListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal recastRiskRule list response: %w", err)
	}

	return &result, nil
}

// Create creates a new recastRiskRule.
func (s *RecastRiskRuleService) Create(ctx context.Context, input *RecastRiskRuleCreateInput) (*RecastRiskRule, error) {
	resp, err := s.client.post(ctx, "/recastRiskRule", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create recastRiskRule: %w", err)
	}

	var result RecastRiskRule
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal recastRiskRule response: %w", err)
	}

	return &result, nil
}

// Get returns the recastRiskRule with the given ID.
func (s *RecastRiskRuleService) Get(ctx context.Context, id string) (*RecastRiskRule, error) {
	resp, err := s.client.get(ctx, "/recastRiskRule"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get recastRiskRule %s: %w", id, err)
	}

	var result RecastRiskRule
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal recastRiskRule response: %w", err)
	}

	return &result, nil
}

// Delete deletes the recastRiskRule with the given ID.
func (s *RecastRiskRuleService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/recastRiskRule"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete recastRiskRule %s: %w", id, err)
	}

	return nil
}

// Update updates the recastRiskRule with the given ID.
func (s *RecastRiskRuleService) Update(ctx context.Context, id string, input *RecastRiskRuleUpdateInput) (*RecastRiskRule, error) {
	resp, err := s.client.patch(ctx, "/recastRiskRule"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update recastRiskRule %s: %w", id, err)
	}

	var result RecastRiskRule
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal recastRiskRule response: %w", err)
	}

	return &result, nil
}

// Apply performs the apply action on the recastRiskRule.
func (s *RecastRiskRuleService) Apply(ctx context.Context) (*RecastRiskRule, error) {
	resp, err := s.client.post(ctx, "/recastRiskRule/apply", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: apply recastRiskRule: %w", err)
	}

	var result RecastRiskRule
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal recastRiskRule apply response: %w", err)
	}

	return &result, nil
}
