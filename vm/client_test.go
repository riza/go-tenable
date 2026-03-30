package vm

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("https://vm.example.com", WithAPIKey("access", "secret"))

	if c.BaseURL != "https://vm.example.com" {
		t.Errorf("BaseURL = %q, want %q", c.BaseURL, "https://vm.example.com")
	}
	if c.accessKey != "access" {
		t.Errorf("accessKey = %q, want %q", c.accessKey, "access")
	}
	if c.secretKey != "secret" {
		t.Errorf("secretKey = %q, want %q", c.secretKey, "secret")
	}
	if c.UserAgent != "go-tenable/0.1" {
		t.Errorf("UserAgent = %q, want %q", c.UserAgent, "go-tenable/0.1")
	}
	if c.HTTPClient == nil {
		t.Error("HTTPClient should not be nil")
	}

	if c.ScansService == nil || c.ScanControlService == nil {
		t.Error("Not all services were initialized")
	}
}

func TestNewClientTrailingSlash(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"https://vm.example.com/", "https://vm.example.com"},
		{"https://vm.example.com", "https://vm.example.com"},
	}

	for _, tt := range tests {
		c := NewClient(tt.input)
		if c.BaseURL != tt.want {
			t.Errorf("NewClient(%q).BaseURL = %q, want %q", tt.input, c.BaseURL, tt.want)
		}
	}
}

func TestAPIErrorParsing(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(apiErrorResponse{
			Error: "Invalid credentials",
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
	if apiErr.ErrorMsg != "Invalid credentials" {
		t.Errorf("ErrorMsg = %q, want %q", apiErr.ErrorMsg, "Invalid credentials")
	}
}

func TestAPIKeyHeader(t *testing.T) {
	var gotHeader string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotHeader = r.Header.Get("x-apikeys")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer ts.Close()

	c := NewClient(ts.URL, WithAPIKey("ak", "sk"))
	_, err := c.get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "accessKey=ak;secretKey=sk"
	if gotHeader != want {
		t.Errorf("x-apikeys header = %q, want %q", gotHeader, want)
	}
}

func TestAPIKeyHeaderAbsentWithoutCredentials(t *testing.T) {
	var gotHeader string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotHeader = r.Header.Get("x-apikeys")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	_, err := c.get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotHeader != "" {
		t.Errorf("x-apikeys header = %q, want empty", gotHeader)
	}
}

func TestWithHTTPClient(t *testing.T) {
	custom := &http.Client{}
	c := NewClient("https://vm.example.com", WithHTTPClient(custom))
	if c.HTTPClient != custom {
		t.Error("expected custom HTTP client to be set")
	}
}

func TestWithInsecureSkipVerify(t *testing.T) {
	c := NewClient("https://vm.example.com", WithInsecureSkipVerify())
	if c.HTTPClient == nil || c.HTTPClient.Transport == nil {
		t.Fatal("expected custom transport")
	}
	transport, ok := c.HTTPClient.Transport.(*http.Transport)
	if !ok || transport.TLSClientConfig == nil || !transport.TLSClientConfig.InsecureSkipVerify {
		t.Error("expected InsecureSkipVerify to be true")
	}
}
