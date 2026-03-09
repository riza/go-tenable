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
	)

	result, err := client.ScansService.List(context.Background())
	if err != nil {
		log.Fatalf("failed to list scans: %v", err)
	}

	fmt.Println("Scans with 'running' or 'stopping' status:")
	fmt.Printf("%-10s %-36s %-15s %s\n", "ID", "Schedule UUID", "Status", "Name")
	fmt.Println("--------------------------------------------------------------------------------")

	for _, s := range result.Scans {
		if s.Status == "running" || s.Status == "stopping" {
			fmt.Printf("%-10d %-36s %-15s %s\n", s.Id, s.ScheduleUuid, s.Status, s.Name)
		}
	}
}
