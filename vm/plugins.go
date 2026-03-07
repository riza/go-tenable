package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// PluginsService handles communication with the Plugins related endpoints of the VM API.
type PluginsService struct {
	client *Client
}

type PluginsServicePluginsListResponseDataPluginDetailsItemAttributesItemCvssVector struct {
	AccessComplexity      string `json:"AccessComplexity,omitempty"`
	AccessVector          string `json:"AccessVector,omitempty"`
	Authentication        string `json:"Authentication,omitempty"`
	AvailabilityImpact    string `json:"Availability-Impact,omitempty"`
	ConfidentialityImpact string `json:"Confidentiality-Impact,omitempty"`
	IntegrityImpact       string `json:"Integrity-Impact,omitempty"`
}

type PluginsServicePluginsListResponseDataPluginDetailsItemAttributesItemVprDrivers struct {
}

type PluginsServicePluginsListResponseDataPluginDetailsItemAttributesItemVpr struct {
	Score   float64                                                                        `json:"score,omitempty"`
	Drivers PluginsServicePluginsListResponseDataPluginDetailsItemAttributesItemVprDrivers `json:"drivers,omitempty"`
	Updated string                                                                         `json:"updated,omitempty"`
}

type PluginsServicePluginsListResponseDataPluginDetailsItemAttributesItem struct {
	PluginModificationDate     string                                                                         `json:"plugin_modification_date,omitempty"`
	PluginVersion              string                                                                         `json:"plugin_version,omitempty"`
	ExploitedByMalware         bool                                                                           `json:"exploited_by_malware,omitempty"`
	Description                string                                                                         `json:"description,omitempty"`
	UnsupportedByVendor        bool                                                                           `json:"unsupported_by_vendor,omitempty"`
	CvssTemporalScore          float64                                                                        `json:"cvss_temporal_score,omitempty"`
	PatchPublicationDate       string                                                                         `json:"patch_publication_date,omitempty"`
	SeeAlso                    []string                                                                       `json:"see_also,omitempty"`
	DefaultAccount             string                                                                         `json:"default_account,omitempty"`
	ExploitAvailable           bool                                                                           `json:"exploit_available,omitempty"`
	Cve                        []string                                                                       `json:"cve,omitempty"`
	ExploitFrameworkCanvas     bool                                                                           `json:"exploit_framework_canvas,omitempty"`
	CvssBaseScore              string                                                                         `json:"cvss_base_score,omitempty"`
	Solution                   string                                                                         `json:"solution,omitempty"`
	CvssVector                 PluginsServicePluginsListResponseDataPluginDetailsItemAttributesItemCvssVector `json:"cvss_vector,omitempty"`
	ExploitFrameworkExploithub bool                                                                           `json:"exploit_framework_exploithub,omitempty"`
	Cpe                        []string                                                                       `json:"cpe,omitempty"`
	PluginPublicationDate      string                                                                         `json:"plugin_publication_date,omitempty"`
	ExploitFrameworkCore       bool                                                                           `json:"exploit_framework_core,omitempty"`
	InTheNews                  bool                                                                           `json:"in_the_news,omitempty"`
	HasPatch                   bool                                                                           `json:"has_patch,omitempty"`
	Xref                       []string                                                                       `json:"xref,omitempty"`
	Malware                    bool                                                                           `json:"malware,omitempty"`
	ExploitFrameworkD2Elliot   bool                                                                           `json:"exploit_framework_d2_elliot,omitempty"`
	Xrefs                      []string                                                                       `json:"xrefs,omitempty"`
	RiskFactor                 string                                                                         `json:"risk_factor,omitempty"`
	Synopsis                   string                                                                         `json:"synopsis,omitempty"`
	Cvss3TemporalScore         float64                                                                        `json:"cvss3_temporal_score,omitempty"`
	ExploitedByNessus          bool                                                                           `json:"exploited_by_nessus,omitempty"`
	Cvss3BaseScore             string                                                                         `json:"cvss3_base_score,omitempty"`
	ExploitFrameworkMetasploit bool                                                                           `json:"exploit_framework_metasploit,omitempty"`
	PluginType                 string                                                                         `json:"plugin_type,omitempty"`
	Vpr                        PluginsServicePluginsListResponseDataPluginDetailsItemAttributesItemVpr        `json:"vpr,omitempty"`
}

