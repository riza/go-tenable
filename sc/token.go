package sc

import (
	"context"
	"encoding/json"
	"fmt"
)

// TokenService handles token (session) operations for Tenable Security Center.
type TokenService struct {
	client *Client
}

// TokenCreateInput holds the parameters for creating a session token (login).
type TokenCreateInput struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	ReleaseSession string `json:"releaseSession,omitempty"`
}

// TokenResponse holds the response from creating a session token.
type TokenResponse struct {
	Token int `json:"token"`
}

// Create logs in to Tenable Security Center and returns a session token.
// POST /token
func (s *TokenService) Create(ctx context.Context, input *TokenCreateInput) (*TokenResponse, error) {
	resp, err := s.client.post(ctx, "/token", input)
	if err != nil {
		return nil, fmt.Errorf("sc: token create: %w", err)
	}

	var token TokenResponse
	if err := json.Unmarshal(resp.Response, &token); err != nil {
		return nil, fmt.Errorf("sc: token create unmarshal: %w", err)
	}

	return &token, nil
}

// Delete destroys the current session token (logout).
// DELETE /token
func (s *TokenService) Delete(ctx context.Context) error {
	_, err := s.client.delete(ctx, "/token")
	if err != nil {
		return fmt.Errorf("sc: token delete: %w", err)
	}

	return nil
}
