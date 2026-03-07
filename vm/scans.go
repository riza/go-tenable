package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// ScansService handles communication with the Scans related endpoints of the VM API.
type ScansService struct {
	client *Client
}

type ScansServiceCreateRequestSettingsTriggersItemOptions struct {
	PeriodicHourlyInterval int    `json:"periodic_hourly_interval,omitempty"`
	Filename               string `json:"filename,omitempty"`
}

type ScansServiceCreateRequestSettingsTriggersItem struct {
	TypeField string                                               `json:"type,omitempty"`
	Options   ScansServiceCreateRequestSettingsTriggersItemOptions `json:"options,omitempty"`
}

type ScansServiceCreateRequestSettingsFiltersItem struct {
	Filter  string      `json:"filter,omitempty"`
	Quality string      `json:"quality,omitempty"`
	Value   interface{} `json:"value,omitempty"`
}

type ScansServiceCreateRequestSettingsAclsItem struct {
	Permissions int    `json:"permissions,omitempty"`
	Owner       int    `json:"owner,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Name        string `json:"name,omitempty"`
	Id          int    `json:"id,omitempty"`
	TypeField   string `json:"type,omitempty"`
}

type ScansServiceCreateRequestSettings struct {
	Name                           string                                          `json:"name,omitempty"`
	Description                    string                                          `json:"description,omitempty"`
	PolicyId                       int                                             `json:"policy_id,omitempty"`
	FolderId                       int                                             `json:"folder_id,omitempty"`
	OwnerId                        int                                             `json:"owner_id,omitempty"`
	ScannerId                      string                                          `json:"scanner_id,omitempty"`
	TargetNetworkUuid              string                                          `json:"target_network_uuid,omitempty"`
	Enabled                        bool                                            `json:"enabled,omitempty"`
	Launch                         string                                          `json:"launch,omitempty"`
	ScanTimeWindow                 int                                             `json:"scan_time_window,omitempty"`
	Starttime                      string                                          `json:"starttime,omitempty"`
	Rrules                         string                                          `json:"rrules,omitempty"`
	Timezone                       string                                          `json:"timezone,omitempty"`
	TextTargets                    string                                          `json:"text_targets,omitempty"`
	TargetGroups                   []int                                           `json:"target_groups,omitempty"`
	FileTargets                    string                                          `json:"file_targets,omitempty"`
	TagTargets                     []string                                        `json:"tag_targets,omitempty"`
	TagTargetsIpSelection          int                                             `json:"tag_targets_ip_selection,omitempty"`
	UseTagRulesAsTargets           int                                             `json:"use_tag_rules_as_targets,omitempty"`
	HostTagging                    string                                          `json:"host_tagging,omitempty"`
	AgentGroupId                   []string                                        `json:"agent_group_id,omitempty"`
	AgentScanLaunchType            string                                          `json:"agent_scan_launch_type,omitempty"`
	Triggers                       []ScansServiceCreateRequestSettingsTriggersItem `json:"triggers,omitempty"`
	RefreshReportingType           string                                          `json:"refresh_reporting_type,omitempty"`
	RefreshReportingFrequencyScans int                                             `json:"refresh_reporting_frequency_scans,omitempty"`
	RefreshReportingFrequencyDays  int                                             `json:"refresh_reporting_frequency_days,omitempty"`
	DisableRefreshReporting        string                                          `json:"disable_refresh_reporting,omitempty"`
	Emails                         string                                          `json:"emails,omitempty"`
	Sms                            string                                          `json:"sms,omitempty"`
	FilterType                     string                                          `json:"filter_type,omitempty"`
	Filters                        []ScansServiceCreateRequestSettingsFiltersItem  `json:"filters,omitempty"`
	Acls                           []ScansServiceCreateRequestSettingsAclsItem     `json:"acls,omitempty"`
	ReportSupersededPatches        string                                          `json:"report_superseded_patches,omitempty"`
	ThoroughTests                  string                                          `json:"thorough_tests,omitempty"`
}

type ScansServiceCreateRequestCredentialsAddHostWindowsItem struct {
	Domain     string `json:"domain,omitempty"`
	Username   string `json:"username,omitempty"`
	AuthMethod string `json:"auth_method,omitempty"`
	Password   string `json:"password,omitempty"`
}

type ScansServiceCreateRequestCredentialsAddHost struct {
	Windows []ScansServiceCreateRequestCredentialsAddHostWindowsItem `json:"Windows,omitempty"`
}

type ScansServiceCreateRequestCredentialsAdd struct {
	Host ScansServiceCreateRequestCredentialsAddHost `json:"Host,omitempty"`
}

type ScansServiceCreateRequestCredentials struct {
	Add ScansServiceCreateRequestCredentialsAdd `json:"add,omitempty"`
}

type ScansServiceCreateRequestPluginsWebServersIndividual struct {
	Num11213 string `json:"11213,omitempty"`
	Num18261 string `json:"18261,omitempty"`
}

type ScansServiceCreateRequestPluginsWebServers struct {
	Individual ScansServiceCreateRequestPluginsWebServersIndividual `json:"individual,omitempty"`
}

type ScansServiceCreateRequestPlugins struct {
	WebServers ScansServiceCreateRequestPluginsWebServers `json:"Web Servers,omitempty"`
}

type ScansServiceCreateRequest struct {
	Uuid        string                               `json:"uuid,omitempty"`
	Settings    ScansServiceCreateRequestSettings    `json:"settings,omitempty"`
	Credentials ScansServiceCreateRequestCredentials `json:"credentials,omitempty"`
	Plugins     ScansServiceCreateRequestPlugins     `json:"plugins,omitempty"`
}

type ScansServiceCreateResponseTriggersItemOptions struct {
	PeriodicHourlyInterval int    `json:"periodic_hourly_interval,omitempty"`
	Filename               string `json:"filename,omitempty"`
}

type ScansServiceCreateResponseTriggersItem struct {
	TypeField string                                        `json:"type,omitempty"`
	Options   ScansServiceCreateResponseTriggersItemOptions `json:"options,omitempty"`
}

type ScansServiceCreateResponseNotificationFiltersItem struct {
	Value   string `json:"value,omitempty"`
	Quality string `json:"quality,omitempty"`
	Filter  string `json:"filter,omitempty"`
}

type ScansServiceCreateResponse struct {
	TagType                string                                              `json:"tag_type,omitempty"`
	ContainerId            string                                              `json:"container_id,omitempty"`
	OwnerUuid              string                                              `json:"owner_uuid,omitempty"`
	Uuid                   string                                              `json:"uuid,omitempty"`
	Name                   string                                              `json:"name,omitempty"`
	Description            string                                              `json:"description,omitempty"`
	PolicyId               int                                                 `json:"policy_id,omitempty"`
	ScannerUuid            string                                              `json:"scanner_uuid,omitempty"`
	TargetNetworkUuid      string                                              `json:"target_network_uuid,omitempty"`
	Emails                 string                                              `json:"emails,omitempty"`
	Sms                    string                                              `json:"sms,omitempty"`
	Enabled                bool                                                `json:"enabled,omitempty"`
	DashboardFile          string                                              `json:"dashboard_file,omitempty"`
	Remediation            int                                                 `json:"remediation,omitempty"`
	IncludeAggregate       bool                                                `json:"include_aggregate,omitempty"`
	ScanTimeWindow         string                                              `json:"scan_time_window,omitempty"`
	CustomTargets          string                                              `json:"custom_targets,omitempty"`
	Triggers               []ScansServiceCreateResponseTriggersItem            `json:"triggers,omitempty"`
	ReportingMode          string                                              `json:"reporting_mode,omitempty"`
	IntervalType           string                                              `json:"interval_type,omitempty"`
	IntervalValue          int                                                 `json:"interval_value,omitempty"`
	BaselineNextScan       string                                              `json:"baseline_next_scan,omitempty"`
	AgentScanLaunchType    string                                              `json:"agent_scan_launch_type,omitempty"`
	Starttime              string                                              `json:"starttime,omitempty"`
	Rrules                 string                                              `json:"rrules,omitempty"`
	Timezone               string                                              `json:"timezone,omitempty"`
	NotificationFilterType string                                              `json:"notification_filter_type,omitempty"`
	NotificationFilters    []ScansServiceCreateResponseNotificationFiltersItem `json:"notification_filters,omitempty"`
	TagTargets             []string                                            `json:"tag_targets,omitempty"`
	Shared                 bool                                                `json:"shared,omitempty"`
	UserPermissions        int                                                 `json:"user_permissions,omitempty"`
	DefaultPermissions     int                                                 `json:"default_permissions,omitempty"`
	Owner                  string                                              `json:"owner,omitempty"`
	OwnerId                int                                                 `json:"owner_id,omitempty"`
	LastModificationDate   int                                                 `json:"last_modification_date,omitempty"`
	CreationDate           int                                                 `json:"creation_date,omitempty"`
	TypeField              string                                              `json:"type,omitempty"`
	Id                     int                                                 `json:"id,omitempty"`
}

type ScansServiceListResponseScansItem struct {
	Control              bool   `json:"control,omitempty"`
	CreationDate         int    `json:"creation_date,omitempty"`
	Enabled              bool   `json:"enabled,omitempty"`
	Id                   int    `json:"id,omitempty"`
	LastModificationDate int    `json:"last_modification_date,omitempty"`
	Legacy               bool   `json:"legacy,omitempty"`
	Name                 string `json:"name,omitempty"`
	Owner                string `json:"owner,omitempty"`
	PolicyId             int    `json:"policy_id,omitempty"`
	Read                 bool   `json:"read,omitempty"`
	ScheduleUuid         string `json:"schedule_uuid,omitempty"`
	Shared               bool   `json:"shared,omitempty"`
	Status               string `json:"status,omitempty"`
	TemplateUuid         string `json:"template_uuid,omitempty"`
	TypeField            string `json:"type,omitempty"`
	Permissions          int    `json:"permissions,omitempty"`
	UserPermissions      int    `json:"user_permissions,omitempty"`
	Uuid                 string `json:"uuid,omitempty"`
	WizardUuid           string `json:"wizard_uuid,omitempty"`
	Progress             int    `json:"progress,omitempty"`
	Timezone             string `json:"timezone,omitempty"`
	Rrules               string `json:"rrules,omitempty"`
	Starttime            string `json:"starttime,omitempty"`
	TotalTargets         int    `json:"total_targets,omitempty"`
}

type ScansServiceListResponse struct {
	Scans     []ScansServiceListResponseScansItem `json:"scans,omitempty"`
	Folders   []interface{}                       `json:"folders,omitempty"`
	Timestamp int                                 `json:"timestamp,omitempty"`
}

type ScansServiceDetailsResponseInfoAclsItem struct {
	Permissions int    `json:"permissions,omitempty"`
	Owner       int    `json:"owner,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Name        string `json:"name,omitempty"`
	Id          int    `json:"id,omitempty"`
	TypeField   string `json:"type,omitempty"`
}

