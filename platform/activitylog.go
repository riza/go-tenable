package platform

import (
	"context"
	"encoding/json"
)

// ActivityLogService handles communication with the Activity Log related endpoints of the Platform API.
type ActivityLogService struct {
	client *Client
}

// ActivityLogEvent represents an activity log event.
type ActivityLogEvent struct {
	Id          string                 `json:"id,omitempty"`
	Action      string                 `json:"action,omitempty"`
	Actor       string                 `json:"actor,omitempty"`
	Target      string                 `json:"target,omitempty"`
	TargetType  string                 `json:"target_type,omitempty"`
	Description string                 `json:"description,omitempty"`
	CreatedAt   string                 `json:"created_time,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// ActivityLogListResponse represents the response from listing activity log events.
type ActivityLogListResponse struct {
	Events []ActivityLogEvent `json:"events,omitempty"`
	Total  int                `json:"total,omitempty"`
}

// ListActivityLog returns activity log events.
func (s *ActivityLogService) ListActivityLog(ctx context.Context) (*ActivityLogListResponse, error) {
	resp, err := s.client.get(ctx, "/audit-log/v1/events")
	if err != nil {
		return nil, err
	}

	var result ActivityLogListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
