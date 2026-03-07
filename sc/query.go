
package sc

import (
	"encoding/json"
	"fmt"
)

// QueryService handles query operations.
type QueryService struct {
	client *Client
}

// Query represents a query resource.
type Query struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// QueryListResponse represents the response from listing querys.
type QueryListResponse struct {
	Usable     []Query `json:"usable"`
	Manageable []Query `json:"manageable"`
}

// QueryCreateInput represents the request body for creating a query.
type QueryCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// QueryUpdateInput represents the request body for updating a query.
type QueryUpdateInput = QueryCreateInput

// List returns all querys.
func (s *QueryService) List() (*QueryListResponse, error) {
	resp, err := s.client.get("/query")
	if err != nil {
		return nil, fmt.Errorf("sc: list querys: %w", err)
	}

	var result QueryListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal query list response: %w", err)
	}

	return &result, nil
}

// Create creates a new query.
func (s *QueryService) Create(input *QueryCreateInput) (*Query, error) {
	resp, err := s.client.post("/query", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create query: %w", err)
	}

	var result Query
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal query response: %w", err)
	}

	return &result, nil
}

// Get returns the query with the given ID.
func (s *QueryService) Get(id string) (*Query, error) {
	resp, err := s.client.get("/query" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get query %s: %w", id, err)
	}

	var result Query
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal query response: %w", err)
	}

	return &result, nil
}

// Update updates the query with the given ID.
func (s *QueryService) Update(id string, input *QueryUpdateInput) (*Query, error) {
	resp, err := s.client.patch("/query" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update query %s: %w", id, err)
	}

	var result Query
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal query response: %w", err)
	}

	return &result, nil
}

// Delete deletes the query with the given ID.
func (s *QueryService) Delete(id string) error {
	_, err := s.client.delete("/query" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete query %s: %w", id, err)
	}

	return nil
}

// Share performs the share action on the query with the given ID.
func (s *QueryService) Share(id string) (*Query, error) {
	resp, err := s.client.post("/query" + "/" + id + "/share", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: share query %s: %w", id, err)
	}

	var result Query
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal query share response: %w", err)
	}

	return &result, nil
}

// Tag performs the tag action on the query.
func (s *QueryService) Tag() (*Query, error) {
	resp, err := s.client.get("/query/tag")
	if err != nil {
		return nil, fmt.Errorf("sc: tag query: %w", err)
	}

	var result Query
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal query tag response: %w", err)
	}

	return &result, nil
}

