package platform

import (
	"context"
	"encoding/json"
)

// ActivityLogService handles communication with the Activity Log related endpoints of the Platform API.
type ActivityLogService struct {
	client *Client
}

// ActivityLogEvent represents an activity log event.
type ActivityLogEvent struct {
	Id          string                 `json:"id,omitempty"`
	Action      string                 `json:"action,omitempty"`
	Actor       string                 `json:"actor,omitempty"`
	Target      string                 `json:"target,omitempty"`
	TargetType  string                 `json:"target_type,omitempty"`
	Description string                 `json:"description,omitempty"`
	CreatedAt   string                 `json:"created_time,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// ActivityLogListResponse represents the response from listing activity log events.
type ActivityLogListResponse struct {
	Events []ActivityLogEvent `json:"events,omitempty"`
	Total  int                `json:"total,omitempty"`
}

// ListActivityLog returns activity log events.
func (s *ActivityLogService) ListActivityLog(ctx context.Context) (*ActivityLogListResponse, error) {
	resp, err := s.client.get(ctx, "/audit-log/v1/events")
	if err != nil {
		return nil, err
	}

	var result ActivityLogListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CloudConnectorsService handles communication with the Cloud Connectors related endpoints of the Platform API.
type CloudConnectorsService struct {
	client *Client
}

// CloudConnector represents a cloud connector.
type CloudConnector struct {
	Id        string                 `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Type      string                 `json:"type,omitempty"`
	Status    string                 `json:"status,omitempty"`
	CreatedAt string                 `json:"created_time,omitempty"`
	UpdatedAt string                 `json:"updated_time,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// CloudConnectorsListResponse represents the response from listing cloud connectors.
type CloudConnectorsListResponse struct {
	Connectors []CloudConnector `json:"cloud_connectors,omitempty"`
	Total      int              `json:"total,omitempty"`
}

// ListCloudConnectors returns all cloud connectors.
func (s *CloudConnectorsService) ListCloudConnectors(ctx context.Context) (*CloudConnectorsListResponse, error) {
	resp, err := s.client.get(ctx, "/cloud-connectors")
	if err != nil {
		return nil, err
	}

	var result CloudConnectorsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateCloudConnector creates a new cloud connector.
func (s *CloudConnectorsService) CreateCloudConnector(ctx context.Context, req *CloudConnector) (*CloudConnector, error) {
	resp, err := s.client.post(ctx, "/cloud-connectors", req)
	if err != nil {
		return nil, err
	}

	var result CloudConnector
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetCloudConnector returns a specific cloud connector by ID.
func (s *CloudConnectorsService) GetCloudConnector(ctx context.Context, connectorId string) (*CloudConnector, error) {
	resp, err := s.client.get(ctx, "/cloud-connectors/"+connectorId)
	if err != nil {
		return nil, err
	}

	var result CloudConnector
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateCloudConnector updates a cloud connector.
func (s *CloudConnectorsService) UpdateCloudConnector(ctx context.Context, connectorId string, req *CloudConnector) (*CloudConnector, error) {
	resp, err := s.client.put(ctx, "/cloud-connectors/"+connectorId, req)
	if err != nil {
		return nil, err
	}

	var result CloudConnector
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteCloudConnector deletes a cloud connector.
func (s *CloudConnectorsService) DeleteCloudConnector(ctx context.Context, connectorId string) error {
	_, err := s.client.delete(ctx, "/cloud-connectors/"+connectorId)
	return err
}

// ExclusionsService handles communication with the Exclusions related endpoints of the Platform API.
type ExclusionsService struct {
	client *Client
}

// Exclusion represents an exclusion.
type Exclusion struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Value       string `json:"value,omitempty"`
	StartTime   int    `json:"start_time,omitempty"`
	EndTime     int    `json:"end_time,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
}

// ExclusionsListResponse represents the response from listing exclusions.
type ExclusionsListResponse struct {
	Exclusions []Exclusion `json:"exclusions,omitempty"`
	Total      int         `json:"total,omitempty"`
}

