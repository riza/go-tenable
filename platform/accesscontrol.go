package platform

import (
	"context"
	"encoding/json"
	"fmt"
)

// AccessControlService handles communication with the Access Control related endpoints of the Platform API.
type AccessControlService struct {
	client *Client
}

// APISecuritySettings represents API security settings.
type APISecuritySettings struct {
	APIAccessKeyEnabled bool `json:"api_access_key_enabled,omitempty"`
	APIKeyIdleTimeout   int  `json:"api_key_idle_timeout,omitempty"`
	APIKeyMaxLifetime   int  `json:"api_key_max_lifetime,omitempty"`
}

// GetAPISecuritySettings returns the API security settings.
func (s *AccessControlService) GetAPISecuritySettings(ctx context.Context) (*APISecuritySettings, error) {
	resp, err := s.client.get(ctx, "/access-control/v1/api-security-settings")
	if err != nil {
		return nil, err
	}

	var result APISecuritySettings
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateAPISecuritySettings updates the API security settings.
func (s *AccessControlService) UpdateAPISecuritySettings(ctx context.Context, req *APISecuritySettings) (*APISecuritySettings, error) {
	resp, err := s.client.put(ctx, "/access-control/v1/api-security-settings", req)
	if err != nil {
		return nil, err
	}

	var result APISecuritySettings
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GroupsService handles communication with the Groups related endpoints of the Platform API.
type GroupsService struct {
	client *Client
}

// Group represents a user group.
type Group struct {
	Id            int    `json:"id,omitempty"`
	Uuid          string `json:"uuid,omitempty"`
	Name          string `json:"name,omitempty"`
	Permissions   int    `json:"permissions,omitempty"`
	ContainerUuid string `json:"container_uuid,omitempty"`
	UserCount     int    `json:"user_count,omitempty"`
	Description   string `json:"description,omitempty"`
	CreatedAt     string `json:"created_time,omitempty"`
	UpdatedAt     string `json:"modified_time,omitempty"`
}

// GroupsListResponse represents the response from listing groups.
type GroupsListResponse struct {
	Groups []Group `json:"groups,omitempty"`
	Total  int     `json:"total,omitempty"`
}

// ListGroups returns all groups.
func (s *GroupsService) ListGroups(ctx context.Context) (*GroupsListResponse, error) {
	resp, err := s.client.get(ctx, "/groups")
	if err != nil {
		return nil, err
	}

	var result GroupsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateGroup creates a new group.
func (s *GroupsService) CreateGroup(ctx context.Context, req *Group) (*Group, error) {
	resp, err := s.client.post(ctx, "/groups", req)
	if err != nil {
		return nil, err
	}

	var result Group
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGroup returns a specific group by ID.
func (s *GroupsService) GetGroup(ctx context.Context, groupId string) (*Group, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/groups/%s", groupId))
	if err != nil {
		return nil, err
	}

	var result Group
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateGroup updates a group.
func (s *GroupsService) UpdateGroup(ctx context.Context, groupId string, req *Group) (*Group, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/groups/%s", groupId), req)
	if err != nil {
		return nil, err
	}

	var result Group
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteGroup deletes a group.
func (s *GroupsService) DeleteGroup(ctx context.Context, groupId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/groups/%s", groupId))
	return err
}

// GroupUsersResponse represents the response from listing group users.
type GroupUsersResponse struct {
	Users []User `json:"users,omitempty"`
	Total int    `json:"total,omitempty"`
}

// ListGroupUsers returns all users in a group.
func (s *GroupsService) ListGroupUsers(ctx context.Context, groupId string) (*GroupUsersResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/groups/%s/users", groupId))
	if err != nil {
		return nil, err
	}

	var result GroupUsersResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// AddUserToGroup adds a user to a group.
func (s *GroupsService) AddUserToGroup(ctx context.Context, groupId, userId string) error {
	_, err := s.client.post(ctx, fmt.Sprintf("/groups/%s/users/%s", groupId, userId), nil)
	return err
}

// RemoveUserFromGroup removes a user from a group.
func (s *GroupsService) RemoveUserFromGroup(ctx context.Context, groupId, userId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/groups/%s/users/%s", groupId, userId))
	return err
}

// PermissionsService handles communication with the Permissions related endpoints of the Platform API.
type PermissionsService struct {
	client *Client
}

// Permission represents a permission.
type Permission struct {
	Uuid       string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
	Permission string `json:"permission,omitempty"`
}

// PermissionsListResponse represents the response from listing permissions.
type PermissionsListResponse struct {
	Permissions []Permission `json:"permissions,omitempty"`
	Total       int          `json:"total,omitempty"`
}

// ListPermissions returns all permissions.
func (s *PermissionsService) ListPermissions(ctx context.Context) (*PermissionsListResponse, error) {
	resp, err := s.client.get(ctx, "/api/v3/access-control/permissions")
	if err != nil {
		return nil, err
	}

	var result PermissionsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreatePermission creates a new permission.
func (s *PermissionsService) CreatePermission(ctx context.Context, req *Permission) (*Permission, error) {
	resp, err := s.client.post(ctx, "/api/v3/access-control/permissions", req)
	if err != nil {
		return nil, err
	}

	var result Permission
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetPermission returns a specific permission by UUID.
func (s *PermissionsService) GetPermission(ctx context.Context, permissionUuid string) (*Permission, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/api/v3/access-control/permissions/%s", permissionUuid))
	if err != nil {
		return nil, err
	}

	var result Permission
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdatePermission updates a permission.
func (s *PermissionsService) UpdatePermission(ctx context.Context, permissionUuid string, req *Permission) (*Permission, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/api/v3/access-control/permissions/%s", permissionUuid), req)
	if err != nil {
		return nil, err
	}

	var result Permission
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeletePermission deletes a permission.
func (s *PermissionsService) DeletePermission(ctx context.Context, permissionUuid string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/api/v3/access-control/permissions/%s", permissionUuid))
	return err
}

// GetUserPermissions returns permissions for a specific user.
func (s *PermissionsService) GetUserPermissions(ctx context.Context, userUuid string) (*PermissionsListResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/api/v3/access-control/permissions/users/%s", userUuid))
	if err != nil {
		return nil, err
	}

	var result PermissionsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetUserGroupPermissions returns permissions for a specific user group.
func (s *PermissionsService) GetUserGroupPermissions(ctx context.Context, userGroupUuid string) (*PermissionsListResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/api/v3/access-control/permissions/user-groups/%s", userGroupUuid))
	if err != nil {
		return nil, err
	}

	var result PermissionsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetMyPermissions returns permissions for the current user.
func (s *PermissionsService) GetMyPermissions(ctx context.Context) (*PermissionsListResponse, error) {
	resp, err := s.client.get(ctx, "/api/v3/access-control/permissions/users/me")
	if err != nil {
		return nil, err
	}

	var result PermissionsListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
