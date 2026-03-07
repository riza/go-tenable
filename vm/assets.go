package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// AssetsService handles communication with the Assets related endpoints of the VM API.
type AssetsService struct {
	client *Client
}

type AssetsServiceListAssetsResponseAssetsItemSourcesItem struct {
	Name      string `json:"name,omitempty"`
	FirstSeen string `json:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty"`
}

type AssetsServiceListAssetsResponseAssetsItemAcrDriversItem struct {
	DriverName  string        `json:"driver_name,omitempty"`
	DriverValue []interface{} `json:"driver_value,omitempty"`
}

type AssetsServiceListAssetsResponseAssetsItemScanFrequencyItem struct {
	Interval  int  `json:"interval,omitempty"`
	Frequency int  `json:"frequency,omitempty"`
	Licensed  bool `json:"licensed,omitempty"`
}

type AssetsServiceListAssetsResponseAssetsItem struct {
	Id                      string                                                       `json:"id,omitempty"`
	HasAgent                bool                                                         `json:"has_agent,omitempty"`
	LastSeen                string                                                       `json:"last_seen,omitempty"`
	LastScanTarget          string                                                       `json:"last_scan_target,omitempty"`
	Sources                 []AssetsServiceListAssetsResponseAssetsItemSourcesItem       `json:"sources,omitempty"`
	AcrScore                int                                                          `json:"acr_score,omitempty"`
	AcrDrivers              []AssetsServiceListAssetsResponseAssetsItemAcrDriversItem    `json:"acr_drivers,omitempty"`
	ExposureScore           int                                                          `json:"exposure_score,omitempty"`
	ScanFrequency           []AssetsServiceListAssetsResponseAssetsItemScanFrequencyItem `json:"scan_frequency,omitempty"`
	Ipv4                    []string                                                     `json:"ipv4,omitempty"`
	Ipv6                    []string                                                     `json:"ipv6,omitempty"`
	Fqdn                    []string                                                     `json:"fqdn,omitempty"`
	MacAddress              []string                                                     `json:"mac_address,omitempty"`
	NetbiosName             []string                                                     `json:"netbios_name,omitempty"`
	OperatingSystem         []string                                                     `json:"operating_system,omitempty"`
	Hostname                []string                                                     `json:"hostname,omitempty"`
	AgentName               []string                                                     `json:"agent_name,omitempty"`
	AwsEc2Name              []string                                                     `json:"aws_ec2_name,omitempty"`
	SecurityProtectionLevel int                                                          `json:"security_protection_level,omitempty"`
	SecurityProtections     []string                                                     `json:"security_protections,omitempty"`
	ExposureConfidenceValue int                                                          `json:"exposure_confidence_value,omitempty"`
}

type AssetsServiceListAssetsResponse struct {
	Assets []AssetsServiceListAssetsResponseAssetsItem `json:"assets,omitempty"`
	Total  int                                         `json:"total,omitempty"`
}

type AssetsServiceAssetInfoResponseSourcesItem struct {
	Name      string `json:"name,omitempty"`
	FirstSeen string `json:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty"`
}

type AssetsServiceAssetInfoResponseTagsItem struct {
	TagUuid  string `json:"tag_uuid,omitempty"`
	TagKey   string `json:"tag_key,omitempty"`
	TagValue string `json:"tag_value,omitempty"`
	AddedBy  string `json:"added_by,omitempty"`
	AddedAt  string `json:"added_at,omitempty"`
}

type AssetsServiceAssetInfoResponseAcrDriversItem struct {
	DriverName  string        `json:"driver_name,omitempty"`
	DriverValue []interface{} `json:"driver_value,omitempty"`
}

type AssetsServiceAssetInfoResponseScanFrequencyItem struct {
	Interval  int  `json:"interval,omitempty"`
	Frequency int  `json:"frequency,omitempty"`
	Licensed  bool `json:"licensed,omitempty"`
}

