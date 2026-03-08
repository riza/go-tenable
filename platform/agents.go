package platform

import (
	"context"
	"encoding/json"
	"fmt"
)

// AgentsService handles communication with the Agents related endpoints of the Platform API.
type AgentsService struct {
	client *Client
}

// Agent represents an agent.
type Agent struct {
	Id           string            `json:"id,omitempty"`
	Uuid         string            `json:"uuid,omitempty"`
	Name         string            `json:"name,omitempty"`
	IpAddress    string            `json:"ip_address,omitempty"`
	Platform     string            `json:"platform,omitempty"`
	Status       string            `json:"status,omitempty"`
	GroupId      string            `json:"group_id,omitempty"`
	ProfileId    string            `json:"profile_id,omitempty"`
	NetworkId    string            `json:"network_id,omitempty"`
	LinkedOn     int               `json:"linked_on,omitempty"`
	LastSeen     string            `json:"last_seen,omitempty"`
	Tags         map[string]string `json:"tags,omitempty"`
	PluginPolicy string            `json:"plugin_policy,omitempty"`
}

// AgentGroup represents an agent group.
type AgentGroup struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_time,omitempty"`
	ModifiedAt  string `json:"modified_time,omitempty"`
}

// AgentExclusion represents an agent exclusion.
type AgentExclusion struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	StartTime   int      `json:"start_time,omitempty"`
	EndTime     int      `json:"end_time,omitempty"`
	Timezone    string   `json:"timezone,omitempty"`
	Acums       []string `json:"acums,omitempty"`
}

// AgentsListResponse represents the response from listing agents.
type AgentsListResponse struct {
	Agents []Agent `json:"agents,omitempty"`
	Total  int     `json:"total,omitempty"`
}

// ListAgents returns all agents.
func (s *AgentsService) ListAgents(ctx context.Context) (*AgentsListResponse, error) {
	resp, err := s.client.get(ctx, "/scanners/null/agents")
	if err != nil {
		return nil, err
	}

	var result AgentsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAgent returns a specific agent by ID.
func (s *AgentsService) GetAgent(ctx context.Context, agentId string) (*Agent, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scanners/null/agents/%s", agentId))
	if err != nil {
		return nil, err
	}

	var result Agent
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateAgent updates an agent.
func (s *AgentsService) UpdateAgent(ctx context.Context, agentId string, req *Agent) (*Agent, error) {
	resp, err := s.client.patch(ctx, fmt.Sprintf("/scanners/null/agents/%s", agentId), req)
	if err != nil {
		return nil, err
	}

	var result Agent
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteAgent deletes an agent.
func (s *AgentsService) DeleteAgent(ctx context.Context, agentId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/scanners/null/agents/%s", agentId))
	return err
}

// GetAgentConfig returns the agent configuration.
func (s *AgentsService) GetAgentConfig(ctx context.Context) (map[string]interface{}, error) {
	resp, err := s.client.get(ctx, "/scanners/null/agents/config")
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateAgentConfig updates the agent configuration.
func (s *AgentsService) UpdateAgentConfig(ctx context.Context, req map[string]interface{}) (map[string]interface{}, error) {
	resp, err := s.client.put(ctx, "/scanners/null/agents/config", req)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListAgentGroups returns all agent groups.
func (s *AgentsService) ListAgentGroups(ctx context.Context) ([]AgentGroup, error) {
	resp, err := s.client.get(ctx, "/scanners/null/agent-groups")
	if err != nil {
		return nil, err
	}

	var result []AgentGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateAgentGroup creates a new agent group.
func (s *AgentsService) CreateAgentGroup(ctx context.Context, req *AgentGroup) (*AgentGroup, error) {
	resp, err := s.client.post(ctx, "/scanners/null/agent-groups", req)
	if err != nil {
		return nil, err
	}

	var result AgentGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAgentGroup returns a specific agent group by ID.
func (s *AgentsService) GetAgentGroup(ctx context.Context, groupId string) (*AgentGroup, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scanners/null/agent-groups/%s", groupId))
	if err != nil {
		return nil, err
	}

	var result AgentGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateAgentGroup updates an agent group.
func (s *AgentsService) UpdateAgentGroup(ctx context.Context, groupId string, req *AgentGroup) (*AgentGroup, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/scanners/null/agent-groups/%s", groupId), req)
	if err != nil {
		return nil, err
	}

	var result AgentGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteAgentGroup deletes an agent group.
func (s *AgentsService) DeleteAgentGroup(ctx context.Context, groupId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/scanners/null/agent-groups/%s", groupId))
	return err
}

// ListAgentsInGroup returns all agents in a specific group.
func (s *AgentsService) ListAgentsInGroup(ctx context.Context, groupId string) (*AgentsListResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scanners/null/agent-groups/%s/agents", groupId))
	if err != nil {
		return nil, err
	}

	var result AgentsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// AssignAgentToGroup assigns an agent to a group.
func (s *AgentsService) AssignAgentToGroup(ctx context.Context, groupId, agentId string) error {
	_, err := s.client.put(ctx, fmt.Sprintf("/scanners/null/agent-groups/%s/agents/%s", groupId, agentId), nil)
	return err
}

// RemoveAgentFromGroup removes an agent from a group.
func (s *AgentsService) RemoveAgentFromGroup(ctx context.Context, groupId, agentId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/scanners/null/agent-groups/%s/agents/%s", groupId, agentId))
	return err
}

// ListAgentExclusions returns all agent exclusions.
func (s *AgentsService) ListAgentExclusions(ctx context.Context) ([]AgentExclusion, error) {
	resp, err := s.client.get(ctx, "/scanners/null/agents/exclusions")
	if err != nil {
		return nil, err
	}

	var result []AgentExclusion
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateAgentExclusion creates a new agent exclusion.
func (s *AgentsService) CreateAgentExclusion(ctx context.Context, req *AgentExclusion) (*AgentExclusion, error) {
	resp, err := s.client.post(ctx, "/scanners/null/agents/exclusions", req)
	if err != nil {
		return nil, err
	}

	var result AgentExclusion
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAgentExclusion returns a specific agent exclusion by ID.
func (s *AgentsService) GetAgentExclusion(ctx context.Context, exclusionId string) (*AgentExclusion, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scanners/null/agents/exclusions/%s", exclusionId))
	if err != nil {
		return nil, err
	}

	var result AgentExclusion
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateAgentExclusion updates an agent exclusion.
func (s *AgentsService) UpdateAgentExclusion(ctx context.Context, exclusionId string, req *AgentExclusion) (*AgentExclusion, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/scanners/null/agents/exclusions/%s", exclusionId), req)
	if err != nil {
		return nil, err
	}

	var result AgentExclusion
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteAgentExclusion deletes an agent exclusion.
func (s *AgentsService) DeleteAgentExclusion(ctx context.Context, exclusionId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/scanners/null/agents/exclusions/%s", exclusionId))
	return err
}

// GetSafeModeSummary returns the safe mode summary.
func (s *AgentsService) GetSafeModeSummary(ctx context.Context) (map[string]interface{}, error) {
	resp, err := s.client.get(ctx, "/agents/safe-mode-summary")
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}
