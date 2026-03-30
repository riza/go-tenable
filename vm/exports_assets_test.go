package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExportAssetsV1(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"export_uuid": "export-123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ExportsAssetsServiceExportAssetsV1Request{
		ChunkSize: 100,
	}
	resp, err := c.ExportsAssetsService.ExportAssetsV1(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/assets/export" {
		t.Errorf("path = %q, want %q", gotPath, "/assets/export")
	}
	if resp.ExportUuid != "export-123" {
		t.Errorf("ExportUuid = %q, want %q", resp.ExportUuid, "export-123")
	}
}

func TestExportAssetsV2(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"export_uuid": "export-v2-123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &ExportsAssetsServiceExportAssetsV2Request{
		ChunkSize: 200,
	}
	resp, err := c.ExportsAssetsService.ExportAssetsV2(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/assets/v2/export" {
		t.Errorf("path = %q, want %q", gotPath, "/assets/v2/export")
	}
	if resp.ExportUuid != "export-v2-123" {
		t.Errorf("ExportUuid = %q, want %q", resp.ExportUuid, "export-v2-123")
	}
}

func TestExportsAssetsExportStatus(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"status": "FINISHED",
			"chunks_available": [1, 2]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportsAssetsService.ExportsAssetsExportStatus(context.Background(), "export-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/assets/export/export-123/status" {
		t.Errorf("path = %q, want %q", gotPath, "/assets/export/export-123/status")
	}
	if resp.Status != "FINISHED" {
		t.Errorf("Status = %q, want %q", resp.Status, "FINISHED")
	}
}

func TestExportsAssetsExportStatusRecent(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"exports": [
				{"uuid": "export-123", "status": "FINISHED"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExportsAssetsService.ExportsAssetsExportStatusRecent(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/assets/export/status" {
		t.Errorf("path = %q, want %q", gotPath, "/assets/export/status")
	}
	if len(resp.Exports) != 1 || resp.Exports[0].Uuid != "export-123" {
		t.Errorf("Exports = %v", resp.Exports)
	}
}

func TestExportsAssetsDownloadChunk(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("chunk data"))
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.ExportsAssetsService.ExportsAssetsDownloadChunk(context.Background(), "export-123", "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/assets/export/export-123/chunks/1" {
		t.Errorf("path = %q, want %q", gotPath, "/assets/export/export-123/chunks/1")
	}
}

func TestExportsAssetsExportCancel(t *testing.T) {
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
	resp, err := c.ExportsAssetsService.ExportsAssetsExportCancel(context.Background(), "export-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/assets/export/export-123/cancel" {
		t.Errorf("path = %q, want %q", gotPath, "/assets/export/export-123/cancel")
	}
	if resp.Status != "CANCELLED" {
		t.Errorf("Status = %q, want %q", resp.Status, "CANCELLED")
	}
}
