package sc

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// apiResponse is the envelope for all SC API responses.
type apiResponse struct {
	Type      string          `json:"type"`
	Response  json.RawMessage `json:"response"`
	ErrorCode int             `json:"error_code"`
	ErrorMsg  string          `json:"error_msg"`
	Warnings  []string        `json:"warnings"`
	Timestamp int64           `json:"timestamp"`
}

// Option configures a Client.
type Option func(*Client)

// Client is the Tenable Security Center API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	UserAgent  string
	accessKey  string
	secretKey  string

	Token *TokenService
	Scan  *ScanService

	AcceptRiskRule              *AcceptRiskRuleService
	AgentGroup                  *AgentGroupService
	AgentResultsSync            *AgentResultsSyncService
	AgentScan                   *AgentScanService
	Alert                       *AlertService
	Analysis                    *AnalysisService
	ARC                         *ARCService
	ARCTemplate                 *ARCTemplateService
	Asset                       *AssetService
	AssetTemplate               *AssetTemplateService
	AttributeSet                *AttributeSetService
	AuditFile                   *AuditFileService
	AuditFileTemplate           *AuditFileTemplateService
	Bulk                        *BulkService
	ConfigSection               *ConfigSectionService
	Configuration               *ConfigurationService
	Credential                  *CredentialService
	CurrentOrganization         *CurrentOrganizationService
	CurrentUser                 *CurrentUserService
	CustomPlugins               *CustomPluginsService
	DashboardComponent          *DashboardComponentService
	DashboardTab                *DashboardTabService
	DashboardTemplate           *DashboardTemplateService
	DeviceInfo                  *DeviceInfoService
	DirectorInsights            *DirectorInsightsService
	DirectorOrganization        *DirectorOrganizationService
	DirectorRepository          *DirectorRepositoryService
	DirectorScan                *DirectorScanService
	DirectorScanPolicy          *DirectorScanPolicyService
	DirectorScanResult          *DirectorScanResultService
	DirectorScanZone            *DirectorScanZoneService
	DirectorScanner             *DirectorScannerService
	DirectorSystem              *DirectorSystemService
	DirectorUser                *DirectorUserService
	Feed                        *FeedService
	File                        *FileService
	FreezeWindow                *FreezeWindowService
	Group                       *GroupService
	Hosts                       *HostsService
	Job                         *JobService
	LCE                         *LCEService
	LCEClient                   *LCEClientService
	LCEPolicy                   *LCEPolicyService
	LDAP                        *LDAPService
	LicenseInfo                 *LicenseInfoService
	Lumin                       *LuminService
	MDM                         *MDMService
	Notification                *NotificationService
	Organization                *OrganizationService
	OrganizationSecurityManager *OrganizationSecurityManagerService
	OrganizationUser            *OrganizationUserService
	PassiveScanner              *PassiveScannerService
	Plugin                      *PluginService
	PluginFamily                *PluginFamilyService
	PublishingSite              *PublishingSiteService
	Query                       *QueryService
	RecastRiskRule              *RecastRiskRuleService
	Report                      *ReportService
	ReportDefinition            *ReportDefinitionService
	ReportImage                 *ReportImageService
	ReportTemplate              *ReportTemplateService
	Repository                  *RepositoryService
	Role                        *RoleService
	SAML                        *SAMLService
	Scanner                     *ScannerService
	ScanPolicy                  *ScanPolicyService
	ScanPolicyTemplate          *ScanPolicyTemplateService
	ScanResult                  *ScanResultService
	ScanZone                    *ScanZoneService
	SensorProxy                 *SensorProxyService
	SoftwareUpdate              *SoftwareUpdateService
	Solutions                   *SolutionsService
	SSHKey                      *SSHKeyService
	Status                      *StatusService
	Style                       *StyleService
	StyleFamily                 *StyleFamilyService
	System                      *SystemService
	TenableSCInstance           *TenableSCInstanceService
	TESAdminRoles               *TESAdminRolesService
	TESUserPermissions          *TESUserPermissionsService
	Ticket                      *TicketService
	User                        *UserService
	VulnerabilityRoutingRule    *VulnerabilityRoutingRuleService
	VulnerabilityRoutingSummary *VulnerabilityRoutingSummaryService
	WASScan                     *WASScanService
	WASScanner                  *WASScannerService
}

