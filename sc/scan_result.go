package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// ScanResultService handles communication with the scan result-related endpoints of the SC API.
type ScanResultService struct {
	client *Client
}

// ScanResult represents a scan result resource from the API.
type ScanResult struct {
	ID                     string          `json:"id"`
	Name                   string          `json:"name"`
	Description            string          `json:"description"`
	Details                string          `json:"details"`
	Status                 string          `json:"status"`
	ImportStatus           string          `json:"importStatus"`
	ImportStart            string          `json:"importStart"`
	ImportFinish           string          `json:"importFinish"`
	ImportDuration         string          `json:"importDuration"`
	DownloadAvailable      string          `json:"downloadAvailable"`
	DownloadFormat         string          `json:"downloadFormat"`
	DataFormat             string          `json:"dataFormat"`
	ResultType             string          `json:"resultType"`
	ResultSource           string          `json:"resultSource"`
	Running                string          `json:"running"`
	ErrorDetails           string          `json:"errorDetails"`
	ImportErrorDetails     string          `json:"importErrorDetails"`
	TotalIPs               string          `json:"totalIPs"`
	ScannedIPs             string          `json:"scannedIPs"`
	StartTime              string          `json:"startTime"`
	FinishTime             string          `json:"finishTime"`
	CreatedTime            string          `json:"createdTime"`
	ScanDuration           string          `json:"scanDuration"`
	CompletedIPs           string          `json:"completedIPs"`
	CompletedChecks        string          `json:"completedChecks"`
	TotalChecks            string          `json:"totalChecks"`
	AgentScanUUID          string          `json:"agentScanUUID"`
	AgentScanContainerUUID string          `json:"agentScanContainerUUID"`
	DiagnosticAvailable    string          `json:"diagnosticAvailable"`
	CanUse                 string          `json:"canUse"`
	CanManage              string          `json:"canManage"`
	Initiator              *IDRef          `json:"initiator"`
	Owner                  *IDRef          `json:"owner"`
	OwnerGroup             *IDRef          `json:"ownerGroup"`
	Scan                   *IDRef          `json:"scan"`
	Repository             *IDRef          `json:"repository"`
	Progress               json.RawMessage `json:"progress"`
}

// ScanResultListResponse holds the usable and manageable scan result arrays returned by the list endpoint.
type ScanResultListResponse struct {
	Usable     []ScanResult `json:"usable"`
	Manageable []ScanResult `json:"manageable"`
}

// ScanResultCopyInput contains the fields for copying a scan result.
type ScanResultCopyInput struct {
	Users []IDRef `json:"users,omitempty"`
}

// ScanResultEmailInput contains the fields for emailing a scan result.
type ScanResultEmailInput struct {
	Email string `json:"email"`
}

// ScanResultImportInput contains the fields for importing a scan result.
type ScanResultImportInput struct {
	Filename   string `json:"filename"`
	Repository *IDRef `json:"repository"`
}

// ScanResultReimportInput contains the fields for reimporting a scan result.
type ScanResultReimportInput struct {
	Repository *IDRef `json:"repository,omitempty"`
}

// ScanResultDownloadInput contains the fields for downloading a scan result.
type ScanResultDownloadInput struct {
	DownloadType string `json:"downloadType"`
}

// List returns the list of scan results (usable and manageable).
func (s *ScanResultService) List(ctx context.Context) (*ScanResultListResponse, error) {
	resp, err := s.client.get(ctx, "/scanResult")
	if err != nil {
		return nil, fmt.Errorf("sc: list scan results: %w", err)
	}

	var result ScanResultListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan result list response: %w", err)
	}

	return &result, nil
}

// Get returns the scan result with the given ID.
func (s *ScanResultService) Get(ctx context.Context, id string) (*ScanResult, error) {
	resp, err := s.client.get(ctx, "/scanResult/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get scan result %s: %w", id, err)
	}

	var scanResult ScanResult
	if err := json.Unmarshal(resp.Response, &scanResult); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan result response: %w", err)
	}

	return &scanResult, nil
}

// Delete deletes the scan result with the given ID.
func (s *ScanResultService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/scanResult/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete scan result %s: %w", id, err)
	}

	return nil
}

// Copy copies the scan result with the given ID to the specified users.
func (s *ScanResultService) Copy(ctx context.Context, id string, input *ScanResultCopyInput) error {
	_, err := s.client.post(ctx, "/scanResult/"+id+"/copy", input)
	if err != nil {
		return fmt.Errorf("sc: copy scan result %s: %w", id, err)
	}

	return nil
}

// Email emails the scan result with the given ID.
func (s *ScanResultService) Email(ctx context.Context, id string, input *ScanResultEmailInput) error {
	_, err := s.client.post(ctx, "/scanResult/"+id+"/email", input)
	if err != nil {
		return fmt.Errorf("sc: email scan result %s: %w", id, err)
	}

	return nil
}

// Import imports a scan result from a file.
func (s *ScanResultService) Import(ctx context.Context, input *ScanResultImportInput) (*ScanResult, error) {
	resp, err := s.client.post(ctx, "/scanResult/import", input)
	if err != nil {
		return nil, fmt.Errorf("sc: import scan result: %w", err)
	}

	var scanResult ScanResult
	if err := json.Unmarshal(resp.Response, &scanResult); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan result response: %w", err)
	}

	return &scanResult, nil
}

// Reimport reimports the scan result with the given ID.
func (s *ScanResultService) Reimport(ctx context.Context, id string, input *ScanResultReimportInput) (*ScanResult, error) {
	resp, err := s.client.post(ctx, "/scanResult/"+id+"/import", input)
	if err != nil {
		return nil, fmt.Errorf("sc: reimport scan result %s: %w", id, err)
	}

	var scanResult ScanResult
	if err := json.Unmarshal(resp.Response, &scanResult); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan result response: %w", err)
	}

	return &scanResult, nil
}

// Stop stops the running scan result with the given ID.
func (s *ScanResultService) Stop(ctx context.Context, id string) (*ScanResult, error) {
	resp, err := s.client.post(ctx, "/scanResult/"+id+"/stop", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: stop scan result %s: %w", id, err)
	}

	var scanResult ScanResult
	if err := json.Unmarshal(resp.Response, &scanResult); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan result response: %w", err)
	}

	return &scanResult, nil
}

// Pause pauses the running scan result with the given ID.
func (s *ScanResultService) Pause(ctx context.Context, id string) (*ScanResult, error) {
	resp, err := s.client.post(ctx, "/scanResult/"+id+"/pause", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: pause scan result %s: %w", id, err)
	}

	var scanResult ScanResult
	if err := json.Unmarshal(resp.Response, &scanResult); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan result response: %w", err)
	}

	return &scanResult, nil
}

// Resume resumes the paused scan result with the given ID.
func (s *ScanResultService) Resume(ctx context.Context, id string) (*ScanResult, error) {
	resp, err := s.client.post(ctx, "/scanResult/"+id+"/resume", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: resume scan result %s: %w", id, err)
	}

	var scanResult ScanResult
	if err := json.Unmarshal(resp.Response, &scanResult); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan result response: %w", err)
	}

	return &scanResult, nil
}

// Download downloads the scan result with the given ID.
func (s *ScanResultService) Download(ctx context.Context, id string, input *ScanResultDownloadInput) error {
	_, err := s.client.post(ctx, "/scanResult/"+id+"/download", input)
	if err != nil {
		return fmt.Errorf("sc: download scan result %s: %w", id, err)
	}

	return nil
}
