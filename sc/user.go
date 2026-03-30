package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// UserService handles communication with the user-related endpoints of the SC API.
type UserService struct {
	client *Client
}

// User represents a user resource from the API.
type User struct {
	ID                         string            `json:"id"`
	UUID                       string            `json:"uuid"`
	Status                     string            `json:"status"`
	Username                   string            `json:"username"`
	LDAPUsername               string            `json:"ldapUsername"`
	Firstname                  string            `json:"firstname"`
	Lastname                   string            `json:"lastname"`
	Title                      string            `json:"title"`
	Email                      string            `json:"email"`
	Address                    string            `json:"address"`
	City                       string            `json:"city"`
	State                      string            `json:"state"`
	Country                    string            `json:"country"`
	Phone                      string            `json:"phone"`
	Fax                        string            `json:"fax"`
	CreatedTime                string            `json:"createdTime"`
	ModifiedTime               string            `json:"modifiedTime"`
	LastLogin                  string            `json:"lastLogin"`
	LastLoginIP                string            `json:"lastLoginIP"`
	MustChangePassword         string            `json:"mustChangePassword"`
	PasswordExpires            string            `json:"passwordExpires"`
	PasswordExpiration         string            `json:"passwordExpiration"`
	PasswordExpirationOverride string            `json:"passwordExpirationOverride"`
	PasswordSetDate            string            `json:"passwordSetDate"`
	Locked                     string            `json:"locked"`
	FailedLogins               string            `json:"failedLogins"`
	AuthType                   string            `json:"authType"`
	Fingerprint                string            `json:"fingerprint"`
	Password                   string            `json:"password"`
	CanUse                     string            `json:"canUse"`
	CanManage                  string            `json:"canManage"`
	Role                       *IDRef            `json:"role"`
	LDAP                       *IDRef            `json:"ldap"`
	ResponsibleAsset           *IDRef            `json:"responsibleAsset"`
	Group                      *IDRef            `json:"group"`
	ManagedUsersGroups         []IDRef           `json:"managedUsersGroups"`
	ManagedObjectsGroups       []IDRef           `json:"managedObjectsGroups"`
	Preferences                []UserPreference  `json:"preferences"`
	APIKeys                    []json.RawMessage `json:"apiKeys"`
	Parent                     json.RawMessage   `json:"parent"`
}

// UserPreference represents a single user preference entry.
type UserPreference struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
}

// UserListResponse holds the paginated response from listing users.
type UserListResponse struct {
	PaginatedResponse
	Results []User `json:"results"`
}

// UserCreateInput contains the fields for creating a new user.
type UserCreateInput struct {
	Username    string           `json:"username"`
	RoleID      string           `json:"roleID"`
	AuthType    string           `json:"authType"`
	Password    string           `json:"password,omitempty"`
	Firstname   string           `json:"firstname,omitempty"`
	Lastname    string           `json:"lastname,omitempty"`
	Title       string           `json:"title,omitempty"`
	Email       string           `json:"email,omitempty"`
	Address     string           `json:"address,omitempty"`
	City        string           `json:"city,omitempty"`
	State       string           `json:"state,omitempty"`
	Country     string           `json:"country,omitempty"`
	Phone       string           `json:"phone,omitempty"`
	Fax         string           `json:"fax,omitempty"`
	Locked      string           `json:"locked,omitempty"`
	Fingerprint string           `json:"fingerprint,omitempty"`
	EmailNotice string           `json:"emailNotice,omitempty"`
	Preferences []UserPreference `json:"preferences,omitempty"`
}

// UserUpdateInput contains the fields for updating an existing user.
type UserUpdateInput = UserCreateInput

// UserDeleteInput contains optional parameters for deleting a user.
type UserDeleteInput struct {
	MigrateUserID string `json:"migrateUserID,omitempty"`
}

// List returns the paginated list of users.
func (s *UserService) List(ctx context.Context) (*UserListResponse, error) {
	resp, err := s.client.get(ctx, "/user")
	if err != nil {
		return nil, fmt.Errorf("sc: list users: %w", err)
	}

	var result UserListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal user list response: %w", err)
	}

	return &result, nil
}

// Get returns the user with the given ID.
func (s *UserService) Get(ctx context.Context, id string) (*User, error) {
	resp, err := s.client.get(ctx, "/user/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get user %s: %w", id, err)
	}

	var user User
	if err := json.Unmarshal(resp.Response, &user); err != nil {
		return nil, fmt.Errorf("sc: unmarshal user response: %w", err)
	}

	return &user, nil
}

// Create creates a new user with the given input.
func (s *UserService) Create(ctx context.Context, input *UserCreateInput) (*User, error) {
	resp, err := s.client.post(ctx, "/user", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create user: %w", err)
	}

	var user User
	if err := json.Unmarshal(resp.Response, &user); err != nil {
		return nil, fmt.Errorf("sc: unmarshal user response: %w", err)
	}

	return &user, nil
}

// Update updates an existing user with the given input.
func (s *UserService) Update(ctx context.Context, id string, input *UserUpdateInput) (*User, error) {
	resp, err := s.client.patch(ctx, "/user/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update user %s: %w", id, err)
	}

	var user User
	if err := json.Unmarshal(resp.Response, &user); err != nil {
		return nil, fmt.Errorf("sc: unmarshal user response: %w", err)
	}

	return &user, nil
}

// Delete deletes the user with the given ID.
// If input is non-nil and MigrateUserID is set, the user's objects are migrated to that user.
func (s *UserService) Delete(ctx context.Context, id string, input *UserDeleteInput) error {
	path := "/user/" + id
	if input != nil && input.MigrateUserID != "" {
		path += "?migrateUserID=" + input.MigrateUserID
	}
	_, err := s.client.delete(ctx, path)
	if err != nil {
		return fmt.Errorf("sc: delete user %s: %w", id, err)
	}

	return nil
}
