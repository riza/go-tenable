package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// ScanTasksService handles communication with the Scan Tasks related endpoints of the VM API.
type ScanTasksService struct {
	client *Client
}

type ScanTasksServiceScheduleRequest struct {
	Enabled bool `json:"enabled,omitempty"`
}

type ScanTasksServiceScheduleResponse struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Control   bool   `json:"control,omitempty"`
	Rrules    string `json:"rrules,omitempty"`
	Starttime string `json:"starttime,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
}

type ScanTasksServiceCopyRequest struct {
	FolderId int    `json:"folder_id,omitempty"`
	Name     string `json:"name,omitempty"`
}

type ScanTasksServiceCopyResponse struct {
	Name                 string `json:"name,omitempty"`
	Status               string `json:"status,omitempty"`
	Id                   int    `json:"id,omitempty"`
	LastModificationDate int    `json:"last_modification_date,omitempty"`
	Uuid                 string `json:"uuid,omitempty"`
	TypeField            string `json:"type,omitempty"`
	Owner                string `json:"owner,omitempty"`
	Enabled              bool   `json:"enabled,omitempty"`
	Read                 bool   `json:"read,omitempty"`
	Shared               bool   `json:"shared,omitempty"`
	UserPermissions      int    `json:"user_permissions,omitempty"`
	CreationDate         int    `json:"creation_date,omitempty"`
	Control              bool   `json:"control,omitempty"`
	Starttime            string `json:"starttime,omitempty"`
	Timezone             string `json:"timezone,omitempty"`
	Rrules               string `json:"rrules,omitempty"`
	ScheduleUuid         string `json:"schedule_uuid,omitempty"`
}

type ScanTasksServiceCredentialsConvertRequestSettings struct {
	Domain     string `json:"domain,omitempty"`
	Username   string `json:"username,omitempty"`
	AuthMethod string `json:"auth_method,omitempty"`
	Password   string `json:"password,omitempty"`
}

type ScanTasksServiceCredentialsConvertRequestPermissionsItem struct {
	GranteeUuid string `json:"grantee_uuid,omitempty"`
	TypeField   string `json:"type,omitempty"`
	Permissions int    `json:"permissions,omitempty"`
	Name        string `json:"name,omitempty"`
}

type ScanTasksServiceCredentialsConvertRequest struct {
	Name        string                                                     `json:"name,omitempty"`
	Settings    ScanTasksServiceCredentialsConvertRequestSettings          `json:"settings,omitempty"`
	TypeField   string                                                     `json:"type,omitempty"`
	Category    string                                                     `json:"category,omitempty"`
	AdHoc       bool                                                       `json:"ad_hoc,omitempty"`
	Permissions []ScanTasksServiceCredentialsConvertRequestPermissionsItem `json:"permissions,omitempty"`
}

type ScanTasksServiceCredentialsConvertResponse struct {
	Uuid string `json:"uuid,omitempty"`
}

type ScanTasksServiceImportRequest struct {
	File     string `json:"file,omitempty"`
	FolderId int    `json:"folder_id,omitempty"`
	Password string `json:"password,omitempty"`
}

type ScanTasksServiceImportResponse struct {
	Name                 string `json:"name,omitempty"`
	Status               string `json:"status,omitempty"`
	Id                   int    `json:"id,omitempty"`
	LastModificationDate int    `json:"last_modification_date,omitempty"`
	Uuid                 string `json:"uuid,omitempty"`
	TypeField            string `json:"type,omitempty"`
	Owner                string `json:"owner,omitempty"`
	Enabled              bool   `json:"enabled,omitempty"`
	Read                 bool   `json:"read,omitempty"`
	Shared               bool   `json:"shared,omitempty"`
	UserPermissions      int    `json:"user_permissions,omitempty"`
	CreationDate         int    `json:"creation_date,omitempty"`
	Control              bool   `json:"control,omitempty"`
	Starttime            string `json:"starttime,omitempty"`
	Timezone             string `json:"timezone,omitempty"`
	Rrules               string `json:"rrules,omitempty"`
	ScheduleUuid         string `json:"schedule_uuid,omitempty"`
}

type ScanTasksServiceCountResponse struct {
	Count int `json:"count,omitempty"`
}

type ScanTasksServiceTimezonesResponseItem struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ScanTasksServiceCheckAutoTargetsRequest struct {
	NetworkUuid string   `json:"network_uuid,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	TargetList  string   `json:"target_list,omitempty"`
}

type ScanTasksServiceCheckAutoTargetsResponse struct {
	MissedTargets             []string `json:"missed_targets,omitempty"`
	TotalMissedTargets        int      `json:"total_missed_targets,omitempty"`
	MatchedResourceUuids      []string `json:"matched_resource_uuids,omitempty"`
	TotalMatchedResourceUuids int      `json:"total_matched_resource_uuids,omitempty"`
}

// Schedule - Enable schedule
func (s *ScanTasksService) Schedule(ctx context.Context, scanId string, req *ScanTasksServiceScheduleRequest) (*ScanTasksServiceScheduleResponse, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/scans/%s/schedule", scanId), req)
	if err != nil {
		return nil, err
	}
	var result ScanTasksServiceScheduleResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Copy - Copy scan
func (s *ScanTasksService) Copy(ctx context.Context, scanId string, req *ScanTasksServiceCopyRequest) (*ScanTasksServiceCopyResponse, error) {
	resp, err := s.client.post(ctx, fmt.Sprintf("/scans/%s/copy", scanId), req)
	if err != nil {
		return nil, err
	}
	var result ScanTasksServiceCopyResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CredentialsConvert - Convert credentials
func (s *ScanTasksService) CredentialsConvert(ctx context.Context, scanId string, credentialsId string, req *ScanTasksServiceCredentialsConvertRequest) (*ScanTasksServiceCredentialsConvertResponse, error) {
	resp, err := s.client.post(ctx, fmt.Sprintf("/scans/%s/credentials/%s/upgrade", scanId, credentialsId), req)
	if err != nil {
		return nil, err
	}
	var result ScanTasksServiceCredentialsConvertResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Import - Import uploaded scan
func (s *ScanTasksService) Import(ctx context.Context, req *ScanTasksServiceImportRequest) (*ScanTasksServiceImportResponse, error) {
	resp, err := s.client.post(ctx, "/scans/import", req)
	if err != nil {
		return nil, err
	}
	var result ScanTasksServiceImportResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Count - Get scan count
func (s *ScanTasksService) Count(ctx context.Context) (*ScanTasksServiceCountResponse, error) {
	resp, err := s.client.get(ctx, "/scans/count")
	if err != nil {
		return nil, err
	}
	var result ScanTasksServiceCountResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Timezones - Get timezones
func (s *ScanTasksService) Timezones(ctx context.Context) ([]ScanTasksServiceTimezonesResponseItem, error) {
	resp, err := s.client.get(ctx, "/scans/timezones")
	if err != nil {
		return nil, err
	}
	var result []ScanTasksServiceTimezonesResponseItem
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CheckAutoTargets - Test scan routes
func (s *ScanTasksService) CheckAutoTargets(ctx context.Context, req *ScanTasksServiceCheckAutoTargetsRequest) (*ScanTasksServiceCheckAutoTargetsResponse, error) {
	resp, err := s.client.post(ctx, "/scans/check-auto-targets", req)
	if err != nil {
		return nil, err
	}
	var result ScanTasksServiceCheckAutoTargetsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
