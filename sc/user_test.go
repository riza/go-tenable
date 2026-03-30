package sc

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(apiResponse{
			Type: "regular",
			Response: json.RawMessage(`{
				"totalRecords": 2,
				"returnedRecords": 2,
				"startOffset": 0,
				"endOffset": 2,
				"results": [
					{"id": "1", "username": "alice", "status": "0", "email": "a@example.com"},
					{"id": "2", "username": "bob", "status": "0", "email": "b@example.com"}
				]
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	result, err := c.User.List(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/rest/user" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/user")
	}
	if result.TotalRecords != 2 {
		t.Errorf("TotalRecords = %d, want 2", result.TotalRecords)
	}
	if len(result.Results) != 2 {
		t.Errorf("len(Results) = %d, want 2", len(result.Results))
	}
	if result.Results[0].ID != "1" || result.Results[0].Username != "alice" {
		t.Errorf("Results[0] = %+v, want id=1 username=alice", result.Results[0])
	}
	if result.Results[1].ID != "2" || result.Results[1].Username != "bob" {
		t.Errorf("Results[1] = %+v, want id=2 username=bob", result.Results[1])
	}
}

func TestUserGet(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(apiResponse{
			Type: "regular",
			Response: json.RawMessage(`{
				"id": "42",
				"uuid": "550E8400-E29B-41D4-A716-446655440000",
				"username": "jdoe",
				"status": "0",
				"firstname": "Jane",
				"lastname": "Doe",
				"email": "jane@example.com",
				"authType": "ldap",
				"role": {"id": "3"},
				"canUse": "true",
				"canManage": "false"
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	user, err := c.User.Get(context.Background(), "42")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/rest/user/42" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/user/42")
	}
	if user.ID != "42" {
		t.Errorf("ID = %q, want %q", user.ID, "42")
	}
	if user.Username != "jdoe" {
		t.Errorf("Username = %q, want %q", user.Username, "jdoe")
	}
	if user.Firstname != "Jane" || user.Lastname != "Doe" {
		t.Errorf("name = %q %q, want Jane Doe", user.Firstname, user.Lastname)
	}
	if user.Role == nil || user.Role.ID != "3" {
		t.Errorf("Role = %v, want id=3", user.Role)
	}
	if user.CanUse != "true" {
		t.Errorf("CanUse = %q, want %q", user.CanUse, "true")
	}
}

func TestUserCreate(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody UserCreateInput
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
		_ = json.NewEncoder(w).Encode(apiResponse{
			Type: "regular",
			Response: json.RawMessage(`{
				"id": "99",
				"username": "newuser",
				"status": "0",
				"firstname": "New",
				"lastname": "User",
				"email": "new@example.com",
				"authType": "password",
				"role": {"id": "2"},
				"canUse": "true",
				"canManage": "false"
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	input := &UserCreateInput{
		Username:  "newuser",
		RoleID:    "2",
		AuthType:  "password",
		Password:  "secret",
		Firstname: "New",
		Lastname:  "User",
		Email:     "new@example.com",
	}
	user, err := c.User.Create(context.Background(), input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/rest/user" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/user")
	}

	if gotBody.Username != "newuser" {
		t.Errorf("request Username = %q, want %q", gotBody.Username, "newuser")
	}
	if gotBody.RoleID != "2" {
		t.Errorf("request RoleID = %q, want %q", gotBody.RoleID, "2")
	}
	if gotBody.AuthType != "password" {
		t.Errorf("request AuthType = %q, want %q", gotBody.AuthType, "password")
	}
	if gotBody.Password != "secret" {
		t.Errorf("request Password = %q, want %q", gotBody.Password, "secret")
	}
	if gotBody.Email != "new@example.com" {
		t.Errorf("request Email = %q, want %q", gotBody.Email, "new@example.com")
	}

	if user.ID != "99" {
		t.Errorf("ID = %q, want %q", user.ID, "99")
	}
	if user.Username != "newuser" {
		t.Errorf("Username = %q, want %q", user.Username, "newuser")
	}
	if user.Role == nil || user.Role.ID != "2" {
		t.Errorf("Role.ID = %v, want %q", user.Role, "2")
	}
}

func TestUserUpdate(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody UserUpdateInput
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
		_ = json.NewEncoder(w).Encode(apiResponse{
			Type: "regular",
			Response: json.RawMessage(`{
				"id": "7",
				"username": "updated",
				"status": "0",
				"email": "upd@example.com",
				"role": {"id": "4"},
				"canUse": "true",
				"canManage": "true"
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	input := &UserUpdateInput{
		Username: "updated",
		RoleID:   "4",
		AuthType: "password",
		Email:    "upd@example.com",
	}
	user, err := c.User.Update(context.Background(), "7", input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodPatch {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPatch)
	}
	if gotPath != "/rest/user/7" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/user/7")
	}

	if gotBody.Username != "updated" {
		t.Errorf("request Username = %q, want %q", gotBody.Username, "updated")
	}
	if gotBody.RoleID != "4" {
		t.Errorf("request RoleID = %q, want %q", gotBody.RoleID, "4")
	}
	if gotBody.Email != "upd@example.com" {
		t.Errorf("request Email = %q, want %q", gotBody.Email, "upd@example.com")
	}

	if user.ID != "7" {
		t.Errorf("ID = %q, want %q", user.ID, "7")
	}
	if user.Username != "updated" {
		t.Errorf("Username = %q, want %q", user.Username, "updated")
	}
	if user.Email != "upd@example.com" {
		t.Errorf("Email = %q, want %q", user.Email, "upd@example.com")
	}
}

func TestUserDelete(t *testing.T) {
	t.Run("without migration user", func(t *testing.T) {
		var gotMethod, gotPath string
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotMethod = r.Method
			gotPath = r.URL.Path
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(apiResponse{
				Type:     "regular",
				Response: json.RawMessage(`""`),
			})
		}))
		defer ts.Close()

		c := NewClient(ts.URL)
		err := c.User.Delete(context.Background(), "5", nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if gotMethod != http.MethodDelete {
			t.Errorf("method = %q, want %q", gotMethod, http.MethodDelete)
		}
		if gotPath != "/rest/user/5" {
			t.Errorf("path = %q, want %q", gotPath, "/rest/user/5")
		}
	})

	t.Run("with migration user", func(t *testing.T) {
		var gotMethod, gotPath, gotQuery string
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotMethod = r.Method
			gotPath = r.URL.Path
			gotQuery = r.URL.RawQuery
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(apiResponse{
				Type:     "regular",
				Response: json.RawMessage(`""`),
			})
		}))
		defer ts.Close()

		c := NewClient(ts.URL)
		err := c.User.Delete(context.Background(), "5", &UserDeleteInput{MigrateUserID: "88"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if gotMethod != http.MethodDelete {
			t.Errorf("method = %q, want %q", gotMethod, http.MethodDelete)
		}
		if gotPath != "/rest/user/5" {
			t.Errorf("path = %q, want %q", gotPath, "/rest/user/5")
		}
		if gotQuery != "migrateUserID=88" {
			t.Errorf("RawQuery = %q, want %q", gotQuery, "migrateUserID=88")
		}
	})
}

func TestUserListAPIError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(apiResponse{
			Type:      "error",
			ErrorCode: 401,
			ErrorMsg:  "Insufficient privileges",
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	_, err := c.User.List(context.Background())
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != http.StatusForbidden {
		t.Errorf("StatusCode = %d, want %d", apiErr.StatusCode, http.StatusForbidden)
	}
	if apiErr.ErrorCode != 401 {
		t.Errorf("ErrorCode = %d, want %d", apiErr.ErrorCode, 401)
	}
	if apiErr.ErrorMsg != "Insufficient privileges" {
		t.Errorf("ErrorMsg = %q, want %q", apiErr.ErrorMsg, "Insufficient privileges")
	}
}
