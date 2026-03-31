package vm

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// StringOrFloat unmarshals a JSON value that may arrive as either a bare number
// or a quoted float string (e.g. both 8.0 and "8.0" are valid in Tenable exports).
type StringOrFloat float64

func (s *StringOrFloat) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		*s = StringOrFloat(f)
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("StringOrFloat: cannot parse %s: %w", data, err)
	}
	if str == "" {
		*s = 0
		return nil
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("StringOrFloat: cannot convert %q to float64: %w", str, err)
	}
	*s = StringOrFloat(f)
	return nil
}

// ExportsAssetsService handles communication with the Exports (Assets) related endpoints of the VM API.
type ExportsAssetsService struct {
	client *Client
}

type ExportsAssetsServiceExportAssetsV1RequestFilters struct {
	CreatedAt                 int64    `json:"created_at,omitempty"`
	UpdatedAt                 int64    `json:"updated_at,omitempty"`
	TerminatedAt              int64    `json:"terminated_at,omitempty"`
	IsTerminated              bool     `json:"is_terminated,omitempty"`
	DeletedAt                 int64    `json:"deleted_at,omitempty"`
	IsDeleted                 bool     `json:"is_deleted,omitempty"`
	IsLicensed                bool     `json:"is_licensed,omitempty"`
	FirstScanTime             int64    `json:"first_scan_time,omitempty"`
	LastAuthenticatedScanTime int64    `json:"last_authenticated_scan_time,omitempty"`
	LastAssessed              int64    `json:"last_assessed,omitempty"`
	LastScanId                string   `json:"last_scan_id,omitempty"`
	HasPluginResults          bool     `json:"has_plugin_results,omitempty"`
	NetworkId                 string   `json:"network_id,omitempty"`
	ServicenowSysid           bool     `json:"servicenow_sysid,omitempty"`
	Sources                   []string `json:"sources,omitempty"`
	TagCategory               []string `json:"tag.<category>,omitempty"`
}

type ExportsAssetsServiceExportAssetsV1Request struct {
	ChunkSize           int                                              `json:"chunk_size,omitempty"`
	IncludeOpenPorts    bool                                             `json:"include_open_ports,omitempty"`
	IncludeResourceTags bool                                             `json:"include_resource_tags,omitempty"`
	Filters             ExportsAssetsServiceExportAssetsV1RequestFilters `json:"filters,omitempty"`
}

type ExportsAssetsServiceExportAssetsV1Response struct {
	ExportUuid string `json:"export_uuid,omitempty"`
}

type ExportsAssetsServiceExportAssetsV2RequestFilters struct {
	CreatedAt                 int64    `json:"created_at,omitempty"`
	UpdatedAt                 int64    `json:"updated_at,omitempty"`
	TerminatedAt              int64    `json:"terminated_at,omitempty"`
	IsTerminated              bool     `json:"is_terminated,omitempty"`
	DeletedAt                 int64    `json:"deleted_at,omitempty"`
	IsDeleted                 bool     `json:"is_deleted,omitempty"`
	IsLicensed                bool     `json:"is_licensed,omitempty"`
	FirstScanTime             int64    `json:"first_scan_time,omitempty"`
	LastAuthenticatedScanTime int64    `json:"last_authenticated_scan_time,omitempty"`
	LastAssessed              int64    `json:"last_assessed,omitempty"`
	LastScanId                string   `json:"last_scan_id,omitempty"`
	HasPluginResults          bool     `json:"has_plugin_results,omitempty"`
	NetworkId                 string   `json:"network_id,omitempty"`
	ServicenowSysid           bool     `json:"servicenow_sysid,omitempty"`
	Since                     int64    `json:"since,omitempty"`
	Sources                   []string `json:"sources,omitempty"`
	Types                     []string `json:"types,omitempty"`
}

type ExportsAssetsServiceExportAssetsV2Request struct {
	ChunkSize           int                                              `json:"chunk_size,omitempty"`
	IncludeOpenPorts    bool                                             `json:"include_open_ports,omitempty"`
	IncludeResourceTags bool                                             `json:"include_resource_tags,omitempty"`
	Filters             ExportsAssetsServiceExportAssetsV2RequestFilters `json:"filters,omitempty"`
}

