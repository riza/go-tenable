package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListAssets(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"assets": [
				{"id": "asset1"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.AssetsService.ListAssets(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/assets" {
		t.Errorf("path = %q, want %q", gotPath, "/assets")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %v, want 1", resp.Total)
	}
}

func TestAssetInfo(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"id": "asset1",
			"name": "Host 1"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.AssetsService.AssetInfo(context.Background(), "asset1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/assets/asset1" {
		t.Errorf("path = %q, want %q", gotPath, "/assets/asset1")
	}
	if resp.Name != "Host 1" {
		t.Errorf("Name = %q, want %q", resp.Name, "Host 1")
	}
}

func TestAssetsImport(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"asset_import_job_uuid": "job-123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &AssetsServiceImportRequest{
		Source: "test",
	}
	resp, err := c.AssetsService.Import(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/import/assets" {
		t.Errorf("path = %q, want %q", gotPath, "/import/assets")
	}
	if resp.AssetImportJobUuid != "job-123" {
		t.Errorf("AssetImportJobUuid = %q, want %q", resp.AssetImportJobUuid, "job-123")
	}
}

func TestListImportJobs(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"asset_import_jobs": [
				{"job_id": "job1"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.AssetsService.ListImportJobs(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/import/asset-jobs" {
		t.Errorf("path = %q, want %q", gotPath, "/import/asset-jobs")
	}
	if len(resp.AssetImportJobs) != 1 || resp.AssetImportJobs[0].JobId != "job1" {
		t.Errorf("AssetImportJobs = %v", resp.AssetImportJobs)
	}
}
