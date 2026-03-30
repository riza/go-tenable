package sc

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("https://sc.example.com", WithAPIKey("myaccess", "mysecret"))

	if c.BaseURL != "https://sc.example.com" {
		t.Errorf("BaseURL = %q, want %q", c.BaseURL, "https://sc.example.com")
	}
	if c.accessKey != "myaccess" {
		t.Errorf("accessKey = %q, want %q", c.accessKey, "myaccess")
	}
	if c.secretKey != "mysecret" {
		t.Errorf("secretKey = %q, want %q", c.secretKey, "mysecret")
	}
	if c.UserAgent != "go-tenable/0.1" {
		t.Errorf("UserAgent = %q, want %q", c.UserAgent, "go-tenable/0.1")
	}
	if c.HTTPClient == nil {
		t.Error("HTTPClient should not be nil")
	}
}

func TestNewClientTrailingSlash(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"https://sc.example.com/", "https://sc.example.com"},
		{"https://sc.example.com///", "https://sc.example.com"},
		{"https://sc.example.com", "https://sc.example.com"},
	}

	for _, tt := range tests {
		c := NewClient(tt.input)
		if c.BaseURL != tt.want {
			t.Errorf("NewClient(%q).BaseURL = %q, want %q", tt.input, c.BaseURL, tt.want)
		}
	}
}

func TestAPIKeyHeader(t *testing.T) {
	var gotHeader string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotHeader = r.Header.Get("x-apikey")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type:     "regular",
			Response: json.RawMessage(`{}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL, WithAPIKey("ak", "sk"))
	_, err := c.get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "accesskey=ak; secretkey=sk;"
	if gotHeader != want {
		t.Errorf("x-apikey header = %q, want %q", gotHeader, want)
	}
}

func TestAPIKeyHeaderAbsentWithoutCredentials(t *testing.T) {
	var gotHeader string
	var hasHeader bool
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotHeader = r.Header.Get("x-apikey")
		_, hasHeader = r.Header["X-Apikey"]
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type:     "regular",
			Response: json.RawMessage(`{}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	_, err := c.get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if hasHeader {
		t.Errorf("x-apikey header should not be set when no API key configured, got %q", gotHeader)
	}
}

func TestAPIErrorParsing(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(apiResponse{
			Type:      "error",
			ErrorCode: 143,
			ErrorMsg:  "Invalid API key",
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	_, err := c.get(context.Background(), "/test")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != http.StatusForbidden {
		t.Errorf("StatusCode = %d, want %d", apiErr.StatusCode, http.StatusForbidden)
	}
	if apiErr.ErrorCode != 143 {
		t.Errorf("ErrorCode = %d, want %d", apiErr.ErrorCode, 143)
	}
	if apiErr.ErrorMsg != "Invalid API key" {
		t.Errorf("ErrorMsg = %q, want %q", apiErr.ErrorMsg, "Invalid API key")
	}

	// Verify the Error() string format.
	wantMsg := "sc: API error (HTTP 403, code 143): Invalid API key"
	if apiErr.Error() != wantMsg {
		t.Errorf("Error() = %q, want %q", apiErr.Error(), wantMsg)
	}
}

func TestAPIErrorParsingWithErrorCodeOnSuccess(t *testing.T) {
	// Even if HTTP status is 200, a non-zero error_code should produce an APIError.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(apiResponse{
			Type:      "error",
			ErrorCode: 42,
			ErrorMsg:  "Something went wrong",
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	_, err := c.get(context.Background(), "/test")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", apiErr.StatusCode, http.StatusOK)
	}
	if apiErr.ErrorCode != 42 {
		t.Errorf("ErrorCode = %d, want %d", apiErr.ErrorCode, 42)
	}
}

func TestDoRequestSetsContentType(t *testing.T) {
	var gotContentType string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotContentType = r.Header.Get("Content-Type")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type:     "regular",
			Response: json.RawMessage(`{}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	_, err := c.get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotContentType != "application/json" {
		t.Errorf("Content-Type = %q, want %q", gotContentType, "application/json")
	}
}

func TestDoRequestSetsUserAgent(t *testing.T) {
	var gotUA string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUA = r.Header.Get("User-Agent")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type:     "regular",
			Response: json.RawMessage(`{}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	_, err := c.get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotUA != "go-tenable/0.1" {
		t.Errorf("User-Agent = %q, want %q", gotUA, "go-tenable/0.1")
	}
}

func TestHTTPMethods(t *testing.T) {
	tests := []struct {
		name       string
		callMethod func(c *Client) (*apiResponse, error)
		wantMethod string
		wantPath   string
	}{
		{
			name:       "GET",
			callMethod: func(c *Client) (*apiResponse, error) { return c.get(context.Background(), "/scan") },
			wantMethod: http.MethodGet,
			wantPath:   "/rest/scan",
		},
		{
			name:       "POST",
			callMethod: func(c *Client) (*apiResponse, error) { return c.post(context.Background(), "/scan", nil) },
			wantMethod: http.MethodPost,
			wantPath:   "/rest/scan",
		},
		{
			name:       "PATCH",
			callMethod: func(c *Client) (*apiResponse, error) { return c.patch(context.Background(), "/scan/1", nil) },
			wantMethod: http.MethodPatch,
			wantPath:   "/rest/scan/1",
		},
		{
			name:       "DELETE",
			callMethod: func(c *Client) (*apiResponse, error) { return c.delete(context.Background(), "/scan/1") },
			wantMethod: http.MethodDelete,
			wantPath:   "/rest/scan/1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotMethod, gotPath string
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(apiResponse{
					Type:     "regular",
					Response: json.RawMessage(`{}`),
				})
			}))
			defer ts.Close()

			c := NewClient(ts.URL)
			_, err := tt.callMethod(c)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if gotMethod != tt.wantMethod {
				t.Errorf("method = %q, want %q", gotMethod, tt.wantMethod)
			}
			if gotPath != tt.wantPath {
				t.Errorf("path = %q, want %q", gotPath, tt.wantPath)
			}
		})
	}
}

