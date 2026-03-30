package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// AuditFileTemplateService handles auditFileTemplate operations.
type AuditFileTemplateService struct {
	client *Client
}

// AuditFileTemplate represents a auditFileTemplate resource.
type AuditFileTemplate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AuditFileTemplateListResponse represents the response from listing auditFileTemplates.
type AuditFileTemplateListResponse struct {
	Usable     []AuditFileTemplate `json:"usable"`
	Manageable []AuditFileTemplate `json:"manageable"`
}

// List returns all auditFileTemplates.
func (s *AuditFileTemplateService) List(ctx context.Context) (*AuditFileTemplateListResponse, error) {
	resp, err := s.client.get(ctx, "/auditFileTemplate")
	if err != nil {
		return nil, fmt.Errorf("sc: list auditFileTemplates: %w", err)
	}

	var result AuditFileTemplateListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal auditFileTemplate list response: %w", err)
	}

	return &result, nil
}
