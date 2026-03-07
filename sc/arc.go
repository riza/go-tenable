
package sc

import (
	"encoding/json"
	"fmt"
)

// ARCService handles aRC operations.
type ARCService struct {
	client *Client
}

// ARC represents a aRC resource.
type ARC struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ARCListResponse represents the response from listing aRCs.
type ARCListResponse struct {
	Usable     []ARC `json:"usable"`
	Manageable []ARC `json:"manageable"`
}

// ARCCreateInput represents the request body for creating a aRC.
type ARCCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// ARCUpdateInput represents the request body for updating a aRC.
type ARCUpdateInput = ARCCreateInput

// List returns all aRCs.
func (s *ARCService) List() (*ARCListResponse, error) {
	resp, err := s.client.get("/arc")
	if err != nil {
		return nil, fmt.Errorf("sc: list aRCs: %w", err)
	}

	var result ARCListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRC list response: %w", err)
	}

	return &result, nil
}

// Create creates a new aRC.
func (s *ARCService) Create(input *ARCCreateInput) (*ARC, error) {
	resp, err := s.client.post("/arc", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create aRC: %w", err)
	}

	var result ARC
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRC response: %w", err)
	}

	return &result, nil
}

// Get returns the aRC with the given ID.
func (s *ARCService) Get(id string) (*ARC, error) {
	resp, err := s.client.get("/arc" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get aRC %s: %w", id, err)
	}

	var result ARC
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRC response: %w", err)
	}

	return &result, nil
}

// Update updates the aRC with the given ID.
func (s *ARCService) Update(id string, input *ARCUpdateInput) (*ARC, error) {
	resp, err := s.client.patch("/arc" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update aRC %s: %w", id, err)
	}

	var result ARC
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRC response: %w", err)
	}

	return &result, nil
}

// Delete deletes the aRC with the given ID.
func (s *ARCService) Delete(id string) error {
	_, err := s.client.delete("/arc" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete aRC %s: %w", id, err)
	}

	return nil
}

// Import performs the import action on the aRC.
func (s *ARCService) Import() (*ARC, error) {
	resp, err := s.client.post("/arc/import", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: import aRC: %w", err)
	}

	var result ARC
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRC import response: %w", err)
	}

	return &result, nil
}

// Export performs the export action on the aRC with the given ID.
func (s *ARCService) Export(id string) (*ARC, error) {
	resp, err := s.client.post("/arc" + "/" + id + "/export", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: export aRC %s: %w", id, err)
	}

	var result ARC
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRC export response: %w", err)
	}

	return &result, nil
}

// Copy performs the copy action on the aRC with the given ID.
func (s *ARCService) Copy(id string) (*ARC, error) {
	resp, err := s.client.post("/arc" + "/" + id + "/copy", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: copy aRC %s: %w", id, err)
	}

	var result ARC
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRC copy response: %w", err)
	}

	return &result, nil
}

// Refresh performs the refresh action on the aRC with the given ID.
func (s *ARCService) Refresh(id string) (*ARC, error) {
	resp, err := s.client.post("/arc" + "/" + id + "/refresh", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: refresh aRC %s: %w", id, err)
	}

	var result ARC
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRC refresh response: %w", err)
	}

	return &result, nil
}

// Share performs the share action on the aRC with the given ID.
func (s *ARCService) Share(id string) (*ARC, error) {
	resp, err := s.client.post("/arc" + "/" + id + "/share", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: share aRC %s: %w", id, err)
	}

	var result ARC
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRC share response: %w", err)
	}

	return &result, nil
}

