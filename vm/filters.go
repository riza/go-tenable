package vm

import (
	"context"
	"encoding/json"
)

// FiltersService handles communication with the Filters related endpoints of the VM API.
type FiltersService struct {
	client *Client
}

type FiltersServiceFiltersAgentsListResponseFiltersItemControl struct {
	ReadableRegex string `json:"readable_regex,omitempty"`
	TypeField     string `json:"type,omitempty"`
	Regex         string `json:"regex,omitempty"`
}

type FiltersServiceFiltersAgentsListResponseFiltersItem struct {
	Name         string                                                    `json:"name,omitempty"`
	ReadableName string                                                    `json:"readable_name,omitempty"`
	Operators    []string                                                  `json:"operators,omitempty"`
	Control      FiltersServiceFiltersAgentsListResponseFiltersItemControl `json:"control,omitempty"`
	GroupName    string                                                    `json:"group_name,omitempty"`
}

type FiltersServiceFiltersAgentsListResponseSort struct {
	MaxSortFields  int      `json:"max_sort_fields,omitempty"`
	SortableFields []string `json:"sortable_fields,omitempty"`
}

type FiltersServiceFiltersAgentsListResponse struct {
	WildcardFields []string                                             `json:"wildcard_fields,omitempty"`
	Filters        []FiltersServiceFiltersAgentsListResponseFiltersItem `json:"filters,omitempty"`
	Sort           FiltersServiceFiltersAgentsListResponseSort          `json:"sort,omitempty"`
}

type FiltersServiceFiltersAssetsListResponseFiltersItemControl struct {
	ReadableRegex string `json:"readable_regex,omitempty"`
	TypeField     string `json:"type,omitempty"`
	Regex         string `json:"regex,omitempty"`
}

type FiltersServiceFiltersAssetsListResponseFiltersItem struct {
	Name         string                                                    `json:"name,omitempty"`
	ReadableName string                                                    `json:"readable_name,omitempty"`
	Operators    []string                                                  `json:"operators,omitempty"`
	Control      FiltersServiceFiltersAssetsListResponseFiltersItemControl `json:"control,omitempty"`
	GroupName    string                                                    `json:"group_name,omitempty"`
}

type FiltersServiceFiltersAssetsListResponseSort struct {
	MaxSortFields  int      `json:"max_sort_fields,omitempty"`
	SortableFields []string `json:"sortable_fields,omitempty"`
}

type FiltersServiceFiltersAssetsListResponse struct {
	WildcardFields []string                                             `json:"wildcard_fields,omitempty"`
	Filters        []FiltersServiceFiltersAssetsListResponseFiltersItem `json:"filters,omitempty"`
	Sort           FiltersServiceFiltersAssetsListResponseSort          `json:"sort,omitempty"`
}

type FiltersServiceFiltersAssetsListV2Request struct {
	TagUuids []string `json:"tag_uuids,omitempty"`
}

type FiltersServiceFiltersAssetsListV2ResponseFiltersItemControl struct {
	ReadableRegex string `json:"readable_regex,omitempty"`
	TypeField     string `json:"type,omitempty"`
	Regex         string `json:"regex,omitempty"`
}

type FiltersServiceFiltersAssetsListV2ResponseFiltersItem struct {
	Name         string                                                      `json:"name,omitempty"`
	ReadableName string                                                      `json:"readable_name,omitempty"`
	Operators    []string                                                    `json:"operators,omitempty"`
	Control      FiltersServiceFiltersAssetsListV2ResponseFiltersItemControl `json:"control,omitempty"`
	GroupName    string                                                      `json:"group_name,omitempty"`
}

type FiltersServiceFiltersAssetsListV2ResponseSort struct {
	MaxSortFields  int      `json:"max_sort_fields,omitempty"`
	SortableFields []string `json:"sortable_fields,omitempty"`
}

type FiltersServiceFiltersAssetsListV2Response struct {
	WildcardFields []string                                               `json:"wildcard_fields,omitempty"`
	Filters        []FiltersServiceFiltersAssetsListV2ResponseFiltersItem `json:"filters,omitempty"`
	Sort           FiltersServiceFiltersAssetsListV2ResponseSort          `json:"sort,omitempty"`
}

type FiltersServiceFiltersCredentialsListResponseFiltersItemControl struct {
	ReadableRegex string `json:"readable_regex,omitempty"`
	TypeField     string `json:"type,omitempty"`
	Regex         string `json:"regex,omitempty"`
}

type FiltersServiceFiltersCredentialsListResponseFiltersItem struct {
	Name         string                                                         `json:"name,omitempty"`
	ReadableName string                                                         `json:"readable_name,omitempty"`
	Operators    []string                                                       `json:"operators,omitempty"`
	Control      FiltersServiceFiltersCredentialsListResponseFiltersItemControl `json:"control,omitempty"`
	GroupName    string                                                         `json:"group_name,omitempty"`
}

