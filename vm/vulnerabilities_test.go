package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVulnerabilitiesImport(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &VulnerabilitiesServiceVulnerabilitiesImportRequest{
		Source: "test",
	}
	err := c.VulnerabilitiesService.VulnerabilitiesImport(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/import/vulnerabilities" {
		t.Errorf("path = %q, want %q", gotPath, "/import/vulnerabilities")
	}
}

func TestVulnerabilitiesImportV2(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"job_uuid": "job-123"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &VulnerabilitiesServiceVulnerabilitiesImportV2Request{
		Vendor: "test",
	}
	resp, err := c.VulnerabilitiesService.VulnerabilitiesImportV2(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/api/v2/vulnerabilities" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v2/vulnerabilities")
	}
	if resp.JobUuid != "job-123" {
		t.Errorf("JobUuid = %q, want %q", resp.JobUuid, "job-123")
	}
}
