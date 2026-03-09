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
	var gotBody interface{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		if r.URL.Query().Get("limit") != "100" {
			t.Errorf("limit query = %q, want 100", r.URL.Query().Get("limit"))
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		if err := json.Unmarshal(data, &gotBody); err != nil {
			t.Fatalf("failed to unmarshal request body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"data": [
				{"id": "path1", "name": "Path A", "score": 95},
				{"id": "path2", "name": "Path B", "score": 80}
			],
			"total": 2
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &APASearchAttackPathsRequest{
		Limit: 100,
		Filter: APAFilterCondition{
			Property: "asset_id",
			Operator: "eq",
			Value:    "asset1",
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

	filter, ok := gotBody.(map[string]interface{})
	if !ok {
		t.Fatalf("gotBody is not map[string]interface{}, got: %T", gotBody)
	}
	if filter["property"] != "asset_id" || filter["operator"] != "eq" || filter["value"] != "asset1" {
		t.Errorf("gotBody = %v, want {property: asset_id, operator: eq, value: asset1}", gotBody)
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
		if r.URL.Query().Get("limit") != "150" {
			t.Errorf("limit query = %q, want 150", r.URL.Query().Get("limit"))
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"data": [
				{"mitre_id": "tech1", "technique_name": "Technique A", "score": 95, "priority": "high", "tactics": ["tacticA"]}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &APASearchAttackTechniquesRequest{
		Limit: 150,
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
	if len(resp.Techniques) != 1 {
		t.Fatalf("Got %d techniques, want 1", len(resp.Techniques))
	}
	if resp.Techniques[0].MitreId != "tech1" {
		t.Errorf("Technique MitreId = %q, want %q", resp.Techniques[0].MitreId, "tech1")
	}
}
