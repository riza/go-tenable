package platform

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListUsers(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"users": [
				{"id": 1, "username": "admin"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.UsersService.ListUsers(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/users" {
		t.Errorf("path = %q, want %q", gotPath, "/users")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %v, want 1", resp.Total)
	}
}

func TestGetUserRoles(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"roles": [
				"admin"
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.UsersService.GetUserRoles(context.Background(), "u1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/access-control/v1/users/u1/roles" {
		t.Errorf("path = %q, want %q", gotPath, "/access-control/v1/users/u1/roles")
	}
	if len(resp.Roles) != 1 || resp.Roles[0] != "admin" {
		t.Errorf("Roles = %v", resp.Roles)
	}
}

func TestImpersonateUser(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"username": "admin"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.UsersService.ImpersonateUser(context.Background(), "2")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/users/2/impersonate" {
		t.Errorf("path = %q, want %q", gotPath, "/users/2/impersonate")
	}
	if resp.Username != "admin" {
		t.Errorf("Username = %v", resp.Username)
	}
}
