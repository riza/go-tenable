package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DeviceInfoService handles deviceInfo operations.
type DeviceInfoService struct {
	client *Client
}

// DeviceInfo represents a deviceInfo resource.
type DeviceInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DeviceInfoListResponse represents the response from listing deviceInfos.
type DeviceInfoListResponse struct {
	Usable     []DeviceInfo `json:"usable"`
	Manageable []DeviceInfo `json:"manageable"`
}

// List returns all deviceInfos.
func (s *DeviceInfoService) List(ctx context.Context) (*DeviceInfoListResponse, error) {
	resp, err := s.client.get(ctx, "/deviceInfo")
	if err != nil {
		return nil, fmt.Errorf("sc: list deviceInfos: %w", err)
	}

	var result DeviceInfoListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal deviceInfo list response: %w", err)
	}

	return &result, nil
}
