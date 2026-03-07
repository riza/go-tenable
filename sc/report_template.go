
package sc

import (
	"encoding/json"
	"fmt"
)

// ReportTemplateService handles reportTemplate operations.
type ReportTemplateService struct {
	client *Client
}

// ReportTemplate represents a reportTemplate resource.
type ReportTemplate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ReportTemplateListResponse represents the response from listing reportTemplates.
type ReportTemplateListResponse struct {
	Usable     []ReportTemplate `json:"usable"`
	Manageable []ReportTemplate `json:"manageable"`
}

// List returns all reportTemplates.
func (s *ReportTemplateService) List() (*ReportTemplateListResponse, error) {
	resp, err := s.client.get("/reportTemplate")
	if err != nil {
		return nil, fmt.Errorf("sc: list reportTemplates: %w", err)
	}

	var result ReportTemplateListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal reportTemplate list response: %w", err)
	}

	return &result, nil
}

// Get returns the reportTemplate with the given ID.
func (s *ReportTemplateService) Get(id string) (*ReportTemplate, error) {
	resp, err := s.client.get("/reportTemplate" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get reportTemplate %s: %w", id, err)
	}

	var result ReportTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal reportTemplate response: %w", err)
	}

	return &result, nil
}

// Image performs the image action on the reportTemplate with the given ID.
func (s *ReportTemplateService) Image(id string) (*ReportTemplate, error) {
	resp, err := s.client.get("/reportTemplate" + "/" + id + "/image")
	if err != nil {
		return nil, fmt.Errorf("sc: image reportTemplate %s: %w", id, err)
	}

	var result ReportTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal reportTemplate image response: %w", err)
	}

	return &result, nil
}

// Categories performs the categories action on the reportTemplate.
func (s *ReportTemplateService) Categories() (*ReportTemplate, error) {
	resp, err := s.client.get("/reportTemplate/categories")
	if err != nil {
		return nil, fmt.Errorf("sc: categories reportTemplate: %w", err)
	}

	var result ReportTemplate
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal reportTemplate categories response: %w", err)
	}

	return &result, nil
}

