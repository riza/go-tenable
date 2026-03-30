package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// AlertService handles communication with the alert-related endpoints of the SC API.
type AlertService struct {
	client *Client
}

// Alert represents an alert resource from the API.
type Alert struct {
	ID                       string          `json:"id"`
	Name                     string          `json:"name"`
	Description              string          `json:"description"`
	TriggerName              string          `json:"triggerName"`
	TriggerOperator          string          `json:"triggerOperator"`
	TriggerValue             string          `json:"triggerValue"`
	ModifiedTime             string          `json:"modifiedTime"`
	CreatedTime              string          `json:"createdTime"`
	LastTriggered            string          `json:"lastTriggered"`
	LastEvaluated            string          `json:"lastEvaluated"`
	ExecuteOnEveryTrigger    string          `json:"executeOnEveryTrigger"`
	DidTriggerLastEvaluation string          `json:"didTriggerLastEvaluation"`
	Status                   string          `json:"status"`
	Action                   json.RawMessage `json:"action,omitempty"`
	Schedule                 *AlertSchedule  `json:"schedule,omitempty"`
	Query                    *AlertQueryRef  `json:"query,omitempty"`
	CanUse                   string          `json:"canUse"`
	CanManage                string          `json:"canManage"`
	Owner                    *IDRef          `json:"owner,omitempty"`
	OwnerGroup               *IDRef          `json:"ownerGroup,omitempty"`
}

// AlertSchedule represents the schedule configuration for an alert.
type AlertSchedule struct {
	Type       string `json:"type"`
	Start      string `json:"start"`
	RepeatRule string `json:"repeatRule"`
}

// AlertQueryRef is a query reference with name and description.
type AlertQueryRef struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// AlertListResponse holds the usable and manageable alert arrays returned by the list endpoint.
type AlertListResponse struct {
	Usable     []Alert `json:"usable"`
	Manageable []Alert `json:"manageable"`
}

// AlertCreateInput contains the fields for creating a new alert.
type AlertCreateInput struct {
	Name                  string          `json:"name"`
	Description           string          `json:"description,omitempty"`
	Query                 *IDRef          `json:"query,omitempty"`
	TriggerName           string          `json:"triggerName,omitempty"`
	TriggerOperator       string          `json:"triggerOperator,omitempty"`
	TriggerValue          string          `json:"triggerValue,omitempty"`
	ExecuteOnEveryTrigger string          `json:"executeOnEveryTrigger,omitempty"`
	Schedule              *AlertSchedule  `json:"schedule,omitempty"`
	Action                json.RawMessage `json:"action,omitempty"`
}

// AlertUpdateInput contains the fields for updating an existing alert.
type AlertUpdateInput = AlertCreateInput

// List returns the list of alerts (usable and manageable).
func (s *AlertService) List(ctx context.Context) (*AlertListResponse, error) {
	resp, err := s.client.get(ctx, "/alert")
	if err != nil {
		return nil, fmt.Errorf("sc: list alerts: %w", err)
	}

	var result AlertListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal alert list response: %w", err)
	}

	return &result, nil
}

// Create creates a new alert with the given input.
func (s *AlertService) Create(ctx context.Context, input *AlertCreateInput) (*Alert, error) {
	resp, err := s.client.post(ctx, "/alert", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create alert: %w", err)
	}

	var alert Alert
	if err := json.Unmarshal(resp.Response, &alert); err != nil {
		return nil, fmt.Errorf("sc: unmarshal alert response: %w", err)
	}

	return &alert, nil
}

// Get returns the alert with the given ID.
func (s *AlertService) Get(ctx context.Context, id string) (*Alert, error) {
	resp, err := s.client.get(ctx, "/alert/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get alert %s: %w", id, err)
	}

	var alert Alert
	if err := json.Unmarshal(resp.Response, &alert); err != nil {
		return nil, fmt.Errorf("sc: unmarshal alert response: %w", err)
	}

	return &alert, nil
}

// Update updates an existing alert with the given input.
func (s *AlertService) Update(ctx context.Context, id string, input *AlertUpdateInput) (*Alert, error) {
	resp, err := s.client.patch(ctx, "/alert/"+id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update alert %s: %w", id, err)
	}

	var alert Alert
	if err := json.Unmarshal(resp.Response, &alert); err != nil {
		return nil, fmt.Errorf("sc: unmarshal alert response: %w", err)
	}

	return &alert, nil
}

// Delete deletes the alert with the given ID.
func (s *AlertService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/alert/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete alert %s: %w", id, err)
	}

	return nil
}

// Execute triggers execution of the alert with the given ID.
func (s *AlertService) Execute(ctx context.Context, id string) (*Alert, error) {
	resp, err := s.client.post(ctx, "/alert/"+id+"/execute", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: execute alert %s: %w", id, err)
	}

	var alert Alert
	if err := json.Unmarshal(resp.Response, &alert); err != nil {
		return nil, fmt.Errorf("sc: unmarshal alert execute response: %w", err)
	}

	return &alert, nil
}