func TestPostBodyMarshaling(t *testing.T) {
	type requestBody struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Count       int    `json:"count"`
	}

	var gotBody requestBody
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		if err := json.Unmarshal(data, &gotBody); err != nil {
			t.Fatalf("failed to unmarshal request body: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type:     "regular",
			Response: json.RawMessage(`{}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	input := requestBody{
		Name:        "test-scan",
		Description: "A test scan",
		Count:       42,
	}
	_, err := c.post(context.Background(), "/scan", input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotBody.Name != input.Name {
		t.Errorf("Name = %q, want %q", gotBody.Name, input.Name)
	}
	if gotBody.Description != input.Description {
		t.Errorf("Description = %q, want %q", gotBody.Description, input.Description)
	}
	if gotBody.Count != input.Count {
		t.Errorf("Count = %d, want %d", gotBody.Count, input.Count)
	}
}

func TestSuccessfulResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type:      "regular",
			Response:  json.RawMessage(`{"id":"123","name":"My Scan"}`),
			Timestamp: 1700000000,
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.get(context.Background(), "/scan/123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Type != "regular" {
		t.Errorf("Type = %q, want %q", resp.Type, "regular")
	}
	if resp.Timestamp != 1700000000 {
		t.Errorf("Timestamp = %d, want %d", resp.Timestamp, 1700000000)
	}

	// Verify the Response field contains the raw JSON.
	var result struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		t.Fatalf("failed to unmarshal Response: %v", err)
	}
	if result.ID != "123" {
		t.Errorf("Response.ID = %q, want %q", result.ID, "123")
	}
	if result.Name != "My Scan" {
		t.Errorf("Response.Name = %q, want %q", result.Name, "My Scan")
	}
}

func TestWithHTTPClient(t *testing.T) {
	custom := &http.Client{}
	c := NewClient("https://sc.example.com", WithHTTPClient(custom))
	if c.HTTPClient != custom {
		t.Error("expected custom HTTP client to be set")
	}
}

func TestUnmarshalError(t *testing.T) {
	// Server returns invalid JSON.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("not valid json"))
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	_, err := c.get(context.Background(), "/test")
	if err == nil {
		t.Fatal("expected error for invalid JSON response, got nil")
	}
	// Should not be an APIError, should be a generic unmarshal error.
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		t.Errorf("expected non-APIError, got *APIError: %v", apiErr)
	}
}

func TestHTTPStatusErrorWithZeroErrorCode(t *testing.T) {
	// Server returns HTTP 500 but with error_code 0.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apiResponse{
			Type:      "error",
			ErrorCode: 0,
			ErrorMsg:  "Internal Server Error",
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	_, err := c.get(context.Background(), "/test")
	if err == nil {
		t.Fatal("expected error for HTTP 500, got nil")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", apiErr.StatusCode, http.StatusInternalServerError)
	}
}

func TestNilBodyProducesNoRequestBody(t *testing.T) {
	var gotContentLength int64
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotContentLength = r.ContentLength
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type:     "regular",
			Response: json.RawMessage(`{}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	_, err := c.get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// GET with nil body should have no content length (0 or -1).
	if gotContentLength > 0 {
		t.Errorf("ContentLength = %d, want 0 or -1 for nil body", gotContentLength)
	}
}
