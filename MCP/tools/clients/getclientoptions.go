package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/openapi-generator-online/mcp-server/config"
	"github.com/openapi-generator-online/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetclientoptionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		languageVal, ok := args["language"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: language"), nil
		}
		language, ok := languageVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: language"), nil
		}
		url := fmt.Sprintf("%s/api/gen/clients/%s", cfg.BaseURL, language)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetclientoptionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_gen_clients_language",
		mcp.WithDescription("Returns options for a client library"),
		mcp.WithString("language", mcp.Required(), mcp.Description("The target language for the client library")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetclientoptionsHandler(cfg),
	}
}
