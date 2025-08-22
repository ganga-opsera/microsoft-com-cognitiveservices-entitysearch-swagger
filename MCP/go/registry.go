package main

import (
	"github.com/entity-search-client/mcp-server/config"
	"github.com/entity-search-client/mcp-server/models"
	tools_entitysearch "github.com/entity-search-client/mcp-server/tools/entitysearch"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_entitysearch.CreateEntities_searchTool(cfg),
	}
}
