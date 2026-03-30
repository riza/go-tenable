package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScansCreate(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"uuid": "test-uuid",
			"name": "Test Scan",
			"id": 42
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ScansServiceCreateRequest{
		Settings: ScansServiceCreateRequestSettings{
			Name: "Test Scan",
		},
	}
	resp, err := c.ScansService.Create(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans" {
		t.Errorf("path = %q, want %q", gotPath, "/scans")
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	if resp.Uuid != "test-uuid" {
		t.Errorf("Uuid = %q, want %q", resp.Uuid, "test-uuid")
	}
	if resp.Name != "Test Scan" {
		t.Errorf("Name = %q, want %q", resp.Name, "Test Scan")
	}
	if resp.Id != 42 {
		t.Errorf("Id = %d, want %d", resp.Id, 42)
	}
}

func TestScansList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"scans": [
				{"id": 1, "name": "Scan 1"}
			],
			"timestamp": 123456
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScansService.List(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans" {
		t.Errorf("path = %q, want %q", gotPath, "/scans")
	}
	if len(resp.Scans) != 1 || resp.Scans[0].Id != 1 {
		t.Errorf("Scans = %v", resp.Scans)
	}
}

func TestScansDetails(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"info": {
				"uuid": "test-uuid",
				"name": "Test Scan"
			}
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScansService.Details(context.Background(), "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/1" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/1")
	}
	if resp.Info.Name != "Test Scan" {
		t.Errorf("Info.Name = %q, want %q", resp.Info.Name, "Test Scan")
	}
}

func TestScansConfigure(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"uuid": "test-uuid",
			"name": "Updated Scan"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ScansServiceConfigureRequest{
		Settings: ScansServiceConfigureRequestSettings{
			Name: "Updated Scan",
		},
	}
	_, err := c.ScansService.Configure(context.Background(), "1", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPut {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPut)
	}
	if gotPath != "/scans/1" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/1")
	}
}

func TestScansDelete(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.ScansService.Delete(context.Background(), "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodDelete {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodDelete)
	}
	if gotPath != "/scans/1" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/1")
	}
}