type AssetsServiceAssetInfoResponse struct {
	Name                          string                                            `json:"name,omitempty"`
	Id                            string                                            `json:"id,omitempty"`
	HasAgent                      bool                                              `json:"has_agent,omitempty"`
	CreatedAt                     string                                            `json:"created_at,omitempty"`
	UpdatedAt                     string                                            `json:"updated_at,omitempty"`
	TerminatedAt                  string                                            `json:"terminated_at,omitempty"`
	DeletedAt                     string                                            `json:"deleted_at,omitempty"`
	AesScoreV3                    float64                                           `json:"aes_score_v3,omitempty"`
	AcrScoreV3                    float64                                           `json:"acr_score_v3,omitempty"`
	FirstSeen                     string                                            `json:"first_seen,omitempty"`
	LastSeen                      string                                            `json:"last_seen,omitempty"`
	LastScanTarget                string                                            `json:"last_scan_target,omitempty"`
	LastAuthenticationAttemptDate string                                            `json:"last_authentication_attempt_date,omitempty"`
	LastAuthenticationSuccessDate string                                            `json:"last_authentication_success_date,omitempty"`
	LastAuthenticatedScanDate     string                                            `json:"last_authenticated_scan_date,omitempty"`
	LastLicensedScanDate          string                                            `json:"last_licensed_scan_date,omitempty"`
	LastScanId                    string                                            `json:"last_scan_id,omitempty"`
	LastScheduleId                string                                            `json:"last_schedule_id,omitempty"`
	Sources                       []AssetsServiceAssetInfoResponseSourcesItem       `json:"sources,omitempty"`
	Tags                          []AssetsServiceAssetInfoResponseTagsItem          `json:"tags,omitempty"`
	AcrScore                      float64                                           `json:"acr_score,omitempty"`
	AcrDrivers                    []AssetsServiceAssetInfoResponseAcrDriversItem    `json:"acr_drivers,omitempty"`
	ExposureScore                 float64                                           `json:"exposure_score,omitempty"`
	ScanFrequency                 []AssetsServiceAssetInfoResponseScanFrequencyItem `json:"scan_frequency,omitempty"`
	NetworkId                     []string                                          `json:"network_id,omitempty"`
	Ipv4                          []string                                          `json:"ipv4,omitempty"`
	Ipv6                          []string                                          `json:"ipv6,omitempty"`
	Fqdn                          []string                                          `json:"fqdn,omitempty"`
	MacAddress                    []string                                          `json:"mac_address,omitempty"`
	NetbiosName                   []string                                          `json:"netbios_name,omitempty"`
	OperatingSystem               []string                                          `json:"operating_system,omitempty"`
	SystemType                    []string                                          `json:"system_type,omitempty"`
	TenableUuid                   []string                                          `json:"tenable_uuid,omitempty"`
	Hostname                      []string                                          `json:"hostname,omitempty"`
	AgentName                     []string                                          `json:"agent_name,omitempty"`
	BiosUuid                      []string                                          `json:"bios_uuid,omitempty"`
	AwsEc2InstanceId              []string                                          `json:"aws_ec2_instance_id,omitempty"`
	AwsEc2InstanceAmiId           []string                                          `json:"aws_ec2_instance_ami_id,omitempty"`
	AwsOwnerId                    []string                                          `json:"aws_owner_id,omitempty"`
	AwsAvailabilityZone           []string                                          `json:"aws_availability_zone,omitempty"`
	AwsRegion                     []string                                          `json:"aws_region,omitempty"`
	AwsVpcId                      []string                                          `json:"aws_vpc_id,omitempty"`
	AwsEc2InstanceGroupName       []string                                          `json:"aws_ec2_instance_group_name,omitempty"`
	AwsEc2InstanceStateName       []string                                          `json:"aws_ec2_instance_state_name,omitempty"`
	AwsEc2InstanceType            []string                                          `json:"aws_ec2_instance_type,omitempty"`
	AwsSubnetId                   []string                                          `json:"aws_subnet_id,omitempty"`
	AwsEc2ProductCode             []string                                          `json:"aws_ec2_product_code,omitempty"`
	AwsEc2Name                    []string                                          `json:"aws_ec2_name,omitempty"`
	AzureVmId                     []string                                          `json:"azure_vm_id,omitempty"`
	AzureResourceId               []string                                          `json:"azure_resource_id,omitempty"`
	GcpProjectId                  []string                                          `json:"gcp_project_id,omitempty"`
	GcpZone                       []string                                          `json:"gcp_zone,omitempty"`
	GcpInstanceId                 []string                                          `json:"gcp_instance_id,omitempty"`
	SshFingerprint                []string                                          `json:"ssh_fingerprint,omitempty"`
	McafeeEpoGuid                 []string                                          `json:"mcafee_epo_guid,omitempty"`
	McafeeEpoAgentGuid            []string                                          `json:"mcafee_epo_agent_guid,omitempty"`
	QualysAssetId                 []string                                          `json:"qualys_asset_id,omitempty"`
	QualysHostId                  []string                                          `json:"qualys_host_id,omitempty"`
	ServicenowSysid               []string                                          `json:"servicenow_sysid,omitempty"`
	InstalledSoftware             []string                                          `json:"installed_software,omitempty"`
	SecurityProtectionLevel       int                                               `json:"security_protection_level,omitempty"`
	SecurityProtections           []string                                          `json:"security_protections,omitempty"`
	ExposureConfidenceValue       int                                               `json:"exposure_confidence_value,omitempty"`
}

