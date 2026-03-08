package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPoliciesList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[
			{"id": 1, "name": "Policy 1"}
		]`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.PoliciesService.PoliciesList(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/policies" {
		t.Errorf("path = %q, want %q", gotPath, "/policies")
	}
	if len(resp) != 1 || resp[0].Id != 1 {
		t.Errorf("Policies = %v", resp)
	}
}

func TestPoliciesCreate(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"policy_id": 1,
			"policy_name": "New Policy"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &PoliciesServicePoliciesCreateRequest{
		Settings: PoliciesServicePoliciesCreateRequestSettings{
			Name: "New Policy",
		},
	}
	resp, err := c.PoliciesService.PoliciesCreate(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/policies" {
		t.Errorf("path = %q, want %q", gotPath, "/policies")
	}
	if resp.PolicyId != 1 {
		t.Errorf("PolicyId = %d", resp.PolicyId)
	}
}

func TestPoliciesDetails(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"uuid": "test-uuid"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.PoliciesService.PoliciesDetails(context.Background(), "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/policies/1" {
		t.Errorf("path = %q, want %q", gotPath, "/policies/1")
	}
	if resp.Uuid != "test-uuid" {
		t.Errorf("Uuid = %q, want %q", resp.Uuid, "test-uuid")
	}
}

func TestPoliciesDelete(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.PoliciesService.PoliciesDelete(context.Background(), "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodDelete {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodDelete)
	}
	if gotPath != "/policies/1" {
		t.Errorf("path = %q, want %q", gotPath, "/policies/1")
	}
}

func TestPoliciesImport(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"id": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &PoliciesServicePoliciesImportRequest{
		File: "test",
	}
	resp, err := c.PoliciesService.PoliciesImport(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/policies/import" {
		t.Errorf("path = %q, want %q", gotPath, "/policies/import")
	}
	if resp.Id != 1 {
		t.Errorf("Id = %v", resp.Id)
	}
}
