package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// ConfigurationService handles communication with the configuration-related
// endpoints of the SC API.
type ConfigurationService struct {
	client *Client
}

// ConfigType represents a configuration category returned by the list endpoint.
type ConfigType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ConfigDetail is a type alias for the dynamic key-value map returned by the
// config detail endpoint. The shape varies per configuration type, so it is
// left as raw JSON.
type ConfigDetail = json.RawMessage

// ConfigQueryItem represents a single item in the config query response.
type ConfigQueryItem struct {
	Item       string          `json:"item"`
	Configured bool            `json:"configured"`
	Details    json.RawMessage `json:"details"`
}

// TestSMTPInput contains the fields for testing SMTP configuration.
type TestSMTPInput struct {
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	Encrypt  string `json:"encrypt,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// TestSMTPResponse holds the response from testing SMTP configuration.
type TestSMTPResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// LicenseRegisterInput contains the fields for registering a license.
type LicenseRegisterInput struct {
	Filename string `json:"filename"`
}

// LicenseRegisterResponse is a type alias for the dynamic response returned
// by the license register endpoint.
type LicenseRegisterResponse = json.RawMessage

// PluginRegisterInput contains the fields for registering plugins.
type PluginRegisterInput struct {
	ActivationCode string `json:"activationCode"`
	UpdateSite     string `json:"updateSite,omitempty"`
	Type           string `json:"type,omitempty"`
}

// PluginRegisterResponse holds the response from registering plugins.
type PluginRegisterResponse struct {
	PluginSubscriptionStatus string `json:"pluginSubscriptionStatus"`
}

// PluginResetInput contains the fields for resetting plugins.
type PluginResetInput struct {
	Type string `json:"type"`
}

// List returns all configuration types.
func (s *ConfigurationService) List(ctx context.Context) ([]ConfigType, error) {
	resp, err := s.client.get(ctx, "/config")
	if err != nil {
		return nil, fmt.Errorf("sc: list configurations: %w", err)
	}

	var result []ConfigType
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal configuration list response: %w", err)
	}

	return result, nil
}

// Get returns the configuration detail for the given ID. The response is a
// dynamic key-value map that varies per configuration type.
func (s *ConfigurationService) Get(ctx context.Context, id string) (json.RawMessage, error) {
	resp, err := s.client.get(ctx, "/config/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get configuration %s: %w", id, err)
	}

	return resp.Response, nil
}

// Update updates the configuration with the given ID using a map of key-value
// pairs. The response is a dynamic key-value map.
func (s *ConfigurationService) Update(ctx context.Context, id string, input map[string]string) (json.RawMessage, error) {
	resp, err := s.client.patch(ctx, "/config/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update configuration %s: %w", id, err)
	}

	return resp.Response, nil
}

// Query queries the configuration for the given comma-separated items
// (e.g. "smtp,ldap").
func (s *ConfigurationService) Query(ctx context.Context, items string) ([]ConfigQueryItem, error) {
	path := fmt.Sprintf("/config/query?item=%s", items)
	resp, err := s.client.get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("sc: query configuration: %w", err)
	}

	var result []ConfigQueryItem
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal configuration query response: %w", err)
	}

	return result, nil
}

// TestSMTP tests the SMTP configuration.
func (s *ConfigurationService) TestSMTP(ctx context.Context, input *TestSMTPInput) (*TestSMTPResponse, error) {
	resp, err := s.client.post(ctx, "/config/testSMTP", input)
	if err != nil {
		return nil, fmt.Errorf("sc: test SMTP configuration: %w", err)
	}

	var result TestSMTPResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal test SMTP response: %w", err)
	}

	return &result, nil
}

// RegisterLicense registers a license file.
func (s *ConfigurationService) RegisterLicense(ctx context.Context, input *LicenseRegisterInput) (json.RawMessage, error) {
	resp, err := s.client.post(ctx, "/config/license/register", input)
	if err != nil {
		return nil, fmt.Errorf("sc: register license: %w", err)
	}

	return resp.Response, nil
}

// RegisterPlugins registers plugins with an activation code.
func (s *ConfigurationService) RegisterPlugins(ctx context.Context, input *PluginRegisterInput) (*PluginRegisterResponse, error) {
	resp, err := s.client.post(ctx, "/config/plugins/register", input)
	if err != nil {
		return nil, fmt.Errorf("sc: register plugins: %w", err)
	}

	var result PluginRegisterResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal plugin register response: %w", err)
	}

	return &result, nil
}

// ResetPlugins resets plugins of the given type.
func (s *ConfigurationService) ResetPlugins(ctx context.Context, input *PluginResetInput) error {
	_, err := s.client.post(ctx, "/config/plugins/reset", input)
	if err != nil {
		return fmt.Errorf("sc: reset plugins: %w", err)
	}

	return nil
}
