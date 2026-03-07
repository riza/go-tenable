package sc

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ScanService handles communication with the scan-related endpoints of the SC API.
type ScanService struct {
	client *Client
}

// Scan represents a scan resource from the API.
type Scan struct {
	ID                   string        `json:"id"`
	Name                 string        `json:"name"`
	Description          string        `json:"description"`
	Type                 string        `json:"type"`
	Status               string        `json:"status"`
	IPList               string        `json:"ipList"`
	DhcpTracking         string        `json:"dhcpTracking"`
	ClassifyMitigatedAge string        `json:"classifyMitigatedAge"`
	EmailOnLaunch        string        `json:"emailOnLaunch"`
	EmailOnFinish        string        `json:"emailOnFinish"`
	TimeoutAction        string        `json:"timeoutAction"`
	ScanningVirtualHosts string        `json:"scanningVirtualHosts"`
	RolloverType         string        `json:"rolloverType"`
	CreatedTime          string        `json:"createdTime"`
	ModifiedTime         string        `json:"modifiedTime"`
	MaxScanTime          string        `json:"maxScanTime"`
	InactivityTimeout    string        `json:"inactivityTimeout"`
	UUID                 string        `json:"uuid"`
	Policy               *IDRef        `json:"policy"`
	Repository           *IDRef        `json:"repository"`
	Zone                 *IDRef        `json:"zone"`
	OwnerGroup           *IDRef        `json:"ownerGroup"`
	Creator              *IDRef        `json:"creator"`
	Owner                *IDRef        `json:"owner"`
	Schedule             *ScanSchedule `json:"schedule"`
	Reports              []IDRef       `json:"reports"`
	Assets               []IDRef       `json:"assets"`
	Credentials          []IDRef       `json:"credentials"`
	NumDependents        string        `json:"numDependents"`
	CanUse               string        `json:"canUse"`
	CanManage            string        `json:"canManage"`
}

// ScanSchedule represents the schedule configuration for a scan.
type ScanSchedule struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Start       string `json:"start"`
	RepeatRule  string `json:"repeatRule"`
	Enabled     string `json:"enabled"`
	NextRun     int    `json:"nextRun"`
	ObjectType  string `json:"objectType"`
	DependentID string `json:"dependentID"`
}

// ScanListResponse holds the usable and manageable scan arrays returned by the list endpoint.
type ScanListResponse struct {
	Usable     []Scan `json:"usable"`
	Manageable []Scan `json:"manageable"`
}

// ScanCreateInput contains the fields for creating a new scan.
type ScanCreateInput struct {
	Name                 string        `json:"name"`
	Type                 string        `json:"type,omitempty"`
	Description          string        `json:"description,omitempty"`
	Repository           *IDRef        `json:"repository"`
	Zone                 *IDRef        `json:"zone,omitempty"`
	DhcpTracking         string        `json:"dhcpTracking,omitempty"`
	ClassifyMitigatedAge int           `json:"classifyMitigatedAge,omitempty"`
	Schedule             *ScanSchedule `json:"schedule,omitempty"`
	Reports              []IDRef       `json:"reports,omitempty"`
	Assets               []IDRef       `json:"assets,omitempty"`
	Credentials          []IDRef       `json:"credentials,omitempty"`
	EmailOnLaunch        string        `json:"emailOnLaunch,omitempty"`
	EmailOnFinish        string        `json:"emailOnFinish,omitempty"`
	TimeoutAction        string        `json:"timeoutAction,omitempty"`
	ScanningVirtualHosts string        `json:"scanningVirtualHosts,omitempty"`
	RolloverType         string        `json:"rolloverType,omitempty"`
	IPList               string        `json:"ipList,omitempty"`
	MaxScanTime          int           `json:"maxScanTime,omitempty"`
	InactivityTimeout    int           `json:"inactivityTimeout,omitempty"`
	Policy               *IDRef        `json:"policy,omitempty"`
}

// ScanUpdateInput contains the fields for updating an existing scan.
type ScanUpdateInput = ScanCreateInput

// ScanLaunchInput contains the optional diagnostic parameters for launching a scan.
type ScanLaunchInput struct {
	DiagnosticTarget   string `json:"diagnosticTarget,omitempty"`
	DiagnosticPassword string `json:"diagnosticPassword,omitempty"`
}

