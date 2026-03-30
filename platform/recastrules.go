package platform

import (
	"context"
	"encoding/json"
)

// RecastRulesService handles communication with the Recast Rules related endpoints of the Platform API.
type RecastRulesService struct {
	client *Client
}

// RecastRule represents a recast rule.
type RecastRule struct {
	Id          string `json:"id,omitempty"`
	PluginId    int    `json:"plugin_id,omitempty"`
	Severity    string `json:"severity,omitempty"`
	RecastType  string `json:"recast_type,omitempty"`
	NewSeverity string `json:"new_severity,omitempty"`
	CreatedBy   string `json:"created_by,omitempty"`
	CreatedAt   string `json:"created_time,omitempty"`
}

// RecastRulesListResponse represents the response from listing recast rules.
type RecastRulesListResponse struct {
	Rules []RecastRule `json:"rules,omitempty"`
	Total int          `json:"total,omitempty"`
}

// ListRecastRules returns all recast rules.
func (s *RecastRulesService) ListRecastRules(ctx context.Context) (*RecastRulesListResponse, error) {
	resp, err := s.client.get(ctx, "/recast-rules")
	if err != nil {
		return nil, err
	}

	var result RecastRulesListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateRecastRule creates a new recast rule.
func (s *RecastRulesService) CreateRecastRule(ctx context.Context, req *RecastRule) (*RecastRule, error) {
	resp, err := s.client.post(ctx, "/recast-rules", req)
	if err != nil {
		return nil, err
	}

	var result RecastRule
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetRecastRule returns a specific recast rule by ID.
func (s *RecastRulesService) GetRecastRule(ctx context.Context, ruleId string) (*RecastRule, error) {
	resp, err := s.client.get(ctx, "/recast-rules/"+ruleId)
	if err != nil {
		return nil, err
	}

	var result RecastRule
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteRecastRule deletes a recast rule.
func (s *RecastRulesService) DeleteRecastRule(ctx context.Context, ruleId string) error {
	_, err := s.client.delete(ctx, "/recast-rules/"+ruleId)
	return err
}
