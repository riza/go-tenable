package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DashboardComponentService handles dashboardComponent operations.
type DashboardComponentService struct {
	client *Client
}

// DashboardComponent represents a dashboardComponent resource.
type DashboardComponent struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DashboardComponentListResponse represents the response from listing dashboardComponents.
type DashboardComponentListResponse struct {
	Usable     []DashboardComponent `json:"usable"`
	Manageable []DashboardComponent `json:"manageable"`
}

// DashboardComponentCreateInput represents the request body for creating a dashboardComponent.
type DashboardComponentCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// DashboardComponentUpdateInput represents the request body for updating a dashboardComponent.
type DashboardComponentUpdateInput = DashboardComponentCreateInput

// List returns all dashboardComponents.
func (s *DashboardComponentService) List(ctx context.Context) (*DashboardComponentListResponse, error) {
	resp, err := s.client.get(ctx, "/dashboard/{dID}/component")
	if err != nil {
		return nil, fmt.Errorf("sc: list dashboardComponents: %w", err)
	}

	var result DashboardComponentListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardComponent list response: %w", err)
	}

	return &result, nil
}

// Create creates a new dashboardComponent.
func (s *DashboardComponentService) Create(ctx context.Context, input *DashboardComponentCreateInput) (*DashboardComponent, error) {
	resp, err := s.client.post(ctx, "/dashboard/{dID}/component", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create dashboardComponent: %w", err)
	}

	var result DashboardComponent
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardComponent response: %w", err)
	}

	return &result, nil
}

// Get returns the dashboardComponent with the given ID.
func (s *DashboardComponentService) Get(ctx context.Context, id string) (*DashboardComponent, error) {
	resp, err := s.client.get(ctx, "/dashboard/{dID}/component"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get dashboardComponent %s: %w", id, err)
	}

	var result DashboardComponent
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardComponent response: %w", err)
	}

	return &result, nil
}

// Update updates the dashboardComponent with the given ID.
func (s *DashboardComponentService) Update(ctx context.Context, id string, input *DashboardComponentUpdateInput) (*DashboardComponent, error) {
	resp, err := s.client.patch(ctx, "/dashboard/{dID}/component"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update dashboardComponent %s: %w", id, err)
	}

	var result DashboardComponent
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardComponent response: %w", err)
	}

	return &result, nil
}

// Delete deletes the dashboardComponent with the given ID.
func (s *DashboardComponentService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/dashboard/{dID}/component"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete dashboardComponent %s: %w", id, err)
	}

	return nil
}

// Copy performs the copy action on the dashboardComponent with the given ID.
func (s *DashboardComponentService) Copy(ctx context.Context, id string) (*DashboardComponent, error) {
	resp, err := s.client.post(ctx, "/dashboard/{dID}/component"+"/"+id+"/copy", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: copy dashboardComponent %s: %w", id, err)
	}

	var result DashboardComponent
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardComponent copy response: %w", err)
	}

	return &result, nil
}

// Refresh performs the refresh action on the dashboardComponent with the given ID.
func (s *DashboardComponentService) Refresh(ctx context.Context, id string) (*DashboardComponent, error) {
	resp, err := s.client.post(ctx, "/dashboard/{dID}/component"+"/"+id+"/refresh", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: refresh dashboardComponent %s: %w", id, err)
	}

	var result DashboardComponent
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardComponent refresh response: %w", err)
	}

	return &result, nil
}
