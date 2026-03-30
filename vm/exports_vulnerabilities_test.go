package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExportsVulnsRequestExport(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"export_uuid": "vuln-export-123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ExportsVulnerabilitiesServiceExportsVulnsRequestExportRequest{
		NumAssets: 100,
	}
	resp, err := c.ExportsVulnerabilitiesService.ExportsVulnsRequestExport(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/vulns/export" {
		t.Errorf("path = %q, want %q", gotPath, "/vulns/export")
	}
	if resp.ExportUuid != "vuln-export-123" {
		t.Errorf("ExportUuid = %q, want %q", resp.ExportUuid, "vuln-export-123")
	}
}

func TestExportsVulnsExportStatus(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"status": "FINISHED",
			"uuid": "vuln-export-123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportsVulnerabilitiesService.ExportsVulnsExportStatus(context.Background(), "vuln-export-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/vulns/export/vuln-export-123/status" {
		t.Errorf("path = %q, want %q", gotPath, "/vulns/export/vuln-export-123/status")
	}
	if resp.Status != "FINISHED" {
		t.Errorf("Status = %q, want %q", resp.Status, "FINISHED")
	}
}

func TestExportsVulnsExportStatusRecent(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"exports": [
				{"uuid": "vuln-export-123", "status": "FINISHED"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportsVulnerabilitiesService.ExportsVulnsExportStatusRecent(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/vulns/export/status" {
		t.Errorf("path = %q, want %q", gotPath, "/vulns/export/status")
	}
	if len(resp.Exports) != 1 || resp.Exports[0].Uuid != "vuln-export-123" {
		t.Errorf("Exports = %v", resp.Exports)
	}
}

func TestExportsVulnsDownloadChunk(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("vuln chunk data"))
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.ExportsVulnerabilitiesService.ExportsVulnsDownloadChunk(context.Background(), "vuln-export-123", "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/vulns/export/vuln-export-123/chunks/1" {
		t.Errorf("path = %q, want %q", gotPath, "/vulns/export/vuln-export-123/chunks/1")
	}
}

func TestExportsVulnsExportCancel(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"status": "CANCELLED"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportsVulnerabilitiesService.ExportsVulnsExportCancel(context.Background(), "vuln-export-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/vulns/export/vuln-export-123/cancel" {
		t.Errorf("path = %q, want %q", gotPath, "/vulns/export/vuln-export-123/cancel")
	}
	if resp.Status != "CANCELLED" {
		t.Errorf("Status = %q, want %q", resp.Status, "CANCELLED")
	}
}
