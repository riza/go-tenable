package sc

import (
	"encoding/json"
	"fmt"
)

// ReportDefinitionService handles communication with the report-definition-related endpoints of the SC API.
type ReportDefinitionService struct {
	client *Client
}

// ReportDefinition represents a report definition resource from the API.
type ReportDefinition struct {
	ID                 string                     `json:"id"`
	Name               string                     `json:"name"`
	Description        string                     `json:"description"`
	Type               string                     `json:"type"`
	StyleFamily        string                     `json:"styleFamily"`
	Definition         json.RawMessage            `json:"definition,omitempty"`
	XMLDefinition      string                     `json:"xmldefinition"`
	EncryptionPassword string                     `json:"encryptionPassword"`
	Status             string                     `json:"status"`
	ShareUsers         []IDRef                    `json:"shareUsers,omitempty"`
	EmailUsers         []IDRef                    `json:"emailUsers,omitempty"`
	EmailTargets       string                     `json:"emailTargets"`
	EmailBCCTargets    string                     `json:"emailBCCTargets"`
	EmailTargetType    string                     `json:"emailTargetType"`
	CreatedTime        string                     `json:"createdTime"`
	ModifiedTime       string                     `json:"modifiedTime"`
	PubSites           []IDRef                    `json:"pubSites,omitempty"`
	Sources            json.RawMessage            `json:"sources,omitempty"`
	Components         json.RawMessage            `json:"components,omitempty"`
	Iterators          json.RawMessage            `json:"iterators,omitempty"`
	ScanResult         json.RawMessage            `json:"scanResult,omitempty"`
	Schedule           *ReportDefinitionSchedule  `json:"schedule,omitempty"`
	Creator            *IDRef                     `json:"creator,omitempty"`
	Owner              *IDRef                     `json:"owner,omitempty"`
	OwnerGroup         *IDRef                     `json:"ownerGroup,omitempty"`
	AttributeSet       *IDRef                     `json:"attributeSet,omitempty"`
	QueryStatus        json.RawMessage            `json:"queryStatus,omitempty"`
}

// ReportDefinitionSchedule represents the schedule configuration for a report definition.
type ReportDefinitionSchedule struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Start      string `json:"start"`
	RepeatRule string `json:"repeatRule"`
	NextRun    int    `json:"nextRun"`
}

// ReportDefinitionListResponse holds the usable and manageable report definition arrays returned by the list endpoint.
type ReportDefinitionListResponse struct {
	Usable     []ReportDefinition `json:"usable"`
	Manageable []ReportDefinition `json:"manageable"`
}

// ReportDefinitionCreateInput contains the fields for creating a new report definition.
type ReportDefinitionCreateInput struct {
	Name               string                    `json:"name"`
	Type               string                    `json:"type,omitempty"`
	Description        string                    `json:"description,omitempty"`
	StyleFamily        string                    `json:"styleFamily,omitempty"`
	Definition         json.RawMessage           `json:"definition,omitempty"`
	EncryptionPassword string                    `json:"encryptionPassword,omitempty"`
	ShareUsers         []IDRef                   `json:"shareUsers,omitempty"`
	EmailUsers         []IDRef                   `json:"emailUsers,omitempty"`
	EmailTargets       string                    `json:"emailTargets,omitempty"`
	EmailBCCTargets    string                    `json:"emailBCCTargets,omitempty"`
	EmailTargetType    string                    `json:"emailTargetType,omitempty"`
	PubSites           []IDRef                   `json:"pubSites,omitempty"`
	Sources            json.RawMessage           `json:"sources,omitempty"`
	Components         json.RawMessage           `json:"components,omitempty"`
	Iterators          json.RawMessage           `json:"iterators,omitempty"`
	Schedule           *ReportDefinitionSchedule `json:"schedule,omitempty"`
	AttributeSet       *IDRef                    `json:"attributeSet,omitempty"`
}

// ReportDefinitionUpdateInput contains the fields for updating an existing report definition.
type ReportDefinitionUpdateInput = ReportDefinitionCreateInput

// ReportDefinitionLaunchResponse holds the response from launching a report definition.
type ReportDefinitionLaunchResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	JobID  string `json:"jobID"`
}

