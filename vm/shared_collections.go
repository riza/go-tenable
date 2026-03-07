package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// SharedCollectionsService handles communication with the Shared Collections related endpoints of the VM API.
type SharedCollectionsService struct {
	client *Client
}

type SharedCollectionsServiceSharedCollectionsCreateRequestAclItem struct {
	Permission string `json:"permission,omitempty"`
	TypeField  string `json:"type,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsCreateRequest struct {
	Name        string                                                          `json:"name,omitempty"`
	Description string                                                          `json:"description,omitempty"`
	Acl         []SharedCollectionsServiceSharedCollectionsCreateRequestAclItem `json:"acl,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsCreateResponse struct {
	Message         string `json:"message,omitempty"`
	RequestStatusId string `json:"request_status_id,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsListResponseItemAclItem struct {
	Permission string `json:"permission,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
	TypeField  string `json:"type,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsListResponseItem struct {
	Name        string                                                             `json:"name,omitempty"`
	Description string                                                             `json:"description,omitempty"`
	Uuid        string                                                             `json:"uuid,omitempty"`
	Acl         []SharedCollectionsServiceSharedCollectionsListResponseItemAclItem `json:"acl,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsDetailsResponseAclItem struct {
	Permission string `json:"permission,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
	TypeField  string `json:"type,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsDetailsResponse struct {
	Name        string                                                            `json:"name,omitempty"`
	Description string                                                            `json:"description,omitempty"`
	Uuid        string                                                            `json:"uuid,omitempty"`
	Acl         []SharedCollectionsServiceSharedCollectionsDetailsResponseAclItem `json:"acl,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsUpdateRequestAclItem struct {
	Permission string `json:"permission,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
	TypeField  string `json:"type,omitempty"`
	Name       string `json:"name,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsUpdateRequest struct {
	Name        string                                                          `json:"name,omitempty"`
	Description string                                                          `json:"description,omitempty"`
	Acl         []SharedCollectionsServiceSharedCollectionsUpdateRequestAclItem `json:"acl,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsUpdateResponse struct {
	Message         string `json:"message,omitempty"`
	RequestStatusId string `json:"request_status_id,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsDeleteResponse struct {
	Message         string `json:"message,omitempty"`
	RequestStatusId string `json:"request_status_id,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsDetailsByNameRequest struct {
	Name string `json:"name,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsDetailsByNameResponseAclItem struct {
	Permission string `json:"permission,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
	TypeField  string `json:"type,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsDetailsByNameResponse struct {
	Name        string                                                                  `json:"name,omitempty"`
	Description string                                                                  `json:"description,omitempty"`
	Uuid        string                                                                  `json:"uuid,omitempty"`
	Acl         []SharedCollectionsServiceSharedCollectionsDetailsByNameResponseAclItem `json:"acl,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsJobStatusResponse struct {
	ContainerUuid string `json:"container_uuid,omitempty"`
	TypeField     string `json:"type,omitempty"`
	Id            string `json:"id,omitempty"`
	Status        string `json:"status,omitempty"`
	Created       string `json:"created,omitempty"`
	Modified      string `json:"modified,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsConfigAddRequest struct {
	ScanConfigs []string `json:"scan_configs,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsConfigAddResponse struct {
	Message         string `json:"message,omitempty"`
	RequestStatusId string `json:"request_status_id,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsConfigListResponseScansItem struct {
	Control              bool   `json:"control,omitempty"`
	CreationDate         int    `json:"creation_date,omitempty"`
	Enabled              bool   `json:"enabled,omitempty"`
	Id                   int    `json:"id,omitempty"`
	LastModificationDate int    `json:"last_modification_date,omitempty"`
	Legacy               bool   `json:"legacy,omitempty"`
	Name                 string `json:"name,omitempty"`
	Owner                string `json:"owner,omitempty"`
	PolicyId             int    `json:"policy_id,omitempty"`
	Read                 bool   `json:"read,omitempty"`
	ScheduleUuid         string `json:"schedule_uuid,omitempty"`
	Shared               bool   `json:"shared,omitempty"`
	Status               string `json:"status,omitempty"`
	TemplateUuid         string `json:"template_uuid,omitempty"`
	TypeField            string `json:"type,omitempty"`
	Permissions          int    `json:"permissions,omitempty"`
	UserPermissions      int    `json:"user_permissions,omitempty"`
	Uuid                 string `json:"uuid,omitempty"`
	WizardUuid           string `json:"wizard_uuid,omitempty"`
	Progress             int    `json:"progress,omitempty"`
	Timezone             string `json:"timezone,omitempty"`
	Rrules               string `json:"rrules,omitempty"`
	Starttime            string `json:"starttime,omitempty"`
	TotalTargets         int    `json:"total_targets,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsConfigListResponsePagination struct {
	Total  int `json:"total,omitempty"`
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsConfigListResponse struct {
	Scans      []SharedCollectionsServiceSharedCollectionsConfigListResponseScansItem `json:"scans,omitempty"`
	Pagination SharedCollectionsServiceSharedCollectionsConfigListResponsePagination  `json:"pagination,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsConfigsRemoveRequest struct {
	ScanConfigs []string `json:"scan_configs,omitempty"`
}

type SharedCollectionsServiceSharedCollectionsConfigsRemoveResponse struct {
	Message         string `json:"message,omitempty"`
	RequestStatusId string `json:"request_status_id,omitempty"`
}

// SharedCollectionsCreate - Create shared collection
func (s *SharedCollectionsService) SharedCollectionsCreate(ctx context.Context, req *SharedCollectionsServiceSharedCollectionsCreateRequest) (*SharedCollectionsServiceSharedCollectionsCreateResponse, error) {
	resp, err := s.client.post(ctx, "/shared-collections", req)
	if err != nil {
		return nil, err
	}
	var result SharedCollectionsServiceSharedCollectionsCreateResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SharedCollectionsList - List shared collections
func (s *SharedCollectionsService) SharedCollectionsList(ctx context.Context) ([]SharedCollectionsServiceSharedCollectionsListResponseItem, error) {
	resp, err := s.client.get(ctx, "/shared-collections")
	if err != nil {
		return nil, err
	}
	var result []SharedCollectionsServiceSharedCollectionsListResponseItem
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SharedCollectionsDetails - Get shared collection details
func (s *SharedCollectionsService) SharedCollectionsDetails(ctx context.Context, sharedCollectionId string) (*SharedCollectionsServiceSharedCollectionsDetailsResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/shared-collections/%s", sharedCollectionId))
	if err != nil {
		return nil, err
	}
	var result SharedCollectionsServiceSharedCollectionsDetailsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SharedCollectionsUpdate - Update shared collection
func (s *SharedCollectionsService) SharedCollectionsUpdate(ctx context.Context, sharedCollectionId string, req *SharedCollectionsServiceSharedCollectionsUpdateRequest) (*SharedCollectionsServiceSharedCollectionsUpdateResponse, error) {
	resp, err := s.client.put(ctx, fmt.Sprintf("/shared-collections/%s", sharedCollectionId), req)
	if err != nil {
		return nil, err
	}
	var result SharedCollectionsServiceSharedCollectionsUpdateResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SharedCollectionsDelete - Delete shared collection
func (s *SharedCollectionsService) SharedCollectionsDelete(ctx context.Context, sharedCollectionId string) (*SharedCollectionsServiceSharedCollectionsDeleteResponse, error) {
	resp, err := s.client.delete(ctx, fmt.Sprintf("/shared-collections/%s", sharedCollectionId))
	if err != nil {
		return nil, err
	}
	var result SharedCollectionsServiceSharedCollectionsDeleteResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SharedCollectionsDetailsByName - Get shared collection details by name
func (s *SharedCollectionsService) SharedCollectionsDetailsByName(ctx context.Context, req *SharedCollectionsServiceSharedCollectionsDetailsByNameRequest) (*SharedCollectionsServiceSharedCollectionsDetailsByNameResponse, error) {
	resp, err := s.client.post(ctx, "/shared-collections/_byName", req)
	if err != nil {
		return nil, err
	}
	var result SharedCollectionsServiceSharedCollectionsDetailsByNameResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SharedCollectionsJobStatus - Get job status
func (s *SharedCollectionsService) SharedCollectionsJobStatus(ctx context.Context, requestStatusId string) (*SharedCollectionsServiceSharedCollectionsJobStatusResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/shared-collections/request-status/%s", requestStatusId))
	if err != nil {
		return nil, err
	}
	var result SharedCollectionsServiceSharedCollectionsJobStatusResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SharedCollectionsConfigAdd - Add scan configs to shared collection
func (s *SharedCollectionsService) SharedCollectionsConfigAdd(ctx context.Context, sharedCollectionId string, req *SharedCollectionsServiceSharedCollectionsConfigAddRequest) (*SharedCollectionsServiceSharedCollectionsConfigAddResponse, error) {
	resp, err := s.client.post(ctx, fmt.Sprintf("/shared-collections/%s/scan-configs", sharedCollectionId), req)
	if err != nil {
		return nil, err
	}
	var result SharedCollectionsServiceSharedCollectionsConfigAddResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SharedCollectionsConfigList - List scan configs in shared collection
func (s *SharedCollectionsService) SharedCollectionsConfigList(ctx context.Context, sharedCollectionId string) (*SharedCollectionsServiceSharedCollectionsConfigListResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/shared-collections/%s/scan-configs", sharedCollectionId))
	if err != nil {
		return nil, err
	}
	var result SharedCollectionsServiceSharedCollectionsConfigListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SharedCollectionsConfigsRemove - Remove scan configs from shared collection
func (s *SharedCollectionsService) SharedCollectionsConfigsRemove(ctx context.Context, sharedCollectionId string, req *SharedCollectionsServiceSharedCollectionsConfigsRemoveRequest) (*SharedCollectionsServiceSharedCollectionsConfigsRemoveResponse, error) {
	resp, err := s.client.delete(ctx, fmt.Sprintf("/shared-collections/%s/scan-configs", sharedCollectionId))
	if err != nil {
		return nil, err
	}
	var result SharedCollectionsServiceSharedCollectionsConfigsRemoveResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
