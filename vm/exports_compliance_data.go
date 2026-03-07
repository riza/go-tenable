package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// ExportsComplianceDataService handles communication with the Exports (Compliance Data) related endpoints of the VM API.
type ExportsComplianceDataService struct {
	client *Client
}

type ExportsComplianceDataServiceExportsComplianceCreateRequestFiltersTagsItem struct {
	Category string   `json:"category,omitempty"`
	Values   []string `json:"values,omitempty"`
}

type ExportsComplianceDataServiceExportsComplianceCreateRequestFilters struct {
	LastSeen          int64                                                                       `json:"last_seen,omitempty"`
	FirstSeen         int64                                                                       `json:"first_seen,omitempty"`
	LastObserved      int64                                                                       `json:"last_observed,omitempty"`
	IndexedAt         int64                                                                       `json:"indexed_at,omitempty"`
	Since             int64                                                                       `json:"since,omitempty"`
	AuditName         string                                                                      `json:"audit_name,omitempty"`
	AuditFileName     string                                                                      `json:"audit_file_name,omitempty"`
	ComplianceResults []string                                                                    `json:"compliance_results,omitempty"`
	Ipv4Addresses     []string                                                                    `json:"ipv4_addresses,omitempty"`
	Ipv6Addresses     []string                                                                    `json:"ipv6_addresses,omitempty"`
	NetworkId         string                                                                      `json:"network_id,omitempty"`
	PluginId          []int                                                                       `json:"plugin_id,omitempty"`
	State             []string                                                                    `json:"state,omitempty"`
	Tags              []ExportsComplianceDataServiceExportsComplianceCreateRequestFiltersTagsItem `json:"tags,omitempty"`
}

type ExportsComplianceDataServiceExportsComplianceCreateRequest struct {
	NumFindings int                                                               `json:"num_findings,omitempty"`
	Asset       []string                                                          `json:"asset,omitempty"`
	Filters     ExportsComplianceDataServiceExportsComplianceCreateRequestFilters `json:"filters,omitempty"`
}

type ExportsComplianceDataServiceExportsComplianceCreateResponse struct {
	ExportUuid string `json:"export_uuid,omitempty"`
}

type ExportsComplianceDataServiceExportsComplianceStatusResponseFilters struct {
}

type ExportsComplianceDataServiceExportsComplianceStatusResponse struct {
	Uuid                 string                                                             `json:"uuid,omitempty"`
	Status               string                                                             `json:"status,omitempty"`
	ChunksAvailable      []int                                                              `json:"chunks_available,omitempty"`
	ChunksFailed         []int                                                              `json:"chunks_failed,omitempty"`
	ChunksCancelled      []int                                                              `json:"chunks_cancelled,omitempty"`
	ChunksAvailableCount int                                                                `json:"chunks_available_count,omitempty"`
	EmptyChunksCount     int                                                                `json:"empty_chunks_count,omitempty"`
	FinishedChunks       int                                                                `json:"finished_chunks,omitempty"`
	Filters              ExportsComplianceDataServiceExportsComplianceStatusResponseFilters `json:"filters,omitempty"`
	NumFindings          int                                                                `json:"num_findings,omitempty"`
	Asset                []string                                                           `json:"asset,omitempty"`
	Created              int                                                                `json:"created,omitempty"`
}

type ExportsComplianceDataServiceExportsComplianceStatusListResponseExportsItemFilters struct {
}

type ExportsComplianceDataServiceExportsComplianceStatusListResponseExportsItem struct {
	Uuid                 string                                                                            `json:"uuid,omitempty"`
	Status               string                                                                            `json:"status,omitempty"`
	ChunksAvailableCount int                                                                               `json:"chunks_available_count,omitempty"`
	EmptyChunksCount     int                                                                               `json:"empty_chunks_count,omitempty"`
	FinishedChunks       int                                                                               `json:"finished_chunks,omitempty"`
	Filters              ExportsComplianceDataServiceExportsComplianceStatusListResponseExportsItemFilters `json:"filters,omitempty"`
	NumFindings          int                                                                               `json:"num_findings,omitempty"`
	Asset                []string                                                                          `json:"asset,omitempty"`
	Created              int                                                                               `json:"created,omitempty"`
}

type ExportsComplianceDataServiceExportsComplianceStatusListResponse struct {
	Exports []ExportsComplianceDataServiceExportsComplianceStatusListResponseExportsItem `json:"exports,omitempty"`
}

type ExportsComplianceDataServiceExportsComplianceDownloadResponseItemReferenceItem struct {
	Framework string `json:"framework,omitempty"`
	Control   string `json:"control,omitempty"`
}

type ExportsComplianceDataServiceExportsComplianceDownloadResponseItemAssetsTagsItem struct {
	Category string   `json:"category,omitempty"`
	Values   []string `json:"values,omitempty"`
}

type ExportsComplianceDataServiceExportsComplianceDownloadResponseItemAssets struct {
	Id               string                                                                            `json:"id,omitempty"`
	Ipv4Addresses    []string                                                                          `json:"ipv4_addresses,omitempty"`
	Ipv6Addresses    []string                                                                          `json:"ipv6_addresses,omitempty"`
	Fqdns            []string                                                                          `json:"fqdns,omitempty"`
	Name             string                                                                            `json:"name,omitempty"`
	AgentName        string                                                                            `json:"agent_name,omitempty"`
	AgentUuid        string                                                                            `json:"agent_uuid,omitempty"`
	NetbiosName      string                                                                            `json:"netbios_name,omitempty"`
	MacAddresses     []string                                                                          `json:"mac_addresses,omitempty"`
	OperatingSystems []string                                                                          `json:"operating_systems,omitempty"`
	SystemType       string                                                                            `json:"system_type,omitempty"`
	NetworkId        string                                                                            `json:"network_id,omitempty"`
	Tags             []ExportsComplianceDataServiceExportsComplianceDownloadResponseItemAssetsTagsItem `json:"tags,omitempty"`
}

