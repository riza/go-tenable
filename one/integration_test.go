//go:build integration

package one

import (
	"context"
	"os"
	"testing"
)

// skipIfMissingEnv skips the test if required environment variables are not set.
func skipIfMissingEnv(t *testing.T, url, accessKey, secretKey string) {
	if url == "" || accessKey == "" || secretKey == "" {
		t.Skip("TENABLE_ONE_* environment variables not set")
	}
}

func TestIntegration_AttackPathSearch(t *testing.T) {
	url := os.Getenv("TENABLE_ONE_URL")
	accessKey := os.Getenv("TENABLE_ONE_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_ONE_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	paths, err := client.AttackPathService.SearchAttackPaths(ctx, &APASearchAttackPathsRequest{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("AttackPathService.SearchAttackPaths() failed: %v", err)
	}

	t.Logf("Found %d attack paths (total: %d)", len(paths.AttackPaths), paths.Total)
}

func TestIntegration_InventorySearchAssets(t *testing.T) {
	url := os.Getenv("TENABLE_ONE_URL")
	accessKey := os.Getenv("TENABLE_ONE_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_ONE_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	limit := 10
	ctx := context.Background()
	result, err := client.InventoryService.SearchAssets(ctx, &InventoryAssetsSearchRequest{
		Limit: &limit,
	})
	if err != nil {
		t.Fatalf("InventoryService.SearchAssets() failed: %v", err)
	}

	t.Logf("Found %d assets (total: %d)", len(result.Assets), result.Total)
}

func TestIntegration_TagsSearch(t *testing.T) {
	url := os.Getenv("TENABLE_ONE_URL")
	accessKey := os.Getenv("TENABLE_ONE_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_ONE_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	limit := 10
	ctx := context.Background()
	result, err := client.TagsService.SearchTags(ctx, &TagsSearchRequest{
		Limit: &limit,
	})
	if err != nil {
		t.Fatalf("TagsService.SearchTags() failed: %v", err)
	}

	t.Logf("Found %d tags (total: %d)", len(result.Tags), result.Total)
}

func TestIntegration_ExportAssets(t *testing.T) {
	url := os.Getenv("TENABLE_ONE_URL")
	accessKey := os.Getenv("TENABLE_ONE_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_ONE_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	export, err := client.ExportService.ExportAssets(ctx, &InventoryExportRequest{})
	if err != nil {
		t.Fatalf("ExportService.ExportAssets() failed: %v", err)
	}

	t.Logf("Export created with ID: %s", export.ExportID)
}

func TestIntegration_ExposureViewCards(t *testing.T) {
	url := os.Getenv("TENABLE_ONE_URL")
	accessKey := os.Getenv("TENABLE_ONE_ACCESS_KEY")
	secretKey := os.Getenv("TENABLE_ONE_SECRET_KEY")
	skipIfMissingEnv(t, url, accessKey, secretKey)

	client := NewClient(url, WithAPIKey(accessKey, secretKey))

	ctx := context.Background()
	cards, err := client.ExposureViewService.ListCards(ctx)
	if err != nil {
		t.Fatalf("ExposureViewService.ListCards() failed: %v", err)
	}

	t.Logf("Found %d exposure cards", len(cards.Cards))
}
