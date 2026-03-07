package sc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CurrentUserService handles communication with the currentUser-related endpoints of the SC API.
type CurrentUserService struct {
	client *Client
}

// CurrentUser represents the currently authenticated user resource from the API.
type CurrentUser struct {
	ID                         string            `json:"id"`
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
	Locked                     string            `json:"locked"`
	FailedLogins               string            `json:"failedLogins"`
	AuthType                   string            `json:"authType"`
	Fingerprint                string            `json:"fingerprint"`
	Password                   string            `json:"password"`
	ManagedUsersGroups         []IDRef           `json:"managedUsersGroups"`
	ManagedObjectsGroups       []IDRef           `json:"managedObjectsGroups"`
	Preferences                []UserPreference  `json:"preferences"`
	Organization               *IDRef            `json:"organization"`
	Role                       *IDRef            `json:"role"`
	Group                      *IDRef            `json:"group"`
	LDAP                       *IDRef            `json:"ldap"`
	ResponsibleAsset           *IDRef            `json:"responsibleAsset"`
	LinkedUserRole             *IDRef            `json:"linkedUserRole"`
	OrgName                    string            `json:"orgName"`
	UUID                       string            `json:"uuid"`
	SwitchableUsers            []json.RawMessage `json:"switchableUsers"`
}

// CurrentUserUpdateInput contains the fields for updating the current user.
type CurrentUserUpdateInput struct {
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
	Fingerprint string           `json:"fingerprint,omitempty"`
	Password    string           `json:"password,omitempty"`
	Preferences []UserPreference `json:"preferences,omitempty"`
}

// PreferenceInput contains the fields for managing user preferences.
type PreferenceInput struct {
	Name  string `json:"name,omitempty"`
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

// SwitchInput contains the fields for switching the current user context.
type SwitchInput struct {
	Username string `json:"username"`
}

// Get returns the currently authenticated user.
func (s *CurrentUserService) Get() (*CurrentUser, error) {
	resp, err := s.client.get("/currentUser")
	if err != nil {
		return nil, fmt.Errorf("sc: get current user: %w", err)
	}

	var result CurrentUser
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal current user response: %w", err)
	}

	return &result, nil
}

// Update updates the currently authenticated user with the given input.
func (s *CurrentUserService) Update(input *CurrentUserUpdateInput) (*CurrentUser, error) {
	resp, err := s.client.patch("/currentUser", input)
	if err != nil {
		return nil, fmt.Errorf("sc: update current user: %w", err)
	}

	var result CurrentUser
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal current user response: %w", err)
	}

	return &result, nil
}

// AssociateCert associates a client certificate with the current user.
func (s *CurrentUserService) AssociateCert() (*CurrentUser, error) {
	resp, err := s.client.post("/currentUser/associateCert", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: associate cert for current user: %w", err)
	}

	var result CurrentUser
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal current user response: %w", err)
	}

	return &result, nil
}

// GetPreferences returns the preferences for the current user.
func (s *CurrentUserService) GetPreferences(input *PreferenceInput) ([]UserPreference, error) {
	resp, err := s.client.get("/currentUser/preferences")
	if err != nil {
		return nil, fmt.Errorf("sc: get current user preferences: %w", err)
	}

	var result []UserPreference
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal current user preferences response: %w", err)
	}

	return result, nil
}

// UpdatePreference updates a preference for the current user.
func (s *CurrentUserService) UpdatePreference(input *PreferenceInput) ([]UserPreference, error) {
	resp, err := s.client.patch("/currentUser/preferences", input)
	if err != nil {
		return nil, fmt.Errorf("sc: update current user preference: %w", err)
	}

	var result []UserPreference
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal current user preferences response: %w", err)
	}

	return result, nil
}

// DeletePreferences deletes preferences for the current user.
// The SC API requires a request body with the DELETE method for this endpoint.
func (s *CurrentUserService) DeletePreferences(input *PreferenceInput) error {
	_, err := s.client.doRequest(http.MethodDelete, "/currentUser/preferences", input)
	if err != nil {
		return fmt.Errorf("sc: delete current user preferences: %w", err)
	}

	return nil
}

// Switch switches the current user context to another user.
func (s *CurrentUserService) Switch(input *SwitchInput) (*CurrentUser, error) {
	resp, err := s.client.post("/currentUser/switch", input)
	if err != nil {
		return nil, fmt.Errorf("sc: switch current user: %w", err)
	}

	var result CurrentUser
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal current user response: %w", err)
	}

	return &result, nil
}