// ReportDefinitionCopyInput contains the fields for copying a report definition.
type ReportDefinitionCopyInput struct {
	Name         string `json:"name,omitempty"`
	TargetUserID string `json:"targetUserID,omitempty"`
}

// ReportDefinitionExportInput contains the fields for exporting a report definition.
type ReportDefinitionExportInput struct {
	ExportType string `json:"exportType,omitempty"`
}

// ReportDefinitionImportInput contains the fields for importing a report definition.
type ReportDefinitionImportInput struct {
	Name     string `json:"name,omitempty"`
	Filename string `json:"filename,omitempty"`
}

// ReportDefinitionImportResponse holds the response from importing a report definition.
type ReportDefinitionImportResponse struct {
	Name     string `json:"name"`
	Filename string `json:"filename"`
}

// List returns the list of report definitions (usable and manageable).
func (s *ReportDefinitionService) List() (*ReportDefinitionListResponse, error) {
	resp, err := s.client.get("/reportDefinition")
	if err != nil {
		return nil, fmt.Errorf("sc: list report definitions: %w", err)
	}

	var result ReportDefinitionListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report definition list response: %w", err)
	}

	return &result, nil
}

// Create creates a new report definition with the given input.
func (s *ReportDefinitionService) Create(input *ReportDefinitionCreateInput) (*ReportDefinition, error) {
	resp, err := s.client.post("/reportDefinition", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create report definition: %w", err)
	}

	var rd ReportDefinition
	if err := json.Unmarshal(resp.Response, &rd); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report definition response: %w", err)
	}

	return &rd, nil
}

// Get returns the report definition with the given ID.
func (s *ReportDefinitionService) Get(id string) (*ReportDefinition, error) {
	resp, err := s.client.get("/reportDefinition/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get report definition %s: %w", id, err)
	}

	var rd ReportDefinition
	if err := json.Unmarshal(resp.Response, &rd); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report definition response: %w", err)
	}

	return &rd, nil
}

// Update updates an existing report definition with the given input.
func (s *ReportDefinitionService) Update(id string, input *ReportDefinitionUpdateInput) (*ReportDefinition, error) {
	resp, err := s.client.patch("/reportDefinition/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update report definition %s: %w", id, err)
	}

	var rd ReportDefinition
	if err := json.Unmarshal(resp.Response, &rd); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report definition response: %w", err)
	}

	return &rd, nil
}

// Delete deletes the report definition with the given ID.
func (s *ReportDefinitionService) Delete(id string) error {
	_, err := s.client.delete("/reportDefinition/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete report definition %s: %w", id, err)
	}

	return nil
}

// Launch launches the report definition with the given ID.
func (s *ReportDefinitionService) Launch(id string) (*ReportDefinitionLaunchResponse, error) {
	resp, err := s.client.post("/reportDefinition/"+id+"/launch", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: launch report definition %s: %w", id, err)
	}

	var result ReportDefinitionLaunchResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report definition launch response: %w", err)
	}

	return &result, nil
}

// Copy copies the report definition with the given ID.
func (s *ReportDefinitionService) Copy(id string, input *ReportDefinitionCopyInput) (*ReportDefinition, error) {
	resp, err := s.client.post("/reportDefinition/"+id+"/copy", input)
	if err != nil {
		return nil, fmt.Errorf("sc: copy report definition %s: %w", id, err)
	}

	var rd ReportDefinition
	if err := json.Unmarshal(resp.Response, &rd); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report definition copy response: %w", err)
	}

	return &rd, nil
}

// Export exports the report definition with the given ID. The response is the raw binary content.
func (s *ReportDefinitionService) Export(id string, input *ReportDefinitionExportInput) (json.RawMessage, error) {
	resp, err := s.client.post("/reportDefinition/"+id+"/export", input)
	if err != nil {
		return nil, fmt.Errorf("sc: export report definition %s: %w", id, err)
	}

	return resp.Response, nil
}

// Import imports a report definition from the given input.
func (s *ReportDefinitionService) Import(input *ReportDefinitionImportInput) (*ReportDefinitionImportResponse, error) {
	resp, err := s.client.post("/reportDefinition/import", input)
	if err != nil {
		return nil, fmt.Errorf("sc: import report definition: %w", err)
	}

	var result ReportDefinitionImportResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report definition import response: %w", err)
	}

	return &result, nil
}
