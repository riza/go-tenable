package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// SSHKeyService handles sSHKey operations.
type SSHKeyService struct {
	client *Client
}

// SSHKey represents a sSHKey resource.
type SSHKey struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// SSHKeyListResponse represents the response from listing sSHKeys.
type SSHKeyListResponse struct {
	Usable     []SSHKey `json:"usable"`
	Manageable []SSHKey `json:"manageable"`
}

// SSHKeyCreateInput represents the request body for creating a sSHKey.
type SSHKeyCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// List returns all sSHKeys.
func (s *SSHKeyService) List(ctx context.Context) (*SSHKeyListResponse, error) {
	resp, err := s.client.get(ctx, "/sshKey")
	if err != nil {
		return nil, fmt.Errorf("sc: list sSHKeys: %w", err)
	}

	var result SSHKeyListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal sSHKey list response: %w", err)
	}

	return &result, nil
}

// Create creates a new sSHKey.
func (s *SSHKeyService) Create(ctx context.Context, input *SSHKeyCreateInput) (*SSHKey, error) {
	resp, err := s.client.post(ctx, "/sshKey", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create sSHKey: %w", err)
	}

	var result SSHKey
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal sSHKey response: %w", err)
	}

	return &result, nil
}

// Delete deletes the sSHKey.
func (s *SSHKeyService) Delete(ctx context.Context) error {
	_, err := s.client.delete(ctx, "/sshKey")
	if err != nil {
		return fmt.Errorf("sc: delete sSHKey: %w", err)
	}

	return nil
}

// Download performs the download action on the sSHKey.
func (s *SSHKeyService) Download(ctx context.Context) (*SSHKey, error) {
	resp, err := s.client.get(ctx, "/sshKey/download")
	if err != nil {
		return nil, fmt.Errorf("sc: download sSHKey: %w", err)
	}

	var result SSHKey
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal sSHKey download response: %w", err)
	}

	return &result, nil
}

// InstallRemoteKey performs the installRemoteKey action on the sSHKey.
func (s *SSHKeyService) InstallRemoteKey(ctx context.Context) (*SSHKey, error) {
	resp, err := s.client.post(ctx, "/sshKey/installRemoteKey", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: installRemoteKey sSHKey: %w", err)
	}

	var result SSHKey
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal sSHKey installRemoteKey response: %w", err)
	}

	return &result, nil
}
