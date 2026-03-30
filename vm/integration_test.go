//go:build integration

package vm

import (
	"context"
	"os"
	"testing"
)

// skipIfMissingEnv skips the test if required environment variables are not set.
func skipIfMissingEnv(t *testing.T, url, accessKey, secretKey string) {
	if url == "" || accessKey == "" || secretKey == "" {
		t.Skip("TENABLE_VM_* environment variables not set")
	}
}

func TestIntegration_AssetsList(t *testing.T) {
	url := os.Getenv("TENABLE_VM_URL")
	accessKey := os.Getenv("TENABLE_VM_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_VM_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	assets, err := client.AssetsService.ListAssets(ctx)
	if err != nil {
		t.Fatalf("AssetsService.ListAssets() failed: %v", err)
	}

	t.Logf("Found %d assets (total: %d)", len(assets.Assets), assets.Total)
}

func TestIntegration_ScansList(t *testing.T) {
	url := os.Getenv("TENABLE_VM_URL")
	accessKey := os.Getenv("TENABLE_VM_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_VM_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	scans, err := client.ScansService.List(ctx)
	if err != nil {
		t.Fatalf("ScansService.List() failed: %v", err)
	}

	t.Logf("Found %d scans", len(scans.Scans))
}

func TestIntegration_FiltersListAssets(t *testing.T) {
	url := os.Getenv("TENABLE_VM_URL")
	accessKey := os.Getenv("TENABLE_VM_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_VM_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	filters, err := client.FiltersService.FiltersAssetsList(ctx)
	if err != nil {
		t.Fatalf("FiltersService.FiltersAssetsList() failed: %v", err)
	}

	t.Logf("Found %d asset filters", len(filters.Filters))
}

func TestIntegration_PoliciesList(t *testing.T) {
	url := os.Getenv("TENABLE_VM_URL")
	accessKey := os.Getenv("TENABLE_VM_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_VM_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	policies, err := client.PoliciesService.PoliciesList(ctx)
	if err != nil {
		t.Fatalf("PoliciesService.PoliciesList() failed: %v", err)
	}

	t.Logf("Found %d policies", len(policies))
}

func TestIntegration_WorkbenchesVulnerabilities(t *testing.T) {
	url := os.Getenv("TENABLE_VM_URL")
	accessKey := os.Getenv("TENABLE_VM_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_VM_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	vulns, err := client.WorkbenchesService.WorkbenchesVulnerabilities(ctx)
	if err != nil {
		t.Fatalf("WorkbenchesService.WorkbenchesVulnerabilities() failed: %v", err)
	}

	t.Logf("Found %d vulnerabilities in workbench", len(vulns.Vulnerabilities))
}
