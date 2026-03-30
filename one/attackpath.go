package one

import (
	"context"
	"encoding/json"
	"fmt"
)

// AttackPathService handles communication with the Attack Path related endpoints of the Tenable One API.
type AttackPathService struct {
	client *Client
}

// APAAttackPath represents an attack path in the APA service.
type APAAttackPath struct {
	ID         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Score      int         `json:"score,omitempty"`
	AssetCount int         `json:"asset_count,omitempty"`
	CreatedAt  string      `json:"created_at,omitempty"`
	UpdatedAt  string      `json:"updated_at,omitempty"`
	Nodes      interface{} `json:"nodes,omitempty"`
	Edges      interface{} `json:"edges,omitempty"`
	Metadata   interface{} `json:"metadata,omitempty"`
}

// APAFilterCondition represents a single search filter for Attack Paths APIs.
type APAFilterCondition struct {
	Property string      `json:"property,omitempty"`
	Operator string      `json:"operator,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}

// APAFilterGroup represents a logical combination (AND/OR) of multiple filter conditions.
type APAFilterGroup struct {
	Operator string               `json:"operator,omitempty"`
	Value    []APAFilterCondition `json:"value,omitempty"`
}

// APASearchAttackPathsRequest represents the request body for searching attack paths.
type APASearchAttackPathsRequest struct {
	Limit  int         `json:"-"`
	Offset int         `json:"-"`
	Filter interface{} `json:"filter,omitempty"`
}

// APASearchAttackPathsResponse represents the response from searching attack paths.
type APASearchAttackPathsResponse struct {
	AttackPaths []APAAttackPath `json:"data,omitempty"`
	Total       int             `json:"total,omitempty"`
	Pagination  *PaginationInfo `json:"pagination,omitempty"`
}

// SearchAttackPaths searches for attack paths.
func (s *AttackPathService) SearchAttackPaths(ctx context.Context, req *APASearchAttackPathsRequest) (*APASearchAttackPathsResponse, error) {
	params := QueryParams{}
	if req != nil {
		if req.Limit > 0 {
			params["limit"] = []string{fmt.Sprintf("%d", req.Limit)}
		}
		if req.Offset > 0 {
			params["offset"] = []string{fmt.Sprintf("%d", req.Offset)}
		}
	}

	var payload interface{}
	if req != nil && req.Filter != nil {
		payload = req.Filter
	} else {
		payload = map[string]interface{}{}
	}

	resp, err := s.client.doRequestWithParams(ctx, "POST", "/api/v1/t1/apa/top-attack-paths/search", payload, params)
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
	MitreID       string   `json:"mitre_id,omitempty"`
	TechniqueName string   `json:"technique_name,omitempty"`
	Tactics       []string `json:"tactics,omitempty"`
	Count         int      `json:"vector_count,omitempty"`
	Score         int      `json:"score,omitempty"`
	Priority      string   `json:"priority,omitempty"`
}

// APASearchAttackTechniquesRequest represents the request body for searching attack techniques.
type APASearchAttackTechniquesRequest struct {
	Limit  int         `json:"-"`
	Offset int         `json:"-"`
	Filter interface{} `json:"filter,omitempty"`
}

// APASearchAttackTechniquesResponse represents the response from searching attack techniques.
type APASearchAttackTechniquesResponse struct {
	Techniques []APAAttackTechnique `json:"data,omitempty"`
	Total      int                  `json:"total,omitempty"`
	Pagination *PaginationInfo      `json:"pagination,omitempty"`
}

// SearchAttackTechniques searches for attack techniques.
func (s *AttackPathService) SearchAttackTechniques(ctx context.Context, req *APASearchAttackTechniquesRequest) (*APASearchAttackTechniquesResponse, error) {
	params := QueryParams{}
	if req != nil {
		if req.Limit > 0 {
			params["limit"] = []string{fmt.Sprintf("%d", req.Limit)}
		}
		if req.Offset > 0 {
			params["offset"] = []string{fmt.Sprintf("%d", req.Offset)}
		}
	}

	var payload interface{}
	if req != nil && req.Filter != nil {
		payload = req.Filter
	} else {
		payload = map[string]interface{}{}
	}

	resp, err := s.client.doRequestWithParams(ctx, "POST", "/api/v1/t1/apa/top-attack-techniques/search", payload, params)
	if err != nil {
		return nil, err
	}

	var result APASearchAttackTechniquesResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
