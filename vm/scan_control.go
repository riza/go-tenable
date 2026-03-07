package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// ScanControlService handles communication with the Scan Control related endpoints of the VM API.
type ScanControlService struct {
	client *Client
}

type ScanControlServiceLaunchRequest struct {
	AltTargets []string `json:"alt_targets,omitempty"`
	Rollover   bool     `json:"rollover,omitempty"`
}

type ScanControlServiceLaunchResponse struct {
	ScanUuid string `json:"scan_uuid,omitempty"`
}

// Launch - Launch scan
func (s *ScanControlService) Launch(ctx context.Context, scanId string, req *ScanControlServiceLaunchRequest) (*ScanControlServiceLaunchResponse, error) {
	resp, err := s.client.post(ctx, fmt.Sprintf("/scans/%s/launch", scanId), req)
	if err != nil {
		return nil, err
	}
	var result ScanControlServiceLaunchResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Pause - Pause scan
func (s *ScanControlService) Pause(ctx context.Context, scanId string) error {
	_, err := s.client.post(ctx, fmt.Sprintf("/scans/%s/pause", scanId), nil)
	if err != nil {
		return err
	}
	return nil
}

// Resume - Resume scan
func (s *ScanControlService) Resume(ctx context.Context, scanId string) error {
	_, err := s.client.post(ctx, fmt.Sprintf("/scans/%s/resume", scanId), nil)
	if err != nil {
		return err
	}
	return nil
}

// Stop - Stop scan
func (s *ScanControlService) Stop(ctx context.Context, scanId string) error {
	_, err := s.client.post(ctx, fmt.Sprintf("/scans/%s/stop", scanId), nil)
	if err != nil {
		return err
	}
	return nil
}

// VmScansStopForce - Force stop scan
func (s *ScanControlService) VmScansStopForce(ctx context.Context, scheduleUuid string) error {
	_, err := s.client.post(ctx, fmt.Sprintf("/scans/%s/force-stop", scheduleUuid), nil)
	if err != nil {
		return err
	}
	return nil
}
