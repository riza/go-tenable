package platform

import (
	"context"
	"encoding/json"
)

// ServerService handles communication with the Server related endpoints of the Platform API.
type ServerService struct {
	client *Client
}

// ServerInfo represents server information.
type ServerInfo struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Version  string `json:"version,omitempty"`
	Build    string `json:"build,omitempty"`
	Edition  string `json:"edition,omitempty"`
	License  string `json:"license,omitempty"`
	Platform string `json:"platform,omitempty"`
}

// GetServerInfo returns server information.
func (s *ServerService) GetServerInfo(ctx context.Context) (*ServerInfo, error) {
	resp, err := s.client.get(ctx, "/server")
	if err != nil {
		return nil, err
	}

	var result ServerInfo
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
