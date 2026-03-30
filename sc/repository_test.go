package sc

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRepositoryList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(apiResponse{
			Type: "regular",
			Response: json.RawMessage(`[
				{
					"id": "1",
					"name": "Main Repo",
					"description": "Primary",
					"type": "local",
					"dataFormat": "IPv4",
					"uuid": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeee1111"
				},
				{
					"id": "2",
					"name": "Secondary",
					"description": "Backup",
					"type": "remote",
					"dataFormat": "IPv4",
					"uuid": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeee2222"
				}
			]`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	repos, err := c.Repository.List(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/rest/repository" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/repository")
	}
	if len(repos) != 2 {
		t.Fatalf("len(repos) = %d, want 2", len(repos))
	}
	if repos[0].ID != "1" || repos[0].Name != "Main Repo" {
		t.Errorf("repos[0] = %+v, want id=1 name=Main Repo", repos[0])
	}
	if repos[1].ID != "2" || repos[1].Type != "remote" {
		t.Errorf("repos[1] = %+v, want id=2 type=remote", repos[1])
	}
}

func TestRepositoryGet(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(apiResponse{
			Type: "regular",
			Response: json.RawMessage(`{
				"id": "3",
				"name": "Detail Repo",
				"description": "Full detail",
				"type": "local",
				"dataFormat": "IPv4",
				"remoteID": "",
				"remoteIP": "",
				"running": "false",
				"downloadFormat": "v2",
				"lastSyncTime": "0",
				"createdTime": "1000",
				"modifiedTime": "2000",
				"uuid": "bbbbbbbb-cccc-dddd-eeee-ffffffff3333",
				"organizations": [
					{"id": "1", "groupAssign": "1", "name": "Org A", "description": "", "uuid": "org-uuid-1"}
				]
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	repo, err := c.Repository.Get(context.Background(), "3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/rest/repository/3" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/repository/3")
	}
	if repo.ID != "3" {
		t.Errorf("ID = %q, want %q", repo.ID, "3")
	}
	if repo.Name != "Detail Repo" {
		t.Errorf("Name = %q, want %q", repo.Name, "Detail Repo")
	}
	if repo.DataFormat != "IPv4" {
		t.Errorf("DataFormat = %q, want %q", repo.DataFormat, "IPv4")
	}
	if len(repo.Organizations) != 1 {
		t.Fatalf("len(Organizations) = %d, want 1", len(repo.Organizations))
	}
	if repo.Organizations[0].Name != "Org A" {
		t.Errorf("Organizations[0].Name = %q, want %q", repo.Organizations[0].Name, "Org A")
	}
}
