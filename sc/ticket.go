
package sc

import (
	"encoding/json"
	"fmt"
)

// TicketService handles ticket operations.
type TicketService struct {
	client *Client
}

// Ticket represents a ticket resource.
type Ticket struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// TicketListResponse represents the response from listing tickets.
type TicketListResponse struct {
	Usable     []Ticket `json:"usable"`
	Manageable []Ticket `json:"manageable"`
}

// TicketCreateInput represents the request body for creating a ticket.
type TicketCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// TicketUpdateInput represents the request body for updating a ticket.
type TicketUpdateInput = TicketCreateInput

// List returns all tickets.
func (s *TicketService) List() (*TicketListResponse, error) {
	resp, err := s.client.get("/ticket")
	if err != nil {
		return nil, fmt.Errorf("sc: list tickets: %w", err)
	}

	var result TicketListResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal ticket list response: %w", err)
	}

	return &result, nil
}

// Create creates a new ticket.
func (s *TicketService) Create(input *TicketCreateInput) (*Ticket, error) {
	resp, err := s.client.post("/ticket", input)
	if err != nil {
		return nil, fmt.Errorf("sc: create ticket: %w", err)
	}

	var result Ticket
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal ticket response: %w", err)
	}

	return &result, nil
}

// Get returns the ticket with the given ID.
func (s *TicketService) Get(id string) (*Ticket, error) {
	resp, err := s.client.get("/ticket" + "/" + id)
	if err != nil {
		return nil, fmt.Errorf("sc: get ticket %s: %w", id, err)
	}

	var result Ticket
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal ticket response: %w", err)
	}

	return &result, nil
}

// Update updates the ticket with the given ID.
func (s *TicketService) Update(id string, input *TicketUpdateInput) (*Ticket, error) {
	resp, err := s.client.patch("/ticket" + "/" + id, input)
	if err != nil {
		return nil, fmt.Errorf("sc: update ticket %s: %w", id, err)
	}

	var result Ticket
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal ticket response: %w", err)
	}

	return &result, nil
}