type ExportsComplianceDataServiceExportsComplianceDownloadResponseItem struct {
	AssetUuid                  string                                                                           `json:"asset_uuid,omitempty"`
	FirstSeen                  string                                                                           `json:"first_seen,omitempty"`
	LastSeen                   string                                                                           `json:"last_seen,omitempty"`
	AuditFile                  string                                                                           `json:"audit_file,omitempty"`
	CheckId                    string                                                                           `json:"check_id,omitempty"`
	CheckName                  string                                                                           `json:"check_name,omitempty"`
	CheckInfo                  string                                                                           `json:"check_info,omitempty"`
	ExpectedValue              string                                                                           `json:"expected_value,omitempty"`
	ActualValue                string                                                                           `json:"actual_value,omitempty"`
	Status                     string                                                                           `json:"status,omitempty"`
	Reference                  []ExportsComplianceDataServiceExportsComplianceDownloadResponseItemReferenceItem `json:"reference,omitempty"`
	SeeAlso                    string                                                                           `json:"see_also,omitempty"`
	Solution                   string                                                                           `json:"solution,omitempty"`
	DbType                     string                                                                           `json:"db_type,omitempty"`
	PluginId                   int                                                                              `json:"plugin_id,omitempty"`
	State                      string                                                                           `json:"state,omitempty"`
	CheckError                 string                                                                           `json:"check_error,omitempty"`
	ProfileName                string                                                                           `json:"profile_name,omitempty"`
	Description                string                                                                           `json:"description,omitempty"`
	ComplianceBenchmarkName    string                                                                           `json:"compliance_benchmark_name,omitempty"`
	ComplianceBenchmarkVersion string                                                                           `json:"compliance_benchmark_version,omitempty"`
	ComplianceControlId        string                                                                           `json:"compliance_control_id,omitempty"`
	ComplianceFullId           string                                                                           `json:"compliance_full_id,omitempty"`
	ComplianceFunctionalId     string                                                                           `json:"compliance_functional_id,omitempty"`
	ComplianceInformationalId  string                                                                           `json:"compliance_informational_id,omitempty"`
	Synopsis                   string                                                                           `json:"synopsis,omitempty"`
	LastFixed                  string                                                                           `json:"last_fixed,omitempty"`
	LastObserved               string                                                                           `json:"last_observed,omitempty"`
	MetadataId                 string                                                                           `json:"metadata_id,omitempty"`
	UnameOutput                string                                                                           `json:"uname_output,omitempty"`
	IndexedAt                  string                                                                           `json:"indexed_at,omitempty"`
	PluginName                 string                                                                           `json:"plugin_name,omitempty"`
	Assets                     ExportsComplianceDataServiceExportsComplianceDownloadResponseItemAssets          `json:"assets,omitempty"`
}

type ExportsComplianceDataServiceExportsComplianceCancelResponse struct {
	Status string `json:"status,omitempty"`
}

// ExportsComplianceCreate - Export compliance data
func (s *ExportsComplianceDataService) ExportsComplianceCreate(ctx context.Context, req *ExportsComplianceDataServiceExportsComplianceCreateRequest) (*ExportsComplianceDataServiceExportsComplianceCreateResponse, error) {
	resp, err := s.client.post(ctx, "/compliance/export", req)
	if err != nil {
		return nil, err
	}
	var result ExportsComplianceDataServiceExportsComplianceCreateResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsComplianceStatus - Get compliance export status
func (s *ExportsComplianceDataService) ExportsComplianceStatus(ctx context.Context, exportUuid string) (*ExportsComplianceDataServiceExportsComplianceStatusResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/compliance/export/%s/status", exportUuid))
	if err != nil {
		return nil, err
	}
	var result ExportsComplianceDataServiceExportsComplianceStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsComplianceStatusList - List compliance export jobs
func (s *ExportsComplianceDataService) ExportsComplianceStatusList(ctx context.Context) (*ExportsComplianceDataServiceExportsComplianceStatusListResponse, error) {
	resp, err := s.client.get(ctx, "/compliance/export/status")
	if err != nil {
		return nil, err
	}
	var result ExportsComplianceDataServiceExportsComplianceStatusListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsComplianceDownload - Download compliance chunk
func (s *ExportsComplianceDataService) ExportsComplianceDownload(ctx context.Context, exportUuid string, chunkId string) ([]ExportsComplianceDataServiceExportsComplianceDownloadResponseItem, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/compliance/export/%s/chunks/%s", exportUuid, chunkId))
	if err != nil {
		return nil, err
	}
	var result []ExportsComplianceDataServiceExportsComplianceDownloadResponseItem
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ExportsComplianceCancel - Cancel compliance export
func (s *ExportsComplianceDataService) ExportsComplianceCancel(ctx context.Context, exportUuid string) (*ExportsComplianceDataServiceExportsComplianceCancelResponse, error) {
	resp, err := s.client.post(ctx, fmt.Sprintf("/compliance/export/%s/cancel", exportUuid), nil)
	if err != nil {
		return nil, err
	}
	var result ExportsComplianceDataServiceExportsComplianceCancelResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