type ScansServiceDetailsResponseInfo struct {
	Owner           string                                    `json:"owner,omitempty"`
	Name            string                                    `json:"name,omitempty"`
	NoTarget        bool                                      `json:"no_target,omitempty"`
	FolderId        int                                       `json:"folder_id,omitempty"`
	Control         bool                                      `json:"control,omitempty"`
	UserPermissions int                                       `json:"user_permissions,omitempty"`
	ScheduleUuid    string                                    `json:"schedule_uuid,omitempty"`
	EditAllowed     bool                                      `json:"edit_allowed,omitempty"`
	ScannerName     string                                    `json:"scanner_name,omitempty"`
	Policy          string                                    `json:"policy,omitempty"`
	Shared          bool                                      `json:"shared,omitempty"`
	ObjectId        int                                       `json:"object_id,omitempty"`
	TagTargets      []string                                  `json:"tag_targets,omitempty"`
	Acls            []ScansServiceDetailsResponseInfoAclsItem `json:"acls,omitempty"`
	Hostcount       int                                       `json:"hostcount,omitempty"`
	Uuid            string                                    `json:"uuid,omitempty"`
	Status          string                                    `json:"status,omitempty"`
	ScanType        string                                    `json:"scan_type,omitempty"`
	Targets         string                                    `json:"targets,omitempty"`
	AltTargetsUsed  bool                                      `json:"alt_targets_used,omitempty"`
	PciCanUpload    bool                                      `json:"pci-can-upload,omitempty"`
	ScanStart       int                                       `json:"scan_start,omitempty"`
	Timestamp       int                                       `json:"timestamp,omitempty"`
	IsArchived      bool                                      `json:"is_archived,omitempty"`
	ScanEnd         int                                       `json:"scan_end,omitempty"`
	Haskb           bool                                      `json:"haskb,omitempty"`
	Hasaudittrail   bool                                      `json:"hasaudittrail,omitempty"`
	ScannerStart    string                                    `json:"scanner_start,omitempty"`
	ScannerEnd      string                                    `json:"scanner_end,omitempty"`
}

