package one

import (
	"context"
	"encoding/json"
)

// InventoryService handles communication with the Inventory related endpoints of the Tenable One API.
type InventoryService struct {
	client *Client
}

// InventoryAsset represents an asset in the inventory.
type InventoryAsset struct {
	Id              string                 `json:"id,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Type            string                 `json:"type,omitempty"`
	Sources         []string               `json:"sources,omitempty"`
	FirstSeen       string                 `json:"first_seen,omitempty"`
	LastSeen        string                 `json:"last_seen,omitempty"`
	Tags            []string               `json:"tags,omitempty"`
	NetworkId       string                 `json:"network_id,omitempty"`
	IpAddresses     []string               `json:"ip_addresses,omitempty"`
	Hostnames       []string               `json:"hostnames,omitempty"`
	MacAddresses    []string               `json:"mac_addresses,omitempty"`
	OperatingSystem []string               `json:"operating_system,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
}

// InventoryAssetsSearchRequest represents the request body for searching assets.
type InventoryAssetsSearchRequest struct {
	Limit   *int                    `json:"limit,omitempty"`
	Offset  *int                    `json:"offset,omitempty"`
	Filters []InventorySearchFilter `json:"filters,omitempty"`
}

// InventorySearchFilter represents a generic filter for inventory search.
type InventorySearchFilter struct {
	Property string      `json:"property"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

// InventoryAssetsSearchResponse represents the response from searching assets.
type InventoryAssetsSearchResponse struct {
	Assets     []InventoryAsset `json:"assets,omitempty"`
	Total      int              `json:"total,omitempty"`
	Pagination interface{}      `json:"pagination,omitempty"`
}

// SearchAssets searches for assets in the inventory.
func (s *InventoryService) SearchAssets(ctx context.Context, req *InventoryAssetsSearchRequest) (*InventoryAssetsSearchResponse, error) {
	resp, err := s.client.post(ctx, "/api/v1/t1/inventory/assets/search", req)
	if err != nil {
		return nil, err
	}

	var result InventoryAssetsSearchResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// InventoryFinding represents a finding in the inventory.
type InventoryFinding struct {
	Id          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Severity    string                 `json:"severity,omitempty"`
	Status      string                 `json:"status,omitempty"`
	PluginId    int                    `json:"plugin_id,omitempty"`
	PluginName  string                 `json:"plugin_name,omitempty"`
	AssetId     string                 `json:"asset_id,omitempty"`
	FirstSeen   string                 `json:"first_seen,omitempty"`
	LastSeen    string                 `json:"last_seen,omitempty"`
	Output      string                 `json:"output,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// InventoryFindingsSearchRequest represents the request body for searching findings.
type InventoryFindingsSearchRequest struct {
	Limit   *int                    `json:"limit,omitempty"`
	Offset  *int                    `json:"offset,omitempty"`
	Filters []InventorySearchFilter `json:"filters,omitempty"`
}

// InventoryFindingsSearchResponse represents the response from searching findings.
type InventoryFindingsSearchResponse struct {
	Findings   []InventoryFinding `json:"findings,omitempty"`
	Total      int                `json:"total,omitempty"`
	Pagination interface{}        `json:"pagination,omitempty"`
}

// SearchFindings searches for findings in the inventory.
func (s *InventoryService) SearchFindings(ctx context.Context, req *InventoryFindingsSearchRequest) (*InventoryFindingsSearchResponse, error) {
	resp, err := s.client.post(ctx, "/api/v1/t1/inventory/findings/search", req)
	if err != nil {
		return nil, err
	}

	var result InventoryFindingsSearchResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// InventorySoftware represents software in the inventory.
type InventorySoftware struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Vendor    string `json:"vendor,omitempty"`
	Version   string `json:"version,omitempty"`
	Cpe       string `json:"cpe,omitempty"`
	AssetId   string `json:"asset_id,omitempty"`
	FirstSeen string `json:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty"`
}

// InventorySoftwareSearchRequest represents the request body for searching software.
type InventorySoftwareSearchRequest struct {
	Limit   *int                    `json:"limit,omitempty"`
	Offset  *int                    `json:"offset,omitempty"`
	Filters []InventorySearchFilter `json:"filters,omitempty"`
}

// InventorySoftwareSearchResponse represents the response from searching software.
type InventorySoftwareSearchResponse struct {
	Software   []InventorySoftware `json:"software,omitempty"`
	Total      int                 `json:"total,omitempty"`
	Pagination interface{}         `json:"pagination,omitempty"`
}

// SearchSoftware searches for software in the inventory.
func (s *InventoryService) SearchSoftware(ctx context.Context, req *InventorySoftwareSearchRequest) (*InventorySoftwareSearchResponse, error) {
	resp, err := s.client.post(ctx, "/api/v1/t1/inventory/software/search", req)
	if err != nil {
		return nil, err
	}

	var result InventorySoftwareSearchResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// InventoryPropertiesResponse represents the response from getting inventory properties.
type InventoryPropertiesResponse struct {
	Properties []map[string]interface{} `json:"properties,omitempty"`
	Total      int                      `json:"total,omitempty"`
}

// GetAssetProperties returns available properties for asset search.
func (s *InventoryService) GetAssetProperties(ctx context.Context) (*InventoryPropertiesResponse, error) {
	resp, err := s.client.get(ctx, "/api/v1/t1/inventory/assets/properties")
	if err != nil {
		return nil, err
	}

	var result InventoryPropertiesResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetFindingProperties returns available properties for finding search.
func (s *InventoryService) GetFindingProperties(ctx context.Context) (*InventoryPropertiesResponse, error) {
	resp, err := s.client.get(ctx, "/api/v1/t1/inventory/findings/properties")
	if err != nil {
		return nil, err
	}

	var result InventoryPropertiesResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetSoftwareProperties returns available properties for software search.
func (s *InventoryService) GetSoftwareProperties(ctx context.Context) (*InventoryPropertiesResponse, error) {
	resp, err := s.client.get(ctx, "/api/v1/t1/inventory/software/properties")
	if err != nil {
		return nil, err
	}

	var result InventoryPropertiesResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
