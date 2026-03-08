package platform

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPISecuritySettings(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"api_access_key_enabled": true,
			"api_key_idle_timeout": 30,
			"api_key_max_lifetime": 60
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.AccessControlService.GetAPISecuritySettings(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/access-control/v1/api-security-settings" {
		t.Errorf("path = %q, want %q", gotPath, "/access-control/v1/api-security-settings")
	}
	if !resp.APIAccessKeyEnabled {
		t.Errorf("APIAccessKeyEnabled = %v, want true", resp.APIAccessKeyEnabled)
	}
	if resp.APIKeyIdleTimeout != 30 {
		t.Errorf("APIKeyIdleTimeout = %v, want 30", resp.APIKeyIdleTimeout)
	}
}

func TestListGroups(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"groups": [
				{"id": 1, "name": "Admin"},
				{"id": 2, "name": "User"}
			],
			"total": 2
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.GroupsService.ListGroups(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/groups" {
		t.Errorf("path = %q, want %q", gotPath, "/groups")
	}
	if resp.Total != 2 {
		t.Errorf("Total = %v, want 2", resp.Total)
	}
	if len(resp.Groups) != 2 || resp.Groups[0].Name != "Admin" {
		t.Errorf("Groups = %v", resp.Groups)
	}
}

func TestListPermissions(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"permissions": [
				{"uuid": "abc", "name": "PermA"}
			],
			"total": 1
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.PermissionsService.ListPermissions(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v3/access-control/permissions" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v3/access-control/permissions")
	}
	if resp.Total != 1 {
		t.Errorf("Total = %v, want 1", resp.Total)
	}
}
