// Search attack paths from Tenable One.
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

	excludeResolved := false
	attackPathsReq := &one.APASearchAttackPathsRequest{
		Limit:           100,
		Sort:            "priority:desc",
		ExcludeResolved: &excludeResolved,
		Filter: one.APAFilterGroup{
			Operator: "AND",
			Value: []one.APAFilterCondition{
				{
					Property: "data_source",
					Operator: "eq",
					Value:    "tenable",
				},
			},
		},
	}

	attackPaths, err := client.AttackPathService.SearchAttackPaths(ctx, attackPathsReq)
	if err != nil {
		log.Fatalf("failed to search attack paths: %v", err)
	}

	fmt.Printf("Found %d attack path(s)\n\n", attackPaths.Total)

	for _, ap := range attackPaths.AttackPaths {
		fmt.Printf("VectorID: %s | Name: %s | Status: %s | AES: %.1f | ACR: %.1f\n",
			ap.VectorID,
			ap.Name,
			ap.PathStatus,
			ap.FirstAES,
			ap.LastACR,
		)
	}

	// Search attack techniques
	fmt.Println("\n--- Attack Techniques ---")

	techniquesReq := &one.APASearchAttackTechniquesRequest{
		Limit: 100,
		Filter: one.APAFilterGroup{
			Operator: "AND",
			Value: []one.APAFilterCondition{
				{
					Property: "priority",
					Operator: "eq",
					Value:    "low",
				},
			},
		},
	}

	techniques, err := client.AttackPathService.SearchAttackTechniques(ctx, techniquesReq)
	if err != nil {
		log.Fatalf("failed to search attack techniques: %v", err)
	}

	fmt.Printf("Found %d attack technique(s)\n\n", techniques.Total)

	for _, t := range techniques.Techniques {
		fmt.Printf("ID: %s | Name: %s | Tactics: %v | Count: %d | Priority: %s\n",
			t.MitreID,
			t.TechniqueName,
			t.Tactics,
			t.Count,
			t.Priority,
		)
	}
}
