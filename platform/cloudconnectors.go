package platform

import (
	"context"
	"encoding/json"
)

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
