package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// AssetTemplateService handles assetTemplate operations.
type AssetTemplateService struct {
	client *Client
}

// AssetTemplate represents a assetTemplate resource.
type AssetTemplate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AssetTemplateListResponse represents the response from listing assetTemplates.
type AssetTemplateListResponse struct {
	Usable     []AssetTemplate `json:"usable"`
	Manageable []AssetTemplate `json:"manageable"`
}

// List returns all assetTemplates.
func (s *AssetTemplateService) List(ctx context.Context) (*AssetTemplateListResponse, error) {
	resp, err := s.client.get(ctx, "/assetTemplate")
	if err != nil {
		return nil, fmt.Errorf("sc: list assetTemplates: %w", err)
	}

	var result AssetTemplateListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal assetTemplate list response: %w", err)
	}

	return &result, nil
}

// Get returns the assetTemplate with the given ID.
func (s *AssetTemplateService) Get(ctx context.Context, id string) (*AssetTemplate, error) {
	resp, err := s.client.get(ctx, "/assetTemplate"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get assetTemplate %s: %w", id, err)
	}

	var result AssetTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal assetTemplate response: %w", err)
	}

	return &result, nil
}

// Categories performs the categories action on the assetTemplate.
func (s *AssetTemplateService) Categories(ctx context.Context) (*AssetTemplate, error) {
	resp, err := s.client.get(ctx, "/assetTemplate/categories")
	if err != nil {
		return nil, fmt.Errorf("sc: categories assetTemplate: %w", err)
	}

	var result AssetTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal assetTemplate categories response: %w", err)
	}

	return &result, nil
}
