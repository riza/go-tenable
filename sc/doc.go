// Package sc provides a Go client for the Tenable Security Center REST API.
//
// # Authentication
//
// The client supports API key authentication via the x-apikey header:
//
//	client := sc.NewClient("https://sc.example.com",
//		sc.WithAPIKey("access-key", "secret-key"),
//	)
//
// Session-based authentication is available through the TokenService:
//
//	client := sc.NewClient("https://sc.example.com")
//	token, err := client.Token.Create(&sc.TokenCreateInput{
//		Username: "admin",
//		Password: "password",
//	})
//
// # Usage
//
// All API resources are exposed as service fields on the Client.
// Each service provides methods matching the available API operations
// (List, Get, Create, Update, Delete, and resource-specific actions).
//
//	scans, err := client.Scan.List(nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, s := range scans.Usable {
//		fmt.Println(s.ID, s.Name)
//	}
//
// # Error Handling
//
// API errors are returned as *APIError values containing the HTTP status code,
// SC-specific error code, and error message:
//
//	scan, err := client.Scan.Get("999")
//	if err != nil {
//		var apiErr *sc.APIError
//		if errors.As(err, &apiErr) {
//			fmt.Printf("HTTP %d, code %d: %s\n",
//				apiErr.StatusCode, apiErr.ErrorCode, apiErr.ErrorMsg)
//		}
//	}
package sc
