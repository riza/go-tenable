.PHONY: build test test-integration vet fmt lint check

## Build all packages
build:
	go build ./...

## Run tests (unit tests only)
test:
	go test ./...

## Run integration tests against live Tenable APIs
## Requires TENABLE_* environment variables to be set
## Example:
##   TENABLE_SC_URL=https://sc.example.com \
##   TENABLE_SC_ACCESS_KEY=xxx \
##   TENABLE_SC_SECRET_KEY=yyy \
##   TENABLE_VM_URL=https://cloud.tenable.com \
##   TENABLE_VM_ACCESS_KEY=xxx \
##   TENABLE_VM_SECRET_KEY=yyy \
##   make test-integration
test-integration:
	go test -tags=integration ./...

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
