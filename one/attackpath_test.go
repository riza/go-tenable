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

func TestSearchAttackPaths(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody APASearchAttackPathsRequest
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
			"attack_paths": [
				{"id": "path1", "name": "Path A", "score": 95},
				{"id": "path2", "name": "Path B", "score": 80}
			],
			"total": 2
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &APASearchAttackPathsRequest{
		Limit: 10,
		Filters: APASearchAttackPathsFilters{
			AssetIds: []string{"asset1"},
		},
	}
	resp, err := c.AttackPathService.SearchAttackPaths(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/api/v1/t1/apa/top-attack-paths/search" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/apa/top-attack-paths/search")
	}
	if gotBody.Limit != 10 {
		t.Errorf("gotBody.Limit = %d, want 10", gotBody.Limit)
	}
	if len(gotBody.Filters.AssetIds) != 1 || gotBody.Filters.AssetIds[0] != "asset1" {
		t.Errorf("gotBody.Filters.AssetIds = %v, want [asset1]", gotBody.Filters.AssetIds)
	}
	if resp.Total != 2 {
		t.Errorf("Total = %d, want 2", resp.Total)
	}
	if len(resp.AttackPaths) != 2 || resp.AttackPaths[0].Id != "path1" {
		t.Errorf("AttackPaths = %v, want 2 elements with first ID path1", resp.AttackPaths)
	}
}

func TestSearchAttackTechniques(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"attack_techniques": [
				{"id": "tech1", "name": "Technique A", "score": 95, "severity": "High"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &APASearchAttackTechniquesRequest{
		Limit: 5,
	}
	resp, err := c.AttackPathService.SearchAttackTechniques(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/api/v1/t1/apa/top-attack-techniques/search" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/apa/top-attack-techniques/search")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %d, want 1", resp.Total)
	}
	if len(resp.Techniques) != 1 || resp.Techniques[0].Id != "tech1" {
		t.Errorf("Techniques = %v, want 1 element with ID tech1", resp.Techniques)
	}
}
