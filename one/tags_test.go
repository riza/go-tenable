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

func TestSearchTags(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody TagsSearchRequest
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
			"tags": [
				{"uuid": "tag1", "key": "env", "value": "prod"},
				{"uuid": "tag2", "key": "env", "value": "dev"}
			],
			"total": 2
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &TagsSearchRequest{
		Limit: 20,
		Filters: TagsSearchFilters{
			Keys: []string{"env"},
		},
	}
	resp, err := c.TagsService.SearchTags(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/api/v1/t1/tags/search" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/tags/search")
	}
	if gotBody.Limit != 20 {
		t.Errorf("Limit = %d, want 20", gotBody.Limit)
	}
	if len(gotBody.Filters.Keys) != 1 || gotBody.Filters.Keys[0] != "env" {
		t.Errorf("Filters.Keys = %v, want [env]", gotBody.Filters.Keys)
	}
	if resp.Total != 2 {
		t.Errorf("Total = %d, want 2", resp.Total)
	}
	if len(resp.Tags) != 2 || resp.Tags[0].Uuid != "tag1" {
		t.Errorf("Tags = %v, want 2 elements with first UUID tag1", resp.Tags)
	}
}

func TestGetTagsProperties(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"properties": [
				{"name": "key", "type": "string"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.TagsService.GetProperties(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v1/t1/tags/properties" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/tags/properties")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %d, want 1", resp.Total)
	}
}
