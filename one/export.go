package one

import (
	"context"
	"encoding/json"
)

// ExportService handles communication with the Export related endpoints of the Tenable One API.
type ExportService struct {
	client *Client
}

// InventoryExportRequest represents the request body for creating an inventory export.
type InventoryExportRequest struct {
	Format   string        `json:"format,omitempty"`
	Limit    int           `json:"limit,omitempty"`
	Filters  interface{}   `json:"filters,omitempty"`
	Columns  []string      `json:"columns,omitempty"`
}

// InventoryExportResponse represents the response from creating an inventory export.
type InventoryExportResponse struct {
	ExportId     string `json:"export_id,omitempty"`
	Format       string `json:"format,omitempty"`
	Status       string `json:"status,omitempty"`
	TotalObjects int    `json:"total_objects,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	ExpiresAt    string `json:"expires_at,omitempty"`
}

// ExportAssets creates an export of assets.
func (s *ExportService) ExportAssets(ctx context.Context, req *InventoryExportRequest) (*InventoryExportResponse, error) {
	resp, err := s.client.post(ctx, "/api/v1/t1/inventory/export/assets", req)
	if err != nil {
		return nil, err
	}

	var result InventoryExportResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ExportFindings creates an export of findings.
func (s *ExportService) ExportFindings(ctx context.Context, req *InventoryExportRequest) (*InventoryExportResponse, error) {
	resp, err := s.client.post(ctx, "/api/v1/t1/inventory/export/findings", req)
	if err != nil {
		return nil, err
	}

	var result InventoryExportResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ExportStatusResponse represents the status of an export.
type ExportStatusResponse struct {
	ExportId     string `json:"export_id,omitempty"`
	Status       string `json:"status,omitempty"`
	TotalObjects int    `json:"total_objects,omitempty"`
	ReadyObjects int    `json:"ready_objects,omitempty"`
	Chunks       []int  `json:"chunks,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
}

// GetAssetsExportStatus returns the status of an assets export.
func (s *ExportService) GetAssetsExportStatus(ctx context.Context, exportId string) (*ExportStatusResponse, error) {
	resp, err := s.client.get(ctx, "/api/v1/t1/inventory/export/assets/status")
	if err != nil {
		return nil, err
	}

	var result ExportStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetFindingsExportStatus returns the status of a findings export.
func (s *ExportService) GetFindingsExportStatus(ctx context.Context, exportId string) (*ExportStatusResponse, error) {
	resp, err := s.client.get(ctx, "/api/v1/t1/inventory/export/findings/status")
	if err != nil {
		return nil, err
	}

	var result ExportStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetExportStatus returns the status of an export by ID.
func (s *ExportService) GetExportStatus(ctx context.Context, exportId string) (*ExportStatusResponse, error) {
	resp, err := s.client.get(ctx, "/api/v1/t1/inventory/export/"+exportId+"/status")
	if err != nil {
		return nil, err
	}

	var result ExportStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DownloadExportChunk downloads a chunk of an export.
func (s *ExportService) DownloadExportChunk(ctx context.Context, exportId string, chunkId int) ([]byte, error) {
	return s.client.get(ctx, "/api/v1/t1/inventory/export/"+exportId+"/download/"+string(rune(chunkId+'0')))
}
