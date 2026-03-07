package vm

import (
	"context"
	"encoding/json"
)

// FileService handles communication with the File related endpoints of the VM API.
type FileService struct {
	client *Client
}

type FileServiceFileUploadResponse struct {
	Fileuploaded string `json:"fileuploaded,omitempty"`
}

// FileUpload - Upload file
func (s *FileService) FileUpload(ctx context.Context) (*FileServiceFileUploadResponse, error) {
	resp, err := s.client.post(ctx, "/file/upload", nil)
	if err != nil {
		return nil, err
	}
	var result FileServiceFileUploadResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
