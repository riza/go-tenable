package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// SystemService handles communication with the system-related endpoints of the SC API.
type SystemService struct {
	client *Client
}

// SystemInfo represents the system information returned by GET /system.
type SystemInfo struct {
	Version                          string          `json:"version"`
	BuildID                          string          `json:"buildID"`
	ReleaseID                        string          `json:"releaseID"`
	Banner                           string          `json:"banner"`
	UUID                             string          `json:"uuid"`
	UUIDSHA256                       string          `json:"uuidsha256"`
	Logo                             string          `json:"logo"`
	ServerAuth                       string          `json:"serverAuth"`
	ServerClassification             string          `json:"serverClassification"`
	SessionTimeout                   string          `json:"sessionTimeout"`
	LicenseStatus                    string          `json:"licenseStatus"`
	ACAS                             string          `json:"ACAS"`
	FreshInstall                     string          `json:"freshInstall"`
	HeaderText                       string          `json:"headerText"`
	RiskRuleCommentsEnabled          string          `json:"riskRuleCommentsEnabled"`
	TelemetryEnabled                 string          `json:"telemetryEnabled"`
	SerializationDisabled            string          `json:"SerializationDisabled"`
	ActiveIPs                        string          `json:"activeIPs"`
	LicensedIPs                      string          `json:"licensedIPs"`
	LicenseExpiration                string          `json:"licenseExpiration"`
	PostgresConnStatus               string          `json:"postgresConnStatus"`
	PostgresConnectionType           string          `json:"postgresConnectionType"`
	VulnerabilityIntelligenceEnabled string          `json:"VulnerabilityIntelligenceEnabled"`
	WASLicense                       string          `json:"wasLicense"`
	WASFQDNCount                     string          `json:"wasFQDNCount"`
	LoginNotifications               json.RawMessage `json:"loginNotifications"`
	PasswordComplexity               json.RawMessage `json:"PasswordComplexity"`
	SAML                             json.RawMessage `json:"SAML"`
	ReportTypes                      []ReportType    `json:"reportTypes"`
	Timezones                        []Timezone      `json:"timezones"`
}

// ReportType represents a report type entry in the system info.
type ReportType struct {
	Name          string          `json:"name"`
	Type          string          `json:"type"`
	Enabled       string          `json:"enabled"`
	AttributeSets json.RawMessage `json:"attributeSets"`
}

// Timezone represents a timezone entry in the system info.
type Timezone struct {
	Name      string `json:"name"`
	GMTOffset string `json:"GMTOffset"`
}

// DebugItem represents a single debug setting item.
type DebugItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	SizeWarning string `json:"sizeWarning"`
	Enabled     string `json:"enabled"`
}

// DebugUpdateInput contains the fields for updating debug settings.
type DebugUpdateInput struct {
	Modules []DebugItem `json:"modules,omitempty"`
}

// SystemDiagnostics represents the system diagnostics information.
type SystemDiagnostics struct {
	StatusJava          string `json:"statusJava"`
	StatusRPM           string `json:"statusRPM"`
	StatusDisk          string `json:"statusDisk"`
	StatusThresholdDisk string `json:"statusThresholdDisk"`
	StatusLastChecked   string `json:"statusLastChecked"`
}

// DiagnosticsGenerateInput contains the fields for generating diagnostics.
type DiagnosticsGenerateInput struct {
	Type string `json:"type,omitempty"`
}

// DebugLogsGenerateInput contains the fields for generating debug logs.
type DebugLogsGenerateInput struct {
	Type string `json:"type,omitempty"`
}

// Locale represents a locale configuration.
type Locale struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// LocaleUpdateInput contains the fields for updating the locale.
type LocaleUpdateInput struct {
	Code string `json:"code"`
}

// FIPSStatus represents the FIPS status information.
type FIPSStatus struct {
	ApacheStatus     string `json:"apacheStatus"`
	FipsMode         string `json:"fipsMode"`
	ApacheStatusText string `json:"apacheStatusText"`
}

// FIPSSetInput contains the fields for setting FIPS mode.
type FIPSSetInput struct {
	FipsMode string `json:"fipsMode"`
}

// LogEntry represents a single log entry from a log search.
type LogEntry struct {
	RawLog       string `json:"rawLog"`
	Message      string `json:"message"`
	Organization string `json:"organization"`
	Severity     string `json:"severity"`
	Module       string `json:"module"`
	Source       string `json:"source"`
	Date         string `json:"date"`
}

// LogSearchResponse holds the response from searching logs.
type LogSearchResponse struct {
	TotalRecords int        `json:"totalRecords"`
	Results      []LogEntry `json:"results"`
}

// LogSearchInput contains the fields for searching logs.
type LogSearchInput struct {
	Module   string `json:"module,omitempty"`
	Severity string `json:"severity,omitempty"`
	Date     string `json:"date,omitempty"`
	Keywords string `json:"keywords,omitempty"`
}

// LogDownloadInput contains the fields for downloading logs.
type LogDownloadInput struct {
	Module   string `json:"module,omitempty"`
	Severity string `json:"severity,omitempty"`
	Date     string `json:"date,omitempty"`
	Keywords string `json:"keywords,omitempty"`
}

// Get returns the system information.
func (s *SystemService) Get(ctx context.Context) (*SystemInfo, error) {
	resp, err := s.client.get(ctx, "/system")
	if err != nil {
		return nil, fmt.Errorf("sc: get system: %w", err)
	}

	var result SystemInfo
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal system response: %w", err)
	}

	return &result, nil
}

