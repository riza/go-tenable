package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// ExportsAssetsService handles communication with the Exports (Assets) related endpoints of the VM API.
type ExportsAssetsService struct {
	client *Client
}

type ExportsAssetsServiceExportAssetsV1RequestFilters struct {
	CreatedAt                 int64    `json:"created_at,omitempty"`
	UpdatedAt                 int64    `json:"updated_at,omitempty"`
	TerminatedAt              int64    `json:"terminated_at,omitempty"`
	IsTerminated              bool     `json:"is_terminated,omitempty"`
	DeletedAt                 int64    `json:"deleted_at,omitempty"`
	IsDeleted                 bool     `json:"is_deleted,omitempty"`
	IsLicensed                bool     `json:"is_licensed,omitempty"`
	FirstScanTime             int64    `json:"first_scan_time,omitempty"`
	LastAuthenticatedScanTime int64    `json:"last_authenticated_scan_time,omitempty"`
	LastAssessed              int64    `json:"last_assessed,omitempty"`
	LastScanId                string   `json:"last_scan_id,omitempty"`
	HasPluginResults          bool     `json:"has_plugin_results,omitempty"`
	NetworkId                 string   `json:"network_id,omitempty"`
	ServicenowSysid           bool     `json:"servicenow_sysid,omitempty"`
	Sources                   []string `json:"sources,omitempty"`
	TagCategory               []string `json:"tag.<category>,omitempty"`
}

type ExportsAssetsServiceExportAssetsV1Request struct {
	ChunkSize           int                                              `json:"chunk_size,omitempty"`
	IncludeOpenPorts    bool                                             `json:"include_open_ports,omitempty"`
	IncludeResourceTags bool                                             `json:"include_resource_tags,omitempty"`
	Filters             ExportsAssetsServiceExportAssetsV1RequestFilters `json:"filters,omitempty"`
}

type ExportsAssetsServiceExportAssetsV1Response struct {
	ExportUuid string `json:"export_uuid,omitempty"`
}

type ExportsAssetsServiceExportAssetsV2RequestFilters struct {
	CreatedAt                 int64    `json:"created_at,omitempty"`
	UpdatedAt                 int64    `json:"updated_at,omitempty"`
	TerminatedAt              int64    `json:"terminated_at,omitempty"`
	IsTerminated              bool     `json:"is_terminated,omitempty"`
	DeletedAt                 int64    `json:"deleted_at,omitempty"`
	IsDeleted                 bool     `json:"is_deleted,omitempty"`
	IsLicensed                bool     `json:"is_licensed,omitempty"`
	FirstScanTime             int64    `json:"first_scan_time,omitempty"`
	LastAuthenticatedScanTime int64    `json:"last_authenticated_scan_time,omitempty"`
	LastAssessed              int64    `json:"last_assessed,omitempty"`
	LastScanId                string   `json:"last_scan_id,omitempty"`
	HasPluginResults          bool     `json:"has_plugin_results,omitempty"`
	NetworkId                 string   `json:"network_id,omitempty"`
	ServicenowSysid           bool     `json:"servicenow_sysid,omitempty"`
	Since                     int64    `json:"since,omitempty"`
	Sources                   []string `json:"sources,omitempty"`
	Types                     []string `json:"types,omitempty"`
}

type ExportsAssetsServiceExportAssetsV2Request struct {
	ChunkSize           int                                              `json:"chunk_size,omitempty"`
	IncludeOpenPorts    bool                                             `json:"include_open_ports,omitempty"`
	IncludeResourceTags bool                                             `json:"include_resource_tags,omitempty"`
	Filters             ExportsAssetsServiceExportAssetsV2RequestFilters `json:"filters,omitempty"`
}

type ExportsAssetsServiceExportAssetsV2Response struct {
	ExportUuid string `json:"export_uuid,omitempty"`
}

type ExportsAssetsServiceExportsAssetsExportStatusResponse struct {
	Status          string `json:"status,omitempty"`
	ChunksAvailable []int  `json:"chunks_available,omitempty"`
}

type ExportsAssetsServiceExportsAssetsExportStatusRecentResponseExportsItemFilters struct {
}

type ExportsAssetsServiceExportsAssetsExportStatusRecentResponseExportsItem struct {
	Uuid              string                                                                        `json:"uuid,omitempty"`
	Status            string                                                                        `json:"status,omitempty"`
	TotalChunks       int                                                                           `json:"total_chunks,omitempty"`
	Filters           ExportsAssetsServiceExportsAssetsExportStatusRecentResponseExportsItemFilters `json:"filters,omitempty"`
	FinishedChunks    int                                                                           `json:"finished_chunks,omitempty"`
	NumAssetsPerChunk int                                                                           `json:"num_assets_per_chunk,omitempty"`
	Created           int                                                                           `json:"created,omitempty"`
}

type ExportsAssetsServiceExportsAssetsExportStatusRecentResponse struct {
	Exports []ExportsAssetsServiceExportsAssetsExportStatusRecentResponseExportsItem `json:"exports,omitempty"`
}

type ExportsAssetsServiceExportsAssetsExportCancelResponse struct {
	Status string `json:"status,omitempty"`
}

// ExportAssetsV1 - Export assets v1
func (s *ExportsAssetsService) ExportAssetsV1(ctx context.Context, req *ExportsAssetsServiceExportAssetsV1Request) (*ExportsAssetsServiceExportAssetsV1Response, error) {
	resp, err := s.client.post(ctx, "/assets/export", req)
	if err != nil {
		return nil, err
	}
	var result ExportsAssetsServiceExportAssetsV1Response
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportAssetsV2 - Export assets v2
func (s *ExportsAssetsService) ExportAssetsV2(ctx context.Context, req *ExportsAssetsServiceExportAssetsV2Request) (*ExportsAssetsServiceExportAssetsV2Response, error) {
	resp, err := s.client.post(ctx, "/assets/v2/export", req)
	if err != nil {
		return nil, err
	}
	var result ExportsAssetsServiceExportAssetsV2Response
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsAssetsExportStatus - Get assets export status
func (s *ExportsAssetsService) ExportsAssetsExportStatus(ctx context.Context, exportUuid string) (*ExportsAssetsServiceExportsAssetsExportStatusResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/assets/export/%s/status", exportUuid))
	if err != nil {
		return nil, err
	}
	var result ExportsAssetsServiceExportsAssetsExportStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsAssetsExportStatusRecent - List asset export jobs
func (s *ExportsAssetsService) ExportsAssetsExportStatusRecent(ctx context.Context) (*ExportsAssetsServiceExportsAssetsExportStatusRecentResponse, error) {
	resp, err := s.client.get(ctx, "/assets/export/status")
	if err != nil {
		return nil, err
	}
	var result ExportsAssetsServiceExportsAssetsExportStatusRecentResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsAssetsDownloadChunk - Download assets chunk
func (s *ExportsAssetsService) ExportsAssetsDownloadChunk(ctx context.Context, exportUuid string, chunkId string) error {
	_, err := s.client.get(ctx, fmt.Sprintf("/assets/export/%s/chunks/%s", exportUuid, chunkId))
	if err != nil {
		return err
	}
	return nil
}

// ExportsAssetsExportCancel - Cancel asset export
func (s *ExportsAssetsService) ExportsAssetsExportCancel(ctx context.Context, exportUuid string) (*ExportsAssetsServiceExportsAssetsExportCancelResponse, error) {
	resp, err := s.client.post(ctx, fmt.Sprintf("/assets/export/%s/cancel", exportUuid), nil)
	if err != nil {
		return nil, err
	}
	var result ExportsAssetsServiceExportsAssetsExportCancelResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
