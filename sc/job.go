
package sc

import (
	"encoding/json"
	"fmt"
)

// JobService handles job operations.
type JobService struct {
	client *Client
}

// Job represents a job resource.
type Job struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// JobListResponse represents the response from listing jobs.
type JobListResponse struct {
	Usable     []Job `json:"usable"`
	Manageable []Job `json:"manageable"`
}

// List returns all jobs.
func (s *JobService) List() (*JobListResponse, error) {
	resp, err := s.client.get("/job")
	if err != nil {
		return nil, fmt.Errorf("sc: list jobs: %w", err)
	}

	var result JobListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal job list response: %w", err)
	}

	return &result, nil
}

// Get returns the job with the given ID.
func (s *JobService) Get(id string) (*Job, error) {
	resp, err := s.client.get("/job" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get job %s: %w", id, err)
	}

	var result Job
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal job response: %w", err)
	}

	return &result, nil
}

// Kill performs the kill action on the job with the given ID.
func (s *JobService) Kill(id string) (*Job, error) {
	resp, err := s.client.post("/job" + "/" + id + "/kill", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: kill job %s: %w", id, err)
	}

	var result Job
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal job kill response: %w", err)
	}

	return &result, nil
}

