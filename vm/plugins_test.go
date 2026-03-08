package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPluginsList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"data": {
				"plugin_details": [
					{"id": 1, "name": "Plugin 1"}
				]
			},
			"total_count": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.PluginsService.PluginsList(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/plugins/plugin" {
		t.Errorf("path = %q, want %q", gotPath, "/plugins/plugin")
	}
	if len(resp.Data.PluginDetails) != 1 || resp.Data.PluginDetails[0].Id != 1 {
		t.Errorf("PluginDetails = %v", resp.Data.PluginDetails)
	}
}

func TestPluginsDetails(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"id": 1,
			"name": "Plugin 1"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.PluginsService.PluginsDetails(context.Background(), "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/plugins/plugin/1" {
		t.Errorf("path = %q, want %q", gotPath, "/plugins/plugin/1")
	}
	if resp.Id != 1 {
		t.Errorf("Id = %d", resp.Id)
	}
}

func TestPluginsFamiliesList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"families": [
				{"id": 1, "name": "Family 1"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.PluginsService.PluginsFamiliesList(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/plugins/families" {
		t.Errorf("path = %q, want %q", gotPath, "/plugins/families")
	}
	if len(resp.Families) != 1 || resp.Families[0].Id != 1 {
		t.Errorf("Families = %v", resp.Families)
	}
}

func TestPluginsFamilyDetailsId(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"id": 1,
			"plugins": [
				{"id": 100, "name": "Plugin 100"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.PluginsService.PluginsFamilyDetailsId(context.Background(), "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/plugins/families/1" {
		t.Errorf("path = %q, want %q", gotPath, "/plugins/families/1")
	}
	if len(resp.Plugins) != 1 || resp.Plugins[0].Id != 100 {
		t.Errorf("Plugins = %v", resp.Plugins)
	}
}

func TestPluginsFamilyDetailsName(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"id": 1,
			"plugins": [
				{"id": 100, "name": "Plugin 100"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &PluginsServicePluginsFamilyDetailsNameRequest{
		Name: "Family 1",
	}
	resp, err := c.PluginsService.PluginsFamilyDetailsName(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/plugins/families/_byName" {
		t.Errorf("path = %q, want %q", gotPath, "/plugins/families/_byName")
	}
	if len(resp.Plugins) != 1 || resp.Plugins[0].Id != 100 {
		t.Errorf("Plugins = %v", resp.Plugins)
	}
}
