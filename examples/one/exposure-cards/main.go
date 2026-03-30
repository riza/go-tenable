// View Exposure Cards from Tenable One.
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

	fmt.Println("--- Exposure View Cards ---")

	// Disclaimer: This endpoint sometimes returns HTTP 500 depending on your Tenable One backend provision state.
	cardsResp, err := client.ExposureViewService.ListCards(ctx)
	if err != nil {
		fmt.Printf("Note: Exposure View API is not fully provisioned or encountering transient errors.\nError: %v\n", err)
		return
	}

	fmt.Printf("Total Cards Found: %d\n", cardsResp.Total)
	for i, card := range cardsResp.Cards {
		fmt.Printf("[%d] ID: %s | Name: %s | Category: %s | Type: %s\n",
			i+1, card.ID, card.Name, card.Category, card.Type)
	}
}
