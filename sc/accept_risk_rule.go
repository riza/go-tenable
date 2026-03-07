
package sc

import (
	"encoding/json"
	"fmt"
)

// AcceptRiskRuleService handles acceptRiskRule operations.
type AcceptRiskRuleService struct {
	client *Client
}

// AcceptRiskRule represents a acceptRiskRule resource.
type AcceptRiskRule struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AcceptRiskRuleListResponse represents the response from listing acceptRiskRules.
type AcceptRiskRuleListResponse struct {
	Usable     []AcceptRiskRule `json:"usable"`
	Manageable []AcceptRiskRule `json:"manageable"`
}

// AcceptRiskRuleCreateInput represents the request body for creating a acceptRiskRule.
type AcceptRiskRuleCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// List returns all acceptRiskRules.
func (s *AcceptRiskRuleService) List() (*AcceptRiskRuleListResponse, error) {
	resp, err := s.client.get("/acceptRiskRule")
	if err != nil {
		return nil, fmt.Errorf("sc: list acceptRiskRules: %w", err)
	}

	var result AcceptRiskRuleListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal acceptRiskRule list response: %w", err)
	}

	return &result, nil
}

// Create creates a new acceptRiskRule.
func (s *AcceptRiskRuleService) Create(input *AcceptRiskRuleCreateInput) (*AcceptRiskRule, error) {
	resp, err := s.client.post("/acceptRiskRule", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create acceptRiskRule: %w", err)
	}

	var result AcceptRiskRule
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal acceptRiskRule response: %w", err)
	}

	return &result, nil
}

// Get returns the acceptRiskRule with the given ID.
func (s *AcceptRiskRuleService) Get(id string) (*AcceptRiskRule, error) {
	resp, err := s.client.get("/acceptRiskRule" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get acceptRiskRule %s: %w", id, err)
	}

	var result AcceptRiskRule
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal acceptRiskRule response: %w", err)
	}

	return &result, nil
}

// Delete deletes the acceptRiskRule with the given ID.
func (s *AcceptRiskRuleService) Delete(id string) error {
	_, err := s.client.delete("/acceptRiskRule" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete acceptRiskRule %s: %w", id, err)
	}

	return nil
}

// Apply performs the apply action on the acceptRiskRule.
func (s *AcceptRiskRuleService) Apply() (*AcceptRiskRule, error) {
	resp, err := s.client.post("/acceptRiskRule/apply", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: apply acceptRiskRule: %w", err)
	}

	var result AcceptRiskRule
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal acceptRiskRule apply response: %w", err)
	}

	return &result, nil
}

