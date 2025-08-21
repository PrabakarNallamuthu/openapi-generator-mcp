package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// AuthorizationValue represents the AuthorizationValue schema from the OpenAPI specification
type AuthorizationValue struct {
	Urlmatcher PredicateOfURL `json:"urlMatcher,omitempty"`
	Value string `json:"value,omitempty"`
	Keyname string `json:"keyName,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// CliOption represents the CliOption schema from the OpenAPI specification
type CliOption struct {
	DefaultField string `json:"default,omitempty"`
	Description string `json:"description,omitempty"`
	Enum map[string]interface{} `json:"enum,omitempty"`
	Opt string `json:"opt,omitempty"`
	Optvalue string `json:"optValue,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// GeneratorInput represents the GeneratorInput schema from the OpenAPI specification
type GeneratorInput struct {
	Options map[string]interface{} `json:"options,omitempty"`
	Spec map[string]interface{} `json:"spec,omitempty"`
	Authorizationvalue AuthorizationValue `json:"authorizationValue,omitempty"`
	Openapiurl string `json:"openAPIUrl,omitempty"`
}

// PredicateOfURL represents the PredicateOfURL schema from the OpenAPI specification
type PredicateOfURL struct {
}

// ResponseCode represents the ResponseCode schema from the OpenAPI specification
type ResponseCode struct {
	Code string `json:"code,omitempty"` // File download code
	Link string `json:"link,omitempty"` // URL for fetching the generated client
}

// URLStreamHandler represents the URLStreamHandler schema from the OpenAPI specification
type URLStreamHandler struct {
}
