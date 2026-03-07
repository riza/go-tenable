package vm

import (
	"context"
	"encoding/json"
)

// VulnerabilitiesService handles communication with the Vulnerabilities related endpoints of the VM API.
type VulnerabilitiesService struct {
	client *Client
}

type VulnerabilitiesServiceVulnerabilitiesImportRequestAssetsItemNetworkInterfacesItem struct {
	Ipv4        []string `json:"ipv4,omitempty"`
	Ipv6        []string `json:"ipv6,omitempty"`
	MacAddress  string   `json:"mac_address,omitempty"`
	NetbiosName string   `json:"netbios_name,omitempty"`
	Fqdn        string   `json:"fqdn,omitempty"`
}

type VulnerabilitiesServiceVulnerabilitiesImportRequestAssetsItemVulnerabilitiesItem struct {
	TenablePluginId string `json:"tenable_plugin_id,omitempty"`
	Cve             string `json:"cve,omitempty"`
	Port            int    `json:"port,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	Authenticated   bool   `json:"authenticated,omitempty"`
	FirstFound      int    `json:"first_found,omitempty"`
	LastFound       int    `json:"last_found,omitempty"`
	LastFixed       int    `json:"last_fixed,omitempty"`
	Output          string `json:"output,omitempty"`
}

type VulnerabilitiesServiceVulnerabilitiesImportRequestAssetsItem struct {
	NetworkInterfaces []VulnerabilitiesServiceVulnerabilitiesImportRequestAssetsItemNetworkInterfacesItem `json:"network_interfaces,omitempty"`
	Hostname          string                                                                              `json:"hostname,omitempty"`
	ServicenowSysid   string                                                                              `json:"servicenow_sysid,omitempty"`
	SshFingerprint    string                                                                              `json:"ssh_fingerprint,omitempty"`
	BiosUuid          string                                                                              `json:"bios_uuid,omitempty"`
	NetbiosName       string                                                                              `json:"netbios_name,omitempty"`
	TenableAgentId    int                                                                                 `json:"tenable_agent_id,omitempty"`
	Vulnerabilities   []VulnerabilitiesServiceVulnerabilitiesImportRequestAssetsItemVulnerabilitiesItem   `json:"vulnerabilities,omitempty"`
}

type VulnerabilitiesServiceVulnerabilitiesImportRequestChecksRanItem struct {
	TenablePluginId string `json:"tenable_plugin_id,omitempty"`
	Port            int    `json:"port,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
}

type VulnerabilitiesServiceVulnerabilitiesImportRequest struct {
	Source    string                                                            `json:"source,omitempty"`
	TypeField string                                                            `json:"type,omitempty"`
	Assets    []VulnerabilitiesServiceVulnerabilitiesImportRequestAssetsItem    `json:"assets,omitempty"`
	ChecksRan []VulnerabilitiesServiceVulnerabilitiesImportRequestChecksRanItem `json:"checks_ran,omitempty"`
}

type VulnerabilitiesServiceVulnerabilitiesImportV2RequestAssetsItemNetworkInterfacesItem struct {
	Ipv4        []string `json:"ipv4,omitempty"`
	Ipv6        []string `json:"ipv6,omitempty"`
	MacAddress  string   `json:"mac_address,omitempty"`
	NetbiosName string   `json:"netbios_name,omitempty"`
	Fqdn        string   `json:"fqdn,omitempty"`
}

type VulnerabilitiesServiceVulnerabilitiesImportV2RequestAssetsItemVulnerabilitiesItem struct {
	TenablePluginId string `json:"tenable_plugin_id,omitempty"`
	Cve             string `json:"cve,omitempty"`
	Port            int    `json:"port,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	Authenticated   bool   `json:"authenticated,omitempty"`
	LastFound       int    `json:"last_found,omitempty"`
	Output          string `json:"output,omitempty"`
}

type VulnerabilitiesServiceVulnerabilitiesImportV2RequestAssetsItem struct {
	NetworkInterfaces []VulnerabilitiesServiceVulnerabilitiesImportV2RequestAssetsItemNetworkInterfacesItem `json:"network_interfaces,omitempty"`
	Hostname          string                                                                                `json:"hostname,omitempty"`
	ServicenowSysid   string                                                                                `json:"servicenow_sysid,omitempty"`
	SshFingerprint    string                                                                                `json:"ssh_fingerprint,omitempty"`
	BiosUuid          string                                                                                `json:"bios_uuid,omitempty"`
	NetbiosName       string                                                                                `json:"netbios_name,omitempty"`
	OperatingSystems  string                                                                                `json:"operating_systems,omitempty"`
	Authenticated     bool                                                                                  `json:"authenticated,omitempty"`
	TenableAgentId    string                                                                                `json:"tenable_agent_id,omitempty"`
	TenableNetworkId  string                                                                                `json:"tenable_network_id,omitempty"`
	Vulnerabilities   []VulnerabilitiesServiceVulnerabilitiesImportV2RequestAssetsItemVulnerabilitiesItem   `json:"vulnerabilities,omitempty"`
}

type VulnerabilitiesServiceVulnerabilitiesImportV2RequestCoverage struct {
	Ids string `json:"ids,omitempty"`
}

type VulnerabilitiesServiceVulnerabilitiesImportV2Request struct {
	Vendor   string                                                           `json:"vendor,omitempty"`
	Product  string                                                           `json:"product,omitempty"`
	DataType string                                                           `json:"data_type,omitempty"`
	Source   string                                                           `json:"source,omitempty"`
	Assets   []VulnerabilitiesServiceVulnerabilitiesImportV2RequestAssetsItem `json:"assets,omitempty"`
	Coverage VulnerabilitiesServiceVulnerabilitiesImportV2RequestCoverage     `json:"coverage,omitempty"`
}

type VulnerabilitiesServiceVulnerabilitiesImportV2Response struct {
	JobUuid string `json:"job_uuid,omitempty"`
}

// VulnerabilitiesImport - Import vulnerabilities v1
func (s *VulnerabilitiesService) VulnerabilitiesImport(ctx context.Context, req *VulnerabilitiesServiceVulnerabilitiesImportRequest) error {
	_, err := s.client.post(ctx, "/import/vulnerabilities", req)
	if err != nil {
		return err
	}
	return nil
}

// VulnerabilitiesImportV2 - Import vulnerabilities v2
func (s *VulnerabilitiesService) VulnerabilitiesImportV2(ctx context.Context, req *VulnerabilitiesServiceVulnerabilitiesImportV2Request) (*VulnerabilitiesServiceVulnerabilitiesImportV2Response, error) {
	resp, err := s.client.post(ctx, "/api/v2/vulnerabilities", req)
	if err != nil {
		return nil, err
	}
	var result VulnerabilitiesServiceVulnerabilitiesImportV2Response
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
