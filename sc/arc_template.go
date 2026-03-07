
package sc

import (
	"encoding/json"
	"fmt"
)

// ARCTemplateService handles aRCTemplate operations.
type ARCTemplateService struct {
	client *Client
}

// ARCTemplate represents a aRCTemplate resource.
type ARCTemplate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ARCTemplateListResponse represents the response from listing aRCTemplates.
type ARCTemplateListResponse struct {
	Usable     []ARCTemplate `json:"usable"`
	Manageable []ARCTemplate `json:"manageable"`
}

// List returns all aRCTemplates.
func (s *ARCTemplateService) List() (*ARCTemplateListResponse, error) {
	resp, err := s.client.get("/arcTemplate")
	if err != nil {
		return nil, fmt.Errorf("sc: list aRCTemplates: %w", err)
	}

	var result ARCTemplateListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRCTemplate list response: %w", err)
	}

	return &result, nil
}

// Get returns the aRCTemplate with the given ID.
func (s *ARCTemplateService) Get(id string) (*ARCTemplate, error) {
	resp, err := s.client.get("/arcTemplate" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get aRCTemplate %s: %w", id, err)
	}

	var result ARCTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRCTemplate response: %w", err)
	}

	return &result, nil
}

// Image performs the image action on the aRCTemplate with the given ID.
func (s *ARCTemplateService) Image(id string) (*ARCTemplate, error) {
	resp, err := s.client.get("/arcTemplate" + "/" + id + "/image")
	if err != nil {
		return nil, fmt.Errorf("sc: image aRCTemplate %s: %w", id, err)
	}

	var result ARCTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRCTemplate image response: %w", err)
	}

	return &result, nil
}

// Categories performs the categories action on the aRCTemplate.
func (s *ARCTemplateService) Categories() (*ARCTemplate, error) {
	resp, err := s.client.get("/arcTemplate/categories")
	if err != nil {
		return nil, fmt.Errorf("sc: categories aRCTemplate: %w", err)
	}

	var result ARCTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal aRCTemplate categories response: %w", err)
	}

	return &result, nil
}

