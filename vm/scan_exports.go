package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// ScanExportsService handles communication with the Scan Exports related endpoints of the VM API.
type ScanExportsService struct {
	client *Client
}

type ScanExportsServiceExportRequestRequest struct {
	Format           string `json:"format,omitempty"`
	Password         string `json:"password,omitempty"`
	Chapters         string `json:"chapters,omitempty"`
	Filter0Filter    string `json:"filter.0.filter,omitempty"`
	Filter0Quality   string `json:"filter.0.quality,omitempty"`
	Filter0Value     string `json:"filter.0.value,omitempty"`
	FilterSearchType string `json:"filter.search_type,omitempty"`
	AssetId          string `json:"asset_id,omitempty"`
}

type ScanExportsServiceExportRequestResponse struct {
	File      string `json:"file,omitempty"`
	TempToken string `json:"temp_token,omitempty"`
}

type ScanExportsServiceExportStatusResponse struct {
	Status string `json:"status,omitempty"`
}

// ExportRequest - Export scan
func (s *ScanExportsService) ExportRequest(ctx context.Context, scanId string, req *ScanExportsServiceExportRequestRequest) (*ScanExportsServiceExportRequestResponse, error) {
	resp, err := s.client.post(ctx, fmt.Sprintf("/scans/%s/export", scanId), req)
	if err != nil {
		return nil, err
	}
	var result ScanExportsServiceExportRequestResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportStatus - Check scan export status
func (s *ScanExportsService) ExportStatus(ctx context.Context, scanId string, fileId string) (*ScanExportsServiceExportStatusResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scans/%s/export/%s/status", scanId, fileId))
	if err != nil {
		return nil, err
	}
	var result ScanExportsServiceExportStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportDownload - Download exported scan
func (s *ScanExportsService) ExportDownload(ctx context.Context, scanId string, fileId string) error {
	_, err := s.client.get(ctx, fmt.Sprintf("/scans/%s/export/%s/download", scanId, fileId))
	if err != nil {
		return err
	}
	return nil
}