// ListExclusions returns all exclusions.
func (s *ExclusionsService) ListExclusions(ctx context.Context) (*ExclusionsListResponse, error) {
	resp, err := s.client.get(ctx, "/exclusions")
	if err != nil {
		return nil, err
	}

	var result ExclusionsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateExclusion creates a new exclusion.
func (s *ExclusionsService) CreateExclusion(ctx context.Context, req *Exclusion) (*Exclusion, error) {
	resp, err := s.client.post(ctx, "/exclusions", req)
	if err != nil {
		return nil, err
	}

	var result Exclusion
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetExclusion returns a specific exclusion by ID.
func (s *ExclusionsService) GetExclusion(ctx context.Context, exclusionId string) (*Exclusion, error) {
	resp, err := s.client.get(ctx, "/exclusions/"+exclusionId)
	if err != nil {
		return nil, err
	}

	var result Exclusion
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateExclusion updates an exclusion.
func (s *ExclusionsService) UpdateExclusion(ctx context.Context, exclusionId string, req *Exclusion) (*Exclusion, error) {
	resp, err := s.client.put(ctx, "/exclusions/"+exclusionId, req)
	if err != nil {
		return nil, err
	}

	var result Exclusion
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteExclusion deletes an exclusion.
func (s *ExclusionsService) DeleteExclusion(ctx context.Context, exclusionId string) error {
	_, err := s.client.delete(ctx, "/exclusions/"+exclusionId)
	return err
}

// ServerService handles communication with the Server related endpoints of the Platform API.
type ServerService struct {
	client *Client
}

// ServerInfo represents server information.
type ServerInfo struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Version  string `json:"version,omitempty"`
	Build    string `json:"build,omitempty"`
	Edition  string `json:"edition,omitempty"`
	License  string `json:"license,omitempty"`
	Platform string `json:"platform,omitempty"`
}

// GetServerInfo returns server information.
func (s *ServerService) GetServerInfo(ctx context.Context) (*ServerInfo, error) {
	resp, err := s.client.get(ctx, "/server")
	if err != nil {
		return nil, err
	}

	var result ServerInfo
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ProfilesService handles communication with the Profiles related endpoints of the Platform API.
type ProfilesService struct {
	client *Client
}

// Profile represents a scan profile.
type Profile struct {
	Id          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Settings    map[string]interface{} `json:"settings,omitempty"`
	CreatedAt   string                 `json:"created_time,omitempty"`
	UpdatedAt   string                 `json:"updated_time,omitempty"`
}

// ProfilesListResponse represents the response from listing profiles.
type ProfilesListResponse struct {
	Profiles []Profile `json:"profiles,omitempty"`
	Total    int       `json:"total,omitempty"`
}

// ListProfiles returns all profiles.
func (s *ProfilesService) ListProfiles(ctx context.Context) (*ProfilesListResponse, error) {
	resp, err := s.client.get(ctx, "/profiles")
	if err != nil {
		return nil, err
	}

	var result ProfilesListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// RecastRulesService handles communication with the Recast Rules related endpoints of the Platform API.
type RecastRulesService struct {
	client *Client
}

// RecastRule represents a recast rule.
type RecastRule struct {
	Id          string `json:"id,omitempty"`
	PluginId    int    `json:"plugin_id,omitempty"`
	Severity    string `json:"severity,omitempty"`
	RecastType  string `json:"recast_type,omitempty"`
	NewSeverity string `json:"new_severity,omitempty"`
	CreatedBy   string `json:"created_by,omitempty"`
	CreatedAt   string `json:"created_time,omitempty"`
}

// RecastRulesListResponse represents the response from listing recast rules.
type RecastRulesListResponse struct {
	Rules []RecastRule `json:"rules,omitempty"`
	Total int          `json:"total,omitempty"`
}

// ListRecastRules returns all recast rules.
func (s *RecastRulesService) ListRecastRules(ctx context.Context) (*RecastRulesListResponse, error) {
	resp, err := s.client.get(ctx, "/recast-rules")
	if err != nil {
		return nil, err
	}

	var result RecastRulesListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateRecastRule creates a new recast rule.
func (s *RecastRulesService) CreateRecastRule(ctx context.Context, req *RecastRule) (*RecastRule, error) {
	resp, err := s.client.post(ctx, "/recast-rules", req)
	if err != nil {
		return nil, err
	}

	var result RecastRule
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetRecastRule returns a specific recast rule by ID.
func (s *RecastRulesService) GetRecastRule(ctx context.Context, ruleId string) (*RecastRule, error) {
	resp, err := s.client.get(ctx, "/recast-rules/"+ruleId)
	if err != nil {
		return nil, err
	}

	var result RecastRule
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteRecastRule deletes a recast rule.
func (s *RecastRulesService) DeleteRecastRule(ctx context.Context, ruleId string) error {
	_, err := s.client.delete(ctx, "/recast-rules/"+ruleId)
	return err
}

// TargetGroupsService handles communication with the Target Groups related endpoints of the Platform API.
type TargetGroupsService struct {
	client *Client
}

// TargetGroup represents a target group.
type TargetGroup struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type,omitempty"`
	Targets     []string `json:"targets,omitempty"`
	CreatedAt   string   `json:"created_time,omitempty"`
	UpdatedAt   string   `json:"updated_time,omitempty"`
}

// TargetGroupsListResponse represents the response from listing target groups.
type TargetGroupsListResponse struct {
	TargetGroups []TargetGroup `json:"target_groups,omitempty"`
	Total        int           `json:"total,omitempty"`
}

// ListTargetGroups returns all target groups.
func (s *TargetGroupsService) ListTargetGroups(ctx context.Context) (*TargetGroupsListResponse, error) {
	resp, err := s.client.get(ctx, "/target-groups")
	if err != nil {
		return nil, err
	}

	var result TargetGroupsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateTargetGroup creates a new target group.
func (s *TargetGroupsService) CreateTargetGroup(ctx context.Context, req *TargetGroup) (*TargetGroup, error) {
	resp, err := s.client.post(ctx, "/target-groups", req)
	if err != nil {
		return nil, err
	}

	var result TargetGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTargetGroup returns a specific target group by ID.
func (s *TargetGroupsService) GetTargetGroup(ctx context.Context, targetGroupId string) (*TargetGroup, error) {
	resp, err := s.client.get(ctx, "/target-groups/"+targetGroupId)
	if err != nil {
		return nil, err
	}

	var result TargetGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateTargetGroup updates a target group.
func (s *TargetGroupsService) UpdateTargetGroup(ctx context.Context, targetGroupId string, req *TargetGroup) (*TargetGroup, error) {
	resp, err := s.client.put(ctx, "/target-groups/"+targetGroupId, req)
	if err != nil {
		return nil, err
	}

	var result TargetGroup
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteTargetGroup deletes a target group.
func (s *TargetGroupsService) DeleteTargetGroup(ctx context.Context, targetGroupId string) error {
	_, err := s.client.delete(ctx, "/target-groups/"+targetGroupId)
	return err
}

// Permissions2Service is an alias for PermissionsService to avoid conflicts.
type Permissions2Service = PermissionsService
