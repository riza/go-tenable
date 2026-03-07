package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// ExportsVulnerabilitiesService handles communication with the Exports (Vulnerabilities) related endpoints of the VM API.
type ExportsVulnerabilitiesService struct {
	client *Client
}

type ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFiltersTimeTakenToFix struct {
	Gte int64 `json:"gte,omitempty"`
	Lte int64 `json:"lte,omitempty"`
}

type ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFiltersCvss4BaseScore struct {
	Eq  []float64 `json:"eq,omitempty"`
	Neq []float64 `json:"neq,omitempty"`
	Gt  float64   `json:"gt,omitempty"`
	Gte float64   `json:"gte,omitempty"`
	Lt  float64   `json:"lt,omitempty"`
	Lte float64   `json:"lte,omitempty"`
}

type ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFiltersEpssScore struct {
	Eq  []float64 `json:"eq,omitempty"`
	Neq []float64 `json:"neq,omitempty"`
	Gt  float64   `json:"gt,omitempty"`
	Gte float64   `json:"gte,omitempty"`
	Lt  float64   `json:"lt,omitempty"`
	Lte float64   `json:"lte,omitempty"`
}

type ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFiltersVprScore struct {
	Eq  []float64 `json:"eq,omitempty"`
	Neq []float64 `json:"neq,omitempty"`
	Gt  float64   `json:"gt,omitempty"`
	Gte float64   `json:"gte,omitempty"`
	Lt  float64   `json:"lt,omitempty"`
	Lte float64   `json:"lte,omitempty"`
}

type ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFiltersVprV2Score struct {
	Eq  []float64 `json:"eq,omitempty"`
	Neq []float64 `json:"neq,omitempty"`
	Gt  float64   `json:"gt,omitempty"`
	Gte float64   `json:"gte,omitempty"`
	Lt  float64   `json:"lt,omitempty"`
	Lte float64   `json:"lte,omitempty"`
}

type ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFilters struct {
	Since                    int64                                                                              `json:"since,omitempty"`
	FirstFound               int64                                                                              `json:"first_found,omitempty"`
	FirstSeen                int64                                                                              `json:"first_seen,omitempty"`
	LastFound                int64                                                                              `json:"last_found,omitempty"`
	LastSeen                 int64                                                                              `json:"last_seen,omitempty"`
	IndexedAt                int64                                                                              `json:"indexed_at,omitempty"`
	LastFixed                int64                                                                              `json:"last_fixed,omitempty"`
	TimeTakenToFix           ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFiltersTimeTakenToFix `json:"time_taken_to_fix,omitempty"`
	ResurfacedDate           int64                                                                              `json:"resurfaced_date,omitempty"`
	CidrRange                string                                                                             `json:"cidr_range,omitempty"`
	CveId                    []string                                                                           `json:"cve_id,omitempty"`
	CveCategory              []string                                                                           `json:"cve_category,omitempty"`
	Cvss4BaseScore           ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFiltersCvss4BaseScore `json:"cvss4_base_score,omitempty"`
	EpssScore                ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFiltersEpssScore      `json:"epss_score,omitempty"`
	ExploitMaturity          []string                                                                           `json:"exploit_maturity,omitempty"`
	InitiativeId             string                                                                             `json:"initiative_id,omitempty"`
	NetworkId                string                                                                             `json:"network_id,omitempty"`
	PluginFamily             []string                                                                           `json:"plugin_family,omitempty"`
	PluginId                 []int                                                                              `json:"plugin_id,omitempty"`
	PluginType               string                                                                             `json:"plugin_type,omitempty"`
	ScanUuid                 string                                                                             `json:"scan_uuid,omitempty"`
	Severity                 []string                                                                           `json:"severity,omitempty"`
	SeverityModificationType []string                                                                           `json:"severity_modification_type,omitempty"`
	State                    []string                                                                           `json:"state,omitempty"`
	Source                   []string                                                                           `json:"source,omitempty"`
	TagCategory              []string                                                                           `json:"tag.<category>,omitempty"`
	VprScore                 ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFiltersVprScore       `json:"vpr_score,omitempty"`
	VprV2Score               ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFiltersVprV2Score     `json:"vpr_v2_score,omitempty"`
	VprThreatIntensity       []string                                                                           `json:"vpr_threat_intensity,omitempty"`
	Weaponization            []string                                                                           `json:"weaponization,omitempty"`
	ZeroDay                  bool                                                                               `json:"zero_day,omitempty"`
}

type ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequest struct {
	NumAssets         int                                                                  `json:"num_assets,omitempty"`
	IncludeUnlicensed bool                                                                 `json:"include_unlicensed,omitempty"`
	Filters           ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequestFilters `json:"filters,omitempty"`
}

type ExportsVulnerabilitiesServiceExportsVulnsRequestExportResponse struct {
	ExportUuid string `json:"export_uuid,omitempty"`
}

