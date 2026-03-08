package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHistory(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"history": [
				{"id": 1, "status": "completed"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScanHistoryService.History(context.Background(), "scan-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/scan-123/history" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/history")
	}
	if len(resp.History) != 1 || resp.History[0].Id != 1 {
		t.Errorf("History = %v", resp.History)
	}
}

func TestHistoryDetails(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"uuid": "history-123",
			"status": "completed"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ScanHistoryService.HistoryDetails(context.Background(), "scan-123", "history-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scans/scan-123/history/history-123" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/history/history-123")
	}
	if resp.Uuid != "history-123" {
		t.Errorf("Uuid = %q, want %q", resp.Uuid, "history-123")
	}
}

func TestDeleteHistory(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.ScanHistoryService.DeleteHistory(context.Background(), "scan-123", "history-123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodDelete {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodDelete)
	}
	if gotPath != "/scans/scan-123/history/history-123" {
		t.Errorf("path = %q, want %q", gotPath, "/scans/scan-123/history/history-123")
	}
}
