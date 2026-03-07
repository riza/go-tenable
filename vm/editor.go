package vm

import (
	"context"
	"encoding/json"
	"fmt"
)

// EditorService handles communication with the Editor related endpoints of the VM API.
type EditorService struct {
	client *Client
}

type EditorServiceDetailsResponseFilterAttributesItemControl struct {
	TypeField      string                   `json:"type,omitempty"`
	ReadableRegest string                   `json:"readable_regest,omitempty"`
	Regex          string                   `json:"regex,omitempty"`
	Options        []map[string]interface{} `json:"options,omitempty"`
}

type EditorServiceDetailsResponseFilterAttributesItem struct {
	Name         string                                                  `json:"name,omitempty"`
	ReadableName string                                                  `json:"readable_name,omitempty"`
	Operators    []map[string]interface{}                                `json:"operators,omitempty"`
	Control      EditorServiceDetailsResponseFilterAttributesItemControl `json:"control,omitempty"`
}

type EditorServiceDetailsResponseSettings struct {
	Basic      map[string]interface{} `json:"basic,omitempty"`
	Discovery  map[string]interface{} `json:"discovery,omitempty"`
	Assessment map[string]interface{} `json:"assessment,omitempty"`
	Advanced   map[string]interface{} `json:"advanced,omitempty"`
}

type EditorServiceDetailsResponseCredentials struct {
}

type EditorServiceDetailsResponseCompliance struct {
}

type EditorServiceDetailsResponsePlugins struct {
}

type EditorServiceDetailsResponse struct {
	IsWas            bool                                               `json:"is_was,omitempty"`
	UserPermissions  int                                                `json:"user_permissions,omitempty"`
	Owner            string                                             `json:"owner,omitempty"`
	Title            string                                             `json:"title,omitempty"`
	IsAgent          bool                                               `json:"is_agent,omitempty"`
	Uuid             string                                             `json:"uuid,omitempty"`
	FilterAttributes []EditorServiceDetailsResponseFilterAttributesItem `json:"filter_attributes,omitempty"`
	Settings         EditorServiceDetailsResponseSettings               `json:"settings,omitempty"`
	Credentials      EditorServiceDetailsResponseCredentials            `json:"credentials,omitempty"`
	Compliance       EditorServiceDetailsResponseCompliance             `json:"compliance,omitempty"`
	Plugins          EditorServiceDetailsResponsePlugins                `json:"plugins,omitempty"`
	Name             string                                             `json:"name,omitempty"`
}

type EditorServiceListTemplatesResponseItem struct {
	Unsupported      bool   `json:"unsupported,omitempty"`
	CloudOnly        bool   `json:"cloud_only,omitempty"`
	Desc             string `json:"desc,omitempty"`
	SubscriptionOnly bool   `json:"subscription_only,omitempty"`
	IsWas            bool   `json:"is_was,omitempty"`
	Title            string `json:"title,omitempty"`
	IsAgent          bool   `json:"is_agent,omitempty"`
	Uuid             string `json:"uuid,omitempty"`
	ManagerOnly      bool   `json:"manager_only,omitempty"`
	Name             string `json:"name,omitempty"`
}

type EditorServiceTemplateDetailsResponseFilterAttributesItemControl struct {
	TypeField      string                   `json:"type,omitempty"`
	ReadableRegest string                   `json:"readable_regest,omitempty"`
	Regex          string                   `json:"regex,omitempty"`
	Options        []map[string]interface{} `json:"options,omitempty"`
}

type EditorServiceTemplateDetailsResponseFilterAttributesItem struct {
	Name         string                                                          `json:"name,omitempty"`
	ReadableName string                                                          `json:"readable_name,omitempty"`
	Operators    []map[string]interface{}                                        `json:"operators,omitempty"`
	Control      EditorServiceTemplateDetailsResponseFilterAttributesItemControl `json:"control,omitempty"`
}

type EditorServiceTemplateDetailsResponseSettings struct {
	Basic      map[string]interface{} `json:"basic,omitempty"`
	Discovery  map[string]interface{} `json:"discovery,omitempty"`
	Assessment map[string]interface{} `json:"assessment,omitempty"`
	Advanced   map[string]interface{} `json:"advanced,omitempty"`
}

type EditorServiceTemplateDetailsResponseCredentials struct {
}

type EditorServiceTemplateDetailsResponseCompliance struct {
}

type EditorServiceTemplateDetailsResponsePlugins struct {
}

type EditorServiceTemplateDetailsResponse struct {
	IsWas            bool                                                       `json:"is_was,omitempty"`
	Title            string                                                     `json:"title,omitempty"`
	Name             string                                                     `json:"name,omitempty"`
	IsAgent          bool                                                       `json:"is_agent,omitempty"`
	FilterAttributes []EditorServiceTemplateDetailsResponseFilterAttributesItem `json:"filter_attributes,omitempty"`
	Settings         EditorServiceTemplateDetailsResponseSettings               `json:"settings,omitempty"`
	Credentials      EditorServiceTemplateDetailsResponseCredentials            `json:"credentials,omitempty"`
	Compliance       EditorServiceTemplateDetailsResponseCompliance             `json:"compliance,omitempty"`
	Plugins          EditorServiceTemplateDetailsResponsePlugins                `json:"plugins,omitempty"`
}

type EditorServicePluginDescriptionResponsePlugindescription struct {
	Severity         string                 `json:"severity,omitempty"`
	Pluginname       string                 `json:"pluginname,omitempty"`
	Pluginattributes map[string]interface{} `json:"pluginattributes,omitempty"`
	Pluginfamily     string                 `json:"pluginfamily,omitempty"`
	Pluginid         int                    `json:"pluginid,omitempty"`
}

type EditorServicePluginDescriptionResponse struct {
	Plugindescription EditorServicePluginDescriptionResponsePlugindescription `json:"plugindescription,omitempty"`
}

// Details - Get configuration details
func (s *EditorService) Details(ctx context.Context, typeParam string, id string) (*EditorServiceDetailsResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/editor/%s/%s", typeParam, id))
	if err != nil {
		return nil, err
	}
	var result EditorServiceDetailsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListTemplates - List templates
func (s *EditorService) ListTemplates(ctx context.Context, typeParam string) ([]EditorServiceListTemplatesResponseItem, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/editor/%s/templates", typeParam))
	if err != nil {
		return nil, err
	}
	var result []EditorServiceListTemplatesResponseItem
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// TemplateDetails - Get template details
func (s *EditorService) TemplateDetails(ctx context.Context, typeParam string, wizardUuid string) (*EditorServiceTemplateDetailsResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/editor/%s/templates/%s", typeParam, wizardUuid))
	if err != nil {
		return nil, err
	}
	var result EditorServiceTemplateDetailsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PluginDescription - Get plugin details
func (s *EditorService) PluginDescription(ctx context.Context, policyId string, familyId string, pluginId string) (*EditorServicePluginDescriptionResponse, error) {
	resp, err := s.client.get(ctx, fmt.Sprintf("/editor/policy/%s/families/%s/plugins/%s", policyId, familyId, pluginId))
	if err != nil {
		return nil, err
	}
	var result EditorServicePluginDescriptionResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Audits - Download audit file
func (s *EditorService) Audits(ctx context.Context, typeParam string, objectId string, fileId string) error {
	_, err := s.client.get(ctx, fmt.Sprintf("/editor/%s/%s/audits/%s", typeParam, objectId, fileId))
	if err != nil {
		return err
	}
	return nil
}
