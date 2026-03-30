# go-tenable

[![Go Reference](https://pkg.go.dev/badge/github.com/riza/go-tenable.svg)](https://pkg.go.dev/github.com/riza/go-tenable)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/riza/go-tenable)](https://goreportcard.com/report/github.com/riza/go-tenable)

Go SDK for the [Tenable](https://www.tenable.com/) APIs.

| Package | Product | Documentation |
|---------|---------|---------------|
| [`sc`](sc/) | Tenable Security Center (on-prem) | [**SC Documentation →**](doc/SC.md) |
| [`vm`](vm/) | Tenable Vulnerability Management (cloud) | [**VM Documentation →**](doc/VM.md) |
| [`platform`](platform/) | Tenable Platform API | [**Platform Documentation →**](doc/PLATFORM.md) |
| [`one`](one/) | Tenable One API | [**One Documentation →**](doc/ONE.md) |

> **Zero external dependencies** — built entirely on Go's standard library.

## Installation

```bash
go get github.com/riza/go-tenable
```

Requires **Go 1.21+**.

## Quick Start

### Security Center

```go
client := sc.NewClient("https://sc.example.com",
	sc.WithAPIKey("access-key", "secret-key"),
)

scans, err := client.Scan.List(context.Background(), nil)
```

### Vulnerability Management

```go
client := vm.NewClient("https://cloud.tenable.com",
	vm.WithAPIKey("access-key", "secret-key"),
)

assets, err := client.AssetsService.ListAssets(context.Background())
```

### Tenable Platform API

```go
client := platform.NewClient("https://cloud.tenable.com",
	platform.WithAPIKey("access-key", "secret-key"),
)

users, err := client.UsersService.ListUsers(context.Background())
```

### Tenable One API

```go
client := one.NewClient("https://cloud.tenable.com",
	one.WithAPIKey("access-key", "secret-key"),
)

paths, err := client.AttackPathService.SearchAttackPaths(context.Background(), &one.APASearchAttackPathsRequest{
	Limit: 10,
})
```

For full examples, authentication options, error handling, and service listings see:
- **[SC Documentation](doc/SC.md)** — Authentication, services list (~85), error handling, examples
- **[VM Documentation](doc/VM.md)** — Authentication, services list (23), error handling, examples
- **[Platform Documentation](doc/PLATFORM.md)** — Authentication, services list, error handling, examples
- **[One Documentation](doc/ONE.md)** — Authentication, services list, error handling, examples

## Testing

### Unit Tests

```bash
go test ./...
```

### Integration Tests

Integration tests run against live Tenable API instances. They are gated behind the `integration` build tag and will skip automatically if credentials are not configured.

```bash
# Set credentials for the packages you want to test
export TENABLE_SC_URL=https://sc.example.com
export TENABLE_SC_ACCESS_KEY=xxx
export TENABLE_SC_SECRET_KEY=yyy

export TENABLE_VM_URL=https://cloud.tenable.com
export TENABLE_VM_ACCESS_KEY=xxx
export TENABLE_VM_SECRET_KEY=yyy

export TENABLE_PLATFORM_URL=https://cloud.tenable.com
export TENABLE_PLATFORM_ACCESS_KEY=xxx
export TENABLE_PLATFORM_SECRET_KEY=yyy

export TENABLE_ONE_URL=https://cloud.tenable.com
export TENABLE_ONE_ACCESS_KEY=xxx
export TENABLE_ONE_SECRET_KEY=yyy

# Run integration tests
go test -tags=integration ./...
```

Or use the Makefile target:

```bash
TENABLE_SC_URL=https://sc.example.com TENABLE_SC_ACCESS_KEY=xxx TENABLE_SC_SECRET_KEY=yyy \
  make test-integration
```

## Examples

Runnable examples are in the [`examples/`](examples/) directory:

| Area | Examples |
|------|----------|
| SC | `scheduled-scans` |
| VM | `list-assets`, `list-scans`, `stop-scan` |
| Platform | `list-groups` |
| One | `search-attack-paths`, `inventory-search`, `tags-search`, `export-assets`, `exposure-cards` |

```bash
# SC — List scheduled scans
SC_URL=https://sc.example.com SC_ACCESS_KEY=xxx SC_SECRET_KEY=yyy \
  go run ./examples/sc/scheduled-scans/

# VM — List assets
VM_URL=https://cloud.tenable.com VM_ACCESS_KEY=xxx VM_SECRET_KEY=yyy \
  go run ./examples/vm/list-assets/

# VM — List scans
VM_URL=https://cloud.tenable.com VM_ACCESS_KEY=xxx VM_SECRET_KEY=yyy \
  go run ./examples/vm/list-scans/

# VM — Stop scan
VM_URL=https://cloud.tenable.com VM_ACCESS_KEY=xxx VM_SECRET_KEY=yyy \
  go run ./examples/vm/stop-scan/

# Platform — List groups
PLATFORM_URL=https://cloud.tenable.com PLATFORM_ACCESS_KEY=xxx PLATFORM_SECRET_KEY=yyy \
  go run ./examples/platform/list-groups/

# One — Search attack paths
ONE_URL=https://cloud.tenable.com ONE_ACCESS_KEY=xxx ONE_SECRET_KEY=yyy \
  go run ./examples/one/search-attack-paths/

# One — Inventory search
ONE_URL=https://cloud.tenable.com ONE_ACCESS_KEY=xxx ONE_SECRET_KEY=yyy \
  go run ./examples/one/inventory-search/

# One — Tags search
ONE_URL=https://cloud.tenable.com ONE_ACCESS_KEY=xxx ONE_SECRET_KEY=yyy \
  go run ./examples/one/tags-search/

# One — Export assets
ONE_URL=https://cloud.tenable.com ONE_ACCESS_KEY=xxx ONE_SECRET_KEY=yyy \
  go run ./examples/one/export-assets/

# One — Exposure cards
ONE_URL=https://cloud.tenable.com ONE_ACCESS_KEY=xxx ONE_SECRET_KEY=yyy \
  go run ./examples/one/exposure-cards/
```

## Project Structure

```
go-tenable/
├── sc/                     # Tenable Security Center SDK
├── vm/                     # Tenable Vulnerability Management SDK
├── platform/               # Tenable Platform API SDK
├── one/                    # Tenable One API SDK
├── examples/               # Runnable usage examples
├── doc/                    # Documentation
│   ├── SC.md               # SC package guide
│   ├── VM.md               # VM package guide
│   ├── PLATFORM.md         # Platform API package guide
│   └── ONE.md              # One API package guide
├── go.mod
├── LICENSE
├── CONTRIBUTING.md
└── Makefile
```

## Roadmap

- **Typed fields** — Replace `interface{}` fields with proper Go structs
- **Pagination helpers** — Iterator utilities for list endpoints
- **Rate limiting** — Configurable rate limiter
- **Retry with backoff** — Automatic retry for transient errors
- **Unified client** — Top-level `tenable.Client` targeting both SC and VM
- **CI pipeline** — GitHub Actions for lint, vet, test

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License — see [LICENSE](LICENSE) for details.
