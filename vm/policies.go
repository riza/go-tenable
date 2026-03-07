package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// PoliciesService handles communication with the Policies related endpoints of the VM API.
type PoliciesService struct {
	client *Client
}

type PoliciesServicePoliciesCreateRequestSettings struct {
	Name string `json:"name,omitempty"`
}

type PoliciesServicePoliciesCreateRequest struct {
	Uuid     string                                       `json:"uuid,omitempty"`
	Settings PoliciesServicePoliciesCreateRequestSettings `json:"settings,omitempty"`
}

type PoliciesServicePoliciesCreateResponse struct {
	PolicyId   int    `json:"policy_id,omitempty"`
	PolicyName string `json:"policy_name,omitempty"`
}

type PoliciesServicePoliciesListResponseItem struct {
	Id                   int    `json:"id,omitempty"`
	TemplateUuid         string `json:"template_uuid,omitempty"`
	Name                 string `json:"name,omitempty"`
	Description          string `json:"description,omitempty"`
	OwnerId              string `json:"owner_id,omitempty"`
	Owner                string `json:"owner,omitempty"`
	Shared               int    `json:"shared,omitempty"`
	UserPermissions      int    `json:"user_permissions,omitempty"`
	CreationDate         int    `json:"creation_date,omitempty"`
	LastModificationDate int    `json:"last_modification_date,omitempty"`
	Visibility           int    `json:"visibility,omitempty"`
	NoTarget             bool   `json:"no_target,omitempty"`
}

type PoliciesServicePoliciesImportRequest struct {
	File string `json:"file,omitempty"`
}

type PoliciesServicePoliciesImportResponse struct {
	Private              int    `json:"private,omitempty"`
	NoTarget             string `json:"no_target,omitempty"`
	TemplateUuid         string `json:"template_uuid,omitempty"`
	Description          string `json:"description,omitempty"`
	Name                 string `json:"name,omitempty"`
	Owner                string `json:"owner,omitempty"`
	Shared               int    `json:"shared,omitempty"`
	UserPermissions      int    `json:"user_permissions,omitempty"`
	LastModificationDate int    `json:"last_modification_date,omitempty"`
	CreationDate         int    `json:"creation_date,omitempty"`
	OwnerId              int    `json:"owner_id,omitempty"`
	Id                   int    `json:"id,omitempty"`
}

type PoliciesServicePoliciesDetailsResponse struct {
	Uuid        string                 `json:"uuid,omitempty"`
	Audits      map[string]interface{} `json:"audits,omitempty"`
	Credentials map[string]interface{} `json:"credentials,omitempty"`
	Plugins     map[string]interface{} `json:"plugins,omitempty"`
	Scap        map[string]interface{} `json:"scap,omitempty"`
	Settings    map[string]interface{} `json:"settings,omitempty"`
}

// PoliciesCreate - Create policy
func (s *PoliciesService) PoliciesCreate(ctx context.Context, req *PoliciesServicePoliciesCreateRequest) (*PoliciesServicePoliciesCreateResponse, error) {
	resp, err := s.client.post(ctx, "/policies", req)
	if err != nil {
		return nil, err
	}
	var result PoliciesServicePoliciesCreateResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PoliciesList - List policies
func (s *PoliciesService) PoliciesList(ctx context.Context) ([]PoliciesServicePoliciesListResponseItem, error) {
	resp, err := s.client.get(ctx, "/policies")
	if err != nil {
		return nil, err
	}
	var result []PoliciesServicePoliciesListResponseItem
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// PoliciesCopy - Copy policy
func (s *PoliciesService) PoliciesCopy(ctx context.Context, policyId string) error {
	_, err := s.client.post(ctx, fmt.Sprintf("/policies/%s/copy", policyId), nil)
	if err != nil {
		return err
	}
	return nil
}

// PoliciesImport - Import policy
func (s *PoliciesService) PoliciesImport(ctx context.Context, req *PoliciesServicePoliciesImportRequest) (*PoliciesServicePoliciesImportResponse, error) {
	resp, err := s.client.post(ctx, "/policies/import", req)
	if err != nil {
		return nil, err
	}
	var result PoliciesServicePoliciesImportResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PoliciesExport - Export policy
func (s *PoliciesService) PoliciesExport(ctx context.Context, policyId string) error {
	_, err := s.client.get(ctx, fmt.Sprintf("/policies/%s/export", policyId))
	if err != nil {
		return err
	}
	return nil
}

// PoliciesDetails - List policy details
func (s *PoliciesService) PoliciesDetails(ctx context.Context, policyId string) (*PoliciesServicePoliciesDetailsResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/policies/%s", policyId))
	if err != nil {
		return nil, err
	}
	var result PoliciesServicePoliciesDetailsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PoliciesConfigure - Update policy
func (s *PoliciesService) PoliciesConfigure(ctx context.Context, policyId string) error {
	_, err := s.client.put(ctx, fmt.Sprintf("/policies/%s", policyId), nil)
	if err != nil {
		return err
	}
	return nil
}

// PoliciesDelete - Delete policy
func (s *PoliciesService) PoliciesDelete(ctx context.Context, policyId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/policies/%s", policyId))
	if err != nil {
		return err
	}
	return nil
}
