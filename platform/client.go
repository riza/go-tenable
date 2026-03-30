package platform

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

// APIError is a Tenable Platform API Error
type APIError struct {
	StatusCode int
	ErrorCode  int
	ErrorMsg   string
}

func (e *APIError) Error() string {
	if e.ErrorCode != 0 {
		return fmt.Sprintf("platform api error (HTTP %d): %d - %s", e.StatusCode, e.ErrorCode, e.ErrorMsg)
	}
	return fmt.Sprintf("platform api error (HTTP %d): %s", e.StatusCode, e.ErrorMsg)
}

// Option configures a Client.
type Option func(*Client)

// Client is the Tenable Platform API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	UserAgent  string
	accessKey  string
	secretKey  string

	// Services
	AccessControlService   *AccessControlService
	GroupsService          *GroupsService
	PermissionsService     *PermissionsService
	UsersService           *UsersService
	AccessGroupsService    *AccessGroupsService
	ActivityLogService     *ActivityLogService
	AgentsService          *AgentsService
	CloudConnectorsService *CloudConnectorsService
	CredentialsService     *CredentialsService
	ExclusionsService      *ExclusionsService
	NetworksService        *NetworksService
	Permissions2Service    *Permissions2Service
	ProfilesService        *ProfilesService
	RecastRulesService     *RecastRulesService
	ScannersService        *ScannersService
	ServerService          *ServerService
	TagsService            *TagsService
	TargetGroupsService    *TargetGroupsService
}

// NewClient creates a new Platform API client.
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

func (c *Client) initServices() {
	c.AccessControlService = &AccessControlService{client: c}
	c.GroupsService = &GroupsService{client: c}
	c.PermissionsService = &PermissionsService{client: c}
	c.UsersService = &UsersService{client: c}
	c.AccessGroupsService = &AccessGroupsService{client: c}
	c.ActivityLogService = &ActivityLogService{client: c}
	c.AgentsService = &AgentsService{client: c}
	c.CloudConnectorsService = &CloudConnectorsService{client: c}
	c.CredentialsService = &CredentialsService{client: c}
	c.ExclusionsService = &ExclusionsService{client: c}
	c.NetworksService = &NetworksService{client: c}
	c.Permissions2Service = &Permissions2Service{client: c}
	c.ProfilesService = &ProfilesService{client: c}
	c.RecastRulesService = &RecastRulesService{client: c}
	c.ScannersService = &ScannersService{client: c}
	c.ServerService = &ServerService{client: c}
	c.TagsService = &TagsService{client: c}
	c.TargetGroupsService = &TargetGroupsService{client: c}
}

// QueryParams holds optional query string parameters for API requests.
type QueryParams map[string][]string

func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	return c.newRequestWithParams(ctx, method, path, body, nil)
}

func (c *Client) newRequestWithParams(ctx context.Context, method, path string, body interface{}, params QueryParams) (*http.Request, error) {
	u := c.BaseURL + path

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
			return nil, fmt.Errorf("platform: marshal request body: %w", err)
		}
		buf = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if c.accessKey != "" && c.secretKey != "" {
		req.Header.Set("x-apikeys", fmt.Sprintf("accessKey=%s;secretKey=%s", c.accessKey, c.secretKey))
	}

	return req, nil
}

// apiErrorResponse is the response envelope for errors.
type apiErrorResponse struct {
	Error string `json:"error"`
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("platform: request failed: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("platform: read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
		}

		var errResp apiErrorResponse
		if err := json.Unmarshal(data, &errResp); err == nil {
			apiErr.ErrorMsg = errResp.Error
		} else {
			apiErr.ErrorMsg = string(data)
		}

		return nil, apiErr
	}

	return data, nil
}

func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	return c.doRequestWithParams(ctx, method, path, body, nil)
}

func (c *Client) doRequestWithParams(ctx context.Context, method, path string, body interface{}, params QueryParams) ([]byte, error) {
	req, err := c.newRequestWithParams(ctx, method, path, body, params)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	return c.doRequest(ctx, http.MethodGet, path, nil)
}

func (c *Client) getWithParams(ctx context.Context, path string, params QueryParams) ([]byte, error) {
	return c.doRequestWithParams(ctx, http.MethodGet, path, nil, params)
}

func (c *Client) post(ctx context.Context, path string, body interface{}) ([]byte, error) {
	return c.doRequest(ctx, http.MethodPost, path, body)
}

func (c *Client) put(ctx context.Context, path string, body interface{}) ([]byte, error) {
	return c.doRequest(ctx, http.MethodPut, path, body)
}

func (c *Client) delete(ctx context.Context, path string) ([]byte, error) {
	return c.doRequest(ctx, http.MethodDelete, path, nil)
}

func (c *Client) patch(ctx context.Context, path string, body interface{}) ([]byte, error) {
	return c.doRequest(ctx, http.MethodPatch, path, body)
}
