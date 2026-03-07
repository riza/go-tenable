package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// ScanHistoryService handles communication with the Scan History related endpoints of the VM API.
type ScanHistoryService struct {
	client *Client
}

type ScanHistoryServiceHistoryResponsePaginationSortItem struct {
	Name  string `json:"name,omitempty"`
	Order string `json:"order,omitempty"`
}

type ScanHistoryServiceHistoryResponsePagination struct {
	Total  int                                                   `json:"total,omitempty"`
	Limit  int                                                   `json:"limit,omitempty"`
	Offset int                                                   `json:"offset,omitempty"`
	Sort   []ScanHistoryServiceHistoryResponsePaginationSortItem `json:"sort,omitempty"`
}

type ScanHistoryServiceHistoryResponseHistoryItemTargets struct {
	Custom       bool `json:"custom,omitempty"`
	DefaultValue bool `json:"default,omitempty"`
}

type ScanHistoryServiceHistoryResponseHistoryItem struct {
	Id            int                                                 `json:"id,omitempty"`
	Status        string                                              `json:"status,omitempty"`
	ReportingMode string                                              `json:"reporting_mode,omitempty"`
	IsArchived    bool                                                `json:"is_archived,omitempty"`
	Visibility    string                                              `json:"visibility,omitempty"`
	Targets       ScanHistoryServiceHistoryResponseHistoryItemTargets `json:"targets,omitempty"`
	TimeEnd       int                                                 `json:"time_end,omitempty"`
	ScanUuid      string                                              `json:"scan_uuid,omitempty"`
	TimeStart     int                                                 `json:"time_start,omitempty"`
}

type ScanHistoryServiceHistoryResponse struct {
	Pagination ScanHistoryServiceHistoryResponsePagination    `json:"pagination,omitempty"`
	History    []ScanHistoryServiceHistoryResponseHistoryItem `json:"history,omitempty"`
}

type ScanHistoryServiceHistoryDetailsResponse struct {
	OwnerId       int    `json:"owner_id,omitempty"`
	ScheduleUuid  string `json:"schedule_uuid,omitempty"`
	Status        string `json:"status,omitempty"`
	ReportingMode string `json:"reporting_mode,omitempty"`
	IsArchived    bool   `json:"is_archived,omitempty"`
	ScanStart     int    `json:"scan_start,omitempty"`
	OwnerUuid     int    `json:"owner_uuid,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Targets       string `json:"targets,omitempty"`
	ObjectId      int    `json:"object_id,omitempty"`
	Uuid          string `json:"uuid,omitempty"`
	ScanEnd       int    `json:"scan_end,omitempty"`
	ScanType      string `json:"scan_type,omitempty"`
	Name          string `json:"name,omitempty"`
}

// History - Get scan history
func (s *ScanHistoryService) History(ctx context.Context, scanId string) (*ScanHistoryServiceHistoryResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scans/%s/history", scanId))
	if err != nil {
		return nil, err
	}
	var result ScanHistoryServiceHistoryResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// HistoryDetails - Get scan history details
func (s *ScanHistoryService) HistoryDetails(ctx context.Context, scanId string, historyUuid string) (*ScanHistoryServiceHistoryDetailsResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scans/%s/history/%s", scanId, historyUuid))
	if err != nil {
		return nil, err
	}
	var result ScanHistoryServiceHistoryDetailsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteHistory - Delete scan history
func (s *ScanHistoryService) DeleteHistory(ctx context.Context, scanId string, historyId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/scans/%s/history/%s", scanId, historyId))
	if err != nil {
		return err
	}
	return nil
}
