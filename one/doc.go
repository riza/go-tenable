// Package one provides a Go client for the Tenable One API.
//
// # Authentication
//
// Create a client with API key authentication:
//
//	client := one.NewClient("https://cloud.tenable.com",
//		one.WithAPIKey(accessKey, secretKey),
//	)
//
// # Services
//
// The client exposes services for each API resource area. Access them
// as fields on the Client struct:
//
//	paths, err := client.AttackPathService.SearchAttackPaths(ctx, req)
//	cards, err := client.ExposureViewService.ListExposureCards(ctx)
//
// # Error Handling
//
// API errors are returned as *APIError values containing the HTTP status code
// and error message:
//
//	var apiErr *one.APIError
//	if errors.As(err, &apiErr) {
//		fmt.Printf("HTTP %d: %s\n", apiErr.StatusCode, apiErr.ErrorMsg)
//	}
package one
