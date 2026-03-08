package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFiltersAgentsList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"filters": [
				{"name": "filter1"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.FiltersService.FiltersAgentsList(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/filters/scans/agents" {
		t.Errorf("path = %q, want %q", gotPath, "/filters/scans/agents")
	}
	if len(resp.Filters) != 1 || resp.Filters[0].Name != "filter1" {
		t.Errorf("Filters = %v", resp.Filters)
	}
}

func TestFiltersAssetsList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"filters": [
				{"name": "filter2"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.FiltersService.FiltersAssetsList(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/filters/workbenches/assets" {
		t.Errorf("path = %q, want %q", gotPath, "/filters/workbenches/assets")
	}
	if len(resp.Filters) != 1 || resp.Filters[0].Name != "filter2" {
		t.Errorf("Filters = %v", resp.Filters)
	}
}

func TestFiltersAssetsListV2(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"filters": [
				{"name": "filter3"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &FiltersServiceFiltersAssetsListV2Request{
		TagUuids: []string{"tag1"},
	}
	resp, err := c.FiltersService.FiltersAssetsListV2(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/filters/workbenches/assets" {
		t.Errorf("path = %q, want %q", gotPath, "/filters/workbenches/assets")
	}
	if len(resp.Filters) != 1 || resp.Filters[0].Name != "filter3" {
		t.Errorf("Filters = %v", resp.Filters)
	}
}
