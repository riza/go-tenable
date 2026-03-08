package one

import (
	"context"
	"encoding/json"
)

// AttackPathService handles communication with the Attack Path related endpoints of the Tenable One API.
type AttackPathService struct {
	client *Client
}

// APAAttackPath represents an attack path in the APA service.
type APAAttackPath struct {
	Id          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Score       int         `json:"score,omitempty"`
	AssetCount  int         `json:"asset_count,omitempty"`
	CreatedAt   string      `json:"created_at,omitempty"`
	UpdatedAt   string      `json:"updated_at,omitempty"`
	Nodes       interface{} `json:"nodes,omitempty"`
	Edges       interface{} `json:"edges,omitempty"`
	Metadata    interface{} `json:"metadata,omitempty"`
}

// APASearchAttackPathsRequest represents the request body for searching attack paths.
type APASearchAttackPathsRequest struct {
	Limit   int                    `json:"limit,omitempty"`
	Offset  int                    `json:"offset,omitempty"`
	Filters APASearchAttackPathsFilters `json:"filters,omitempty"`
}

// APASearchAttackPathsFilters represents the filters for attack path search.
type APASearchAttackPathsFilters struct {
	AssetIds    []string `json:"asset_ids,omitempty"`
	TagIds      []string `json:"tag_ids,omitempty"`
	TagTargets  []string `json:"tag_targets,omitempty"`
	Sources     []string `json:"sources,omitempty"`
	NetworkIds  []string `json:"network_ids,omitempty"`
}

// APASearchAttackPathsResponse represents the response from searching attack paths.
type APASearchAttackPathsResponse struct {
	AttackPaths []APAAttackPath `json:"attack_paths,omitempty"`
	Total       int             `json:"total,omitempty"`
	Pagination  interface{}     `json:"pagination,omitempty"`
}

// SearchAttackPaths searches for attack paths.
func (s *AttackPathService) SearchAttackPaths(ctx context.Context, req *APASearchAttackPathsRequest) (*APASearchAttackPathsResponse, error) {
	resp, err := s.client.post(ctx, "/api/v1/t1/apa/top-attack-paths/search", req)
	if err != nil {
		return nil, err
	}

	var result APASearchAttackPathsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// APAAttackTechnique represents an attack technique in the APA service.
type APAAttackTechnique struct {
	Id          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Category    string      `json:"category,omitempty"`
	Count       int         `json:"count,omitempty"`
	Score       int         `json:"score,omitempty"`
	Severity    string      `json:"severity,omitempty"`
}

// APASearchAttackTechniquesRequest represents the request body for searching attack techniques.
type APASearchAttackTechniquesRequest struct {
	Limit   int                           `json:"limit,omitempty"`
	Offset  int                           `json:"offset,omitempty"`
	Filters APASearchAttackTechniquesFilters `json:"filters,omitempty"`
}

// APASearchAttackTechniquesFilters represents the filters for attack technique search.
type APASearchAttackTechniquesFilters struct {
	AssetIds    []string `json:"asset_ids,omitempty"`
	TagIds      []string `json:"tag_ids,omitempty"`
	TagTargets  []string `json:"tag_targets,omitempty"`
	Sources     []string `json:"sources,omitempty"`
	NetworkIds  []string `json:"network_ids,omitempty"`
}

// APASearchAttackTechniquesResponse represents the response from searching attack techniques.
type APASearchAttackTechniquesResponse struct {
	Techniques []APAAttackTechnique `json:"attack_techniques,omitempty"`
	Total      int                  `json:"total,omitempty"`
	Pagination interface{}          `json:"pagination,omitempty"`
}

// SearchAttackTechniques searches for attack techniques.
func (s *AttackPathService) SearchAttackTechniques(ctx context.Context, req *APASearchAttackTechniquesRequest) (*APASearchAttackTechniquesResponse, error) {
	resp, err := s.client.post(ctx, "/api/v1/t1/apa/top-attack-techniques/search", req)
	if err != nil {
		return nil, err
	}

	var result APASearchAttackTechniquesResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