// GetDebug returns the list of debug settings.
func (s *SystemService) GetDebug(ctx context.Context) ([]DebugItem, error) {
	resp, err := s.client.get(ctx, "/system/debug")
	if err != nil {
		return nil, fmt.Errorf("sc: get system debug: %w", err)
	}

	var result []DebugItem
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal system debug response: %w", err)
	}

	return result, nil
}

// UpdateDebug updates the debug settings.
func (s *SystemService) UpdateDebug(ctx context.Context, input *DebugUpdateInput) error {
	_, err := s.client.patch(ctx, "/system/debug", input)
	if err != nil {
		return fmt.Errorf("sc: update system debug: %w", err)
	}

	return nil
}

// GetDiagnostics returns the system diagnostics information.
func (s *SystemService) GetDiagnostics(ctx context.Context) (*SystemDiagnostics, error) {
	resp, err := s.client.get(ctx, "/system/diagnostics")
	if err != nil {
		return nil, fmt.Errorf("sc: get system diagnostics: %w", err)
	}

	var result SystemDiagnostics
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal system diagnostics response: %w", err)
	}

	return &result, nil
}

// GenerateDiagnostics starts generating system diagnostics.
func (s *SystemService) GenerateDiagnostics(ctx context.Context, input *DiagnosticsGenerateInput) error {
	_, err := s.client.post(ctx, "/system/diagnostics/generate", input)
	if err != nil {
		return fmt.Errorf("sc: generate system diagnostics: %w", err)
	}

	return nil
}

// DownloadDiagnostics downloads the system diagnostics file.
func (s *SystemService) DownloadDiagnostics(ctx context.Context) error {
	_, err := s.client.post(ctx, "/system/diagnostics/download", nil)
	if err != nil {
		return fmt.Errorf("sc: download system diagnostics: %w", err)
	}

	return nil
}

// GenerateDebugLogs starts generating debug logs.
func (s *SystemService) GenerateDebugLogs(ctx context.Context, input *DebugLogsGenerateInput) error {
	_, err := s.client.post(ctx, "/system/debuglogs/generate", input)
	if err != nil {
		return fmt.Errorf("sc: generate system debug logs: %w", err)
	}

	return nil
}

// DownloadDebugLogs downloads the debug logs file.
func (s *SystemService) DownloadDebugLogs(ctx context.Context) error {
	_, err := s.client.post(ctx, "/system/debuglogs/download", nil)
	if err != nil {
		return fmt.Errorf("sc: download system debug logs: %w", err)
	}

	return nil
}

// GetLocale returns the current locale configuration.
func (s *SystemService) GetLocale(ctx context.Context) (*Locale, error) {
	resp, err := s.client.get(ctx, "/system/locale")
	if err != nil {
		return nil, fmt.Errorf("sc: get system locale: %w", err)
	}

	var result Locale
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal system locale response: %w", err)
	}

	return &result, nil
}

// UpdateLocale updates the system locale.
func (s *SystemService) UpdateLocale(ctx context.Context, input *LocaleUpdateInput) error {
	_, err := s.client.patch(ctx, "/system/locale", input)
	if err != nil {
		return fmt.Errorf("sc: update system locale: %w", err)
	}

	return nil
}

// GetLocales returns the available locales as a raw JSON map.
func (s *SystemService) GetLocales(ctx context.Context) (json.RawMessage, error) {
	resp, err := s.client.get(ctx, "/system/locales")
	if err != nil {
		return nil, fmt.Errorf("sc: get system locales: %w", err)
	}

	return resp.Response, nil
}

// GetLogFiles returns the list of log files as raw JSON.
func (s *SystemService) GetLogFiles(ctx context.Context) (json.RawMessage, error) {
	resp, err := s.client.get(ctx, "/system/logFiles")
	if err != nil {
		return nil, fmt.Errorf("sc: get system log files: %w", err)
	}

	return resp.Response, nil
}

// SearchLogs searches the system logs with the given criteria.
func (s *SystemService) SearchLogs(ctx context.Context, input *LogSearchInput) (*LogSearchResponse, error) {
	resp, err := s.client.post(ctx, "/system/logs", input)
	if err != nil {
		return nil, fmt.Errorf("sc: search system logs: %w", err)
	}

	var result LogSearchResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal system logs response: %w", err)
	}

	return &result, nil
}

// DownloadLogs downloads system logs matching the given criteria.
func (s *SystemService) DownloadLogs(ctx context.Context, input *LogDownloadInput) error {
	_, err := s.client.post(ctx, "/system/logs/download", input)
	if err != nil {
		return fmt.Errorf("sc: download system logs: %w", err)
	}

	return nil
}

// GetLogModules returns the available log modules as raw JSON.
func (s *SystemService) GetLogModules(ctx context.Context) (json.RawMessage, error) {
	resp, err := s.client.get(ctx, "/system/logs/modules")
	if err != nil {
		return nil, fmt.Errorf("sc: get system log modules: %w", err)
	}

	return resp.Response, nil
}

// GetFIPS returns the current FIPS status.
func (s *SystemService) GetFIPS(ctx context.Context) (*FIPSStatus, error) {
	resp, err := s.client.get(ctx, "/system/fips")
	if err != nil {
		return nil, fmt.Errorf("sc: get system fips: %w", err)
	}

	var result FIPSStatus
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal system fips response: %w", err)
	}

	return &result, nil
}

// SetFIPS sets the FIPS mode and returns the updated status.
func (s *SystemService) SetFIPS(ctx context.Context, input *FIPSSetInput) (*FIPSStatus, error) {
	resp, err := s.client.post(ctx, "/system/fips", input)
	if err != nil {
		return nil, fmt.Errorf("sc: set system fips: %w", err)
	}

	var result FIPSStatus
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal system fips response: %w", err)
	}

	return &result, nil
}
