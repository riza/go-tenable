// Search and manage tags in Tenable One.
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

	// 1. Search Tags
	fmt.Println("--- Searching Platform Tags ---")
	tagsReq := &one.TagsSearchRequest{
		Limit: intPtr(10),
	}

	tagsResp, err := client.TagsService.SearchTags(ctx, tagsReq)
	if err != nil {
		log.Fatalf("failed to search tags: %v", err)
	}

	fmt.Printf("Total Tags Found: %d\n", tagsResp.Total)
	for i, tag := range tagsResp.Tags {
		fmt.Printf("[%d] UUID: %s | Category: %s | Value: %s | Sources: %v\n",
			i+1, tag.Uuid, tag.Key, tag.Value, tag.Sources)
	}

	// 2. See Tag Properties (Metadata about what filters are valid)
	fmt.Println("\n--- Valid Tag Properties ---")
	props, err := client.TagsService.GetProperties(ctx)
	if err != nil {
		log.Fatalf("failed to get tag properties: %v", err)
	}

	fmt.Printf("Total Properties Categories: %d\n", props.Total)
	for _, prop := range props.Properties {
		if name, ok := prop["name"]; ok {
			fmt.Printf("- Property: %s\n", name)
		}
	}
}
