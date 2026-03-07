.PHONY: build test vet fmt lint check

## Build all packages
build:
	go build ./...

## Run tests
test:
	go test ./...

## Run tests with race detector
test-race:
	go test -race ./...

## Run go vet
vet:
	go vet ./...

## Format code
fmt:
	gofmt -w .
	goimports -w . 2>/dev/null || true

## Run linter (requires golangci-lint)
lint:
	golangci-lint run ./...

## Run all checks (vet + test)
check: vet test
	@echo "All checks passed."
