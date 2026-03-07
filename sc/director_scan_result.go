
package sc

import (
	"encoding/json"
	"fmt"
)

// DirectorScanResultService handles directorScanResult operations.
type DirectorScanResultService struct {
	client *Client
}

// DirectorScanResult represents a directorScanResult resource.
type DirectorScanResult struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DirectorScanResultListResponse represents the response from listing directorScanResults.
type DirectorScanResultListResponse struct {
	Usable     []DirectorScanResult `json:"usable"`
	Manageable []DirectorScanResult `json:"manageable"`
}

// List returns all directorScanResults.
func (s *DirectorScanResultService) List() (*DirectorScanResultListResponse, error) {
	resp, err := s.client.get("/mgmt/scanResult")
	if err != nil {
		return nil, fmt.Errorf("sc: list directorScanResults: %w", err)
	}

	var result DirectorScanResultListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanResult list response: %w", err)
	}

	return &result, nil
}

// Get returns the directorScanResult with the given ID.
func (s *DirectorScanResultService) Get(id string) (*DirectorScanResult, error) {
	resp, err := s.client.get("/mgmt/scanResult" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get directorScanResult %s: %w", id, err)
	}

	var result DirectorScanResult
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanResult response: %w", err)
	}

	return &result, nil
}

// Email performs the email action on the directorScanResult with the given ID.
func (s *DirectorScanResultService) Email(id string) (*DirectorScanResult, error) {
	resp, err := s.client.post("/mgmt/scanResult" + "/" + id + "/email", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: email directorScanResult %s: %w", id, err)
	}

	var result DirectorScanResult
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanResult email response: %w", err)
	}

	return &result, nil
}

// Pause performs the pause action on the directorScanResult with the given ID.
func (s *DirectorScanResultService) Pause(id string) (*DirectorScanResult, error) {
	resp, err := s.client.post("/mgmt/scanResult" + "/" + id + "/pause", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: pause directorScanResult %s: %w", id, err)
	}

	var result DirectorScanResult
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanResult pause response: %w", err)
	}

	return &result, nil
}

// Resume performs the resume action on the directorScanResult with the given ID.
func (s *DirectorScanResultService) Resume(id string) (*DirectorScanResult, error) {
	resp, err := s.client.post("/mgmt/scanResult" + "/" + id + "/resume", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: resume directorScanResult %s: %w", id, err)
	}

	var result DirectorScanResult
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanResult resume response: %w", err)
	}

	return &result, nil
}

// Download performs the download action on the directorScanResult with the given ID.
func (s *DirectorScanResultService) Download(id string) (*DirectorScanResult, error) {
	resp, err := s.client.post("/mgmt/scanResult" + "/" + id + "/download", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: download directorScanResult %s: %w", id, err)
	}

	var result DirectorScanResult
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorScanResult download response: %w", err)
	}

	return &result, nil
}