type AssetsServiceBulkUpdateAcrRequestItemAssetItem struct {
	Id          string   `json:"id,omitempty"`
	Fqdn        []string `json:"fqdn,omitempty"`
	MacAddress  []string `json:"mac_address,omitempty"`
	NetbiosName []string `json:"netbios_name,omitempty"`
	Ipv4        []string `json:"ipv4,omitempty"`
}

type AssetsServiceBulkUpdateAcrRequestItem struct {
	AcrScore int                                              `json:"acr_score,omitempty"`
	Reason   []string                                         `json:"reason,omitempty"`
	Note     string                                           `json:"note,omitempty"`
	Asset    []AssetsServiceBulkUpdateAcrRequestItemAssetItem `json:"asset,omitempty"`
}

type AssetsServiceBulkMoveRequest struct {
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
	Targets     string `json:"targets,omitempty"`
}

type AssetsServiceBulkDeleteRequest struct {
	HardDelete bool        `json:"hard_delete,omitempty"`
	Query      interface{} `json:"query,omitempty"`
}

type AssetsServiceImportRequestAssetsItem struct {
	MacAddress              []string `json:"mac_address,omitempty"`
	NetbiosName             string   `json:"netbios_name,omitempty"`
	Fqdn                    []string `json:"fqdn,omitempty"`
	IpAddress               []string `json:"ip_address,omitempty"`
	Ipv4                    []string `json:"ipv4,omitempty"`
	Ipv6                    []string `json:"ipv6,omitempty"`
	Hostname                []string `json:"hostname,omitempty"`
	OperatingSystem         []string `json:"operating_system,omitempty"`
	SshFingerprint          string   `json:"ssh_fingerprint,omitempty"`
	BiosUuid                string   `json:"bios_uuid,omitempty"`
	ManufacturerTpmId       string   `json:"manufacturer_tpm_id,omitempty"`
	McafeeEpoGuid           string   `json:"mcafee_epo_guid,omitempty"`
	McafeeEpoAgentGuid      string   `json:"mcafee_epo_agent_guid,omitempty"`
	SymantecEpHardwareKey   string   `json:"symantec_ep_hardware_key,omitempty"`
	QualysAssetId           string   `json:"qualys_asset_id,omitempty"`
	QualysHostId            string   `json:"qualys_host_id,omitempty"`
	ServicenowSysId         string   `json:"servicenow_sys_id,omitempty"`
	GcpProjectId            string   `json:"gcp_project_id,omitempty"`
	GcpZone                 string   `json:"gcp_zone,omitempty"`
	GcpInstanceId           string   `json:"gcp_instance_id,omitempty"`
	AzureVmId               string   `json:"azure_vm_id,omitempty"`
	AzureResourceId         string   `json:"azure_resource_id,omitempty"`
	AwsAvailabilityZone     string   `json:"aws_availability_zone,omitempty"`
	AwsEc2InstanceId        string   `json:"aws_ec2_instance_id,omitempty"`
	AwsEc2InstanceAmiId     string   `json:"aws_ec2_instance_ami_id,omitempty"`
	AwsEc2InstanceGroupName string   `json:"aws_ec2_instance_group_name,omitempty"`
	AwsEc2InstanceStateName string   `json:"aws_ec2_instance_state_name,omitempty"`
	AwsEc2InstanceType      string   `json:"aws_ec2_instance_type,omitempty"`
	AwsEc2Name              string   `json:"aws_ec2_name,omitempty"`
	AwsEc2ProductCode       string   `json:"aws_ec2_product_code,omitempty"`
	AwsOwnerId              string   `json:"aws_owner_id,omitempty"`
	AwsRegion               string   `json:"aws_region,omitempty"`
	AwsSubnetId             string   `json:"aws_subnet_id,omitempty"`
	AwsVpcId                string   `json:"aws_vpc_id,omitempty"`
	InstalledSoftware       []string `json:"installed_software,omitempty"`
	BigfixAssetId           []string `json:"bigfix_asset_id,omitempty"`
}

