package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// FoldersService handles communication with the Folders related endpoints of the VM API.
type FoldersService struct {
	client *Client
}

type FoldersServiceFoldersCreateRequest struct {
	Name string `json:"name,omitempty"`
}

type FoldersServiceFoldersCreateResponse struct {
	Id int `json:"id,omitempty"`
}

type FoldersServiceFoldersListResponseItem struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	TypeField   string `json:"type,omitempty"`
	DefaultTag  int    `json:"default_tag,omitempty"`
	Custom      int    `json:"custom,omitempty"`
	UnreadCount int    `json:"unread_count,omitempty"`
}

type FoldersServiceFoldersEditRequest struct {
	Name string `json:"name,omitempty"`
}

// FoldersCreate - Create folder
func (s *FoldersService) FoldersCreate(ctx context.Context, req *FoldersServiceFoldersCreateRequest) (*FoldersServiceFoldersCreateResponse, error) {
	resp, err := s.client.post(ctx, "/folders", req)
	if err != nil {
		return nil, err
	}
	var result FoldersServiceFoldersCreateResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FoldersList - List folders
func (s *FoldersService) FoldersList(ctx context.Context) ([]FoldersServiceFoldersListResponseItem, error) {
	resp, err := s.client.get(ctx, "/folders")
	if err != nil {
		return nil, err
	}
	var result []FoldersServiceFoldersListResponseItem
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// FoldersEdit - Rename folder
func (s *FoldersService) FoldersEdit(ctx context.Context, folderId string, req *FoldersServiceFoldersEditRequest) error {
	_, err := s.client.put(ctx, fmt.Sprintf("/folders/%s", folderId), req)
	if err != nil {
		return err
	}
	return nil
}

// FoldersDelete - Delete folder
func (s *FoldersService) FoldersDelete(ctx context.Context, folderId string) error {
	_, err := s.client.delete(ctx, fmt.Sprintf("/folders/%s", folderId))
	if err != nil {
		return err
	}
	return nil
}
