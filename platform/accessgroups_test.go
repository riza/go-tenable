package platform

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListAccessGroups(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"access_groups": [
				{"id": "ag1", "name": "Group A"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.AccessGroupsService.ListAccessGroups(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/access-groups" {
		t.Errorf("path = %q, want %q", gotPath, "/access-groups")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %v, want 1", resp.Total)
	}
}

func TestGetAccessGroupV2(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"id": "ag2",
			"name": "Group B"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.AccessGroupsService.GetAccessGroupV2(context.Background(), "ag2")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/v2/access-groups/ag2" {
		t.Errorf("path = %q, want %q", gotPath, "/v2/access-groups/ag2")
	}
	if resp.Id != "ag2" {
		t.Errorf("Id = %v, want ag2", resp.Id)
	}
}
