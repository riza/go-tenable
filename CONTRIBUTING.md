# Contributing to go-tenable

Thank you for your interest in contributing! This document provides guidelines for contributing to the project.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/<your-username>/go-tenable.git`
3. Create a feature branch: `git checkout -b feature/my-feature`
4. Make your changes
5. Run tests and checks: `make check`
6. Commit your changes: `git commit -m "feat: add my feature"`
7. Push to your fork: `git push origin feature/my-feature`
8. Open a Pull Request

## Development

### Prerequisites

- Go 1.21+
- (Optional) [golangci-lint](https://golangci-lint.run/) for linting

### Running Tests

```bash
make test
```

### Running All Checks

```bash
make check
```

This runs `go vet`, `go test`, and formatting checks.

## Code Style

- Follow standard Go conventions and [Effective Go](https://go.dev/doc/effective-go)
- Run `gofmt` before committing (or use `make fmt`)
- Add godoc comments for all exported types, functions, and methods
- Keep package APIs consistent with the existing service pattern

## Adding a New Service

Each API resource maps to a single Go file in the corresponding package (`sc/` or `vm/`):

1. Create `sc/<resource>.go` or `vm/<resource>.go`
2. Define the service struct with a `client *Client` field
3. Add exported methods for each API endpoint
4. Register the service in `initServices()` (in `services.go` or `client.go`)
5. Add tests in `<resource>_test.go`

## Reporting Issues

- Use GitHub Issues
- Include Go version, OS, and a minimal reproduction case
- For security vulnerabilities, please email the maintainer directly instead of opening a public issue

## Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` new features
- `fix:` bug fixes
- `docs:` documentation changes
- `refactor:` code refactoring
- `test:` adding or updating tests
- `chore:` maintenance tasks
