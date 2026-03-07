
package sc

import (
	"encoding/json"
	"fmt"
)

// MDMService handles mDM operations.
type MDMService struct {
	client *Client
}

// MDM represents a mDM resource.
type MDM struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// MDMListResponse represents the response from listing mDMs.
type MDMListResponse struct {
	Usable     []MDM `json:"usable"`
	Manageable []MDM `json:"manageable"`
}

// List returns all mDMs.
func (s *MDMService) List() (*MDMListResponse, error) {
	resp, err := s.client.get("/mdm")
	if err != nil {
		return nil, fmt.Errorf("sc: list mDMs: %w", err)
	}

	var result MDMListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal mDM list response: %w", err)
	}

	return &result, nil
}

// Get returns the mDM with the given ID.
func (s *MDMService) Get(id string) (*MDM, error) {
	resp, err := s.client.get("/mdm" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get mDM %s: %w", id, err)
	}

	var result MDM
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal mDM response: %w", err)
	}

	return &result, nil
}

