# Tenable Vulnerability Management (VM) — Go SDK

The `vm` package provides a Go client for the [Tenable Vulnerability Management](https://www.tenable.com/products/vulnerability-management) (cloud) REST API, covering **23 services** including assets, scans, exports, vulnerabilities, and workbenches.

> All methods accept a `context.Context` as the first parameter for cancellation and timeout support.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/riza/go-tenable/vm"
)

func main() {
	client := vm.NewClient("https://cloud.tenable.com",
		vm.WithAPIKey("your-access-key", "your-secret-key"),
	)

	assets, err := client.AssetsService.ListAssets(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d assets\n", assets.Total)
}
```

## Authentication

```go
client := vm.NewClient(baseURL, vm.WithAPIKey(accessKey, secretKey))
```

## Client Options

| Option | Description |
|--------|-------------|
| `WithAPIKey(access, secret)` | API key authentication |
| `WithHTTPClient(client)` | Custom `*http.Client` |
| `WithInsecureSkipVerify()` | Disable TLS certificate verification |

## Error Handling

API errors are returned as `*vm.APIError` with HTTP status code and error message:

```go
_, err := client.AssetsService.ListAssets(ctx)
if err != nil {
	var apiErr *vm.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("HTTP %d: %s\n", apiErr.StatusCode, apiErr.ErrorMsg)
	}
}
```

## Available Services

`AssetAttributes` · `Assets` · `Editor` · `ExportsAssets` · `ExportsComplianceData` · `ExportsVulnerabilities` · `File` · `Filters` · `Folders` · `Plugins` · `Policies` · `RemediationScans` · `Reports` · `ScanControl` · `ScanExports` · `ScanHistory` · `ScanResults` · `ScanStatus` · `ScanTasks` · `Scans` · `SharedCollections` · `Vulnerabilities` · `Workbenches`

## Examples

```bash
VM_URL=https://cloud.tenable.com \
VM_ACCESS_KEY=xxx \
VM_SECRET_KEY=yyy \
go run ./examples/vm/list-assets/
```

See [`examples/vm/`](../examples/vm/) for runnable examples.
