# go-tenable

[![Go Reference](https://pkg.go.dev/badge/github.com/riza/go-tenable.svg)](https://pkg.go.dev/github.com/riza/go-tenable)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/riza/go-tenable)](https://goreportcard.com/report/github.com/riza/go-tenable)

Go SDK for the [Tenable](https://www.tenable.com/) platform APIs.

| Package | Product | Documentation |
|---------|---------|---------------|
| [`sc`](sc/) | Tenable Security Center (on-prem) | [**SC Documentation →**](doc/SC.md) |
| [`vm`](vm/) | Tenable Vulnerability Management (cloud) | [**VM Documentation →**](doc/VM.md) |

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

scans, err := client.Scan.List(nil)
```

### Vulnerability Management

```go
client := vm.NewClient("https://cloud.tenable.com",
	vm.WithAPIKey("access-key", "secret-key"),
)

assets, err := client.AssetsService.ListAssets(context.Background())
```

For full examples, authentication options, error handling, and service listings see:
- **[SC Documentation](doc/SC.md)** — Authentication, services list (~85), error handling, examples
- **[VM Documentation](doc/VM.md)** — Authentication, services list (23), error handling, examples

## Examples

Runnable examples are in the [`examples/`](examples/) directory:

```bash
# SC — List scheduled scans
SC_URL=https://sc.example.com SC_ACCESS_KEY=xxx SC_SECRET_KEY=yyy \
  go run ./examples/sc/scheduled-scans/

# VM — List assets
VM_URL=https://cloud.tenable.com VM_ACCESS_KEY=xxx VM_SECRET_KEY=yyy \
  go run ./examples/vm/list-assets/
```

## Project Structure

```
go-tenable/
├── sc/                     # Tenable Security Center SDK
│   ├── client.go           # HTTP client, auth, request helpers
│   ├── error.go            # APIError and shared types
│   ├── services.go         # Service initialization
│   └── *.go                # One file per API resource (~85 services)
├── vm/                     # Tenable Vulnerability Management SDK
│   ├── client.go           # HTTP client, auth, request helpers
│   └── *.go                # One file per API resource (~23 services)
├── examples/               # Runnable usage examples
├── doc/                    # Documentation
│   ├── SC.md               # SC package guide
│   ├── VM.md               # VM package guide
│   └── sc/                 # SC API endpoint reference (markdown)
├── go.mod
├── LICENSE
├── CONTRIBUTING.md
└── Makefile
```

## Roadmap

- **Context support for SC** — Add `context.Context` to all SC service methods (VM already has it)
- **Typed fields** — Replace `interface{}` fields with proper Go structs
- **Pagination helpers** — Iterator utilities for list endpoints
- **Rate limiting** — Configurable rate limiter
- **Retry with backoff** — Automatic retry for transient errors
- **Export workflow helpers** — Polling utilities for VM async export APIs
- **Unified client** — Top-level `tenable.Client` targeting both SC and VM
- **Integration tests** — Tests against live instances (behind build tag)
- **CI pipeline** — GitHub Actions for lint, vet, test

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License — see [LICENSE](LICENSE) for details.
