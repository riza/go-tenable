
package sc

import (
	"encoding/json"
	"fmt"
)

// FileService handles file operations.
type FileService struct {
	client *Client
}

// File represents a file resource.
type File struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Upload performs the upload action on the file.
func (s *FileService) Upload() (*File, error) {
	resp, err := s.client.post("/file/upload", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: upload file: %w", err)
	}

	var result File
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal file upload response: %w", err)
	}

	return &result, nil
}

// Clear performs the clear action on the file.
func (s *FileService) Clear() (*File, error) {
	resp, err := s.client.post("/file/clear", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: clear file: %w", err)
	}

	var result File
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal file clear response: %w", err)
	}

	return &result, nil
}

