package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHostDetails(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"vulnerabilities": [
				{"plugin_id": 100}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScanResultsService.HostDetails(context.Background(), "scan-123", "host-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/scan-123/hosts/host-123" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/hosts/host-123")
	}
	if len(resp.Vulnerabilities) != 1 || resp.Vulnerabilities[0].PluginId != 100 {
		t.Errorf("Vulnerabilities = %v", resp.Vulnerabilities)
	}
}

func TestPluginOutput(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"output": [
				{"plugin_output": "test output"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScanResultsService.PluginOutput(context.Background(), "scan-123", "host-123", "100")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/scan-123/hosts/host-123/plugins/100" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/hosts/host-123/plugins/100")
	}
	if len(resp.Output) != 1 || resp.Output[0].PluginOutput != "test output" {
		t.Errorf("Output = %v", resp.Output)
	}
}

func TestAttachments(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("attachment data"))
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.ScanResultsService.Attachments(context.Background(), "scan-123", "attachment-1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/scan-123/attachments/attachment-1" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/attachments/attachment-1")
	}
}