// ScanLaunchResponse holds the response from launching a scan.
type ScanLaunchResponse struct {
	ScanID     string `json:"scanID"`
	ScanResult struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
		JobID  string `json:"jobID"`
	} `json:"scanResult"`
}

// ScanCopyInput contains the fields for copying a scan.
type ScanCopyInput struct {
	Name       string `json:"name,omitempty"`
	TargetUser *IDRef `json:"targetUser,omitempty"`
}

// scanCopyResponse is an internal wrapper for the copy endpoint response
// which nests the scan under a "scan" key.
type scanCopyResponse struct {
	Scan Scan `json:"scan"`
}

// ScanListOptions specifies optional parameters for listing scans.
type ScanListOptions struct {
	Fields []string // API fields to include in the response.
}

// ScanAllFields is the list of all known scan fields that can be requested.
var ScanAllFields = []string{
	"id", "name", "description", "type", "status", "ipList",
	"dhcpTracking", "classifyMitigatedAge", "emailOnLaunch", "emailOnFinish",
	"timeoutAction", "scanningVirtualHosts", "rolloverType",
	"createdTime", "modifiedTime", "maxScanTime", "inactivityTimeout",
	"uuid", "policy", "repository", "zone", "ownerGroup", "creator", "owner",
	"schedule", "reports", "assets", "credentials",
	"numDependents", "canUse", "canManage",
}

// List returns the list of scans (usable and manageable).
// Pass nil for opts to use the API's default fields.
func (s *ScanService) List(opts *ScanListOptions) (*ScanListResponse, error) {
	var params QueryParams
	if opts != nil && len(opts.Fields) > 0 {
		params = QueryParams{"fields": strings.Join(opts.Fields, ",")}
	}

	resp, err := s.client.getWithParams("/scan", params)
	if err != nil {
		return nil, fmt.Errorf("sc: list scans: %w", err)
	}

	var result ScanListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan list response: %w", err)
	}

	return &result, nil
}

// Get returns the scan with the given ID.
func (s *ScanService) Get(id string) (*Scan, error) {
	resp, err := s.client.get("/scan/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get scan %s: %w", id, err)
	}

	var scan Scan
	if err := json.Unmarshal(resp.Response, &scan); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan response: %w", err)
	}

	return &scan, nil
}

// Create creates a new scan with the given input.
func (s *ScanService) Create(input *ScanCreateInput) (*Scan, error) {
	resp, err := s.client.post("/scan", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create scan: %w", err)
	}

	var scan Scan
	if err := json.Unmarshal(resp.Response, &scan); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan response: %w", err)
	}

	return &scan, nil
}

// Update updates an existing scan with the given input.
func (s *ScanService) Update(id string, input *ScanUpdateInput) (*Scan, error) {
	resp, err := s.client.patch("/scan/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update scan %s: %w", id, err)
	}

	var scan Scan
	if err := json.Unmarshal(resp.Response, &scan); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan response: %w", err)
	}

	return &scan, nil
}

// Delete deletes the scan with the given ID.
func (s *ScanService) Delete(id string) error {
	_, err := s.client.delete("/scan/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete scan %s: %w", id, err)
	}

	return nil
}

// Launch launches the scan with the given ID.
func (s *ScanService) Launch(id string, input *ScanLaunchInput) (*ScanLaunchResponse, error) {
	resp, err := s.client.post("/scan/"+id+"/launch", input)
	if err != nil {
		return nil, fmt.Errorf("sc: launch scan %s: %w", id, err)
	}

	var result ScanLaunchResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan launch response: %w", err)
	}

	return &result, nil
}

// Copy copies the scan with the given ID.
func (s *ScanService) Copy(id string, input *ScanCopyInput) (*Scan, error) {
	resp, err := s.client.post("/scan/"+id+"/copy", input)
	if err != nil {
		return nil, fmt.Errorf("sc: copy scan %s: %w", id, err)
	}

	var wrapper scanCopyResponse
	if err := json.Unmarshal(resp.Response, &wrapper); err != nil {
		return nil, fmt.Errorf("sc: unmarshal scan copy response: %w", err)
	}

	return &wrapper.Scan, nil
}
