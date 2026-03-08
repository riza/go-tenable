package one

import (
	"context"
	"encoding/json"
)

// ExposureViewService handles communication with the Exposure View related endpoints of the Tenable One API.
type ExposureViewService struct {
	client *Client
}

// ExposureViewCard represents an exposure view card.
type ExposureViewCard struct {
	Id          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Category    string                 `json:"category,omitempty"`
	Title       string                 `json:"title,omitempty"`
	Description string                 `json:"description,omitempty"`
	Stats       map[string]interface{} `json:"stats,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// ExposureViewCardsResponse represents the response from getting exposure view cards.
type ExposureViewCardsResponse struct {
	Cards []ExposureViewCard `json:"cards,omitempty"`
	Total int                `json:"total,omitempty"`
}

// ListCards returns all exposure view cards.
func (s *ExposureViewService) ListCards(ctx context.Context) (*ExposureViewCardsResponse, error) {
	resp, err := s.client.get(ctx, "/api/v1/t1/exposure-view/cards")
	if err != nil {
		return nil, err
	}

	var result ExposureViewCardsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetCard returns a specific exposure view card by ID.
func (s *ExposureViewService) GetCard(ctx context.Context, cardId string) (*ExposureViewCard, error) {
	resp, err := s.client.get(ctx, "/api/v1/t1/exposure-view/cards/"+cardId)
	if err != nil {
		return nil, err
	}

	var result ExposureViewCard
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
