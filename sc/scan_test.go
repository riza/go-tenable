package sc

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScanList(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type: "regular",
			Response: json.RawMessage(`{
				"usable": [
					{"id": "1", "name": "Scan A", "status": "0"},
					{"id": "2", "name": "Scan B", "status": "1"}
				],
				"manageable": [
					{"id": "1", "name": "Scan A", "status": "0"}
				]
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	result, err := c.Scan.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/rest/scan" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/scan")
	}
	if len(result.Usable) != 2 {
		t.Errorf("len(Usable) = %d, want 2", len(result.Usable))
	}
	if len(result.Manageable) != 1 {
		t.Errorf("len(Manageable) = %d, want 1", len(result.Manageable))
	}
	if result.Usable[0].ID != "1" {
		t.Errorf("Usable[0].ID = %q, want %q", result.Usable[0].ID, "1")
	}
	if result.Usable[0].Name != "Scan A" {
		t.Errorf("Usable[0].Name = %q, want %q", result.Usable[0].Name, "Scan A")
	}
	if result.Usable[1].ID != "2" {
		t.Errorf("Usable[1].ID = %q, want %q", result.Usable[1].ID, "2")
	}
	if result.Manageable[0].ID != "1" {
		t.Errorf("Manageable[0].ID = %q, want %q", result.Manageable[0].ID, "1")
	}
}

func TestScanGet(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type: "regular",
			Response: json.RawMessage(`{
				"id": "4",
				"name": "POSTtest",
				"description": "This is a test for POST",
				"ipList": "100.100.100.100",
				"type": "policy",
				"status": "0",
				"dhcpTracking": "false",
				"emailOnLaunch": "false",
				"emailOnFinish": "false",
				"timeoutAction": "import",
				"scanningVirtualHosts": "false",
				"rolloverType": "template",
				"createdTime": "1406815242",
				"modifiedTime": "1406815242",
				"uuid": "29F2B9E1-ADE9-4550-B63C-CEA1423E52FC",
				"policy": {"id": "1000002"},
				"repository": {"id": "2"},
				"owner": {"id": "1"},
				"creator": {"id": "1"},
				"ownerGroup": {"id": "0"},
				"schedule": {
					"id": "17",
					"type": "dependent",
					"start": "",
					"repeatRule": "",
					"enabled": "true",
					"nextRun": 0,
					"objectType": "scan",
					"dependentID": "14"
				},
				"reports": [],
				"assets": [],
				"credentials": [],
				"numDependents": "0",
				"canUse": "true",
				"canManage": "true"
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	scan, err := c.Scan.Get(context.Background(), "4")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/rest/scan/4" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/scan/4")
	}
	if scan.ID != "4" {
		t.Errorf("ID = %q, want %q", scan.ID, "4")
	}
	if scan.Name != "POSTtest" {
		t.Errorf("Name = %q, want %q", scan.Name, "POSTtest")
	}
	if scan.Description != "This is a test for POST" {
		t.Errorf("Description = %q, want %q", scan.Description, "This is a test for POST")
	}
	if scan.IPList != "100.100.100.100" {
		t.Errorf("IPList = %q, want %q", scan.IPList, "100.100.100.100")
	}
	if scan.Type != "policy" {
		t.Errorf("Type = %q, want %q", scan.Type, "policy")
	}
	if scan.Policy == nil || scan.Policy.ID != "1000002" {
		t.Errorf("Policy.ID = %v, want %q", scan.Policy, "1000002")
	}
	if scan.Repository == nil || scan.Repository.ID != "2" {
		t.Errorf("Repository.ID = %v, want %q", scan.Repository, "2")
	}
	if scan.Schedule == nil {
		t.Fatal("Schedule should not be nil")
	}
	if scan.Schedule.Type != "dependent" {
		t.Errorf("Schedule.Type = %q, want %q", scan.Schedule.Type, "dependent")
	}
	if scan.Schedule.DependentID != "14" {
		t.Errorf("Schedule.DependentID = %q, want %q", scan.Schedule.DependentID, "14")
	}
	if scan.CanUse != "true" {
		t.Errorf("CanUse = %q, want %q", scan.CanUse, "true")
	}
}

