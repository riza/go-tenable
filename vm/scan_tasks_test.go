package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSchedule(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"enabled": true
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ScanTasksServiceScheduleRequest{
		Enabled: true,
	}
	resp, err := c.ScanTasksService.Schedule(context.Background(), "scan-123", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPut {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPut)
	}
	if gotPath != "/scans/scan-123/schedule" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/schedule")
	}
	if !resp.Enabled {
		t.Errorf("Enabled = %v, want %v", resp.Enabled, true)
	}
}

func TestCopy(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"name": "Copy of Scan",
			"id": 100
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ScanTasksServiceCopyRequest{
		Name: "Copy of Scan",
	}
	resp, err := c.ScanTasksService.Copy(context.Background(), "scan-123", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans/scan-123/copy" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/copy")
	}
	if resp.Id != 100 {
		t.Errorf("Id = %d, want %d", resp.Id, 100)
	}
}

func TestCredentialsConvert(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"uuid": "cred-123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ScanTasksServiceCredentialsConvertRequest{
		Name: "New Cred",
	}
	resp, err := c.ScanTasksService.CredentialsConvert(context.Background(), "scan-123", "cred-id", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans/scan-123/credentials/cred-id/upgrade" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/credentials/cred-id/upgrade")
	}
	if resp.Uuid != "cred-123" {
		t.Errorf("Uuid = %q, want %q", resp.Uuid, "cred-123")
	}
}

func TestImport(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"id": 100
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ScanTasksServiceImportRequest{
		File: "import.nessus",
	}
	resp, err := c.ScanTasksService.Import(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans/import" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/import")
	}
	if resp.Id != 100 {
		t.Errorf("Id = %d, want %d", resp.Id, 100)
	}
}

func TestCount(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"count": 10
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScanTasksService.Count(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/count" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/count")
	}
	if resp.Count != 10 {
		t.Errorf("Count = %d, want %d", resp.Count, 10)
	}
}

func TestTimezones(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[
			{"name": "UTC", "value": "UTC"}
		]`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScanTasksService.Timezones(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/timezones" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/timezones")
	}
	if len(resp) != 1 || resp[0].Name != "UTC" {
		t.Errorf("Timezones = %v", resp)
	}
}

func TestCheckAutoTargets(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"total_missed_targets": 2
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ScanTasksServiceCheckAutoTargetsRequest{
		TargetList: "127.0.0.1",
	}
	resp, err := c.ScanTasksService.CheckAutoTargets(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans/check-auto-targets" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/check-auto-targets")
	}
	if resp.TotalMissedTargets != 2 {
		t.Errorf("TotalMissedTargets = %d, want %d", resp.TotalMissedTargets, 2)
	}
}
