//go:build integration

package sc

import (
	"context"
	"os"
	"testing"
)

// skipIfMissingEnv skips the test if required environment variables are not set.
func skipIfMissingEnv(t *testing.T, url, accessKey, secretKey string) {
	if url == "" || accessKey == "" || secretKey == "" {
		t.Skip("TENABLE_SC_* environment variables not set")
	}
}

func TestIntegration_ScanList(t *testing.T) {
	url := os.Getenv("TENABLE_SC_URL")
	accessKey := os.Getenv("TENABLE_SC_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_SC_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	scans, err := client.Scan.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("Scan.List() failed: %v", err)
	}

	t.Logf("Found %d usable scans, %d manageable scans",
		len(scans.Usable), len(scans.Manageable))
}

func TestIntegration_ScanListWithFields(t *testing.T) {
	url := os.Getenv("TENABLE_SC_URL")
	accessKey := os.Getenv("TENABLE_SC_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_SC_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	scans, err := client.Scan.List(context.Background(), &ScanListOptions{
		Fields: []string{"id", "name", "status"},
	})
	if err != nil {
		t.Fatalf("Scan.List() with fields failed: %v", err)
	}

	t.Logf("Found %d usable scans with fields filter", len(scans.Usable))
}

func TestIntegration_ScanGet(t *testing.T) {
	url := os.Getenv("TENABLE_SC_URL")
	accessKey := os.Getenv("TENABLE_SC_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_SC_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	// First, get the list of scans to find a valid ID
	scans, err := client.Scan.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("Scan.List() failed: %v", err)
	}

	if len(scans.Usable) == 0 {
		t.Skip("No scans available to test Scan.Get()")
	}

	scanID := scans.Usable[0].ID
	scan, err := client.Scan.Get(context.Background(), scanID)
	if err != nil {
		t.Fatalf("Scan.Get(%s) failed: %v", scanID, err)
	}

	if scan.ID != scanID {
		t.Errorf("Scan.ID = %q, want %q", scan.ID, scanID)
	}

	t.Logf("Successfully retrieved scan: %s (%s)", scan.Name, scan.ID)
}

func TestIntegration_UserList(t *testing.T) {
	url := os.Getenv("TENABLE_SC_URL")
	accessKey := os.Getenv("TENABLE_SC_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_SC_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	users, err := client.User.List(context.Background())
	if err != nil {
		t.Fatalf("User.List() failed: %v", err)
	}

	t.Logf("Found %d users", len(users.Results))
}

func TestIntegration_RepositoryList(t *testing.T) {
	url := os.Getenv("TENABLE_SC_URL")
	accessKey := os.Getenv("TENABLE_SC_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_SC_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	repos, err := client.Repository.List(context.Background())
	if err != nil {
		t.Fatalf("Repository.List() failed: %v", err)
	}

	t.Logf("Found %d repositories", len(repos))
}

func TestIntegration_ScannerList(t *testing.T) {
	url := os.Getenv("TENABLE_SC_URL")
	accessKey := os.Getenv("TENABLE_SC_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_SC_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	scanners, err := client.Scanner.List(context.Background())
	if err != nil {
		t.Fatalf("Scanner.List() failed: %v", err)
	}

	t.Logf("Found %d usable scanners, %d manageable scanners",
		len(scanners.Usable), len(scanners.Manageable))
}

func TestIntegration_Status(t *testing.T) {
	url := os.Getenv("TENABLE_SC_URL")
	accessKey := os.Getenv("TENABLE_SC_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_SC_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	status, err := client.Status.List(context.Background())
	if err != nil {
		t.Fatalf("Status.List() failed: %v", err)
	}

	t.Logf("Found %d usable statuses, %d manageable statuses",
		len(status.Usable), len(status.Manageable))
}
