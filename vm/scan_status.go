package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// ScanStatusService handles communication with the Scan Status related endpoints of the VM API.
type ScanStatusService struct {
	client *Client
}

type ScanStatusServiceGetLatestStatusResponse struct {
	Status string `json:"status,omitempty"`
}

type ScanStatusServiceReadStatusRequest struct {
	Read bool `json:"read,omitempty"`
}

type ScanStatusServiceVmScansProgressGetResponse struct {
	Progress int `json:"progress,omitempty"`
}

// GetLatestStatus - Get latest scan status
func (s *ScanStatusService) GetLatestStatus(ctx context.Context, scanId string) (*ScanStatusServiceGetLatestStatusResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scans/%s/latest-status", scanId))
	if err != nil {
		return nil, err
	}
	var result ScanStatusServiceGetLatestStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReadStatus - Update scan status
func (s *ScanStatusService) ReadStatus(ctx context.Context, scanId string, req *ScanStatusServiceReadStatusRequest) error {
	_, err := s.client.put(ctx, fmt.Sprintf("/scans/%s/status", scanId), req)
	if err != nil {
		return err
	}
	return nil
}

// VmScansProgressGet - Get scan progress
func (s *ScanStatusService) VmScansProgressGet(ctx context.Context, scanId string) (*ScanStatusServiceVmScansProgressGetResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scans/%s/progress", scanId))
	if err != nil {
		return nil, err
	}
	var result ScanStatusServiceVmScansProgressGetResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