func TestScanCreate(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody ScanCreateInput
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
			Type: "regular",
			Response: json.RawMessage(`{
				"id": "4",
				"name": "POSTtest",
				"description": "This is a test for POST",
				"ipList": "100.100.100.100",
				"type": "policy",
				"status": "0",
				"dhcpTracking": "false",
				"emailOnLaunch": "false",
				"emailOnFinish": "false",
				"timeoutAction": "import",
				"scanningVirtualHosts": "false",
				"rolloverType": "template",
				"createdTime": "1406815242",
				"modifiedTime": "1406815242",
				"uuid": "29F2B9E1-ADE9-4550-B63C-CEA1423E52FC",
				"policy": {"id": "1000002"},
				"repository": {"id": "2"},
				"owner": {"id": "1"},
				"creator": {"id": "1"},
				"ownerGroup": {"id": "0"},
				"schedule": {
					"id": "17",
					"type": "dependent",
					"start": "",
					"repeatRule": "",
					"enabled": "true",
					"nextRun": 0,
					"objectType": "scan",
					"dependentID": "14"
				},
				"reports": [],
				"assets": [],
				"credentials": [],
				"numDependents": "0",
				"canUse": "true",
				"canManage": "true"
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	input := &ScanCreateInput{
		Name:        "POSTtest",
		Type:        "policy",
		Description: "This is a test for POST",
		Repository:  &IDRef{ID: "2"},
		IPList:      "100.100.100.100",
		Policy:      &IDRef{ID: "1000002"},
		Schedule: &ScanSchedule{
			Type: "dependent",
		},
	}
	scan, err := c.Scan.Create(context.Background(), input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/rest/scan" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/scan")
	}

	// Verify request body
	if gotBody.Name != "POSTtest" {
		t.Errorf("request Name = %q, want %q", gotBody.Name, "POSTtest")
	}
	if gotBody.Type != "policy" {
		t.Errorf("request Type = %q, want %q", gotBody.Type, "policy")
	}
	if gotBody.Repository == nil || gotBody.Repository.ID != "2" {
		t.Errorf("request Repository.ID = %v, want %q", gotBody.Repository, "2")
	}
	if gotBody.IPList != "100.100.100.100" {
		t.Errorf("request IPList = %q, want %q", gotBody.IPList, "100.100.100.100")
	}
	if gotBody.Policy == nil || gotBody.Policy.ID != "1000002" {
		t.Errorf("request Policy.ID = %v, want %q", gotBody.Policy, "1000002")
	}
	if gotBody.Schedule == nil || gotBody.Schedule.Type != "dependent" {
		t.Errorf("request Schedule.Type = %v, want %q", gotBody.Schedule, "dependent")
	}

	// Verify response parsing
	if scan.ID != "4" {
		t.Errorf("ID = %q, want %q", scan.ID, "4")
	}
	if scan.Name != "POSTtest" {
		t.Errorf("Name = %q, want %q", scan.Name, "POSTtest")
	}
	if scan.Type != "policy" {
		t.Errorf("Type = %q, want %q", scan.Type, "policy")
	}
}

func TestScanUpdate(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody map[string]interface{}
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
			Type: "regular",
			Response: json.RawMessage(`{
				"id": "4",
				"name": "UpdatedScan",
				"description": "Updated description",
				"ipList": "10.0.0.1",
				"type": "policy",
				"status": "0",
				"uuid": "29F2B9E1-ADE9-4550-B63C-CEA1423E52FC",
				"repository": {"id": "2"},
				"canUse": "true",
				"canManage": "true"
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	input := &ScanUpdateInput{
		Name:        "UpdatedScan",
		Description: "Updated description",
		Repository:  &IDRef{ID: "2"},
		IPList:      "10.0.0.1",
	}
	scan, err := c.Scan.Update(context.Background(), "4", input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodPatch {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPatch)
	}
	if gotPath != "/rest/scan/4" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/scan/4")
	}

	// Verify request body
	if gotBody["name"] != "UpdatedScan" {
		t.Errorf("request name = %v, want %q", gotBody["name"], "UpdatedScan")
	}
	if gotBody["description"] != "Updated description" {
		t.Errorf("request description = %v, want %q", gotBody["description"], "Updated description")
	}
	if gotBody["ipList"] != "10.0.0.1" {
		t.Errorf("request ipList = %v, want %q", gotBody["ipList"], "10.0.0.1")
	}

	// Verify response parsing
	if scan.ID != "4" {
		t.Errorf("ID = %q, want %q", scan.ID, "4")
	}
	if scan.Name != "UpdatedScan" {
		t.Errorf("Name = %q, want %q", scan.Name, "UpdatedScan")
	}
	if scan.Description != "Updated description" {
		t.Errorf("Description = %q, want %q", scan.Description, "Updated description")
	}
}

func TestScanDelete(t *testing.T) {
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
	err := c.Scan.Delete(context.Background(), "4")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodDelete {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodDelete)
	}
	if gotPath != "/rest/scan/4" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/scan/4")
	}
}

