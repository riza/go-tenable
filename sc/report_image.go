package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// ReportImageService handles reportImage operations.
type ReportImageService struct {
	client *Client
}

// ReportImage represents a reportImage resource.
type ReportImage struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ReportImageListResponse represents the response from listing reportImages.
type ReportImageListResponse struct {
	Usable     []ReportImage `json:"usable"`
	Manageable []ReportImage `json:"manageable"`
}

// ReportImageCreateInput represents the request body for creating a reportImage.
type ReportImageCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// List returns all reportImages.
func (s *ReportImageService) List(ctx context.Context) (*ReportImageListResponse, error) {
	resp, err := s.client.get(ctx, "/report/image")
	if err != nil {
		return nil, fmt.Errorf("sc: list reportImages: %w", err)
	}

	var result ReportImageListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal reportImage list response: %w", err)
	}

	return &result, nil
}

// Create creates a new reportImage.
func (s *ReportImageService) Create(ctx context.Context, input *ReportImageCreateInput) (*ReportImage, error) {
	resp, err := s.client.post(ctx, "/report/image", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create reportImage: %w", err)
	}

	var result ReportImage
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal reportImage response: %w", err)
	}

	return &result, nil
}
