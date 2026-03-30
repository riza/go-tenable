package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// LCEPolicyService handles lCEPolicy operations.
type LCEPolicyService struct {
	client *Client
}

// LCEPolicy represents a lCEPolicy resource.
type LCEPolicy struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// LCEPolicyListResponse represents the response from listing lCEPolicys.
type LCEPolicyListResponse struct {
	Usable     []LCEPolicy `json:"usable"`
	Manageable []LCEPolicy `json:"manageable"`
}

// List returns all lCEPolicys.
func (s *LCEPolicyService) List(ctx context.Context) (*LCEPolicyListResponse, error) {
	resp, err := s.client.get(ctx, "/lce/{id}/policy")
	if err != nil {
		return nil, fmt.Errorf("sc: list lCEPolicys: %w", err)
	}

	var result LCEPolicyListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lCEPolicy list response: %w", err)
	}

	return &result, nil
}

// Delete deletes the lCEPolicy.
func (s *LCEPolicyService) Delete(ctx context.Context) error {
	_, err := s.client.delete(ctx, "/lce/{id}/policy")
	if err != nil {
		return fmt.Errorf("sc: delete lCEPolicy: %w", err)
	}

	return nil
}
