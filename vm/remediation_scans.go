package vm

import (
	"context"
	"encoding/json"
)

// RemediationScansService handles communication with the Remediation Scans related endpoints of the VM API.
type RemediationScansService struct {
	client *Client
}

type RemediationScansServiceRemediationCreateRequestSettingsAclsItem struct {
	Permissions int    `json:"permissions,omitempty"`
	Owner       int    `json:"owner,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Name        string `json:"name,omitempty"`
	Id          int    `json:"id,omitempty"`
	TypeField   string `json:"type,omitempty"`
}

type RemediationScansServiceRemediationCreateRequestSettings struct {
	Name              string                                                            `json:"name,omitempty"`
	Description       string                                                            `json:"description,omitempty"`
	ScannerId         string                                                            `json:"scanner_id,omitempty"`
	TargetNetworkUuid string                                                            `json:"target_network_uuid,omitempty"`
	ScanTimeWindow    int                                                               `json:"scan_time_window,omitempty"`
	TextTargets       string                                                            `json:"text_targets,omitempty"`
	TargetGroups      []int                                                             `json:"target_groups,omitempty"`
	FileTargets       string                                                            `json:"file_targets,omitempty"`
	TagTargets        []string                                                          `json:"tag_targets,omitempty"`
	AgentGroupId      []string                                                          `json:"agent_group_id,omitempty"`
	Emails            string                                                            `json:"emails,omitempty"`
	Acls              []RemediationScansServiceRemediationCreateRequestSettingsAclsItem `json:"acls,omitempty"`
}

type RemediationScansServiceRemediationCreateRequestCredentialsAddHostWindowsItem struct {
	Domain     string `json:"domain,omitempty"`
	Username   string `json:"username,omitempty"`
	AuthMethod string `json:"auth_method,omitempty"`
	Password   string `json:"password,omitempty"`
}

type RemediationScansServiceRemediationCreateRequestCredentialsAddHost struct {
	Windows []RemediationScansServiceRemediationCreateRequestCredentialsAddHostWindowsItem `json:"Windows,omitempty"`
}

type RemediationScansServiceRemediationCreateRequestCredentialsAdd struct {
	Host RemediationScansServiceRemediationCreateRequestCredentialsAddHost `json:"Host,omitempty"`
}

type RemediationScansServiceRemediationCreateRequestCredentials struct {
	Add RemediationScansServiceRemediationCreateRequestCredentialsAdd `json:"add,omitempty"`
}

type RemediationScansServiceRemediationCreateRequest struct {
	Uuid           string                                                     `json:"uuid,omitempty"`
	Settings       RemediationScansServiceRemediationCreateRequestSettings    `json:"settings,omitempty"`
	Credentials    RemediationScansServiceRemediationCreateRequestCredentials `json:"credentials,omitempty"`
	EnabledPlugins []int                                                      `json:"enabled_plugins,omitempty"`
}

type RemediationScansServiceRemediationCreateResponseTriggersItemOptions struct {
	PeriodicHourlyInterval int    `json:"periodic_hourly_interval,omitempty"`
	Filename               string `json:"filename,omitempty"`
}

type RemediationScansServiceRemediationCreateResponseTriggersItem struct {
	TypeField string                                                              `json:"type,omitempty"`
	Options   RemediationScansServiceRemediationCreateResponseTriggersItemOptions `json:"options,omitempty"`
}

type RemediationScansServiceRemediationCreateResponseNotificationFiltersItem struct {
	Value   string `json:"value,omitempty"`
	Quality string `json:"quality,omitempty"`
	Filter  string `json:"filter,omitempty"`
}

type RemediationScansServiceRemediationCreateResponse struct {
	TagType                string                                                                    `json:"tag_type,omitempty"`
	ContainerId            string                                                                    `json:"container_id,omitempty"`
	OwnerUuid              string                                                                    `json:"owner_uuid,omitempty"`
	Uuid                   string                                                                    `json:"uuid,omitempty"`
	Name                   string                                                                    `json:"name,omitempty"`
	Description            string                                                                    `json:"description,omitempty"`
	PolicyId               int                                                                       `json:"policy_id,omitempty"`
	ScannerUuid            string                                                                    `json:"scanner_uuid,omitempty"`
	TargetNetworkUuid      string                                                                    `json:"target_network_uuid,omitempty"`
	Emails                 string                                                                    `json:"emails,omitempty"`
	Sms                    string                                                                    `json:"sms,omitempty"`
	Enabled                bool                                                                      `json:"enabled,omitempty"`
	DashboardFile          string                                                                    `json:"dashboard_file,omitempty"`
	Remediation            int                                                                       `json:"remediation,omitempty"`
	IncludeAggregate       bool                                                                      `json:"include_aggregate,omitempty"`
	ScanTimeWindow         string                                                                    `json:"scan_time_window,omitempty"`
	CustomTargets          string                                                                    `json:"custom_targets,omitempty"`
	Triggers               []RemediationScansServiceRemediationCreateResponseTriggersItem            `json:"triggers,omitempty"`
	ReportingMode          string                                                                    `json:"reporting_mode,omitempty"`
	IntervalType           string                                                                    `json:"interval_type,omitempty"`
	IntervalValue          int                                                                       `json:"interval_value,omitempty"`
	BaselineNextScan       string                                                                    `json:"baseline_next_scan,omitempty"`
	AgentScanLaunchType    string                                                                    `json:"agent_scan_launch_type,omitempty"`
	Starttime              string                                                                    `json:"starttime,omitempty"`
	Rrules                 string                                                                    `json:"rrules,omitempty"`
	Timezone               string                                                                    `json:"timezone,omitempty"`
	NotificationFilterType string                                                                    `json:"notification_filter_type,omitempty"`
	NotificationFilters    []RemediationScansServiceRemediationCreateResponseNotificationFiltersItem `json:"notification_filters,omitempty"`
	TagTargets             []string                                                                  `json:"tag_targets,omitempty"`
	Shared                 bool                                                                      `json:"shared,omitempty"`
	UserPermissions        int                                                                       `json:"user_permissions,omitempty"`
	DefaultPermissions     int                                                                       `json:"default_permissions,omitempty"`
	Owner                  string                                                                    `json:"owner,omitempty"`
	OwnerId                int                                                                       `json:"owner_id,omitempty"`
	LastModificationDate   int                                                                       `json:"last_modification_date,omitempty"`
	CreationDate           int                                                                       `json:"creation_date,omitempty"`
	TypeField              string                                                                    `json:"type,omitempty"`
	Id                     int                                                                       `json:"id,omitempty"`
}

type RemediationScansServiceRemediationListResponsePaginationSortItem struct {
	Name  string `json:"name,omitempty"`
	Order string `json:"order,omitempty"`
}

type RemediationScansServiceRemediationListResponsePagination struct {
	Total  int                                                                `json:"total,omitempty"`
	Limit  int                                                                `json:"limit,omitempty"`
	Offset int                                                                `json:"offset,omitempty"`
	Sort   []RemediationScansServiceRemediationListResponsePaginationSortItem `json:"sort,omitempty"`
}

type RemediationScansServiceRemediationListResponseScansItem struct {
	Control              bool   `json:"control,omitempty"`
	CreationDate         int    `json:"creation_date,omitempty"`
	Enabled              bool   `json:"enabled,omitempty"`
	Id                   int    `json:"id,omitempty"`
	LastModificationDate int    `json:"last_modification_date,omitempty"`
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
	ScanCreationDate     int    `json:"scan_creation_date,omitempty"`
	Remediation          int    `json:"remediation,omitempty"`
	TotalTargets         int    `json:"total_targets,omitempty"`
}

type RemediationScansServiceRemediationListResponse struct {
	Pagination RemediationScansServiceRemediationListResponsePagination  `json:"pagination,omitempty"`
	Scans      []RemediationScansServiceRemediationListResponseScansItem `json:"scans,omitempty"`
}

// RemediationCreate - Create remediation scan
func (s *RemediationScansService) RemediationCreate(ctx context.Context, req *RemediationScansServiceRemediationCreateRequest) (*RemediationScansServiceRemediationCreateResponse, error) {
	resp, err := s.client.post(ctx, "/scans/remediation", req)
	if err != nil {
		return nil, err
	}
	var result RemediationScansServiceRemediationCreateResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RemediationList - List remediation scans
func (s *RemediationScansService) RemediationList(ctx context.Context) (*RemediationScansServiceRemediationListResponse, error) {
	resp, err := s.client.get(ctx, "/scans/remediation")
	if err != nil {
		return nil, err
	}
	var result RemediationScansServiceRemediationListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