// NewClient creates a new SC API client.
func NewClient(baseURL string, opts ...Option) *Client {
	baseURL = strings.TrimRight(baseURL, "/")
	c := &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
		UserAgent:  "go-tenable/0.1",
	}
	for _, opt := range opts {
		opt(c)
	}
	c.initServices()
	return c
}

// WithAPIKey sets API key authentication.
func WithAPIKey(accessKey, secretKey string) Option {
	return func(c *Client) {
		c.accessKey = accessKey
		c.secretKey = secretKey
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = hc
	}
}

// WithInsecureSkipVerify disables TLS certificate verification.
func WithInsecureSkipVerify() Option {
	return func(c *Client) {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		c.HTTPClient = &http.Client{Transport: transport}
	}
}

// QueryParams holds optional query string parameters for API requests.
type QueryParams map[string][]string

func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	return c.newRequestWithParams(ctx, method, path, body, nil)
}

func (c *Client) newRequestWithParams(ctx context.Context, method, path string, body interface{}, params QueryParams) (*http.Request, error) {
	u := c.BaseURL + "/rest" + path

	if len(params) > 0 {
		q := url.Values{}
		for k, v := range params {
			for _, val := range v {
				q.Add(k, val)
			}
		}
		u += "?" + q.Encode()
	}

	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("sc: marshal request body: %w", err)
		}
		buf = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if c.accessKey != "" {
		req.Header.Set("x-apikey", fmt.Sprintf("accesskey=%s; secretkey=%s;", c.accessKey, c.secretKey))
	}

	return req, nil
}

func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*apiResponse, error) {
	return c.doRequestWithParams(ctx, method, path, body, nil)
}

func (c *Client) doRequestWithParams(ctx context.Context, method, path string, body interface{}, params QueryParams) (*apiResponse, error) {
	req, err := c.newRequestWithParams(ctx, method, path, body, params)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sc: request failed: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("sc: read response body: %w", err)
	}

	var apiResp apiResponse
	if err := json.Unmarshal(data, &apiResp); err != nil {
		return nil, fmt.Errorf("sc: unmarshal response (HTTP %d): %w", resp.StatusCode, err)
	}

	if apiResp.ErrorCode != 0 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			ErrorCode:  apiResp.ErrorCode,
			ErrorMsg:   apiResp.ErrorMsg,
		}
	}

	if resp.StatusCode >= 400 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			ErrorCode:  apiResp.ErrorCode,
			ErrorMsg:   apiResp.ErrorMsg,
		}
	}

	return &apiResp, nil
}

func (c *Client) get(ctx context.Context, path string) (*apiResponse, error) {
	return c.doRequest(ctx, http.MethodGet, path, nil)
}

func (c *Client) getWithParams(ctx context.Context, path string, params QueryParams) (*apiResponse, error) {
	return c.doRequestWithParams(ctx, http.MethodGet, path, nil, params)
}

func (c *Client) post(ctx context.Context, path string, body interface{}) (*apiResponse, error) {
	return c.doRequest(ctx, http.MethodPost, path, body)
}

func (c *Client) put(ctx context.Context, path string, body interface{}) (*apiResponse, error) { //nolint:unused // kept for API completeness
	return c.doRequest(ctx, http.MethodPut, path, body)
}

func (c *Client) patch(ctx context.Context, path string, body interface{}) (*apiResponse, error) {
	return c.doRequest(ctx, http.MethodPatch, path, body)
}

func (c *Client) delete(ctx context.Context, path string) (*apiResponse, error) {
	return c.doRequest(ctx, http.MethodDelete, path, nil)
}
