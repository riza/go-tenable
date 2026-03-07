package sc

import (
	"encoding/json"
	"fmt"
)

// AssetService handles communication with the asset-related endpoints of the SC API.
type AssetService struct {
	client *Client
}

// Asset represents an asset resource from the API.
type Asset struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Type            string          `json:"type"`
	Description     string          `json:"description"`
	Tags            string          `json:"tags"`
	Context         string          `json:"context"`
	Status          string          `json:"status"`
	TemplateID      string          `json:"templateID"`
	CreatedTime     string          `json:"createdTime"`
	ModifiedTime    string          `json:"modifiedTime"`
	TypeFields      json.RawMessage `json:"typeFields,omitempty"`
	Repositories    []AssetRepository `json:"repositories,omitempty"`
	IPCount         string          `json:"ipCount"`
	Groups          []IDRef         `json:"groups,omitempty"`
	AssetDataFields []IDRef         `json:"assetDataFields,omitempty"`
	CanUse          string          `json:"canUse"`
	CanManage       string          `json:"canManage"`
	Creator         *AssetUserRef   `json:"creator,omitempty"`
	Owner           *AssetUserRef   `json:"owner,omitempty"`
	OwnerGroup      *AssetGroupRef  `json:"ownerGroup,omitempty"`
	TargetGroup     *AssetGroupRef  `json:"targetGroup,omitempty"`
}

// AssetUserRef is a user reference with extended fields returned by asset endpoints.
type AssetUserRef struct {
	ID        string `json:"id"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

// AssetGroupRef is a group reference with name and description.
type AssetGroupRef struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// AssetRepository represents a repository entry within an asset's repository list.
type AssetRepository struct {
	IPCount    string `json:"ipCount"`
	Repository IDRef  `json:"repository"`
}

// AssetListResponse holds the usable and manageable asset arrays returned by the list endpoint.
type AssetListResponse struct {
	Usable     []Asset `json:"usable"`
	Manageable []Asset `json:"manageable"`
}

// AssetCreateInput contains the fields for creating a new asset.
type AssetCreateInput struct {
	Name        string          `json:"name"`
	Type        string          `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	Tags        string          `json:"tags,omitempty"`
	Context     string          `json:"context,omitempty"`
	TypeFields  json.RawMessage `json:"typeFields,omitempty"`
	Groups      []IDRef         `json:"groups,omitempty"`
}

// AssetUpdateInput contains the fields for updating an existing asset.
type AssetUpdateInput = AssetCreateInput

// AssetImportInput contains the fields for importing an asset.
type AssetImportInput struct {
	Filename string `json:"filename"`
	Name     string `json:"name"`
}

// AssetRefreshInput contains the fields for refreshing an asset.
type AssetRefreshInput struct {
	OrgID  string   `json:"orgID,omitempty"`
	RepIDs []string `json:"repIDs,omitempty"`
}

// AssetLDAPTestInput contains the fields for testing an LDAP query.
type AssetLDAPTestInput struct {
	LDAPQuery string `json:"ldapQuery,omitempty"`
	LDAPID    string `json:"ldapID,omitempty"`
}

// AssetLDAPTestResponse holds the response from testing an LDAP query.
type AssetLDAPTestResponse struct {
	Hostnames []string `json:"hostnames"`
}

// AssetShareInput contains the fields for sharing an asset with groups.
type AssetShareInput struct {
	Groups []IDRef `json:"groups"`
}

// List returns the list of assets (usable and manageable).
func (s *AssetService) List() (*AssetListResponse, error) {
	resp, err := s.client.get("/asset")
	if err != nil {
		return nil, fmt.Errorf("sc: list assets: %w", err)
	}

	var result AssetListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal asset list response: %w", err)
	}

	return &result, nil
}

// Get returns the asset with the given ID.
func (s *AssetService) Get(id string) (*Asset, error) {
	resp, err := s.client.get("/asset/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get asset %s: %w", id, err)
	}

	var asset Asset
	if err := json.Unmarshal(resp.Response, &asset); err != nil {
		return nil, fmt.Errorf("sc: unmarshal asset response: %w", err)
	}

	return &asset, nil
}

