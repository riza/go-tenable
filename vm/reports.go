package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// ReportsService handles communication with the Reports related endpoints of the VM API.
type ReportsService struct {
	client *Client
}

type ReportsServiceVmReportsCreateRequestFiltersItem struct {
	Property string      `json:"property,omitempty"`
	Operator string      `json:"operator,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}

type ReportsServiceVmReportsCreateRequest struct {
	Name         string                                            `json:"name,omitempty"`
	TemplateName string                                            `json:"template_name,omitempty"`
	Filters      []ReportsServiceVmReportsCreateRequestFiltersItem `json:"filters,omitempty"`
}

type ReportsServiceVmReportsCreateResponse struct {
	Uuid string `json:"uuid,omitempty"`
}

type ReportsServiceVmReportsStatusResponse struct {
	Status string `json:"status,omitempty"`
}

// VmReportsCreate - Create report
func (s *ReportsService) VmReportsCreate(ctx context.Context, req *ReportsServiceVmReportsCreateRequest) (*ReportsServiceVmReportsCreateResponse, error) {
	resp, err := s.client.post(ctx, "/reports/export", req)
	if err != nil {
		return nil, err
	}
	var result ReportsServiceVmReportsCreateResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// VmReportsStatus - Get report status
func (s *ReportsService) VmReportsStatus(ctx context.Context, reportUuid string) (*ReportsServiceVmReportsStatusResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/reports/export/%s/status", reportUuid))
	if err != nil {
		return nil, err
	}
	var result ReportsServiceVmReportsStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// VmReportsDownload - Download report
func (s *ReportsService) VmReportsDownload(ctx context.Context, reportUuid string) error {
	_, err := s.client.get(ctx, fmt.Sprintf("/reports/export/%s/download", reportUuid))
	if err != nil {
		return err
	}
	return nil
}