type ExportsVulnerabilitiesServiceExportsVulnsExportStatusResponseFilters struct {
}

type ExportsVulnerabilitiesServiceExportsVulnsExportStatusResponse struct {
	Uuid                 string                                                               `json:"uuid,omitempty"`
	Status               string                                                               `json:"status,omitempty"`
	ChunksAvailable      []int                                                                `json:"chunks_available,omitempty"`
	ChunksFailed         []int                                                                `json:"chunks_failed,omitempty"`
	ChunksCancelled      []int                                                                `json:"chunks_cancelled,omitempty"`
	TotalChunks          int                                                                  `json:"total_chunks,omitempty"`
	ChunksAvailableCount int                                                                  `json:"chunks_available_count,omitempty"`
	EmptyChunksCount     int                                                                  `json:"empty_chunks_count,omitempty"`
	FinishedChunks       int                                                                  `json:"finished_chunks,omitempty"`
	Filters              ExportsVulnerabilitiesServiceExportsVulnsExportStatusResponseFilters `json:"filters,omitempty"`
	NumAssetsPerChunk    int                                                                  `json:"num_assets_per_chunk,omitempty"`
	Created              int                                                                  `json:"created,omitempty"`
}

type ExportsVulnerabilitiesServiceExportsVulnsExportStatusRecentResponseExportsItemFilters struct {
}

type ExportsVulnerabilitiesServiceExportsVulnsExportStatusRecentResponseExportsItem struct {
	Uuid                 string                                                                                `json:"uuid,omitempty"`
	Status               string                                                                                `json:"status,omitempty"`
	TotalChunks          int                                                                                   `json:"total_chunks,omitempty"`
	ChunksAvailableCount int                                                                                   `json:"chunks_available_count,omitempty"`
	EmptyChunksCount     int                                                                                   `json:"empty_chunks_count,omitempty"`
	FinishedChunks       int                                                                                   `json:"finished_chunks,omitempty"`
	Filters              ExportsVulnerabilitiesServiceExportsVulnsExportStatusRecentResponseExportsItemFilters `json:"filters,omitempty"`
	NumAssetsPerChunk    int                                                                                   `json:"num_assets_per_chunk,omitempty"`
	Created              int                                                                                   `json:"created,omitempty"`
}

type ExportsVulnerabilitiesServiceExportsVulnsExportStatusRecentResponse struct {
	Exports []ExportsVulnerabilitiesServiceExportsVulnsExportStatusRecentResponseExportsItem `json:"exports,omitempty"`
}

type ExportsVulnerabilitiesServiceExportsVulnsExportCancelResponse struct {
	Status string `json:"status,omitempty"`
}

// ExportsVulnsRequestExport - Export vulnerabilities
func (s *ExportsVulnerabilitiesService) ExportsVulnsRequestExport(ctx context.Context, req *ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequest) (*ExportsVulnerabilitiesServiceExportsVulnsRequestExportResponse, error) {
	resp, err := s.client.post(ctx, "/vulns/export", req)
	if err != nil {
		return nil, err
	}
	var result ExportsVulnerabilitiesServiceExportsVulnsRequestExportResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsVulnsExportStatus - Get vulnerability export status
func (s *ExportsVulnerabilitiesService) ExportsVulnsExportStatus(ctx context.Context, exportUuid string) (*ExportsVulnerabilitiesServiceExportsVulnsExportStatusResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/vulns/export/%s/status", exportUuid))
	if err != nil {
		return nil, err
	}
	var result ExportsVulnerabilitiesServiceExportsVulnsExportStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsVulnsExportStatusRecent - List vuln export jobs
func (s *ExportsVulnerabilitiesService) ExportsVulnsExportStatusRecent(ctx context.Context) (*ExportsVulnerabilitiesServiceExportsVulnsExportStatusRecentResponse, error) {
	resp, err := s.client.get(ctx, "/vulns/export/status")
	if err != nil {
		return nil, err
	}
	var result ExportsVulnerabilitiesServiceExportsVulnsExportStatusRecentResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsVulnsDownloadChunk - Download vulnerabilities chunk
func (s *ExportsVulnerabilitiesService) ExportsVulnsDownloadChunk(ctx context.Context, exportUuid string, chunkId string) error {
	_, err := s.client.get(ctx, fmt.Sprintf("/vulns/export/%s/chunks/%s", exportUuid, chunkId))
	if err != nil {
		return err
	}
	return nil
}

// ExportsVulnsExportCancel - Cancel vuln export
func (s *ExportsVulnerabilitiesService) ExportsVulnsExportCancel(ctx context.Context, exportUuid string) (*ExportsVulnerabilitiesServiceExportsVulnsExportCancelResponse, error) {
	resp, err := s.client.post(ctx, fmt.Sprintf("/vulns/export/%s/cancel", exportUuid), nil)
	if err != nil {
		return nil, err
	}
	var result ExportsVulnerabilitiesServiceExportsVulnsExportCancelResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
