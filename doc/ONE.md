# Tenable One API — Go SDK

The `one` package provides a Go client for the [Tenable One API](https://developer.tenable.com), covering platform-level services such as attack paths, exposure views, global inventory, and tags.

> All methods accept a `context.Context` as the first parameter for cancellation and timeout support.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/riza/go-tenable/one"
)

func main() {
	client := one.NewClient("https://cloud.tenable.com",
		one.WithAPIKey("your-access-key", "your-secret-key"),
	)

	searchReq := &one.APASearchAttackPathsRequest{
		Limit: 10,
	}

	paths, err := client.AttackPathService.SearchAttackPaths(context.Background(), searchReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d attack paths\n", paths.Total)
}
```

## Authentication

```go
client := one.NewClient(baseURL, one.WithAPIKey(accessKey, secretKey))
```

## Client Options

| Option | Description |
|--------|-------------|
| `WithAPIKey(access, secret)` | API key authentication |
| `WithHTTPClient(client)` | Custom `*http.Client` |
| `WithInsecureSkipVerify()` | Disable TLS certificate verification |

## Error Handling

API errors are returned as `*one.APIError` with HTTP status code and error message:

```go
_, err := client.AttackPathService.SearchAttackPaths(ctx, req)
if err != nil {
	var apiErr *one.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("HTTP %d: %s\n", apiErr.StatusCode, apiErr.ErrorMsg)
	}
}
```

## Available Services

`AttackPath` · `Export` · `ExposureView` · `Inventory` · `Tags`

## Examples

Run the search attack paths example:

```bash
ONE_URL=https://cloud.tenable.com \
ONE_ACCESS_KEY=xxx \
ONE_SECRET_KEY=yyy \
go run ./examples/one/search-attack-paths/
```

See [`examples/one/`](../examples/one/) for runnable examples.
