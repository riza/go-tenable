package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// ScanResultsService handles communication with the Scan Results related endpoints of the VM API.
type ScanResultsService struct {
	client *Client
}

type ScanResultsServiceHostDetailsResponseComplianceItem struct {
	HostId        int    `json:"host_id,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	PluginId      int    `json:"plugin_id,omitempty"`
	PluginName    string `json:"plugin_name,omitempty"`
	PluginFamily  string `json:"plugin_family,omitempty"`
	Count         int    `json:"count,omitempty"`
	SeverityIndex int    `json:"severity_index,omitempty"`
	Severity      int    `json:"severity,omitempty"`
}

type ScanResultsServiceHostDetailsResponseVulnerabilitiesItem struct {
	HostId        int    `json:"host_id,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	PluginId      int    `json:"plugin_id,omitempty"`
	PluginName    string `json:"plugin_name,omitempty"`
	PluginFamily  string `json:"plugin_family,omitempty"`
	Count         int    `json:"count,omitempty"`
	VulnIndex     int    `json:"vuln_index,omitempty"`
	SeverityIndex int    `json:"severity_index,omitempty"`
	Severity      int    `json:"severity,omitempty"`
}

type ScanResultsServiceHostDetailsResponse struct {
	Info            map[string]interface{}                                     `json:"info,omitempty"`
	Compliance      []ScanResultsServiceHostDetailsResponseComplianceItem      `json:"compliance,omitempty"`
	Vulnerabilities []ScanResultsServiceHostDetailsResponseVulnerabilitiesItem `json:"vulnerabilities,omitempty"`
}

type ScanResultsServicePluginOutputResponseOutputItemPorts struct {
}

type ScanResultsServicePluginOutputResponseOutputItem struct {
	Ports             ScanResultsServicePluginOutputResponseOutputItemPorts `json:"ports,omitempty"`
	HasAttachment     int                                                   `json:"has_attachment,omitempty"`
	CustomDescription string                                                `json:"custom_description,omitempty"`
	PluginOutput      string                                                `json:"plugin_output,omitempty"`
	Hosts             string                                                `json:"hosts,omitempty"`
	Severity          int                                                   `json:"severity,omitempty"`
}

type ScanResultsServicePluginOutputResponseInfo struct {
	HostFqdn    string `json:"host-fqdn,omitempty"`
	HostFqdnAlt string `json:"host_fqdn,omitempty"`
	HostIp      string `json:"host-ip,omitempty"`
	HostUuid    string `json:"host-uuid,omitempty"`
	HostStart   string `json:"host_start,omitempty"`
	HostEnd     string `json:"host_end,omitempty"`
	MacAddress  string `json:"mac-address,omitempty"`
}

type ScanResultsServicePluginOutputResponse struct {
	Output []ScanResultsServicePluginOutputResponseOutputItem `json:"output,omitempty"`
	Info   ScanResultsServicePluginOutputResponseInfo         `json:"info,omitempty"`
}

// HostDetails - Get host details
func (s *ScanResultsService) HostDetails(ctx context.Context, scanUuid string, hostId string) (*ScanResultsServiceHostDetailsResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scans/%s/hosts/%s", scanUuid, hostId))
	if err != nil {
		return nil, err
	}
	var result ScanResultsServiceHostDetailsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PluginOutput - Get plugin output
func (s *ScanResultsService) PluginOutput(ctx context.Context, scanUuid string, hostId string, pluginId string) (*ScanResultsServicePluginOutputResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scans/%s/hosts/%s/plugins/%s", scanUuid, hostId, pluginId))
	if err != nil {
		return nil, err
	}
	var result ScanResultsServicePluginOutputResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Attachments - Get scan attachment file
func (s *ScanResultsService) Attachments(ctx context.Context, scanId string, attachmentId string) error {
	_, err := s.client.get(ctx, fmt.Sprintf("/scans/%s/attachments/%s", scanId, attachmentId))
	if err != nil {
		return err
	}
	return nil
}
