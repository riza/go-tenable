package platform

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListAgents(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"agents": [
				{"id": "ag1", "name": "Agent A"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.AgentsService.ListAgents(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scanners/null/agents" {
		t.Errorf("path = %q, want %q", gotPath, "/scanners/null/agents")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %v, want 1", resp.Total)
	}
}

func TestGetAgentConfig(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"setting": "value"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.AgentsService.GetAgentConfig(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/scanners/null/agents/config" {
		t.Errorf("path = %q, want %q", gotPath, "/scanners/null/agents/config")
	}
	if resp["setting"] != "value" {
		t.Errorf("setting = %v, want value", resp["setting"])
	}
}
