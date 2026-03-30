//go:build integration

package platform

import (
	"context"
	"os"
	"testing"
)

// skipIfMissingEnv skips the test if required environment variables are not set.
func skipIfMissingEnv(t *testing.T, url, accessKey, secretKey string) {
	if url == "" || accessKey == "" || secretKey == "" {
		t.Skip("TENABLE_PLATFORM_* environment variables not set")
	}
}

func TestIntegration_UsersList(t *testing.T) {
	url := os.Getenv("TENABLE_PLATFORM_URL")
	accessKey := os.Getenv("TENABLE_PLATFORM_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_PLATFORM_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	users, err := client.UsersService.ListUsers(ctx)
	if err != nil {
		t.Fatalf("UsersService.ListUsers() failed: %v", err)
	}

	t.Logf("Found %d users", len(users.Users))
}

func TestIntegration_GroupsList(t *testing.T) {
	url := os.Getenv("TENABLE_PLATFORM_URL")
	accessKey := os.Getenv("TENABLE_PLATFORM_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_PLATFORM_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	groups, err := client.GroupsService.ListGroups(ctx)
	if err != nil {
		t.Fatalf("GroupsService.ListGroups() failed: %v", err)
	}

	t.Logf("Found %d groups", len(groups.Groups))
}

func TestIntegration_AgentsList(t *testing.T) {
	url := os.Getenv("TENABLE_PLATFORM_URL")
	accessKey := os.Getenv("TENABLE_PLATFORM_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_PLATFORM_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	agents, err := client.AgentsService.ListAgents(ctx)
	if err != nil {
		t.Fatalf("AgentsService.ListAgents() failed: %v", err)
	}

	t.Logf("Found %d agents", len(agents.Agents))
}

func TestIntegration_ScannersList(t *testing.T) {
	url := os.Getenv("TENABLE_PLATFORM_URL")
	accessKey := os.Getenv("TENABLE_PLATFORM_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_PLATFORM_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	scanners, err := client.ScannersService.ListScanners(ctx)
	if err != nil {
		t.Fatalf("ScannersService.ListScanners() failed: %v", err)
	}

	t.Logf("Found %d scanners", len(scanners.Scanners))
}
