package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWorkbenchesVulnerabilities(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"vulnerabilities": [
				{"plugin_id": 12345, "plugin_name": "Test Vuln"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.WorkbenchesService.WorkbenchesVulnerabilities(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/workbenches/vulnerabilities" {
		t.Errorf("path = %q, want %q", gotPath, "/workbenches/vulnerabilities")
	}
	if len(resp.Vulnerabilities) != 1 || resp.Vulnerabilities[0].PluginId != 12345 {
		t.Errorf("Vulnerabilities = %v", resp.Vulnerabilities)
	}
}

func TestWorkbenchesAssets(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"assets": [
				{"id": "asset1"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.WorkbenchesService.WorkbenchesAssets(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/workbenches/assets" {
		t.Errorf("path = %q, want %q", gotPath, "/workbenches/assets")
	}
	if len(resp.Assets) != 1 || resp.Assets[0].Id != "asset1" {
		t.Errorf("Assets = %v", resp.Assets)
	}
}

func TestWorkbenchesExportRequest(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"file": 123
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.WorkbenchesService.WorkbenchesExportRequest(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/workbenches/export" {
		t.Errorf("path = %q, want %q", gotPath, "/workbenches/export")
	}
	if resp.File != 123 {
		t.Errorf("File = %d", resp.File)
	}
}

func TestWorkbenchesExportStatus(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"status": "ready"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.WorkbenchesService.WorkbenchesExportStatus(context.Background(), "123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/workbenches/export/123/status" {
		t.Errorf("path = %q, want %q", gotPath, "/workbenches/export/123/status")
	}
	if resp.Status != "ready" {
		t.Errorf("Status = %s", resp.Status)
	}
}
