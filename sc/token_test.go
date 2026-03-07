package sc

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTokenCreate(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody TokenCreateInput

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path

		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		if err := json.Unmarshal(data, &gotBody); err != nil {
			t.Fatalf("failed to unmarshal request body: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type:     "regular",
			Response: json.RawMessage(`{"token":123456789}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.Token.Create(&TokenCreateInput{
		Username: "admin",
		Password: "secret",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify HTTP method.
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}

	// Verify path.
	if gotPath != "/rest/token" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/token")
	}

	// Verify request body was marshaled correctly.
	if gotBody.Username != "admin" {
		t.Errorf("Username = %q, want %q", gotBody.Username, "admin")
	}
	if gotBody.Password != "secret" {
		t.Errorf("Password = %q, want %q", gotBody.Password, "secret")
	}

	// Verify response was parsed correctly.
	if resp.Token != 123456789 {
		t.Errorf("Token = %d, want %d", resp.Token, 123456789)
	}
}

func TestTokenCreateWithReleaseSession(t *testing.T) {
	var gotBody TokenCreateInput

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		if err := json.Unmarshal(data, &gotBody); err != nil {
			t.Fatalf("failed to unmarshal request body: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type:     "regular",
			Response: json.RawMessage(`{"token":987654321}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	resp, err := c.Token.Create(&TokenCreateInput{
		Username:       "admin",
		Password:       "secret",
		ReleaseSession: "true",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotBody.ReleaseSession != "true" {
		t.Errorf("ReleaseSession = %q, want %q", gotBody.ReleaseSession, "true")
	}
	if resp.Token != 987654321 {
		t.Errorf("Token = %d, want %d", resp.Token, 987654321)
	}
}

func TestTokenDelete(t *testing.T) {
	var gotMethod, gotPath string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type:     "regular",
			Response: json.RawMessage(`""`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.Token.Delete()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify HTTP method.
	if gotMethod != http.MethodDelete {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodDelete)
	}

	// Verify path.
	if gotPath != "/rest/token" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/token")
	}
}

func TestTokenCreateAPIError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(apiResponse{
			Type:      "error",
			ErrorCode: 141,
			ErrorMsg:  "Invalid login credentials.",
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	_, err := c.Token.Create(&TokenCreateInput{
		Username: "admin",
		Password: "wrong",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestTokenDeleteAPIError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(apiResponse{
			Type:      "error",
			ErrorCode: 143,
			ErrorMsg:  "Invalid token.",
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	err := c.Token.Delete()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
