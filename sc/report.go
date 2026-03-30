package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// ReportService handles communication with the report-related endpoints of the SC API.
type ReportService struct {
	client *Client
}

// Report represents a report resource from the API.
type Report struct {
	ID                 string          `json:"id"`
	ReportDefinitionID string          `json:"reportDefinitionID"`
	JobID              string          `json:"jobID"`
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	Type               string          `json:"type"`
	Status             string          `json:"status"`
	Running            string          `json:"running"`
	ErrorDetails       string          `json:"errorDetails"`
	TotalSteps         string          `json:"totalSteps"`
	CompletedSteps     string          `json:"completedSteps"`
	StartTime          string          `json:"startTime"`
	FinishTime         string          `json:"finishTime"`
	OwnerGID           string          `json:"ownerGID"`
	DisplayType        string          `json:"displayType"`
	PubSites           []IDRef         `json:"pubSites,omitempty"`
	TxLogs             json.RawMessage `json:"txLogs,omitempty"`
	CanUse             string          `json:"canUse"`
	CanManage          string          `json:"canManage"`
	Creator            *IDRef          `json:"creator,omitempty"`
	Owner              *IDRef          `json:"owner,omitempty"`
	OwnerGroup         *IDRef          `json:"ownerGroup,omitempty"`
}

// ReportListResponse holds the usable and manageable report arrays returned by the list endpoint.
type ReportListResponse struct {
	Usable     []Report `json:"usable"`
	Manageable []Report `json:"manageable"`
}

// ReportCopyInput contains the fields for copying a report.
type ReportCopyInput struct {
	Users []IDRef `json:"users,omitempty"`
}

// ReportCopyResponse holds the response from copying a report.
type ReportCopyResponse struct {
	Users []IDRef `json:"users"`
}

// ReportEmailInput contains the fields for emailing a report.
type ReportEmailInput struct {
	Email string `json:"email"`
}

// ReportEmailResponse holds the response from emailing a report.
type ReportEmailResponse struct {
	Email string `json:"email"`
}

// ReportSendInput contains the fields for sending a report to publishing sites.
type ReportSendInput struct {
	PubSites []int `json:"pubSites,omitempty"`
}

// ReportSendResponse holds the response from sending a report.
type ReportSendResponse struct {
	PubSites []int `json:"pubSites"`
}

// List returns the list of reports (usable and manageable).
func (s *ReportService) List(ctx context.Context) (*ReportListResponse, error) {
	resp, err := s.client.get(ctx, "/report")
	if err != nil {
		return nil, fmt.Errorf("sc: list reports: %w", err)
	}

	var result ReportListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report list response: %w", err)
	}

	return &result, nil
}

// Get returns the report with the given ID.
func (s *ReportService) Get(ctx context.Context, id string) (*Report, error) {
	resp, err := s.client.get(ctx, "/report/"+id)
	if err != nil {
		return nil, fmt.Errorf("sc: get report %s: %w", id, err)
	}

	var report Report
	if err := json.Unmarshal(resp.Response, &report); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report response: %w", err)
	}

	return &report, nil
}

// Delete deletes the report with the given ID.
func (s *ReportService) Delete(ctx context.Context, id string) error {
	_, err := s.client.delete(ctx, "/report/"+id)
	if err != nil {
		return fmt.Errorf("sc: delete report %s: %w", id, err)
	}

	return nil
}

// Copy copies the report with the given ID to the specified users.
func (s *ReportService) Copy(ctx context.Context, id string, input *ReportCopyInput) (*ReportCopyResponse, error) {
	resp, err := s.client.post(ctx, "/report/"+id+"/copy", input)
	if err != nil {
		return nil, fmt.Errorf("sc: copy report %s: %w", id, err)
	}

	var result ReportCopyResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report copy response: %w", err)
	}

	return &result, nil
}

// Email emails the report with the given ID to the specified address.
func (s *ReportService) Email(ctx context.Context, id string, input *ReportEmailInput) (*ReportEmailResponse, error) {
	resp, err := s.client.post(ctx, "/report/"+id+"/email", input)
	if err != nil {
		return nil, fmt.Errorf("sc: email report %s: %w", id, err)
	}

	var result ReportEmailResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report email response: %w", err)
	}

	return &result, nil
}

// Download downloads the report with the given ID. The response is the raw binary content.
func (s *ReportService) Download(ctx context.Context, id string) (json.RawMessage, error) {
	resp, err := s.client.post(ctx, "/report/"+id+"/download", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: download report %s: %w", id, err)
	}

	return resp.Response, nil
}

// Stop stops the running report with the given ID.
func (s *ReportService) Stop(ctx context.Context, id string) error {
	_, err := s.client.post(ctx, "/report/"+id+"/stop", nil)
	if err != nil {
		return fmt.Errorf("sc: stop report %s: %w", id, err)
	}

	return nil
}

// Send sends the report with the given ID to the specified publishing sites.
func (s *ReportService) Send(ctx context.Context, id string, input *ReportSendInput) (*ReportSendResponse, error) {
	resp, err := s.client.post(ctx, "/report/"+id+"/send", input)
	if err != nil {
		return nil, fmt.Errorf("sc: send report %s: %w", id, err)
	}

	var result ReportSendResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal report send response: %w", err)
	}

	return &result, nil
}
