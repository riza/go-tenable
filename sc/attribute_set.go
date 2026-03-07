
package sc

import (
	"encoding/json"
	"fmt"
)

// AttributeSetService handles attributeSet operations.
type AttributeSetService struct {
	client *Client
}

// AttributeSet represents a attributeSet resource.
type AttributeSet struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AttributeSetListResponse represents the response from listing attributeSets.
type AttributeSetListResponse struct {
	Usable     []AttributeSet `json:"usable"`
	Manageable []AttributeSet `json:"manageable"`
}

// AttributeSetCreateInput represents the request body for creating a attributeSet.
type AttributeSetCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// AttributeSetUpdateInput represents the request body for updating a attributeSet.
type AttributeSetUpdateInput = AttributeSetCreateInput

// List returns all attributeSets.
func (s *AttributeSetService) List() (*AttributeSetListResponse, error) {
	resp, err := s.client.get("/attributeSet")
	if err != nil {
		return nil, fmt.Errorf("sc: list attributeSets: %w", err)
	}

	var result AttributeSetListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal attributeSet list response: %w", err)
	}

	return &result, nil
}

// Create creates a new attributeSet.
func (s *AttributeSetService) Create(input *AttributeSetCreateInput) (*AttributeSet, error) {
	resp, err := s.client.post("/attributeSet", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create attributeSet: %w", err)
	}

	var result AttributeSet
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal attributeSet response: %w", err)
	}

	return &result, nil
}

// Get returns the attributeSet with the given ID.
func (s *AttributeSetService) Get(id string) (*AttributeSet, error) {
	resp, err := s.client.get("/attributeSet" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get attributeSet %s: %w", id, err)
	}

	var result AttributeSet
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal attributeSet response: %w", err)
	}

	return &result, nil
}

// Update updates the attributeSet with the given ID.
func (s *AttributeSetService) Update(id string, input *AttributeSetUpdateInput) (*AttributeSet, error) {
	resp, err := s.client.patch("/attributeSet" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update attributeSet %s: %w", id, err)
	}

	var result AttributeSet
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal attributeSet response: %w", err)
	}

	return &result, nil
}

// Delete deletes the attributeSet with the given ID.
func (s *AttributeSetService) Delete(id string) error {
	_, err := s.client.delete("/attributeSet" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete attributeSet %s: %w", id, err)
	}

	return nil
}

// Types performs the types action on the attributeSet.
func (s *AttributeSetService) Types() (*AttributeSet, error) {
	resp, err := s.client.get("/attributeSet/types")
	if err != nil {
		return nil, fmt.Errorf("sc: types attributeSet: %w", err)
	}

	var result AttributeSet
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal attributeSet types response: %w", err)
	}

	return &result, nil
}