// Create creates a new asset with the given input.
func (s *AssetService) Create(input *AssetCreateInput) (*Asset, error) {
	resp, err := s.client.post("/asset", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create asset: %w", err)
	}

	var asset Asset
	if err := json.Unmarshal(resp.Response, &asset); err != nil {
		return nil, fmt.Errorf("sc: unmarshal asset response: %w", err)
	}

	return &asset, nil
}

// Update updates an existing asset with the given input.
func (s *AssetService) Update(id string, input *AssetUpdateInput) (*Asset, error) {
	resp, err := s.client.patch("/asset/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update asset %s: %w", id, err)
	}

	var asset Asset
	if err := json.Unmarshal(resp.Response, &asset); err != nil {
		return nil, fmt.Errorf("sc: unmarshal asset response: %w", err)
	}

	return &asset, nil
}

// Delete deletes the asset with the given ID.
func (s *AssetService) Delete(id string) error {
	_, err := s.client.delete("/asset/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete asset %s: %w", id, err)
	}

	return nil
}

// Import imports an asset definition from a file.
func (s *AssetService) Import(input *AssetImportInput) (*Asset, error) {
	resp, err := s.client.post("/asset/import", input)
	if err != nil {
		return nil, fmt.Errorf("sc: import asset: %w", err)
	}

	var asset Asset
	if err := json.Unmarshal(resp.Response, &asset); err != nil {
		return nil, fmt.Errorf("sc: unmarshal asset import response: %w", err)
	}

	return &asset, nil
}

// Export exports the asset with the given ID. The response is raw XML data.
func (s *AssetService) Export(id string) (json.RawMessage, error) {
	resp, err := s.client.get("/asset/" + id + "/export")
	if err != nil {
		return nil, fmt.Errorf("sc: export asset %s: %w", id, err)
	}

	return resp.Response, nil
}

// Refresh triggers a refresh of the asset with the given ID.
func (s *AssetService) Refresh(id string, input *AssetRefreshInput) (*Asset, error) {
	resp, err := s.client.post("/asset/"+id+"/refresh", input)
	if err != nil {
		return nil, fmt.Errorf("sc: refresh asset %s: %w", id, err)
	}

	var asset Asset
	if err := json.Unmarshal(resp.Response, &asset); err != nil {
		return nil, fmt.Errorf("sc: unmarshal asset refresh response: %w", err)
	}

	return &asset, nil
}

// TestLDAPQuery tests an LDAP query and returns matching hostnames.
func (s *AssetService) TestLDAPQuery(input *AssetLDAPTestInput) (*AssetLDAPTestResponse, error) {
	resp, err := s.client.post("/asset/testLDAPQuery", input)
	if err != nil {
		return nil, fmt.Errorf("sc: test LDAP query: %w", err)
	}

	var result AssetLDAPTestResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal LDAP test response: %w", err)
	}

	return &result, nil
}

// Share shares the asset with the given ID to the specified groups.
func (s *AssetService) Share(id string, input *AssetShareInput) (*Asset, error) {
	resp, err := s.client.post("/asset/"+id+"/share", input)
	if err != nil {
		return nil, fmt.Errorf("sc: share asset %s: %w", id, err)
	}

	var asset Asset
	if err := json.Unmarshal(resp.Response, &asset); err != nil {
		return nil, fmt.Errorf("sc: unmarshal asset share response: %w", err)
	}

	return &asset, nil
}

// Tags returns the list of asset tags.
func (s *AssetService) Tags() ([]string, error) {
	resp, err := s.client.get("/asset/tag")
	if err != nil {
		return nil, fmt.Errorf("sc: list asset tags: %w", err)
	}

	var tags []string
	if err := json.Unmarshal(resp.Response, &tags); err != nil {
		return nil, fmt.Errorf("sc: unmarshal asset tags response: %w", err)
	}

	return tags, nil
}