type ScansServiceDetailsResponseComphostsItemSeveritycount struct {
}

type ScansServiceDetailsResponseComphostsItem struct {
	Totalchecksconsidered int                                                   `json:"totalchecksconsidered,omitempty"`
	Numchecksconsidered   int                                                   `json:"numchecksconsidered,omitempty"`
	Scanprogresstotal     int                                                   `json:"scanprogresstotal,omitempty"`
	Scanprogresscurrent   int                                                   `json:"scanprogresscurrent,omitempty"`
	HostIndex             string                                                `json:"host_index,omitempty"`
	Score                 int                                                   `json:"score,omitempty"`
	Severitycount         ScansServiceDetailsResponseComphostsItemSeveritycount `json:"severitycount,omitempty"`
	Progress              string                                                `json:"progress,omitempty"`
	Critical              int                                                   `json:"critical,omitempty"`
	High                  int                                                   `json:"high,omitempty"`
	Medium                int                                                   `json:"medium,omitempty"`
	Low                   int                                                   `json:"low,omitempty"`
	Info                  int                                                   `json:"info,omitempty"`
	HostId                int                                                   `json:"host_id,omitempty"`
	Hostname              string                                                `json:"hostname,omitempty"`
}

type ScansServiceDetailsResponseHostsItemSeveritycount struct {
}

