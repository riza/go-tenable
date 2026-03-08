package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAssetAttributesList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"attributes": [
				{"id": "attr1", "name": "Environment"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.AssetAttributesService.AssetAttributesList(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v3/assets/attributes" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v3/assets/attributes")
	}
	if len(resp.Attributes) != 1 || resp.Attributes[0].Id != "attr1" {
		t.Errorf("Attributes = %v", resp.Attributes)
	}
}

func TestAssetAttributesCreate(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &AssetAttributesServiceAssetAttributesCreateRequest{
		Attributes: []AssetAttributesServiceAssetAttributesCreateRequestAttributesItem{
			{Name: "Environment"},
		},
	}
	err := c.AssetAttributesService.AssetAttributesCreate(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/api/v3/assets/attributes" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v3/assets/attributes")
	}
}

func TestAssetAttributesUpdate(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &AssetAttributesServiceAssetAttributesUpdateRequest{
		Description: "Updated desc",
	}
	err := c.AssetAttributesService.AssetAttributesUpdate(context.Background(), "attr1", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPut {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPut)
	}
	if gotPath != "/api/v3/assets/attributes/attr1" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v3/assets/attributes/attr1")
	}
}

func TestAssetAttributesAssign(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &AssetAttributesServiceAssetAttributesAssignRequest{
		Attributes: []AssetAttributesServiceAssetAttributesAssignRequestAttributesItem{
			{Id: "attr1", Value: "Prod"},
		},
	}
	err := c.AssetAttributesService.AssetAttributesAssign(context.Background(), "asset1", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPut {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPut)
	}
	if gotPath != "/api/v3/assets/asset1/attributes" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v3/assets/asset1/attributes")
	}
}

func TestAssetAttributesAssignedList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"attributes": [
				{"id": "attr1", "value": "Prod"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.AssetAttributesService.AssetAttributesAssignedList(context.Background(), "asset1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v3/assets/asset1/attributes" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v3/assets/asset1/attributes")
	}
	if len(resp.Attributes) != 1 || resp.Attributes[0].Value != "Prod" {
		t.Errorf("Attributes = %v", resp.Attributes)
	}
}