func TestScanLaunch(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody map[string]interface{}
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
			Type: "regular",
			Response: json.RawMessage(`{
				"scanID": "2",
				"scanResult": {
					"id": "3",
					"name": "test",
					"status": "Queued",
					"jobID": "143301"
				}
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	input := &ScanLaunchInput{
		DiagnosticTarget:   "10.0.0.1",
		DiagnosticPassword: "secret123",
	}
	result, err := c.Scan.Launch(context.Background(), "2", input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/rest/scan/2/launch" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/scan/2/launch")
	}

	// Verify request body
	if gotBody["diagnosticTarget"] != "10.0.0.1" {
		t.Errorf("request diagnosticTarget = %v, want %q", gotBody["diagnosticTarget"], "10.0.0.1")
	}
	if gotBody["diagnosticPassword"] != "secret123" {
		t.Errorf("request diagnosticPassword = %v, want %q", gotBody["diagnosticPassword"], "secret123")
	}

	// Verify response parsing
	if result.ScanID != "2" {
		t.Errorf("ScanID = %q, want %q", result.ScanID, "2")
	}
	if result.ScanResult.ID != "3" {
		t.Errorf("ScanResult.ID = %q, want %q", result.ScanResult.ID, "3")
	}
	if result.ScanResult.Name != "test" {
		t.Errorf("ScanResult.Name = %q, want %q", result.ScanResult.Name, "test")
	}
	if result.ScanResult.Status != "Queued" {
		t.Errorf("ScanResult.Status = %q, want %q", result.ScanResult.Status, "Queued")
	}
	if result.ScanResult.JobID != "143301" {
		t.Errorf("ScanResult.JobID = %q, want %q", result.ScanResult.JobID, "143301")
	}
}

func TestScanLaunchNilInput(t *testing.T) {
	var gotMethod, gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse{
			Type: "regular",
			Response: json.RawMessage(`{
				"scanID": "2",
				"scanResult": {
					"id": "3",
					"name": "test",
					"status": "Queued",
					"jobID": "143301"
				}
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	result, err := c.Scan.Launch(context.Background(), "2", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/rest/scan/2/launch" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/scan/2/launch")
	}
	if result.ScanID != "2" {
		t.Errorf("ScanID = %q, want %q", result.ScanID, "2")
	}
}

func TestScanCopy(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody map[string]interface{}
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
			Type: "regular",
			Response: json.RawMessage(`{
				"scan": {
					"id": "5",
					"name": "CopiedScan",
					"description": "This is a test for POST",
					"ipList": "100.100.100.100",
					"type": "policy",
					"status": "0",
					"dhcpTracking": "false",
					"emailOnLaunch": "false",
					"emailOnFinish": "false",
					"timeoutAction": "import",
					"scanningVirtualHosts": "false",
					"rolloverType": "template",
					"createdTime": "1406815242",
					"modifiedTime": "1406815242",
					"uuid": "29F2B9E1-ADE9-4550-B63C-CEA1423E52FC",
					"policy": {"id": "1000002"},
					"repository": {"id": "2"},
					"owner": {"id": "1"},
					"creator": {"id": "1"},
					"ownerGroup": {"id": "0"},
					"numDependents": "0",
					"canUse": "true",
					"canManage": "true"
				}
			}`),
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL)
	input := &ScanCopyInput{
		Name:       "CopiedScan",
		TargetUser: &IDRef{ID: "1"},
	}
	scan, err := c.Scan.Copy(context.Background(), "4", input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/rest/scan/4/copy" {
		t.Errorf("path = %q, want %q", gotPath, "/rest/scan/4/copy")
	}

	// Verify request body
	if gotBody["name"] != "CopiedScan" {
		t.Errorf("request name = %v, want %q", gotBody["name"], "CopiedScan")
	}
	targetUser, ok := gotBody["targetUser"].(map[string]interface{})
	if !ok {
		t.Fatalf("request targetUser is not an object: %v", gotBody["targetUser"])
	}
	if targetUser["id"] != "1" {
		t.Errorf("request targetUser.id = %v, want %q", targetUser["id"], "1")
	}

	// Verify response parsing
	if scan.ID != "5" {
		t.Errorf("ID = %q, want %q", scan.ID, "5")
	}
	if scan.Name != "CopiedScan" {
		t.Errorf("Name = %q, want %q", scan.Name, "CopiedScan")
	}
	if scan.Type != "policy" {
		t.Errorf("Type = %q, want %q", scan.Type, "policy")
	}
	if scan.Repository == nil || scan.Repository.ID != "2" {
		t.Errorf("Repository.ID = %v, want %q", scan.Repository, "2")
	}
}
