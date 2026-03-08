package one

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExportAssets(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"export_id": "exp123",
			"format": "json",
			"status": "QUEUED"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &InventoryExportRequest{
		Format: "json",
	}
	resp, err := c.ExportService.ExportAssets(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/api/v1/t1/inventory/export/assets" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/export/assets")
	}
	if resp.ExportId != "exp123" {
		t.Errorf("ExportId = %q, want %q", resp.ExportId, "exp123")
	}
}

func TestExportFindings(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody InventoryExportRequest
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		if err := json.Unmarshal(data, &gotBody); err != nil {
			t.Fatalf("failed to unmarshal request body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"export_id": "exp456",
			"format": "csv",
			"status": "PROCESSING"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &InventoryExportRequest{
		Format: "csv",
		Limit:  1000,
	}
	resp, err := c.ExportService.ExportFindings(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/api/v1/t1/inventory/export/findings" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/export/findings")
	}
	if gotBody.Format != "csv" {
		t.Errorf("Format = %q, want csv", gotBody.Format)
	}
	if resp.ExportId != "exp456" {
		t.Errorf("ExportId = %q, want %q", resp.ExportId, "exp456")
	}
}

func TestGetAssetsExportStatus(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"export_id": "exp123",
			"status": "FINISHED",
			"total_objects": 500,
			"chunks": [1, 2]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportService.GetAssetsExportStatus(context.Background(), "exp123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v1/t1/inventory/export/assets/status" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/export/assets/status")
	}
	if resp.Status != "FINISHED" {
		t.Errorf("Status = %q, want FINISHED", resp.Status)
	}
	if len(resp.Chunks) != 2 || resp.Chunks[0] != 1 {
		t.Errorf("Chunks = %v, want [1, 2]", resp.Chunks)
	}
}

func TestGetFindingsExportStatus(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"export_id": "exp456",
			"status": "FINISHED"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportService.GetFindingsExportStatus(context.Background(), "exp456")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v1/t1/inventory/export/findings/status" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/export/findings/status")
	}
	if resp.Status != "FINISHED" {
		t.Errorf("Status = %q, want FINISHED", resp.Status)
	}
}

func TestGetExportStatus(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"export_id": "exp789",
			"status": "QUEUED"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportService.GetExportStatus(context.Background(), "exp789")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v1/t1/inventory/export/exp789/status" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/export/exp789/status")
	}
	if resp.ExportId != "exp789" {
		t.Errorf("ExportId = %q, want exp789", resp.ExportId)
	}
}

func TestDownloadExportChunk(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/octet-stream")
		fmt.Fprint(w, "chunk_data")
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportService.DownloadExportChunk(context.Background(), "exp123", 1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	// Note: string(rune(1+'0')) is "1"
	if gotPath != "/api/v1/t1/inventory/export/exp123/download/1" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/export/exp123/download/1")
	}
	if string(resp) != "chunk_data" {
		t.Errorf("Response = %q, want chunk_data", string(resp))
	}
}
