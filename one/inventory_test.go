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

func ptr[T any](v T) *T {
	return &v
}

func TestInventoryAssetsSearchRequest(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody InventoryAssetsSearchRequest
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
			"assets": [
				{"id": "asset1", "name": "Asset A", "type": "host"},
				{"id": "asset2", "name": "Asset B", "type": "cloud"}
			],
			"total": 2
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &InventoryAssetsSearchRequest{
		Limit: ptr(50),
		Filters: []InventorySearchFilter{
			{Property: "sources", Operator: "eq", Value: []string{"tvm"}},
		},
	}
	resp, err := c.InventoryService.SearchAssets(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/api/v1/t1/inventory/assets/search" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/assets/search")
	}
	if gotBody.Limit == nil || *gotBody.Limit != 50 {
		t.Errorf("Limit = %v, want 50", gotBody.Limit)
	}
	if len(gotBody.Filters) == 0 {
		t.Errorf("Filters is empty, want non-empty")
	} else if gotBody.Filters[0].Property != "sources" || gotBody.Filters[0].Operator != "eq" {
		t.Errorf("Filters.Sources = %v, want [tvm]", gotBody.Filters[0])
	}
	if resp.Total != 2 {
		t.Errorf("Total = %d, want 2", resp.Total)
	}
	if len(resp.Assets) != 2 || resp.Assets[0].ID != "asset1" {
		t.Errorf("Assets = %v, want 2 elements with first ID asset1", resp.Assets)
	}
}

func TestSearchFindings(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody InventoryFindingsSearchRequest
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
			"findings": [
				{"id": "vuln1", "name": "Vuln A", "severity": "High"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &InventoryFindingsSearchRequest{
		Limit: ptr(25),
		Filters: []InventorySearchFilter{
			{Property: "finding_severity", Operator: "=", Value: []string{"HIGH", "CRITICAL"}},
		},
	}
	resp, err := c.InventoryService.SearchFindings(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/api/v1/t1/inventory/findings/search" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/findings/search")
	}
	if gotBody.Limit == nil || *gotBody.Limit != 25 {
		t.Errorf("Limit = %v, want 25", gotBody.Limit)
	}
	if len(gotBody.Filters) == 0 {
		t.Errorf("Filters is empty, want non-empty")
	} else if gotBody.Filters[0].Property != "finding_severity" {
		t.Errorf("Filters.Severities = %v, want [HIGH, CRITICAL]", gotBody.Filters[0])
	}
	if resp.Total != 1 {
		t.Errorf("Total = %d, want 1", resp.Total)
	}
	if len(resp.Findings) != 1 || resp.Findings[0].ID != "vuln1" {
		t.Errorf("Findings = %v, want 1 element with ID vuln1", resp.Findings)
	}
}

func TestSearchSoftware(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"software": [
				{"id": "soft1", "name": "Nginx", "version": "1.18.0"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &InventorySoftwareSearchRequest{
		Limit: ptr(10),
	}
	resp, err := c.InventoryService.SearchSoftware(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/api/v1/t1/inventory/software/search" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/software/search")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %d, want 1", resp.Total)
	}
	if len(resp.Software) != 1 || resp.Software[0].Name != "Nginx" {
		t.Errorf("Software = %v, want 1 element with Name Nginx", resp.Software)
	}
}

func TestGetAssetProperties(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"properties": [
				{"name": "os", "type": "string"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.InventoryService.GetAssetProperties(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v1/t1/inventory/assets/properties" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/assets/properties")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %d, want 1", resp.Total)
	}
}

func TestGetFindingProperties(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"properties": [
				{"name": "severity", "type": "string"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.InventoryService.GetFindingProperties(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v1/t1/inventory/findings/properties" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/findings/properties")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %d, want 1", resp.Total)
	}
}

func TestGetSoftwareProperties(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"properties": [
				{"name": "vendor", "type": "string"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.InventoryService.GetSoftwareProperties(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v1/t1/inventory/software/properties" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/inventory/software/properties")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %d, want 1", resp.Total)
	}
}
