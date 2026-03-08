package platform

import (
	"context"
	"encoding/json"
	"fmt"
)

// UsersService handles communication with the Users related endpoints of the Platform API.
type UsersService struct {
	client *Client
}

// User represents a user.
type User struct {
	Id             int                    `json:"id,omitempty"`
	Uuid           string                 `json:"uuid,omitempty"`
	Username       string                 `json:"username,omitempty"`
	Email          string                 `json:"email,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Enabled        bool                   `json:"enabled,omitempty"`
	CreatedAt      string                 `json:"created_time,omitempty"`
	LastAuth       string                 `json:"last_auth,omitempty"`
	Permission     string                 `json:"permission,omitempty"`
	Role           string                 `json:"role,omitempty"`
	AuthType       string                 `json:"auth_type,omitempty"`
	Timezone       string                 `json:"timezone,omitempty"`
}

// UsersListResponse represents the response from listing users.
type UsersListResponse struct {
	Users []User `json:"users,omitempty"`
	Total int    `json:"total,omitempty"`
}

// ListUsers returns all users.
func (s *UsersService) ListUsers(ctx context.Context) (*UsersListResponse, error) {
	resp, err := s.client.get(ctx, "/users")
	if err != nil {
		return nil, err
	}

	var result UsersListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateUser creates a new user.
func (s *UsersService) CreateUser(ctx context.Context, req *User) (*User, error) {
	resp, err := s.client.post(ctx, "/users", req)
	if err != nil {
		return nil, err
	}

	var result User
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetUser returns a specific user by ID.
func (s *UsersService) GetUser(ctx context.Context, userId string) (*User, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/users/%s", userId))
	if err != nil {
		return nil, err
	}

	var result User
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateUser updates a user.
func (s *UsersService) UpdateUser(ctx context.Context, userId string, req *User) (*User, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/users/%s", userId), req)
	if err != nil {
		return nil, err
	}

	var result User
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteUser deletes a user.
func (s *UsersService) DeleteUser(ctx context.Context, userId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/users/%s", userId))
	return err
}

// UserRolesResponse represents the response from getting user roles.
type UserRolesResponse struct {
	Roles []string `json:"roles,omitempty"`
}

// GetUserRoles returns roles for a specific user.
func (s *UsersService) GetUserRoles(ctx context.Context, userUuid string) (*UserRolesResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/access-control/v1/users/%s/roles", userUuid))
	if err != nil {
		return nil, err
	}

	var result UserRolesResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateUserRoles updates roles for a specific user.
func (s *UsersService) UpdateUserRoles(ctx context.Context, userUuid string, req *UserRolesResponse) (*UserRolesResponse, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/access-control/v1/users/%s/roles", userUuid), req)
	if err != nil {
		return nil, err
	}

	var result UserRolesResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ChangePasswordRequest represents a change password request.
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password,omitempty"`
	NewPassword     string `json:"new_password,omitempty"`
}

// ChangePassword changes the password for a user.
func (s *UsersService) ChangePassword(ctx context.Context, userUuid string, req *ChangePasswordRequest) error {
	_, err := s.client.put(ctx, fmt.Sprintf("/users/%s/chpasswd", userUuid), req)
	return err
}

// SetUserEnabled enables or disables a user.
func (s *UsersService) SetUserEnabled(ctx context.Context, userId string, enabled bool) error {
	_, err := s.client.put(ctx, fmt.Sprintf("/users/%s/enabled", userId), map[string]bool{"enabled": enabled})
	return err
}

// UserAuthorizationsResponse represents the response from getting user authorizations.
type UserAuthorizationsResponse struct {
	Authorizations []string `json:"authorizations,omitempty"`
}

// GetUserAuthorizations returns authorizations for a specific user.
func (s *UsersService) GetUserAuthorizations(ctx context.Context, userId string) (*UserAuthorizationsResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/users/%s/authorizations", userId))
	if err != nil {
		return nil, err
	}

	var result UserAuthorizationsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateUserAuthorizations updates authorizations for a specific user.
func (s *UsersService) UpdateUserAuthorizations(ctx context.Context, userId string, req *UserAuthorizationsResponse) (*UserAuthorizationsResponse, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/users/%s/authorizations", userId), req)
	if err != nil {
		return nil, err
	}

	var result UserAuthorizationsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ImpersonateUser impersonates a user.
func (s *UsersService) ImpersonateUser(ctx context.Context, userId string) (*User, error) {
	resp, err := s.client.post(ctx, fmt.Sprintf("/users/%s/impersonate", userId), nil)
	if err != nil {
		return nil, err
	}

	var result User
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
