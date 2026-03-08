package vm

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSharedCollectionsCreate(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"message": "created",
			"request_status_id": "req-1"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &SharedCollectionsServiceSharedCollectionsCreateRequest{
		Name: "New Collection",
	}
	resp, err := c.SharedCollectionsService.SharedCollectionsCreate(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/shared-collections" {
		t.Errorf("path = %q, want %q", gotPath, "/shared-collections")
	}
	if resp.RequestStatusId != "req-1" {
		t.Errorf("RequestStatusId = %q, want %q", resp.RequestStatusId, "req-1")
	}
}

func TestSharedCollectionsList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[
			{"uuid": "col-1", "name": "Collection 1"}
		]`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.SharedCollectionsService.SharedCollectionsList(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/shared-collections" {
		t.Errorf("path = %q, want %q", gotPath, "/shared-collections")
	}
	if len(resp) != 1 || resp[0].Uuid != "col-1" {
		t.Errorf("Collections = %v", resp)
	}
}

func TestSharedCollectionsDetails(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"uuid": "col-1",
			"name": "Collection 1"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.SharedCollectionsService.SharedCollectionsDetails(context.Background(), "col-1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/shared-collections/col-1" {
		t.Errorf("path = %q, want %q", gotPath, "/shared-collections/col-1")
	}
	if resp.Uuid != "col-1" {
		t.Errorf("Uuid = %q, want %q", resp.Uuid, "col-1")
	}
}

func TestSharedCollectionsUpdate(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"message": "updated"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &SharedCollectionsServiceSharedCollectionsUpdateRequest{
		Name: "Updated Collection",
	}
	resp, err := c.SharedCollectionsService.SharedCollectionsUpdate(context.Background(), "col-1", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPut {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPut)
	}
	if gotPath != "/shared-collections/col-1" {
		t.Errorf("path = %q, want %q", gotPath, "/shared-collections/col-1")
	}
	if resp.Message != "updated" {
		t.Errorf("Message = %q, want %q", resp.Message, "updated")
	}
}

func TestSharedCollectionsDelete(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"message": "deleted"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.SharedCollectionsService.SharedCollectionsDelete(context.Background(), "col-1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodDelete {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodDelete)
	}
	if gotPath != "/shared-collections/col-1" {
		t.Errorf("path = %q, want %q", gotPath, "/shared-collections/col-1")
	}
	if resp.Message != "deleted" {
		t.Errorf("Message = %q, want %q", resp.Message, "deleted")
	}
}

func TestSharedCollectionsDetailsByName(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"uuid": "col-1",
			"name": "Collection 1"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &SharedCollectionsServiceSharedCollectionsDetailsByNameRequest{
		Name: "Collection 1",
	}
	resp, err := c.SharedCollectionsService.SharedCollectionsDetailsByName(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/shared-collections/_byName" {
		t.Errorf("path = %q, want %q", gotPath, "/shared-collections/_byName")
	}
	if resp.Uuid != "col-1" {
		t.Errorf("Uuid = %q, want %q", resp.Uuid, "col-1")
	}
}

func TestSharedCollectionsJobStatus(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"status": "COMPLETED"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.SharedCollectionsService.SharedCollectionsJobStatus(context.Background(), "req-1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/shared-collections/request-status/req-1" {
		t.Errorf("path = %q, want %q", gotPath, "/shared-collections/request-status/req-1")
	}
	if resp.Status != "COMPLETED" {
		t.Errorf("Status = %q, want %q", resp.Status, "COMPLETED")
	}
}

func TestSharedCollectionsConfigAdd(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"message": "added"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &SharedCollectionsServiceSharedCollectionsConfigAddRequest{
		ScanConfigs: []string{"scan-1"},
	}
	resp, err := c.SharedCollectionsService.SharedCollectionsConfigAdd(context.Background(), "col-1", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/shared-collections/col-1/scan-configs" {
		t.Errorf("path = %q, want %q", gotPath, "/shared-collections/col-1/scan-configs")
	}
	if resp.Message != "added" {
		t.Errorf("Message = %q, want %q", resp.Message, "added")
	}
}

func TestSharedCollectionsConfigList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"scans": [
				{"uuid": "scan-1"}
			]
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.SharedCollectionsService.SharedCollectionsConfigList(context.Background(), "col-1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/shared-collections/col-1/scan-configs" {
		t.Errorf("path = %q, want %q", gotPath, "/shared-collections/col-1/scan-configs")
	}
	if len(resp.Scans) != 1 || resp.Scans[0].Uuid != "scan-1" {
		t.Errorf("Scans = %v", resp.Scans)
	}
}

func TestSharedCollectionsConfigsRemove(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"message": "removed"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	req := &SharedCollectionsServiceSharedCollectionsConfigsRemoveRequest{
		ScanConfigs: []string{"scan-1"},
	}
	resp, err := c.SharedCollectionsService.SharedCollectionsConfigsRemove(context.Background(), "col-1", req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodDelete {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodDelete)
	}
	if gotPath != "/shared-collections/col-1/scan-configs" {
		t.Errorf("path = %q, want %q", gotPath, "/shared-collections/col-1/scan-configs")
	}
	if resp.Message != "removed" {
		t.Errorf("Message = %q, want %q", resp.Message, "removed")
	}
}
