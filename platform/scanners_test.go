package platform

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListScanners(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"scanners": [
				{"id": 1, "name": "Scanner A"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScannersService.ListScanners(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scanners" {
		t.Errorf("path = %q, want %q", gotPath, "/scanners")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %v, want 1", resp.Total)
	}
}

func TestListScannerGroups(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"scanner_pools": [
				{"id": 2, "name": "Pool A"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScannersService.ListScannerGroups(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scanner-groups" {
		t.Errorf("path = %q, want %q", gotPath, "/scanner-groups")
	}
	if len(resp.ScannerPools) != 1 || resp.ScannerPools[0].Id != 2 {
		t.Errorf("ScannerPools = %v", resp.ScannerPools)
	}
}

func TestGetScannerTasks(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[
			{"id": "t1", "status": "running"}
		]`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScannersService.GetScannerTasks(context.Background(), 1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scanners/1/tasks" {
		t.Errorf("path = %q, want %q", gotPath, "/scanners/1/tasks")
	}
	if len(resp) != 1 || resp[0].Id != "t1" {
		t.Errorf("Tasks = %v", resp)
	}
}
