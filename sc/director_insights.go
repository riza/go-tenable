package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DirectorInsightsService handles directorInsights operations.
type DirectorInsightsService struct {
	client *Client
}

// DirectorInsights represents a directorInsights resource.
type DirectorInsights struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DirectorInsightsListResponse represents the response from listing directorInsightss.
type DirectorInsightsListResponse struct {
	Usable     []DirectorInsights `json:"usable"`
	Manageable []DirectorInsights `json:"manageable"`
}

// List returns all directorInsightss.
func (s *DirectorInsightsService) List(ctx context.Context) (*DirectorInsightsListResponse, error) {
	resp, err := s.client.get(ctx, "/mgmt/insights")
	if err != nil {
		return nil, fmt.Errorf("sc: list directorInsightss: %w", err)
	}

	var result DirectorInsightsListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorInsights list response: %w", err)
	}

	return &result, nil
}