type ScansServiceDetailsResponseHostsItem struct {
	Totalchecksconsidered int                                               `json:"totalchecksconsidered,omitempty"`
	Numchecksconsidered   int                                               `json:"numchecksconsidered,omitempty"`
	Scanprogresstotal     int                                               `json:"scanprogresstotal,omitempty"`
	Scanprogresscurrent   int                                               `json:"scanprogresscurrent,omitempty"`
	HostIndex             string                                            `json:"host_index,omitempty"`
	Score                 int                                               `json:"score,omitempty"`
	Severitycount         ScansServiceDetailsResponseHostsItemSeveritycount `json:"severitycount,omitempty"`
	Progress              string                                            `json:"progress,omitempty"`
	Critical              int                                               `json:"critical,omitempty"`
	High                  int                                               `json:"high,omitempty"`
	Medium                int                                               `json:"medium,omitempty"`
	Low                   int                                               `json:"low,omitempty"`
	Info                  int                                               `json:"info,omitempty"`
	HostId                int                                               `json:"host_id,omitempty"`
	Hostname              string                                            `json:"hostname,omitempty"`
}

type ScansServiceDetailsResponseNotesItem struct {
	Title    string `json:"title,omitempty"`
	Message  string `json:"message,omitempty"`
	Severity int    `json:"severity,omitempty"`
}

type ScansServiceDetailsResponseVulnerabilitiesItem struct {
	Count         int    `json:"count,omitempty"`
	PluginName    string `json:"plugin_name,omitempty"`
	VulnIndex     int    `json:"vuln_index,omitempty"`
	Severity      int    `json:"severity,omitempty"`
	PluginId      int    `json:"plugin_id,omitempty"`
	SeverityIndex int    `json:"severity_index,omitempty"`
	PluginFamily  string `json:"plugin_family,omitempty"`
}

type ScansServiceDetailsResponseFiltersItemControl struct {
	TypeField      string   `json:"type,omitempty"`
	ReadableRegest string   `json:"readable_regest,omitempty"`
	Regex          string   `json:"regex,omitempty"`
	Options        []string `json:"options,omitempty"`
}

type ScansServiceDetailsResponseFiltersItem struct {
	Name         string                                        `json:"name,omitempty"`
	ReadableName string                                        `json:"readable_name,omitempty"`
	Operators    []string                                      `json:"operators,omitempty"`
	Control      ScansServiceDetailsResponseFiltersItemControl `json:"control,omitempty"`
}

