// Fetch all assets belonging to a specific Tenable tag using the VM Export API.
//
// The program triggers an async asset export filtered by the given tag, polls
// until the export finishes, downloads every chunk, and prints a summary line
// for each asset.
//
// Usage:
//
//	VM_URL=https://cloud.tenable.com \
//	VM_ACCESS_KEY=xxx \
//	VM_SECRET_KEY=yyy \
//	VM_TAG_CATEGORY=Environment \
//	VM_TAG_VALUE=Production \
//	go run .
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/riza/go-tenable/vm"
)

func main() {
	baseURL := os.Getenv("VM_URL")
	accessKey := os.Getenv("VM_ACCESS_KEY")
	secretKey := os.Getenv("VM_SECRET_KEY")
	tagCategory := os.Getenv("VM_TAG_CATEGORY")
	tagValue := os.Getenv("VM_TAG_VALUE")

	if baseURL == "" || accessKey == "" || secretKey == "" {
		log.Fatal("VM_URL, VM_ACCESS_KEY, and VM_SECRET_KEY environment variables are required")
	}
	if tagCategory == "" || tagValue == "" {
		log.Fatal("VM_TAG_CATEGORY and VM_TAG_VALUE environment variables are required")
	}

	client := vm.NewClient(baseURL,
		vm.WithAPIKey(accessKey, secretKey),
	)

	ctx := context.Background()

	// Step 1: Start export filtered by tag
	fmt.Printf("Starting asset export — tag: %s = %s\n", tagCategory, tagValue)
	export, err := client.ExportsAssetsService.ExportAssetsByTag(ctx, tagCategory, []string{tagValue}, 1000)
	if err != nil {
		log.Fatalf("failed to start asset export: %v", err)
	}
	fmt.Printf("Export started — UUID: %s\n", export.ExportUuid)

	// Step 2: Poll until FINISHED
	var chunks []int
	for {
		status, err := client.ExportsAssetsService.ExportsAssetsExportStatus(ctx, export.ExportUuid)
		if err != nil {
			log.Fatalf("failed to get export status: %v", err)
		}

		fmt.Printf("  Status: %-12s chunks ready: %v\n", status.Status, status.ChunksAvailable)

		switch status.Status {
		case "FINISHED":
			chunks = status.ChunksAvailable
		case "ERROR":
			log.Fatal("export failed with ERROR status")
		default:
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	if len(chunks) == 0 {
		fmt.Println("Export finished but no chunks available — no assets matched the tag filter.")
		return
	}

	// Step 3: Download and print each chunk
	var total int
	for _, chunkID := range chunks {
		fmt.Printf("\nDownloading chunk %d…\n", chunkID)

		data, err := client.ExportsAssetsService.ExportsAssetsDownloadChunk(ctx, export.ExportUuid, fmt.Sprintf("%d", chunkID))
		if err != nil {
			log.Printf("failed to download chunk %d: %v", chunkID, err)
			continue
		}

		var assets []vm.ExportsAssetsChunkAsset
		if err := json.Unmarshal(data, &assets); err != nil {
			log.Printf("failed to decode chunk %d: %v", chunkID, err)
			continue
		}

		for _, a := range assets {
			fmt.Printf("  %-36s  %-40s  os=%-25s  lastSeen=%s\n",
				a.Id,
				truncate(primaryIdentifier(a), 40),
				truncate(first(a.OperatingSystems), 25),
				a.LastSeen,
			)
		}
		total += len(assets)
	}

	fmt.Printf("\nTotal assets with tag %s=%s: %d\n", tagCategory, tagValue, total)
}

// primaryIdentifier returns the most human-readable identifier available.
func primaryIdentifier(a vm.ExportsAssetsChunkAsset) string {
	if len(a.Fqdns) > 0 {
		return a.Fqdns[0]
	}
	if len(a.Hostnames) > 0 {
		return a.Hostnames[0]
	}
	if a.AwsEc2Name != "" {
		return a.AwsEc2Name
	}
	if len(a.Ipv4s) > 0 {
		return strings.Join(a.Ipv4s, ", ")
	}
	return "unknown"
}

func first(ss []string) string {
	if len(ss) > 0 {
		return ss[0]
	}
	return ""
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
