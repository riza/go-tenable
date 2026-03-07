package sc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// AnalysisService handles communication with the analysis-related endpoints of the SC API.
type AnalysisService struct {
	client *Client
}

// AnalysisInput contains the fields for running an analysis query.
type AnalysisInput struct {
	Type        string          `json:"type,omitempty"`
	Query       json.RawMessage `json:"query,omitempty"`
	SourceType  string          `json:"sourceType,omitempty"`
	SortField   string          `json:"sortField,omitempty"`
	SortDir     string          `json:"sortDir,omitempty"`
	Columns     json.RawMessage `json:"columns,omitempty"`
	StartOffset int             `json:"startOffset,omitempty"`
	EndOffset   int             `json:"endOffset,omitempty"`
}

// AnalysisResponse holds the response from running an analysis query.
type AnalysisResponse struct {
	TotalRecords    string            `json:"totalRecords"`
	ReturnedRecords int               `json:"returnedRecords"`
	Results         []json.RawMessage `json:"results"`
}

// Run executes an analysis query and returns the results.
func (s *AnalysisService) Run(input *AnalysisInput) (*AnalysisResponse, error) {
	resp, err := s.client.post("/analysis", input)
	if err != nil {
		return nil, fmt.Errorf("sc: run analysis: %w", err)
	}

	var result AnalysisResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal analysis response: %w", err)
	}

	return &result, nil
}

// Download executes an analysis query and returns the binary response body.
// The caller is responsible for closing the returned io.ReadCloser.
func (s *AnalysisService) Download(input *AnalysisInput) (io.ReadCloser, error) {
	req, err := s.client.newRequest(http.MethodPost, "/analysis/download", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create analysis download request: %w", err)
	}

	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sc: analysis download request failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		resp.Body.Close()
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			ErrorMsg:   fmt.Sprintf("HTTP %d", resp.StatusCode),
		}
	}

	return resp.Body, nil
}
