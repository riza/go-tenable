package sc

import (
	"encoding/json"
	"fmt"
)

// BulkService handles communication with the bulk operation endpoint of the
// SC API.
type BulkService struct {
	client *Client
}

// BulkOperation represents a single operation within a bulk request.
type BulkOperation struct {
	API    string          `json:"api"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params,omitempty"`
}

// BulkInput contains the operations to execute in a single bulk request.
type BulkInput struct {
	Operations []BulkOperation `json:"operations"`
}

// BulkResponse is a type alias for the dynamic array of sub-responses
// returned by the bulk endpoint.
type BulkResponse = json.RawMessage

// Execute sends a bulk request containing multiple API operations.
func (s *BulkService) Execute(input *BulkInput) (json.RawMessage, error) {
	resp, err := s.client.post("/bulk", input)
	if err != nil {
		return nil, fmt.Errorf("sc: execute bulk: %w", err)
	}

	return resp.Response, nil
}