type ScansServiceDetailsResponseHistoryItem struct {
	HistoryId            int    `json:"history_id,omitempty"`
	OwnerId              int    `json:"owner_id,omitempty"`
	CreationDate         int    `json:"creation_date,omitempty"`
	LastModificationDate int    `json:"last_modification_date,omitempty"`
	Uuid                 string `json:"uuid,omitempty"`
	TypeField            string `json:"type,omitempty"`
	Status               string `json:"status,omitempty"`
	Scheduler            int    `json:"scheduler,omitempty"`
	AltTargetsUsed       bool   `json:"alt_targets_used,omitempty"`
	IsArchived           bool   `json:"is_archived,omitempty"`
	ReportingMode        string `json:"reporting_mode,omitempty"`
}

type ScansServiceDetailsResponseComplianceItem struct {
	Count         int    `json:"count,omitempty"`
	PluginName    string `json:"plugin_name,omitempty"`
	VulnIndex     int    `json:"vuln_index,omitempty"`
	Severity      int    `json:"severity,omitempty"`
	PluginId      int    `json:"plugin_id,omitempty"`
	SeverityIndex int    `json:"severity_index,omitempty"`
	PluginFamily  string `json:"plugin_family,omitempty"`
}

type ScansServiceDetailsResponse struct {
	Info            ScansServiceDetailsResponseInfo                  `json:"info,omitempty"`
	Comphosts       []ScansServiceDetailsResponseComphostsItem       `json:"comphosts,omitempty"`
	Hosts           []ScansServiceDetailsResponseHostsItem           `json:"hosts,omitempty"`
	Notes           []ScansServiceDetailsResponseNotesItem           `json:"notes,omitempty"`
	Remediations    map[string]interface{}                           `json:"remediations,omitempty"`
	Vulnerabilities []ScansServiceDetailsResponseVulnerabilitiesItem `json:"vulnerabilities,omitempty"`
	Filters         []ScansServiceDetailsResponseFiltersItem         `json:"filters,omitempty"`
	History         []ScansServiceDetailsResponseHistoryItem         `json:"history,omitempty"`
	Compliance      []ScansServiceDetailsResponseComplianceItem      `json:"compliance,omitempty"`
	Progress        int                                              `json:"progress,omitempty"`
}

type ScansServiceConfigureRequestSettingsTriggersItemOptions struct {
	PeriodicHourlyInterval int    `json:"periodic_hourly_interval,omitempty"`
	Filename               string `json:"filename,omitempty"`
}

type ScansServiceConfigureRequestSettingsTriggersItem struct {
	TypeField string                                                  `json:"type,omitempty"`
	Options   ScansServiceConfigureRequestSettingsTriggersItemOptions `json:"options,omitempty"`
}

type ScansServiceConfigureRequestSettingsFiltersItem struct {
	Filter  string      `json:"filter,omitempty"`
	Quality string      `json:"quality,omitempty"`
	Value   interface{} `json:"value,omitempty"`
}

