package sc

import (
	"context"
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

// QueryListResponse represents the response from listing queries.
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

// List returns all queries.
func (s *QueryService) List(ctx context.Context) (*QueryListResponse, error) {
	resp, err := s.client.get(ctx, "/query")
	if err != nil {
		return nil, fmt.Errorf("sc: list queries: %w", err)
	}

	var result QueryListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal query list response: %w", err)
	}

	return &result, nil
}

// Create creates a new query.
func (s *QueryService) Create(ctx context.Context, input *QueryCreateInput) (*Query, error) {
	resp, err := s.client.post(ctx, "/query", input)
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
func (s *QueryService) Get(ctx context.Context, id string) (*Query, error) {
	resp, err := s.client.get(ctx, "/query"+"/"+id)
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
func (s *QueryService) Update(ctx context.Context, id string, input *QueryUpdateInput) (*Query, error) {
	resp, err := s.client.patch(ctx, "/query"+"/"+id, input)
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
func (s *QueryService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/query"+"/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete query %s: %w", id, err)
	}

	return nil
}

// Share performs the share action on the query with the given ID.
func (s *QueryService) Share(ctx context.Context, id string) (*Query, error) {
	resp, err := s.client.post(ctx, "/query"+"/"+id+"/share", nil)
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
func (s *QueryService) Tag(ctx context.Context) (*Query, error) {
	resp, err := s.client.get(ctx, "/query/tag")
	if err != nil {
		return nil, fmt.Errorf("sc: tag query: %w", err)
	}

	var result Query
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal query tag response: %w", err)
	}

	return &result, nil
}
