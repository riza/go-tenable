package sc

import "fmt"

// APIError represents an error returned by the SC API.
type APIError struct {
	StatusCode int
	ErrorCode  int
	ErrorMsg   string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("sc: API error (HTTP %d, code %d): %s", e.StatusCode, e.ErrorCode, e.ErrorMsg)
}

// IDRef is a common reference type for nested objects like {"id": "123"}.
type IDRef struct {
	ID   string `json:"id,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

// PaginatedResponse holds pagination metadata for list endpoints that support it.
type PaginatedResponse struct {
	TotalRecords    int `json:"totalRecords"`
	ReturnedRecords int `json:"returnedRecords"`
	StartOffset     int `json:"startOffset"`
	EndOffset       int `json:"endOffset"`
}
