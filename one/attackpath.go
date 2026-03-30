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

// APAFilterCondition represents a single search filter for Attack Paths APIs.
type APAFilterCondition struct {
	Property string      `json:"property"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

// APAFilterGroup represents a logical combination (AND/OR) of multiple filter conditions.
type APAFilterGroup struct {
	Operator string               `json:"operator"`
	Value    []APAFilterCondition `json:"value"`
}

// APAPathNode represents a node (start or end) in an attack path.
type APAPathNode struct {
	Name         string            `json:"name,omitempty"`
	Description  string            `json:"description,omitempty"`
	Builtin      bool              `json:"_builtin,omitempty"`
	WhenCreated  float64           `json:"whencreated,omitempty"`
	ACR          float64           `json:"acr,omitempty"`
	AES          float64           `json:"aes,omitempty"`
	AVR          float64           `json:"avr,omitempty"`
	APAAES       float64           `json:"apa_aes,omitempty"`
	AssetID      string            `json:"asset_id,omitempty"`
	LastScanned  float64           `json:"last_scanned,omitempty"`
	LastObserved float64           `json:"last_observed,omitempty"`
	LastLicensed float64           `json:"last_licensed,omitempty"`
	Sensors      []string          `json:"sensors,omitempty"`
	Sources      []APANodeSource   `json:"sources,omitempty"`
}

// APANodeSource represents a source associated with a path node.
type APANodeSource struct {
	Name            string   `json:"name,omitempty"`
	FullName        string   `json:"fullname,omitempty"`
	ID              string   `json:"id,omitempty"`
	Labels          []string `json:"labels,omitempty"`
	IsCrownJewel    bool     `json:"isCrownJewel,omitempty"`
	VulnerabilityID string   `json:"vulnerability_id,omitempty"`
	AssetID         string   `json:"asset_id,omitempty"`
}

// APASegmentRelationship represents the relationship between two nodes in a path segment.
type APASegmentRelationship struct {
	ID         float64                `json:"id,omitempty"`
	Start      float64                `json:"start,omitempty"`
	End        float64                `json:"end,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// APAPathSegment represents a segment of an attack path between two nodes.
type APAPathSegment struct {
	Start        *APAPathNode            `json:"start,omitempty"`
	End          *APAPathNode            `json:"end,omitempty"`
	Relationship *APASegmentRelationship `json:"relationship,omitempty"`
}

// APAPath represents the graph path information for an attack path vector.
type APAPath struct {
	Start    *APAPathNode     `json:"start,omitempty"`
	End      *APAPathNode     `json:"end,omitempty"`
	Segments []APAPathSegment `json:"segments,omitempty"`
}

// APATechniqueMethod represents the procedure node for an attack technique.
type APATechniqueMethod struct {
	Name              string                 `json:"name,omitempty"`
	UUID              string                 `json:"uuid,omitempty"`
	VectorCount       int                    `json:"vector_count,omitempty"`
	Status            string                 `json:"status,omitempty"`
	State             string                 `json:"state,omitempty"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	PriorityScore     float64                `json:"priority_score,omitempty"`
	SourceInformation string                 `json:"source_information,omitempty"`
	Updated           float64                `json:"_updated,omitempty"`
}

// APATechnique represents an attack technique associated with a vector.
type APATechnique struct {
	Nodes             []string            `json:"nodes,omitempty"`
	Method            *APATechniqueMethod `json:"method,omitempty"`
	Technique         string              `json:"technique,omitempty"`
	SourceInformation string              `json:"source_information,omitempty"`
	WeaknessesIDs     []string            `json:"weaknessesIds,omitempty"`
	ExternalID        string              `json:"external_id,omitempty"`
	Target            string              `json:"target,omitempty"`
}

// APAAttackPath represents a single attack path (vector) returned by the search endpoint.
type APAAttackPath struct {
	IsNew         bool                     `json:"is_new,omitempty"`
	VectorID      string                   `json:"vector_id,omitempty"`
	Path          *APAPath                 `json:"path,omitempty"`
	Techniques    []APATechnique           `json:"techniques,omitempty"`
	Nodes         []map[string]interface{} `json:"nodes,omitempty"`
	FindingsNames []string                 `json:"findings_names,omitempty"`
	Name          string                   `json:"name,omitempty"`
	Summary       string                   `json:"summary,omitempty"`
	FirstAES      float64                  `json:"first_aes,omitempty"`
	LastACR       float64                  `json:"last_acr,omitempty"`
	PathStatus    string                   `json:"path_status,omitempty"`
}

// APASearchAttackPathsRequest represents the request parameters for searching attack paths.
type APASearchAttackPathsRequest struct {
	// Limit is the number of records to retrieve (query param). Min: 100, Max: 10000, Default: 1000.
	Limit int
	// Offset is the number of records to skip (query param). Default: 0.
	Offset int
	// Sort specifies the sort order in "property:direction" format, e.g. "priority:desc".
	Sort string
	// ExcludeResolved excludes resolved attack paths (done, chain_prevented, accepted). Default: true.
	ExcludeResolved *bool
	// RunAISummarization enables AI summarization for missing paths. Default: false.
	RunAISummarization *bool
	// Filter is the request body: either an APAFilterCondition or APAFilterGroup.
	// Sent directly as the JSON body (not wrapped).
	Filter interface{}
}

// APASearchAttackPathsResponse represents the response from searching attack paths.
type APASearchAttackPathsResponse struct {
	AttackPaths []APAAttackPath `json:"data,omitempty"`
	Total       int             `json:"total,omitempty"`
}

// SearchAttackPaths searches for attack paths.
//
// The filter (request body) is optional. When nil, no body is sent.
// Query parameters (limit, offset, sort, exclude_resolved, run_ai_summarization) are sent as URL params.
func (s *AttackPathService) SearchAttackPaths(ctx context.Context, req *APASearchAttackPathsRequest) (*APASearchAttackPathsResponse, error) {
	params := QueryParams{}
	var body interface{}

	if req != nil {
		if req.Limit > 0 {
			params["limit"] = []string{fmt.Sprintf("%d", req.Limit)}
		}
		if req.Offset > 0 {
			params["offset"] = []string{fmt.Sprintf("%d", req.Offset)}
		}
		if req.Sort != "" {
			params["sort"] = []string{req.Sort}
		}
		if req.ExcludeResolved != nil {
			params["exclude_resolved"] = []string{fmt.Sprintf("%t", *req.ExcludeResolved)}
		}
		if req.RunAISummarization != nil {
			params["run_ai_summarization"] = []string{fmt.Sprintf("%t", *req.RunAISummarization)}
		}
		body = req.Filter
	}

	resp, err := s.client.doRequestWithParams(ctx, "POST", "/api/v1/t1/apa/top-attack-paths/search", body, params)
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
