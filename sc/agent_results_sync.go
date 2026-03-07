
package sc

import (
	"encoding/json"
	"fmt"
)

// AgentResultsSyncService handles agentResultsSync operations.
type AgentResultsSyncService struct {
	client *Client
}

// AgentResultsSync represents a agentResultsSync resource.
type AgentResultsSync struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AgentResultsSyncListResponse represents the response from listing agentResultsSyncs.
type AgentResultsSyncListResponse struct {
	Usable     []AgentResultsSync `json:"usable"`
	Manageable []AgentResultsSync `json:"manageable"`
}

// AgentResultsSyncCreateInput represents the request body for creating a agentResultsSync.
type AgentResultsSyncCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// AgentResultsSyncUpdateInput represents the request body for updating a agentResultsSync.
type AgentResultsSyncUpdateInput = AgentResultsSyncCreateInput

// List returns all agentResultsSyncs.
func (s *AgentResultsSyncService) List() (*AgentResultsSyncListResponse, error) {
	resp, err := s.client.get("/agentResultsSync")
	if err != nil {
		return nil, fmt.Errorf("sc: list agentResultsSyncs: %w", err)
	}

	var result AgentResultsSyncListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentResultsSync list response: %w", err)
	}

	return &result, nil
}

// Create creates a new agentResultsSync.
func (s *AgentResultsSyncService) Create(input *AgentResultsSyncCreateInput) (*AgentResultsSync, error) {
	resp, err := s.client.post("/agentResultsSync", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create agentResultsSync: %w", err)
	}

	var result AgentResultsSync
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentResultsSync response: %w", err)
	}

	return &result, nil
}

// Get returns the agentResultsSync with the given ID.
func (s *AgentResultsSyncService) Get(id string) (*AgentResultsSync, error) {
	resp, err := s.client.get("/agentResultsSync" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get agentResultsSync %s: %w", id, err)
	}

	var result AgentResultsSync
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentResultsSync response: %w", err)
	}

	return &result, nil
}

// Update updates the agentResultsSync with the given ID.
func (s *AgentResultsSyncService) Update(id string, input *AgentResultsSyncUpdateInput) (*AgentResultsSync, error) {
	resp, err := s.client.patch("/agentResultsSync" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update agentResultsSync %s: %w", id, err)
	}

	var result AgentResultsSync
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentResultsSync response: %w", err)
	}

	return &result, nil
}

// Delete deletes the agentResultsSync with the given ID.
func (s *AgentResultsSyncService) Delete(id string) error {
	_, err := s.client.delete("/agentResultsSync" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete agentResultsSync %s: %w", id, err)
	}

	return nil
}

// Launch performs the launch action on the agentResultsSync with the given ID.
func (s *AgentResultsSyncService) Launch(id string) (*AgentResultsSync, error) {
	resp, err := s.client.post("/agentResultsSync" + "/" + id + "/launch", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: launch agentResultsSync %s: %w", id, err)
	}

	var result AgentResultsSync
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentResultsSync launch response: %w", err)
	}

	return &result, nil
}

