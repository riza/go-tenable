package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// FeedService handles communication with the feed-related endpoints of the SC API.
type FeedService struct {
	client *Client
}

// FeedStatus represents the status of a single feed type (sc, active, passive, lce).
type FeedStatus struct {
	UpdateTime    string `json:"updateTime"`
	Stale         string `json:"stale"`
	UpdateRunning string `json:"updateRunning"`
}

// FeedResponse holds the response from the feed status endpoint.
type FeedResponse struct {
	SC      FeedStatus `json:"sc"`
	Active  FeedStatus `json:"active"`
	Passive FeedStatus `json:"passive"`
	LCE     FeedStatus `json:"lce"`
}

// FeedProcessInput contains the fields for processing a feed update.
type FeedProcessInput struct {
	Filename string `json:"filename"`
}

// Get returns the status of all feeds.
func (s *FeedService) Get(ctx context.Context) (*FeedResponse, error) {
	resp, err := s.client.get(ctx, "/feed")
	if err != nil {
		return nil, fmt.Errorf("sc: get feed: %w", err)
	}

	var result FeedResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal feed response: %w", err)
	}

	return &result, nil
}

// GetType returns the status of a specific feed type.
func (s *FeedService) GetType(ctx context.Context, feedType string) (*FeedStatus, error) {
	resp, err := s.client.get(ctx, "/feed/"+feedType)
	if err != nil {
		return nil, fmt.Errorf("sc: get feed %s: %w", feedType, err)
	}

	var result FeedStatus
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal feed response: %w", err)
	}

	return &result, nil
}

// Update triggers an update for the specified feed type.
func (s *FeedService) Update(ctx context.Context, feedType string) error {
	_, err := s.client.post(ctx, "/feed/"+feedType+"/update", nil)
	if err != nil {
		return fmt.Errorf("sc: update feed %s: %w", feedType, err)
	}

	return nil
}

// Process triggers processing of an uploaded feed file for the specified feed type.
func (s *FeedService) Process(ctx context.Context, feedType string, input *FeedProcessInput) error {
	_, err := s.client.post(ctx, "/feed/"+feedType+"/process", input)
	if err != nil {
		return fmt.Errorf("sc: process feed %s: %w", feedType, err)
	}

	return nil
}