type ScansServiceConfigureRequestSettingsAclsItem struct {
	Permissions int    `json:"permissions,omitempty"`
	Owner       int    `json:"owner,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Name        string `json:"name,omitempty"`
	Id          int    `json:"id,omitempty"`
	TypeField   string `json:"type,omitempty"`
}

type ScansServiceConfigureRequestSettings struct {
	Name                           string                                             `json:"name,omitempty"`
	Description                    string                                             `json:"description,omitempty"`
	PolicyId                       int                                                `json:"policy_id,omitempty"`
	FolderId                       int                                                `json:"folder_id,omitempty"`
	OwnerId                        int                                                `json:"owner_id,omitempty"`
	ScannerId                      string                                             `json:"scanner_id,omitempty"`
	TargetNetworkUuid              string                                             `json:"target_network_uuid,omitempty"`
	Enabled                        bool                                               `json:"enabled,omitempty"`
	Launch                         string                                             `json:"launch,omitempty"`
	ScanTimeWindow                 int                                                `json:"scan_time_window,omitempty"`
	Starttime                      string                                             `json:"starttime,omitempty"`
	Rrules                         string                                             `json:"rrules,omitempty"`
	Timezone                       string                                             `json:"timezone,omitempty"`
	TextTargets                    string                                             `json:"text_targets,omitempty"`
	TargetGroups                   []int                                              `json:"target_groups,omitempty"`
	FileTargets                    string                                             `json:"file_targets,omitempty"`
	TagTargets                     []string                                           `json:"tag_targets,omitempty"`
	TagTargetsIpSelection          int                                                `json:"tag_targets_ip_selection,omitempty"`
	UseTagRulesAsTargets           int                                                `json:"use_tag_rules_as_targets,omitempty"`
	HostTagging                    string                                             `json:"host_tagging,omitempty"`
	AgentGroupId                   []string                                           `json:"agent_group_id,omitempty"`
	AgentScanLaunchType            string                                             `json:"agent_scan_launch_type,omitempty"`
	Triggers                       []ScansServiceConfigureRequestSettingsTriggersItem `json:"triggers,omitempty"`
	RefreshReportingType           string                                             `json:"refresh_reporting_type,omitempty"`
	RefreshReportingFrequencyScans int                                                `json:"refresh_reporting_frequency_scans,omitempty"`
	RefreshReportingFrequencyDays  int                                                `json:"refresh_reporting_frequency_days,omitempty"`
	DisableRefreshReporting        string                                             `json:"disable_refresh_reporting,omitempty"`
	Emails                         string                                             `json:"emails,omitempty"`
	Sms                            string                                             `json:"sms,omitempty"`
	FilterType                     string                                             `json:"filter_type,omitempty"`
	Filters                        []ScansServiceConfigureRequestSettingsFiltersItem  `json:"filters,omitempty"`
	Acls                           []ScansServiceConfigureRequestSettingsAclsItem     `json:"acls,omitempty"`
	ReportSupersededPatches        string                                             `json:"report_superseded_patches,omitempty"`
	ThoroughTests                  string                                             `json:"thorough_tests,omitempty"`
}

type ScansServiceConfigureRequestCredentialsAddHostWindowsItem struct {
	Domain     string `json:"domain,omitempty"`
	Username   string `json:"username,omitempty"`
	AuthMethod string `json:"auth_method,omitempty"`
	Password   string `json:"password,omitempty"`
}

type ScansServiceConfigureRequestCredentialsAddHost struct {
	Windows []ScansServiceConfigureRequestCredentialsAddHostWindowsItem `json:"Windows,omitempty"`
}

type ScansServiceConfigureRequestCredentialsAdd struct {
	Host ScansServiceConfigureRequestCredentialsAddHost `json:"Host,omitempty"`
}

type ScansServiceConfigureRequestCredentialsEdit struct {
}

type ScansServiceConfigureRequestCredentials struct {
	Add    ScansServiceConfigureRequestCredentialsAdd  `json:"add,omitempty"`
	Edit   ScansServiceConfigureRequestCredentialsEdit `json:"edit,omitempty"`
	Delete []string                                    `json:"delete,omitempty"`
}

type ScansServiceConfigureRequestPluginsWebServersIndividual struct {
	Num11213 string `json:"11213,omitempty"`
	Num18261 string `json:"18261,omitempty"`
}

type ScansServiceConfigureRequestPluginsWebServers struct {
	Individual ScansServiceConfigureRequestPluginsWebServersIndividual `json:"individual,omitempty"`
}

type ScansServiceConfigureRequestPlugins struct {
	WebServers ScansServiceConfigureRequestPluginsWebServers `json:"Web Servers,omitempty"`
}

type ScansServiceConfigureRequest struct {
	Uuid        string                                  `json:"uuid,omitempty"`
	Settings    ScansServiceConfigureRequestSettings    `json:"settings,omitempty"`
	Credentials ScansServiceConfigureRequestCredentials `json:"credentials,omitempty"`
	Plugins     ScansServiceConfigureRequestPlugins     `json:"plugins,omitempty"`
}

type ScansServiceConfigureResponseTriggersItemOptions struct {
	PeriodicHourlyInterval int    `json:"periodic_hourly_interval,omitempty"`
	Filename               string `json:"filename,omitempty"`
}

type ScansServiceConfigureResponseTriggersItem struct {
	TypeField string                                           `json:"type,omitempty"`
	Options   ScansServiceConfigureResponseTriggersItemOptions `json:"options,omitempty"`
}

type ScansServiceConfigureResponseNotificationFiltersItem struct {
	Value   string `json:"value,omitempty"`
	Quality string `json:"quality,omitempty"`
	Filter  string `json:"filter,omitempty"`
}

type ScansServiceConfigureResponse struct {
	TagType                string                                                 `json:"tag_type,omitempty"`
	ContainerId            string                                                 `json:"container_id,omitempty"`
	OwnerUuid              string                                                 `json:"owner_uuid,omitempty"`
	Uuid                   string                                                 `json:"uuid,omitempty"`
	Name                   string                                                 `json:"name,omitempty"`
	Description            string                                                 `json:"description,omitempty"`
	PolicyId               int                                                    `json:"policy_id,omitempty"`
	ScannerUuid            string                                                 `json:"scanner_uuid,omitempty"`
	TargetNetworkUuid      string                                                 `json:"target_network_uuid,omitempty"`
	Emails                 string                                                 `json:"emails,omitempty"`
	Sms                    string                                                 `json:"sms,omitempty"`
	Enabled                bool                                                   `json:"enabled,omitempty"`
	DashboardFile          string                                                 `json:"dashboard_file,omitempty"`
	Remediation            int                                                    `json:"remediation,omitempty"`
	IncludeAggregate       bool                                                   `json:"include_aggregate,omitempty"`
	ScanTimeWindow         string                                                 `json:"scan_time_window,omitempty"`
	CustomTargets          string                                                 `json:"custom_targets,omitempty"`
	Triggers               []ScansServiceConfigureResponseTriggersItem            `json:"triggers,omitempty"`
	ReportingMode          string                                                 `json:"reporting_mode,omitempty"`
	IntervalType           string                                                 `json:"interval_type,omitempty"`
	IntervalValue          int                                                    `json:"interval_value,omitempty"`
	BaselineNextScan       string                                                 `json:"baseline_next_scan,omitempty"`
	AgentScanLaunchType    string                                                 `json:"agent_scan_launch_type,omitempty"`
	Starttime              string                                                 `json:"starttime,omitempty"`
	Rrules                 string                                                 `json:"rrules,omitempty"`
	Timezone               string                                                 `json:"timezone,omitempty"`
	NotificationFilterType string                                                 `json:"notification_filter_type,omitempty"`
	NotificationFilters    []ScansServiceConfigureResponseNotificationFiltersItem `json:"notification_filters,omitempty"`
	TagTargets             []string                                               `json:"tag_targets,omitempty"`
	Shared                 bool                                                   `json:"shared,omitempty"`
	UserPermissions        int                                                    `json:"user_permissions,omitempty"`
	DefaultPermissions     int                                                    `json:"default_permissions,omitempty"`
	Owner                  string                                                 `json:"owner,omitempty"`
	OwnerId                int                                                    `json:"owner_id,omitempty"`
	LastModificationDate   int                                                    `json:"last_modification_date,omitempty"`
	CreationDate           int                                                    `json:"creation_date,omitempty"`
	TypeField              string                                                 `json:"type,omitempty"`
	Id                     int                                                    `json:"id,omitempty"`
}

// Create - Create scan
func (s *ScansService) Create(ctx context.Context, req *ScansServiceCreateRequest) (*ScansServiceCreateResponse, error) {
	resp, err := s.client.post(ctx, "/scans", req)
	if err != nil {
		return nil, err
	}
	var result ScansServiceCreateResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List - List scans
func (s *ScansService) List(ctx context.Context) (*ScansServiceListResponse, error) {
	resp, err := s.client.get(ctx, "/scans")
	if err != nil {
		return nil, err
	}
	var result ScansServiceListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Details - Get scan details
func (s *ScansService) Details(ctx context.Context, scanId string) (*ScansServiceDetailsResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/scans/%s", scanId))
	if err != nil {
		return nil, err
	}
	var result ScansServiceDetailsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Configure - Update scan
func (s *ScansService) Configure(ctx context.Context, scanId string, req *ScansServiceConfigureRequest) (*ScansServiceConfigureResponse, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/scans/%s", scanId), req)
	if err != nil {
		return nil, err
	}
	var result ScansServiceConfigureResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete - Delete scan
func (s *ScansService) Delete(ctx context.Context, scanId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/scans/%s", scanId))
	if err != nil {
		return err
	}
	return nil
}
