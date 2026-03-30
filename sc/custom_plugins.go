package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// CustomPluginsService handles customPlugins operations.
type CustomPluginsService struct {
	client *Client
}

// CustomPlugins represents a customPlugins resource.
type CustomPlugins struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CustomPluginsListResponse represents the response from listing customPluginss.
type CustomPluginsListResponse struct {
	Usable     []CustomPlugins `json:"usable"`
	Manageable []CustomPlugins `json:"manageable"`
}

// List returns all customPluginss.
func (s *CustomPluginsService) List(ctx context.Context) (*CustomPluginsListResponse, error) {
	resp, err := s.client.get(ctx, "/customPlugins")
	if err != nil {
		return nil, fmt.Errorf("sc: list customPluginss: %w", err)
	}

	var result CustomPluginsListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal customPlugins list response: %w", err)
	}

	return &result, nil
}

// Process performs the process action on the customPlugins with the given ID.
func (s *CustomPluginsService) Process(ctx context.Context, id string) (*CustomPlugins, error) {
	resp, err := s.client.post(ctx, "/customPlugins"+"/"+id+"/process", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: process customPlugins %s: %w", id, err)
	}

	var result CustomPlugins
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal customPlugins process response: %w", err)
	}

	return &result, nil
}
