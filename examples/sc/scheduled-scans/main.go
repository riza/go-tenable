// List all scheduled scans from Tenable Security Center.
//
// Usage:
//
//	SC_URL=https://sc.example.com SC_ACCESS_KEY=xxx SC_SECRET_KEY=yyy go run .
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/riza/go-tenable/sc"
)

func main() {
	baseURL := os.Getenv("SC_URL")
	accessKey := os.Getenv("SC_ACCESS_KEY")
	secretKey := os.Getenv("SC_SECRET_KEY")

	if baseURL == "" || accessKey == "" || secretKey == "" {
		log.Fatal("SC_URL, SC_ACCESS_KEY, and SC_SECRET_KEY environment variables are required")
	}

	client := sc.NewClient(baseURL,
		sc.WithAPIKey(accessKey, secretKey),
		sc.WithInsecureSkipVerify(),
	)

	result, err := client.Scan.List(context.Background(), &sc.ScanListOptions{
		Fields: sc.ScanAllFields,
	})
	if err != nil {
		log.Fatalf("failed to list scans: %v", err)
	}

	fmt.Printf("Found %d usable scan(s)\n\n", len(result.Usable))

	var scheduled int
	for _, s := range result.Usable {
		schedType := "N/A"
		enabled := "N/A"
		nextRun := "N/A"

		if s.Schedule != nil {
			schedType = s.Schedule.Type
			enabled = s.Schedule.Enabled
			if s.Schedule.NextRun > 0 {
				nextRun = time.Unix(int64(s.Schedule.NextRun), 0).Format(time.RFC3339)
			}
		}

		isScheduled := s.Schedule != nil && s.Schedule.Type != "" && s.Schedule.Type != "never"
		if isScheduled {
			scheduled++
		}

		fmt.Printf("%-6s %-40s schedule=%-12s enabled=%-5s next=%s\n",
			s.ID,
			truncate(s.Name, 40),
			schedType,
			enabled,
			nextRun,
		)
	}
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
