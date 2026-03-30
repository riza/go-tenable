# Tenable Security Center (SC) — Go SDK

The `sc` package provides a Go client for the [Tenable Security Center](https://www.tenable.com/products/security-center) REST API, covering **~85 endpoints** with full CRUD operations.

> All methods accept a `context.Context` as the first parameter for cancellation and timeout support.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/riza/go-tenable/sc"
)

func main() {
	client := sc.NewClient("https://sc.example.com",
		sc.WithAPIKey("your-access-key", "your-secret-key"),
	)

	scans, err := client.Scan.List(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range scans.Usable {
		fmt.Printf("%s  %s\n", s.ID, s.Name)
	}
}
```

## Authentication

### API Key

```go
client := sc.NewClient(baseURL, sc.WithAPIKey(accessKey, secretKey))
```

### Session Token

```go
client := sc.NewClient(baseURL)
token, err := client.Token.Create(context.Background(), &sc.TokenCreateInput{
	Username: "admin",
	Password: "password",
})
```

## Client Options

| Option | Description |
|--------|-------------|
| `WithAPIKey(access, secret)` | API key authentication |
| `WithHTTPClient(client)` | Custom `*http.Client` |
| `WithInsecureSkipVerify()` | Disable TLS certificate verification |

## Error Handling

API errors are returned as `*sc.APIError` with HTTP status code, SC error code, and message:

```go
scan, err := client.Scan.Get(context.Background(), "999999")
if err != nil {
	var apiErr *sc.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("HTTP %d — SC error %d: %s\n",
			apiErr.StatusCode, apiErr.ErrorCode, apiErr.ErrorMsg)
	}
}
```

## Available Services

All services support the standard CRUD operations available for their endpoint, plus any resource-specific sub-actions (launch, copy, share, export, etc.).

`AcceptRiskRule` · `AgentGroup` · `AgentResultsSync` · `AgentScan` · `Alert` · `Analysis` · `ARC` · `ARCTemplate` · `Asset` · `AssetTemplate` · `AttributeSet` · `AuditFile` · `AuditFileTemplate` · `Bulk` · `ConfigSection` · `Configuration` · `Credential` · `CurrentOrganization` · `CurrentUser` · `CustomPlugins` · `DashboardComponent` · `DashboardTab` · `DashboardTemplate` · `DeviceInfo` · `DirectorInsights` · `DirectorOrganization` · `DirectorRepository` · `DirectorScan` · `DirectorScanPolicy` · `DirectorScanResult` · `DirectorScanZone` · `DirectorScanner` · `DirectorSystem` · `DirectorUser` · `Feed` · `File` · `FreezeWindow` · `Group` · `Hosts` · `Job` · `LCE` · `LCEClient` · `LCEPolicy` · `LDAP` · `LicenseInfo` · `Lumin` · `MDM` · `Notification` · `Organization` · `OrganizationSecurityManager` · `OrganizationUser` · `PassiveScanner` · `Plugin` · `PluginFamily` · `PublishingSite` · `Query` · `RecastRiskRule` · `Report` · `ReportDefinition` · `ReportImage` · `ReportTemplate` · `Repository` · `Role` · `SAML` · `Scan` · `Scanner` · `ScanPolicy` · `ScanPolicyTemplate` · `ScanResult` · `ScanZone` · `SensorProxy` · `SoftwareUpdate` · `Solutions` · `SSHKey` · `Status` · `Style` · `StyleFamily` · `System` · `TenableSCInstance` · `TESAdminRoles` · `TESUserPermissions` · `Ticket` · `Token` · `User` · `VulnerabilityRoutingRule` · `VulnerabilityRoutingSummary` · `WASScan` · `WASScanner`

## Examples

```bash
SC_URL=https://sc.example.com \
SC_ACCESS_KEY=xxx \
SC_SECRET_KEY=yyy \
go run ./examples/sc/scheduled-scans/
```

See [`examples/sc/`](../examples/sc/) for runnable examples.