type ExportsAssetsServiceExportAssetsV2Response struct {
	ExportUuid string `json:"export_uuid,omitempty"`
}

type ExportsAssetsServiceExportsAssetsExportStatusResponse struct {
	Status          string `json:"status,omitempty"`
	ChunksAvailable []int  `json:"chunks_available,omitempty"`
}

type ExportsAssetsServiceExportsAssetsExportStatusRecentResponseExportsItemFilters struct {
}

type ExportsAssetsServiceExportsAssetsExportStatusRecentResponseExportsItem struct {
	Uuid              string                                                                        `json:"uuid,omitempty"`
	Status            string                                                                        `json:"status,omitempty"`
	TotalChunks       int                                                                           `json:"total_chunks,omitempty"`
	Filters           ExportsAssetsServiceExportsAssetsExportStatusRecentResponseExportsItemFilters `json:"filters,omitempty"`
	FinishedChunks    int                                                                           `json:"finished_chunks,omitempty"`
	NumAssetsPerChunk int                                                                           `json:"num_assets_per_chunk,omitempty"`
	Created           int                                                                           `json:"created,omitempty"`
}

type ExportsAssetsServiceExportsAssetsExportStatusRecentResponse struct {
	Exports []ExportsAssetsServiceExportsAssetsExportStatusRecentResponseExportsItem `json:"exports,omitempty"`
}

type ExportsAssetsServiceExportsAssetsExportCancelResponse struct {
	Status string `json:"status,omitempty"`
}

