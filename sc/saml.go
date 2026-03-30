package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// SAMLService handles sAML operations.
type SAMLService struct {
	client *Client
}

// SAML represents a sAML resource.
type SAML struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// SAMLListResponse represents the response from listing sAMLs.
type SAMLListResponse struct {
	Usable     []SAML `json:"usable"`
	Manageable []SAML `json:"manageable"`
}

// SAMLUpdateInput represents the request body for updating a sAML.
type SAMLUpdateInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// List returns all sAMLs.
func (s *SAMLService) List(ctx context.Context) (*SAMLListResponse, error) {
	resp, err := s.client.get(ctx, "/saml")
	if err != nil {
		return nil, fmt.Errorf("sc: list sAMLs: %w", err)
	}

	var result SAMLListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal sAML list response: %w", err)
	}

	return &result, nil
}

// Get returns the sAML with the given ID.
func (s *SAMLService) Get(ctx context.Context, id string) (*SAML, error) {
	resp, err := s.client.get(ctx, "/saml"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get sAML %s: %w", id, err)
	}

	var result SAML
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal sAML response: %w", err)
	}

	return &result, nil
}

// Update updates the sAML with the given ID.
func (s *SAMLService) Update(ctx context.Context, id string, input *SAMLUpdateInput) (*SAML, error) {
	resp, err := s.client.patch(ctx, "/saml"+"/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update sAML %s: %w", id, err)
	}

	var result SAML
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal sAML response: %w", err)
	}

	return &result, nil
}

// GetMetadata performs the getMetadata action on the sAML.
func (s *SAMLService) GetMetadata(ctx context.Context) (*SAML, error) {
	resp, err := s.client.get(ctx, "/saml/getMetadata")
	if err != nil {
		return nil, fmt.Errorf("sc: getMetadata sAML: %w", err)
	}

	var result SAML
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal sAML getMetadata response: %w", err)
	}

	return &result, nil
}
