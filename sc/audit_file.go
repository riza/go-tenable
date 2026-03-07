
package sc

import (
	"encoding/json"
	"fmt"
)

// AuditFileService handles auditFile operations.
type AuditFileService struct {
	client *Client
}

// AuditFile represents a auditFile resource.
type AuditFile struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AuditFileListResponse represents the response from listing auditFiles.
type AuditFileListResponse struct {
	Usable     []AuditFile `json:"usable"`
	Manageable []AuditFile `json:"manageable"`
}

// AuditFileCreateInput represents the request body for creating a auditFile.
type AuditFileCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// AuditFileUpdateInput represents the request body for updating a auditFile.
type AuditFileUpdateInput = AuditFileCreateInput

// List returns all auditFiles.
func (s *AuditFileService) List() (*AuditFileListResponse, error) {
	resp, err := s.client.get("/auditFile")
	if err != nil {
		return nil, fmt.Errorf("sc: list auditFiles: %w", err)
	}

	var result AuditFileListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal auditFile list response: %w", err)
	}

	return &result, nil
}

// Create creates a new auditFile.
func (s *AuditFileService) Create(input *AuditFileCreateInput) (*AuditFile, error) {
	resp, err := s.client.post("/auditFile", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create auditFile: %w", err)
	}

	var result AuditFile
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal auditFile response: %w", err)
	}

	return &result, nil
}

// Get returns the auditFile with the given ID.
func (s *AuditFileService) Get(id string) (*AuditFile, error) {
	resp, err := s.client.get("/auditFile" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get auditFile %s: %w", id, err)
	}

	var result AuditFile
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal auditFile response: %w", err)
	}

	return &result, nil
}

// Update updates the auditFile with the given ID.
func (s *AuditFileService) Update(id string, input *AuditFileUpdateInput) (*AuditFile, error) {
	resp, err := s.client.patch("/auditFile" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update auditFile %s: %w", id, err)
	}

	var result AuditFile
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal auditFile response: %w", err)
	}

	return &result, nil
}

// Delete deletes the auditFile with the given ID.
func (s *AuditFileService) Delete(id string) error {
	_, err := s.client.delete("/auditFile" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete auditFile %s: %w", id, err)
	}

	return nil
}

// Refresh performs the refresh action on the auditFile with the given ID.
func (s *AuditFileService) Refresh(id string) (*AuditFile, error) {
	resp, err := s.client.post("/auditFile" + "/" + id + "/refresh", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: refresh auditFile %s: %w", id, err)
	}

	var result AuditFile
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal auditFile refresh response: %w", err)
	}

	return &result, nil
}

// Share performs the share action on the auditFile with the given ID.
func (s *AuditFileService) Share(id string) (*AuditFile, error) {
	resp, err := s.client.post("/auditFile" + "/" + id + "/share", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: share auditFile %s: %w", id, err)
	}

	var result AuditFile
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal auditFile share response: %w", err)
	}

	return &result, nil
}

// Export performs the export action on the auditFile with the given ID.
func (s *AuditFileService) Export(id string) (*AuditFile, error) {
	resp, err := s.client.get("/auditFile" + "/" + id + "/export")
	if err != nil {
		return nil, fmt.Errorf("sc: export auditFile %s: %w", id, err)
	}

	var result AuditFile
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal auditFile export response: %w", err)
	}

	return &result, nil
}

