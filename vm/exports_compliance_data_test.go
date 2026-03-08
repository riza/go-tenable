package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExportsComplianceCreate(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"export_uuid": "comp-export-123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ExportsComplianceDataServiceExportsComplianceCreateRequest{
		NumFindings: 100,
	}
	resp, err := c.ExportsComplianceDataService.ExportsComplianceCreate(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/compliance/export" {
		t.Errorf("path = %q, want %q", gotPath, "/compliance/export")
	}
	if resp.ExportUuid != "comp-export-123" {
		t.Errorf("ExportUuid = %q, want %q", resp.ExportUuid, "comp-export-123")
	}
}

func TestExportsComplianceStatus(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"status": "FINISHED",
			"uuid": "comp-export-123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportsComplianceDataService.ExportsComplianceStatus(context.Background(), "comp-export-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/compliance/export/comp-export-123/status" {
		t.Errorf("path = %q, want %q", gotPath, "/compliance/export/comp-export-123/status")
	}
	if resp.Status != "FINISHED" {
		t.Errorf("Status = %q, want %q", resp.Status, "FINISHED")
	}
}

func TestExportsComplianceStatusList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"exports": [
				{"uuid": "comp-export-123", "status": "FINISHED"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportsComplianceDataService.ExportsComplianceStatusList(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/compliance/export/status" {
		t.Errorf("path = %q, want %q", gotPath, "/compliance/export/status")
	}
	if len(resp.Exports) != 1 || resp.Exports[0].Uuid != "comp-export-123" {
		t.Errorf("Exports = %v", resp.Exports)
	}
}

func TestExportsComplianceDownload(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[
			{"asset_uuid": "asset-1", "plugin_id": 1000}
		]`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportsComplianceDataService.ExportsComplianceDownload(context.Background(), "export-123", "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/compliance/export/export-123/chunks/1" {
		t.Errorf("path = %q, want %q", gotPath, "/compliance/export/export-123/chunks/1")
	}
	if len(resp) != 1 || resp[0].AssetUuid != "asset-1" {
		t.Errorf("DownloadResponse = %v", resp)
	}
}

func TestExportsComplianceCancel(t *testing.T) {
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
	resp, err := c.ExportsComplianceDataService.ExportsComplianceCancel(context.Background(), "export-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/compliance/export/export-123/cancel" {
		t.Errorf("path = %q, want %q", gotPath, "/compliance/export/export-123/cancel")
	}
	if resp.Status != "CANCELLED" {
		t.Errorf("Status = %q, want %q", resp.Status, "CANCELLED")
	}
}
