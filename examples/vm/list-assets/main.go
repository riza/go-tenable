// List all assets from Tenable Vulnerability Management.
//
// Usage:
//
//	VM_URL=https://cloud.tenable.com VM_ACCESS_KEY=xxx VM_SECRET_KEY=yyy go run .
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/riza/go-tenable/vm"
)

func main() {
	baseURL := os.Getenv("VM_URL")
	accessKey := os.Getenv("VM_ACCESS_KEY")
	secretKey := os.Getenv("VM_SECRET_KEY")

	if baseURL == "" || accessKey == "" || secretKey == "" {
		log.Fatal("VM_URL, VM_ACCESS_KEY, and VM_SECRET_KEY environment variables are required")
	}

	client := vm.NewClient(baseURL,
		vm.WithAPIKey(accessKey, secretKey),
		vm.WithInsecureSkipVerify(), // if needed for intercepting proxies
	)

	// Fetch assets using the generated AssetsService
	result, err := client.AssetsService.ListAssets(context.Background())
	if err != nil {
		log.Fatalf("failed to list assets: %v", err)
	}

	fmt.Printf("Found %d asset(s)\n\n", result.Total)

	for _, a := range result.Assets {
		fmt.Printf("%-36s %-40s hasAgent=%-5t lastSeen=%s\n",
			a.Id,
			truncate(getPrimaryHostname(a), 40),
			a.HasAgent,
			a.LastSeen,
		)
	}
}

func getPrimaryHostname(asset vm.AssetsServiceListAssetsResponseAssetsItem) string {
	if len(asset.Fqdn) > 0 {
		return asset.Fqdn[0]
	}
	if len(asset.Hostname) > 0 {
		return asset.Hostname[0]
	}
	if len(asset.Ipv4) > 0 {
		return asset.Ipv4[0]
	}
	if len(asset.Ipv6) > 0 {
		return asset.Ipv6[0]
	}
	return "Unknown"
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
