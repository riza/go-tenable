package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFoldersCreate(t *testing.T) {
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
	req := &FoldersServiceFoldersCreateRequest{
		Name: "New Folder",
	}
	resp, err := c.FoldersService.FoldersCreate(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/folders" {
		t.Errorf("path = %q, want %q", gotPath, "/folders")
	}
	if resp.Id != 1 {
		t.Errorf("Id = %d", resp.Id)
	}
}

func TestFoldersList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[
			{"id": 1, "name": "Folder 1"}
		]`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.FoldersService.FoldersList(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/folders" {
		t.Errorf("path = %q, want %q", gotPath, "/folders")
	}
	if len(resp) != 1 || resp[0].Id != 1 {
		t.Errorf("Folders = %v", resp)
	}
}

func TestFoldersEdit(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &FoldersServiceFoldersEditRequest{
		Name: "Renamed Folder",
	}
	err := c.FoldersService.FoldersEdit(context.Background(), "1", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPut {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPut)
	}
	if gotPath != "/folders/1" {
		t.Errorf("path = %q, want %q", gotPath, "/folders/1")
	}
}

func TestFoldersDelete(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.FoldersService.FoldersDelete(context.Background(), "1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodDelete {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodDelete)
	}
	if gotPath != "/folders/1" {
		t.Errorf("path = %q, want %q", gotPath, "/folders/1")
	}
}
