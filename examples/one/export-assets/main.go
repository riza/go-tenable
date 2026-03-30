// Export inventory assets from Tenable One.
// This example demonstrates to trigger an asynchronous export task, wait for it to process, and then download chunks.
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
	"time"

	"github.com/riza/go-tenable/one"
)

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

	fmt.Println("1. Requesting Inventory Export (Assets)...")
	exportReq := &one.InventoryExportRequest{
		Limit: 100, // Number of items per chunk
	}

	// Step 1: Trigger Export
	exportResp, err := client.ExportService.ExportAssets(ctx, exportReq)
	if err != nil {
		log.Fatalf("failed to request asset export: %v", err)
	}

	fmt.Printf("Export Task Created: UUID = %s | Format = %s\n", exportResp.ExportID, exportResp.Format)

	// Step 2: Poll Export Status
	for {
		fmt.Printf("\nPolling Export Task [%s] Status...\n", exportResp.ExportID)
		status, err := client.ExportService.GetAssetsExportStatus(ctx, exportResp.ExportID)
		if err != nil {
			log.Fatalf("failed to retrieve export status: %v", err)
		}

		fmt.Printf("Status: %s\n", status.Status)
		if status.Status == "FINISHED" {
			fmt.Printf("Export Completed! Available Chunks = %v\n", status.Chunks)

			// Step 3: Fetch Data Chunks
			for _, chunkID := range status.Chunks {
				fmt.Printf("\nDownloading Chunk #%d\n", chunkID)

				data, err := client.ExportService.DownloadExportChunk(ctx, exportResp.ExportID, chunkID)
				if err != nil {
					log.Printf("Failed to download chunk %d: %v", chunkID, err)
					continue
				}

				// The JSON array starts directly in the block (newline delimited objects for export)
				fmt.Printf("Chunk Size: %d bytes\n", len(data))
				if len(data) > 0 {
					// Just print a little preview instead of terminal flood
					preview := string(data)
					if len(preview) > 1000 {
						preview = preview[:1000] + "...\n(TRUNCATED)"
					}
					fmt.Printf("Data Preview:\n%s\n", preview)
				}
			}
			break
		} else if status.Status == "ERROR" {
			log.Fatalf("Export Failed!")
		}

		fmt.Printf("Total Objects Exported so Far: %d\n", status.TotalObjects)
		fmt.Println("Sleeping for 5 seconds before checking again...")
		time.Sleep(5 * time.Second)
	}
}
