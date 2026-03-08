// Package tenable provides Go SDK packages for the Tenable platform APIs.
//
// The SDK is organized into sub-packages, each targeting a specific Tenable product:
//
//   - [github.com/riza/go-tenable/sc] — Tenable Security Center (on-prem)
//   - [github.com/riza/go-tenable/vm] — Tenable Vulnerability Management (cloud)
//   - [github.com/riza/go-tenable/platform] — Tenable Platform API (access control, agents, scanners, etc.)
//   - [github.com/riza/go-tenable/one] — Tenable One API (Attack Path, Exposure View, Inventory, Exports)
//
// All packages follow the same design pattern: create a [Client] with
// authentication options, then call methods on the service fields.
//
// Zero external dependencies — built entirely on Go's standard library.
package tenable