type FiltersServiceFiltersCredentialsListResponseSort struct {
	MaxSortFields  int      `json:"max_sort_fields,omitempty"`
	SortableFields []string `json:"sortable_fields,omitempty"`
}

type FiltersServiceFiltersCredentialsListResponse struct {
	WildcardFields []string                                                  `json:"wildcard_fields,omitempty"`
	Filters        []FiltersServiceFiltersCredentialsListResponseFiltersItem `json:"filters,omitempty"`
	Sort           FiltersServiceFiltersCredentialsListResponseSort          `json:"sort,omitempty"`
}

type FiltersServiceVmFiltersReportsListResponseFilters struct {
	Name          string   `json:"name,omitempty"`
	Operators     []string `json:"operators,omitempty"`
	TypeField     string   `json:"type,omitempty"`
	AllowedValues []string `json:"allowedValues,omitempty"`
}

type FiltersServiceVmFiltersReportsListResponse struct {
	Filters FiltersServiceVmFiltersReportsListResponseFilters `json:"filters,omitempty"`
}

type FiltersServiceFiltersScanListResponseFiltersItemControl struct {
	ReadableRegex string `json:"readable_regex,omitempty"`
	TypeField     string `json:"type,omitempty"`
	Regex         string `json:"regex,omitempty"`
}

type FiltersServiceFiltersScanListResponseFiltersItem struct {
	Name         string                                                  `json:"name,omitempty"`
	ReadableName string                                                  `json:"readable_name,omitempty"`
	Operators    []string                                                `json:"operators,omitempty"`
	Control      FiltersServiceFiltersScanListResponseFiltersItemControl `json:"control,omitempty"`
	GroupName    string                                                  `json:"group_name,omitempty"`
}

type FiltersServiceFiltersScanListResponse struct {
	Filters []FiltersServiceFiltersScanListResponseFiltersItem `json:"filters,omitempty"`
}

type FiltersServiceFiltersScanHistoryListResponseFiltersItemControl struct {
	ReadableRegex string `json:"readable_regex,omitempty"`
	TypeField     string `json:"type,omitempty"`
	Regex         string `json:"regex,omitempty"`
}

type FiltersServiceFiltersScanHistoryListResponseFiltersItem struct {
	Name         string                                                         `json:"name,omitempty"`
	ReadableName string                                                         `json:"readable_name,omitempty"`
	Operators    []string                                                       `json:"operators,omitempty"`
	Control      FiltersServiceFiltersScanHistoryListResponseFiltersItemControl `json:"control,omitempty"`
	GroupName    string                                                         `json:"group_name,omitempty"`
}

type FiltersServiceFiltersScanHistoryListResponseSort struct {
	MaxSortFields  int      `json:"max_sort_fields,omitempty"`
	SortableFields []string `json:"sortable_fields,omitempty"`
}

type FiltersServiceFiltersScanHistoryListResponse struct {
	WildcardFields []string                                                  `json:"wildcard_fields,omitempty"`
	Filters        []FiltersServiceFiltersScanHistoryListResponseFiltersItem `json:"filters,omitempty"`
	Sort           FiltersServiceFiltersScanHistoryListResponseSort          `json:"sort,omitempty"`
}

type FiltersServiceFiltersVulnerabilitiesWorkbenchListResponseFiltersItemControl struct {
	ReadableRegex string `json:"readable_regex,omitempty"`
	TypeField     string `json:"type,omitempty"`
	Regex         string `json:"regex,omitempty"`
}

type FiltersServiceFiltersVulnerabilitiesWorkbenchListResponseFiltersItem struct {
	Name         string                                                                      `json:"name,omitempty"`
	ReadableName string                                                                      `json:"readable_name,omitempty"`
	Operators    []string                                                                    `json:"operators,omitempty"`
	Control      FiltersServiceFiltersVulnerabilitiesWorkbenchListResponseFiltersItemControl `json:"control,omitempty"`
	GroupName    string                                                                      `json:"group_name,omitempty"`
}

type FiltersServiceFiltersVulnerabilitiesWorkbenchListResponseSort struct {
	MaxSortFields  int      `json:"max_sort_fields,omitempty"`
	SortableFields []string `json:"sortable_fields,omitempty"`
}

type FiltersServiceFiltersVulnerabilitiesWorkbenchListResponse struct {
	WildcardFields []string                                                               `json:"wildcard_fields,omitempty"`
	Filters        []FiltersServiceFiltersVulnerabilitiesWorkbenchListResponseFiltersItem `json:"filters,omitempty"`
	Sort           FiltersServiceFiltersVulnerabilitiesWorkbenchListResponseSort          `json:"sort,omitempty"`
}

type FiltersServiceFiltersVulnerabilitiesWorkbenchListV2Request struct {
	TagUuids []string `json:"tag_uuids,omitempty"`
}

type FiltersServiceFiltersVulnerabilitiesWorkbenchListV2ResponseFiltersItemControl struct {
	ReadableRegex string `json:"readable_regex,omitempty"`
	TypeField     string `json:"type,omitempty"`
	Regex         string `json:"regex,omitempty"`
}

