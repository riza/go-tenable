package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/riza/go-tenable/vm"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run . <schedule_uuid>")
	}
	scheduleUuid := os.Args[1]

	baseURL := os.Getenv("VM_URL")
	accessKey := os.Getenv("VM_ACCESS_KEY")
	secretKey := os.Getenv("VM_SECRET_KEY")

	if baseURL == "" || accessKey == "" || secretKey == "" {
		log.Fatal("VM_URL, VM_ACCESS_KEY, and VM_SECRET_KEY environment variables are required")
	}

	client := vm.NewClient(baseURL,
		vm.WithAPIKey(accessKey, secretKey),
	)

	err := client.ScanControlService.VmScansStopForce(context.Background(), scheduleUuid)
	if err != nil {
		log.Fatalf("failed to force stop scan: %v", err)
	}

	fmt.Printf("Successfully sent force stop command for scan with schedule UUID: %s\n", scheduleUuid)
}
