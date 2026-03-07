
package sc

import (
	"encoding/json"
	"fmt"
)

// AgentScanService handles agentScan operations.
type AgentScanService struct {
	client *Client
}

// AgentScan represents a agentScan resource.
type AgentScan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AgentScanListResponse represents the response from listing agentScans.
type AgentScanListResponse struct {
	Usable     []AgentScan `json:"usable"`
	Manageable []AgentScan `json:"manageable"`
}

// AgentScanCreateInput represents the request body for creating a agentScan.
type AgentScanCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// AgentScanUpdateInput represents the request body for updating a agentScan.
type AgentScanUpdateInput = AgentScanCreateInput

// List returns all agentScans.
func (s *AgentScanService) List() (*AgentScanListResponse, error) {
	resp, err := s.client.get("/agentScan")
	if err != nil {
		return nil, fmt.Errorf("sc: list agentScans: %w", err)
	}

	var result AgentScanListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentScan list response: %w", err)
	}

	return &result, nil
}

// Create creates a new agentScan.
func (s *AgentScanService) Create(input *AgentScanCreateInput) (*AgentScan, error) {
	resp, err := s.client.post("/agentScan", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create agentScan: %w", err)
	}

	var result AgentScan
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentScan response: %w", err)
	}

	return &result, nil
}

// Get returns the agentScan with the given ID.
func (s *AgentScanService) Get(id string) (*AgentScan, error) {
	resp, err := s.client.get("/agentScan" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get agentScan %s: %w", id, err)
	}

	var result AgentScan
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentScan response: %w", err)
	}

	return &result, nil
}

// Update updates the agentScan with the given ID.
func (s *AgentScanService) Update(id string, input *AgentScanUpdateInput) (*AgentScan, error) {
	resp, err := s.client.patch("/agentScan" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update agentScan %s: %w", id, err)
	}

	var result AgentScan
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentScan response: %w", err)
	}

	return &result, nil
}

// Delete deletes the agentScan with the given ID.
func (s *AgentScanService) Delete(id string) error {
	_, err := s.client.delete("/agentScan" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete agentScan %s: %w", id, err)
	}

	return nil
}

// Launch performs the launch action on the agentScan with the given ID.
func (s *AgentScanService) Launch(id string) (*AgentScan, error) {
	resp, err := s.client.post("/agentScan" + "/" + id + "/launch", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: launch agentScan %s: %w", id, err)
	}

	var result AgentScan
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal agentScan launch response: %w", err)
	}

	return &result, nil
}

