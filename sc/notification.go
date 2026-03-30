package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// NotificationService handles notification operations.
type NotificationService struct {
	client *Client
}

// Notification represents a notification resource.
type Notification struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// NotificationListResponse represents the response from listing notifications.
type NotificationListResponse struct {
	Usable     []Notification `json:"usable"`
	Manageable []Notification `json:"manageable"`
}

// List returns all notifications.
func (s *NotificationService) List(ctx context.Context) (*NotificationListResponse, error) {
	resp, err := s.client.get(ctx, "/notification")
	if err != nil {
		return nil, fmt.Errorf("sc: list notifications: %w", err)
	}

	var result NotificationListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal notification list response: %w", err)
	}

	return &result, nil
}

// Get returns the notification with the given ID.
func (s *NotificationService) Get(ctx context.Context, id string) (*Notification, error) {
	resp, err := s.client.get(ctx, "/notification"+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get notification %s: %w", id, err)
	}

	var result Notification
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal notification response: %w", err)
	}

	return &result, nil
}
