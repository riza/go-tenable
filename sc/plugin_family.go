package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// PluginFamilyService handles pluginFamily operations.
type PluginFamilyService struct {
	client *Client
}

// PluginFamily represents a pluginFamily resource.
type PluginFamily struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PluginFamilyListResponse represents the response from listing pluginFamilys.
type PluginFamilyListResponse struct {
	Usable     []PluginFamily `json:"usable"`
	Manageable []PluginFamily `json:"manageable"`
}

// List returns all pluginFamilys.
func (s *PluginFamilyService) List(ctx context.Context) (*PluginFamilyListResponse, error) {
	resp, err := s.client.get(ctx, "/pluginFamily")
	if err != nil {
		return nil, fmt.Errorf("sc: list pluginFamilys: %w", err)
	}

	var result PluginFamilyListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal pluginFamily list response: %w", err)
	}

	return &result, nil
}
