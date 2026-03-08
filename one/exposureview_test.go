package one

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListCards(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"cards": [
				{"id": "card1", "name": "Card A", "type": "exposure"},
				{"id": "card2", "name": "Card B", "type": "trend"}
			],
			"total": 2
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExposureViewService.ListCards(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v1/t1/exposure-view/cards" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/exposure-view/cards")
	}
	if resp.Total != 2 {
		t.Errorf("Total = %d, want 2", resp.Total)
	}
	if len(resp.Cards) != 2 || resp.Cards[0].Id != "card1" {
		t.Errorf("Cards = %v, want 2 elements with first ID card1", resp.Cards)
	}
}

func TestGetCard(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
			"id": "card5",
			"name": "Card E",
			"title": "Exposure trend"
		}`)
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.ExposureViewService.GetCard(context.Background(), "card5")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/api/v1/t1/exposure-view/cards/card5" {
		t.Errorf("path = %q, want %q", gotPath, "/api/v1/t1/exposure-view/cards/card5")
	}
	if resp.Id != "card5" {
		t.Errorf("Id = %q, want card5", resp.Id)
	}
	if resp.Title != "Exposure trend" {
		t.Errorf("Title = %q, want Exposure trend", resp.Title)
	}
}
