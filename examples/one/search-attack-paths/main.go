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
		one.WithInsecureSkipVerify(), // if needed for intercepting proxies
	)

	ctx := context.Background()

	// Search attack paths
	attackPathsReq := &one.APASearchAttackPathsRequest{
		Limit:  100,
		Offset: 0,
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
		fmt.Printf("ID: %s | Name: %s | Score: %d | Assets: %d\n",
			ap.ID,
			ap.Name,
			ap.Score,
			ap.AssetCount,
		)
	}

	// Search attack techniques
	fmt.Println("\n--- Attack Techniques ---")

	techniquesReq := &one.APASearchAttackTechniquesRequest{
		Limit:  100,
		Offset: 0,
		Filter: one.APAFilterGroup{
			Operator: "and",
			Value: []one.APAFilterCondition{
				{
					Property: "priority",
					Operator: "==",
					Value:    "low",
				},
			},
		},
	}

	techniques, err := client.AttackPathService.SearchAttackTechniques(ctx, techniquesReq)
	if err != nil {
		fmt.Printf("API Error: %+v\n", err)
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

	/*
		// List exposure view cards
		fmt.Println("\n--- Exposure View Cards ---")

		cards, err := client.ExposureViewService.ListCards(ctx)
		if err != nil {
			log.Fatalf("failed to list exposure view cards: %v", err)
		}

		fmt.Printf("Found %d card(s)\n\n", cards.Total)

		for _, card := range cards.Cards {
			fmt.Printf("ID: %s | Name: %s | Type: %s | Category: %s\n",
				card.ID,
				card.Name,
				card.Type,
				card.Category,
			)
		}
	*/
}
