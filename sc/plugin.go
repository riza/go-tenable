
package sc

import (
	"encoding/json"
	"fmt"
)

// PluginService handles plugin operations.
type PluginService struct {
	client *Client
}

// Plugin represents a plugin resource.
type Plugin struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PluginListResponse represents the response from listing plugins.
type PluginListResponse struct {
	Usable     []Plugin `json:"usable"`
	Manageable []Plugin `json:"manageable"`
}

// List returns all plugins.
func (s *PluginService) List() (*PluginListResponse, error) {
	resp, err := s.client.get("/plugin")
	if err != nil {
		return nil, fmt.Errorf("sc: list plugins: %w", err)
	}

	var result PluginListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal plugin list response: %w", err)
	}

	return &result, nil
}

// Get returns the plugin with the given ID.
func (s *PluginService) Get(id string) (*Plugin, error) {
	resp, err := s.client.get("/plugin" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get plugin %s: %w", id, err)
	}

	var result Plugin
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal plugin response: %w", err)
	}

	return &result, nil
}

