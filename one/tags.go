package one

import (
	"context"
	"encoding/json"
)

// TagsService handles communication with the Tags related endpoints of the Tenable One API.
type TagsService struct {
	client *Client
}

// Tag represents a tag in Tenable One.
type Tag struct {
	Uuid      string                 `json:"uuid,omitempty"`
	Key       string                 `json:"key,omitempty"`
	Value     string                 `json:"value,omitempty"`
	Sources   []string               `json:"sources,omitempty"`
	CreatedAt string                 `json:"created_at,omitempty"`
	UpdatedAt string                 `json:"updated_at,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// TagsSearchRequest represents the request body for searching tags.
type TagsSearchRequest struct {
	Limit   *int               `json:"limit,omitempty"`
	Offset  *int               `json:"offset,omitempty"`
	Filters []TagsSearchFilter `json:"filters,omitempty"`
}

// TagsSearchFilter represents a generic filter for tag search.
type TagsSearchFilter struct {
	Property string      `json:"property"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

// TagsSearchResponse represents the response from searching tags.
type TagsSearchResponse struct {
	Tags       []Tag       `json:"tags,omitempty"`
	Total      int         `json:"total,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

// SearchTags searches for tags.
func (s *TagsService) SearchTags(ctx context.Context, req *TagsSearchRequest) (*TagsSearchResponse, error) {
	resp, err := s.client.post(ctx, "/api/v1/t1/tags/search", req)
	if err != nil {
		return nil, err
	}

	var result TagsSearchResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// TagsPropertiesResponse represents the response from getting tag properties.
type TagsPropertiesResponse struct {
	Properties []map[string]interface{} `json:"properties,omitempty"`
	Total      int                      `json:"total,omitempty"`
}

// GetProperties returns available properties for tag search.
func (s *TagsService) GetProperties(ctx context.Context) (*TagsPropertiesResponse, error) {
	resp, err := s.client.get(ctx, "/api/v1/t1/tags/properties")
	if err != nil {
		return nil, err
	}

	var result TagsPropertiesResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
