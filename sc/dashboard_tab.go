package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DashboardTabService handles dashboardTab operations.
type DashboardTabService struct {
	client *Client
}

// DashboardTab represents a dashboardTab resource.
type DashboardTab struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DashboardTabListResponse represents the response from listing dashboardTabs.
type DashboardTabListResponse struct {
	Usable     []DashboardTab `json:"usable"`
	Manageable []DashboardTab `json:"manageable"`
}

// DashboardTabCreateInput represents the request body for creating a dashboardTab.
type DashboardTabCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// DashboardTabUpdateInput represents the request body for updating a dashboardTab.
type DashboardTabUpdateInput = DashboardTabCreateInput

// List returns all dashboardTabs.
func (s *DashboardTabService) List(ctx context.Context) (*DashboardTabListResponse, error) {
	resp, err := s.client.get(ctx, "/dashboard")
	if err != nil {
		return nil, fmt.Errorf("sc: list dashboardTabs: %w", err)
	}

	var result DashboardTabListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTab list response: %w", err)
	}

	return &result, nil
}

// Create creates a new dashboardTab.
func (s *DashboardTabService) Create(ctx context.Context, input *DashboardTabCreateInput) (*DashboardTab, error) {
	resp, err := s.client.post(ctx, "/dashboard", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create dashboardTab: %w", err)
	}

	var result DashboardTab
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTab response: %w", err)
	}

	return &result, nil
}

// Get returns the dashboardTab with the given ID.
func (s *DashboardTabService) Get(ctx context.Context, id string) (*DashboardTab, error) {
	resp, err := s.client.get(ctx, "/dashboard"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get dashboardTab %s: %w", id, err)
	}

	var result DashboardTab
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTab response: %w", err)
	}

	return &result, nil
}

// Update updates the dashboardTab with the given ID.
func (s *DashboardTabService) Update(ctx context.Context, id string, input *DashboardTabUpdateInput) (*DashboardTab, error) {
	resp, err := s.client.patch(ctx, "/dashboard"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update dashboardTab %s: %w", id, err)
	}

	var result DashboardTab
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTab response: %w", err)
	}

	return &result, nil
}

// Delete deletes the dashboardTab with the given ID.
func (s *DashboardTabService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/dashboard"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete dashboardTab %s: %w", id, err)
	}

	return nil
}

// Copy performs the copy action on the dashboardTab with the given ID.
func (s *DashboardTabService) Copy(ctx context.Context, id string) (*DashboardTab, error) {
	resp, err := s.client.post(ctx, "/dashboard"+"/"+id+"/copy", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: copy dashboardTab %s: %w", id, err)
	}

	var result DashboardTab
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTab copy response: %w", err)
	}

	return &result, nil
}

// Import performs the import action on the dashboardTab.
func (s *DashboardTabService) Import(ctx context.Context) (*DashboardTab, error) {
	resp, err := s.client.post(ctx, "/dashboard/import", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: import dashboardTab: %w", err)
	}

	var result DashboardTab
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTab import response: %w", err)
	}

	return &result, nil
}

// Export performs the export action on the dashboardTab with the given ID.
func (s *DashboardTabService) Export(ctx context.Context, id string) (*DashboardTab, error) {
	resp, err := s.client.post(ctx, "/dashboard"+"/"+id+"/export", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: export dashboardTab %s: %w", id, err)
	}

	var result DashboardTab
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTab export response: %w", err)
	}

	return &result, nil
}

// Share performs the share action on the dashboardTab with the given ID.
func (s *DashboardTabService) Share(ctx context.Context, id string) (*DashboardTab, error) {
	resp, err := s.client.post(ctx, "/dashboard"+"/"+id+"/share", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: share dashboardTab %s: %w", id, err)
	}

	var result DashboardTab
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal dashboardTab share response: %w", err)
	}

	return &result, nil
}
