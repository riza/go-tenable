# Tenable Platform API — Go SDK

The `platform` package provides a Go client for the [Tenable Platform API](https://developer.tenable.com), covering core services including access control, agents, credentials, scanners, and user management.

> All methods accept a `context.Context` as the first parameter for cancellation and timeout support.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/riza/go-tenable/platform"
)

func main() {
	client := platform.NewClient("https://cloud.tenable.com",
		platform.WithAPIKey("your-access-key", "your-secret-key"),
	)

	users, err := client.UsersService.ListUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d users\n", len(users.Users))
}
```

## Authentication

```go
client := platform.NewClient(baseURL, platform.WithAPIKey(accessKey, secretKey))
```

## Client Options

| Option | Description |
|--------|-------------|
| `WithAPIKey(access, secret)` | API key authentication |
| `WithHTTPClient(client)` | Custom `*http.Client` |
| `WithInsecureSkipVerify()` | Disable TLS certificate verification |

## Error Handling

API errors are returned as `*platform.APIError` with HTTP status code and error message:

```go
_, err := client.UsersService.ListUsers(ctx)
if err != nil {
	var apiErr *platform.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("HTTP %d: %s\n", apiErr.StatusCode, apiErr.ErrorMsg)
	}
}
```

## Available Services

`AccessControl` · `AccessGroups` · `ActivityLog` · `Agents` · `CloudConnectors` · `Credentials` · `Exclusions` · `Groups` · `Networks` · `Permissions` · `Profiles` · `RecastRules` · `Scanners` · `Server` · `Tags` · `TargetGroups` · `Users`

## Examples

Run the list groups example:

```bash
PLATFORM_URL=https://cloud.tenable.com \
PLATFORM_ACCESS_KEY=xxx \
PLATFORM_SECRET_KEY=yyy \
go run ./examples/platform/list-groups/
```

See [`examples/platform/`](../examples/platform/) for runnable examples.