type AssetsServiceImportRequest struct {
	Assets []AssetsServiceImportRequestAssetsItem `json:"assets,omitempty"`
	Source string                                 `json:"source,omitempty"`
}

type AssetsServiceImportResponse struct {
	AssetImportJobUuid string `json:"asset_import_job_uuid,omitempty"`
}

type AssetsServiceListImportJobsResponseAssetImportJobsItem struct {
	JobId          string `json:"job_id,omitempty"`
	ContainerId    string `json:"container_id,omitempty"`
	Source         string `json:"source,omitempty"`
	Batches        int    `json:"batches,omitempty"`
	UploadedAssets int    `json:"uploaded_assets,omitempty"`
	FailedAssets   int    `json:"failed_assets,omitempty"`
	StartTime      int    `json:"start_time,omitempty"`
	LastUpdateTime int    `json:"last_update_time,omitempty"`
	EndTime        int    `json:"end_time,omitempty"`
	Status         string `json:"status,omitempty"`
	StatusMessage  string `json:"status_message,omitempty"`
}

type AssetsServiceListImportJobsResponse struct {
	AssetImportJobs []AssetsServiceListImportJobsResponseAssetImportJobsItem `json:"asset_import_jobs,omitempty"`
}

type AssetsServiceImportJobInfoResponse struct {
	JobId          string `json:"job_id,omitempty"`
	ContainerId    string `json:"container_id,omitempty"`
	Source         string `json:"source,omitempty"`
	Batches        int    `json:"batches,omitempty"`
	UploadedAssets int    `json:"uploaded_assets,omitempty"`
	FailedAssets   int    `json:"failed_assets,omitempty"`
	StartTime      int    `json:"start_time,omitempty"`
	LastUpdateTime int    `json:"last_update_time,omitempty"`
	EndTime        int    `json:"end_time,omitempty"`
	Status         string `json:"status,omitempty"`
	StatusMessage  string `json:"status_message,omitempty"`
}

// ListAssets - List assets
func (s *AssetsService) ListAssets(ctx context.Context) (*AssetsServiceListAssetsResponse, error) {
	resp, err := s.client.get(ctx, "/assets")
	if err != nil {
		return nil, err
	}
	var result AssetsServiceListAssetsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AssetInfo - Get asset details
func (s *AssetsService) AssetInfo(ctx context.Context, assetUuid string) (*AssetsServiceAssetInfoResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/assets/%s", assetUuid))
	if err != nil {
		return nil, err
	}
	var result AssetsServiceAssetInfoResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BulkUpdateAcr - Update ACR
func (s *AssetsService) BulkUpdateAcr(ctx context.Context, req []AssetsServiceBulkUpdateAcrRequestItem) error {
	_, err := s.client.post(ctx, "/api/v2/assets/bulk-jobs/acr", req)
	if err != nil {
		return err
	}
	return nil
}

// BulkMove - Move assets
func (s *AssetsService) BulkMove(ctx context.Context, req *AssetsServiceBulkMoveRequest) error {
	_, err := s.client.post(ctx, "/api/v2/assets/bulk-jobs/move-to-network", req)
	if err != nil {
		return err
	}
	return nil
}

// BulkDelete - Bulk delete assets
func (s *AssetsService) BulkDelete(ctx context.Context, req *AssetsServiceBulkDeleteRequest) error {
	_, err := s.client.post(ctx, "/api/v2/assets/bulk-jobs/delete", req)
	if err != nil {
		return err
	}
	return nil
}

// Import - Import assets
func (s *AssetsService) Import(ctx context.Context, req *AssetsServiceImportRequest) (*AssetsServiceImportResponse, error) {
	resp, err := s.client.post(ctx, "/import/assets", req)
	if err != nil {
		return nil, err
	}
	var result AssetsServiceImportResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListImportJobs - List asset import jobs
func (s *AssetsService) ListImportJobs(ctx context.Context) (*AssetsServiceListImportJobsResponse, error) {
	resp, err := s.client.get(ctx, "/import/asset-jobs")
	if err != nil {
		return nil, err
	}
	var result AssetsServiceListImportJobsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ImportJobInfo - Get import job status
func (s *AssetsService) ImportJobInfo(ctx context.Context, assetImportJobUuid string) (*AssetsServiceImportJobInfoResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/import/asset-jobs/%s", assetImportJobUuid))
	if err != nil {
		return nil, err
	}
	var result AssetsServiceImportJobInfoResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
