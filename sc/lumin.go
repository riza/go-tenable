
package sc

import (
	"encoding/json"
	"fmt"
)

// LuminService handles lumin operations.
type LuminService struct {
	client *Client
}

// Lumin represents a lumin resource.
type Lumin struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Repositories performs the repositories action on the lumin.
func (s *LuminService) Repositories() (*Lumin, error) {
	resp, err := s.client.post("/lumin/repositories", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: repositories lumin: %w", err)
	}

	var result Lumin
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lumin repositories response: %w", err)
	}

	return &result, nil
}

// Assets performs the assets action on the lumin.
func (s *LuminService) Assets() (*Lumin, error) {
	resp, err := s.client.post("/lumin/assets", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: assets lumin: %w", err)
	}

	var result Lumin
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lumin assets response: %w", err)
	}

	return &result, nil
}

// Schedule performs the schedule action on the lumin.
func (s *LuminService) Schedule() (*Lumin, error) {
	resp, err := s.client.get("/lumin/schedule")
	if err != nil {
		return nil, fmt.Errorf("sc: schedule lumin: %w", err)
	}

	var result Lumin
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lumin schedule response: %w", err)
	}

	return &result, nil
}

// Metrics performs the metrics action on the lumin.
func (s *LuminService) Metrics() (*Lumin, error) {
	resp, err := s.client.get("/lumin/metrics")
	if err != nil {
		return nil, fmt.Errorf("sc: metrics lumin: %w", err)
	}

	var result Lumin
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lumin metrics response: %w", err)
	}

	return &result, nil
}

// Test performs the test action on the lumin.
func (s *LuminService) Test() (*Lumin, error) {
	resp, err := s.client.get("/lumin/test")
	if err != nil {
		return nil, fmt.Errorf("sc: test lumin: %w", err)
	}

	var result Lumin
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal lumin test response: %w", err)
	}

	return &result, nil
}

