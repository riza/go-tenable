package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// DirectorSystemService handles directorSystem operations.
type DirectorSystemService struct {
	client *Client
}

// DirectorSystem represents a directorSystem resource.
type DirectorSystem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// LogFiles performs the logFiles action on the directorSystem.
func (s *DirectorSystemService) LogFiles(ctx context.Context) (*DirectorSystem, error) {
	resp, err := s.client.get(ctx, "/mgmt/system/logFiles")
	if err != nil {
		return nil, fmt.Errorf("sc: logFiles directorSystem: %w", err)
	}

	var result DirectorSystem
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorSystem logFiles response: %w", err)
	}

	return &result, nil
}

// Logs performs the logs action on the directorSystem.
func (s *DirectorSystemService) Logs(ctx context.Context) (*DirectorSystem, error) {
	resp, err := s.client.post(ctx, "/mgmt/system/logs", nil)
	if err != nil {
		return nil, fmt.Errorf("sc: logs directorSystem: %w", err)
	}

	var result DirectorSystem
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal directorSystem logs response: %w", err)
	}

	return &result, nil
}