type FiltersServiceFiltersVulnerabilitiesWorkbenchListV2ResponseFiltersItem struct {
	Name         string                                                                        `json:"name,omitempty"`
	ReadableName string                                                                        `json:"readable_name,omitempty"`
	Operators    []string                                                                      `json:"operators,omitempty"`
	Control      FiltersServiceFiltersVulnerabilitiesWorkbenchListV2ResponseFiltersItemControl `json:"control,omitempty"`
	GroupName    string                                                                        `json:"group_name,omitempty"`
}

type FiltersServiceFiltersVulnerabilitiesWorkbenchListV2ResponseSort struct {
	MaxSortFields  int      `json:"max_sort_fields,omitempty"`
	SortableFields []string `json:"sortable_fields,omitempty"`
}

type FiltersServiceFiltersVulnerabilitiesWorkbenchListV2Response struct {
	WildcardFields []string                                                                 `json:"wildcard_fields,omitempty"`
	Filters        []FiltersServiceFiltersVulnerabilitiesWorkbenchListV2ResponseFiltersItem `json:"filters,omitempty"`
	Sort           FiltersServiceFiltersVulnerabilitiesWorkbenchListV2ResponseSort          `json:"sort,omitempty"`
}

// FiltersAgentsList - List agent filters
func (s *FiltersService) FiltersAgentsList(ctx context.Context) (*FiltersServiceFiltersAgentsListResponse, error) {
	resp, err := s.client.get(ctx, "/filters/scans/agents")
	if err != nil {
		return nil, err
	}
	var result FiltersServiceFiltersAgentsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FiltersAssetsList - List asset filters
func (s *FiltersService) FiltersAssetsList(ctx context.Context) (*FiltersServiceFiltersAssetsListResponse, error) {
	resp, err := s.client.get(ctx, "/filters/workbenches/assets")
	if err != nil {
		return nil, err
	}
	var result FiltersServiceFiltersAssetsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FiltersAssetsListV2 - List asset filters v2
func (s *FiltersService) FiltersAssetsListV2(ctx context.Context, req *FiltersServiceFiltersAssetsListV2Request) (*FiltersServiceFiltersAssetsListV2Response, error) {
	resp, err := s.client.post(ctx, "/filters/workbenches/assets", req)
	if err != nil {
		return nil, err
	}
	var result FiltersServiceFiltersAssetsListV2Response
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FiltersCredentialsList - List credential filters
func (s *FiltersService) FiltersCredentialsList(ctx context.Context) (*FiltersServiceFiltersCredentialsListResponse, error) {
	resp, err := s.client.get(ctx, "/filters/credentials")
	if err != nil {
		return nil, err
	}
	var result FiltersServiceFiltersCredentialsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// VmFiltersReportsList - List report filters
func (s *FiltersService) VmFiltersReportsList(ctx context.Context) (*FiltersServiceVmFiltersReportsListResponse, error) {
	resp, err := s.client.get(ctx, "/filters/reports/export")
	if err != nil {
		return nil, err
	}
	var result FiltersServiceVmFiltersReportsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FiltersScanList - List scan filters
func (s *FiltersService) FiltersScanList(ctx context.Context) (*FiltersServiceFiltersScanListResponse, error) {
	resp, err := s.client.get(ctx, "/filters/scans/reports")
	if err != nil {
		return nil, err
	}
	var result FiltersServiceFiltersScanListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FiltersScanHistoryList - List scan history filters
func (s *FiltersService) FiltersScanHistoryList(ctx context.Context) (*FiltersServiceFiltersScanHistoryListResponse, error) {
	resp, err := s.client.get(ctx, "/filters/scans/reports/history")
	if err != nil {
		return nil, err
	}
	var result FiltersServiceFiltersScanHistoryListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FiltersVulnerabilitiesWorkbenchList - List vulnerability filters
func (s *FiltersService) FiltersVulnerabilitiesWorkbenchList(ctx context.Context) (*FiltersServiceFiltersVulnerabilitiesWorkbenchListResponse, error) {
	resp, err := s.client.get(ctx, "/filters/workbenches/vulnerabilities")
	if err != nil {
		return nil, err
	}
	var result FiltersServiceFiltersVulnerabilitiesWorkbenchListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FiltersVulnerabilitiesWorkbenchListV2 - List vulnerability filters v2
func (s *FiltersService) FiltersVulnerabilitiesWorkbenchListV2(ctx context.Context, req *FiltersServiceFiltersVulnerabilitiesWorkbenchListV2Request) (*FiltersServiceFiltersVulnerabilitiesWorkbenchListV2Response, error) {
	resp, err := s.client.post(ctx, "/filters/workbenches/vulnerabilities", req)
	if err != nil {
		return nil, err
	}
	var result FiltersServiceFiltersVulnerabilitiesWorkbenchListV2Response
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
