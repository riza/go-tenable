// Search inventory assets from Tenable One.
//
// Usage:
//
//	ONE_URL=https://cloud.tenable.com ONE_ACCESS_KEY=xxx ONE_SECRET_KEY=yyy go run .
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/riza/go-tenable/one"
)

func intPtr(i int) *int {
	return &i
}

func main() {
	baseURL := os.Getenv("ONE_URL")
	accessKey := os.Getenv("ONE_ACCESS_KEY")
	secretKey := os.Getenv("ONE_SECRET_KEY")

	if baseURL == "" || accessKey == "" || secretKey == "" {
		log.Fatal("ONE_URL, ONE_ACCESS_KEY, and ONE_SECRET_KEY environment variables are required")
	}

	client := one.NewClient(baseURL,
		one.WithAPIKey(accessKey, secretKey),
	)

	ctx := context.Background()

	// 1. Search Assets
	fmt.Println("--- Searching Inventory Assets ---")
	assetsReq := &one.InventoryAssetsSearchRequest{
		Limit: intPtr(10),
	}

	assetsResp, err := client.InventoryService.SearchAssets(ctx, assetsReq)
	if err != nil {
		log.Fatalf("failed to search assets: %v", err)
	}

	fmt.Printf("Total Assets Found (across platform): %d\n", assetsResp.Total)
	for i, asset := range assetsResp.Assets {
		fmt.Printf("[%d] ID: %s | Name: %s | Type: %s | IPv4: %v\n",
			i+1, asset.Id, asset.Name, asset.Type, asset.IpAddresses)
	}

	// 2. Search Findings
	fmt.Println("\n--- Searching Inventory Findings ---")
	findingsReq := &one.InventoryFindingsSearchRequest{
		Limit: intPtr(10),
		Filters: []one.InventorySearchFilter{
			{Property: "finding_severity", Operator: "=", Value: []string{"HIGH", "CRITICAL"}},
		},
	}

	findingsResp, err := client.InventoryService.SearchFindings(ctx, findingsReq)
	if err != nil {
		log.Fatalf("failed to search findings: %v", err)
	}

	fmt.Printf("Total High/Critical Findings: %d\n", findingsResp.Total)
	for i, finding := range findingsResp.Findings {
		fmt.Printf("[%d] ID: %s | Plugin: %s (%d) | Severity: %s | Asset: %s\n",
			i+1, finding.Id, finding.PluginName, finding.PluginId, finding.Severity, finding.AssetId)
	}
}
