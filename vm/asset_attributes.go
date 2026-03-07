package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// AssetAttributesService handles communication with the Asset Attributes related endpoints of the VM API.
type AssetAttributesService struct {
	client *Client
}

type AssetAttributesServiceAssetAttributesCreateRequestAttributesItem struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type AssetAttributesServiceAssetAttributesCreateRequest struct {
	Attributes []AssetAttributesServiceAssetAttributesCreateRequestAttributesItem `json:"attributes,omitempty"`
}

type AssetAttributesServiceAssetAttributesListResponseAttributesItem struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type AssetAttributesServiceAssetAttributesListResponse struct {
	Attributes []AssetAttributesServiceAssetAttributesListResponseAttributesItem `json:"attributes,omitempty"`
}

type AssetAttributesServiceAssetAttributesUpdateRequest struct {
	Description string `json:"description,omitempty"`
}

type AssetAttributesServiceAssetAttributesAssignRequestAttributesItem struct {
	Id    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

type AssetAttributesServiceAssetAttributesAssignRequest struct {
	Attributes []AssetAttributesServiceAssetAttributesAssignRequestAttributesItem `json:"attributes,omitempty"`
}

type AssetAttributesServiceAssetAttributesAssignedListResponseAttributesItem struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type AssetAttributesServiceAssetAttributesAssignedListResponse struct {
	Attributes []AssetAttributesServiceAssetAttributesAssignedListResponseAttributesItem `json:"attributes,omitempty"`
}

type AssetAttributesServiceAssetAttributesSingleUpdateRequest struct {
	Value string `json:"value,omitempty"`
}

// AssetAttributesCreate - Create attribute
func (s *AssetAttributesService) AssetAttributesCreate(ctx context.Context, req *AssetAttributesServiceAssetAttributesCreateRequest) error {
	_, err := s.client.post(ctx, "/api/v3/assets/attributes", req)
	if err != nil {
		return err
	}
	return nil
}

// AssetAttributesList - List attributes
func (s *AssetAttributesService) AssetAttributesList(ctx context.Context) (*AssetAttributesServiceAssetAttributesListResponse, error) {
	resp, err := s.client.get(ctx, "/api/v3/assets/attributes")
	if err != nil {
		return nil, err
	}
	var result AssetAttributesServiceAssetAttributesListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AssetAttributesUpdate - Update attribute
func (s *AssetAttributesService) AssetAttributesUpdate(ctx context.Context, attributeId string, req *AssetAttributesServiceAssetAttributesUpdateRequest) error {
	_, err := s.client.put(ctx, fmt.Sprintf("/api/v3/assets/attributes/%s", attributeId), req)
	if err != nil {
		return err
	}
	return nil
}

// AssetAttributesDelete - Delete attribute
func (s *AssetAttributesService) AssetAttributesDelete(ctx context.Context, attributeId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/api/v3/assets/attributes/%s", attributeId))
	if err != nil {
		return err
	}
	return nil
}

// AssetAttributesAssign - Assign attributes to asset
func (s *AssetAttributesService) AssetAttributesAssign(ctx context.Context, assetId string, req *AssetAttributesServiceAssetAttributesAssignRequest) error {
	_, err := s.client.put(ctx, fmt.Sprintf("/api/v3/assets/%s/attributes", assetId), req)
	if err != nil {
		return err
	}
	return nil
}

// AssetAttributesAssignedList - List attributes assigned to asset
func (s *AssetAttributesService) AssetAttributesAssignedList(ctx context.Context, assetId string) (*AssetAttributesServiceAssetAttributesAssignedListResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/api/v3/assets/%s/attributes", assetId))
	if err != nil {
		return nil, err
	}
	var result AssetAttributesServiceAssetAttributesAssignedListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AssetAttributesAssignedDelete - Delete attributes from asset
func (s *AssetAttributesService) AssetAttributesAssignedDelete(ctx context.Context, assetId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/api/v3/assets/%s/attributes", assetId))
	if err != nil {
		return err
	}
	return nil
}

// AssetAttributesSingleUpdate - Assign single attribute to asset
func (s *AssetAttributesService) AssetAttributesSingleUpdate(ctx context.Context, assetId string, attributeId string, req *AssetAttributesServiceAssetAttributesSingleUpdateRequest) error {
	_, err := s.client.put(ctx, fmt.Sprintf("/api/v3/assets/%s/attributes/%s", assetId, attributeId), req)
	if err != nil {
		return err
	}
	return nil
}

// AssetAttributesSingleDelete - Delete attribute from asset
func (s *AssetAttributesService) AssetAttributesSingleDelete(ctx context.Context, assetId string, attributeId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/api/v3/assets/%s/attributes/%s", assetId, attributeId))
	if err != nil {
		return err
	}
	return nil
}
