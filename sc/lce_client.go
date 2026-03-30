package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// LCEClientService handles lCEClient operations.
type LCEClientService struct {
	client *Client
}

// LCEClient represents a lCEClient resource.
type LCEClient struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// LCEClientListResponse represents the response from listing lCEClients.
type LCEClientListResponse struct {
	Usable     []LCEClient `json:"usable"`
	Manageable []LCEClient `json:"manageable"`
}

// List returns all lCEClients.
func (s *LCEClientService) List(ctx context.Context) (*LCEClientListResponse, error) {
	resp, err := s.client.get(ctx, "/lce/{id}/client")
	if err != nil {
		return nil, fmt.Errorf("sc: list lCEClients: %w", err)
	}

	var result LCEClientListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lCEClient list response: %w", err)
	}

	return &result, nil
}

// Types performs the types action on the lCEClient.
func (s *LCEClientService) Types(ctx context.Context) (*LCEClient, error) {
	resp, err := s.client.get(ctx, "/lce/{id}/client/types")
	if err != nil {
		return nil, fmt.Errorf("sc: types lCEClient: %w", err)
	}

	var result LCEClient
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lCEClient types response: %w", err)
	}

	return &result, nil
}

// OsTypes performs the osTypes action on the lCEClient.
func (s *LCEClientService) OsTypes(ctx context.Context) (*LCEClient, error) {
	resp, err := s.client.get(ctx, "/lce/{id}/client/osTypes")
	if err != nil {
		return nil, fmt.Errorf("sc: osTypes lCEClient: %w", err)
	}

	var result LCEClient
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lCEClient osTypes response: %w", err)
	}

	return &result, nil
}
