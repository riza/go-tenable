
package sc

import (
	"encoding/json"
	"fmt"
)

// LDAPService handles lDAP operations.
type LDAPService struct {
	client *Client
}

// LDAP represents a lDAP resource.
type LDAP struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// LDAPListResponse represents the response from listing lDAPs.
type LDAPListResponse struct {
	Usable     []LDAP `json:"usable"`
	Manageable []LDAP `json:"manageable"`
}

// LDAPCreateInput represents the request body for creating a lDAP.
type LDAPCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// LDAPUpdateInput represents the request body for updating a lDAP.
type LDAPUpdateInput = LDAPCreateInput

// List returns all lDAPs.
func (s *LDAPService) List() (*LDAPListResponse, error) {
	resp, err := s.client.get("/ldap")
	if err != nil {
		return nil, fmt.Errorf("sc: list lDAPs: %w", err)
	}

	var result LDAPListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lDAP list response: %w", err)
	}

	return &result, nil
}

// Create creates a new lDAP.
func (s *LDAPService) Create(input *LDAPCreateInput) (*LDAP, error) {
	resp, err := s.client.post("/ldap", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create lDAP: %w", err)
	}

	var result LDAP
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lDAP response: %w", err)
	}

	return &result, nil
}

// Get returns the lDAP with the given ID.
func (s *LDAPService) Get(id string) (*LDAP, error) {
	resp, err := s.client.get("/ldap" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get lDAP %s: %w", id, err)
	}

	var result LDAP
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lDAP response: %w", err)
	}

	return &result, nil
}

// Update updates the lDAP with the given ID.
func (s *LDAPService) Update(id string, input *LDAPUpdateInput) (*LDAP, error) {
	resp, err := s.client.patch("/ldap" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update lDAP %s: %w", id, err)
	}

	var result LDAP
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lDAP response: %w", err)
	}

	return &result, nil
}

// Delete deletes the lDAP with the given ID.
func (s *LDAPService) Delete(id string) error {
	_, err := s.client.delete("/ldap" + "/" + id)
	if err != nil {
		return fmt.Errorf("sc: delete lDAP %s: %w", id, err)
	}

	return nil
}

// Query performs the query action on the lDAP with the given ID.
func (s *LDAPService) Query(id string) (*LDAP, error) {
	resp, err := s.client.post("/ldap" + "/" + id + "/query", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: query lDAP %s: %w", id, err)
	}

	var result LDAP
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lDAP query response: %w", err)
	}

	return &result, nil
}

// Test performs the test action on the lDAP.
func (s *LDAPService) Test() (*LDAP, error) {
	resp, err := s.client.post("/ldap/test", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: test lDAP: %w", err)
	}

	var result LDAP
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lDAP test response: %w", err)
	}

	return &result, nil
}

