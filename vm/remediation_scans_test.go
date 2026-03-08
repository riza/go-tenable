package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemediationCreate(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"uuid": "rem-scan-123",
			"name": "Remediation Scan"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &RemediationScansServiceRemediationCreateRequest{
		Settings: RemediationScansServiceRemediationCreateRequestSettings{
			Name: "Remediation Scan",
		},
	}
	resp, err := c.RemediationScansService.RemediationCreate(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/scans/remediation" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/remediation")
	}
	if resp.Uuid != "rem-scan-123" {
		t.Errorf("Uuid = %q, want %q", resp.Uuid, "rem-scan-123")
	}
}

func TestRemediationList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"scans": [
				{"id": 1, "name": "Scan 1"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.RemediationScansService.RemediationList(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/remediation" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/remediation")
	}
	if len(resp.Scans) != 1 || resp.Scans[0].Id != 1 {
		t.Errorf("Scans = %v", resp.Scans)
	}
}
