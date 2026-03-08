package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExportRequest(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"file": "export_file.pdf",
			"temp_token": "token123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ScanExportsServiceExportRequestRequest{
		Format: "pdf",
	}
	resp, err := c.ScanExportsService.ExportRequest(context.Background(), "scan-123", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans/scan-123/export" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/export")
	}
	if resp.File != "export_file.pdf" {
		t.Errorf("File = %q, want %q", resp.File, "export_file.pdf")
	}
}

func TestExportStatus(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"status": "READY"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScanExportsService.ExportStatus(context.Background(), "scan-123", "file-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/scan-123/export/file-123/status" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/export/file-123/status")
	}
	if resp.Status != "READY" {
		t.Errorf("Status = %q, want %q", resp.Status, "READY")
	}
}

func TestExportDownload(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("scan export data"))
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.ScanExportsService.ExportDownload(context.Background(), "scan-123", "file-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/scan-123/export/file-123/download" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/export/file-123/download")
	}
}
