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
		if r.URL.Query().Get("sort") != "priority:desc" {
			t.Errorf("sort query = %q, want priority:desc", r.URL.Query().Get("sort"))
		}
		if r.URL.Query().Get("exclude_resolved") != "false" {
			t.Errorf("exclude_resolved query = %q, want false", r.URL.Query().Get("exclude_resolved"))
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		if len(data) > 0 {
			if err := json.Unmarshal(data, &gotBody); err != nil {
				t.Fatalf("failed to unmarshal request body: %v", err)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"data": [
				{
					"vector_id": "vec1",
					"name": "Path A",
					"path_status": "to_do",
					"first_aes": 8.5,
					"last_acr": 9.0,
					"techniques": [
						{"technique": "Initial Access", "external_id": "T1190"}
					]
				},
				{
					"vector_id": "vec2",
					"name": "Path B",
					"path_status": "in_progress",
					"is_new": true
				}
			],
			"total": 2
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	excludeResolved := false
	req := &APASearchAttackPathsRequest{
		Limit:           100,
		Sort:            "priority:desc",
		ExcludeResolved: &excludeResolved,
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
	if len(resp.AttackPaths) != 2 {
		t.Fatalf("AttackPaths count = %d, want 2", len(resp.AttackPaths))
	}
	if resp.AttackPaths[0].VectorID != "vec1" {
		t.Errorf("AttackPaths[0].VectorID = %q, want %q", resp.AttackPaths[0].VectorID, "vec1")
	}
	if resp.AttackPaths[0].PathStatus != "to_do" {
		t.Errorf("AttackPaths[0].PathStatus = %q, want %q", resp.AttackPaths[0].PathStatus, "to_do")
	}
	if len(resp.AttackPaths[0].Techniques) != 1 {
		t.Fatalf("AttackPaths[0].Techniques count = %d, want 1", len(resp.AttackPaths[0].Techniques))
	}
	if resp.AttackPaths[0].Techniques[0].Technique != "Initial Access" {
		t.Errorf("Techniques[0].Technique = %q, want %q", resp.AttackPaths[0].Techniques[0].Technique, "Initial Access")
	}
	if !resp.AttackPaths[1].IsNew {
		t.Errorf("AttackPaths[1].IsNew = false, want true")
	}
}

func TestSearchAttackPathsNoFilter(t *testing.T) {
	var gotContentLength int64
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotContentLength = r.ContentLength
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"data": [], "total": 0}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.AttackPathService.SearchAttackPaths(context.Background(), &APASearchAttackPathsRequest{
		Limit: 100,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotContentLength > 0 {
		t.Errorf("expected no body when filter is nil, got Content-Length %d", gotContentLength)
	}
	if resp.Total != 0 {
		t.Errorf("Total = %d, want 0", resp.Total)
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
	if resp.Techniques[0].MitreID != "tech1" {
		t.Errorf("Technique MitreID = %q, want %q", resp.Techniques[0].MitreID, "tech1")
	}
}
