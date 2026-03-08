package platform

import (
	"context"
	"encoding/json"
	"fmt"
)

// ScannersService handles communication with the Scanners related endpoints of the Platform API.
type ScannersService struct {
	client *Client
}

// Scanner represents a scanner.
type Scanner struct {
	Id            int                    `json:"id,omitempty"`
	Uuid          string                 `json:"uuid,omitempty"`
	Name          string                 `json:"name,omitempty"`
	IpAddresses   []string               `json:"ip_addresses,omitempty"`
	Hostname      string                 `json:"hostname,omitempty"`
	Status        string                 `json:"status,omitempty"`
	Platform      string                 `json:"platform,omitempty"`
	Distro        string                 `json:"distro,omitempty"`
	EngineVersion string                 `json:"engine_version,omitempty"`
	Group         bool                   `json:"group,omitempty"`
	Groups        []ScannerGroupRef      `json:"groups,omitempty"`
	License       map[string]interface{} `json:"license,omitempty"`
	CreationDate  int64                  `json:"creation_date,omitempty"`
	LastSeen      int64                  `json:"last_scan,omitempty"`
	Settings      map[string]interface{} `json:"settings,omitempty"`
}

// ScannerGroupRef represents a reference to a scanner group.
type ScannerGroupRef struct {
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

// ScannerGroup represents a scanner group (scanner pool).
type ScannerGroup struct {
	Id                   int    `json:"id,omitempty"`
	Uuid                 string `json:"uuid,omitempty"`
	Name                 string `json:"name,omitempty"`
	Description          string `json:"description,omitempty"`
	CreationDate         int64  `json:"creation_date,omitempty"`
	LastModificationDate int64  `json:"last_modification_date,omitempty"`
	OwnerId              int    `json:"owner_id,omitempty"`
	Owner                string `json:"owner,omitempty"`
	OwnerName            string `json:"owner_name,omitempty"`
	DefaultPermissions   int    `json:"default_permissions,omitempty"`
	UserPermissions      int    `json:"user_permissions,omitempty"`
	Shared               int    `json:"shared,omitempty"`
	ScanCount            int    `json:"scan_count,omitempty"`
	ScannerCount         string `json:"scanner_count,omitempty"`
	Token                string `json:"token,omitempty"`
	Flag                 string `json:"flag,omitempty"`
	Type                 string `json:"type,omitempty"`
	NetworkName          string `json:"network_name,omitempty"`
	SupportsWebapp       bool   `json:"supports_webapp,omitempty"`
	ScannerId            int    `json:"scanner_id,omitempty"`
	ScannerUuid          string `json:"scanner_uuid,omitempty"`
}

// ScannersListResponse represents the response from listing scanners.
type ScannersListResponse struct {
	Scanners []Scanner `json:"scanners,omitempty"`
	Total    int       `json:"total,omitempty"`
}

// ListScanners returns all scanners.
func (s *ScannersService) ListScanners(ctx context.Context) (*ScannersListResponse, error) {
	resp, err := s.client.get(ctx, "/scanners")
	if err != nil {
		return nil, err
	}

	var result ScannersListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetScanner returns a specific scanner by ID.
func (s *ScannersService) GetScanner(ctx context.Context, scannerId int) (*Scanner, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scanners/%d", scannerId))
	if err != nil {
		return nil, err
	}

	var result Scanner
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ScannerConfig represents scanner configuration.
type ScannerConfig struct {
	Settings map[string]interface{} `json:"settings,omitempty"`
}

// GetScannerConfig returns the configuration for a scanner.
func (s *ScannersService) GetScannerConfig(ctx context.Context, scannerId int) (*ScannerConfig, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scanners/%d/config", scannerId))
	if err != nil {
		return nil, err
	}

	var result ScannerConfig
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateScannerConfig updates the configuration for a scanner.
func (s *ScannersService) UpdateScannerConfig(ctx context.Context, scannerId int, req *ScannerConfig) (*ScannerConfig, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/scanners/%d/config", scannerId), req)
	if err != nil {
		return nil, err
	}

	var result ScannerConfig
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ScannerTask represents a scanner task.
type ScannerTask struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Status      string `json:"status,omitempty"`
	Progress    int    `json:"progress,omitempty"`
	StartedAt   string `json:"started_at,omitempty"`
	CompletedAt string `json:"completed_at,omitempty"`
}

// GetScannerTasks returns tasks for a scanner.
func (s *ScannersService) GetScannerTasks(ctx context.Context, scannerId int) ([]ScannerTask, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scanners/%d/tasks", scannerId))
	if err != nil {
		return nil, err
	}

	var result []ScannerTask
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ScannerGroupsResponse represents the response from listing scanner groups.
type ScannerGroupsResponse struct {
	ScannerPools []ScannerGroup `json:"scanner_pools,omitempty"`
}

// ListScannerGroups returns all scanner groups.
func (s *ScannersService) ListScannerGroups(ctx context.Context) (*ScannerGroupsResponse, error) {
	resp, err := s.client.get(ctx, "/scanner-groups")
	if err != nil {
		return nil, err
	}

	var result ScannerGroupsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateScannerGroup creates a new scanner group.
func (s *ScannersService) CreateScannerGroup(ctx context.Context, req *ScannerGroup) (*ScannerGroup, error) {
	resp, err := s.client.post(ctx, "/scanner-groups", req)
	if err != nil {
		return nil, err
	}

	var result ScannerGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetScannerGroup returns a specific scanner group by ID.
func (s *ScannersService) GetScannerGroup(ctx context.Context, groupId int) (*ScannerGroup, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scanner-groups/%d", groupId))
	if err != nil {
		return nil, err
	}

	var result ScannerGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateScannerGroup updates a scanner group.
func (s *ScannersService) UpdateScannerGroup(ctx context.Context, groupId int, req *ScannerGroup) (*ScannerGroup, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/scanner-groups/%d", groupId), req)
	if err != nil {
		return nil, err
	}

	var result ScannerGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteScannerGroup deletes a scanner group.
func (s *ScannersService) DeleteScannerGroup(ctx context.Context, groupId int) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/scanner-groups/%d", groupId))
	return err
}