type PluginsServicePluginsListResponseDataPluginDetailsItem struct {
	Id         int                                                                    `json:"id,omitempty"`
	Name       string                                                                 `json:"name,omitempty"`
	Attributes []PluginsServicePluginsListResponseDataPluginDetailsItemAttributesItem `json:"attributes,omitempty"`
}

type PluginsServicePluginsListResponseData struct {
	PluginDetails []PluginsServicePluginsListResponseDataPluginDetailsItem `json:"plugin_details,omitempty"`
}

type PluginsServicePluginsListResponseParams struct {
	Page        int    `json:"page,omitempty"`
	Size        int    `json:"size,omitempty"`
	LastUpdated string `json:"last_updated,omitempty"`
}

type PluginsServicePluginsListResponse struct {
	Data       PluginsServicePluginsListResponseData   `json:"data,omitempty"`
	Size       int                                     `json:"size,omitempty"`
	Params     PluginsServicePluginsListResponseParams `json:"params,omitempty"`
	TotalCount int                                     `json:"total_count,omitempty"`
}

type PluginsServicePluginsDetailsResponseAttributesItem struct {
	AttributeName  string `json:"attribute_name,omitempty"`
	AttributeValue string `json:"attribute_value,omitempty"`
}

type PluginsServicePluginsDetailsResponse struct {
	Attributes []PluginsServicePluginsDetailsResponseAttributesItem `json:"attributes,omitempty"`
	FamilyName string                                               `json:"family_name,omitempty"`
	Name       string                                               `json:"name,omitempty"`
	Id         int                                                  `json:"id,omitempty"`
}

type PluginsServicePluginsFamiliesListResponseFamiliesItem struct {
	Count int    `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Id    int    `json:"id,omitempty"`
}

type PluginsServicePluginsFamiliesListResponse struct {
	Families []PluginsServicePluginsFamiliesListResponseFamiliesItem `json:"families,omitempty"`
}

type PluginsServicePluginsFamilyDetailsIdResponsePluginsItem struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PluginsServicePluginsFamilyDetailsIdResponse struct {
	Plugins []PluginsServicePluginsFamilyDetailsIdResponsePluginsItem `json:"plugins,omitempty"`
	Name    string                                                    `json:"name,omitempty"`
	Id      int                                                       `json:"id,omitempty"`
}

type PluginsServicePluginsFamilyDetailsNameRequest struct {
	Name string `json:"name,omitempty"`
}

type PluginsServicePluginsFamilyDetailsNameResponsePluginsItem struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PluginsServicePluginsFamilyDetailsNameResponse struct {
	Plugins []PluginsServicePluginsFamilyDetailsNameResponsePluginsItem `json:"plugins,omitempty"`
	Name    string                                                      `json:"name,omitempty"`
	Id      int                                                         `json:"id,omitempty"`
}

// PluginsList - List plugins
func (s *PluginsService) PluginsList(ctx context.Context) (*PluginsServicePluginsListResponse, error) {
	resp, err := s.client.get(ctx, "/plugins/plugin")
	if err != nil {
		return nil, err
	}
	var result PluginsServicePluginsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PluginsDetails - Get plugin details
func (s *PluginsService) PluginsDetails(ctx context.Context, id string) (*PluginsServicePluginsDetailsResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/plugins/plugin/%s", id))
	if err != nil {
		return nil, err
	}
	var result PluginsServicePluginsDetailsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PluginsFamiliesList - List plugin families
func (s *PluginsService) PluginsFamiliesList(ctx context.Context) (*PluginsServicePluginsFamiliesListResponse, error) {
	resp, err := s.client.get(ctx, "/plugins/families")
	if err != nil {
		return nil, err
	}
	var result PluginsServicePluginsFamiliesListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PluginsFamilyDetailsId - List plugins in family (ID)
func (s *PluginsService) PluginsFamilyDetailsId(ctx context.Context, id string) (*PluginsServicePluginsFamilyDetailsIdResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/plugins/families/%s", id))
	if err != nil {
		return nil, err
	}
	var result PluginsServicePluginsFamilyDetailsIdResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PluginsFamilyDetailsName - List plugins in family (name)
func (s *PluginsService) PluginsFamilyDetailsName(ctx context.Context, req *PluginsServicePluginsFamilyDetailsNameRequest) (*PluginsServicePluginsFamilyDetailsNameResponse, error) {
	resp, err := s.client.post(ctx, "/plugins/families/_byName", req)
	if err != nil {
		return nil, err
	}
	var result PluginsServicePluginsFamilyDetailsNameResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
