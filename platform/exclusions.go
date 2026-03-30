package platform

import (
	"context"
	"encoding/json"
)

// ExclusionsService handles communication with the Exclusions related endpoints of the Platform API.
type ExclusionsService struct {
	client *Client
}

// Exclusion represents an exclusion.
type Exclusion struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Value       string `json:"value,omitempty"`
	StartTime   int    `json:"start_time,omitempty"`
	EndTime     int    `json:"end_time,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
}

// ExclusionsListResponse represents the response from listing exclusions.
type ExclusionsListResponse struct {
	Exclusions []Exclusion `json:"exclusions,omitempty"`
	Total      int         `json:"total,omitempty"`
}

// ListExclusions returns all exclusions.
func (s *ExclusionsService) ListExclusions(ctx context.Context) (*ExclusionsListResponse, error) {
	resp, err := s.client.get(ctx, "/exclusions")
	if err != nil {
		return nil, err
	}

	var result ExclusionsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateExclusion creates a new exclusion.
func (s *ExclusionsService) CreateExclusion(ctx context.Context, req *Exclusion) (*Exclusion, error) {
	resp, err := s.client.post(ctx, "/exclusions", req)
	if err != nil {
		return nil, err
	}

	var result Exclusion
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetExclusion returns a specific exclusion by ID.
func (s *ExclusionsService) GetExclusion(ctx context.Context, exclusionId string) (*Exclusion, error) {
	resp, err := s.client.get(ctx, "/exclusions/"+exclusionId)
	if err != nil {
		return nil, err
	}

	var result Exclusion
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateExclusion updates an exclusion.
func (s *ExclusionsService) UpdateExclusion(ctx context.Context, exclusionId string, req *Exclusion) (*Exclusion, error) {
	resp, err := s.client.put(ctx, "/exclusions/"+exclusionId, req)
	if err != nil {
		return nil, err
	}

	var result Exclusion
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteExclusion deletes an exclusion.
func (s *ExclusionsService) DeleteExclusion(ctx context.Context, exclusionId string) error {
	_, err := s.client.delete(ctx, "/exclusions/"+exclusionId)
	return err
}
