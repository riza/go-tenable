package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVmReportsCreate(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"uuid": "report-123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ReportsServiceVmReportsCreateRequest{
		Name: "Test Report",
	}
	resp, err := c.ReportsService.VmReportsCreate(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/reports/export" {
		t.Errorf("path = %q, want %q", gotPath, "/reports/export")
	}
	if resp.Uuid != "report-123" {
		t.Errorf("Uuid = %q, want %q", resp.Uuid, "report-123")
	}
}

func TestVmReportsStatus(t *testing.T) {
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
	resp, err := c.ReportsService.VmReportsStatus(context.Background(), "report-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/reports/export/report-123/status" {
		t.Errorf("path = %q, want %q", gotPath, "/reports/export/report-123/status")
	}
	if resp.Status != "READY" {
		t.Errorf("Status = %q, want %q", resp.Status, "READY")
	}
}

func TestVmReportsDownload(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("report data"))
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.ReportsService.VmReportsDownload(context.Background(), "report-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/reports/export/report-123/download" {
		t.Errorf("path = %q, want %q", gotPath, "/reports/export/report-123/download")
	}
}
