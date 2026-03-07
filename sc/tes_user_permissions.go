
package sc

import (
	"encoding/json"
	"fmt"
)

// TESUserPermissionsService handles tESUserPermissions operations.
type TESUserPermissionsService struct {
	client *Client
}

// TESUserPermissions represents a tESUserPermissions resource.
type TESUserPermissions struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// TESUserPermissionsListResponse represents the response from listing tESUserPermissionss.
type TESUserPermissionsListResponse struct {
	Usable     []TESUserPermissions `json:"usable"`
	Manageable []TESUserPermissions `json:"manageable"`
}

// List returns all tESUserPermissionss.
func (s *TESUserPermissionsService) List() (*TESUserPermissionsListResponse, error) {
	resp, err := s.client.get("/tes/userPermissions")
	if err != nil {
		return nil, fmt.Errorf("sc: list tESUserPermissionss: %w", err)
	}

	var result TESUserPermissionsListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal tESUserPermissions list response: %w", err)
	}

	return &result, nil
}

