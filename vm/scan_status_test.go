package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLatestStatus(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"status": "running"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScanStatusService.GetLatestStatus(context.Background(), "scan-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/scan-123/latest-status" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/latest-status")
	}
	if resp.Status != "running" {
		t.Errorf("Status = %q, want %q", resp.Status, "running")
	}
}

func TestReadStatus(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ScanStatusServiceReadStatusRequest{
		Read: true,
	}
	err := c.ScanStatusService.ReadStatus(context.Background(), "scan-123", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPut {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPut)
	}
	if gotPath != "/scans/scan-123/status" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/status")
	}
}

func TestVmScansProgressGet(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"progress": 50
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScanStatusService.VmScansProgressGet(context.Background(), "scan-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/scan-123/progress" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/progress")
	}
	if resp.Progress != 50 {
		t.Errorf("Progress = %d, want %d", resp.Progress, 50)
	}
}
