package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileUpload(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"fileuploaded": "upload-ok"}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.FileService.FileUpload(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/file/upload" {
		t.Errorf("path = %q, want %q", gotPath, "/file/upload")
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	if resp.Fileuploaded != "upload-ok" {
		t.Errorf("Fileuploaded = %q, want %q", resp.Fileuploaded, "upload-ok")
	}
}
