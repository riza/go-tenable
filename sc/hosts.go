package sc

import (
	"encoding/json"
	"fmt"
)

// HostsService handles communication with the hosts-related endpoints of the
// SC API.
type HostsService struct {
	client *Client
}

// Host represents a host resource from the API.
type Host struct {
	ID              string       `json:"id"`
	UUID            string       `json:"uuid"`
	TenableUUID     string       `json:"tenableUUID"`
	Name            string       `json:"name"`
	IPAddress       string       `json:"ipAddress"`
	OS              string       `json:"os"`
	FirstSeen       string       `json:"firstSeen"`
	LastSeen        string       `json:"lastSeen"`
	DNS             string       `json:"dns"`
	FQDNIndex       string       `json:"fqdnIndex"`
	NetBios         string       `json:"netBios"`
	NetBiosWorkgroup string      `json:"netBiosWorkgroup"`
	MacAddress      string       `json:"macAddress"`
	SystemType      string       `json:"systemType"`
	CreatedTime     string       `json:"createdTime"`
	ModifiedTime    string       `json:"modifiedTime"`
	Source          []HostSource `json:"source"`
	Repository      *IDRef       `json:"repository"`
	ACR             *ACRInfo     `json:"acr"`
}

// HostSource represents the source of a host record.
type HostSource struct {
	Type string `json:"type"`
}

// ACRInfo represents the Asset Criticality Rating information for a host.
type ACRInfo struct {
	HostUUID           string            `json:"hostUUID"`
	Score              string            `json:"score"`
	Overwritten        string            `json:"overwritten"`
	Notes              string            `json:"notes"`
	OverwrittenScore   string            `json:"overwrittenScore"`
	LastEditedUserID   string            `json:"lastEditedUserID"`
	LastEditedOrgID    string            `json:"lastEditedOrgID"`
	LastEvaluatedTime  string            `json:"lastEvaluatedTime"`
	InternetExposure   string            `json:"internetExposure"`
	Capability         string            `json:"capability"`
	DeviceType         string            `json:"deviceType"`
	Reasoning          []json.RawMessage `json:"reasoning"`
	KeyDrivers         json.RawMessage   `json:"keyDrivers"`
	User               *IDRef            `json:"user"`
}

// ACRUpdateInput contains the fields for updating a host's ACR.
type ACRUpdateInput struct {
	OverwrittenScore int         `json:"overwrittenScore"`
	Reasoning        []ACRReason `json:"reasoning,omitempty"`
	Notes            string      `json:"notes,omitempty"`
	Overwritten      string      `json:"overwritten,omitempty"`
}

// ACRReason represents a reason entry for an ACR override.
type ACRReason struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
}

// HostSearchInput contains the filters for searching hosts.
type HostSearchInput struct {
	Filters json.RawMessage `json:"filters"`
}

// HostSearchResponse holds the response from a host search.
type HostSearchResponse struct {
	TotalRecords    string            `json:"totalRecords"`
	ReturnedRecords int               `json:"returnedRecords"`
	StartOffset     string            `json:"startOffset"`
	EndOffset       string            `json:"endOffset"`
	Results         []json.RawMessage `json:"results"`
}

// HostDownloadInput contains the filters for downloading hosts.
type HostDownloadInput struct {
	Filters json.RawMessage `json:"filters,omitempty"`
}

// List returns all hosts.
func (s *HostsService) List() ([]Host, error) {
	resp, err := s.client.get("/hosts")
	if err != nil {
		return nil, fmt.Errorf("sc: list hosts: %w", err)
	}

	var result []Host
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal hosts list response: %w", err)
	}

	return result, nil
}

// UpdateACR updates the ACR for the host with the given UUID.
func (s *HostsService) UpdateACR(uuid string, input *ACRUpdateInput) (*Host, error) {
	resp, err := s.client.patch("/hosts/"+uuid+"/acr", input)
	if err != nil {
		return nil, fmt.Errorf("sc: update ACR for host %s: %w", uuid, err)
	}

	var host Host
	if err := json.Unmarshal(resp.Response, &host); err != nil {
		return nil, fmt.Errorf("sc: unmarshal host ACR response: %w", err)
	}

	return &host, nil
}

// Search searches for hosts matching the given filters.
func (s *HostsService) Search(input *HostSearchInput) (*HostSearchResponse, error) {
	resp, err := s.client.post("/hosts/search", input)
	if err != nil {
		return nil, fmt.Errorf("sc: search hosts: %w", err)
	}

	var result HostSearchResponse
	if err := json.Unmarshal(resp.Response, &result); err != nil {
		return nil, fmt.Errorf("sc: unmarshal hosts search response: %w", err)
	}

	return &result, nil
}

// Download requests a host data download. The response body from the API is
// binary (CSV), but since the client processes responses as JSON envelopes,
// the raw response bytes are returned.
func (s *HostsService) Download(input *HostDownloadInput) (json.RawMessage, error) {
	resp, err := s.client.post("/hosts/download", input)
	if err != nil {
		return nil, fmt.Errorf("sc: download hosts: %w", err)
	}

	return resp.Response, nil
}
