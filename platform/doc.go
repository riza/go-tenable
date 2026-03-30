// Package platform provides a Go client for the Tenable Platform API.
//
// # Authentication
//
// Create a client with API key authentication:
//
//	client := platform.NewClient("https://cloud.tenable.com",
//		platform.WithAPIKey(accessKey, secretKey),
//	)
//
// # Services
//
// The client exposes services for each API resource area. Access them
// as fields on the Client struct:
//
//	users, err := client.UsersService.ListUsers(ctx)
//	groups, err := client.GroupsService.ListGroups(ctx)
//
// # Error Handling
//
// API errors are returned as *APIError values containing the HTTP status code
// and error message:
//
//	var apiErr *platform.APIError
//	if errors.As(err, &apiErr) {
//		fmt.Printf("HTTP %d: %s\n", apiErr.StatusCode, apiErr.ErrorMsg)
//	}
package platform
