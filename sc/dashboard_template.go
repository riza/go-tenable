package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DashboardTemplateService handles dashboardTemplate operations.
type DashboardTemplateService struct {
	client *Client
}

// DashboardTemplate represents a dashboardTemplate resource.
type DashboardTemplate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DashboardTemplateListResponse represents the response from listing dashboardTemplates.
type DashboardTemplateListResponse struct {
	Usable     []DashboardTemplate `json:"usable"`
	Manageable []DashboardTemplate `json:"manageable"`
}

// List returns all dashboardTemplates.
func (s *DashboardTemplateService) List(ctx context.Context) (*DashboardTemplateListResponse, error) {
	resp, err := s.client.get(ctx, "/dashboardTemplate")
	if err != nil {
		return nil, fmt.Errorf("sc: list dashboardTemplates: %w", err)
	}

	var result DashboardTemplateListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTemplate list response: %w", err)
	}

	return &result, nil
}

// Get returns the dashboardTemplate with the given ID.
func (s *DashboardTemplateService) Get(ctx context.Context, id string) (*DashboardTemplate, error) {
	resp, err := s.client.get(ctx, "/dashboardTemplate"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get dashboardTemplate %s: %w", id, err)
	}

	var result DashboardTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTemplate response: %w", err)
	}

	return &result, nil
}

// Image performs the image action on the dashboardTemplate with the given ID.
func (s *DashboardTemplateService) Image(ctx context.Context, id string) (*DashboardTemplate, error) {
	resp, err := s.client.get(ctx, "/dashboardTemplate"+"/"+id+"/image")
	if err != nil {
		return nil, fmt.Errorf("sc: image dashboardTemplate %s: %w", id, err)
	}

	var result DashboardTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTemplate image response: %w", err)
	}

	return &result, nil
}

// Categories performs the categories action on the dashboardTemplate.
func (s *DashboardTemplateService) Categories(ctx context.Context) (*DashboardTemplate, error) {
	resp, err := s.client.get(ctx, "/dashboardTemplate/categories")
	if err != nil {
		return nil, fmt.Errorf("sc: categories dashboardTemplate: %w", err)
	}

	var result DashboardTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTemplate categories response: %w", err)
	}

	return &result, nil
}
