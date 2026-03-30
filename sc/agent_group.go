package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// AgentGroupService handles agentGroup operations.
type AgentGroupService struct {
	client *Client
}

// AgentGroup represents a agentGroup resource.
type AgentGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Get returns the agentGroup with the given ID.
func (s *AgentGroupService) Get(ctx context.Context, id string) (*AgentGroup, error) {
	resp, err := s.client.get(ctx, "/agentGroup"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get agentGroup %s: %w", id, err)
	}

	var result AgentGroup
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentGroup response: %w", err)
	}

	return &result, nil
}

// Remote performs the remote action on the agentGroup with the given ID.
func (s *AgentGroupService) Remote(ctx context.Context, id string) (*AgentGroup, error) {
	resp, err := s.client.get(ctx, "/agentGroup"+"/"+id+"/remote")
	if err != nil {
		return nil, fmt.Errorf("sc: remote agentGroup %s: %w", id, err)
	}

	var result AgentGroup
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentGroup remote response: %w", err)
	}

	return &result, nil
}
