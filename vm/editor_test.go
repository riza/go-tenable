package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEditorDetails(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"uuid": "test-uuid",
			"name": "Editor Name"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.EditorService.Details(context.Background(), "scan", "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/editor/scan/1" {
		t.Errorf("path = %q, want %q", gotPath, "/editor/scan/1")
	}
	if resp.Uuid != "test-uuid" {
		t.Errorf("Uuid = %q, want %q", resp.Uuid, "test-uuid")
	}
}

func TestEditorListTemplates(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[
			{"uuid": "template-1", "name": "Basic Scan"}
		]`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.EditorService.ListTemplates(context.Background(), "scan")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/editor/scan/templates" {
		t.Errorf("path = %q, want %q", gotPath, "/editor/scan/templates")
	}
	if len(resp) != 1 || resp[0].Uuid != "template-1" {
		t.Errorf("Templates = %v", resp)
	}
}

func TestEditorTemplateDetails(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"name": "Basic Scan"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.EditorService.TemplateDetails(context.Background(), "scan", "template-1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/editor/scan/templates/template-1" {
		t.Errorf("path = %q, want %q", gotPath, "/editor/scan/templates/template-1")
	}
	if resp.Name != "Basic Scan" {
		t.Errorf("Name = %q, want %q", resp.Name, "Basic Scan")
	}
}

func TestEditorPluginDescription(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"plugindescription": {
				"pluginname": "Test Plugin",
				"pluginid": 12345
			}
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.EditorService.PluginDescription(context.Background(), "policy1", "family1", "12345")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/editor/policy/policy1/families/family1/plugins/12345" {
		t.Errorf("path = %q, want %q", gotPath, "/editor/policy/policy1/families/family1/plugins/12345")
	}
	if resp.Plugindescription.Pluginid != 12345 {
		t.Errorf("Pluginid = %d, want %d", resp.Plugindescription.Pluginid, 12345)
	}
}

func TestEditorAudits(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("audit file content"))
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.EditorService.Audits(context.Background(), "scan", "scan1", "file1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/editor/scan/scan1/audits/file1" {
		t.Errorf("path = %q, want %q", gotPath, "/editor/scan/scan1/audits/file1")
	}
}
