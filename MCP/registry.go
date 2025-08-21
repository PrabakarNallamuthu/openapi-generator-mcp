package main

import (
	"github.com/openapi-generator-online/mcp-server/config"
	"github.com/openapi-generator-online/mcp-server/models"
	tools_clients "github.com/openapi-generator-online/mcp-server/tools/clients"
	tools_servers "github.com/openapi-generator-online/mcp-server/tools/servers"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_clients.CreateClientoptionsTool(cfg),
		tools_clients.CreateGetclientoptionsTool(cfg),
		tools_clients.CreateGenerateclientTool(cfg),
		tools_servers.CreateServeroptionsTool(cfg),
		tools_servers.CreateGetserveroptionsTool(cfg),
		tools_servers.CreateGenerateserverforlanguageTool(cfg),
	}
}
