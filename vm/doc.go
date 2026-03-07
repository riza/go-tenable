// Package vm provides a Go client for the Tenable Vulnerability Management
// (cloud) REST API.
//
// # Authentication
//
// The client supports API key authentication:
//
//	client := vm.NewClient("https://cloud.tenable.com",
//		vm.WithAPIKey("access-key", "secret-key"),
//	)
//
// # Usage
//
// All API resources are exposed as service fields on the Client.
// Methods accept a context.Context as the first parameter for cancellation
// and timeout support.
//
//	assets, err := client.AssetsService.ListAssets(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, a := range assets.Assets {
//		fmt.Println(a.Id)
//	}
//
// # Error Handling
//
// API errors are returned as *APIError values containing the HTTP status code
// and error message:
//
//	_, err := client.AssetsService.ListAssets(ctx)
//	if err != nil {
//		var apiErr *vm.APIError
//		if errors.As(err, &apiErr) {
//			fmt.Printf("HTTP %d: %s\n", apiErr.StatusCode, apiErr.ErrorMsg)
//		}
//	}
package vm
