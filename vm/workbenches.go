package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// WorkbenchesService handles communication with the Workbenches related endpoints of the VM API.
type WorkbenchesService struct {
	client *Client
}

type WorkbenchesServiceWorkbenchesVulnerabilitiesResponseVulnerabilitiesItemCountsBySeverityItem struct {
	Count int `json:"count,omitempty"`
	Value int `json:"value,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilitiesResponseVulnerabilitiesItem struct {
	Count              int                                                                                           `json:"count,omitempty"`
	PluginFamily       string                                                                                        `json:"plugin_family,omitempty"`
	PluginId           int                                                                                           `json:"plugin_id,omitempty"`
	PluginName         string                                                                                        `json:"plugin_name,omitempty"`
	VulnerabilityState string                                                                                        `json:"vulnerability_state,omitempty"`
	VprScore           float64                                                                                       `json:"vpr_score,omitempty"`
	Severity           int                                                                                           `json:"severity,omitempty"`
	AcceptedCount      int                                                                                           `json:"accepted_count,omitempty"`
	RecastedCount      int                                                                                           `json:"recasted_count,omitempty"`
	CountsBySeverity   []WorkbenchesServiceWorkbenchesVulnerabilitiesResponseVulnerabilitiesItemCountsBySeverityItem `json:"counts by severity,omitempty"`
	CvssBaseScore      float64                                                                                       `json:"cvss_base_score,omitempty"`
	Cvss3BaseScore     float64                                                                                       `json:"cvss3_base_score,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilitiesResponse struct {
	Vulnerabilities         []WorkbenchesServiceWorkbenchesVulnerabilitiesResponseVulnerabilitiesItem `json:"vulnerabilities,omitempty"`
	TotalVulnerabilityCount int                                                                       `json:"total_vulnerability_count,omitempty"`
	TotalAssetCount         int                                                                       `json:"total_asset_count,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseDiscovery struct {
	SeenFirst string `json:"seen_first,omitempty"`
	SeenLast  string `json:"seen_last,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityInfoResponsePluginDetails struct {
	Family           string `json:"family,omitempty"`
	ModificationDate string `json:"modification_date,omitempty"`
	Name             string `json:"name,omitempty"`
	PublicationDate  string `json:"publication_date,omitempty"`
	TypeField        string `json:"type,omitempty"`
	Version          string `json:"version,omitempty"`
	Severity         int    `json:"severity,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseReferenceInformationItem struct {
	Name   string   `json:"name,omitempty"`
	Url    string   `json:"url,omitempty"`
	Values []string `json:"values,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseRiskInformation struct {
	RiskFactor          string `json:"risk_factor,omitempty"`
	CvssVector          string `json:"cvss_vector,omitempty"`
	CvssBaseScore       string `json:"cvss_base_score,omitempty"`
	CvssTemporalVector  string `json:"cvss_temporal_vector,omitempty"`
	CvssTemporalScore   string `json:"cvss_temporal_score,omitempty"`
	Cvss3Vector         string `json:"cvss3_vector,omitempty"`
	Cvss3BaseScore      string `json:"cvss3_base_score,omitempty"`
	Cvss3TemporalVector string `json:"cvss3_temporal_vector,omitempty"`
	Cvss3TemporalScore  string `json:"cvss3_temporal_score,omitempty"`
	StigSeverity        string `json:"stig_severity,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseVulnerabilityInformationExploitFrameworksItemExploitsItem struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseVulnerabilityInformationExploitFrameworksItem struct {
	Name     string                                                                                                            `json:"name,omitempty"`
	Exploits []WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseVulnerabilityInformationExploitFrameworksItemExploitsItem `json:"exploits,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseVulnerabilityInformation struct {
	VulnerabilityPublicationDate string                                                                                                `json:"vulnerability_publication_date,omitempty"`
	ExploitedByMalware           bool                                                                                                  `json:"exploited_by_malware,omitempty"`
	PatchPublicationDate         string                                                                                                `json:"patch_publication_date,omitempty"`
	ExploitAvailable             bool                                                                                                  `json:"exploit_available,omitempty"`
	ExploitabilityEase           string                                                                                                `json:"exploitability_ease,omitempty"`
	AssetInventory               string                                                                                                `json:"asset_inventory,omitempty"`
	DefaultAccount               string                                                                                                `json:"default_account,omitempty"`
	ExploitedByNessus            bool                                                                                                  `json:"exploited_by_nessus,omitempty"`
	InTheNews                    bool                                                                                                  `json:"in_the_news,omitempty"`
	Malware                      string                                                                                                `json:"malware,omitempty"`
	UnsupportedByVendor          bool                                                                                                  `json:"unsupported_by_vendor,omitempty"`
	Cpe                          string                                                                                                `json:"cpe,omitempty"`
	ExploitFrameworks            []WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseVulnerabilityInformationExploitFrameworksItem `json:"exploit_frameworks,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseVprDrivers struct {
}

type WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseVpr struct {
	Score   float64                                                          `json:"score,omitempty"`
	Drivers WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseVprDrivers `json:"drivers,omitempty"`
	Updated string                                                           `json:"updated,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityInfoResponse struct {
	Count                    int                                                                              `json:"count,omitempty"`
	VulnCount                int                                                                              `json:"vuln_count,omitempty"`
	Description              string                                                                           `json:"description,omitempty"`
	Synopsis                 string                                                                           `json:"synopsis,omitempty"`
	Solution                 string                                                                           `json:"solution,omitempty"`
	Discovery                WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseDiscovery                  `json:"discovery,omitempty"`
	Severity                 int                                                                              `json:"severity,omitempty"`
	PluginDetails            WorkbenchesServiceWorkbenchesVulnerabilityInfoResponsePluginDetails              `json:"plugin_details,omitempty"`
	ReferenceInformation     []WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseReferenceInformationItem `json:"reference_information,omitempty"`
	RiskInformation          WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseRiskInformation            `json:"risk_information,omitempty"`
	SeeAlso                  []string                                                                         `json:"see_also,omitempty"`
	VulnerabilityInformation WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseVulnerabilityInformation   `json:"vulnerability_information,omitempty"`
	Vpr                      WorkbenchesServiceWorkbenchesVulnerabilityInfoResponseVpr                        `json:"vpr,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityOutputResponseItemStatesItemResultsItemAssetsItem struct {
	Hostname    string `json:"hostname,omitempty"`
	Id          string `json:"id,omitempty"`
	Uuid        string `json:"uuid,omitempty"`
	NetbiosName string `json:"netbios_name,omitempty"`
	Fqdn        string `json:"fqdn,omitempty"`
	Ipv4        string `json:"ipv4,omitempty"`
	FirstSeen   string `json:"first_seen,omitempty"`
	LastSeen    string `json:"last_seen,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityOutputResponseItemStatesItemResultsItem struct {
	ApplicationProtocol string                                                                                        `json:"application_protocol,omitempty"`
	Port                int                                                                                           `json:"port,omitempty"`
	TransportProtocol   string                                                                                        `json:"transport_protocol,omitempty"`
	Assets              []WorkbenchesServiceWorkbenchesVulnerabilityOutputResponseItemStatesItemResultsItemAssetsItem `json:"assets,omitempty"`
	Severity            int                                                                                           `json:"severity,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityOutputResponseItemStatesItem struct {
	Name    string                                                                              `json:"name,omitempty"`
	Results []WorkbenchesServiceWorkbenchesVulnerabilityOutputResponseItemStatesItemResultsItem `json:"results,omitempty"`
}

type WorkbenchesServiceWorkbenchesVulnerabilityOutputResponseItem struct {
	PluginOutput string                                                                   `json:"plugin_output,omitempty"`
	States       []WorkbenchesServiceWorkbenchesVulnerabilityOutputResponseItemStatesItem `json:"states,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetsResponseAssetsItemSourcesItem struct {
	Name      string `json:"name,omitempty"`
	FirstSeen string `json:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetsResponseAssetsItemAcrDriversItem struct {
	DriverName  string   `json:"driver_name,omitempty"`
	DriverValue []string `json:"driver_value,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetsResponseAssetsItemScanFrequencyItem struct {
	Interval  int  `json:"interval,omitempty"`
	Frequency int  `json:"frequency,omitempty"`
	Licensed  bool `json:"licensed,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetsResponseAssetsItem struct {
	Id              string                                                                   `json:"id,omitempty"`
	HasAgent        bool                                                                     `json:"has_agent,omitempty"`
	LastSeen        string                                                                   `json:"last_seen,omitempty"`
	LastScanTarget  string                                                                   `json:"last_scan_target,omitempty"`
	Sources         []WorkbenchesServiceWorkbenchesAssetsResponseAssetsItemSourcesItem       `json:"sources,omitempty"`
	AcrScore        int                                                                      `json:"acr_score,omitempty"`
	AcrDrivers      []WorkbenchesServiceWorkbenchesAssetsResponseAssetsItemAcrDriversItem    `json:"acr_drivers,omitempty"`
	ExposureScore   int                                                                      `json:"exposure_score,omitempty"`
	ScanFrequency   []WorkbenchesServiceWorkbenchesAssetsResponseAssetsItemScanFrequencyItem `json:"scan_frequency,omitempty"`
	Ipv4            []string                                                                 `json:"ipv4,omitempty"`
	Ipv6            []string                                                                 `json:"ipv6,omitempty"`
	Fqdn            []string                                                                 `json:"fqdn,omitempty"`
	NetbiosName     []string                                                                 `json:"netbios_name,omitempty"`
	OperatingSystem []string                                                                 `json:"operating_system,omitempty"`
	AgentName       []string                                                                 `json:"agent_name,omitempty"`
	AwsEc2Name      []string                                                                 `json:"aws_ec2_name,omitempty"`
	MacAddress      []string                                                                 `json:"mac_address,omitempty"`
	BigfixAssetId   []string                                                                 `json:"bigfix_asset_id,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetsResponse struct {
	Assets []WorkbenchesServiceWorkbenchesAssetsResponseAssetsItem `json:"assets,omitempty"`
	Total  int                                                     `json:"total,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetsVulnerabilitiesResponseItemSeveritiesItem struct {
	Count int    `json:"count,omitempty"`
	Level int    `json:"level,omitempty"`
	Name  string `json:"name,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetsVulnerabilitiesResponseItem struct {
	Id          string                                                                         `json:"id,omitempty"`
	Severities  []WorkbenchesServiceWorkbenchesAssetsVulnerabilitiesResponseItemSeveritiesItem `json:"severities,omitempty"`
	Total       int                                                                            `json:"total,omitempty"`
	Fqdn        []string                                                                       `json:"fqdn,omitempty"`
	Ipv4        []string                                                                       `json:"ipv4,omitempty"`
	Ipv6        []string                                                                       `json:"ipv6,omitempty"`
	LastSeen    string                                                                         `json:"last_seen,omitempty"`
	NetbiosName []string                                                                       `json:"netbios_name,omitempty"`
	AgentName   []string                                                                       `json:"agent_name,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetInfoResponseInfoCounts struct {
}

type WorkbenchesServiceWorkbenchesAssetInfoResponseInfoSourcesItem struct {
	Name      string `json:"name,omitempty"`
	FirstSeen string `json:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetInfoResponseInfoTagsItem struct {
	TagUuid  string `json:"tag_uuid,omitempty"`
	TagKey   string `json:"tag_key,omitempty"`
	TagValue string `json:"tag_value,omitempty"`
	AddedBy  string `json:"added_by,omitempty"`
	AddedAt  string `json:"added_at,omitempty"`
	Source   string `json:"source,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetInfoResponseInfoAcrDriversItem struct {
	DriverName  string   `json:"driver_name,omitempty"`
	DriverValue []string `json:"driver_value,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetInfoResponseInfoScanFrequencyItem struct {
	Interval  int  `json:"interval,omitempty"`
	Frequency int  `json:"frequency,omitempty"`
	Licensed  bool `json:"licensed,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetInfoResponseInfo struct {
	TimeEnd                   string                                                                `json:"time_end,omitempty"`
	TimeStart                 string                                                                `json:"time_start,omitempty"`
	Id                        string                                                                `json:"id,omitempty"`
	Uuid                      string                                                                `json:"uuid,omitempty"`
	OperatingSystem           []string                                                              `json:"operating_system,omitempty"`
	Counts                    WorkbenchesServiceWorkbenchesAssetInfoResponseInfoCounts              `json:"counts,omitempty"`
	HasAgent                  bool                                                                  `json:"has_agent,omitempty"`
	CreatedAt                 string                                                                `json:"created_at,omitempty"`
	UpdatedAt                 string                                                                `json:"updated_at,omitempty"`
	FirstSeen                 string                                                                `json:"first_seen,omitempty"`
	LastSeen                  string                                                                `json:"last_seen,omitempty"`
	LastAuthenticatedScanDate string                                                                `json:"last_authenticated_scan_date,omitempty"`
	LastLicensedScanDate      string                                                                `json:"last_licensed_scan_date,omitempty"`
	LastScanTarget            string                                                                `json:"last_scan_target,omitempty"`
	Sources                   []WorkbenchesServiceWorkbenchesAssetInfoResponseInfoSourcesItem       `json:"sources,omitempty"`
	Tags                      []WorkbenchesServiceWorkbenchesAssetInfoResponseInfoTagsItem          `json:"tags,omitempty"`
	AcrScore                  float64                                                               `json:"acr_score,omitempty"`
	AcrDrivers                []WorkbenchesServiceWorkbenchesAssetInfoResponseInfoAcrDriversItem    `json:"acr_drivers,omitempty"`
	ExposureScore             float64                                                               `json:"exposure_score,omitempty"`
	ScanFrequency             []WorkbenchesServiceWorkbenchesAssetInfoResponseInfoScanFrequencyItem `json:"scan_frequency,omitempty"`
	Ipv4                      []string                                                              `json:"ipv4,omitempty"`
	Ipv6                      []string                                                              `json:"ipv6,omitempty"`
	Fqdn                      []string                                                              `json:"fqdn,omitempty"`
	MacAddress                []string                                                              `json:"mac_address,omitempty"`
	NetbiosName               []string                                                              `json:"netbios_name,omitempty"`
	SystemType                []string                                                              `json:"system_type,omitempty"`
	TenableUuid               []string                                                              `json:"tenable_uuid,omitempty"`
	Hostname                  []string                                                              `json:"hostname,omitempty"`
	AgentName                 []string                                                              `json:"agent_name,omitempty"`
	BiosUuid                  []string                                                              `json:"bios_uuid,omitempty"`
	AwsEc2InstanceId          []string                                                              `json:"aws_ec2_instance_id,omitempty"`
	AwsEc2InstanceAmiId       []string                                                              `json:"aws_ec2_instance_ami_id,omitempty"`
	AwsOwnerId                []string                                                              `json:"aws_owner_id,omitempty"`
	AwsAvailabilityZone       []string                                                              `json:"aws_availability_zone,omitempty"`
	AwsRegion                 []string                                                              `json:"aws_region,omitempty"`
	AwsVpcId                  []string                                                              `json:"aws_vpc_id,omitempty"`
	AwsEc2InstanceGroupName   []string                                                              `json:"aws_ec2_instance_group_name,omitempty"`
	AwsEc2InstanceStateName   []string                                                              `json:"aws_ec2_instance_state_name,omitempty"`
	AwsEc2InstanceType        []string                                                              `json:"aws_ec2_instance_type,omitempty"`
	AwsSubnetId               []string                                                              `json:"aws_subnet_id,omitempty"`
	AwsEc2ProductCode         []string                                                              `json:"aws_ec2_product_code,omitempty"`
	AwsEc2Name                []string                                                              `json:"aws_ec2_name,omitempty"`
	AzureVmId                 []string                                                              `json:"azure_vm_id,omitempty"`
	AzureResourceId           []string                                                              `json:"azure_resource_id,omitempty"`
	GcpProjectId              []string                                                              `json:"gcp_project_id,omitempty"`
	GcpZone                   []string                                                              `json:"gcp_zone,omitempty"`
	GcpInstanceId             []string                                                              `json:"gcp_instance_id,omitempty"`
	SshFingerprint            []string                                                              `json:"ssh_fingerprint,omitempty"`
	McafeeEpoGuid             []string                                                              `json:"mcafee_epo_guid,omitempty"`
	McafeeEpoAgentGuid        []string                                                              `json:"mcafee_epo_agent_guid,omitempty"`
	QualysAssetId             []string                                                              `json:"qualys_asset_id,omitempty"`
	QualysHostId              []string                                                              `json:"qualys_host_id,omitempty"`
	ServicenowSysid           []string                                                              `json:"servicenow_sysid,omitempty"`
	InstalledSoftware         []string                                                              `json:"installed_software,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetInfoResponse struct {
	Info WorkbenchesServiceWorkbenchesAssetInfoResponseInfo `json:"info,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetsActivityResponseItemDetailsSourcesItem struct {
	Name      string `json:"name,omitempty"`
	FirstSeen string `json:"firstSeen,omitempty"`
	LastSeen  string `json:"lastSeen,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetsActivityResponseItemDetailsProperties struct {
}

type WorkbenchesServiceWorkbenchesAssetsActivityResponseItemDetails struct {
	AssetId                   string                                                                      `json:"assetId,omitempty"`
	ContainerId               string                                                                      `json:"containerId,omitempty"`
	CreatedAt                 int                                                                         `json:"createdAt,omitempty"`
	UpdatedAt                 int                                                                         `json:"updatedAt,omitempty"`
	HasAgent                  bool                                                                        `json:"hasAgent,omitempty"`
	HasPluginResults          bool                                                                        `json:"hasPluginResults,omitempty"`
	FirstScanTime             int                                                                         `json:"firstScanTime,omitempty"`
	LastScanTime              int                                                                         `json:"lastScanTime,omitempty"`
	LastAuthenticatedScanTime int                                                                         `json:"lastAuthenticatedScanTime,omitempty"`
	LastLicensedScanTime      int                                                                         `json:"lastLicensedScanTime,omitempty"`
	LastLicensedScanTimeV2    int                                                                         `json:"lastLicensedScanTimeV2,omitempty"`
	Sources                   []WorkbenchesServiceWorkbenchesAssetsActivityResponseItemDetailsSourcesItem `json:"sources,omitempty"`
	TerminatedAt              int                                                                         `json:"terminatedAt,omitempty"`
	TerminatedBy              string                                                                      `json:"terminatedBy,omitempty"`
	DeletedAt                 int                                                                         `json:"deletedAt,omitempty"`
	DeletedBy                 string                                                                      `json:"deletedBy,omitempty"`
	Properties                WorkbenchesServiceWorkbenchesAssetsActivityResponseItemDetailsProperties    `json:"properties,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetsActivityResponseItemUpdatesItem struct {
	Method   string `json:"method,omitempty"`
	Property string `json:"property,omitempty"`
	Value    string `json:"value,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetsActivityResponseItem struct {
	TypeField  string                                                               `json:"type,omitempty"`
	Timestamp  int                                                                  `json:"timestamp,omitempty"`
	ScanId     string                                                               `json:"scan_id,omitempty"`
	ScheduleId string                                                               `json:"schedule_id,omitempty"`
	Source     string                                                               `json:"source,omitempty"`
	Details    WorkbenchesServiceWorkbenchesAssetsActivityResponseItemDetails       `json:"details,omitempty"`
	Updates    []WorkbenchesServiceWorkbenchesAssetsActivityResponseItemUpdatesItem `json:"updates,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilitiesResponseVulnerabilitiesItemCountsBySeverityItem struct {
	Count int `json:"count,omitempty"`
	Value int `json:"value,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilitiesResponseVulnerabilitiesItem struct {
	Count              int                                                                                                `json:"count,omitempty"`
	PluginFamily       string                                                                                             `json:"plugin_family,omitempty"`
	PluginId           int                                                                                                `json:"plugin_id,omitempty"`
	PluginName         string                                                                                             `json:"plugin_name,omitempty"`
	VulnerabilityState string                                                                                             `json:"vulnerability_state,omitempty"`
	VprScore           float64                                                                                            `json:"vpr_score,omitempty"`
	Severity           int                                                                                                `json:"severity,omitempty"`
	AcceptedCount      int                                                                                                `json:"accepted_count,omitempty"`
	RecastedCount      int                                                                                                `json:"recasted_count,omitempty"`
	CountsBySeverity   []WorkbenchesServiceWorkbenchesAssetVulnerabilitiesResponseVulnerabilitiesItemCountsBySeverityItem `json:"counts by severity,omitempty"`
	CvssBaseScore      float64                                                                                            `json:"cvss_base_score,omitempty"`
	Cvss3BaseScore     float64                                                                                            `json:"cvss3_base_score,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilitiesResponse struct {
	Vulnerabilities         []WorkbenchesServiceWorkbenchesAssetVulnerabilitiesResponseVulnerabilitiesItem `json:"vulnerabilities,omitempty"`
	TotalVulnerabilityCount int                                                                            `json:"total_vulnerability_count,omitempty"`
	TotalAssetCount         int                                                                            `json:"total_asset_count,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoDiscovery struct {
	SeenFirst string `json:"seen_first,omitempty"`
	SeenLast  string `json:"seen_last,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoPluginDetails struct {
	Family           string `json:"family,omitempty"`
	ModificationDate string `json:"modification_date,omitempty"`
	Name             string `json:"name,omitempty"`
	PublicationDate  string `json:"publication_date,omitempty"`
	TypeField        string `json:"type,omitempty"`
	Version          string `json:"version,omitempty"`
	Severity         int    `json:"severity,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoReferenceInformationItem struct {
	Name   string   `json:"name,omitempty"`
	Url    string   `json:"url,omitempty"`
	Values []string `json:"values,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoRiskInformation struct {
	RiskFactor          string `json:"risk_factor,omitempty"`
	CvssVector          string `json:"cvss_vector,omitempty"`
	CvssBaseScore       string `json:"cvss_base_score,omitempty"`
	CvssTemporalVector  string `json:"cvss_temporal_vector,omitempty"`
	CvssTemporalScore   string `json:"cvss_temporal_score,omitempty"`
	Cvss3Vector         string `json:"cvss3_vector,omitempty"`
	Cvss3BaseScore      string `json:"cvss3_base_score,omitempty"`
	Cvss3TemporalVector string `json:"cvss3_temporal_vector,omitempty"`
	Cvss3TemporalScore  string `json:"cvss3_temporal_score,omitempty"`
	StigSeverity        string `json:"stig_severity,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoVulnerabilityInformationExploitFrameworksItemExploitsItem struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoVulnerabilityInformationExploitFrameworksItem struct {
	Name     string                                                                                                                     `json:"name,omitempty"`
	Exploits []WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoVulnerabilityInformationExploitFrameworksItemExploitsItem `json:"exploits,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoVulnerabilityInformation struct {
	VulnerabilityPublicationDate string                                                                                                         `json:"vulnerability_publication_date,omitempty"`
	ExploitedByMalware           bool                                                                                                           `json:"exploited_by_malware,omitempty"`
	PatchPublicationDate         string                                                                                                         `json:"patch_publication_date,omitempty"`
	ExploitAvailable             bool                                                                                                           `json:"exploit_available,omitempty"`
	ExploitabilityEase           string                                                                                                         `json:"exploitability_ease,omitempty"`
	AssetInventory               string                                                                                                         `json:"asset_inventory,omitempty"`
	DefaultAccount               string                                                                                                         `json:"default_account,omitempty"`
	ExploitedByNessus            bool                                                                                                           `json:"exploited_by_nessus,omitempty"`
	InTheNews                    bool                                                                                                           `json:"in_the_news,omitempty"`
	Malware                      string                                                                                                         `json:"malware,omitempty"`
	UnsupportedByVendor          bool                                                                                                           `json:"unsupported_by_vendor,omitempty"`
	Cpe                          string                                                                                                         `json:"cpe,omitempty"`
	ExploitFrameworks            []WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoVulnerabilityInformationExploitFrameworksItem `json:"exploit_frameworks,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoVprDrivers struct {
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoVpr struct {
	Score   float64                                                                   `json:"score,omitempty"`
	Drivers WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoVprDrivers `json:"drivers,omitempty"`
	Updated string                                                                    `json:"updated,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfo struct {
	Count                    int                                                                                       `json:"count,omitempty"`
	VulnCount                int                                                                                       `json:"vuln_count,omitempty"`
	Description              string                                                                                    `json:"description,omitempty"`
	Synopsis                 string                                                                                    `json:"synopsis,omitempty"`
	Solution                 string                                                                                    `json:"solution,omitempty"`
	Discovery                WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoDiscovery                  `json:"discovery,omitempty"`
	Severity                 int                                                                                       `json:"severity,omitempty"`
	PluginDetails            WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoPluginDetails              `json:"plugin_details,omitempty"`
	ReferenceInformation     []WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoReferenceInformationItem `json:"reference_information,omitempty"`
	RiskInformation          WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoRiskInformation            `json:"risk_information,omitempty"`
	SeeAlso                  []string                                                                                  `json:"see_also,omitempty"`
	VulnerabilityInformation WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoVulnerabilityInformation   `json:"vulnerability_information,omitempty"`
	Vpr                      WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfoVpr                        `json:"vpr,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponse struct {
	Info WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponseInfo `json:"info,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityOutputResponseOutputsItemStatesItemResultsItemAssetsItem struct {
	Hostname    string `json:"hostname,omitempty"`
	Id          string `json:"id,omitempty"`
	Uuid        string `json:"uuid,omitempty"`
	NetbiosName string `json:"netbios_name,omitempty"`
	Fqdn        string `json:"fqdn,omitempty"`
	Ipv4        string `json:"ipv4,omitempty"`
	FirstSeen   string `json:"first_seen,omitempty"`
	LastSeen    string `json:"last_seen,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityOutputResponseOutputsItemStatesItemResultsItem struct {
	ApplicationProtocol string                                                                                                    `json:"application_protocol,omitempty"`
	Port                int                                                                                                       `json:"port,omitempty"`
	TransportProtocol   string                                                                                                    `json:"transport_protocol,omitempty"`
	Assets              []WorkbenchesServiceWorkbenchesAssetVulnerabilityOutputResponseOutputsItemStatesItemResultsItemAssetsItem `json:"assets,omitempty"`
	Severity            int                                                                                                       `json:"severity,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityOutputResponseOutputsItemStatesItem struct {
	Name    string                                                                                          `json:"name,omitempty"`
	Results []WorkbenchesServiceWorkbenchesAssetVulnerabilityOutputResponseOutputsItemStatesItemResultsItem `json:"results,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityOutputResponseOutputsItem struct {
	PluginOutput string                                                                               `json:"plugin_output,omitempty"`
	States       []WorkbenchesServiceWorkbenchesAssetVulnerabilityOutputResponseOutputsItemStatesItem `json:"states,omitempty"`
}

type WorkbenchesServiceWorkbenchesAssetVulnerabilityOutputResponse struct {
	Outputs []WorkbenchesServiceWorkbenchesAssetVulnerabilityOutputResponseOutputsItem `json:"outputs,omitempty"`
}

type WorkbenchesServiceWorkbenchesExportRequestResponse struct {
	File int `json:"file,omitempty"`
}

type WorkbenchesServiceWorkbenchesExportStatusResponse struct {
	Status        string `json:"status,omitempty"`
	ProgressTotal string `json:"progress_total,omitempty"`
	Progress      string `json:"progress,omitempty"`
}

// WorkbenchesVulnerabilities - List vulnerabilities
func (s *WorkbenchesService) WorkbenchesVulnerabilities(ctx context.Context) (*WorkbenchesServiceWorkbenchesVulnerabilitiesResponse, error) {
	resp, err := s.client.get(ctx, "/workbenches/vulnerabilities")
	if err != nil {
		return nil, err
	}
	var result WorkbenchesServiceWorkbenchesVulnerabilitiesResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WorkbenchesVulnerabilityInfo - Get plugin details
func (s *WorkbenchesService) WorkbenchesVulnerabilityInfo(ctx context.Context, pluginId string) (*WorkbenchesServiceWorkbenchesVulnerabilityInfoResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/workbenches/vulnerabilities/%s/info", pluginId))
	if err != nil {
		return nil, err
	}
	var result WorkbenchesServiceWorkbenchesVulnerabilityInfoResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WorkbenchesVulnerabilityOutput - List plugin outputs
func (s *WorkbenchesService) WorkbenchesVulnerabilityOutput(ctx context.Context, pluginId string) ([]WorkbenchesServiceWorkbenchesVulnerabilityOutputResponseItem, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/workbenches/vulnerabilities/%s/outputs", pluginId))
	if err != nil {
		return nil, err
	}
	var result []WorkbenchesServiceWorkbenchesVulnerabilityOutputResponseItem
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// WorkbenchesAssets - List assets
func (s *WorkbenchesService) WorkbenchesAssets(ctx context.Context) (*WorkbenchesServiceWorkbenchesAssetsResponse, error) {
	resp, err := s.client.get(ctx, "/workbenches/assets")
	if err != nil {
		return nil, err
	}
	var result WorkbenchesServiceWorkbenchesAssetsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WorkbenchesAssetsVulnerabilities - List assets with vulnerabilities
func (s *WorkbenchesService) WorkbenchesAssetsVulnerabilities(ctx context.Context) ([]WorkbenchesServiceWorkbenchesAssetsVulnerabilitiesResponseItem, error) {
	resp, err := s.client.get(ctx, "/workbenches/assets/vulnerabilities")
	if err != nil {
		return nil, err
	}
	var result []WorkbenchesServiceWorkbenchesAssetsVulnerabilitiesResponseItem
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// WorkbenchesAssetInfo - Get asset information
func (s *WorkbenchesService) WorkbenchesAssetInfo(ctx context.Context, assetId string) (*WorkbenchesServiceWorkbenchesAssetInfoResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/workbenches/assets/%s/info", assetId))
	if err != nil {
		return nil, err
	}
	var result WorkbenchesServiceWorkbenchesAssetInfoResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WorkbenchesAssetsActivity - Get asset activity log
func (s *WorkbenchesService) WorkbenchesAssetsActivity(ctx context.Context, assetUuid string) ([]WorkbenchesServiceWorkbenchesAssetsActivityResponseItem, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/workbenches/assets/%s/activity", assetUuid))
	if err != nil {
		return nil, err
	}
	var result []WorkbenchesServiceWorkbenchesAssetsActivityResponseItem
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// WorkbenchesAssetVulnerabilities - List asset vulnerabilities
func (s *WorkbenchesService) WorkbenchesAssetVulnerabilities(ctx context.Context, assetId string) (*WorkbenchesServiceWorkbenchesAssetVulnerabilitiesResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/workbenches/assets/%s/vulnerabilities", assetId))
	if err != nil {
		return nil, err
	}
	var result WorkbenchesServiceWorkbenchesAssetVulnerabilitiesResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WorkbenchesAssetVulnerabilityInfo - Get asset vulnerability details
func (s *WorkbenchesService) WorkbenchesAssetVulnerabilityInfo(ctx context.Context, assetId string, pluginId string) (*WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/workbenches/assets/%s/vulnerabilities/%s/info", assetId, pluginId))
	if err != nil {
		return nil, err
	}
	var result WorkbenchesServiceWorkbenchesAssetVulnerabilityInfoResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WorkbenchesAssetVulnerabilityOutput - List asset vulnerabilities for plugin
func (s *WorkbenchesService) WorkbenchesAssetVulnerabilityOutput(ctx context.Context, assetId string, pluginId string) (*WorkbenchesServiceWorkbenchesAssetVulnerabilityOutputResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/workbenches/assets/%s/vulnerabilities/%s/outputs", assetId, pluginId))
	if err != nil {
		return nil, err
	}
	var result WorkbenchesServiceWorkbenchesAssetVulnerabilityOutputResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WorkbenchesAssetsDelete - Delete asset
func (s *WorkbenchesService) WorkbenchesAssetsDelete(ctx context.Context, assetUuid string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/workbenches/assets/%s", assetUuid))
	if err != nil {
		return err
	}
	return nil
}

// WorkbenchesExportRequest - Export workbench
func (s *WorkbenchesService) WorkbenchesExportRequest(ctx context.Context) (*WorkbenchesServiceWorkbenchesExportRequestResponse, error) {
	resp, err := s.client.get(ctx, "/workbenches/export")
	if err != nil {
		return nil, err
	}
	var result WorkbenchesServiceWorkbenchesExportRequestResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WorkbenchesExportStatus - Check export status
func (s *WorkbenchesService) WorkbenchesExportStatus(ctx context.Context, fileId string) (*WorkbenchesServiceWorkbenchesExportStatusResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/workbenches/export/%s/status", fileId))
	if err != nil {
		return nil, err
	}
	var result WorkbenchesServiceWorkbenchesExportStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WorkbenchesExportDownload - Download export file
func (s *WorkbenchesService) WorkbenchesExportDownload(ctx context.Context, fileId string) error {
	_, err := s.client.get(ctx, fmt.Sprintf("/workbenches/export/%s/download", fileId))
	if err != nil {
		return err
	}
	return nil
}
