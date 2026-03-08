package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScanControlLaunch(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"scan_uuid": "test-uuid-123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ScanControlServiceLaunchRequest{
		AltTargets: []string{"192.168.1.1"},
	}
	resp, err := c.ScanControlService.Launch(context.Background(), "1", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans/1/launch" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/1/launch")
	}
	if resp.ScanUuid != "test-uuid-123" {
		t.Errorf("ScanUuid = %q, want %q", resp.ScanUuid, "test-uuid-123")
	}
}

func TestScanControlPause(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.ScanControlService.Pause(context.Background(), "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans/1/pause" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/1/pause")
	}
}

func TestScanControlResume(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.ScanControlService.Resume(context.Background(), "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans/1/resume" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/1/resume")
	}
}

func TestScanControlStop(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.ScanControlService.Stop(context.Background(), "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans/1/stop" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/1/stop")
	}
}

func TestScanControlVmScansStopForce(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.ScanControlService.VmScansStopForce(context.Background(), "schedule-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans/schedule-123/force-stop" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/schedule-123/force-stop")
	}
}
