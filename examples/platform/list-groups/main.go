// List groups and users from Tenable Platform API.
//
// Usage:
//
//	PLATFORM_URL=https://platform.tenable.com PLATFORM_ACCESS_KEY=xxx PLATFORM_SECRET_KEY=yyy go run .
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/riza/go-tenable/platform"
)

func main() {
	baseURL := os.Getenv("PLATFORM_URL")
	accessKey := os.Getenv("PLATFORM_ACCESS_KEY")
	secretKey := os.Getenv("PLATFORM_SECRET_KEY")

	if baseURL == "" || accessKey == "" || secretKey == "" {
		log.Fatal("PLATFORM_URL, PLATFORM_ACCESS_KEY, and PLATFORM_SECRET_KEY environment variables are required")
	}

	client := platform.NewClient(baseURL,
		platform.WithAPIKey(accessKey, secretKey),
		platform.WithInsecureSkipVerify(), // if needed for intercepting proxies
	)

	ctx := context.Background()

	// List groups
	fmt.Println("--- Groups ---")

	groups, err := client.GroupsService.ListGroups(ctx)
	if err != nil {
		log.Fatalf("failed to list groups: %v", err)
	}

	fmt.Printf("Found %d group(s)\n\n", len(groups.Groups))

	for _, g := range groups.Groups {
		fmt.Printf("ID: %d | UUID: %s | Name: %s | User Count: %d\n",
			g.Id,
			g.Uuid,
			g.Name,
			g.UserCount,
		)
	}

	// List users
	fmt.Println("\n--- Users ---")

	users, err := client.UsersService.ListUsers(ctx)
	if err != nil {
		log.Fatalf("failed to list users: %v", err)
	}

	fmt.Printf("Found %d user(s)\n\n", users.Total)

	for _, u := range users.Users {
		fmt.Printf("ID: %d | UUID: %s | Username: %s | Email: %s | Enabled: %t\n",
			u.Id,
			u.Uuid,
			u.Username,
			u.Email,
			u.Enabled,
		)
	}

	// List access groups
	fmt.Println("\n--- Access Groups (v1) ---")

	accessGroups, err := client.AccessGroupsService.ListAccessGroups(ctx)
	if err != nil {
		log.Fatalf("failed to list access groups: %v", err)
	}

	fmt.Printf("Found %d access group(s)\n\n", accessGroups.Total)

	for _, ag := range accessGroups.AccessGroups {
		fmt.Printf("ID: %s | Name: %s | Description: %s\n",
			ag.Id,
			ag.Name,
			ag.Description,
		)
	}

	// List scanners
	fmt.Println("\n--- Scanners ---")

	scanners, err := client.ScannersService.ListScanners(ctx)
	if err != nil {
		log.Fatalf("failed to list scanners: %v", err)
	}

	fmt.Printf("Found %d scanner(s)\n\n", scanners.Total)

	for _, s := range scanners.Scanners {
		ip := ""
		if len(s.IpAddresses) > 0 {
			ip = s.IpAddresses[0]
		}
		fmt.Printf("ID: %d | Name: %s | IP: %s | Status: %s\n",
			s.Id,
			s.Name,
			ip,
			s.Status,
		)
	}

	// Get server info
	fmt.Println("\n--- Server Info ---")

	serverInfo, err := client.ServerService.GetServerInfo(ctx)
	if err != nil {
		log.Fatalf("failed to get server info: %v", err)
	}

	fmt.Printf("Name: %s | Version: %s | Build: %s | Edition: %s\n",
		serverInfo.Name,
		serverInfo.Version,
		serverInfo.Build,
		serverInfo.Edition,
	)
}