// ExportAssetsV1 - Export assets v1
func (s *ExportsAssetsService) ExportAssetsV1(ctx context.Context, req *ExportsAssetsServiceExportAssetsV1Request) (*ExportsAssetsServiceExportAssetsV1Response, error) {
	resp, err := s.client.post(ctx, "/assets/export", req)
	if err != nil {
		return nil, err
	}
	var result ExportsAssetsServiceExportAssetsV1Response
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportAssetsV2 - Export assets v2
func (s *ExportsAssetsService) ExportAssetsV2(ctx context.Context, req *ExportsAssetsServiceExportAssetsV2Request) (*ExportsAssetsServiceExportAssetsV2Response, error) {
	resp, err := s.client.post(ctx, "/assets/v2/export", req)
	if err != nil {
		return nil, err
	}
	var result ExportsAssetsServiceExportAssetsV2Response
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsAssetsExportStatus - Get assets export status
func (s *ExportsAssetsService) ExportsAssetsExportStatus(ctx context.Context, exportUuid string) (*ExportsAssetsServiceExportsAssetsExportStatusResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/assets/export/%s/status", exportUuid))
	if err != nil {
		return nil, err
	}
	var result ExportsAssetsServiceExportsAssetsExportStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsAssetsExportStatusRecent - List asset export jobs
func (s *ExportsAssetsService) ExportsAssetsExportStatusRecent(ctx context.Context) (*ExportsAssetsServiceExportsAssetsExportStatusRecentResponse, error) {
	resp, err := s.client.get(ctx, "/assets/export/status")
	if err != nil {
		return nil, err
	}
	var result ExportsAssetsServiceExportsAssetsExportStatusRecentResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsAssetsChunkAssetTag represents a single Tenable tag on an exported asset.
type ExportsAssetsChunkAssetTag struct {
	Uuid    string `json:"uuid,omitempty"`
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
	AddedBy string `json:"added_by,omitempty"`
	AddedAt string `json:"added_at,omitempty"`
}

// ExportsAssetsChunkAssetNetworkInterface represents a network interface on an exported asset.
type ExportsAssetsChunkAssetNetworkInterface struct {
	Name         string   `json:"name,omitempty"`
	MacAddresses []string `json:"mac_addresses,omitempty"`
	Ipv4s        []string `json:"ipv4s,omitempty"`
	Ipv6s        []string `json:"ipv6s,omitempty"`
	Fqdns        []string `json:"fqdns,omitempty"`
	Virtual      *bool    `json:"virtual,omitempty"`
	Aliased      *bool    `json:"aliased,omitempty"`
}

// ExportsAssetsChunkAssetOpenPort represents an open port on an exported asset.
type ExportsAssetsChunkAssetOpenPort struct {
	Port         int      `json:"port,omitempty"`
	Protocol     string   `json:"protocol,omitempty"`
	ServiceNames []string `json:"service_names,omitempty"`
	FirstSeen    string   `json:"first_seen,omitempty"`
	LastSeen     string   `json:"last_seen,omitempty"`
}

// ExportsAssetsChunkAssetSource represents a data source that contributed to an exported asset.
type ExportsAssetsChunkAssetSource struct {
	Name      string `json:"name,omitempty"`
	FirstSeen string `json:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty"`
}

// ExportsAssetsChunkAssetResourceTag represents a cloud provider resource tag on an exported asset.
type ExportsAssetsChunkAssetResourceTag struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// ExportsAssetsChunkAsset represents a single asset record returned within an export chunk.
// Field names match the plural JSON keys used by the Tenable VM export API
// (e.g. "ipv4s", "hostnames", "operating_systems").
type ExportsAssetsChunkAsset struct {
	Id                             string                                    `json:"id,omitempty"`
	HasAgent                       bool                                      `json:"has_agent,omitempty"`
	HasPluginResults               *bool                                     `json:"has_plugin_results,omitempty"`
	CreatedAt                      string                                    `json:"created_at,omitempty"`
	UpdatedAt                      string                                    `json:"updated_at,omitempty"`
	TerminatedAt                   *string                                   `json:"terminated_at,omitempty"`
	DeletedAt                      *string                                   `json:"deleted_at,omitempty"`
	FirstSeen                      string                                    `json:"first_seen,omitempty"`
	LastSeen                       string                                    `json:"last_seen,omitempty"`
	FirstScanTime                  string                                    `json:"first_scan_time,omitempty"`
	LastScanTime                   string                                    `json:"last_scan_time,omitempty"`
	LastScanTarget                 string                                    `json:"last_scan_target,omitempty"`
	LastScanId                     string                                    `json:"last_scan_id,omitempty"`
	LastScheduleId                 string                                    `json:"last_schedule_id,omitempty"`
	LastAuthenticatedScanDate      *string                                   `json:"last_authenticated_scan_date,omitempty"`
	LastAuthenticationAttemptDate  string                                    `json:"last_authentication_attempt_date,omitempty"`
	LastAuthenticationSuccessDate  string                                    `json:"last_authentication_success_date,omitempty"`
	LastAuthenticationScanStatus   string                                    `json:"last_authentication_scan_status,omitempty"`
	LastLicensedScanDate           *string                                   `json:"last_licensed_scan_date,omitempty"`
	AcrScore                       StringOrFloat                             `json:"acr_score,omitempty"`
	ExposureScore                  StringOrFloat                             `json:"exposure_score,omitempty"`
	NetworkId                      string                                    `json:"network_id,omitempty"`
	NetworkName                    string                                    `json:"network_name,omitempty"`
	AgentUuid                      *string                                   `json:"agent_uuid,omitempty"`
	BiosUuid                       *string                                   `json:"bios_uuid,omitempty"`
	SerialNumber                   string                                    `json:"serial_number,omitempty"`
	Ipv4s                          []string                                  `json:"ipv4s,omitempty"`
	Ipv6s                          []string                                  `json:"ipv6s,omitempty"`
	Fqdns                          []string                                  `json:"fqdns,omitempty"`
	MacAddresses                   []string                                  `json:"mac_addresses,omitempty"`
	NetbiosNames                   []string                                  `json:"netbios_names,omitempty"`
	OperatingSystems               []string                                  `json:"operating_systems,omitempty"`
	SystemTypes                    []string                                  `json:"system_types,omitempty"`
	Hostnames                      []string                                  `json:"hostnames,omitempty"`
	AgentNames                     []string                                  `json:"agent_names,omitempty"`
	SshFingerprints                []string                                  `json:"ssh_fingerprints,omitempty"`
	InstalledSoftware              []string                                  `json:"installed_software,omitempty"`
	AwsEc2Name                     string                                    `json:"aws_ec2_name,omitempty"`
	AwsEc2InstanceId               string                                    `json:"aws_ec2_instance_id,omitempty"`
	AwsEc2InstanceAmiId            string                                    `json:"aws_ec2_instance_ami_id,omitempty"`
	AwsEc2InstanceType             string                                    `json:"aws_ec2_instance_type,omitempty"`
	AwsEc2InstanceStateName        string                                    `json:"aws_ec2_instance_state_name,omitempty"`
	AwsEc2InstanceGroupName        string                                    `json:"aws_ec2_instance_group_name,omitempty"`
	AwsOwnerId                     string                                    `json:"aws_owner_id,omitempty"`
	AwsRegion                      string                                    `json:"aws_region,omitempty"`
	AwsAvailabilityZone            string                                    `json:"aws_availability_zone,omitempty"`
	AwsVpcId                       string                                    `json:"aws_vpc_id,omitempty"`
	AwsSubnetId                    string                                    `json:"aws_subnet_id,omitempty"`
	AzureVmId                      *string                                   `json:"azure_vm_id,omitempty"`
	AzureResourceId                *string                                   `json:"azure_resource_id,omitempty"`
	GcpProjectId                   *string                                   `json:"gcp_project_id,omitempty"`
	GcpZone                        *string                                   `json:"gcp_zone,omitempty"`
	GcpInstanceId                  *string                                   `json:"gcp_instance_id,omitempty"`
	ServicenowSysid                *string                                   `json:"servicenow_sysid,omitempty"`
	Tags                           []ExportsAssetsChunkAssetTag              `json:"tags,omitempty"`
	Sources                        []ExportsAssetsChunkAssetSource           `json:"sources,omitempty"`
	NetworkInterfaces              []ExportsAssetsChunkAssetNetworkInterface `json:"network_interfaces,omitempty"`
	OpenPorts                      []ExportsAssetsChunkAssetOpenPort         `json:"open_ports,omitempty"`
	ResourceTags                   []ExportsAssetsChunkAssetResourceTag      `json:"resource_tags,omitempty"`
	TenableAgentDaysSinceActive    int                                       `json:"tenable_agent_days_since_active,omitempty"`
}

// ExportAssetsByTag starts an asset export filtered by the given tag category and values.
// tagCategory is the tag category name (e.g. "Environment"), tagValues are the matching values
// (e.g. ["Production", "Staging"]), and chunkSize controls how many assets are returned per chunk.
func (s *ExportsAssetsService) ExportAssetsByTag(ctx context.Context, tagCategory string, tagValues []string, chunkSize int) (*ExportsAssetsServiceExportAssetsV1Response, error) {
	if chunkSize <= 0 {
		chunkSize = 1000
	}
	body := map[string]interface{}{
		"chunk_size": chunkSize,
		"filters": map[string]interface{}{
			"tag." + tagCategory: tagValues,
		},
	}
	resp, err := s.client.post(ctx, "/assets/export", body)
	if err != nil {
		return nil, err
	}
	var result ExportsAssetsServiceExportAssetsV1Response
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExportsAssetsDownloadChunk - Download assets chunk and return raw JSON bytes.
func (s *ExportsAssetsService) ExportsAssetsDownloadChunk(ctx context.Context, exportUuid string, chunkId string) ([]byte, error) {
	data, err := s.client.get(ctx, fmt.Sprintf("/assets/export/%s/chunks/%s", exportUuid, chunkId))
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ExportsAssetsExportCancel - Cancel asset export
func (s *ExportsAssetsService) ExportsAssetsExportCancel(ctx context.Context, exportUuid string) (*ExportsAssetsServiceExportsAssetsExportCancelResponse, error) {
	resp, err := s.client.post(ctx, fmt.Sprintf("/assets/export/%s/cancel", exportUuid), nil)
	if err != nil {
		return nil, err
	}
	var result ExportsAssetsServiceExportsAssetsExportCancelResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
